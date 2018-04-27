package base

import (
	"fmt"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Game struct {
}

func (g Game) Input(){
	if GetKeyDown(glfw.KeyUp){
		fmt.Println("Key Press")
	}
	if GetKeyUp(glfw.KeyUp){
		fmt.Println("Key Realse")
	}
	if GetMouseDown(glfw.MouseButtonLeft) {
		fmt.Println("Left Mouse Button Press")
	}

	if GetMouseUp(glfw.MouseButtonLeft) {
		fmt.Println("Left Mouse Button Relase")
	}
}

func (g Game) Update(){
	InputUpdate()

}

func (g Game) Render(){

}