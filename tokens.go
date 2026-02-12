package main

// Creating custom type for token called TokenType
type TokenType string

// DEFINE TOKEN TYPES FOR JSON
const (
	// EOF & Illegal
	EOF     TokenType = "EOF"
	ILLEGAL TokenType = "ILLEGAL"

	// JSON Structural Symbols
	L_BRACKET TokenType = "["
	R_BRACKET TokenType = "]"
	L_BRACE   TokenType = "{"
	R_BRACE   TokenType = "}"
	COLON     TokenType = ":"
	COMMA     TokenType = ","

	// Literals
	STRING  TokenType = "STRING"
	NUMBER  TokenType = "NUMBER"
	BOOLEAN TokenType = "BOOLEAN"
	NULL    TokenType = "NULL"
)

// Creating struct for holding data for every type
type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Col     int
}
