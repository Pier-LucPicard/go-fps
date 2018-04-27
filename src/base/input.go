package base
import (
	"github.com/go-gl/glfw/v3.2/glfw"
)
var angle float32;
var position float32
func onKey( w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey){
	if key == glfw.KeyEscape && action==glfw.Press {
		w.SetShouldClose(true)
	}else if key == glfw.KeyLeft{
		angle = angle - 0.1
	}else if key == glfw.KeyRight {
		angle = angle + 0.1
	}else if key == glfw.KeyUp {
		position = position + 0.1
	}else if key == glfw.KeyDown {
		position = position - 0.1
	}
}
