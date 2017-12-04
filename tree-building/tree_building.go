package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

type Mismatch struct{}

func (m Mismatch) Error() string {
	return "c"
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := make([]*Node, len(records))
	for id, rec := range records {
		if id != rec.ID {
			return nil, errors.New("invalid id")
		}
		if rec.ID < rec.Parent {
			return nil, errors.New("invalid parent id")
		}
		if rec.ID != 0 && rec.ID == rec.Parent {
			return nil, errors.New("invalid parent id")
		}
		node := &Node{ID: rec.ID}
		nodes[id] = node
		if rec.ID != 0 {
			parent := nodes[rec.Parent]
			parent.Children = append(parent.Children, node)
		}
	}

	return nodes[0], nil
}
