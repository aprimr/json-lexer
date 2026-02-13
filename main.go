package main

import (
	"fmt"
	"gojsonlexer/tokens"
	"os"
)

// struct for lexer
type Lexer struct {
	Input        string // raw JSON string
	position     int    // points to the char we are looking at
	nextPosition int    // points to the next index (peek)
	char         byte   // current character byte present at current position
	line         int
	lineStart    int
}

func main() {
	// Open json file and read the data
	dataBytes, err := os.ReadFile("./data.json")
	if err != nil {
		panic(err)
	}
	data := string(dataBytes) // convert []byte into string

	// Init Lexer
	l := &Lexer{Input: data}
	l.readChar()

	for {
		tok := l.NextToken()
		fmt.Printf("%+v \n", tok)
		if tok.Type == tokens.EOF {
			break
		}
	}

}

// ReadChar method to read the current char and move the cursor to next char
func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.Input) {
		// The pointer reached to EOF
		// Set the char to 0 which represent NULL
		l.char = 0
	} else {
		l.char = l.Input[l.nextPosition] // set Char the value the position pointer is pointing
	}
	// Increment pointers
	l.position = l.nextPosition
	l.nextPosition++
}

// Skip whitespaces
func (l *Lexer) skipWhiteSpaces() {
	for l.char == ' ' || l.char == '\n' || l.char == '\t' || l.char == '\r' {
		if l.char == '\n' {
			l.line++
			l.lineStart = l.nextPosition
		}
		l.readChar()
	}
}

func newToken(tokenType tokens.TokenType, line int, col int, char byte) tokens.Token {
	return tokens.Token{
		Type:    tokenType,
		Literal: string(char),
		Line:    line,
		Col:     col,
	}
}

// NextToken switches through the lexer's current char and creates a new token.
func (l *Lexer) NextToken() tokens.Token {
	l.skipWhiteSpaces()
	var t tokens.Token

	switch l.char {
	case '[':
		t = newToken(tokens.L_BRACKET, l.line+1, l.GetCol(), l.char)
	case ']':
		t = newToken(tokens.R_BRACKET, l.line+1, l.GetCol(), l.char)
	case '{':
		t = newToken(tokens.L_BRACE, l.line+1, l.GetCol(), l.char)
	case '}':
		t = newToken(tokens.R_BRACE, l.line+1, l.GetCol(), l.char)
	case ':':
		t = newToken(tokens.COLON, l.line+1, l.GetCol(), l.char)
	case ',':
		t = newToken(tokens.COMMA, l.line+1, l.GetCol(), l.char)
	case '"':
		col := l.GetCol()
		t.Literal = l.readString()
		t.Type = tokens.STRING
		t.Line = l.line + 1
		t.Col = col
		return t
	case 0:
		col := l.GetCol()
		t.Literal = ""
		t.Type = tokens.EOF
		t.Line = l.line + 1
		t.Col = col
	default:
		if isLetter(l.char) {
			col := l.GetCol()
			ident := l.readIdentifier()
			t.Literal = ident
			t.Line = l.line + 1
			t.Col = col
			tokType, err := tokens.LookupIdentifier(ident)
			if err != nil {
				t.Type = tokens.ILLEGAL
				return t
			}
			t.Type = tokType
			return t
		} else if isNumber(l.char) {
			col := l.GetCol()
			t.Literal = l.readNumber()
			t.Type = tokens.NUMBER
			t.Line = l.line + 1
			t.Col = col
			return t
		} else {
			t = newToken(tokens.ILLEGAL, l.line+1, 1, l.char)
		}
	}

	l.readChar()

	return t
}

// readString sets a starting position and reads
// string till it finds closing `"`,  and returns the
// string between the starting position and ending `"`
func (l *Lexer) readString() string {
	position := l.position + 1

	for {
		l.readChar()
		if l.char == '"' || l.char == 0 {
			break
		}
	}
	str := string(l.Input[position:l.position])
	l.readChar()
	return str
}

// readNumber sets a starting position and reads
// numbers and when it finds a char that isnt a number
// it stops and return data betn start and end positions
func (l *Lexer) readNumber() string {
	position := l.position

	for isNumber(l.char) {
		l.readChar()
	}

	return string(l.Input[position:l.position])
}

// calculate and return the column
func (l *Lexer) GetCol() int {
	return l.position - l.lineStart + 1
}

// Check if the data is number or not
func isNumber(data byte) bool {
	return '0' <= data && data <= '9' || data == '.'
}

// Check if the data is letter or not
func isLetter(data byte) bool {
	return ('a' <= data && data <= 'z') || ('A' <= data && data <= 'Z') || data == '_'
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.char) {
		l.readChar()
	}
	return string(l.Input[position:l.position])
}
