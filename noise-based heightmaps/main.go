package main

import (
	"log"
	"runtime"
	"time"

	"github.com/JJ2820/3d-game-engine/engine"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var vertices = []float32{
	-1.0, 1.0, 0.0, // top-left
	-1.0, -1.0, 0.0, // bottom-left
	1.0, -1.0, 0.0, // bottom-right
	1.0, 1.0, 0.0, // top-right
}

var indices = []uint32{
	0, 1, 2, // First triangle
	0, 2, 3, // Second triangle
}

func init() {
	// GLFW and OpenGL need to run on the main thread
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalf("failed to initialize glfw: %v", err)
	}
	defer glfw.Terminate()

	window := initWindow(800, 600, "Terrain Generation")
	defer window.Destroy()

	program, err := engine.NewShader("shaders/basic.vert", "shaders/terrain.frag")
	if err != nil {
		log.Fatalf("failed to create shader: %v", err)
	}

	var vao, vbo, ebo uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	// Game Loop
	for !window.ShouldClose() {
		// Clear the screen
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Use the shader program
		gl.UseProgram(program)

		// Set uniform values (iResolution, iTime)
		resUniform := gl.GetUniformLocation(program, gl.Str("iResolution\x00"))
		gl.Uniform2f(resUniform, float32(800), float32(600))

		timeUniform := gl.GetUniformLocation(program, gl.Str("iTime\x00"))
		gl.Uniform1f(timeUniform, float32(time.Since(time.Now()).Seconds()))

		// Bind and draw the terrain quad
		gl.BindVertexArray(vao)
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.PtrOffset(0))

		// Swap buffers and poll events
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func initWindow(width, height int, title string) *glfw.Window {
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.Resizable, glfw.False)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	gl.Init()

	return window
}
