package main

import (
	"fmt"
	"runtime"
)

// Abstract factory

type WidgetFactory interface {
	CreateWindow() Window
	CreateButton() Button
	CreateMenu() Menu
}

type Window interface {
	Create()
}

type Button interface {
	Create()
}

type Menu interface {
	Create()
}

// Windows factory

type WinFactory struct{}

type WinWindow struct{}

func (w *WinWindow) Create() {
	fmt.Println("Windows window created")
}

type WinButton struct{}

func (w *WinButton) Create() {
	fmt.Println("Windows buttom created")
}

type WinMenu struct{}

func (w *WinMenu) Create() {
	fmt.Println("Windows menu created")
}

func (w *WinFactory) CreateWindow() Window {
	return &WinWindow{}
}
func (w *WinFactory) CreateButton() Button {
	return &WinButton{}
}

func (w *WinFactory) CreateMenu() Menu {
	return &WinMenu{}
}

// Mac Factory

type MacFactory struct{}

type MacWindow struct{}

func (w *MacWindow) Create() {
	fmt.Println("Mac window created")
}

type MacButton struct{}

func (w *MacButton) Create() {
	fmt.Println("Mac buttom created")
}

type MacMenu struct{}

func (w *MacMenu) Create() {
	fmt.Println("Mac menu created")
}

func (m *MacFactory) CreateWindow() Window {
	return &MacWindow{}
}

func (m *MacFactory) CreateButton() Button {
	return &MacButton{}
}

func (m *MacFactory) CreateMenu() Menu {
	return &MacMenu{}
}

func main() {
	var factory WidgetFactory
	if runtime.GOOS == "windows" {
		factory = &WinFactory{}
	} else {
		factory = &MacFactory{}
	}
	window := factory.CreateWindow()
	button := factory.CreateButton()
	menu := factory.CreateMenu()

	window.Create()
	button.Create()
	menu.Create()

}
