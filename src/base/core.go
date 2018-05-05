package base

import (
	"fmt"
	"runtime"
	"../fileUtil"
	"time"
	"github.com/go-gl/glfw/v3.2/glfw"
)
func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}


type CoreEngine struct {
	config fileUtil.Configuration
	Window *glfw.Window

	IsRunning bool
	time Time
	game Game
}

func (e CoreEngine) isRunning()(bool){
	return e.IsRunning
}


func CreateCoreEngine(config fileUtil.Configuration) (e CoreEngine){
	e.config = config

	InitGLFW()

	e.Window = CreateWindow(config.WINDOW_WIDTH, config.WINDOW_HEIGHT, config.NAME)
	e.Window.MakeContextCurrent()
	
	e.Window.SetMouseButtonCallback(onMouse)
	e.Window.SetKeyCallback(onKey)
	e.game = NewGame(config)


	e.IsRunning = false;
	e.time = Time{delta: 0}
	return
}


func (e CoreEngine)Start(){

	if e.isRunning() {
		return
	}

	e.run()
}


func (e CoreEngine)Stop(){
	if !e.isRunning() {
		return
	}

	e.IsRunning = false
}


func (e CoreEngine)run(){
	e.IsRunning = true
	frames:= 0
	var frameCounter int64 =0
	frameTime := 1/float64(e.config.FRAME_CAP)
	lastTime := GetTime()
	var unprocessedTime float64 =0


	for e.isRunning() == true {
		var render bool = false
		var startTime = GetTime()
		passedTime := startTime - lastTime
		lastTime = startTime

		unprocessedTime += float64(passedTime)/  float64(SECOND)
		frameCounter += passedTime
		for unprocessedTime > frameTime {
			render = true;
			unprocessedTime -= frameTime

			e.time.SetDelta(int64(frameTime));
			e.game.Input()
			e.game.Update()

			if float64(frameCounter) > float64(SECOND){
				fmt.Println(frames)
				frames = 0;
				frameCounter = 0
			}
		}

		if e.Window.ShouldClose(){
			e.Stop()
			break;
		}

		if render {
			e.render()
			frames++;
		}else{
			time.Sleep(1)
		}

	}
	e.cleanUp()
}

func (e CoreEngine)render(){

	ClearScreen()

	e.game.Render()

	window.SwapBuffers()
	glfw.PollEvents()
}


func (e CoreEngine)cleanUp(){
	defer glfw.Terminate()
}

