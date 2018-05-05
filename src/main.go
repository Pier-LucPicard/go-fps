package main;

import (
	"./fileUtil"
	"./base"
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
	game := base.CreateCoreEngine(fileUtil.LoadConfig())
	game.Start()
}

