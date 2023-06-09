package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"log"
	"github.com/jroimartin/gocui"
)

func copy_password_to_clipboard(password string) {
	clipboard.WriteAll(password);	
}

// func main() {
// 	pass, err := generate_password(24, true, true, true, true);
// 	if err != nil {
// 		return;
// 	}

// 	out := fmt.Sprintf("password is: %s", pass);
// 	fmt.Println(out);
// }

var isChecked bool

func toggleCheckbox(g *gocui.Gui, v *gocui.View) error {
	isChecked = !isChecked
	v.Clear()
	if isChecked {
		v.SetCursor(0, 0)
		fmt.Fprint(v, "[x] Checkbox")
	} else {
		v.SetCursor(0, 0)
		fmt.Fprint(v, "[ ] Checkbox")
	}
	return nil
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	// if v, err := g.SetView("main", 0, 0, maxX - 1, maxY - 1); err != nil {
	// 	if err != gocui.ErrUnknownView {
	// 		return err
	// 	}
	// 	fmt.Fprintln(v, "Hello world!")
	// }
	if v, err := g.SetView("checkbox", maxX/2-10, maxY/2, maxX/2+10, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		
		v.Title = "Checkbox Example"
		v.Highlight = true
		v.Wrap = true
		v.SetCursor(0, 0)
		fmt.Fprint(v, "[ ] Checkbox")
		if _, err := g.SetCurrentView("checkbox"); err != nil {
			return err
		}
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("checkbox", gocui.KeyEnter, gocui.ModNone, toggleCheckbox); err != nil {
		fmt.Println("Error setting keybinding:", err)
		return
	}

	err = g.MainLoop();
	if err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
