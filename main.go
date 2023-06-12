package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"log"
	"net/http"
)

type keyed struct {
	app.Compo
}

func (k *keyed) Render() app.UI {
	return app.Div().
	Class("container").
		Body(
			app.H1().
				Class("title").
				Text("Keyed - A Simple Password Generator"),
			app.Div().
				Class("checkbox").
				Body(
					app.Input().
						Type("checkbox"),
					app.Text("Include Uppercase Letters"),
				),
			app.Div().
				Class("checkbox").
				Body(
					app.Input().
						Type("checkbox"),
					app.Text("Include Digits"),
				),
			app.Div().Body(
				app.Input().
					Type("checkbox"),
				app.Text("Include Special Characters ('&', '*', '+', '/') BLAH"),
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
			),
			app.Div().Body(
				app.Button().Body(app.Text("Generate")),
				app.Button().Body(app.Text("Copy")),
			),
		)
}

func main() {
	app.Route("/", &keyed{})

	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Author:      "James Moreau",
		Version:     "1.0",
		Name:        "Keyed",
		Description: "A Simple Password Generator",
		Styles: []string{
			// "https://unpkg.com/spectre.css",
			"/web/styles.css",
		},
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
