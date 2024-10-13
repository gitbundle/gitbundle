# this starlark script should be used to generate the build.yml
# configuration file.

import yaml

prodTrigger = {"ref": {"include": ["refs/tags/**", "refs/tags/**-beta", "refs/tags/**-nightly"], "exclude": ["refs/tags/**-dev"]}}
devTrigger = {"ref": ["refs/tags/**-dev"]}

def main():
    return prod(prodTrigger) + devel(devTrigger)


def devel(trigger):
    image = repository("%s/" % "bundle.lo", "gitbundle/bundle-docker")
    repo  = repository("%s/" % "bundle.lo", "bundle/magit-server")
    manifestImage = repository("%s/" % "bundle.lo", "gitbundle/bundle-manifest")
    stages = [
        build("linux", "arm64", "bundle-private", "private_edition", image, repo, "bundle.lo", "bundle.lo", trigger),
        build("linux", "amd64", "bundle-private", "private_edition", image, repo, "bundle.lo", "bundle.lo", trigger),
    ]
    after = manifest("bundle", "private_edition", manifestImage, "bundle.lo", trigger)

    for stage in stages:
        after["depends_on"].append(stage["name"])

    # artefact = binaries("private_edition", [
    #     binariesRelease("bundle", "bundle.lo/gitbundle/builds-gitbundle-release", "https://bundle.lo", "bundle", "magit-server")
    # ], trigger)
    # artefact["depends_on"].append(after["name"])

    return stages + [after] # + [artefact]


def prod(trigger):
    # TODO consider running unit tests before building and
    # publishing docker images.
    # before = testing("linux", "amd64")

    image = repository("", "gitbundle/bundle-docker")
    repo  = repository("", "gitbundle/gitbundle")
    manifestImage = repository("", "gitbundle/bundle-manifest")
    stages = [
        build("linux", "arm64", "gitbundle-private", "private_edition", image, repo, "docker.com", "", trigger),
        build("linux", "amd64", "gitbundle-private", "private_edition", image, repo, "docker.com", "", trigger),
    ]
    after = manifest("gitbundle-private", "private_edition", manifestImage, "docker.com", trigger)

    artefact = binaries("private_edition", [
        binariesRelease("github", "gitbundle/builds-github-release", "https://github.com", "gitbundle", "gitbundle"),
        binariesRelease("gitbundle", "gitbundle/builds-gitbundle-release", "https://gitbundle.com", "gitbundle", "gitbundle")
    ], trigger)
    artefact["depends_on"].append(after["name"])

    # the after stage should only execute after all previous
    # stages complete. this builds the dependency graph.
    for stage in stages:
        # stage["depends_on"].append(before["name"])
        after["depends_on"].append(stage["name"])

    image = repository("bundle.lo/", "gitbundle/bundle-docker")
    repo  = repository("gitbundle.com/", "bundle/magit-server")
    manifestImage = repository("bundle.lo/", "gitbundle/bundle-manifest")
    stageBeta = [
        build("linux", "arm64", "gitbundle-beta", "beta_edition", image, repo, "gitbundle.com", "gitbundle.com", trigger),
        build("linux", "amd64", "gitbundle-beta", "beta_edition", image, repo, "gitbundle.com", "gitbundle.com", trigger),
    ]
    afterBeta = manifest("gitbundle-beta", "beta_edition", manifestImage, "gitbundle.com", trigger)

    # artefactBeta = release("bundle", "beta_edition", "gitbundle", "gitbundle/builds-gitbundle-release", "https://gitbundle.com", "bundle", "magit-server")
    # artefactBeta["depends_on"].append(afterBeta["name"])

    for stage in stageBeta:
        # stage["depends_on"].append(before["name"])
        afterBeta["depends_on"].append(stage["name"])

    return stages + stageBeta + [after] + [afterBeta] + [artefact] # + [artefactBeta]


def repository(host, slug):
    return "%s%s" % (host, slug)


