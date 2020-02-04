package html_test

import (
	"testing"

	sitter "github.com/kiteco/go-tree-sitter"
	"github.com/kiteco/go-tree-sitter/html"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	parser := sitter.NewParser()
	parser.SetLanguage(html.GetLanguage())

	sourceCode := []byte(`
<html>
	<head></head>
	<body>
	</body>
</html>`)
	tree := parser.Parse(sourceCode)

	assert.Equal(
		"(fragment (text) (element (start_tag (tag_name)) (text) (element (start_tag (tag_name)) (end_tag (tag_name))) (text) (element (start_tag (tag_name)) (text) (end_tag (tag_name))) (text) (end_tag (tag_name))))",
		tree.RootNode().String(),
	)
}
