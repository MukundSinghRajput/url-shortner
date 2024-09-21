package components

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func Shotner() []g.Node {
	return []g.Node{
		h.Class("min-h-screen bg-gradient-to-br from-purple-400 via-pink-500 to-red-500 flex items-center justify-center px-4 sm:px-6 lg:px-8"),
		h.Div(
			h.Class("max-w-md w-full space-y-8 bg-white bg-opacity-90 p-10 rounded-xl shadow-2xl transition-all duration-500 ease-in-out transform hover:scale-105"),
			h.Div(
				g.Raw(`<h1 class="text-center text-4xl sm:text-5xl font-extrabold text-transparent bg-clip-text bg-gradient-to-r from-purple-600 to-pink-600 mb-2">Shrink &nbsp; <span class="inline-block transform rotate-90">:)</span></h1>`),
				h.P(
					h.Class("text-center text-sm text-gray-600"),
					g.Text("Shorten your URLs with style!"),
				),
			),
			h.Form(
				h.Class("mt-8 space-y-6"),
				h.Action("#"),
				h.Method("POST"),
				h.Div(
					h.Class("rounded-md shadow-sm -space-y-px"),
					h.Div(
						h.Label(
							h.For("url"),
							h.Class("sr-only"),
							g.Text("URL to shorten"),
						),
						h.Input(
							h.ID("url"),
							h.Name("url"),
							h.Required(),
							h.Class("appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-purple-500 focus:border-purple-500 focus:z-10 sm:text-sm"),
							h.Placeholder("Enter the URL here"),
						),
					),
				),
				h.Div(
					h.Button(
						h.Type("submit"),
						h.Class("group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-purple-600 hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-purple-500 transition duration-150 ease-in-out transform hover:-translate-y-1 hover:scale-105 active:scale-95 active:shadow-inner"),
						h.Span(
							h.Class("absolute left-0 inset-y-0 flex items-center pl-3"),
							g.Raw(`
						<svg class="h-5 w-5 text-purple-500 group-hover:text-purple-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                            <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-8.707l-3-3a1 1 0 00-1.414 1.414L10.586 9H7a1 1 0 100 2h3.586l-1.293 1.293a1 1 0 101.414 1.414l3-3a1 1 0 000-1.414z" clip-rule="evenodd" />
                        </svg>`),
						),
						g.Text("Shrink It!"),
					),
				),
			),
			h.Div(
				h.ID("result"),
				h.Class(`mt-4 hidden`),
				h.P(
					h.Class("text-sm text-gray-600 mb-2"),
					g.Text("Your shortened URL:"),
				),
				h.Div(
					h.Class("flex rounded-md shadow-sm"),
					h.Input(
						h.Type("text"),
						h.ReadOnly(),
						h.ID("shortened-url"),
						h.Class("flex-1 min-w-0 block w-full px-3 py-2 rounded-none rounded-l-md bg-gray-100 text-gray-900 sm:text-sm border border-gray-300 focus:outline-none focus:ring-purple-500 focus:border-purple-500"),
						h.Placeholder("https://shrk.in/abc123"),
					),
					g.Raw(`
				<button type="button" onclick="copyToClipboard()" class="inline-flex items-center px-3 rounded-r-md border border-l-0 border-gray-300 bg-purple-600 text-sm font-medium text-white hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-purple-500">
                    Copy
                </button>`),
				),
			),
			h.Div(
				h.ID("error"),
				h.Class("mt-4 hidden"),
				h.P(
					h.Class("text-sm text-red-600"),
					g.Text("Error: Please enter a valid URL."),
				),
			),
			h.Div(
				h.ID("nok"),
				h.Class("mt-4 hidden"),
				h.P(
					h.Class("text-sm text-red-600"),
					g.Text(""),
				),
			),
		),
	}
}
