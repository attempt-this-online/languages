package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"sync"
	"time"
)

func getInputImages() []string {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lines
}

type result struct {
	ok        bool
	imageName string
}

func main() {
	dependencyGraph := make(map[string][]string)
	contexts := make(map[string]string)

	fromRegexp := regexp.MustCompile(`(?m)^\s*FROM\s+attemptthisonline/(\S+).*$`)
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		if filepath.Base(path) == "Dockerfile" {
			context := filepath.Dir(path)
			name := filepath.Base(context)
			if _, exists := contexts[name]; exists {
				panic("breadmaker: error: duplicate dockerfiles")
			}
			contexts[name] = context
			data, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}
			matches := fromRegexp.FindAllStringSubmatch(string(data), -1)
			for _, m := range matches {
				if len(m) > 1 {
					dependency := m[1]
					dependencyGraph[dependency] = append(dependencyGraph[dependency], name)
				}
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	var queue []string
	switch len(os.Args) {
	case 1:
		queue = getInputImages()
	case 2:
		if os.Args[1] == "--all" {
			queue = []string{"base"}
		} else {
			panic("unknown option")
		}
	default:
		panic("too many options")
	}

	targets := make(map[string]struct{})
	// breadth-first search
	for len(queue) > 0 {
		image := queue[0]
		targets[image] = struct{}{}
		queue = queue[1:]
		// if _, ok := contexts[image]; !ok {
		//   panic("unknown image: " + image)
		// }
		for _, dependent := range dependencyGraph[image] {
			if _, exists := targets[dependent]; !exists {
				queue = append(queue, dependent)
			}
		}
	}

	waitGroups := make(map[string]*sync.WaitGroup)
	for image := range targets {
		var wg sync.WaitGroup
		waitGroups[image] = &wg
	}

	for image := range targets {
		for _, dependent := range dependencyGraph[image] {
			waitGroups[dependent].Add(1)
		}
	}

	now := time.Now().UTC().Format("2006-01-02-15-04-05")
	resultsChan := make(chan result)
	for image := range targets {
		go func(image string) {
			waitGroups[image].Wait()

			build(image, contexts[image], now, resultsChan)

			for _, dependent := range dependencyGraph[image] {
				waitGroups[dependent].Done()
			}
		}(image)
	}
	fmt.Println("starting build at", now)

	buildStatuses := make(map[string]bool)
	numFailed := 0
	numDone := 0
	for numDone < len(targets) {
		result := <-resultsChan
		buildStatuses[result.imageName] = result.ok
		if !result.ok {
			numFailed += 1
		}
		numDone += 1
		fmt.Printf("%s done (%v/%v)\n", result.imageName, numDone, len(targets))
	}

	printLogs(buildStatuses)
	fmt.Printf(RED+"%v failures"+RESET+"\n", numFailed)
	if numFailed != 0 {
		os.Exit(1)
	}
}

const RED = "\x1B[91m"
const RESET = "\x1B[0m"

func printLogs(buildStatuses map[string]bool) {
	// syntax from https://docs.gitlab.com/ee/ci/jobs/#expand-and-collapse-job-log-sections
	for name, succeeded := range buildStatuses {
		unixTimestamp := strconv.FormatInt(time.Now().UTC().Unix(), 10)
		if !succeeded {
			// red text
			fmt.Print(RED)
		}
		fmt.Println("\x1B[0Ksection_start:" + unixTimestamp + ":" + name + "[collapsed=true]\r\x1B[0KBuild " + name)
		if !succeeded {
			fmt.Print(RESET)
		}
		printFile(name)
		fmt.Println("\x1B[0Ksection_end:" + unixTimestamp + ":" + name + "\r\x1B[0K")
	}
}

func printFile(name string) {
	file, err := os.Open("output/" + name + ".log")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		panic(err)
	}
}

func build(
	name string,
	context string,
	now string,
	resultChan chan result,
) {
	tagBase1 := "registry.gitlab.pxeger.com/attempt-this-online/languages/" + name + ":"
	tagBase2 := "attemptthisonline/" + name + ":"
	cmd := exec.Command(
		"podman",
		"build",
		"--no-cache",
		"-t",
		tagBase1+now,
		"-t",
		tagBase1+"latest",
		"-t",
		tagBase2+now,
		"-t",
		tagBase2+"latest",
		context,
	)
	waitForOutputLoggers := logOutput("output/"+name+".log", cmd)
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	waitForOutputLoggers()
	err = cmd.Wait()
	if err == nil {
		resultChan <- result{true, name}
	} else if _, is := err.(*exec.ExitError); is {
		resultChan <- result{false, name}
	} else {
		panic(err)
	}
}

func logOutput(filename string, cmd *exec.Cmd) func() {
	var wg sync.WaitGroup

	linesChan := make(chan string)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	wg.Add(1)
	go readLines(stdout, linesChan, &wg)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}
	wg.Add(1)
	go readLines(stderr, linesChan, &wg)

	output, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	go func() {
		defer output.Close()
		for s := range linesChan {
			_, err := output.WriteString(s)
			if err != nil {
				panic(err)
			}
		}
	}()

	return func() {
		wg.Wait()
		close(linesChan)
	}
}

func readLines(pipe io.Reader, c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	reader := bufio.NewReader(pipe)
	for {
		line, err := reader.ReadString('\n')
		if err == nil || line != "" {
			c <- time.Now().UTC().Format("15-04-05 ") + line
		}
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
}
