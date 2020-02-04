package css_test

import (
	"testing"

	sitter "github.com/kiteco/go-tree-sitter"
	"github.com/kiteco/go-tree-sitter/css"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	parser := sitter.NewParser()
	parser.SetLanguage(css.GetLanguage())

	sourceCode := []byte(`
div {
	background-color: #010101;
}
`)
	tree := parser.Parse(sourceCode)

	assert.Equal(
		"(stylesheet (rule_set (selectors (tag_name)) (block (declaration (property_name) (color_value)))))",
		tree.RootNode().String(),
	)
}
