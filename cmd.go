package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
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

	m := parseMods(graph)
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

func parseMods(s string) *Mod {
	lines := strings.Split(s, "\n")
	if len(lines) == 1 && len(lines[0]) == 0 {
		return nil
	}
	mod := NewMod(lines[0][:strings.Index(lines[0], " ")])
	for i := 0; i < len(lines); i++ {
		m, d, ok := strings.Cut(lines[i], " ")
		if ok {
			mod.AddDependency(m, d)
		}
	}
	return mod
}
