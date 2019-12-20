package cmd

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/gsamokovarov/jump/config"
)

var td string

func tempConfig(t *testing.T) config.Config {
	conf, err := config.Temporary(td, ".tmp")
	if err != nil {
		t.Fatalf("Cannot setup temporary testing config: %v", err)
	}

	return conf
}

func capture(stream **os.File, fn func()) string {
	rescue := *stream
	r, w, _ := os.Pipe()

	*stream = w
	defer func() {
		*stream = rescue
	}()

	fn()

	w.Close()
	out, _ := ioutil.ReadAll(r)

	return string(out)
}

func inside(dir string, fn func()) {
	previousCwd, err := os.Getwd()
	if err != nil {
		return
	}

	if os.Chdir(dir) != nil {
		return
	}

	fn()

	os.Chdir(previousCwd)
}

func init() {
	cwd, _ := os.Getwd()
	td = path.Join(cwd, "testdata")
}
