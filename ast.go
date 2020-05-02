package gosqlshell

import (
	"fmt"
)

type expressionKind uint

const (
	literalKind expressionKind = iota
	binaryKind
)

type binaryExpression struct {
	a  expression
	b  expression
	op token
}

func (be binaryExpression) generateCode() string {
	return fmt.Sprintf("(%s %s %s)", be.a.generateCode(), be.op.value, be.b.generateCode())
}

type expression struct {
	literal *token
	binary  *binaryExpression
	kind    expressionKind
}

func (e expression) generateCode() string {
	switch e.kind {
	case literalKind:
		switch e.literal.kind {
		case identifierKind:
			return fmt.Sprintf("\"%s\"", e.literal.value)
		case stringKind:
			return fmt.Sprintf("'%s'", e.literal.value)
		default:
			return fmt.Sprintf(e.literal.value)
		}

	case binaryKind:
		return e.binary.generateCode()
	}

	return ""
}
