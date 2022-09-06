package main

import (
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

func (m *Mod) Tree(depth int) string {
	if depth == 0 {
		return "..."
	}
	sb := strings.Builder{}
	sb.WriteString(m.name)
	sb.WriteRune('\n')
	depSb := strings.Builder{}
	i := 0
	for _, d := range m.dependencies {
		if i != len(m.dependencies)-1 {
			depSb.WriteString("├──")
		} else {
			depSb.WriteString("└──")
		}
		depSb.WriteString(d.Tree(depth - 1))
		i++
	}
	sb.WriteString(tab(depSb.String()))
	return sb.String()
}

func tab(s string) string {
	lines := strings.Split(s, "\n")
	sb := strings.Builder{}
	for _, line := range lines {
		sb.WriteString("\t" + line + "\n")
	}
	return sb.String()
}
