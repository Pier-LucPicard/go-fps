package main;

import (
	"./fileUtil"
	"./base"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const WINDOW_WIDTH = 800;
const WINDOW_HEIGHT = 600;

var model mgl32.Mat4;
var camera mgl32.Mat4;
var modelUniform int32;
var cameraUniform int32;
var window *glfw.Window;



var angle float32 = 0
var position float32 =0

func main() {


	//var err error	
	//Load the configuration

	// //Load the shaders
	// vertexShader:=fileUtil.LoadShader(config.SHADER.VERTEX)
	// fragmentShader:=fileUtil.LoadShader(config.SHADER.FRAGMENT)

	// cubeVertices:=geometry.NewCube()

	game := base.CreateCoreEngine(fileUtil.LoadConfig())
	game.Start()

	// window = game.Window




	// // Configure the vertex and fragment shaders
	// program, err := newProgram(vertexShader, fragmentShader)
	// if err != nil {
	// 	panic(err)
	// }
	
	// gl.UseProgram(program)

	// projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(config.WINDOW_WIDTH)/float32(config.WINDOW_HEIGHT), 0.1, 10.0)
	// projectionUniform := gl.GetUniformLocation(program, gl.Str("projection\x00"))
	// gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	// camera = mgl32.LookAtV(mgl32.Vec3{3, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	// cameraUniform = gl.GetUniformLocation(program, gl.Str("camera\x00"))
	// gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])


	// model = mgl32.Ident4()
	// modelUniform = gl.GetUniformLocation(program, gl.Str("model\x00"))
	// gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

	// textureUniform := gl.GetUniformLocation(program, gl.Str("tex\x00"))
	// gl.Uniform1i(textureUniform, 0)

	// gl.BindFragDataLocation(program, 0, gl.Str("outputColor\x00"))


	// var vao uint32
	// gl.GenVertexArrays(1, &vao)
	// gl.BindVertexArray(vao)

	// var vbo uint32
	// gl.GenBuffers(1, &vbo)
	// gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	// gl.BufferData(gl.ARRAY_BUFFER, len(cubeVertices)*4, gl.Ptr(cubeVertices), gl.STATIC_DRAW)


	// vertAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vert\x00")))
	// gl.EnableVertexAttribArray(vertAttrib)
	// gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))
	
	// texCoordAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertTexCoord\x00")))
	// gl.EnableVertexAttribArray(texCoordAttrib)
	// gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))



	// fmt.Println("Press left and right arrow to rotate the cube!");
	// for !window.ShouldClose() {
	// 	drawScene( program, vao)
	// }
}

func drawScene(program, vao uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)


	// Update

	model = mgl32.HomogRotate3D(float32(angle), mgl32.Vec3{0, 1, 0})

	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])


	gl.ActiveTexture(gl.TEXTURE0)
	gl.UseProgram(program)


	// Render
	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)

	// Maintenance
	window.SwapBuffers()
	glfw.PollEvents()


}
