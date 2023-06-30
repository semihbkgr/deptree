package main

import (
	"os"
	"strings"
	"testing"
)

func TestParseMods(t *testing.T) {
	graph, err := readGraph()
	if err != nil {
		t.Fatal(err)
	}

	m := ParseMods(graph)
	if m == nil {
		t.Fatal("returned mod is nil")
	}

	if m.name != "github.com/gin-gonic/gin" {
		t.Error("mod name is not as expected")
	}

	expectedTree, err := readTree()
	if err != nil {
		t.Fatal(err)
	}

	tree := m.Tree(10, false)
	if treeMatch(expectedTree, tree) {
		t.Error("tree does not match with the expected tree")
	}
}

func readGraph() (string, error) {
	return readFile("testdata/graph")
}

func readTree() (string, error) {
	return readFile("testdata/tree")
}

func readFile(f string) (string, error) {
	b, err := os.ReadFile(f)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func treeMatch(t1, t2 string) bool {
	lineCountMap := make(map[string]int)
	for _, line := range strings.Split(t1, "\n") {
		if len(line) == 0 {
			continue
		}
		lineCountMap[line]++
	}
	for _, line := range strings.Split(t2, "\n") {
		if len(line) == 0 {
			continue
		}
		lineCountMap[line]--
	}
	for _, v := range lineCountMap {
		if v != 0 {
			return false
		}
	}
	return true
}
