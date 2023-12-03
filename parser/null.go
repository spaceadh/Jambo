package parser

import (
	"github.com/spaceadh/Jambo/ast"
)

func (p *Parser) parseNull() ast.Expression {
	return &ast.Null{Token: p.curToken}
}
