package vue_test

import (
	"testing"

	sitter "github.com/kiteco/go-tree-sitter"
	"github.com/kiteco/go-tree-sitter/vue"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	parser := sitter.NewParser()
	parser.SetLanguage(vue.GetLanguage())

	sourceCode := []byte(`
<template>
  <div class="split"></div>
</template>

<script>
export default {
  name: 'split'
}
</script>`)
	tree := parser.Parse(sourceCode)

	assert.Equal(
		"(component (template_element (start_tag (tag_name)) (text) (element (start_tag (tag_name) (attribute (attribute_name) (quoted_attribute_value (attribute_value)))) (end_tag (tag_name))) (text) (end_tag (tag_name))) (script_element (start_tag (tag_name)) (raw_text) (end_tag (tag_name))))",
		tree.RootNode().String(),
	)
}
