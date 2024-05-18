package token

var keywords = map[string]TokenType{
	"let": LET,
	"fn":  FUNCTION,
}

func LookupIdentifier(literal string) TokenType {
	if tokenType, ok := keywords[literal]; ok {
		return tokenType
	}

	return IDENTIFIER
}
