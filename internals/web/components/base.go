package components

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
)

func Base(title string, body []g.Node) g.Node {
	return c.HTML5(
		c.HTML5Props{
			Title:    title,
			Language: "en",
			Head: []g.Node{
				h.Meta(h.Charset("UTF-8")),
				h.Meta(h.Name("viewport"), h.Content("width=device-width, initial-scale=1.0")),
				h.Link(h.Rel("stylesheet"), h.Href("/assets/css/base.min.css")),
				h.Link(h.Rel("stylesheet"), h.Href("/assets/css/components.min.css")),
				h.Link(h.Rel("stylesheet"), h.Href("/assets/css/typography.min.css")),
				h.Link(h.Rel("stylesheet"), h.Href("/assets/css/utilities.min.css")),
				h.Link(h.Rel("stylesheet"), h.Href("/assets/css/style.css")),
				h.Script(h.Src("/assets/js/script.js")),
			},
			Body: body,
		},
	)
}
