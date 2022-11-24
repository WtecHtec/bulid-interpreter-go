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

let x = 10;
let <identifier> = <expression>;
x 标识符 10 数值、函数表达式