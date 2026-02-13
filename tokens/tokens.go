package tokens

import "fmt"

// Creating custom type for token called TokenType
type TokenType string

// DEFINE TOKEN TYPES FOR JSON
const (
	// EOF & Illegal
	EOF     TokenType = "EOF"
	ILLEGAL TokenType = "ILLEGAL"

	// JSON Structural Symbols
	L_BRACKET TokenType = "LEFT BRACKET"
	R_BRACKET TokenType = "RIGHT BRACKET"
	L_BRACE   TokenType = "LEFT BRACES"
	R_BRACE   TokenType = "RIGHT BRACES"
	COLON     TokenType = "COLON"
	COMMA     TokenType = "COMMA"

	// Literals
	STRING  TokenType = "STRING"
	NUMBER  TokenType = "NUMBER"
	BOOLEAN TokenType = "BOOLEAN"
	NULL    TokenType = "NULL"

	// identifiers values
	TRUE  TokenType = "TRUE"
	FALSE TokenType = "FALSE"
)

// Creating struct for holding data for every type
type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Col     int
}

var keywords = map[string]TokenType{
	"true":  BOOLEAN,
	"false": BOOLEAN,
	"null":  NULL,
}

func LookupIdentifier(ident string) (TokenType, error) {
	tokenType, ok := keywords[ident]
	if ok {
		return tokenType, nil
	}

	return ILLEGAL, fmt.Errorf("unknown identifier: %s", ident)
}
