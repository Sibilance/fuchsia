package c

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenderDeclaration(t *testing.T) {
	type RenderDeclarationTest struct {
		Type    Type `yaml:",inline"`
		VarName string
		Output  string
	}

	testCases := loadTestData[RenderDeclarationTest](t)

	for _, testCase := range testCases {
		t.Run(testCase.Output, func(t *testing.T) {
			var b strings.Builder
			testCase.Type.RenderDeclaration(&b, testCase.VarName)
			assert.Equal(t, testCase.Output, b.String())
		})
	}
}
