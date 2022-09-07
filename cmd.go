package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	depth = flag.Int("d", 5, "depth of dependency tree")
)

func main() {
	flag.Parse()
	s, err := execGoModGraphCommand()
	checkErr(err)
	m, err := parseMods(s)
	checkErr(err)
	if m != nil {
		fmt.Println(m.Tree(*depth, false))
	} else {
		fmt.Println("module has no dependency")
	}
}

func execGoModGraphCommand() (string, error) {
	c := exec.Command("go", "mod", "graph")
	b, err := c.Output()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func parseMods(s string) (*Mod, error) {
	lines := strings.Split(s, "\n")
	if len(lines) == 1 && len(lines[0]) == 0 {
		return nil, nil
	}
	mod := NewMod(lines[0][:strings.Index(lines[0], " ")])
	for i := 0; i < len(lines); i++ {
		m, d, ok := strings.Cut(lines[i], " ")
		if ok {
			mod.AddDependency(m, d)
		}
	}
	fmt.Println(mod)
	return mod, nil
}
