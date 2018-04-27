package geometry

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Mesh struct{
	size int
	vbo uint32
	buffer []float32
}

func CreateMesh() Mesh{
	var vbo uint32

	gl.GenBuffers(1, &vbo)
	return Mesh{vbo: vbo, size:0, buffer: []float32{}}
}

func (m Mesh) AddVertices( vertices []Vertex, program uint32) Mesh{



	m.buffer = CreateFlippedBuffer(vertices)
	points := CreateFlippedBuffer(vertices)
	m.size = len(m.buffer) //* VERTEX_SIZE

    var vbo uint32
    gl.GenBuffers(1, &vbo)
    gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
    gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)
    
    var vao uint32
    gl.GenVertexArrays(1, &vao)
    gl.BindVertexArray(vao)

	m.vbo=vbo
	 return m
}

func (m Mesh) Draw(){
	gl.EnableVertexAttribArray(0)

	gl.BindBuffer(gl.ARRAY_BUFFER, m.vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

    gl.DrawArrays(gl.TRIANGLES, 0, int32(m.size))
	gl.DisableVertexAttribArray(0)

}