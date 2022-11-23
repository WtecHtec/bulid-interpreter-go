package token

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // 开始
	EOF     = "EOF"     // 结束

	// 标识符（变量、函数名等） 定义符（数值、字符串、日期等）
	IDENT = "IDENT" // add, foobar, x, y,...
	INT   = "INT"   // 1343456

	// 操作符
	ASSIGN = "="
	PLUS   = "+"

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
)
