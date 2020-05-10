# goldmark-fenced_codeblock_extension

[![Sync to Gitee](https://github.com/OhYee/goldmark-fenced_codeblock_extension/workflows/Sync%20to%20Gitee/badge.svg)](https://gitee.com/OhYee/goldmark-fenced_codeblock_extension)  [![version](https://img.shields.io/github/v/tag/OhYee/goldmark-fenced_codeblock_extension)](https://github.com/OhYee/goldmark-fenced_codeblock_extension/tags)

A extension for [goldmark](http://github.com/yuin/goldmark) to enhance fenced codeblock


```go
ext.RenderMap{
    Language:       []string{"red", "green"},
    RenderFunction: Renderer,
}
```

```go
func Renderer(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*ast.FencedCodeBlock)
	language := string(n.Language(source))

    buf := bytes.NewBuffer([]byte{})
	l := n.Lines().Len()
	for i := 0; i < l; i++ {
		line := n.Lines().At(i)
		buf.Write(line.Value(source))
	}
    raw := buf.String()

    switch language {
        case "red":
            w.WriteString(fmt.Sprintf("<span style=\"color:red\">%s</span>", raw))
        case "green":
            w.WriteString(fmt.Sprintf("<span style=\"color:green\">%s</span>", raw))
        default: 
            w.WriteString(fmt.Sprintf("<span>%s</span>", raw))
    }
       
	return ast.WalkContinue, nil
}
```

Then, you can use this in markdown


```markdown
　```red
　RED
　```
```

and get `<span style="color:red">RED</span>`

