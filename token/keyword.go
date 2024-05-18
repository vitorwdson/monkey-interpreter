package token

var keywords = map[string]TokenType{
	"let":    LET,
	"fn":     FUNCTION,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdentifier(literal string) TokenType {
	if tokenType, ok := keywords[literal]; ok {
		return tokenType
	}

	return IDENTIFIER
}
