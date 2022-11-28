const env = {
  store: {},
  path: null
}

function getIden(env, val) {
  if (!env) return null
  if (env.store[val]) {
    return env.store[val]
  }
  return getIden(env.path, val)
}

function run(node, env) {
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

function Eval(node) {
  // console.log('ast node', node)
  node.forEach(item => {
    run(item, env)
  })
  // console.log(env)
}

module.exports = Eval