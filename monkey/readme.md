### 词法分析
###### 计算机科学中将字符序列转换为记号（token）序列的过程
定义一个结构体(词法标记)：
```
  type TokenType string
  type Token struct {
    Type TokenType // token 类型
    Literal string  // token 属性值
  }
```
设想代码需要的Tokens：
例如：

### 词法分析器 将输入转换为tokens
语法分析器读取输入字符流、从中识别出词法标记、最后生成不同类型的记号
例如：
sum=3+2;
（
sum	标识符
=	赋值操作符
3	数字
+	加法操作符
2	数字）
###### 实现扫描
扫码代码，其实代码就是一个字符串，
let num = 3 + 2
从字符串索引0开始，当扫描到一个单词时，就判断是不是关键字，或者是变量
相当于从l扫描到下一个空格或者下一个字符是定义的关键字、字符时，切片这一段字符，
例如 ：let num => let , num
###### == 、!= 关键字
当扫描到字符 = 、！时，判断下一个字符是否 符合==、！=
```
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
```
### 实现一个 repl 控制台

#### 解析器
递归下降解析（“自顶向下运算符优先级”解析器）
（自顶向下和自底向上解析器的区别在于，前者从构建AST的根节点开始，然后下降，而后者相反）
一般是指把某种格式的文本（字符串）转换成某种数据结构的过程。
程序文本转换成编译器内部的一种叫做抽象语法树（AST）的数据结构，也叫做语法分析器（js）

##### let
let x = 10;
let <identifier> = <expression>;
x =>标识符; 10 =>数值、函数表达式

#### return
return f
return <expression>
f =>数值、函数表达式

#### 表达式
关键字 前缀：
```
--1
-1
fn () {}
let a = 1
``` 
中缀：
```
1 + 2
a > 1
```
后缀：
```
2--
2++
```
每种token类型最多可以有两个与之关联的解析函数；定义两种类型的函数一个前缀解析函数和一个中缀解析函数。
	prefixParseFns map[token.TokenType]prefixParseFn // 检查是否合适的map(前缀)有一个和curToken.Type联系的解析函数。
	infixParseFns  map[token.TokenType]infixParseFn  // 检查是否合适的map(中缀)有一个和curToken.Type联系的解析函数。

  例如：
  ```
	p.registerPrefix(token.IDENT, p.parseIdentifier)
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}
{
  token.IDENT: parseIdentifier
}

```
当扫描到 token 类型type 为 IDENT,触发parseIdentifier，返回一个IDENT的AST抽象树
#### 前缀运算 +5， -5
<prefix operator><expression>;

type PrefixExpression struct {
	Token    token.Token 
	Operator string
	Right    Expression
}

当遇到 + 符号 前缀时，去扫描下一个字符，负值给 Right
```
// Right = 表达式
expression.Right = p.parseExpression(PREFIX)
```


#### 中缀运算  5 + 5
<expression> <infix operator> <expression>

```
type InfixExpression struct {
	Token    token.Token 
	Left     Expression
	Operator string
	Right    Expression
}
```
扫描到字符5时，预检测下一个字符，如果运算优先级大于 5 ,执行 parseInfixExpression 
```
  // expression_statement.go
	//   当前优先级 小于 下一个字符优先级
	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp)
	}

```


#### 解析 -1 * 5 => (-1) * 5
扫描到字符- ,确认为前缀运算，此时运算优先级为PREFIX，下一个扫描到* 的时，运算优先级比它小，所以这个时候没有运行中缀运算。
```
// 运算优先级排序
const (
	_           int = iota
	LOWEST          // 最小
	EQUALS          // ==
	LESSGREATER     // > or <
	SUM             // +
	PRODUCT         // *
	PREFIX          // -X or !X
	CALL            // myFunction(X)
)
```

```
// 中缀运算逻辑
	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()
		leftExp = infix(leftExp)
	}

```
#### 解析 块级 1 * (1 - 2)
以扫描( 开始，到) 结束，中间为表达式,( 设置前缀运算
( <expression> )
```

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

  // <expression>
	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}

```

#### 解析 if else 
```
 if ( x > y) { return x } else { return y }
 if (<condition>) <consequence> else <alternative>

 ```
添加 ast BlockStatement ，代码运行块， consequence、  alternative都属于 BlockStatement，
condition 则为 表达式 expression
```
type BlockStatement struct {
	Token      token.Token // the { token
	Statements []Statement
}
```
#### 解析 fn() {}
fn <parameters> <block statement>
抽象ast: 
```
type FunctionLiteral struct {
	Token      token.Token // The 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
}

```
#### 调用表达式

<expression>(<comma separated expressions>)
抽象ast:
```
type CallExpression struct {
	Token     token.Token // The '(' token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

```
add(1 + 2)

扫描到add 为ident，下一个token为 (, 此时优先级大于add，执行中缀运算
```
	p.registerInfix(token.LPAREN, p.parseCallExpression)

```

```

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Token: p.curToken, Function: function}
	exp.Arguments = p.parseCallArguments()
	return exp
}

```
#### 表达式eval 自上而下递归
解析器的概念是不会留下可执行文件的东西(与编译器相反，编译器可以留下可执行文件)在查看时变得非常模糊在现实世界和高度优化的编程语言实现中。

```
func Eval(node ast.Node) object.Object {
		//  node.(type) 获取节点类型
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

		...
}

```
从根节点开始遍历，根据每个token节点类型执行不同的操作。当遇到表达式，执行递归。

