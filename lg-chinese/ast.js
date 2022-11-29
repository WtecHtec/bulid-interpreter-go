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

const AST_FN = (params, body, env) => {
  return {
		env,
    params,
    body,
    type: 'FUNCTION',
  }
}

const AST_BLOCK = (values) => {
	return {
		values,
		type: 'BLOCK',
	}
}

const AST_RETURN = (value) => {
	return {
		value: value,
		type: 'RETURN',
	}
}

const AST_CALL = (fn, params) => {
	return {
		params,
		fn,
		type: 'CALL'
	}
}

module.exports = {
  Identifier,
  NumericLiteral,
  AST_LET,
  AST_IN_OPT,
  AST_PRT,
  AST_FN,
	AST_BLOCK,
	AST_RETURN,
	AST_CALL,
}