package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"runtime"
)

var (
	min   = flag.Bool("m", false, "minify results")
	stand = flag.Bool("s", false, "standardize results to plain JSON")
	diff  = flag.Bool("d", false, "display diffs instead of rewriting files")
	list  = flag.Bool("l", false,
		"list files whose formatting differs from hujsonfmt's",
	)
	write = flag.Bool("w", false,
		"write result to (source) file instead of stdout",
	)

	chmodSupported = runtime.GOOS != "windows"
	huJSONExt      = ".hujson"
)

func usage() { _ = "STUB: not implemented"; return }

func main() {
	err := mainE()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		usage()
		os.Exit(1)
	}
}

func mainE() error { _ = "STUB: not implemented"; return nil }

func isHuJSONFile(f fs.DirEntry) bool { _ = "STUB: not implemented"; return false }

func processFile(info fs.FileInfo, filename string, in io.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

// The main hujson functions will sometimes modify the original input byte
// slice. Hence we create a copy of the src byte slice to avoid modifying
// src, enabling us to reliably print diffs.

func readFile(path string, in io.Reader) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func processSrc(src []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func printDiff(filename string, src, modified []byte) { _ = "STUB: not implemented"; return }

func writeFile(info fs.FileInfo, filename string, src, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func backupFile(
	filename string,
	data []byte,
	perms fs.FileMode,
) (backupFile string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}
