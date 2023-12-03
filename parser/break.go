package parser

import (
	"github.com/spaceadh/Jambo/ast"
	"github.com/spaceadh/Jambo/token"
)

func (p *Parser) parseBreak() *ast.Break {
	stmt := &ast.Break{Token: p.curToken}
	for p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
