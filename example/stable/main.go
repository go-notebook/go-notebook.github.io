// Command stable copies what `notebook run` printed from stdin to stdout
// without the parts that change from run to run: a cell's timing
// ("ran in 41µs") and the "kept" marker of a cached cell. What remains is
// what the landing shows and CI compares.
//
//	go tool notebook run waits.md | go run ./stable > output.txt
package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
)

// A cell header is "── cell 2 (line 19, ran in 30µs)" or
// "── cell 0 (line 4, kept)"; both become "── cell N (line L)".
var volatile = regexp.MustCompile(`(?m)^(── cell \d+ \(line \d+), (?:ran in [^)]*|kept)\)$`)

// Stable returns out without its timings and kept markers.
func Stable(out []byte) []byte {
	return volatile.ReplaceAll(out, []byte("$1)"))
}

// osExit is replaced by the tests.
var osExit = os.Exit

func run(in io.Reader, out io.Writer) error {
	b, err := io.ReadAll(in)
	if err != nil {
		return err
	}
	_, err = out.Write(Stable(b))
	return err
}

func main() {
	if err := run(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "stable:", err)
		osExit(1)
	}
}
