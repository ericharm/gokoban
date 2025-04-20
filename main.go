package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, WINDOW_TITLE)
	rl.SetTargetFPS(TARGET_FPS)

	entity := NewEntity(400, 300, 20, rl.Red)

	app := GetApplication()
	app.RootNode.AddChild(entity)

	for !rl.WindowShouldClose() {

		app.RootNode.Update()

		app.RootNode.Draw()

		ProcessInputEvents()
		ProcessRealTimeInput()

		event := DequeueInputEvent()
		for event != -1 {
			if event == LeftHeld {
				entity.Move(-5, 0)
			} else if event == RightHeld {
				entity.Move(5, 0)
			}
			event = DequeueInputEvent()
		}

	}

	rl.CloseWindow()

}
