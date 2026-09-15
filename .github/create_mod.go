// Command create_mod packs a directory into a Go module zip.
//
// Usage: create_mod <module-path> <version> <dir> <out.zip>
//
// It exists only to reproduce the module zip that the module proxy serves for a
// tagged release; it is not part of golang.org/x/text and is removed from the
// tree before the zip is created.
package main

import (
	"fmt"
	"os"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintln(os.Stderr, "usage: create_mod <module-path> <version> <dir> <out.zip>")
		os.Exit(2)
	}

	mv := module.Version{Path: os.Args[1], Version: os.Args[2]}

	f, err := os.Create(os.Args[4])
	if err != nil {
		fmt.Fprintln(os.Stderr, "create:", err)
		os.Exit(1)
	}

	if err := zip.CreateFromDir(f, mv, os.Args[3]); err != nil {
		fmt.Fprintln(os.Stderr, "CreateFromDir:", err)
		os.Exit(1)
	}

	if err := f.Close(); err != nil {
		fmt.Fprintln(os.Stderr, "close:", err)
		os.Exit(1)
	}

	fmt.Printf("wrote %s for %s@%s\n", os.Args[4], mv.Path, mv.Version)
}
