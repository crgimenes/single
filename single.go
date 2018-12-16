package single

import (
	"fmt"
	"os"
	"path/filepath"
)

func Start(name string) (err error) {
	f, err := os.OpenFile(filepath.Join(os.TempDir(), name), os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return
	}
	fmt.Fprintf(f, "%10d", os.Getpid())
	err = f.Close()
	return
}

func Stop(name string) (err error) {
	err = os.Remove(filepath.Join(os.TempDir(), name))
	return
}
