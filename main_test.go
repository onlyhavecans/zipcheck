package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

const (
	testdata = "testdata"
)

func testfile(name string) string {
	return fmt.Sprintf("%s/%s", testdata, name)
}

func TestRun(t *testing.T) {
	type args struct {
		stdin string
		args  []string
	}

	tests := []struct {
		name       string
		args       args
		wantStdout string
		wantErr    bool
	}{
		{"no args no stdin", args{"empty", []string{"zipcheck", ""}}, "emptyStdout", false},
		{"fake directory", args{"empty", []string{"zipcheck", "aaaaaaa"}}, "fakeStdout", false},
		{"good directory", args{"empty", []string{"zipcheck", "testdata"}}, "goodStdout", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			stdinB, err := os.ReadFile(testfile(tt.args.stdin))
			if err != nil {
				t.Errorf("Failed reading test file %s: %v", tt.args.stdin, err)
			}
			stdin := bytes.NewReader(stdinB)

			wantStdoutB, err := os.ReadFile(testfile(tt.wantStdout))
			if err != nil {
				t.Errorf("Failed opening test file %s: %v", tt.wantStdout, err)
			}
			wantStdout := string(wantStdoutB)

			err = run(tt.args.args, stdin, stdout)
			if (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStdout := stdout.String(); gotStdout != wantStdout {
				t.Errorf("run() gotStdout = %v, want %v", gotStdout, wantStdout)
			}
		})
	}
}
