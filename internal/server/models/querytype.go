package models

//go:generate stringer -linecomment -type QueryType
type QueryType int

const (
	Word   QueryType = iota // word
	Number                  // number
)

func (i QueryType) WhereColumn() string {
	return i.String()
}

func (i QueryType) DistinctColumn() string {
	switch i {
	case Word:
		return Number.String()
	case Number:
		return Word.String()
	}
	return i.String()
}
