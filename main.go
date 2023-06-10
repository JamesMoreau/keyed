package main

import (
	"fmt"
	"log"
	_"strconv"
	"github.com/jroimartin/gocui"
)

// func copy_password_to_clipboard(password string) {
// 	clipboard.WriteAll(password)
// }

// func main() {
// 	pass, err := generate_password(24, true, true, true, true);
// 	if err != nil {
// 		return;
// 	}

// 	out := fmt.Sprintf("password is: %s", pass);
// 	fmt.Println(out);
// }

// var (
// 	length int = 10;
// 	uppercase_letters_checkbox bool = true;
// 	digits_checkbox bool = false;
// 	special_characters_checkbox bool = false;
// )

var (
	checkboxes = []Checkbox{
		{name: "Uppercase Letters", checked: false, position: 0},
		{name: "Digits", checked: false, position: 1},
		{name: "Special Characters", checked: false, position: 2},
	}

	activeCheckboxIndex = 0
)

type Checkbox struct {
	name     string
	checked  bool
	position int
}

func toggleCheckbox(g *gocui.Gui, v *gocui.View) {
	checkboxes[activeCheckboxIndex].checked = !checkboxes[activeCheckboxIndex].checked
	v.Clear()
	if checkboxes[activeCheckboxIndex].checked {
		v.SetCursor(0, 0)
		fmt.Fprint(v, "[x] "+checkboxes[activeCheckboxIndex].name)
	} else {
		v.SetCursor(0, 0)
		fmt.Fprint(v, "[ ] "+checkboxes[activeCheckboxIndex].name)
	}
}

func next_checkbox(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (activeCheckboxIndex + 1) % len(checkboxes);
	fmt.Println("Going from view " + checkboxes[activeCheckboxIndex].name + " to " + checkboxes[nextIndex].name)
	activeCheckboxIndex = nextIndex;
	return nil;
}

// func toggleCheckbox(g *gocui.Gui, v *gocui.View) error {
// 	isChecked = !isChecked
// 	v.Clear()
// 	if isChecked {
// 		v.SetCursor(0, 0)
// 		fmt.Fprint(v, "[x] Checkbox")
// 	} else {
// 		v.SetCursor(0, 0)
// 		fmt.Fprint(v, "[ ] Checkbox")
// 	}
// 	return nil
// }

// func handleMouse(g *gocui.Gui, v *gocui.View) {
// 	ox, oy := v.Origin()
// 	cx, cy := v.Cursor()
// 	x, y := ox+cx, oy+cy
// 	width, _ := v.Size()

// 	if y == 0 && x > 0 && x < width-2 {
// 		// Clicked on the checkbox area
// 		// toggleCheckbox(g, v)
// 	}

// }

func layout(g *gocui.Gui) error {
	max_x, max_y := g.Size();

	// fmt.Println("Length of checkboxes: " + strconv.Itoa(len(checkboxes)));
	for i, checkbox := range checkboxes {
		x0 := 0;
		y0 := max_y/len(checkboxes) * i;
		x1 := max_x - 1;
		y1 := (max_y/len(checkboxes) * (i+1)) - 1;

		v, err := g.SetView(checkbox.name, x0, y0, x1, y1); 
		if err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}

			v.Highlight = true
			v.Frame = true;
			v.Highlight = true
			if checkbox.checked {
				fmt.Fprint(v, "[x] "+checkbox.name)
			} else {
				fmt.Fprint(v, "[ ] "+checkbox.name)
			}
		}
	}

	// if _, err := g.SetCurrentView(checkboxes[activeCheckboxIndex].name); err != nil {
	// 	return err
	// }

	// fmt.Println("hello, sailor!");

	return nil;
}

// func layout(g *gocui.Gui) error {
// 	maxX, maxY := g.Size()
// 	v, err := g.SetView("main", 0, 0, maxX-1, maxY-1)
// 	if err != nil {
// 		if err != gocui.ErrUnknownView {
// 			return err
// 		}

// 		v.Title = "Keyed - A Simple Password Generator"
// 		v.Highlight = true
// 		v.Wrap = true
// 		v.SetCursor(1, 1)
// 		v.Frame = false

// 		fmt.Fprintln(v, "Parameters:")
// 		fmt.Fprintln(v, "\t[X] Length (← or → to increase/decrease)")
// 		fmt.Fprintln(v, "\t[ ] Uppercase Letters")
// 		fmt.Fprintln(v, "\t[ ] Digits")
// 		fmt.Fprintln(v, "\t[ ] Special Characters ('#', '&', '*')")

// 		fmt.Fprintln(v, "Password:")
// 		fmt.Fprintln(v, "\t| ... |")

// 	}

// 	return nil
// }

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	// g.Cursor = true;
	// g.Mouse = true
	g.Highlight = true
	// g.Cursor = true
	g.SelFgColor = gocui.ColorGreen

	g.SetManagerFunc(layout)

	// Keybinds
	err = g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); 
	if err != nil {
		log.Panicln(err);
	}

	err = g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, next_checkbox); 
	if err != nil {
		log.Fatal(err);
	}

	// if err := g.SetKeybinding("checkbox", gocui.MouseLeft, gocui.ModNone, handleMouse); err != nil {
	// 	fmt.Println("Error setting keybinding:", err)
	// 	return
	// }

	err = g.MainLoop()
	if err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
