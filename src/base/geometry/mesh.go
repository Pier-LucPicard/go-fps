package geometry

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Mesh struct{
	size int
	vbo uint32
	ibo uint32
	buffer []float32
}

func CreateMesh() Mesh{
	var vbo uint32
	var ibo uint32

	gl.GenBuffers(1, &vbo)
	gl.GenBuffers(1, &ibo)
	return Mesh{vbo: vbo, size:0, buffer: []float32{}, ibo: ibo}
}

func (m Mesh) AddVertices( vertices []Vertex, indices []int32) Mesh{



	m.buffer = CreateFlippedBuffer(vertices)
	points := CreateFlippedBuffer(vertices)


	m.size = len(indices)// * VERTEX_SIZE

	var vbo uint32
	var ibo uint32

    gl.GenBuffers(1, &vbo)
    gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)
	
    var vao uint32
    gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)	
	
    gl.GenBuffers(1, &ibo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ibo)
    gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(indices), gl.Ptr(indices), gl.STATIC_DRAW)

	m.vbo=vbo
	m.ibo=ibo
	 return m
}

func (m Mesh) Draw(){
	gl.EnableVertexAttribArray(0)

	gl.BindBuffer(gl.ARRAY_BUFFER, m.vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, VERTEX_SIZE*4, nil)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.ibo)
	gl.DrawElements(gl.TRIANGLES, int32(m.size), gl.UNSIGNED_INT, nil)

	gl.DisableVertexAttribArray(0)

}