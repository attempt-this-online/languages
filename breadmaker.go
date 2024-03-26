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
	"strings"
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

type options struct {
	all            bool
	dockerOrPodman string
	extraOptions   []string
	tagBases       []string
}

func printUsage() {
	fmt.Println("usage: breadmaker (--podman | --docker) [--all] [-O=<builder option> | -t=<tag base>] [...]")
}

func parseOptions() options {
	var options options
	for _, arg := range os.Args[1:] {
		switch arg {
		case "--all":
			options.all = true
		case "--podman":
			options.dockerOrPodman = "podman"
		case "--docker":
			options.dockerOrPodman = "docker"
		default:
			if strings.HasPrefix(arg, "-O=") {
				options.extraOptions = append(options.extraOptions, arg[3:])
			} else if strings.HasPrefix(arg, "-t=") {
				options.tagBases = append(options.tagBases, arg[3:])
			} else {
				printUsage()
				panic("unknown option: " + arg)
			}
		}
	}
	if options.dockerOrPodman == "" {
		printUsage()
		panic("You must provide either --docker or --podman")
	}
	return options
}

func main() {
	options := parseOptions()

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
	if options.all {
		queue = []string{"base"}
	} else {
		queue = getInputImages()
	}

	targets := make(map[string]struct{})
	// breadth-first search
	for len(queue) > 0 {
		image := queue[0]
		targets[image] = struct{}{}
		queue = queue[1:]
		if _, ok := contexts[image]; !ok {
			panic("unknown image: " + image)
		}
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

			build(image, contexts[image], options, now, resultsChan)

			for _, dependent := range dependencyGraph[image] {
				waitGroups[dependent].Done()
			}
		}(image)
	}
	fmt.Println("starting build at", now)

	numFailed := 0
	numDone := 0
	for numDone < len(targets) {
		result := <-resultsChan
		numDone += 1
		var statusString string
		if result.ok {
			statusString = "succeeded"
		} else {
			numFailed += 1
			statusString = RED + "failed" + RESET
		}
		fmt.Printf("%v %v (%v/%v)\n", result.imageName, statusString, numDone, len(targets))
	}

	if numFailed != 0 {
		var plural string
		if numFailed != 1 {
			plural = "s"
		}
		fmt.Printf("%v%v failure%v%v\n", RED, numFailed, plural, RESET)
		os.Exit(1)
	} else {
		fmt.Println("0 failures")
	}
}

const RED = "\x1B[91m"
const RESET = "\x1B[0m"

func build(
	name string,
	context string,
	options options,
	now string,
	resultChan chan result,
) {
	combinedOptions := append([]string{
		"build",
		context,
	}, options.extraOptions...)
	for _, tagBase := range options.tagBases {
		for _, version := range []string{now, "latest"} {
			combinedOptions = append(combinedOptions, "-t", tagBase+name+":"+version)
		}
	}
	cmd := exec.Command(options.dockerOrPodman, combinedOptions...)
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
