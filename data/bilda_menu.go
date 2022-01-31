package main

import (
	"math/rand"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//MARK:var
var (
	//TXT
	txts = int32(10)
	txtm = int32(20)
	txtl = int32(40)

	//SETINGS
	settingson bool

	//MOUSE
	inmenu bool

	//MENU
	menu    = menumain{}
	infobar = menuinfobar{}
	//OBJS
	objs []obj

	//TIMERS
	fadeblink = float32(0.2)

	fadeblinkon, onoff1, onoff2, onoff3, onoff6, onoff10, onoff15, onoff30, onoff60 bool

	//CORE
	scrwf32, scrhf32              float32
	scrw, scrh                    int32
	scrhint, scrwint, frames      int
	imgs                          rl.Texture2D
	mousev2                       rl.Vector2
	camera, camintro, camintrotxt rl.Camera2D
	fps                           = int32(60)

	dev, pause, uioff, grid, cntr, sprites bool

	//IMG
	cursorimg   = rl.NewRectangle(0, 0, 12, 12)
	settingsimg = rl.NewRectangle(12, 0, 18, 18)
	closewinimg = rl.NewRectangle(30, 0, 14, 14)
	tileimg1    = rl.NewRectangle(0, 18, 32, 32)
	larrowimg   = rl.NewRectangle(44, 0, 10, 14)
	rarrowimg   = rl.NewRectangle(53, 0, 10, 14)
	uarrowimg   = rl.NewRectangle(63, 0, 14, 10)
	darrowimg   = rl.NewRectangle(77, 0, 14, 10)
)

type menumain struct {
	rec rl.Rectangle
	lr  bool
	wid float32
}
type menuinfobar struct {
	rec rl.Rectangle
	tb  bool
	wid float32
}

type obj struct {
}

func timers() { //MARK: timers
	if frames%1 == 0 {
		if onoff1 {
			onoff1 = false
		} else {
			onoff1 = true
		}
	}

	if frames%2 == 0 {
		if onoff2 {
			onoff2 = false
		} else {
			onoff2 = true
		}
	}
	if frames%3 == 0 {
		if onoff3 {
			onoff3 = false
		} else {
			onoff3 = true
		}
	}
	if frames%6 == 0 {
		if onoff6 {
			onoff6 = false
		} else {
			onoff6 = true
		}
	}
	if frames%10 == 0 {
		if onoff10 {
			onoff10 = false
		} else {
			onoff10 = true
		}
	}
	if frames%15 == 0 {
		if onoff15 {
			onoff15 = false
		} else {
			onoff15 = true
		}
	}
	if frames%30 == 0 {
		if onoff30 {
			onoff30 = false
		} else {
			onoff30 = true
		}
	}
	if frames%60 == 0 {
		if onoff60 {
			onoff60 = false
		} else {
			onoff60 = true
		}
	}
	if fadeblinkon {
		if fadeblink > 0.2 {
			fadeblink -= 0.05
		} else {
			fadeblinkon = false
		}
	} else {
		if fadeblink < 0.6 {
			fadeblink += 0.05
		} else {
			fadeblinkon = true
		}
	}

}
func nocam() { //MARK: nocam

	//centerlines
	if cntr {
		rl.DrawLine(scrw/2, 0, scrw/2, scrh, brightyellow())
		rl.DrawLine(0, scrh/2, scrw, scrh/2, brightyellow())
	}

}

func cam() { //MARK: cam

}
func nocamui() { //MARK: nocam

	rl.DrawRectangleRec(menu.rec, rl.Fade(darkred(), 0.4))
	rl.DrawRectangleRec(infobar.rec, rl.Fade(darkred(), 0.4))
	topmenuborder := int32(10)

	//closewin icon
	x := scrwf32 - (closewinimg.Width + 10)
	v2 := rl.NewVector2(x, (infobar.rec.Y+(infobar.rec.Height/2))-closewinimg.Height/2)
	rec := rl.NewRectangle(v2.X, v2.Y, closewinimg.Width, closewinimg.Height)
	if rl.CheckCollisionPointRec(mousev2, rec) {
		rl.DrawTextureRec(imgs, closewinimg, v2, brightred())
		helptxt("closewin")
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			rl.CloseWindow()
		}
	} else {
		rl.DrawTextureRec(imgs, closewinimg, v2, rl.White)
	}
	//time
	currentTime := time.Now()
	txtlen := rl.MeasureText(currentTime.Format("15:04"), txtm)
	x -= float32(txtlen + 10)
	rl.DrawText(currentTime.Format("15:04"), int32(x), int32(infobar.rec.Y)+topmenuborder, 20, rl.White)
	if infobar.tb {
		//upmenu icon
		x -= uarrowimg.Width + 10
		v2 = rl.NewVector2(x, (infobar.rec.Y+(infobar.rec.Height/2))-uarrowimg.Height/2)
		rec = rl.NewRectangle(v2.X, v2.Y, uarrowimg.Width, uarrowimg.Height)
		if rl.CheckCollisionPointRec(mousev2, rec) {
			rl.DrawTextureRec(imgs, uarrowimg, v2, brightred())
			helptxt("uarrow")
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				infobar.tb = false
				upmenu()
			}
		} else {
			rl.DrawTextureRec(imgs, uarrowimg, v2, rl.White)
		}
	} else {
		//downmenu icon
		x -= darrowimg.Width + 10
		v2 = rl.NewVector2(x, (infobar.rec.Y+(infobar.rec.Height/2))-darrowimg.Height/2)
		rec = rl.NewRectangle(v2.X, v2.Y, darrowimg.Width, darrowimg.Height)
		if rl.CheckCollisionPointRec(mousev2, rec) {
			rl.DrawTextureRec(imgs, darrowimg, v2, brightred())
			helptxt("darrow")
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				infobar.tb = true
				upmenu()
			}
		} else {
			rl.DrawTextureRec(imgs, darrowimg, v2, rl.White)
		}
	}

	//settings icon
	x -= settingsimg.Width + 10
	v2 = rl.NewVector2(x, (infobar.rec.Y+(infobar.rec.Height/2))-settingsimg.Height/2)
	rec = rl.NewRectangle(v2.X, v2.Y, settingsimg.Width, settingsimg.Height)
	if rl.CheckCollisionPointRec(mousev2, rec) {
		rl.DrawTextureRec(imgs, settingsimg, v2, brightred())
		helptxt("settings")
		settingson = true
	} else {
		rl.DrawTextureRec(imgs, settingsimg, v2, rl.White)
	}

}
func devui() {
	x := int32(scrwf32 - 140)
	y := int32(infobar.wid) + txts
	if menu.lr {
		x = 10
	}
	if infobar.tb {
		y = 10
	}
	txt := strconv.FormatBool(infobar.tb)
	rl.DrawText(txt, x, y, txts, rl.White)
	x2 := x + 70
	rl.DrawText("infobar.tb", x2, y, txts, rl.White)
	y += txts
	txt = strconv.FormatBool(menu.lr)
	rl.DrawText(txt, x, y, txts, rl.White)
	x2 = x + 70
	rl.DrawText("menu.lr", x2, y, txts, rl.White)
	x += txts

}
func helptxt(name string) {
	switch name {
	case "uarrow":
		txtlen := rl.MeasureText("menu down", txts)
		rl.DrawText("menu up", int32(mousev2.X-float32(txtlen+txts)), int32(mousev2.Y-40), txts, rl.White)
	case "darrow":
		txtlen := rl.MeasureText("menu down", txts)
		rl.DrawText("menu down", int32(mousev2.X-float32(txtlen+txts)), int32(mousev2.Y+25), txts, rl.White)
	case "closewin":
		txtlen := rl.MeasureText("exit", txts)
		if infobar.tb {
			rl.DrawText("exit", int32(mousev2.X-float32(txtlen+txts)), int32(mousev2.Y-40), txts, rl.White)
		} else {
			rl.DrawText("exit", int32(mousev2.X-float32(txtlen+txts)), int32(mousev2.Y+25), txts, rl.White)
		}
	case "settings":
		txtlen := rl.MeasureText("settings", txts)
		if infobar.tb {
			rl.DrawText("settings", int32(mousev2.X-float32(txtlen+txts)), int32(mousev2.Y-40), txts, rl.White)
		} else {
			rl.DrawText("settings", int32(mousev2.X-float32(txtlen+txts)), int32(mousev2.Y+25), txts, rl.White)
		}

	}

}
func up() { //MARK: up

	inp()
	timers()

}
func upmenu() {
	if infobar.tb {
		infobar.rec = rl.NewRectangle(0, scrhf32-infobar.wid, scrwf32, infobar.wid)
		if menu.lr {
			menu.rec = rl.NewRectangle(scrwf32-menu.wid, 0, menu.wid, scrhf32-(infobar.rec.Height+1))
		} else {
			menu.rec = rl.NewRectangle(0, 0, menu.wid, scrhf32-(infobar.rec.Height+1))
		}
	} else {
		infobar.rec = rl.NewRectangle(0, 0, scrwf32, infobar.wid)
		if menu.lr {
			menu.rec = rl.NewRectangle(scrwf32-menu.wid, infobar.rec.Height+1, menu.wid, scrhf32)
		} else {
			menu.rec = rl.NewRectangle(0, infobar.rec.Height+1, menu.wid, scrhf32)
		}
	}
}
func initial() { //MARK: initial
	dev = true
	camera.Zoom = 1.0
	makemenus()
}
func makemenus() {
	infobar.wid = 40
	menu.wid = 200
	infobar.rec = rl.NewRectangle(0, 0, scrwf32, infobar.wid)
	menu.rec = rl.NewRectangle(0, infobar.rec.Height+1, menu.wid, scrhf32)
}
func inp() { //MARK: inp

	//DEV KEY
	if rl.IsKeyPressed(rl.KeyF10) {
		if menu.lr {
			menu.lr = false
			upmenu()
		} else {
			menu.lr = true
			upmenu()
		}
	}

	if rl.IsKeyPressed(rl.KeyEscape) {
		if uioff {
			uioff = false
		} else {
			uioff = true
		}
	}

	if rl.IsKeyPressed(rl.KeyF1) {
		if dev {
			dev = false
		} else {
			dev = true
		}
	}
}

