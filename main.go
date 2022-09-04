package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var (
	depth = flag.Int("depth", 5, "dependency tree depth")
)

func main() {
	flag.Parse()
	o, err := execGoModGraphCommand()
	checkErr(err)
	fmt.Println(o)
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
