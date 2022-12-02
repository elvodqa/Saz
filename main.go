package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	ClearColor   = rl.Color{R: 88, G: 85, B: 83, A: 255}
	menuFileOpen = false
	menuEditOpen = false
	menuViewOpen = false
	menuHelpOpen = false
)

func main() {
	rl.InitWindow(1200, 1000, "Saz")
	rl.SetExitKey(0)
	rl.SetWindowState(rl.FlagWindowResizable)
	rl.SetTargetFPS(60)
	rl.SetWindowMinSize(640, 480)
	jb_mono_normal := rl.LoadFont("resources/fonts/JetBrainsMono-Medium.ttf")
	rl.SetTraceLog(rl.LogError | rl.LogDebug | rl.LogFatal | rl.LogWarning)

	menuBarRect := rl.Rectangle{X: 0, Y: 0, Width: float32(rl.GetScreenWidth()), Height: 20}
	menuBarColor := rl.Color{R: 54, G: 54, B: 54, A: 255}

	menuFileRect := rl.Rectangle{X: 0, Y: 20, Width: 200, Height: 300}
	menuEditRect := rl.Rectangle{X: 50, Y: 20, Width: 200, Height: 300}
	menuViewRect := rl.Rectangle{X: 100, Y: 20, Width: 200, Height: 300}
	menuHelpRect := rl.Rectangle{X: 150, Y: 20, Width: 200, Height: 125}

	camera := rl.Camera3D{}
	camera.Position = rl.Vector3{0.0, 5.0, 10.0}
	camera.Target = rl.Vector3{0.0, 0.0, 0.0}
	camera.Up = rl.Vector3{0.0, 1.0, 0.0}
	camera.Fovy = 45
	camera.Projection = rl.CameraPerspective

	rl.SetTextureFilter(jb_mono_normal.Texture, rl.FilterPoint)

	for !rl.WindowShouldClose() {
		menuBarRect.Width = float32(rl.GetScreenWidth())

		rl.BeginDrawing()
		rl.ClearBackground(ClearColor)
		rl.DrawRectangleRec(menuBarRect, menuBarColor)

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(0, 0, 49, 20)) {
				SetMenuOpen("file")
			} else if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(50, 0, 49, 20)) {
				SetMenuOpen("edit")
			} else if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(100, 0, 49, 20)) {
				SetMenuOpen("view")
			} else if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.NewRectangle(150, 0, 49, 20)) {
				SetMenuOpen("help")
			} else {
				SetMenuOpen("none")
			}
		}

		if menuFileOpen {
			rl.DrawRectangleRec(menuFileRect, menuBarColor)
			rl.DrawRectangleRec(rl.NewRectangle(0, 0, 49, 20), rl.Black)
			rl.DrawTextEx(jb_mono_normal, "New", rl.Vector2{X: 5, Y: 20}, 20, 1, rl.White)
			rl.DrawTextEx(jb_mono_normal, "Open", rl.Vector2{X: 5, Y: 40}, 20, 1, rl.White)
			rl.DrawTextEx(jb_mono_normal, "Save", rl.Vector2{X: 5, Y: 60}, 20, 1, rl.White)
			rl.DrawTextEx(jb_mono_normal, "Recent", rl.Vector2{X: 5, Y: 80}, 20, 1, rl.White)
			rl.DrawTextEx(jb_mono_normal, "Close Project", rl.Vector2{X: 5, Y: 100}, 20, 1, rl.White)
			rl.DrawTextEx(jb_mono_normal, "Rename Project", rl.Vector2{X: 5, Y: 120}, 20, 1, rl.White)
			if rl.IsKeyPressed(rl.KeyRight) {
				SetMenuOpen("edit")
			}
			if rl.IsKeyReleased(rl.KeyEscape) {
				SetMenuOpen("none")
			}
		}
		if menuEditOpen {
			rl.DrawRectangleRec(menuEditRect, menuBarColor)
			rl.DrawRectangleRec(rl.NewRectangle(50, 0, 49, 20), rl.Black)
			rl.DrawTextEx(jb_mono_normal, "Settings", rl.Vector2{X: 55, Y: 20}, 20, 1, rl.White)
			if rl.IsKeyPressed(rl.KeyRight) {
				SetMenuOpen("view")
			} else if rl.IsKeyPressed(rl.KeyLeft) {
				SetMenuOpen("file")
			}
			if rl.IsKeyReleased(rl.KeyEscape) {
				SetMenuOpen("none")
			}
		}

		if menuViewOpen {
			rl.DrawRectangleRec(menuViewRect, menuBarColor)
			rl.DrawRectangleRec(rl.NewRectangle(100, 0, 49, 20), rl.Black)
			if rl.IsKeyPressed(rl.KeyRight) {
				SetMenuOpen("help")
			} else if rl.IsKeyPressed(rl.KeyLeft) {
				SetMenuOpen("edit")
			}
			if rl.IsKeyReleased(rl.KeyEscape) {
				SetMenuOpen("none")
			}
		}

		if menuHelpOpen {
			rl.DrawRectangleRec(menuHelpRect, menuBarColor)
			rl.DrawRectangleRec(rl.NewRectangle(150, 0, 49, 20), rl.Black)
			rl.DrawTextEx(jb_mono_normal, "Check GitHub", rl.Vector2{X: 155, Y: 20}, 20, 1, rl.White)
			rl.DrawTextEx(jb_mono_normal, "Tutorial", rl.Vector2{X: 155, Y: 40}, 20, 1, rl.White)
			rl.DrawTextEx(jb_mono_normal, "About", rl.Vector2{X: 155, Y: 60}, 20, 1, rl.White)
			if rl.IsKeyPressed(rl.KeyLeft) {
				SetMenuOpen("view")
			}
			if rl.IsKeyReleased(rl.KeyEscape) {
				SetMenuOpen("none")
			}
		}

		rl.DrawTextEx(jb_mono_normal, "File", rl.Vector2{X: 3, Y: 1}, 20, 1, rl.White)
		rl.DrawTextEx(jb_mono_normal, "Edit", rl.Vector2{X: 50, Y: 1}, 20, 1, rl.White)
		rl.DrawTextEx(jb_mono_normal, "View", rl.Vector2{X: 100, Y: 1}, 20, 1, rl.White)
		rl.DrawTextEx(jb_mono_normal, "Help", rl.Vector2{X: 150, Y: 1}, 20, 1, rl.White)

		rl.EndDrawing()

	}

	rl.CloseWindow()
}

func SetMenuOpen(menu string) {
	switch menu {
	case "file":
		menuFileOpen = true
		menuEditOpen = false
		menuViewOpen = false
		menuHelpOpen = false
		break
	case "edit":
		menuFileOpen = false
		menuEditOpen = true
		menuViewOpen = false
		menuHelpOpen = false
		break
	case "view":
		menuFileOpen = false
		menuEditOpen = false
		menuViewOpen = true
		menuHelpOpen = false
		break
	case "help":
		menuFileOpen = false
		menuEditOpen = false
		menuViewOpen = false
		menuHelpOpen = true
		break
	case "none":
		menuFileOpen = false
		menuEditOpen = false
		menuViewOpen = false
		menuHelpOpen = false
		break
	}
}
