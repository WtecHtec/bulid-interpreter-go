const { AST_LET, Identifier, NumericLiteral, AST_IN_OPT, AST_PRT, AST_FN, AST_BLOCK, AST_RETURN,AST_CALL } = require("./ast")
const { TOKEN_TYPE } = require("./token")

class Paser {
  constructor(lexer) {
    this.lexer = lexer
    this.curToken = null
    this.peekToken = null
    this.errors = []
    this.prefixParseFns = {
      [TOKEN_TYPE.IDENT]: Identifier,
      [TOKEN_TYPE.INT]: NumericLiteral,
      [TOKEN_TYPE.FUNCTION]: this.parseFuncAst.bind(this),
      [TOKEN_TYPE.LPAREN]: this.parseLpAst.bind(this),
    }

    this.infixParseFns = {
      [TOKEN_TYPE.ASSIGN]: this.parseInfixAst.bind(this),
      [TOKEN_TYPE.ASTERISK]: this.parseInfixAst.bind(this),
			[TOKEN_TYPE.LPAREN]: this.parseCallAst.bind(this),
    }
  }

  PaserParams() {
    const ps = []
    this.nextToken()
    this.nextToken()
    while( this.curToken && this.curToken.type != TOKEN_TYPE.EOF) {
      const ast = this.getAstNode()
      ast && ps.push(ast)
      this.nextToken()
    }
    return ps
  }

  getAstNode() {
    let ast = null
    if (this.curToken.type === TOKEN_TYPE.LET) {
      ast = this.paserLet()
    } else if (this.curToken.type === TOKEN_TYPE.PRINT) {
      ast = this.paserPrint()
    } else if(this.curToken.type === TOKEN_TYPE.RETURN) {
			ast = this.paserReturn()
		} else {
      ast = this.paserExAst()
    }
    return ast
  }

  nextToken() {
    this.curToken = this.peekToken
    this.peekToken = this.lexer.nextToken()
  }

  // let x = 0; <let> <id> <=> <ex> ;
  paserLet() {
    this.nextToken()
    if (this.curToken.type !== TOKEN_TYPE.IDENT) {
      this.pushError('没有变量')
    }
    const id = Identifier(this.curToken.value) 
    this.nextToken()
    if (this.curToken.type !== TOKEN_TYPE.ASSIGN) {
      this.pushError('没有赋值')
    }

    this.nextToken()
    const value = this.paserExAst()

    this.nextToken()
    if (this.curToken.type !== TOKEN_TYPE.SEMICOLON) {
      this.pushError('要以 ; 结尾')
    }
    return AST_LET(id, value)
  }

  // a = 1; <id> <opt> <ex>
  parseInfixAst(leftExp) {
    const ct = leftExp;
    this.nextToken()
    const opt = this.curToken.value
    this.nextToken()
    const right =  this.paserExAst()
    return AST_IN_OPT( ct, opt,  right);
  }
  // re <ex>
	paserReturn() {
		const res = AST_RETURN()
		this.nextToken()
		res.value = this.paserExAst()
    this.nextToken()
		return res
	}
  paserExAst() {
    const prefn = this.prefixParseFns[this.curToken.type]
    let leftExp = null
    if (typeof prefn === 'function' ) {
      leftExp = prefn(this.curToken.value)
    }
    const infixfn = this.infixParseFns[this.peekToken.type]
    if (typeof infixfn === 'function') {
      leftExp = infixfn(leftExp)
    }
    return leftExp;
  }
  // <cl> <ex> ,<ex>....
  paserPrint() {
    const values = []
    this.nextToken()
    while(this.curToken.type !== TOKEN_TYPE.SEMICOLON && this.curToken.type !== TOKEN_TYPE.EOF) {
      values.push( this.paserExAst(this.curToken) )
      this.nextToken()
      if (this.curToken.type === TOKEN_TYPE.COMMA ) {
        this.nextToken()
      }
    }
    if (this.curToken.type !== TOKEN_TYPE.SEMICOLON) {
      this.pushError('缺少;')
    }
    return AST_PRT(values)
  }
  // fn (p) {ex}
  parseFuncAst() {
    const fnAst = AST_FN()
    this.nextToken()
    if (this.curToken.type !== TOKEN_TYPE.LPAREN) {
      this.pushError('方程 缺 (')
    }
   
    fnAst.params = this.parsePramaAst(this.curToken)

    if (this.curToken.type !== TOKEN_TYPE.LBRACE) {
      this.pushError('方程 缺 {')
    }
    fnAst.body = this.parseBodyAst()
    return fnAst;
  }

  parsePramaAst() {
    const res = []
    this.nextToken()
    while(this.peekToken.type === TOKEN_TYPE.COMMA ) {
      res.push(Identifier(this.curToken.value))
      this.nextToken()
      this.nextToken()
    }
    if (this.curToken.type !== TOKEN_TYPE.COMMA) {
      res.push(Identifier(this.curToken.value))
      this.nextToken()
    }
    if (this.curToken.type !== TOKEN_TYPE.COMMA) {
      this.pushError('方程 缺 )')
    }
    this.nextToken()
    return res
  }
  parseBodyAst() {
    const res = []
		this.nextToken()

		while(this.curToken.type !== TOKEN_TYPE.EOF && this.curToken.type !== TOKEN_TYPE.RBRACE) {
			const ast = this.getAstNode()
			res.push(ast)
			this.nextToken()
		}
		// this.nextToken()
    // console.log('parseBodyAst===', this.curToken)
		
		return AST_BLOCK(res)

  }
	// <fn>(<ex>)
	parseCallAst(leftExp) {
		const res = AST_CALL(leftExp)
		if (this.peekToken.type !== TOKEN_TYPE.LPAREN) {
			this.pushError('方程缺失 （')
		}
		res.params = this.parseLpAst()
    if (this.curToken.type !== TOKEN_TYPE.RPAREN) {
      this.pushError('方程缺失 )')
    }
    this.nextToken()
		return res
	}
	// （<ex>, <ex>）
  parseLpAst() {
		const res = []
		this.nextToken()
		this.nextToken()
    while(this.peekToken.type === TOKEN_TYPE.COMMA ) {
      res.push(this.paserExAst())
      this.nextToken()
      this.nextToken()
    }
		if (this.curToken.type !== TOKEN_TYPE.COMMA) {
      res.push(this.paserExAst())
      this.nextToken()
    }
    if (this.curToken.type !== TOKEN_TYPE.RPAREN) {
      this.pushError('方程 缺 )')
    }
    // this.nextToken()
    return res
  }

  pushError(error) {
    this.errors.push(error)
  }
}

module.exports =  Paser
