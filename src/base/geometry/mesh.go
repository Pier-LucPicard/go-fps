package geometry

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Mesh struct{
	size int
	vbo uint32
	vao uint32
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
    gl.EnableVertexAttribArray(0)
    gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	
	 m.vao = vao
	m.vbo=vbo
	 return m
}

func (m Mesh) Draw(){



	gl.BindVertexArray(m.vao)
    gl.DrawArrays(gl.TRIANGLES, 0, int32(len(m.buffer) / 3))


}