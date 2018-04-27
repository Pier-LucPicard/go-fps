package base

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

var window *glfw.Window;

func CreateWindow(width, height int, title string)(*glfw.Window){
	var err error;

	window, err = glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(err)
	}

	return window
}

