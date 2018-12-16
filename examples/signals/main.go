package main

import (
	"fmt"
	"os"
	"os/signal"
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
		signal.Notify(sc, os.Interrupt)
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
