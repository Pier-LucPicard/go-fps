package base

import (
	"fmt"
	"math"
	"./geometry"
	"./shader"
	"./transform"
	"../fileUtil"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type Game struct {
	mesh geometry.Mesh
	Shader shader.Shader
}
var t transform.Transform


func NewGame(config fileUtil.Configuration) Game{


	var g Game = Game{}
	g.Shader = shader.InitGraphic(fileUtil.LoadShader(config.SHADER.VERTEX), fileUtil.LoadShader(config.SHADER.FRAGMENT))
	g.Shader.AddUniform("transform");
	t= transform.NewTransform();
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

var tempAmount float64= 0.0;
var temp float64= 0.0;

func (g Game) Update(){
	InputUpdate()
	temp= temp+0.0001
	tempAmount= math.Sin(temp)

	t= t.SetTranslationFull(float32(tempAmount), 0.0,0.0)
	t=t.SetRotationFull(0,0,float32(tempAmount) * 180)
	t=t.SetScaleFull(float32(tempAmount),float32(tempAmount),0)

}

func (g Game) Render(){
	g.Shader.Bind();
	g.Shader.SetUniformMatrix("transform", t.GetTransformation())
	g.mesh.Draw()
}