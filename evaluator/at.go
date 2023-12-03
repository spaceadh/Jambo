package evaluator

import (
	"github.com/spaceadh/Jambo/ast"
	"github.com/spaceadh/Jambo/object"
)

func evalAt(node *ast.At, env *object.Environment) object.Object {
	if at, ok := env.Get("@"); ok {
		return at
	}
	return newError("Iko nje ya scope")
}
