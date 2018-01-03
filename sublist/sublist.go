package sublist

type Relation string

const (
	EqualRelation     Relation = "equal"
	SuperlistRelation          = "superlist"
	SublistRelation   Relation = "sublist"
	UnequalRelation   Relation = "unequal"
)

func Sublist(a, b []int) Relation {
	aLength := len(a)
	bLength := len(b)
	switch {
	case aLength == bLength && isEqual(a, b):
		return EqualRelation
	case aLength > bLength && isSuperlist(a, b):
		return SuperlistRelation
	case aLength < bLength && isSuperlist(b, a):
		return SublistRelation
	}
	return UnequalRelation
}

func isSuperlist(a, b []int) bool {
	bLength := len(b)
	if bLength == 0 {
		return true
	}

	for i := 0; i <= len(a)-len(b); i++ {
		if a[i] == b[0] {
			if isEqual(a[i:i+bLength], b) {
				return true
			}
		}
	}

	return false
}

func isEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
