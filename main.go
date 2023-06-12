package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type keyed struct {
	app.Compo
}

type keyedPageState struct {
	passwordLengthInput             int
	includeUppercaseLettersInput    bool
	includeDigitsInput              bool
	includeSpecialCharactersInput   bool
	excludeAmbiguousCharactersInput bool
	passwordInput                   string
}

var (
	myPageState = keyedPageState{
		passwordLengthInput:             16,
		includeUppercaseLettersInput:    true,
		includeDigitsInput:              true,
		includeSpecialCharactersInput:   false,
		excludeAmbiguousCharactersInput: false,
		passwordInput:                   ""}
)

func (k *keyed) Render() app.UI {
	return app.Div().Class("my-container").ID("parent").Body(
		app.H1().Class("title").Text("Keyed - A Simple Password Generator"),

		app.Div().Style("padding-bottom", "10px").Body(

			app.Div().Body(
				app.Text("Length: "),
				app.Input().Type("number").Min(1).Max(30).Value(16).ID("LENGTH_NUMBER").Style("margin-right", "10px"),
				app.Input().Type("range").Min(1).Max(30).Value(16).ID("LENGTH_SLIDER"),
			),
		),

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
				app.Text("Include Special Characters ('&', '*', '+', '/', ...)"),
			),
		),

		app.Div().Class("checkbox-container").Body(
			app.Input().Type("checkbox").Type("checkbox").ID("CH3"),
			app.Label().For("CH3").Body(
				app.Text("Exclude Ambiguous Characters ('l', 'L', 'O', '1', ...)"),
			),
		),

		app.Div().Body(
			app.Form().Body(
				app.Label().Body(app.Text("Password: ")),
				app.Input().Type("text").Style("margin-right", "5px"),
			),
		),
		
		app.Div().Style("display", "flex").Style("justify-content", "center").Body(
			app.Button().Class("styled-button").OnClick(k.doNothing).Body(app.Text("Copy")),
			app.Button().Class("styled-button").Body(app.Text("Generate")),
		),
	)
}

func (k *keyed) doNothing(ctx app.Context, e app.Event) {
	// v := ctx.JSSrc().Get("value").String()
	// fmt.Println(v);
	fmt.Println("hello from handler");
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
			"/web/styles.css",
		},
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}