package ext

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

var defaultRenderer = GetFencedCodeBlockRendererFunc(html.NewRenderer())

type RenderMap struct {
	Language       []string
	RenderFunction renderer.NodeRendererFunc
}

type ext struct {
	Map []RenderMap
}

var Ext = NewExt()

func NewExt(maps ...RenderMap) goldmark.Extender {
	return &ext{Map: maps}
}

func (e *ext) Extend(m goldmark.Markdown) {
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(e, 0),
	))
}

func (e *ext) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindFencedCodeBlock, e.RenderFencedCodeBlock)
}

func (e *ext) RenderFencedCodeBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.FencedCodeBlock)
	language := string(n.Language(source))

	for _, m := range e.Map {
		for _, lang := range m.Language {
			if lang == language {
				return m.RenderFunction(w, source, node, entering)
			}
		}
	}
	return defaultRenderer(w, source, node, entering)
}