func raylib() { //MARK: raylib
	rl.SetConfigFlags(rl.FlagMsaa4xHint) // enable 4X anti-aliasing
	rl.InitWindow(scrw, scrh, "GAME TITLE")
	rl.SetExitKey(rl.KeyEnd) // key to end the game and close window

	//rl.ToggleFullscreen()
	rl.HideCursor()
	imgs = rl.LoadTexture("data/imgs.png") // load images
	//makeimgs()
	scr(1)
	initial()
	rl.SetTargetFPS(fps)

	for !rl.WindowShouldClose() {
		frames++
		mousev2 = rl.GetMousePosition()

		if rl.CheckCollisionPointRec(mousev2, menu.rec) || rl.CheckCollisionPointRec(mousev2, infobar.rec) {
			inmenu = true
		} else {
			inmenu = false
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)
		cam()

		rl.EndMode2D()

		nocam()
		//centerlines
		if cntr {
			rl.DrawLine(scrw/2, 0, scrw/2, scrh, rl.Magenta)
			rl.DrawLine(0, scrh/2, scrw, scrh/2, rl.Magenta)
		}
		if !uioff {
			nocamui()
		}
		if dev {
			devui()
		}
		//cursor
		rl.DrawTextureRec(imgs, cursorimg, mousev2, rl.White)
		rl.EndDrawing()
		up()
	}
	rl.CloseWindow()

}

