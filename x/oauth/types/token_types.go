package types

type TokenType int

const (
	Basic TokenType = iota
	Bearer
)

func (tt TokenType) String() string {
	switch tt {
	case Basic:
		return "Basic"
	case Bearer:
		return "Bearer"
	}

	return "unknown"
}
