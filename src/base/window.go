package base

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

var window *glfw.Window;

func CreateWindow(width, height int, title string)(*glfw.Window){
	var err error;

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err = glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(err)
	}

	return window
}

