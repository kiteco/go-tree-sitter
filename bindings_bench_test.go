package sitter_test

import (
	"testing"

	sitter "github.com/kiteco/go-tree-sitter"
	"github.com/kiteco/go-tree-sitter/javascript"
)

var (
	ResultUint32 uint32
	ResultSymbol sitter.Symbol
	ResultPoint  sitter.Point
	ResultString string
	ResultNode   *sitter.Node
)

func BenchmarkNode(b *testing.B) {
	src := []byte("let a = 1")
	p := sitter.NewParser()
	defer p.Close()
	p.SetLanguage(javascript.GetLanguage())
	tree := p.Parse(src)
	defer tree.Close()

	root := tree.RootNode()
	first := root.Child(0)

	b.Run("Symbol", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ResultSymbol = first.Symbol()
		}
	})

	b.Run("StartByte", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ResultUint32 = first.StartByte()
		}
	})

	b.Run("EndByte", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ResultUint32 = first.EndByte()
		}
	})

	b.Run("StartPoint", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ResultPoint = first.StartPoint()
		}
	})

	b.Run("EndPoint", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ResultPoint = first.EndPoint()
		}
	})

	b.Run("ChildCount", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ResultUint32 = first.ChildCount()
		}
	})

	b.Run("Content", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ResultString = first.Content(src)
		}
	})

	b.Run("Type", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ResultString = first.Type()
		}
	})

	b.Run("Parent", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ResultNode = first.Parent()
		}
	})

	b.Run("Child", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ResultNode = first.Child(0)
		}
	})
}
