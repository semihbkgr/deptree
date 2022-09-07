package main

import (
	"fmt"
	"strings"
)

type Mod struct {
	name         string
	dependencies map[string]*Mod
}

func NewMod(name string) *Mod {
	return &Mod{
		name:         name,
		dependencies: make(map[string]*Mod),
	}
}

func (m *Mod) AddDependency(mod string, dependency string) {
	if mod == m.name {
		d := &Mod{name: dependency, dependencies: make(map[string]*Mod)}
		m.dependencies[dependency] = d
	}
	for _, d := range m.dependencies {
		d.AddDependency(mod, dependency)
	}
}

func (m *Mod) Tree(depth int, vl bool) string {
	sb := strings.Builder{}
	sb.WriteString(m.name)
	sb.WriteRune('\n')
	depSb := strings.Builder{}
	i := 0
	for _, d := range m.dependencies {
		last := i == len(m.dependencies)-1
		if !last {
			depSb.WriteString("├───")
		} else {
			depSb.WriteString("└───")
		}
		if depth > 1 {
			depSb.WriteString(d.Tree(depth-1, !last))
		} else {
			depSb.WriteString(fmt.Sprintf("%d more ...\n", len(m.dependencies)))
			break
		}
		i++
	}
	if vl {
		sb.WriteString(linePrefix(depSb.String(), "│    "))
	} else {
		sb.WriteString(linePrefix(depSb.String(), "     "))
	}
	return sb.String()
}

func linePrefix(s string, p string) string {
	lines := strings.Split(s, "\n")
	sb := strings.Builder{}
	for _, line := range lines {
		if len(line) != 0 {
			sb.WriteString(p + line + "\n")
		}
	}
	return sb.String()
}
