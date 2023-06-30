package main

import (
	"fmt"
	"strings"
)

type Mod struct {
	name string
	deps map[string]*Mod
}

func NewMod(name string) *Mod {
	return &Mod{
		name: name,
		deps: make(map[string]*Mod),
	}
}

func (m *Mod) AddDependency(mod string, dependency string) {
	if mod == m.name {
		d := &Mod{
			name: dependency,
			deps: make(map[string]*Mod),
		}
		m.deps[dependency] = d
	}
	for _, d := range m.deps {
		d.AddDependency(mod, dependency)
	}
}

func (m *Mod) Tree(depth int, vl bool) string {
	sb := strings.Builder{}
	sb.WriteString(m.name)
	sb.WriteRune('\n')
	depSb := strings.Builder{}
	i := 0
	for _, d := range m.deps {
		last := i == len(m.deps)-1
		if !last && depth > 0 {
			depSb.WriteString("├───")
		} else {
			depSb.WriteString("└───")
		}
		if depth > 0 {
			depSb.WriteString(d.Tree(depth-1, !last))
		} else {
			depSb.WriteString(fmt.Sprintf("%d more ...\n", len(m.deps)))
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

func ParseMods(graph string) *Mod {
	lines := strings.Split(graph, "\n")
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
