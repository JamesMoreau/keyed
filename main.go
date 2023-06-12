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
		Body(
			app.H1().Class("title").Text("Keyed - A Simple Password Generator"),

			app.Div().Class("checkbox-container").Body(
				app.Input().Class("checkbox").Type("checkbox").ID("CH0"),
				app.Label().For("CH0").Body(
					app.Text("Include Uppercase Letters"),
				),
			),

			app.Div().Class("checkbox-container").Body(
				app.Input().Class("checkbox").Type("checkbox").ID("CH1"),
				app.Label().For("CH1").Body(
					app.Text("Include Digits"),
				),
			),

			app.Div().Class("checkbox-container").Body(
				app.Input().Type("checkbox").Type("checkbox").ID("CH2"),
				app.Label().For("CH2").Body(
					app.Text("Include Special Characters ('&', '*', '+', '/')"),
				),
			),

			app.Div().Class("checkbox-container").Body(
				app.Input().Type("checkbox").Type("checkbox").ID("CH3"),
				app.Label().For("CH3").Body(
					app.Text("Exclude Ambiguous Characters ('l', 'L', 'O', '1')"),
				),
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
