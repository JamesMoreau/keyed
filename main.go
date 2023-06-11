package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"log"
	"net/http"
)

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type hello struct {
	app.Compo
}

func (h *hello) Render() app.UI {
	return app.Div().Body(
		app.H1().
			Class("title").
			Text("Keyed - A Simple Password Generator"),
		app.Div().Body(
			app.Input().
				Type("checkbox"),
			app.Text("Include Uppercase Letters"),
		),
		app.Div().Body(
			app.Input().
				Type("checkbox"),
			app.Text("Include Digits"),
		),
		app.Div().Body(
			app.Input().
				Type("checkbox"),
			app.Text("Include Special Characters ('&', '*', '+', '/')"),
		),
		app.Div().Body(
			app.Input().
				Type("checkbox"),
			app.Text("Exclude Ambiguous Characters ('l', 'L', 'O', '1', ...)"),
		),
		app.Div().Body(
			app.Form().Body(
				app.Label().Body(app.Text("Password")),
				app.Input().
					Type("text"),
			),
			app.Button().Body(app.Text("Copy")),
		),
	)
}

func main() {
	app.Route("/", &hello{})

	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
