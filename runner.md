# Overview

The **GitBundle Runner** is an agent designed to execute **GitBundle Server** workflows. Written in Rust, it provides a lightweight, high-performance runtime for automating GitBundle tasks and reporting execution logs.

This release focuses on the core functionality of self-hosted runners. In upcoming releases, Docker-based builds will be supported for easier deployment and isolation.

# Features
- **Self-hosted**: Deploy the runner on your own infrastructure.
- **Workflow Execution**: Runs GitBundle Server workflows, supporting complex automation pipelines.
- **Log Reporting**: Captures and reports logs for workflow execution, making it easier to debug and monitor tasks.
- **Cross-platform**: Build and run the runner on Linux, Windows, and macOS with native performance.
- **Rust-based**: Implemented in Rust, ensuring memory safety, high concurrency, and minimal runtime overhead.

# Upcoming Enhancements
- **Docker Build Support**: Future versions will allow running the runner inside Docker containers, providing a reproducible environment across platforms.
- **Extended Workflow Integrations**: Enhanced support for GitBundle Server advanced workflows and triggers.
- **Modernize**: Transitioned to a pure Rust implementation, removing all Go dependencies.

# Getting Started
1. Download the latest release from [GitBundle Releases](https://github.com/gitbundle/gitbundle/releases)￼.
2. Install the runner on your preferred machine.
3. Follow the ui guide, register the runner with your GitBundle Server:
```bash
mkdir gitbundle-runner && cd gitbundle-runner
./runner register --server-url <SERVER_URL> --token <TOKEN> --labels <LABELS>
```
4. Start the runner:
```bash
./runner start
```
5. Verify logs and workflow execution in the GitBundle Server action pages.

# System Requirements
- Linux, macOS, or Windows (self-hosted)
- Optional Docker (for future releases)
- Git (for repository clone, push ...)

# Notes
- The current release does not support Docker deployment yet.
- Logs are reported asynchronously to the GitBundle Server. Ensure network connectivity for reliable logging.
- The runner is optimized for performance and concurrency, leveraging Rust’s async runtime.
- It is strongly recommended to deploy **GitBundle Server** and **GitBundle Runner** on separate machines for better experience.

# FAQ
Q: If I want to use a custom runner cache directory, how can I do it?

A: Export a global environment variable `XDG_CACHE_HOME` to specify the custom cache directory. The runner cache will be stored in the subdirectory named act. Users should not use a symbolic link for the cache directory, because the github action plugins may not work correctly. Instead, use a regular directory.
