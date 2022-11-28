
const Eval = require('./eval')
const NewLexer = require('./lexer')
const Paser = require('./paser')
const code = ` 
 设 x 为  2; 
 y 等于 200 乘 x; 
 打印 x,y,x 乘 y;`
const lexer = new NewLexer(code)
const paser = new Paser(lexer)
if (paser.errors.length) {
  console.log('编译失败==', paser.errors.join(','))
  return
}
const astNodes = paser.PaserParams()
Eval(astNodes)

