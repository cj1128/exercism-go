package foodchain

import (
	"bytes"
	"fmt"
)

type item struct {
	name  string
	desc  string
	extra string
}

var items = []*item{
	{"fly", "I don't know why she swallowed the fly. Perhaps she'll die.", ""},
	{"spider", "It wriggled and jiggled and tickled inside her.", " that wriggled and jiggled and tickled inside her"},
	{"bird", "How absurd to swallow a bird!", ""},
	{"cat", "Imagine that, to swallow a cat!", ""},
	{"dog", "What a hog, to swallow a dog!", ""},
	{"goat", "Just opened her throat and swallowed a goat!", ""},
	{"cow", "I don't know how she swallowed a cow!", ""},
	{"horse", "She's dead, of course!", ""},
}

// start from 1
func Verse(n int) string {
	result := &bytes.Buffer{}

	for i := n; i >= 1; i-- {
		item := items[i-1]

		if i == n {
			result.WriteString(fmt.Sprintf("I know an old lady who swallowed a %s.", item.name))
			result.WriteString("\n")
			result.WriteString(item.desc)
			if n == len(items) {
				break
			} else {
				continue
			}
		}

		result.WriteRune('\n')

		next := items[i]

		result.WriteString(fmt.Sprintf(
			"She swallowed the %s to catch the %s%s.",
			next.name,
			item.name,
			item.extra,
		))

		if i == 1 {
			result.WriteRune('\n')
			result.WriteString(items[0].desc)
		}
	}

	return result.String()
}

// start from 1
func Verses(start, end int) string {
	result := &bytes.Buffer{}
	for n := start; n <= end; n++ {
		if n != start {
			result.WriteString("\n\n")
		}
		result.WriteString(Verse(n))
	}
	return result.String()
}

func Song() string {
	return Verses(1, len(items))
}
