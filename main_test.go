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
		args []string
	}

	tests := []struct {
		name       string
		args       args
		wantStdout string
		wantStderr string
		wantErr    bool
	}{
		{"no args no stdin", args{[]string{"zipcheck", ""}}, "emptyStdout", "emptyStderr", false},
		{
			"fake directory",
			args{[]string{"zipcheck", "aaaaaaa"}},
			"fakeStdout",
			"fakeStderr",
			false,
		},
		{
			"good directory",
			args{[]string{"zipcheck", "testdata"}},
			"goodStdout",
			"goodStderr",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			stderr := &bytes.Buffer{}

			wantStdoutB, err := os.ReadFile(testfile(tt.wantStdout))
			if err != nil {
				t.Errorf("Failed opening test file %s: %v", tt.wantStdout, err)
			}
			wantStdout := string(wantStdoutB)

			wantStderrB, err := os.ReadFile(testfile(tt.wantStderr))
			if err != nil {
				t.Errorf("Failed reading test file %s: %v", tt.wantStderr, err)
			}
			wantStderr := string(wantStderrB)

			err = run(tt.args.args, stdout, stderr)
			if (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStdout := stdout.String(); gotStdout != wantStdout {
				t.Errorf("run() gotStdout = %v, want %v", gotStdout, wantStdout)
			}
			if gotStderr := stderr.String(); gotStderr != wantStderr {
				t.Errorf("run() gotStderr = %v, want %v", gotStderr, wantStderr)
			}
		})
	}
}
