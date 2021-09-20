package single_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/crgimenes/single"
)

const filename = "testlock.pid"

func TestStartStop(t *testing.T) {
	defer func() {
		os.Remove(filepath.Join(os.TempDir(), filename))
	}()

	err := single.Start(filename)
	if err != nil {
		t.Errorf("expected no errors but got %q\n", err)
		return
	}

	err = single.Start(filename)
	if !os.IsExist(err) {
		t.Errorf("expected file exist error but got %q\n", err)
		return
	}

	err = single.Stop(filename)
	if err != nil {
		t.Errorf("expected no errors but got %q\n", err)
		return
	}

	err = single.Stop(filename)
	if !os.IsNotExist(err) {
		t.Errorf("expected expected not exist but got %q\n", err)
		return
	}
}
