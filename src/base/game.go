package base

import (
	"fmt"
	"./geometry"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type Game struct {
	mesh geometry.Mesh
}

func NewGame(program uint32) Game{
	mesh := geometry.CreateMesh()
	var g Game = Game{}


	vertex1:= geometry.NewVertex(mgl32.Vec3{-1,-1,0})
	vertex2:= geometry.NewVertex(mgl32.Vec3{0,1,0})
	vertex3:= geometry.NewVertex(mgl32.Vec3{1,-1,0})
	data := []geometry.Vertex{ vertex2, vertex1, vertex3 }
	mesh = mesh.AddVertices(data, program)
	g.mesh= mesh
	return g
}

func (g Game) Input(){
	if GetKeyDown(glfw.KeyUp){
		fmt.Println("Key Press")
	}

	if GetKeyUp(glfw.KeyUp){
		fmt.Println("Key Realse")
	}
	if GetMouseDown(glfw.MouseButtonLeft) {
		fmt.Print("Left Mouse Button Press (")
		fmt.Print(GetMousePosition())
		fmt.Println(")")
	}

	if GetMouseUp(glfw.MouseButtonLeft) {
		fmt.Print("Left Mouse Button Relase (")
		fmt.Print(GetMousePosition())
		fmt.Println(")")
	}
}

func (g Game) Update(){
	InputUpdate()

}

func (g Game) Render(){

	g.mesh.Draw()
}