# GitBundle

**GitBundle is a modern, Rust-powered code hosting and workflow automation platform built for performance, security, and full self-hosting.**
It provides a streamlined workflow engine, an efficient Rust backend, and an architecture designed for teams that need reliability, control, and speed.

[demo](https://demo.gitbundle.com)

## Key Highlights

- ⚡ High-performance core powered by Rust with async I/O
- 🔒 Reliable and memory-safe architecture
- 🔄 Workflow engine compatible with GitHub Actions syntax
- 🧩 Extensible by design for custom CI/CD pipelines and integrations
- 🌐 Multi-language API support for automation and tooling
- 🏢 Ideal for teams and enterprises requiring a secure, self-hosted platform

---

## Overview

The **GitBundle Server** uses release tags such as `server-v3.1.0` for stable builds and `server-v3.1.0-beta` for pre-release builds. Beta releases indicate that the build is feature-complete and near production-ready.

The **GitBundle Runner** follows the same convention, using tags like `runner-v1.0.0` for stable releases and `runner-v1.0.0-beta` for releases that are close to production use.

## Running GitBundle Server

GitBundle Server loads configuration from the `.env` file in the working directory.

requirements:

- git
- [gitleaks](https://github.com/gitleaks/gitleaks)

```bash
# Copy the template and start the server
cp .env.slim .env
gitbundle server
```

## Running GitBundle Server in Container (**recommended**)

Using containerized deployment ensures isolation, consistency, and strong performance.

```bash
docker pull ghcr.io/gitbundle/server:v3-beta
# or
docker pull gitbundle/server:v3-beta
```

## Running the GitBundle Runner

The GitBundle Runner is responsible for executing workflows and reporting logs to the server.

requirements:

- git

```bash
mkdir ~/.gitbundle-runner
cd ~/.gitbundle-runner

# Follow the UI instructions for registration
runner register --server-url <SERVER_URL> --token <TOKEN>

# Start the runner
runner start
```

## Operational Recommendation

For stability and performance, deploy GitBundle Server and GitBundle Runner on separate machines.
Running both on the same host may cause resource contention during workflow execution.

## Github Workflow Syntax Compatibility

GitBundle provides broad compatibility with GitHub Actions workflow syntax, including:

  - Standard GitHub Actions YAML structure
  - Jobs, steps, runs-on, needs, and conditionals
  - Matrix expansion
  - Expressions using GitHub’s expression language
  - All official GitHub contexts and variables
    - github
    - env
    - vars
    - job
    - jobs
    - steps
    - runner (only runner.arch, runner.os, runner.environment are currently supported)
    - secrets
    - strategy
    - matrix
    - needs
    - inputs

## Supported Github Event Types & Operation

- **branch_protection_rule**
  - created
  - edited
  - deleted

- **check_run** (not ready yet)
  - created
  - edited
  - deleted

- **check_suite** (not ready yet)

- **branch**
  - created
  - renamed
  - updated
  - deleted
  
- **tag**
  - created
  - updated
  - deleted

- **pull_request**
  - synchronize
  - commented
  - reviewed
  - assigned
  - opened
  - closed
  - reopened
  - merged
  - edited
  - review_requested
  - review_commented
  - labeled (not ready yet)
  - locked (not ready yet)

- **release**
  - created
  - deleted
  - edited
  - published
  - prereleased

- **repository**
  - renamed
  - imported
  - created
  - deleted
  - forked
  - default_branch_updated
  - restored
  - public
  - watch_started (not ready yet)
  - transferred (not ready yet)
  - dispatch
	
- **schedule** (not ready yet)

- **status** (not ready yet)

- **workflow_call**

- **workflow_dispatch**

- **workflow_run**
  - completed
  - requested
  - in_progress
  
## Gitlab Workflow Syntax Compatibility (WIP)

# FAQ

For issues or inquiries, please report through the official issue tracker:
https://github.com/gitbundle/gitbundle/issues

Thank you for your support.
