import glob
import json
import os.path
import re

result = {}

for dockerfile in glob.iglob("./**/Dockerfile", recursive=True):
    context = os.path.dirname(dockerfile)
    name = os.path.basename(context)
    assert re.fullmatch(r"[\w-]+", name), f"unsafe name {name!r}"
    with open(dockerfile) as f:
        dependencies = [m.group(1) for m in re.finditer(r"^\s*FROM\s+attemptthisonline/(\S+).*$", f.read(), re.M)]
    tags = [
        tag_base + name + ":" + tag
        for tag_base in (
            "registry.gitlab.pxeger.com/attempt-this-online/languages/",
            "docker.io/attemptthisonline/",
        )
        for tag in ("$now", "latest")
    ]
    result[f"build {name}"] = {
        "extends": ".base",  # `.build` template is defined in main YAML file
        "needs": [f"build {dep}" for dep in dependencies] or ["init"],
        "script": [
            "now=$(echo $CI_PIPELINE_CREATED_AT | tr T: - | head -c 19)",
            "podman build --no-cache " + context + "".join(" -t " + t for t in tags),
            "podman push " + " ".join(tags),
        ],
    }

# note YAML 1.2 is backwards-compatible with JSON
with open(".images.gitlab-ci.yml", "w") as f:
    json.dump(result, f)
