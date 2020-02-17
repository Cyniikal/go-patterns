package strategy

import "strings"

type Strategy interface {
	Apply(string) string
}

type Tactic struct {
	Strategy
}

func (t *Tactic) Run(s string) string {
	return t.Apply(s)
}

type CapsTactic struct{}
func (CapsTactic) Apply(s string) string {
	return strings.ToUpper(s)
}

type LowerTactic struct{}
func (LowerTactic) Apply(s string) string {
	return strings.ToLower(s)
}

