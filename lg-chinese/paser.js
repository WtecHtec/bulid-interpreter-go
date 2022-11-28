const { AST_LET, Identifier, NumericLiteral, AST_IN_OPT, AST_PRT } = require("./ast")
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
    }

    this.infixParseFns = {
      [TOKEN_TYPE.ASSIGN]: this.parseInfixAst.bind(this),
      [TOKEN_TYPE.ASTERISK]: this.parseInfixAst.bind(this),
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
      if (this.curToken.type === TOKEN_TYPE.COMMA) {
        this.nextToken()
      }
    }
    if (this.curToken.type !== TOKEN_TYPE.SEMICOLON) {
      this.pushError('缺少;')
    }
    return AST_PRT(values)
  }

  pushError(error) {
    this.errors.push(error)
  }
}

module.exports =  Paser
