package object

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

type Environment struct {
	store map[string]Object
	outer *Environment
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	// プログラム自体の評価。必ずここから開始される
	case *ast.Program:
		return evalProgram(node, env)
	// 式文の評価。式自体を再帰的に評価する
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)

	// 式（数値リテラル）の評価。数値オブジェクトを値として返す
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	// 前置式の評価。右項を再帰的に評価した上で前置式としての結果を返す
	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	// 中置式の評価。左右項を再帰的に評価した上で中置式としての結果を返す
	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right)
	return nil
}