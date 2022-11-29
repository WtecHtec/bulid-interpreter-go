const TOKEN_TYPE = {
  ILLEGAL : "ILLEGAL",
	EOF     : "EOF" , 

	IDENT  : "IDENT" ,
	INT    : "INT"  , 
	STRING : "STRING",

  PRINT: 'PRINT',
  FUNCTION: 'FUNCTION',
  LET: 'LET',
   
  ASSIGN: "=",
  ASTERISK : "*",

  LPAREN  :  "(",
	RPAREN   : ")",
	LBRACE  :  "{",
	RBRACE   : "}",


  COMMA: ",",
	SEMICOLON: ";",
}

const NewToken = (tokenType, value) => {
  return {
    value,
    type: tokenType,
  }
}

module.exports = {
  NewToken,
  TOKEN_TYPE
}