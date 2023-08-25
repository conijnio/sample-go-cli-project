# Golang Template

![Go Version](https://img.shields.io/github/go-mod/go-version/conijnio/golang-template)
[![License](https://img.shields.io/badge/License-Apache2-green.svg)](./LICENSE)
[![Maintenance](https://img.shields.io/badge/Maintained-yes-green.svg)](https://github.com/conijnio/golang-template/graphs/commit-activity)
[![Workflow: go](https://github.com/conijnio/golang-template/actions/workflows/go.yml/badge.svg)](https://github.com/conijnio/golang-template/actions/workflows/go.yml)
[![Workflow: goreleaser](https://github.com/conijnio/golang-template/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/conijnio/golang-template/actions/workflows/goreleaser.yml)
![Release](https://img.shields.io/github/v/release/conijnio/golang-template)
[![Go Report Card](https://goreportcard.com/badge/github.com/conijnio/golang-template)](https://goreportcard.com/report/github.com/conijnio/golang-template)


Template repository for Golang projects.

## Prerequisites

You will need to install the following tools to successfully run the make targets: 

```shell
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
go install github.com/uudashr/gocognit/cmd/gocognit@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
go install github.com/go-critic/go-critic/cmd/gocritic@latest
```

## Commands

- `make build`, builds the project.
- `make complexity`, perform complexity scans on the codebase.
- `make coverage`, create and displays the code coverage report in HTML.
- `make help`, displays all the available options.   
- `make lint`, performs linting actions on the codebase.
- `make test`, runs all the unit tests.

### Run goreleaser locally

Because we enabled signing you need to supply a `GPG_FINGERPRINT` environment variable. You will be prompted for a passphrase. 

```shell
GPG_FINGERPRINT=C490C64E6938FD0C goreleaser release --snapshot --clean
```

Afterward, you can validate the signature using the following command:

```shell
gpg --verify [Signature] [File]
```

## License

This project is free and open source software licensed under the [Apache 2.0 License](./LICENSE).