# Introduction to Go

This repo contains a demo that I prepared for my talk on GDG Melbourne:
https://plus.google.com/events/cahdv4rtr3gbunmbjobtarbsvb0

# Installation

```
$ go get github.com/dooman87/go-intro
```

# Running in Docker

```
$ docker build -t go-intro .
$ docker run  -it --rm -p 8080:8080 -v `pwd`:/go/src/github.com/dooman87/go-intro go-intro
```

# Description

All demos run from `cmd/main.go`.

Demos include:

```
* Package manager               - showing go get
* Hello World & Compilation     - main.go
* Static Types & Type Inference - types.go
* Slices                        - slices.go
* Maps                          - maps.go
* Structs                       - interfaces.go
* Interfaces                    - interfaces.go
* Closures                      - closures.go
* Go routines                   - goroutines.go
* Channels                      - channels.go
* Standard library              - http.go
* Tools                         - show go-vet, go-test, go-lint, go-replace, ...
```