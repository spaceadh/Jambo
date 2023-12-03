package evaluator

import (
	"github.com/spaceadh/Jambo/ast"
	"github.com/spaceadh/Jambo/object"
)

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}
	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}

	return newError("Mstari %d: Neno Halifahamiki: %s", node.Token.Line, node.Value)
}