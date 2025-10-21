package c

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestType_Render(t *testing.T) {
	type TestCase struct {
		Type   Type `yaml:",inline"`
		Output string
	}

	testCases := loadTestData[TestCase](t)

	for _, testCase := range testCases {
		t.Run(testCase.Output, func(t *testing.T) {
			var b strings.Builder
			testCase.Type.Render(&b)
			assert.Equal(t, testCase.Output, b.String())
		})
	}
}

func TestType_RenderDeclaration(t *testing.T) {
	type TestCase struct {
		Type    Type `yaml:",inline"`
		VarName string
		Output  string
	}

	testCases := loadTestData[TestCase](t)

	for _, testCase := range testCases {
		t.Run(testCase.Output, func(t *testing.T) {
			var b strings.Builder
			testCase.Type.RenderDeclaration(&b, testCase.VarName)
			assert.Equal(t, testCase.Output, b.String())
		})
	}
}
