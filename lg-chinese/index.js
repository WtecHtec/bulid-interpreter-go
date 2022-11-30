
const { Eval, env} = require('./eval')
const NewLexer = require('./lexer')
const Paser = require('./paser')
const code = ` 
 设 x 为  5; 
 y 等于 8 乘 x; 
 z 等于 8;
 设 add 为 方程(x, y) {
  y 等于 8;
  设 db 为 方程(y) {
    返回 y 乘 2;
  };
  打印 x,y,z;
  返回 x 乘 y 乘 z;
 };
 res 等于 add(2, y);
 打印 res,  x 乘 add(2, y);`
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
// console.log(astNodes)
Eval(astNodes, env)

