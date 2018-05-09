package geometry

import (

	"strings"
	"strconv"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
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

func (m *Mesh) AddVertices( vertices []Vertex, indices []int32){



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
	
}

func (m *Mesh) Draw(){
	gl.EnableVertexAttribArray(0)

	gl.BindBuffer(gl.ARRAY_BUFFER, m.vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, VERTEX_SIZE*4, nil)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.ibo)
	gl.DrawElements(gl.TRIANGLES, int32(m.size), gl.UNSIGNED_INT, nil)

	gl.DisableVertexAttribArray(0)

}

func ParseObj(rawObj string) Mesh { 

	mesh:=CreateMesh()
	var indices []int32;
	var vertices []Vertex;

	fileLines := strings.Split(rawObj, "\n");

	
	for _,l:= range fileLines {

		if len(l) == 0 {
			continue;
		} 
		tokens := strings.Split(l, " ");
		tokens = removeEmptyString(tokens);
		
		if len(tokens) == 0 {
			continue;
		} else if tokens[0] == "#" || tokens[0] == "vt"{
			continue;
		}


		if tokens[0] == "v" {
			x,_:=strconv.ParseFloat(tokens[1],64)
			y,_:=strconv.ParseFloat(tokens[2],64)
			z,_:=strconv.ParseFloat(tokens[3],64)
			vertices = append(vertices, NewVertex(mgl32.Vec3{float32(x),float32(y),float32(z)}))
			
		}


		if tokens[0] == "f" {
			
			if len(tokens) == 4 {
				microTokensX:=strings.Split(tokens[1], "/");
				microTokensY:=strings.Split(tokens[2], "/");
				microTokensZ:=strings.Split(tokens[3], "/");
	
				x,_:=strconv.ParseInt(microTokensX[0],10,64)
				y,_:=strconv.ParseInt(microTokensY[0],10,64)
				z,_:=strconv.ParseInt(microTokensZ[0],10,64)
				indices = append(indices, int32(x)-1)
				indices = append(indices, int32(y)-1)
				indices = append(indices, int32(z)-1)
			}else if len(tokens) == 5 {
				microTokensX:=strings.Split(tokens[1], "/");
				microTokensY:=strings.Split(tokens[2], "/");
				microTokensZ:=strings.Split(tokens[3], "/");
				microTokensW:=strings.Split(tokens[4], "/");
			
				x,_:=strconv.ParseInt(microTokensX[0],10,64)
				y,_:=strconv.ParseInt(microTokensY[0],10,64)
				z,_:=strconv.ParseInt(microTokensZ[0],10,64)
				w,_:=strconv.ParseInt(microTokensW[0],10,64)
				indices = append(indices, int32(x)-1)
				indices = append(indices, int32(y)-1)
				indices = append(indices, int32(w)-1)

				indices = append(indices, int32(y)-1)
				indices = append(indices, int32(z)-1)
				indices = append(indices, int32(w)-1)
			}

		}
	}

	mesh.AddVertices(vertices, indices)

	return mesh
}

func removeEmptyString( items []string) (res []string){
	for _,i:= range items {
		if i != "" {
			res = append(res, i);
		}
	}

	return
}