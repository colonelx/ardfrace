package interfaces

import (
	"log"

	"github.com/colonelx/ardfrace/helpers"
	"github.com/gotk3/gotk3/gtk"
)

type Main struct {
	Application *gtk.Application
}

func (w Main) Init() {
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
	obj, err := builder.GetObject("main")
	helpers.HandleError(err)

	// menuObj, err := builder.GetObject("menuQuit")
	// menuQuit, err := helpers.IsMenuItem(menuObj)
	// helpers.HandleError(err)

	// menuQuit.Connect("activate", func() { w.Application.Quit() })

	// Verify that the object is a pointer to a gtk.ApplicationWindow.
	win, err := helpers.IsWindow(obj)
	helpers.HandleError(err)

	// Show the Window and all of its components.
	win.Show()
	w.Application.AddWindow(win)
}
func onMainWindowDestroy() {
	log.Println("onMainWindowDestroy")
}
