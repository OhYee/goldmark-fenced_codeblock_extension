package ext

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

var defaultRenderer = GetFencedCodeBlockRendererFunc(html.NewRenderer())

// RenderMap languages and these renderer
type RenderMap struct {
	Language       []string // language name, for all language use "*"
	RenderFunction renderer.NodeRendererFunc
}

type ext struct {
	Map []RenderMap
}

// Ext default Extension
var Ext = NewExt()

// NewExt initial an extension for goldmark fenced codeblock
func NewExt(maps ...RenderMap) goldmark.Extender {
	return &ext{Map: maps}
}

// AddLanguage add a new fenced codeblock language renderer
func (e *ext) AddLanguage(m RenderMap) {
	e.Map = append(e.Map, m)
}

// Extend implement of goldmark.Extender
func (e *ext) Extend(m goldmark.Markdown) {
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(e, 0),
	))
}

// RegisterFuncs implement of goldmark.renderer.NodeRenderer
func (e *ext) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindFencedCodeBlock, e.RenderFencedCodeBlock)
}

// RenderFencedCodeBlock render codeblock
func (e *ext) RenderFencedCodeBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.FencedCodeBlock)
	language := string(n.Language(source))

	for _, m := range e.Map {
		for _, lang := range m.Language {
			if lang == language || lang == "*" {
				return m.RenderFunction(w, source, node, entering)
			}
		}
	}
	return defaultRenderer(w, source, node, entering)
}
