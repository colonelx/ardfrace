package main

import (
	"log"
	"os"

	"github.com/colonelx/ardfrace/helpers"
	"github.com/colonelx/ardfrace/interfaces"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const appId = "com.github.gotk3.gotk3-examples.glade"

func main() {
	log.Println("testing")
	// Create a new application.
	application, err := gtk.ApplicationNew(appId, glib.APPLICATION_FLAGS_NONE)
	helpers.HandleError(err)

	// Connect function to application startup event, this is not required.
	application.Connect("startup", func() {
		log.Println("application startup")
	})

	// Connect function to application activate event
	application.Connect("activate", func() {
		window := interfaces.Main{application}
		window.Init()

		windowLincense := interfaces.License{application}
		windowLincense.Init()
	})

	// Connect function to application shutdown event, this is not required.
	application.Connect("shutdown", func() {
		log.Println("application shutdown")
	})

	// Launch the application
	os.Exit(application.Run(os.Args))
}
