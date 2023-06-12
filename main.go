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
}

var elementIds = []string { "LENGTH_NUMBER" , "CH0", "CH1", "CH2", "CH3", "PASS"};

func (k *keyed) Render() app.UI {

	return app.Div().Class("my-container").ID("parent").Body(
		app.H1().Class("title").Text("Keyed - A Simple Password Generator"),

		app.Div().Style("padding-bottom", "10px").Body(

			app.Div().Body(
				app.Text("Length: "),
				app.Input().Type("number").Min(1).Max(30).Value(16).ID(elementIds[0]).Style("margin-right", "10px"),
				// app.Input().Type("range").Min(1).Max(30).Value(16).ID("LENGTH_SLIDER"),
			),
		),

		app.Div().Class("checkbox-container").Body(
			app.Input().Class("checkbox").Type("checkbox").ID(elementIds[1]).Checked(true),
			app.Label().For("CH0").Body(
				app.Text("Include Uppercase Letters"),
			),
		),

		app.Div().Class("checkbox-container").Body(
			app.Input().Class("checkbox").Type("checkbox").ID(elementIds[2]).Checked(true),
			app.Label().For("CH1").Body(
				app.Text("Include Digits"),
			),
		),

		app.Div().Class("checkbox-container").Body(
			app.Input().Type("checkbox").Type("checkbox").ID(elementIds[3]),
			app.Label().For("CH2").Body(
				app.Text("Include Special Characters ('&', '*', '+', '/', ...)"),
			),
		),

		app.Div().Class("checkbox-container").Body(
			app.Input().Type("checkbox").Type("checkbox").ID(elementIds[4]),
			app.Label().For("CH3").Body(
				app.Text("Exclude Ambiguous Characters ('l', 'L', 'O', '1', ...)"),
			),
		),

		app.Div().Body(
			app.Form().Body(
				app.Label().Body(app.Text("Password: ")),
				app.Input().Type("text").Style("margin-right", "5px").Value("blah!").ID(elementIds[5]),
			),
		),

		app.Div().Style("display", "flex").Style("justify-content", "center").Body(
			app.Button().Class("styled-button").OnClick(k.copyPassword).Body(app.Text("Copy")),
			app.Button().Class("styled-button").OnClick(k.generatePassword).Body(app.Text("Generate")),
		),
	)
}

func (k *keyed) generatePassword(ctx app.Context, e app.Event) {
	lengthStr := app.Window().GetElementByID("LENGTH_NUMBER").Get("value").String()
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		fmt.Println("Keyed: Unable to convert length string to integer.")
	}

	includeUppercaseLettersInput := app.Window().GetElementByID("CH0").Get("checked").Bool()
	includeDigitsInput := app.Window().GetElementByID("CH1").Get("checked").Bool()
	includeSpeciaCharactersInput := app.Window().GetElementByID("CH2").Get("checked").Bool()
	excludeAmbiguousCharactersInput := app.Window().GetElementByID("CH3").Get("checked").Bool()

	// fmt.Println(length)
	// fmt.Println(strconv.FormatBool(includeUppercaseLettersInput))
	// fmt.Println(strconv.FormatBool(includeDigitsInput))
	// fmt.Println(strconv.FormatBool(includeSpeciaCharactersInput))
	// fmt.Println(strconv.FormatBool(excludeAmbiguousCharactersInput))

	password, err := generate_password(length, includeUppercaseLettersInput, includeDigitsInput, includeSpeciaCharactersInput, excludeAmbiguousCharactersInput);
	if err != nil {
		app.Window().Call("alert", "Unable to generate a password. :(");
	}

	app.Window().GetElementByID("PASS").Set("value", password);
}

func (k *keyed) copyPassword(ctx app.Context, e app.Event) {
	password := app.Window().GetElementByID("PASS").Get("value").String();
	// fmt.Println("Writing " + password + " to clipboard.")
	app.Window().Call("copyToClipboard", password);
	// app.Value.Call()
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
		RawHeaders: []string{
			`
			<script>
			function copyToClipboard(text) {
				navigator.clipboard.writeText(text);
			}
		  	</script>
			`,
		},
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
