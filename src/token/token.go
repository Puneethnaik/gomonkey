package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + Literals
	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"

	//Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "#"
	SLASH = "/"
	LT = "<"
	GT = ">"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "func"
	LET      = "let"
)

var keywords = map[string]TokenType{
	"func": FUNCTION,
	"let":  LET,
}

// LookupIdent function to lookup whether ident matches one of the keywords of the language.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENTIFIER
}
