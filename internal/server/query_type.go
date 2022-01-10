package server

type QueryType int

const (
	QueryWord QueryType = iota
	QueryNumber
)

func (col QueryType) WhereColumn() string {
	switch col {
	case QueryWord:
		return "word"
	case QueryNumber:
		return "number"
	}
	return ""
}

func (col QueryType) DistinctColumn() string {
	switch col {
	case QueryWord:
		return "number"
	case QueryNumber:
		return "word"
	}
	return ""
}
