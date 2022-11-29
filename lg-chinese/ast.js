const Identifier = (value) => {
  return {
    value,
    type: 'Identifier',
  }
}

const NumericLiteral = (value) => {
  return {
    value: Number(value),
    type: 'NumericLiteral',
  }
}

const AST_LET = (id, value) => {
  return {
    id,
    value,
    type: 'LET'
  }
}

const AST_IN_OPT = (left, opt, right) => {
  return {
    left,
    opt,
    right,
    type: 'IN_OPT'
  }
}

const AST_PRT = (vlues) => {
  return {
    vlues,
    type: 'PRINT'
  }
}

const AST_FN = (params, body) => {
  return {
    params,
    body,
    type: 'FUNCTION',
  }
}


module.exports = {
  Identifier,
  NumericLiteral,
  AST_LET,
  AST_IN_OPT,
  AST_PRT,
  AST_FN,
}