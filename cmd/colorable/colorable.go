package main

import (
	"io"
	"os"

	"github.com/pschlump/go-colorable"
)

func main() {
	io.Copy(colorable.NewColorableStdout(), os.Stdin)
}
