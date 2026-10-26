package main

import (
	"bytes"
	"errors"
	"os"
	"strings"
	"testing"
)

func TestStable(t *testing.T) {
	in := "── cell 0 (line 4, ran in 610µs)\n" +
		"── cell 1 (line 15, kept)\n" +
		"   \"shape 2.441, scale 1.190\"\n" +
		"── cell 2 (line 21, ran in 1.2s)\n" +
		"   (line 3, ran in 5ms)\n" // a value that looks like a header is not one
	want := "── cell 0 (line 4)\n" +
		"── cell 1 (line 15)\n" +
		"   \"shape 2.441, scale 1.190\"\n" +
		"── cell 2 (line 21)\n" +
		"   (line 3, ran in 5ms)\n"
	if got := string(Stable([]byte(in))); got != want {
		t.Fatalf("Stable:\n%s\nwant:\n%s", got, want)
	}
}

// The committed output.txt is already stable: running it through again changes nothing.
func TestCommittedOutputIsStable(t *testing.T) {
	b, err := os.ReadFile("../output.txt")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(Stable(b), b) {
		t.Fatal("output.txt still has a timing or a kept marker")
	}
	if !bytes.Contains(b, []byte("── cell 1 (line 13)")) {
		t.Fatal("output.txt does not have the Ruby cell's header")
	}
}

func TestRun(t *testing.T) {
	var out bytes.Buffer
	if err := run(strings.NewReader("── cell 0 (line 4, kept)\n"), &out); err != nil {
		t.Fatal(err)
	}
	if out.String() != "── cell 0 (line 4)\n" {
		t.Fatalf("run wrote %q", out.String())
	}
}

type failReader struct{}

func (failReader) Read([]byte) (int, error) { return 0, errors.New("broken") }

type failWriter struct{}

func (failWriter) Write([]byte) (int, error) { return 0, errors.New("full") }

func TestRunErrors(t *testing.T) {
	if err := run(failReader{}, &bytes.Buffer{}); err == nil {
		t.Fatal("a read error was not reported")
	}
	if err := run(strings.NewReader("x"), failWriter{}); err == nil {
		t.Fatal("a write error was not reported")
	}
}

func TestMainExitsOnError(t *testing.T) {
	code := -1
	osExit = func(c int) { code = c }
	defer func() { osExit = os.Exit }()
	stdin := os.Stdin
	defer func() { os.Stdin = stdin }()
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	r.Close()
	w.Close()
	os.Stdin = r // reading a closed file fails
	main()
	if code != 1 {
		t.Fatalf("exit code %d, want 1", code)
	}
}
