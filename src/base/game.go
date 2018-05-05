package base

import (
	"fmt"
	"./geometry"
	"./shader"
	"../fileUtil"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type Game struct {
	mesh geometry.Mesh
	Shader shader.Shader
}

func NewGame(config fileUtil.Configuration) Game{


	var g Game = Game{}
	g.Shader = shader.InitGraphic(fileUtil.LoadShader(config.SHADER.VERTEX), fileUtil.LoadShader(config.SHADER.FRAGMENT))
	g.Shader.AddUniform("uniformFloat");

	mesh := geometry.CreateMesh()



	vertex1:= geometry.NewVertex(mgl32.Vec3{-1,-1,0})
	vertex2:= geometry.NewVertex(mgl32.Vec3{0,1,0})
	vertex3:= geometry.NewVertex(mgl32.Vec3{1,-1,0})
	data := []geometry.Vertex{ vertex3, vertex1, vertex2 }
	mesh = mesh.AddVertices(data)
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
	g.Shader.SetUniformf("uniformFloat", 2)
}

func (g Game) Render(){
	g.Shader.Bind();
	g.mesh.Draw()
}