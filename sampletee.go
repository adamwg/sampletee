package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	newline = []byte{'\n'}
)

func main() {
	var (
		n      = flag.Int("n", 10, "echo every nth line")
		append = flag.Bool("a", false, "append to output files rather than overwriting")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: sampletee [-n n] [-a] [files...]\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	openFlags := os.O_CREATE | os.O_WRONLY
	if *append {
		openFlags |= os.O_APPEND
	} else {
		openFlags |= os.O_TRUNC
	}
	ostreams := make([]io.Writer, flag.NArg())
	var err error
	for i, ofile := range flag.Args() {
		ostreams[i], err = os.OpenFile(ofile, openFlags, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open output file %q: %v\n", ofile, err)
			os.Exit(1)
		}
	}
	ostream := io.MultiWriter(ostreams...)

	istream := bufio.NewScanner(os.Stdin)
	for i := 0; istream.Scan(); i++ {
		line := istream.Bytes()
		if _, err := ostream.Write(line); err != nil {
			fmt.Fprintf(os.Stderr, "failed to write line to output file: %v\n", err)
			os.Exit(1)
		}
		if _, err := ostream.Write(newline); err != nil {
			fmt.Fprintf(os.Stderr, "failed to write line to output file: %v\n", err)
			os.Exit(1)
		}
		if i%*n == 0 {
			os.Stdout.Write(line)
			os.Stdout.Write(newline)
		}
	}
	if err := istream.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read from stdin: %v\n", err)
		os.Exit(1)
	}
}
