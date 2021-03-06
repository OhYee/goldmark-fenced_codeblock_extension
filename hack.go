package ext

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
)

type hack struct {
	target   ast.NodeKind
	receiver *renderer.NodeRendererFunc
}

func (h hack) Register(node ast.NodeKind, f renderer.NodeRendererFunc) {
	if node.String() == h.target.String() {
		*h.receiver = f
	}
}

// GetRenderFunc get the render function from a goldmark renderer
func GetRenderFunc(target ast.NodeKind, r renderer.NodeRenderer) renderer.NodeRendererFunc {
	var receiver renderer.NodeRendererFunc
	h := hack{target, &receiver}
	r.RegisterFuncs(h)
	return receiver
}

// GetFencedCodeBlockRendererFunc get the fenced codeblock render function from a goldmark renderer
func GetFencedCodeBlockRendererFunc(r renderer.NodeRenderer) renderer.NodeRendererFunc {
	return GetRenderFunc(ast.KindFencedCodeBlock, r)
}
