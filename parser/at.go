package parser

import "github.com/spaceadh/Jambo/ast"

func (p *Parser) parseAt() ast.Expression {
	return &ast.At{Token: p.curToken}
}