def testing(os, arch):
    return {
        "kind": "pipeline",
        "type": "docker",
        "name": "testing",
        "platform": {
            "os": os,
            "arch": arch,
        },
        "steps": [
            {
                "name": "vet",
                "pull": "always",
                "image": "golang:1.21-bullseye",
                "environment": {
                    "GO111MODULE": "on",
                    "GOPROXY": "https://goproxy.cn",
                },
                "commands": [
                    "go vet ./...",
                ],
                "volumes": [
                    {"name": "gopath", "path": "/go"},
                ],
            },
            {
                "name": "test",
                "pull": "always",
                "image": "golang:1.21-bullseye",
                "environment": {
                    "GO111MODULE": "on",
                    "GOPROXY": "https://goproxy.cn",
                },
                "commands": [
                    "go test -cover ./...",
                ],
                "volumes": [
                    {"name": "gopath", "path": "/go"},
                ],
            },
        ],
        "volumes": [{"name": "gopath", "temp": {}}],
        # "trigger": {"ref": ["refs/heads/master", "refs/tags/**", "refs/pull/**"]},
    }


def build(os, arch, org, type, image, repo, distHost, registry, trigger):
    return {
        "kind": "pipeline",
        "type": "docker",
        "name": "%s-%s-%s" % (org, os, arch),
        "platform": {
            "os": os,
            "arch": arch,
        },
        "volumes": [{"name": "deps", "temp": {}}],
        "steps": [
            {
                "name": "build",
                # "pull": "always",
                "image": "golang:1.21-bullseye",
                "commands": [
                    "sed -i 's/deb.debian.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list",
                    "apt-get update && apt-get -qqy install ca-certificates curl gnupg",
                    "mkdir -p /etc/apt/keyrings && curl -fsSL https://deb.nodesource.com/gpgkey/nodesource-repo.gpg.key | gpg --dearmor -o /etc/apt/keyrings/nodesource.gpg",
                    "echo \"deb [signed-by=/etc/apt/keyrings/nodesource.gpg] https://deb.nodesource.com/node_18.x nodistro main\" > /etc/apt/sources.list.d/nodesource.list",
                    "apt-get update && apt-get -qqy install nodejs",
                    # "curl -sL https://deb.nodesource.com/setup_18.x | bash - && apt-get -qqy install nodejs",
                    # "npm config set registry https://registry.npmmirror.com",
                    "npm config get registry",
                    "apt-get update && apt-get -qqy install python3-pip",
                    "pip3 config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple",
                    "pip3 install -U jc",
                    "bash hack.sh",
                    "make clean-all build",
                ],
                "environment": {
                    "CGO_ENABLED": 0,
                    "GO111MODULE": "on",
                    "GOPROXY": "https://goproxy.cn",
                    "GONOPROXY": "bundle.lo",
                    "GONOSUMDB": "bundle.lo",
                    "GOPRIVATE": "bundle.lo",
                    "TAGS": "bindata timetzdata %s" % type,
                    "GIT_SSH_KEY": {"from_secret": "git_ssh_key"},
                    "SSL_DEV_CERT": {"from_secret": "ssl_dev_cert"},
                },
                "volumes": [{"name": "deps", "path": "/go"}],
            },
            {
                "name": "executable-environment-to-ini",
                # 'pull': 'always',
                "image": "golang:1.21-bullseye",
                "commands": [
                    "./release/%s/%s/environment-to-ini --help" % (os, arch),
                ],
            },
            {
                "name": "executable-gitbundle",
                # 'pull': 'always',
                "image": "golang:1.21-bullseye",
                "commands": [
                    "./release/%s/%s/gitbundle --help" % (os, arch),
                ],
            },
            {
                "name": "dryrun",
                "pull": "always",
                "image": image,
                "settings": {
                    "daemon_off": False,
                    "dockerfile": "docker/Dockerfile.%s.%s" % (os, arch),
                    "dry_run": True,
                    "repo": repo,
                    "registry": registry,
                    "tags": "%s-%s" % (os, arch),
                    "username": {"from_secret": "%s_docker_username_sk" % distHost},
                    "password": {"from_secret": "%s_docker_password_sk" % distHost},
                },
            },
            {
                "name": "publish",
                "pull": "always",
                "image": image,
                "settings": {
                    "auto_tag": True,
                    "auto_tag_suffix": "%s-%s" % (os, arch),
                    "daemon_off": False,
                    "dockerfile": "docker/Dockerfile.%s.%s" % (os, arch),
                    "repo": repo,
                    "registry": registry,
                    "tags": "%s-%s" % (os, arch),
                    "username": {"from_secret": "%s_docker_username_sk" % distHost},
                    "password": {"from_secret": "%s_docker_password_sk" % distHost},
                },
                "when": {"event": {"exclude": ["pull_request"]}},
            },
        ],
        "depends_on": [],
        "trigger": trigger,
        # "trigger": {"event": ["tag"]},
    }


