package shader

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"strings"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Shader struct {
	program uint32
	uniforms map[string]int32
}

func (s Shader) AddUniform(uniform string){
	unifromLocation := gl.GetUniformLocation(s.program, gl.Str(uniform+"\x00"))

	if(unifromLocation == -1){
		fmt.Println("Error could not find uniform ", uniform)
	}

	s.uniforms[uniform] = unifromLocation

}

func InitGraphic(vertexShaderSource, fragmentShaderSource string) (s Shader) {
	fmt.Println("Init GL")
	s.uniforms= make(map[string]int32)
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
	s.program = prog

	return s
}

func (s Shader) Bind(){
	gl.UseProgram(s.program)
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func (s Shader) SetUniformi(uniformName string, value int32){
	gl.Uniform1i(s.uniforms[uniformName], value)
}

func (s Shader) SetUniformf(uniformName string, value float32){
	gl.Uniform1f(s.uniforms[uniformName], value)
}

func (s Shader) SetUniform(uniformName string, value mgl32.Vec3){
	gl.Uniform3f(s.uniforms[uniformName], value.X(), value.Y(), value.Z())
}

func (s Shader) SetUniformMatrix(uniformName string, value mgl32.Mat4){
	gl.UniformMatrix4fv(s.uniforms[uniformName],1 ,false, &value[0])
}

func GetOpenglVersion()(version string){
	version = gl.GoStr(gl.GetString(gl.VERSION))
	return
}