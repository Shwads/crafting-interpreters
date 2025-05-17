package parser

type Token struct {
    typ TokenType
    lexeme string
    literal string
    line int
}
