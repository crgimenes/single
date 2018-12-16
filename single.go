package single

import (
	"fmt"
	"os"
)

func Start(name string) (err error) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return
	}
	fmt.Fprintf(f, "%10d", os.Getpid())
	f.Close()
}

func Stop(name string) {
	os.Remove(name)
}
