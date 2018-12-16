package single

import (
	"fmt"
	"os"
	"path/filepath"
)

// Start checks to see if a process lock file exists,
// if it does, it returns an error, otherwise, it creates the file.
func Start(name string) (err error) {
	f, err := os.OpenFile(filepath.Join(os.TempDir(), name), os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return
	}
	_, err = fmt.Fprintf(f, "%10d", os.Getpid())
	if err != nil {
		return
	}
	err = f.Close()
	return
}

// Stop remove the lock file
func Stop(name string) (err error) {
	err = os.Remove(filepath.Join(os.TempDir(), name))
	return
}
