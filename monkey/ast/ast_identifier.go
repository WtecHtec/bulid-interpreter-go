package ast

import (
	"bytes"
	"monkey/token"
)

type Identifier struct {
	Token token.Token // IDENT token 标记
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (ls *Identifier) String() string {
	var out bytes.Buffer
	out.WriteString(ls.Value)
	return out.String()
}
