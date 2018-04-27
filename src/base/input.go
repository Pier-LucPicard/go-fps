package base

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

const NUM_KEYCODES = 256


var	downKeys map[glfw.Key]int = make(map[glfw.Key]int)
var	currentKeys map[glfw.Key]int = make(map[glfw.Key]int)
var	upKeys map[glfw.Key]int = make(map[glfw.Key]int)

var	downMouse map[glfw.MouseButton]int = make(map[glfw.MouseButton]int)
var	currentMouse map[glfw.MouseButton]int = make(map[glfw.MouseButton]int)
var	upMouse map[glfw.MouseButton]int = make(map[glfw.MouseButton]int)

func InputUpdate(){

	//Keyboard Update
	for key, value := range downKeys {
		if value == 1 {
			currentKeys[key] = 1
		} 
	}

	for key, value := range upKeys {
		if value == 1 {
			currentKeys[key] = 0
		}
	}

	for key, _ := range upKeys {
		upKeys[key] = 0
	}

	for key, _ := range downKeys {
		downKeys[key] = 0
	}


	//Mouse Update

	for key, value := range downMouse {
		if value == 1 {
			currentMouse[key] = 1
		} 
	}
	
	for key, value := range upMouse {
		if value == 1 {
			currentMouse[key] = 0
		}
	}

	for key, _ := range upMouse {
		upMouse[key] = 0
	}

	for key, _ := range downMouse {
		downMouse[key] = 0
	}

}


func GetKeyDown(key glfw.Key) bool{
	return downKeys[key] == 1
}

func GetKeyUp(key glfw.Key) bool {
	return upKeys[key] == 1
}

func onMouse( w *glfw.Window, button glfw.MouseButton,  action glfw.Action, mods glfw.ModifierKey){


	if action == glfw.Release {
		upMouse[button] = 1
	} 

	if action == glfw.Press {
		downMouse[button] = 1
	}

}
func onKey( w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey){
	if key == glfw.KeyEscape && action==glfw.Press {
		w.SetShouldClose(true)
	}

	if action == glfw.Release {
		upKeys[key] = 1
	} 

	if action == glfw.Press {
		downKeys[key] = 1
	}

}

func GetMouseDown(btn glfw.MouseButton) bool{
	return downMouse[btn] == 1
}

func GetMouseUp(btn glfw.MouseButton) bool {
	return upMouse[btn] == 1
}

func GetMousePosition()(x, y float64){
	return window.GetCursorPos()
}

