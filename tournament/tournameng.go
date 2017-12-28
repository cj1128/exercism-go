package tournament

import (
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

type match struct {
	team1   string
	team2   string
	outcome string
}

type teamResult struct {
	name   string
	played int
	won    int
	tied   int
	lost   int
	points int
}

func Tally(r io.Reader, w io.Writer) error {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	content := string(buf)

	matches, err := parseMatches(content)
	if err != nil {
		return err
	}

	results := getTeamResults(matches)

	sort.Slice(results, func(i, j int) bool {
		if results[i].points == results[j].points {
			return results[i].name < results[j].name
		}
		return results[i].points > results[j].points
	})

	return printTeamResults(results, w)
}

func parseMatches(input string) ([]*match, error) {
	var result []*match
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		if strings.HasPrefix(line, "#") || len(line) == 0 {
			continue
		}

		team1, team2, outcome, err := parseMatch(line)
		if err != nil {
			return nil, err
		}

		result = append(result, &match{team1, team2, outcome})
	}
	return result, nil
}

func parseMatch(input string) (string, string, string, error) {
	parts := strings.Split(input, ";")
	if len(parts) != 3 {
		return "", "", "", fmt.Errorf("bad input: %s", input)
	}

	outcome := parts[2]
	if outcome != "win" && outcome != "draw" && outcome != "loss" {
		return "", "", "", fmt.Errorf("bad outcome: %s", outcome)
	}
	return parts[0], parts[1], outcome, nil
}

func getTeamResults(matches []*match) []*teamResult {
	m := make(map[string]*teamResult)
	var r1, r2 *teamResult
	for _, match := range matches {
		if m[match.team1] == nil {
			m[match.team1] = &teamResult{name: match.team1}
		}
		if m[match.team2] == nil {
			m[match.team2] = &teamResult{name: match.team2}
		}
		r1 = m[match.team1]
		r2 = m[match.team2]
		updateTeamResult(r1, r2, match.outcome)
	}
	var results []*teamResult
	for _, result := range m {
		results = append(results, result)
	}
	return results
}

func printTeamResults(results []*teamResult, w io.Writer) error {
	_, err := fmt.Fprintln(w, "Team                           | MP |  W |  D |  L |  P")
	if err != nil {
		return err
	}

	for _, result := range results {
		_, err := fmt.Fprintf(
			w,
			"%-30s | %2d | %2d | %2d | %2d | %2d\n",
			result.name,
			result.played,
			result.won,
			result.tied,
			result.lost,
			result.points,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func updateTeamResult(r1, r2 *teamResult, outcome string) {
	r1.played++
	r2.played++
	switch outcome {
	case "win":
		r1.won++
		r1.points += 3
		r2.lost++
	case "draw":
		r1.tied++
		r1.points += 1
		r2.tied++
		r2.points += 1
	case "loss":
		r1.lost++
		r2.won++
		r2.points += 3
	}
}
