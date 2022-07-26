package token

type TokenType string

type Token struct {
	Type    TokenType // 操作类型
	Literal string    // 属性值
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 标识符 定义符
	IDENT = "IDENT" // add, foobar, x, y,...
	INT   = "INT"   // 1343456

	// 操作符
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	EQ       = "=="
	NOT_EQ   = "!="
	LT       = "<"
	GT       = ">"

	// 分隔符
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// 关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var Keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

/*判断当前字符是否是关键词*/
func LookupIdent(ident string) TokenType {
	if tok, ok := Keywords[ident]; ok {
		return tok
	}
	return IDENT
}
