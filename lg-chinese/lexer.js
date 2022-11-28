
const { NewToken, TOKEN_TYPE} = require('./token')
module.exports = class NewLexer {
   constructor(code) {
    this.code = code
    this.low = 0
    this.fast = 0
    this.ch = ' '
    this.lexerError = []
   }

   readChar() {
    if (this.fast >= this.code.length) {
      this.ch = 'EOF'
    } else {
      this.ch = this.code[this.fast]
    }
    this.low = this.fast
    this.fast += 1
   }

   nextToken() {
    let token = null;
    this.trimChar()
     switch(this.ch) {
      case 'EOF':
        token = NewToken(TOKEN_TYPE.EOF, this.ch);
        break
      case '设':
        token = NewToken(TOKEN_TYPE.LET, this.ch);
        break
      case '为':
          token = NewToken(TOKEN_TYPE.ASSIGN, this.ch);
          break 
      case '等':
        const ch = this.ch
        if (this.peekChar() === '于') {
          this.readChar()
          token = NewToken(TOKEN_TYPE.ASSIGN, ch + this.ch);
        } else {
          this.pushErrors('关键字 等于 缺少')
        }
        break
      case '乘':
          token = NewToken(TOKEN_TYPE.ASTERISK, this.ch);
        
          break
      case '打':
          const pch = this.ch
          if (this.peekChar() === '印') {
            this.readChar()
            token = NewToken(TOKEN_TYPE.PRINT, pch + this.ch);
          } else {
            this.pushErrors('关键字 打印 缺少')
          }
          break
      case ';':
            token = NewToken(TOKEN_TYPE.SEMICOLON, this.ch);
            break 
      case ',':
            token = NewToken(TOKEN_TYPE.COMMA, this.ch);
            break 
      default:
        if (this.isLetter()) {
          token = NewToken(TOKEN_TYPE.IDENT, this.readLetter())
        } else if (this.isNumber()) {
          token = NewToken(TOKEN_TYPE.INT, this.readNumber())
        } else {
          token = NewToken(TOKEN_TYPE.ILLEGAL, this.ch);
        }
     }
     this.readChar()
     return token
   }
   
   trimChar() {
    while(this.ch === ' ' || this.ch === '\t' || this.ch === '\n') {
      this.readChar()
    }
   }

   isNumber() {
    return /\d/.test(this.ch)
   }

   readNumber() {
    let num = ''
    while(this.isNumber(this.ch) && this.ch != 'EOF') {
      num = `${num}${this.ch}`
      this.readChar()
    }
    if (!this.isNumber(this.ch) && this.ch != 'EOF') {
      this.low -= 1
      this.fast -= 1
    }
    return num
   }


   isLetter() {
    return /[a-zA-Z]/.test(this.ch)
   }
   
   readLetter() {
    let leter = ''
    while(this.isLetter(this.ch) && this.ch != 'EOF') {
      leter = `${leter}${this.ch}`
      this.readChar()
    }
    if (!this.isLetter(this.ch) && this.ch != 'EOF') {
      this.low -= 1
      this.fast -= 1
    }
    return leter
   }

   peekChar() {
    return this.code[this.fast]
   }

   pushErrors(error) {
    this.lexerError.push(error)
   }
}