package base

import (
	"fmt"
	"log"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/gl/v4.1-core/gl"
)

func ClearScreen() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}


func InitGLFW(){
	fmt.Println("Init GLFW")
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
}

func InitGraphic(vertexShaderSource, fragmentShaderSource string) uint32 {
	fmt.Println("Init GL")
	if err := gl.Init(); err != nil {
		panic(err)
	}


	fmt.Println("OpenGL version", GetOpenglVersion())


	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
	gl.FrontFace(gl.CW)
	gl.CullFace(gl.BACK)
	gl.Enable(gl.CULL_FACE) //Back face culling
	gl.Enable(gl.DEPTH_TEST) //Order of drawing
	gl.Enable(gl.FRAMEBUFFER_SRGB)

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	return prog
	//gl.DepthFunc(gl.LESS)


}

func GetOpenglVersion()(version string){
	version = gl.GoStr(gl.GetString(gl.VERSION))
	return
}