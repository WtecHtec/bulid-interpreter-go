const env = {
  store: {},
  path: null
}
// 设置函数环境
function setFnEnv(fn , ps, env) {
	const fnenv = {
		path: env,
		store: {},
	}
	const { params } = fn
	params.forEach((item ,i) => {
		fnenv.store[item.value] = ps[i] || null
	})
	return fnenv;
}

function getIden(env, val) {
  if (!env) return null
  if (env.store[val]) {
    return env.store[val]
  }
  return getIden(env.path, val)
}

function run(node, env) {
	if (!node) return null
  const type = node.type
  switch(type) {
    case 'LET': 
      return runLet(node, env)
    case 'NumericLiteral':
      return node.value;
    case 'Identifier':
      return runIden(node, env)
    case 'IN_OPT':
      return  runInfix(node, env)
    case 'PRINT':
      return runPrint(node, env)
		case 'FUNCTION': 
			return runFuntion(node, env)
		case 'CALL':
			 return runCall(node, env)
		case 'RETURN':
			 return run(node.value, env)
  }
  return null
}

function runLet(node, env) {
  const id = node.id.value
  env.store[id] = run(node.value, env)
  return null
}

function runInfix(node, env) {
  const { left, opt, right} = node
  switch(opt) {
    case '等于':
      return runLet( { id: left, value: right }, env)
    case '乘':
      return run(node.left, env) * run(node.right, env)
  }
}

function runPrint(node, env) {
  const { vlues } = node
  const result = []
  if (Array.isArray(vlues)) {
    vlues.forEach(item => {
      result.push(run(item, env) || 'undefine')
    })
  } 
  console.log(result.join(','))
  return null
}

function runIden(node, env) {
  return getIden(env, node.value)
}

function runFuntion(node, env) {
	node.env = env
	return node
}

function runCall(node, env) {

	const { fn, params } = node
	const func =  run(fn, env)

	const fnps = []
	params.forEach(item => {
		fnps.push(run(item, env))
	})

	const result = applyFunction(func, fnps, env)
	return result
}

function applyFunction(fn, params, env) {
	const fnenv = setFnEnv(fn, params, env)

	const { body } = fn
	// console.log('fnenv==', body, fnenv)
	const result = Eval(body.values, fnenv)
	return result
}
function Eval(node, env) {
  // console.log(' Eval ast node', node)
	let res = null
  node.forEach(item => {
    res = run(item, env)
  })
  // console.log('Eval==', env)
	return res
}

module.exports = {
	env,
	Eval
} 