package main

import (
	"context"
	"fmt"
	"io"
	"os"
)

func run(ctx context.Context, args []string, stdin io.Reader, stdout io.Writer, getwd func() (string, error)) error {
	fmt.Println("Hello, world!")
	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args, os.Stdin, os.Stdout, os.Getwd); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1) // comply with xxd
	}
}
