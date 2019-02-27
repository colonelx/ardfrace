package helpers

import (
	"errors"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func IsWindow(obj glib.IObject) (*gtk.Window, error) {
	// Make type assertion (as per gtk.go).
	if win, ok := obj.(*gtk.Window); ok {
		return win, nil
	}
	return nil, errors.New("not a *gtk.Window")
}

func IsMenuItem(obj glib.IObject) (*gtk.MenuItem, error) {
	if menuItem, ok := obj.(*gtk.MenuItem); ok {
		return menuItem, nil
	}
	return nil, errors.New("Not a *gtk.MenuItem")
}
func IsApplicationWindow(obj glib.IObject) (*gtk.ApplicationWindow, error) {
	// Make type assertion (as per gtk.go).
	if win, ok := obj.(*gtk.ApplicationWindow); ok {
		return win, nil
	}
	return nil, errors.New("not a *gtk.ApplicationWindow")
}
