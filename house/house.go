package house

import (
	"bytes"
	"fmt"
	"strings"
)

const testVersion = 1

type Record struct {
	name   string
	action string
}

var records = []*Record{
	{"house that Jack built", "lay in"},
	{"malt", "ate"},
	{"rat", "killed"},
	{"cat", "worried"},
	{"dog", "tossed"},
	{"cow with the crumpled horn", "milked"},
	{"maiden all forlorn", "kissed"},
	{"man all tattered and torn", "married"},
	{"priest all shaven and shorn", "woke"},
	{"rooster that crowed in the morn", "kept"},
	{"farmer sowing his corn", "belonged to"},
	{"horse and the hound and the horn", ""},
}

func Verse(n int) string {
	result := &bytes.Buffer{}
	result.WriteString("This is the " + records[n-1].name)

	for i := n - 1; i > 0; i-- {
		result.WriteString("\n")
		result.WriteString(fmt.Sprintf("that %s the %s", records[i-1].action, records[i-1].name))
	}

	result.WriteString(".")
	return result.String()
}

func Song() string {
	var verses []string
	for i := 1; i <= len(records); i++ {
		verses = append(verses, Verse(i))
	}

	return strings.Join(verses, "\n\n")
}