func main() { //MARK: main
	rand.Seed(time.Now().UnixNano()) // random numbers
	rl.SetTraceLog(rl.LogError)      // hides info window
	scr(0)
	raylib()
}
func scr(num int) { //MARK: scr
	switch num {
	case 0:
		scrh = int32(rl.GetMonitorHeight(0))
		scrw = int32(rl.GetMonitorWidth(0))
	case 1:
		scrh = int32(rl.GetMonitorHeight(0))
		scrw = int32(rl.GetMonitorWidth(0))
		scrhf32 = float32(scrh)
		scrwf32 = float32(scrw)
		scrhint = int(scrh)
		scrwint = int(scrw)
	}

}

// MARK: colors
// https://www.rapidtables.com/web/color/RGB_Color.html
func darkred() rl.Color {
	color := rl.NewColor(55, 0, 0, 255)
	return color
}
func semidarkred() rl.Color {
	color := rl.NewColor(70, 0, 0, 255)
	return color
}
func brightred() rl.Color {
	color := rl.NewColor(230, 0, 0, 255)
	return color
}
func randomgrey() rl.Color {
	color := rl.NewColor(uint8(rInt(160, 193)), uint8(rInt(160, 193)), uint8(rInt(160, 193)), uint8(rInt(0, 255)))
	return color
}
func randombluelight() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 180)), uint8(rInt(120, 256)), uint8(rInt(120, 256)), 255)
	return color
}
func randombluedark() rl.Color {
	color := rl.NewColor(0, 0, uint8(rInt(120, 250)), 255)
	return color
}
func randomyellow() rl.Color {
	color := rl.NewColor(255, uint8(rInt(150, 256)), 0, 255)
	return color
}
func randomorange() rl.Color {
	color := rl.NewColor(uint8(rInt(250, 256)), uint8(rInt(60, 210)), 0, 255)
	return color
}
func randomred() rl.Color {
	color := rl.NewColor(uint8(rInt(128, 256)), uint8(rInt(0, 129)), uint8(rInt(0, 129)), 255)
	return color
}
func randomgreen() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 170)), uint8(rInt(100, 256)), uint8(rInt(0, 50)), 255)
	return color
}
func randomcolor() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 256)), uint8(rInt(0, 256)), uint8(rInt(0, 256)), 255)
	return color
}
func brightyellow() rl.Color {
	color := rl.NewColor(uint8(255), uint8(255), uint8(0), 255)
	return color
}
func brightbrown() rl.Color {
	color := rl.NewColor(uint8(218), uint8(165), uint8(32), 255)
	return color
}
func brightgrey() rl.Color {
	color := rl.NewColor(uint8(212), uint8(212), uint8(213), 255)
	return color
}

// random numbers
func rInt(min, max int) int {
	return rand.Intn(max-min) + min
}
func rInt32(min, max int32) int32 {
	return (rand.Int31() * (max - min)) + min

}
func rFloat32(min, max float32) float32 {
	return (rand.Float32() * (max - min)) + min
}
func flipcoin() bool {
	var b bool
	a := rInt(0, 10001)
	if a < 5000 {
		b = true
	}
	return b
}
func rolldice() int {
	a := rInt(1, 7)
	return a
}
