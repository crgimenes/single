package single

import (
	"os"
	"path/filepath"
	"testing"
)

const filename = "testlock.pid"

func TestStartStop(t *testing.T) {
	defer func() {
		os.Remove(filepath.Join(os.TempDir(), filename))
	}()

	err := Start(filename)
	if err != nil {
		t.Errorf("expected no errors but got %q\n", err)
		return
	}

	err = Start(filename)
	if !os.IsExist(err) {
		t.Errorf("expected file exist error but got %q\n", err)
		return
	}

	err = Stop(filename)
	if err != nil {
		t.Errorf("expected no errors but got %q\n", err)
		return
	}

	err = Stop(filename)
	if !os.IsNotExist(err) {
		t.Errorf("expected expected not exist but got %q\n", err)
		return
	}
}
