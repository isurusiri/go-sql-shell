package gosqlshell

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseExpression(t *testing.T) {
	tests := []struct {
		source string
		ast    *expression
	}{
		{
			source: "2 = 3 AND 4 = 5",
			ast: &expression{
				binary: &binaryExpression{
					a: expression{
						binary: &binaryExpression{
							a: expression{
								literal: &token{"2", numericKind, location{0, 0}},
								kind:    literalKind,
							},
							b: expression{
								literal: &token{"3", numericKind, location{0, 5}},
								kind:    literalKind,
							},
							op: token{"=", symbolKind, location{0, 3}},
						},
						kind: binaryKind,
					},
					b: expression{
						binary: &binaryExpression{
							a: expression{
								literal: &token{"4", numericKind, location{0, 12}},
								kind:    literalKind,
							},
							b: expression{
								literal: &token{"5", numericKind, location{0, 17}},
								kind:    literalKind,
							},
							op: token{"=", symbolKind, location{0, 15}},
						},
						kind: binaryKind,
					},
					op: token{"and", keywordKind, location{0, 8}},
				},
				kind: binaryKind,
			},
		},
	}

	for _, test := range tests {
		fmt.Println("(Parser) Testing: ", test.source)
		tokens, err := lex(test.source)
		assert.Nil(t, err)

		parser := Parser{}
		ast, cursor, ok := parser.parseExpression(tokens, 0, []token{}, 0)
		assert.True(t, ok, err, test.source)
		assert.Equal(t, cursor, uint(len(tokens)))
		assert.Equal(t, ast, test.ast, test.source)
	}
}
