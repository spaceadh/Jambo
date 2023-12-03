package parser

import (
	"github.com/spaceadh/Jambo/ast"
	"github.com/spaceadh/Jambo/token"
)

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.curToken, Value: p.curTokenIs(token.TRUE)}
}
