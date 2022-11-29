
const Eval = require('./eval')
const NewLexer = require('./lexer')
const Paser = require('./paser')
const code = ` 
 设 x 为  2; 
 y 等于 200 乘 x; 
 打印 x,y,x 乘 y;
 设 add 为 方程(x, y) {
  返回 x 乘 2
 };
 add(2)`
const lexer = new NewLexer(code)
if (lexer.errors.length) {
  console.log('词法失败==', lexer.errors.join(','))
  return
}
const paser = new Paser(lexer)
if (paser.errors.length) {
  console.log('语法失败==', paser.errors.join(','))
  return
}
const astNodes = paser.PaserParams()
console.log(astNodes)
// Eval(astNodes)

