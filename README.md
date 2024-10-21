# confirm

![Build status](https://github.com/romnn/confirm/actions/workflows/build.yaml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/romnn/confirm)](https://goreportcard.com/report/github.com/romnn/confirm)

A dead simple CLI app that asks the user for a yes/no confirmation.

```bash
go install github.com/romnn/confirm@main
```

#### Usage

```bash
# will exit with code 1 when user does not confirm
confirm "are you sure you want to proceed?"
```

## Development

To use the provided tasks in `taskfile.yaml`, install [task](https://taskfile.dev/).

```bash
# view all available tasks
task --list-all
# install development tools
task dependencies:install
```

After setup, you can use the following tasks during development:

```bash
task tidy
task run:race
task run:race -- validate -s ./schema.json -v ./values.json
task build:race
task test
task lint
task format
```
