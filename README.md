# Engine with Terrain Generation

A high-performance, real-time 3D game engine built using Golang and OpenGL/WebGL, with a focus on procedural terrain generation, optimized rendering, and multiplayer support(Working on multiplayer support).

## Features

- **Real-Time Terrain Generation**: Uses Perlin Noise for procedural terrain creation.
- **Optimized Shaders**: Efficient rendering with custom shaders for lighting, textures, and shadows.
- **Physics and Collision Detection**: Lightweight physics engine with real-time collision handling.
- **AI Opponents**: Implemented AI for non-player characters (NPCs), inspired by decision trees and heuristic evaluation.
- **Multiplayer Support**: Integrated multiplayer using Golang's concurrency features.
- **Object-Oriented Design**: Structured game objects using OOP principles for better scalability.
- **Custom Golang Interpreter**: Script handling within the engine for game logic.

## Tech Stack

- **Languages**: Golang, GLSL (OpenGL Shader Language)
- **Graphics**: OpenGL/WebGL for rendering
- **Concurrency**: Golang’s goroutines for handling multiplayer and physics.
- **GPU Acceleration**: CUDA for high-performance computing tasks(Working on integration).

## Setup

### Prerequisites

- **Golang**: Version 1.19+
- **OpenGL/WebGL**: Ensure OpenGL libraries are set up. For WebGL, use a modern browser.
- **CUDA Toolkit**: For systems with CUDA support.
- **Git**: Version control for cloning the repository.




## Usage

### Controls:
- **Move**: W, A, S, D
- **Look**: Mouse movement
- **Jump**: Space
- **Interact**: E


## Project Structure
3d-game-engine/
│
├── assets/               # Textures, models, and other game assets
├── cmd/                  # Command-line tools and server files
├── engine/               # Core game engine logic (rendering, physics, AI)
├── shaders/              # GLSL shaders for OpenGL
├── web/                  # WebGL version of the game engine
├── README.md             # Project documentation
└── go.mod                # Go modules for dependencies