def manifest(org, type, image, distHost, trigger):
    return {
        "kind": "pipeline",
        "type": "docker",
        "name": "manifest-%s" % org,
        "platform": {
            "os": "linux",
            "arch": "amd64",
        },
        "steps": [
            {
                "name": "manifest",
                "image": image,
                "pull": "always",
                "settings": {
                    "dump": True,
                    "auto_tag": True,
                    # 'ignore_missing': True,
                    "spec": "docker/manifest.%s.tmpl" % org,
                    "username": {"from_secret": "%s_docker_username_sk" % distHost},
                    "password": {"from_secret": "%s_docker_password_sk" % distHost},
                },
            },
        ],
        "depends_on": [],
        "trigger": trigger,
    }


def binaries(type, releases, trigger):
    result = {
        "kind": "pipeline",
        "type": "docker",
        "name": "release",
        "platform": {
            "os": "linux",
            "arch": "amd64",
        },
        "steps": [
            {
                "name": "build",
                "pull": "always",
                "image": "techknowlogick/xgo:go-1.21.x",
                "commands": [
                    "sed -i 's/archive.ubuntu.com/mirrors.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list",
                    "sed -i 's/security.ubuntu.com/mirrors.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list",
                    "apt-get update && apt-get -qqy install ca-certificates curl gnupg",
                    "mkdir -p /etc/apt/keyrings && curl -fsSL https://deb.nodesource.com/gpgkey/nodesource-repo.gpg.key | gpg --dearmor -o /etc/apt/keyrings/nodesource.gpg",
                    "echo \"deb [signed-by=/etc/apt/keyrings/nodesource.gpg] https://deb.nodesource.com/node_18.x nodistro main\" > /etc/apt/sources.list.d/nodesource.list",
                    "apt-get update && apt-get -qqy install nodejs",
                    # "curl -sL https://deb.nodesource.com/setup_18.x | bash - && apt-get -qqy install nodejs",
                    # "npm config set registry https://registry.npmmirror.com",
                    "npm config get registry",
                    "apt-get update && apt-get -qqy install python3-pip",
                    "pip3 config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple",
                    "pip3 install -U jc",
                    "bash hack.sh",
                    "make clean-all release",
                ],
                "environment": {
                    "CGO_ENABLED": 0,
                    "GO111MODULE": "on",
                    "GOPROXY": "https://goproxy.cn",
                    "GONOPROXY": "bundle.lo",
                    "GONOSUMDB": "bundle.lo",
                    "GOPRIVATE": "bundle.lo",
                    "TAGS": "bindata timetzdata %s" % type,
                    "GIT_SSH_KEY": {"from_secret": "git_ssh_key"},
                    "SSL_DEV_CERT": {"from_secret": "ssl_dev_cert"},
                },
            },
            {
                "name": "gpg-sign",
                "image": "plugins/gpgsign",
                "pull": "always",
                "settings": {
                    "detach_sign": True,
                    "excludes": ["dist/release/*.sha256"],
                    "files": ["dist/release/*"],
                },
                "environment": {
                    "GPGSIGN_KEY": {"from_secret": "gpgsign_key"},
                    "GPGSIGN_PASSPHRASE": {"from_secret": "gpgsign_passphrase"},
                },
                "depends_on": ["build"]
            },
        ],
        "depends_on": [],
        "trigger": trigger,
    }

    result["steps"].extend(releases)

    return result


def binariesRelease(github, githubImage, baseUrl, repoowner, reponame):
    return {
        "name": github,
        "image": githubImage,
        "pull": "always",
        "settings": {
            "base_url": baseUrl,
            "files": ["dist/release/*"],
            "file_exists": "overwrite",
        },
        "environment": {
            "%s_TOKEN" % github.upper(): {"from_secret": "%s_token" % github},
            "BUNDLE_BUILDS_REPO_LINK": "%s/%s/%s" % (baseUrl, repoowner, reponame),
            "BUNDLE_BUILDS_REPO_OWNER": repoowner,
            "BUNDLE_BUILDS_REPO_NAME": reponame,
        },
        "depends_on": ["gpg-sign"]
    }


data = main()
with open("build.yml", "w") as outfile:
    outfile.write("# this file is automatically generated. DO NOT EDIT")
    outfile.write("\n")
    yaml.dump_all(data, outfile, default_flow_style=False, sort_keys=False)
