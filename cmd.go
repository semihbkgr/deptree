package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
)

var (
	depth = flag.Int("d", 5, "depth of dependency tree")
)

func main() {
	flag.Parse()

	var graph string
	if len(flag.Args()) > 0 && flag.Args()[0] == "-" {
		b, err := io.ReadAll(os.Stdin)
		checkErr(err)
		graph = string(b)
	} else {
		s, err := execGoModGraphCommand()
		checkErr(err)
		graph = s
	}

	m := ParseMods(graph)
	if m != nil {
		fmt.Println(m.Tree(*depth, false))
	} else {
		fmt.Println("module has no dependency")
	}
}

func execGoModGraphCommand() (string, error) {
	c := exec.Command("go", "mod", "graph")
	b, err := c.CombinedOutput()
	if err != nil {
		return "", errors.New(string(b))
	}
	return string(b), nil
}

func checkErr(err error) {
	if err != nil {
		_, err = fmt.Fprintln(os.Stderr, err.Error())
		if err != nil {
			panic(err)
		}
		os.Exit(1)
	}
}
