package evaluator

import (
	"github.com/ktny/monkey/ast"
	"github.com/ktny/monkey/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
