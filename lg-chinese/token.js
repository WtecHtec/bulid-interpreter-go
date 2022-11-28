const TOKEN_TYPE = {
  ILLEGAL : "ILLEGAL",
	EOF     : "EOF" , 

	IDENT  : "IDENT" ,
	INT    : "INT"  , 
	STRING : "STRING",

  PRINT: 'PRINT',

  LET: 'LET',
   
  ASSIGN: "=",
  ASTERISK : "*",


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