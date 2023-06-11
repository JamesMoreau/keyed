package main

import (
	"fmt"
	"log"
	"github.com/jroimartin/gocui"
	"github.com/atotto/clipboard"
)

var (
	// ui_widgets = []interface{} {
	// 	Checkbox{name: "Uppercase Letters", checked: false, position: 0},
	// 	Checkbox{name: "Digits", checked: false, position: 1},
	// 	Checkbox{name: "Special Characters", checked: false, position: 2},
	// 	Button{name: "Copy", onClick: copy_password_to_clipboard, position: 3},
	// }

	password_field string = "MY PASSWORD";
	
	myViews = []string { "v0", "v1" };
	currentViewIndex = 0;
)

type CheckboxWidget struct {
	name     string;
	checked  bool;
	x, y 	 int;
}

func NewCheckboxWidget(name string, x int, y int) *CheckboxWidget {
	return &CheckboxWidget{name: name, checked: false, x: x, y: y};
}

func (w *CheckboxWidget) toggle_checkbox_widget() {
	w.checked = !w.checked;
}

func (w *CheckboxWidget) Layout(g *gocui.Gui) error {
	check := " "
	if w.checked { check = "x"; }
	
	str := fmt.Sprintf("[%s] %s", check, w.name);
	v, err := g.SetView(w.name, w.x, w.y, w.x + len(str) + 2, w.y + 3);
	if err != nil && err != gocui.ErrUnknownView{
		return err;
	}
	v.Clear();

	fmt.Fprint(v, str);
	return nil;
}

type Button struct {
	name string;
	onClick func();
	position int;
}

func copy_password_to_clipboard() {
	clipboard.WriteAll(password_field);
}

func nextView(g *gocui.Gui, v *gocui.View) error {
	fmt.Println("Hello")
	nextViewIndex := (currentViewIndex + 1) % len(myViews);
	nextViewName := myViews[nextViewIndex];

	fmt.Println("Going from view " + v.Name() + " to " + nextViewName);

	if _, err := g.SetCurrentView(nextViewName); err != nil {
		return err;
	}

	currentViewIndex = nextViewIndex;
	return nil
}


// func layout(g *gocui.Gui) error {
// 	max_x, max_y := g.Size();

// 	// fmt.Println("Length of checkboxes: " + strconv.Itoa(len(checkboxes)));
// 	v, err := g.SetView("main", 0, 0, max_x - 1, max_y - 1); 
// 	if err != nil {
// 		if err != gocui.ErrUnknownView {
// 			return err;
// 		}

// 		// Rendering
// 		v.Frame = true;

// 	}

// 	return nil;
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

	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	uppercaseLettersCheckbox := NewCheckboxWidget("v0", 1, 1);
	digitsCheckbox := NewCheckboxWidget("v1", 1, 6);
	// special_characters_checkbox := NewCheckboxWidget("Special Characters", 1, 
	g.SetManager(uppercaseLettersCheckbox, digitsCheckbox);

	// Keybinds
	err = g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); 
	if err != nil {
		log.Panicln(err);
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		log.Panicln(err)
	}


	err = g.MainLoop()
	if err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
