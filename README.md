# GitBundle

[![Discord](https://img.shields.io/badge/Discord-5865F2?logo=discord&logoColor=white)](https://discord.gg/8rF95ZVh9)

**GitBundle is an out-of-the-box, zero-config, modern, Rust-powered code hosting and workflow automation platform built for performance, security, and full self-hosting.**
It provides a streamlined workflow engine, an efficient Rust backend, and an architecture designed for teams that need reliability, control, and speed.

<div align="center">
  <a href="https://www.youtube.com/watch?v=Z2UU4QM2C0o">
    <img src="https://img.youtube.com/vi/Z2UU4QM2C0o/maxresdefault.jpg" alt="GitBundle Demo" width="100%" />
  </a>
</div>

> **❤️ A Note from the Maintainer**
>
> Maintaining this project alone has taken far more effort than I ever imagined, and the cost has been huge.
> If GitBundle has helped you, please **star ⭐** it, this means a lot. 
> And if you'd recommend it to your friends and colleagues, I'd be deeply grateful. Thank you.

## Key Highlights

- ⚡ High-performance core powered by Rust with async I/O
- 🔒 Reliable and memory-safe architecture
- 🔄 Workflow engine compatible with GitHub Actions syntax
- 🧩 Extensible by design for custom CI/CD pipelines and integrations
- 🌐 Multi-language API support for automation and tooling
- 🏢 Ideal for teams and enterprises requiring a secure, self-hosted platform

## Screenshots

### Light Theme

<table>
  <tr>
    <td><img src="https://raw.githubusercontent.com/gitbundle/gitbundle/main/screenshot/light-repo-home.png" alt="Repository" width="100%" /></td>
    <td><img src="https://raw.githubusercontent.com/gitbundle/gitbundle/main/screenshot/light-workflow-detail.png" alt="Workflow" width="100%" /></td>
    <td><img src="https://raw.githubusercontent.com/gitbundle/gitbundle/main/screenshot/light-action-home.png" alt="Actions" width="100%" /></td>
    <td><img src="https://raw.githubusercontent.com/gitbundle/gitbundle/main/screenshot/light-admin-runners.png" alt="Runners" width="100%" /></td>
  </tr>
  <tr>
    <td><img src="https://raw.githubusercontent.com/gitbundle/gitbundle/main/screenshot/light-group-home.png" alt="Group" width="100%" /></td>
    <td><img src="https://raw.githubusercontent.com/gitbundle/gitbundle/main/screenshot/light-admin-groups.png" alt="Admin Groups" width="100%" /></td>
    <td><img src="https://raw.githubusercontent.com/gitbundle/gitbundle/main/screenshot/light-dashboard.png" alt="Dashboard" width="100%" /></td>
    <td><img src="https://raw.githubusercontent.com/gitbundle/gitbundle/main/screenshot/light-person-profile.png" alt="Profile" width="100%" /></td>
  </tr>
</table>

<details>
<summary><b>Dark Theme</b></summary>

<table>
  <tr>
    <td><img src="https://raw.githubusercontent.com/gitbundle/gitbundle/main/screenshot/dark-repo-home.png" alt="Repository" width="100%" /></td>
    <td><img src="https://raw.githubusercontent.com/gitbundle/gitbundle/main/screenshot/dark-workflow-detail.png" alt="Workflow" width="100%" /></td>
    <td><img src="https://raw.githubusercontent.com/gitbundle/gitbundle/main/screenshot/dark-action-home.png" alt="Actions" width="100%" /></td>
    <td><img src="https://raw.githubusercontent.com/gitbundle/gitbundle/main/screenshot/dark-admin-runners.png" alt="Runners" width="100%" /></td>
  </tr>
  <tr>
    <td><img src="https://raw.githubusercontent.com/gitbundle/gitbundle/main/screenshot/dark-group-home.png" alt="Group" width="100%" /></td>
    <td><img src="https://raw.githubusercontent.com/gitbundle/gitbundle/main/screenshot/dark-admin-groups.png" alt="Admin Groups" width="100%" /></td>
    <td><img src="https://raw.githubusercontent.com/gitbundle/gitbundle/main/screenshot/dark-person-profile.png" alt="Profile" width="100%" /></td>
    <td></td>
  </tr>
</table>

</details>

---

## Overview

The **GitBundle Server** uses release tags such as `server-v3.6.6` for stable builds and `server-v3.6.6-beta` for pre-release builds. Beta releases indicate that the build is feature-complete and near production-ready.

The **GitBundle Runner** follows the same convention, using tags like `runner-v1.2.3` for stable releases and `runner-v1.2.3-beta` for releases that are close to production use.

### Binary and Image Release Versioning

GitBundle binaries and container images follow the **Semantic Versioning** convention: `MAJOR.MINOR.PATCH`.

- **MAJOR** – Incremented when there are incompatible breaking changes.
- **MINOR** – Incremented when new functionality is added in a backward-compatible manner.
- **PATCH** – Incremented when backward-compatible bug fixes or minor improvements are made.

When a defect is discovered in a published artifact, the fix will be applied and the artifact will be re-released with an incremented `PATCH` version (e.g., from `v1.2.3` to `v1.2.4`).

## Requirements

For both GitBundle Server and GitBundle Runner:

- git
- [gitleaks](https://github.com/gitleaks/gitleaks) @v8.30.0+
- nodejs @18.16.0+
- docker @28.5.2+

## Running GitBundle Server

GitBundle Server loads configuration from the `.env` file in the working directory.

```bash
# Copy the template and start the server
cp .env.slim .env
gitbundle server
```

## Running GitBundle Server in Container (**recommended**)

Using containerized deployment ensures isolation, consistency, and strong performance.

```bash
docker pull ghcr.io/gitbundle/server:v3
# or
docker pull gitbundle/server:v3
```

## Running the GitBundle Runner

The GitBundle Runner is responsible for executing workflows and reporting logs to the server.

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
    - runner
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

- **create** (extended)
  - **branch**
    - created
    
  - **tag**
    - created
  
  ```yaml
  on:
    delete:
      branches:
        - beta/*
      branches-ignore:
        - alpha/*
      tags:
        - v*-beta
      tags-ignore:
        - v*-alpha
      paths:
        - 'src/**/*.rs'
      paths-ignore:
        - 'alpha/**/*.rs'
  ```

- **delete** (extended, refer to create event)
  - **branch**
    - deleted
    
  - **tag**
    - deleted

- **push**
  - **branch**
    - created
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

## License

[![License](https://img.shields.io/badge/license-Elastic%202.0-blue)](LICENSE)
  
# FAQ

For issues or inquiries, please report through the official issue tracker:
https://github.com/gitbundle/gitbundle/issues

Thank you for your support.
