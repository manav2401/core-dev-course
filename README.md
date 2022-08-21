# core-dev-course
This repository contains solutions for the core dev course. 

The original course repo with questions and examples can be found [here](https://github.com/maticnetwork/core_golang_course). 

### Installation

- Go 1.18
- make

### Running the solutions

Every week's solution is present in a sub-folder and can be run simply by `go run <exercise-name>.go`.

### Linters
Installing dependencies
```
make lintci-deps
```
Linting
```
make lint
```
The github CI performes linting on every commit. Enabled linters can be found in [./.golangci.yml](./.golangci.yml)
