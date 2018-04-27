package geometry

import (
	"github.com/go-gl/mathgl/mgl32"
)

const VERTEX_SIZE = 3;

type Vertex struct {

	position mgl32.Vec3
}

func NewVertex(pos mgl32.Vec3) Vertex{
	var vertex Vertex= Vertex{}
	vertex.position = pos
	return vertex;
}

func (v Vertex) GetPosition()(mgl32.Vec3){
	return v.position
}

func (v Vertex) SetPosition(pos mgl32.Vec3){
	v.position = mgl32.Vec3{pos.X(), pos.Y(), pos.Z()}
}

func CreateFlippedBuffer(vertices []Vertex) ([]float32) {
	buffer := []float32{}

	for _, v := range vertices{
		buffer=append(buffer, v.GetPosition().X())
		buffer=append(buffer, v.GetPosition().Y())
		buffer= append(buffer, v.GetPosition().Z())
	}
	return buffer

}
