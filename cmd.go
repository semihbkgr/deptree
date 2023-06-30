package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/awalterschulze/gographviz"
)

var (
	depth = flag.Int("depth", 5, "depth of dependency tree")
	dot   = flag.Bool("dot", false, "output in dot format")
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

	if *dot {
		dotOut, err := modDot(graph)
		checkErr(err)
		fmt.Println(dotOut)
		return
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

func modDot(modGraph string) (string, error) {
	graphAst, _ := gographviz.ParseString(`digraph mod {}`)
	graph := gographviz.NewGraph()
	if err := gographviz.Analyse(graphAst, graph); err != nil {
		return "", err
	}
	graph.AddAttr("mod", "rankdir", "LR")
	graph.AddAttr("mod", "splines", "ortho")
	nodeAttr := map[string]string{
		"shape": "rect",
	}
	for _, l := range strings.Split(modGraph, "\n") {
		m, d, ok := strings.Cut(l, " ")
		m = fmt.Sprintf("%q", m)
		d = fmt.Sprintf("%q", d)
		if ok {
			if !graph.IsNode(m) {
				graph.AddNode("mod", m, nodeAttr)
			}
			if !graph.IsNode(d) {
				graph.AddNode("mod", d, nodeAttr)
			}
			graph.AddEdge(m, d, true, nil)
		}
	}
	return graph.String(), nil
}
