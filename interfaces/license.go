package interfaces

import (
	"log"

	"github.com/colonelx/ardfrace/helpers"
	"github.com/gotk3/gotk3/gtk"
)

type License struct {
	Application *gtk.Application
}

func (w License) Init() {
	log.Println("application activate")

	// Get the GtkBuilder UI definition in the glade file.
	builder, err := gtk.BuilderNewFromFile("resources/ardfrace.glade")
	helpers.HandleError(err)

	// Map the handlers to callback functions, and connect the signals
	// to the Builder.
	signals := map[string]interface{}{
		"on_main_window_destroy": onMainWindowDestroy,
	}
	builder.ConnectSignals(signals)

	// Get the object with the id of "main_window".
	obj, err := builder.GetObject("dialog")
	helpers.HandleError(err)

	// Verify that the object is a pointer to a gtk.ApplicationWindow.
	win := obj.(*gtk.Dialog)
	helpers.HandleError(err)

	// Show the Window and all of its components.
	win.ShowAll()
}
