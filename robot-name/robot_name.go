package robotname

import (
	"fmt"
	"math/rand"
)

var allNames = make(map[string]struct{})

type Robot struct {
	name string
}

func (r *Robot) Name() string {
	if r.name != "" {
		return r.name
	}
	for {
		name := newName()
		if _, ok := allNames[name]; !ok {
			r.name = name
			allNames[name] = struct{}{}
			return name
		}
	}
}

func (r *Robot) Reset() {
	r.name = ""
}

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func newName() string {
	l1 := letters[rand.Intn(26)]
	l2 := letters[rand.Intn(26)]
	d1 := rand.Intn(10)
	d2 := rand.Intn(10)
	d3 := rand.Intn(10)
	return fmt.Sprintf("%c%c%d%d%d", l1, l2, d1, d2, d3)
}
