package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_handleFile(t *testing.T) {

	// determine input files
	match, err := filepath.Glob("testdata/*.yaml")

	if err != nil {
		t.Fatal(err)
	}

	for _, in := range match {
		out := in + ".golden"
		runTest(t, in, out)
	}
}

func runTest(t *testing.T, in, out string) {
	t.Run(in, func(t *testing.T) {

		f := filepath.Base(in)

		dir, err := ioutil.TempDir(os.TempDir(), f+"-")
		if err != nil {
			log.Fatal(err)
		}
		defer os.RemoveAll(dir)

		handleFile(in, dir)

		wantFiles, err := filepath.Glob(filepath.Join(out, "*"))
		gotFiles, err := filepath.Glob(filepath.Join(dir, "*"))

		if len(gotFiles) != len(wantFiles) {
			t.Errorf("handleFile() = %v, want %v", len(gotFiles), len(wantFiles))
		}

		for _, wantFile := range wantFiles {
			want, _ := ioutil.ReadFile(wantFile)
			got, _ := ioutil.ReadFile(filepath.Join(dir, filepath.Base(wantFile)))

			if !reflect.DeepEqual(got, want) {
				t.Errorf(cmp.Diff(got, want))
			}
		}
	})
}
