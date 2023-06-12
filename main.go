package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type keyed struct {
	app.Compo
	myState keyedPageState
	temp    string
}

type keyedPageState struct {
	passwordLengthInput             int
	includeUppercaseLettersInput    bool
	includeDigitsInput              bool
	includeSpecialCharactersInput   bool
	excludeAmbiguousCharactersInput bool
	passwordInput                   string
}

func (k *keyed) OnMount(ctx app.Context) {
	ctx.ObserveState("greet-name").Value(&k.temp)
	k.temp = "blah!@"
	k.myState.passwordLengthInput = 16
	k.myState.includeUppercaseLettersInput = true
	k.myState.includeDigitsInput = true
	k.myState.includeSpecialCharactersInput = false
	k.myState.excludeAmbiguousCharactersInput = false
	k.myState.passwordInput = "blah!"
}

func (k *keyed) Render() app.UI {

	return app.Div().Class("my-container").ID("parent").Body(
		app.H1().Class("title").Text("Keyed - A Simple Password Generator"),

		app.Div().Body(
			app.Text(k.temp),
		),

		app.Div().Style("padding-bottom", "10px").Body(

			app.Div().Body(
				app.Text("Length: "),
				app.Input().Type("number").Min(1).Max(30).Value(16).ID("LENGTH_NUMBER").Value(k.myState.passwordLengthInput).Style("margin-right", "10px"),
				app.Input().Type("range").Min(1).Max(30).Value(16).ID("LENGTH_SLIDER").Value(k.myState.passwordLengthInput),
			),
		),

		app.Div().Class("checkbox-container").Body(
			app.Input().Class("checkbox").Type("checkbox").ID("CH0").Checked(k.myState.includeUppercaseLettersInput).OnClick(k.clickIncludeUppercaseLettersInputCheckbox),
			app.Label().For("CH0").Body(
				app.Text("Include Uppercase Letters"),
			),
		),

		app.Div().Class("checkbox-container").Body(
			app.Input().Class("checkbox").Type("checkbox").ID("CH1").Checked(k.myState.includeDigitsInput),
			app.Label().For("CH1").Body(
				app.Text("Include Digits"),
			),
		),

		app.Div().Class("checkbox-container").Body(
			app.Input().Type("checkbox").Type("checkbox").ID("CH2").Checked(k.myState.includeSpecialCharactersInput),
			app.Label().For("CH2").Body(
				app.Text("Include Special Characters ('&', '*', '+', '/', ...)"),
			),
		),

		app.Div().Class("checkbox-container").Body(
			app.Input().Type("checkbox").Type("checkbox").ID("CH3").Checked(k.myState.excludeAmbiguousCharactersInput),
			app.Label().For("CH3").Body(
				app.Text("Exclude Ambiguous Characters ('l', 'L', 'O', '1', ...)"),
			),
		),

		app.Div().Body(
			app.Form().Body(
				app.Label().Body(app.Text("Password: ")),
				app.Input().Type("text").Style("margin-right", "5px").Value(k.myState.passwordInput),
			),
		),

		app.Div().Style("display", "flex").Style("justify-content", "center").Body(
			app.Button().Class("styled-button").OnClick(k.setName).Body(app.Text("Copy")),
			app.Button().Class("styled-button").Body(app.Text("Generate")),
		),
	)
}

func (k *keyed) setName(ctx app.Context, e app.Event) {
	// v := ctx.JSSrc().Get("value").String()
	// fmt.Println(v);
	fmt.Println("hello from handler")
}

func (k *keyed) clickIncludeUppercaseLettersInputCheckbox(ctx app.Context, e app.Event) {
	k.myState.includeUppercaseLettersInput = !k.myState.includeUppercaseLettersInput;
	fmt.Println("Value includeUppercaseLettersInput is now " + strconv.FormatBool(k.myState.includeUppercaseLettersInput))
}

func (k *keyed) blah(ctx app.Context, e app.Event) {
	
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
