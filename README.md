# single
[![Build Status](https://travis-ci.org/crgimenes/single.svg?branch=master)](https://travis-ci.org/crgimenes/single)
[![Go Report Card](https://goreportcard.com/badge/github.com/crgimenes/single)](https://goreportcard.com/report/github.com/crgimenes/single)
[![GoDoc](https://godoc.org/github.com/crgimenes/single?status.png)](https://godoc.org/github.com/crgimenes/single)
[![Go project version](https://badge.fury.io/go/github.com%2Fcrgimenes%2Fsingle.svg)](https://badge.fury.io/go/github.com%2Fcrgimenes%2Fsingle)
[![MIT Licensed](https://img.shields.io/badge/license-MIT-green.svg)](https://tldrlegal.com/license/mit-license)

Ensures that only one instance of the executable is running

## Install

```console
go get github.com/crgimenes/single
```

## Examples

```go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/crgimenes/single"
)

const lockfile = "signal.pid"

func main() {
	err := single.Start(lockfile)
	if err != nil {
		if os.IsExist(err) {
			fmt.Println("another instance of the executable is already running")
			os.Exit(1)
		}
		fmt.Println(err)
		os.Exit(1)
	}
	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
		<-sc
		errc := single.Stop(lockfile)
		if errc != nil {
			fmt.Println(errc)
			os.Exit(1)
		}
		fmt.Println("\n\nhave a nice day!")
		os.Exit(0)
	}()

	fmt.Println("Press ^C to terminate")

	for {
		time.Sleep(1 * time.Second)
		fmt.Print(".")
	}
}
```

---

For a process that does not run indefinitely, besides intercepting the signals use defer in the main function to ensure that street completion routine will always be executed

```go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/crgimenes/single"
)

const lockfile = "signal.pid"

func shutdown() {
	err := single.Stop(lockfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("\n\nhave a nice day!")
	os.Exit(0)
}

func main() {
	err := single.Start(lockfile)
	defer shutdown()
	if err != nil {
		if os.IsExist(err) {
			fmt.Println("another instance of the executable is already running")
			os.Exit(1)
		}
		fmt.Println(err)
		os.Exit(1)
	}
	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
		<-sc
		shutdown()
	}()

	fmt.Println("Processing...")
	time.Sleep(3 * time.Second)
}
```