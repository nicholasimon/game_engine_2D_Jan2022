package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//MARK:var
var (
	//FX
	defaultfx                                                                                                                                  []fx
	fxdemoghost, fxdemoshadow, fxdemoscanlines, fxdemopixelnoise, fxdemoghostmenu, fxdemoshadowmenu, fxdemoscanlinesmenu, fxdemopixelnoisemenu bool
	fxdemorec                                                                                                                                  rl.Rectangle
	fxmenuon                                                                                                                                   bool
	fxdemoobjs                                                                                                                                 []fxdemoobj
	//LAYERS
	currentlayer int
	//LABELS
	labeleventnametxt        string
	deflabeleventname        = "label event"
	deflabeleventnamechanged = deflabeleventname

	labelchangeimgonobj2, labelchangeimgon, addlabelon, addinglabel, editlabelactivobjon, editlabelallon, labeleventon, labelspawnon, labelspawnon2, labelspawnon3, labelspawnon4, viewlabelsliston, addexistinglabelon bool

	editlabelnum, labeleventnum, labelsalllistnum, labelactionnum, labelactionnum2, labelactionnum3, labelactionnum2obj2, labelactionnum3obj2, labelactionnumobj2, usrsaveobjnum, usrsaveobjnum2, usrsaveobjnum3, usrsaveobjnum4 int

	labelspawnnum, labelspawnnum2, labelspawnnum3, labelspawnnum4 = 1, 1, 1, 1

	labeleventslist, labelactionslist, labelsalllist, labelactionslistobj2 []string

	//PATH
	addobjpath, createpathon bool

	//SCREEN
	addscreendirec int

	//TIMERS
	newtimerobj bool
	//EVENTS
	neweventon, addevent, addingevent bool
	eventslist                        []string
	eventsonoff                       []bool
	//COPY OBJS
	copyobjon bool
	//UIOBJS
	uiobjs        []uiobj
	activuiobjnum = blanknum
	uiobjmenuon   bool
	//TXT
	newtxtobj                   bool
	keybinput                   = make([]string, 100)
	keybcount                   int
	txts                        = int32(10)
	txtm                        = int32(20)
	txtl                        = int32(40)
	gettxton, charlimiton       bool
	gettxtcharlimit, gettxttype int
	gettxtname, gettxtreturn    string
	//SETINGS
	settingson bool
	settings   = settingsstore{}
	//INP
	ukey, dkey, lkey, rkey, ulkey, urkey, dlkey, drkey int32

	up_change_def, down_change_def, left_change_def, right_change_def, upleft_change_def, upright_change_def, downleft_change_def, downright_change_def float32
	//MOUSE
	inmenu                                                                  bool
	startv2mouse, startv2mouseworld, mousepointgridscr, mousepointgridworld rl.Vector2
	selrec, newobjrec                                                       rl.Rectangle
	clickpause                                                              int32
	//TILE SELECT MENU
	tileselecton  bool
	tilemenuonoff []menulist
	//MENUS
	rightnavrec rl.Rectangle
	togglewid   = float32(10)
	menu        = menumain{}
	infobar     = menuinfobar{}

	menufocus, newobjon, newuiobjon, colorpalon bool
	//COLORS
	colorscreated, editcoloron   bool
	changecolorsel               int
	colorpalrec, colorpalbackrec rl.Rectangle
	colorpalrand                 []color
	colorpaluser                 = make([]color, 96)
	colorpalstand                []rl.Color
	selcolor                     = color{}
	colorusercurrentnum          int
	//OBJS

	objs, backobjs, foreobjs, usrsaveobjs []obj

	activobjnum               = blanknum
	addobjcontrols, objmenuon bool
	playobjs                  bool

	//CORE
	playon                         bool
	scrwf32, scrhf32, frames32     float32
	scrw, scrh                     int32
	scrhint, scrwint, frames       int
	imgs                           rl.Texture2D
	mousev2, mousev2world, cntrscr rl.Vector2
	camera, camtileselect          rl.Camera2D
	fps                            = int32(60)

	fadeblink                                                                       = float32(0.2)
	fadeblinkon, onoff1, onoff2, onoff3, onoff6, onoff10, onoff15, onoff30, onoff60 bool

	dev, pause, uioff, cntr, sprites bool
	//BLANKS
	blankevent = event{}
	blankcolor = rl.Color{}

	blankrec  = rl.NewRectangle(0, 0, 0, 0)
	blankv2   = rl.NewVector2(0, 0)
	blanknum  = 7777777777777777777
	blankbool bool
	//IMG
	camarrowupimg    = rl.NewRectangle(128, 0, 14, 16)
	camarrowdownimg  = rl.NewRectangle(160, 0, 14, 16)
	camarrowrightimg = rl.NewRectangle(143, 1, 16, 14)
	camarrowleftimg  = rl.NewRectangle(175, 1, 16, 14)

	refreshimg  = rl.NewRectangle(111, 0, 15, 15)
	tickimg     = rl.NewRectangle(92, 0, 17, 14)
	cursorimg   = rl.NewRectangle(0, 0, 12, 12)
	settingsimg = rl.NewRectangle(12, 0, 18, 18)
	closewinimg = rl.NewRectangle(30, 0, 14, 14)
	tile1img    = rl.NewRectangle(0, 18, 32, 32)
	larrowimg   = rl.NewRectangle(44, 0, 10, 14)
	rarrowimg   = rl.NewRectangle(53, 0, 10, 14)
	uarrowimg   = rl.NewRectangle(63, 0, 14, 10)
	darrowimg   = rl.NewRectangle(77, 0, 14, 10)
	playimg     = rl.NewRectangle(193, 1, 16, 14)

	onebit16pxtiles   []rl.Rectangle
	onebitkenneytiles []rl.Rectangle
	onebitvar2tiles   []rl.Rectangle
)

//MARK: structs
type obj struct { // len to display 4
	name      string
	hp, shape int

	rotates, rotate_lr, rotate_rand bool
	rotate_speed, rotate_timer      float32

	outline_only, background_obj, middleground_obj, foreground_obj bool

	rotation, width, height, topleft_x, topleft_y float32

	path_move                bool
	direction_x, direction_y float32

	random_direction                                                                                bool
	rand_direc_timer_max, rand_direc_timer_min, rand_direc_max_x, rand_direc_max_y, rand_direc_time float32

	orig_direcx, orig_direcy float32

	complex                   bool
	tile_w, tile_h, outline_w float32

	outline_color, fill_color1, fill_color2 rl.Color

	fade float32

	gradient_v, gradient_h                bool
	img                                   rl.Rectangle
	img_rotation                          float32
	img_rotates, img_rotate_lr            bool
	img_rotate_speed                      float32
	ghosting                              bool
	ghosting_x, ghosting_y, ghosting_fade float32
	ghosting_color                        rl.Color
	shadow                                bool
	shadow_x, shadow_y, shadow_fade       float32
	shadow_color                          rl.Color

	rec, collisrec rl.Rectangle //break

	labelchangelayerobj2, labelchangelayer, labelchangehp, labelchangehpobj2 int

	labelchangeghosting, labelchangeshadow, labelchangefillgradienton, labelchangefillgradienthv, labelevent2, labelevent3, labeleventobj2, labelevent2obj2, labelchangefillgradientonobj2, labelchangefillgradienthvobj2, labelchangeghostingobj2, labelchangeshadowobj2, labelchangedirecxplus, labelchangedirecyplus, labelchangerotationplus, labelchangewidthplus, labelchangeheightplus, labelchangeposxplus, labelchangeposyplus, labelchangehpplus, labelchangehpplusobj2, labelchangeposxplusobj2, labelchangeposyplusobj2, labelchangewidthplusobj2, labelchangeheightplusobj2, labelchangerotationplusobj2, labelchangedirecxplusobj2, labelchangedirecyplusobj2 bool

	labelchangeshadowx, labelchangeshadowy, labelchangeshadowfade, labelchangeghostfade, labelchangeghostx, labelchangeghosty, labelchangedirecx, labelchangedirecy, labelchangerotation, labelchangewidth, labelchangeheight, labelchangeposx, labelchangeposy float32

	labelchangerotationobj2, labelchangedirecyobj2, labelchangedirecxobj2, labelchangewidthobj2, labelchangeheightobj2, labelchangeghostxobj2, labelchangeghostyobj2, labelchangeghostfadeobj2, labelchangeshadowxobj2, labelchangeshadowyobj2, labelchangeshadowfadeobj2, labelchangeposxobj2, labelchangeposyobj2 float32

	labelchangeimg, labelchangeimgobj2 rl.Rectangle

	labelchangeshadowcolor, labelchangefill1, labelchangefill2, labelchangeghostcolor, labelchangefill1obj2, labelchangefill2obj2, labelchangeghostcolorobj2, labelchangeshadowcolorobj2 rl.Color

	labeleventson bool
	labelevents   []labeleventdetails

	orig_rotatetimer, orig_rotatespeed, circrad float32

	tilenumw, tilenumh, currentpathv2 int
	onscreen, controlson, reversepath bool

	circv2 rl.Vector2

	objsin   []obj
	controls []controlinp
	events   []event
	path     []rl.Vector2
	labels   []string
}
type labeleventdetails struct {
	name, event, labelobj2, action, actionobj2 string
	action2, action2obj2                       string
	action3, action3obj2                       string

	spawnmore, spawnmore2, spawnmore3, spawnmore4 bool
	spawnobj, spawnobj2, spawnobj3, spawnobj4     obj
	spawnnum, spawnnum2, spawnnum3, spawnnum4     int
}
type fxdemoobj struct {
	imgon, rotates                              bool
	img                                         rl.Rectangle
	shape, x, y                                 int32
	width, radius, rotation, rotateamount, fade float32
	color                                       rl.Color
}
type fx struct {
	num                    int
	color                  rl.Color
	height, width, spacing int32
	min, max               int
	fade                   float32
	x, y                   []int32
	xf32, yf32             []float32
	x2, y2                 float32
}
type controlinp struct {
	ukey, dkey, lkey, rkey, ulkey, urkey, dlkey, drkey          int32
	up, down, left, right, upleft, upright, downleft, downright float32

	inptype int
}

type event struct {
	name string

	bounce, bounce_random, stop_moving, explode, destroy, invisible, visible bool

	score              int
	change_x, change_y float32
	event_type         int
}
type menulist struct {
	name  string
	onoff bool
}

type uiobj struct {
	name  string
	shape int

	rotation, rec_width, rec_height, polygon_circ_radius, topleft_x, topleft_y, center_x, center_y float32

	shadow                                  bool
	shadow_x, shadow_y, shadow_fade         float32
	shadow_color                            rl.Color
	text                                    string
	txt_topleft_x, txt_topleft_y            float32
	displays                                int
	complex                                 bool
	tile_w, tile_h, outline_w               float32
	outline_color, fill_color1, fill_color2 rl.Color
	gradient_v, gradient_h                  bool
	img                                     rl.Rectangle
	img_rotation                            float32
	img_rotates, img_rotate_lr              bool
	img_rotate_speed                        float32

	rec      rl.Rectangle //break
	circv2   rl.Vector2
	circrad  float32
	uiobjsin []uiobj
}
type color struct {
	r, g, b, fade uint8
	color         rl.Color
	name          string
}
type settingsstore struct {
	level_x_left, level_y_top, level_x_right, level_y_bottom, border_top, border_left, border_right, border_bottom, camera_y_change, camera_x_change, screen_width_multiplier, screen_height_multiplier, tilew, tileh, ghosting_x, ghosting_y, ghosting_fade, outline_w, img_rotate_speed, random_direc_time_max, random_direc_xy_minmax, rotate_speed, rotate_timer float32

	outline_color rl.Color
	animate       bool

	gridon, snapon, outlineson, ruleron, colorpallockon, showpathson bool //break

	recgrid []rl.Rectangle
}
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

//MARK: CAMERAS CAMERAS CAMERAS CAMERAS CAMERAS CAMERAS CAMERAS CAMERAS CAMERAS CAMERAS CAMERAS
func cam() { //MARK: cam

	//drawobjs
	if len(objs) > 0 {
		for a := 0; a < len(objs); a++ {
			if objs[a].background_obj {
				dobjs()
			}
		}
		for a := 0; a < len(objs); a++ {
			if objs[a].middleground_obj {
				dobjs()
			}
		}
		for a := 0; a < len(objs); a++ {
			if objs[a].foreground_obj {
				dobjs()
			}
		}
	}

	//grid
	if settings.gridon {
		x := int32(camera.Target.X)
		y := int32(camera.Target.Y)

		for {
			rl.DrawLine(x, y, x, y+scrh, rl.Fade(brightorange(), 0.2))
			x += int32(settings.tilew)
			if x > int32(camera.Target.X+scrwf32) {
				break
			}
		}
		x = int32(camera.Target.X)
		for {
			rl.DrawLine(x, y, x+scrw, y, rl.Fade(brightorange(), 0.2))
			y += int32(settings.tileh)
			if y > int32(camera.Target.Y+scrhf32) {
				break
			}
		}

	}
}
func nocam() { //MARK: nocam

	//drawuiobjs
	duiobjs()

	//select rec
	if selrec != blankrec {
		rl.DrawRectangleRec(selrec, rl.Fade(brightorange(), 0.1))
		rl.DrawRectangleLinesEx(selrec, 1.0, brightorange())
	}

	//centerlines
	if cntr {
		rl.DrawLine(scrw/2, 0, scrw/2, scrh, brightyellow())
		rl.DrawLine(0, scrh/2, scrw, scrh/2, brightyellow())
	}

	//ruler
	if settings.ruleron {

		x := int32(camera.Target.X)
		y := int32(scrhf32)
		if infobar.tb {
			y = 0
		}
		txtonoff := false

		for {
			if infobar.tb {
				rl.DrawLine(x, y, x, y+8, rl.White)
			} else {
				rl.DrawLine(x, y, x, y-8, rl.White)
			}
			x += int32(settings.tilew)
			if txtonoff {
				txtonoff = false
			} else {
				txtonoff = true
			}
			if txtonoff {
				txt := strconv.Itoa(int(x))
				txtlen := rl.MeasureText(txt, txts)
				if infobar.tb {
					rl.DrawText(txt, x-(txtlen/2), y+20, txts, rl.White)
				} else {
					rl.DrawText(txt, x-(txtlen/2), y-20, txts, rl.White)
				}
			}
			if x >= int32(camera.Target.X)+scrw {
				break
			}
		}

	}

}
func nocamui() { //MARK: nocamui

	if settingson {
		rl.DrawRectangle(0, 0, scrw, scrh, rl.Black)
		rl.DrawRectangle(0, 0, scrw, scrh, rl.Fade(darkred(), 0.4))
		dsettings()
	} else if tileselecton {
		rl.DrawRectangle(0, 0, scrw, scrh, rl.Black)
		rl.DrawRectangle(0, 0, scrw, scrh, rl.Fade(darkred(), 0.4))
		dtileselect()
	} else if fxmenuon {
		rl.DrawRectangle(0, 0, scrw, scrh, rl.Black)
		rl.DrawRectangle(0, 0, scrw, scrh, rl.Fade(darkred(), 0.4))
		dfxmenu()
	} else {
		rl.DrawRectangleRec(menu.rec, rl.Fade(darkred(), 0.4))
		rl.DrawRectangleRec(infobar.rec, rl.Fade(darkred(), 0.4))
		//infobar
		infobarmenu()
		inforicons()
		//cam icons
		dcamicons()
		//main menus
		mainmenu()
	}

	//labels
	if addexistinglabelon {
		listgeneric("addexistinglabel")
	}
	if viewlabelsliston {
		listgeneric("labelslist")
	}
	if editlabelallon {
		menugeneric("editlabelall")
	}
	if editlabelactivobjon {
		menugeneric("editlabelactivobj")
	}
	if labeleventon {
		menugeneric("labelevent")
	}
	if labelchangeimgon || labelchangeimgonobj2 {
		rl.DrawRectangle(0, 0, scrw, scrh, rl.Black)
		rl.DrawRectangle(0, 0, scrw, scrh, rl.Fade(darkred(), 0.4))
		dtileselect()
	}

	//colorpal
	if colorpalon {
		dcolorpal()
	}
	//gettxt
	if gettxton {
		gettxt()
	}
}
func devui() { //MARK: devui

	x := int32(scrwf32 - 140)
	y := int32(infobar.wid) + txts

	rl.DrawRectangle(x-20, y, 160, scrh-(int32(infobar.wid)+txts), rl.Fade(rl.Black, 0.9))

	if menu.lr {
		x = 10
	}
	if infobar.tb {
		y = 10
	}
	txt := strconv.FormatBool(infobar.tb)
	rl.DrawText(txt, x, y, txts, brightred())
	x2 := x + 70
	rl.DrawText("infobar.tb", x2, y, txts, brightred())
	y += txts
	txt = strconv.FormatBool(menu.lr)
	rl.DrawText(txt, x, y, txts, brightred())
	x2 = x + 70
	rl.DrawText("menu.lr", x2, y, txts, brightred())
	y += txts
	txt = strconv.FormatBool(newobjon)
	rl.DrawText(txt, x, y, txts, brightred())
	x2 = x + 70
	rl.DrawText("newobjon", x2, y, txts, brightred())
	y += txts
	txt = fmt.Sprint(camtileselect.Target.Y)
	rl.DrawText(txt, x, y, txts, brightred())
	x2 = x + 70
	rl.DrawText("camtile Y", x2, y, txts, brightred())
	y += txts
	txt = fmt.Sprint(mousepointgridscr.Y)
	rl.DrawText(txt, x, y, txts, brightred())
	x2 = x + 70
	rl.DrawText("mousepointgridscr Y", x2, y, txts, brightred())
	y += txts
	txt = fmt.Sprint(mousepointgridworld.Y)
	rl.DrawText(txt, x, y, txts, brightred())
	x2 = x + 70
	rl.DrawText("mousepointgridworld Y", x2, y, txts, brightred())
	y += txts
	txt = fmt.Sprint(startv2mouse.Y)
	rl.DrawText(txt, x, y, txts, brightred())
	x2 = x + 70
	rl.DrawText("startv2mouse Y", x2, y, txts, brightred())
	y += txts
	txt = strconv.Itoa(colorusercurrentnum)
	rl.DrawText(txt, x, y, txts, brightred())
	x2 = x + 70
	rl.DrawText("colorusercurrentnum", x2, y, txts, brightred())
	y += txts
	txt = strconv.FormatBool(labelchangeimgon)
	rl.DrawText(txt, x, y, txts, brightred())
	x2 = x + 70
	rl.DrawText("labelchangeimgon", x2, y, txts, brightred())
	y += txts
	txt = fmt.Sprint(clickpause)
	rl.DrawText(txt, x, y, txts, brightred())
	x2 = x + 70
	rl.DrawText("clickpause", x2, y, txts, brightred())
	y += txts
}
func addscreen() { //MARK: addscreen

	switch addscreendirec {
	case 3:
		if camera.Target.Y+scrhf32 == settings.level_y_bottom {
			gettxton = true
			gettxtcharlimit = 2
			gettxtname = "addscreen_height"
			gettxttype = 2
		} else {
			camera.Target.Y += settings.camera_y_change
		}
	case 2:
		if camera.Target.X+scrwf32 == settings.level_x_right {
			gettxton = true
			gettxtcharlimit = 2
			gettxtname = "addscreen_width"
			gettxttype = 2
		} else {
			camera.Target.X += settings.camera_x_change
		}
	case 4:
		if camera.Target.X == settings.level_x_left {
			gettxton = true
			gettxtcharlimit = 2
			gettxtname = "addscreen_width"
			gettxttype = 2
		} else {
			camera.Target.X -= settings.camera_x_change
		}
	case 1:
		if camera.Target.Y == settings.level_y_top {
			gettxton = true
			gettxtcharlimit = 2
			gettxtname = "addscreen_height"
			gettxttype = 2
		} else {
			camera.Target.Y -= settings.camera_y_change

		}

	}

}

//MARK: DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW
func dcamicons() { //MARK: dcamicons

	v2 := rl.NewVector2(scrwf32-60, 60)
	if colorpalon {
		v2.X -= menu.rec.Width
	}

	if menu.lr {
		v2.X = 40
		if colorpalon {
			v2.X += menu.rec.Width
		}
	}

	uprec := rl.NewRectangle(v2.X-2, v2.Y-2, camarrowupimg.Width+4, camarrowupimg.Height+4)
	if rl.CheckCollisionPointRec(mousev2, uprec) {
		rl.DrawTextureRec(imgs, camarrowupimg, v2, rl.Fade(brightorange(), fadeblink))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			addscreendirec = 1
			addscreen()
		}
	} else {
		rl.DrawTextureRec(imgs, camarrowupimg, v2, rl.White)
	}
	v2.Y += camarrowupimg.Height
	v2.X -= camarrowleftimg.Width

	leftrec := rl.NewRectangle(v2.X-1, v2.Y-2, camarrowleftimg.Width+4, camarrowleftimg.Height+4)
	if rl.CheckCollisionPointRec(mousev2, leftrec) {
		rl.DrawTextureRec(imgs, camarrowleftimg, v2, rl.Fade(brightorange(), fadeblink))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			addscreendirec = 4
			addscreen()
		}
	} else {
		rl.DrawTextureRec(imgs, camarrowleftimg, v2, rl.White)
	}
	v2.X += camarrowleftimg.Width * 2

	rightrec := rl.NewRectangle(v2.X-1, v2.Y-2, camarrowrightimg.Width+4, camarrowrightimg.Height+4)

	if rl.CheckCollisionPointRec(mousev2, rightrec) {
		rl.DrawTextureRec(imgs, camarrowrightimg, v2, rl.Fade(brightorange(), fadeblink))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			addscreendirec = 2
			addscreen()
		}
	} else {
		rl.DrawTextureRec(imgs, camarrowrightimg, v2, rl.White)
	}
	v2.X -= camarrowleftimg.Width
	v2.Y += camarrowupimg.Height

	downrec := rl.NewRectangle(v2.X-2, v2.Y-2, camarrowdownimg.Width+4, camarrowdownimg.Height+4)

	if rl.CheckCollisionPointRec(mousev2, downrec) {
		rl.DrawTextureRec(imgs, camarrowdownimg, v2, rl.Fade(brightorange(), fadeblink))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			addscreendirec = 3
			addscreen()
		}
	} else {
		rl.DrawTextureRec(imgs, camarrowdownimg, v2, rl.White)
	}

}
func dmenurightnav(endy float32) { //MARK: dmenurightnav

	//	mouseworld := rl.GetWorldToScreen2D(mousev2, camtileselect)
	mousescreen := rl.GetScreenToWorld2D(mousev2, camtileselect)

	rl.DrawRectangleRec(rightnavrec, rl.Black)

	navrecinner := rl.NewRectangle(rightnavrec.X+3, rightnavrec.Y+3, rightnavrec.Width-6, rightnavrec.Height-6)
	//down up arrows
	v2 := rl.NewVector2(navrecinner.X+3, navrecinner.Y+5)
	rl.DrawTextureRec(imgs, uarrowimg, v2, rl.White)

	v2 = rl.NewVector2(navrecinner.X+3, navrecinner.Y+navrecinner.Height-15)
	rl.DrawTextureRec(imgs, darrowimg, v2, rl.White)

	rl.DrawRectangleLinesEx(navrecinner, 1.0, rl.White)

	if rl.CheckCollisionPointRec(mousescreen, navrecinner) {
		rec2 := rl.NewRectangle(navrecinner.X, mousescreen.Y, 20, 20)
		if rec2.Y+rec2.Height > (navrecinner.Y+navrecinner.Height)-rec2.Height {
			rec2.Y = (navrecinner.Y + navrecinner.Height) - rec2.Height
		}
		rl.DrawRectangleLinesEx(rec2, 1.0, rl.White)

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if mousev2.Y > scrhf32/2 {
				camtileselect.Target.Y += 20
				navrecinner.Y += 20
				rightnavrec.Y += 20
			} else {
				if camtileselect.Target.Y > 0 {
					camtileselect.Target.Y -= 20
					navrecinner.Y -= 20
					rightnavrec.Y -= 20
				}
			}
		} else if rl.IsMouseButtonDown(rl.MouseLeftButton) {

			if mousev2.Y > scrhf32/2 {

				if camtileselect.Target.Y+scrhf32 < endy {
					camtileselect.Target.Y += 20
					navrecinner.Y += 20
					rightnavrec.Y += 20
				}
			} else {
				if camtileselect.Target.Y > 0 {
					camtileselect.Target.Y -= 20
					navrecinner.Y -= 20
					rightnavrec.Y -= 20
				}
			}

		}
	}

}
func dtileselect() { //MARK: dtileselect

	//draw tiles camera
	rl.BeginMode2D(camtileselect)

	x := float32(10)
	y := float32(60)

	for _, tilemenuitem := range tilemenuonoff {

		if tilemenuitem.onoff {

			switch tilemenuitem.name {
			case "1 bit various":

				mousescreen := rl.GetScreenToWorld2D(mousev2, camtileselect)

				for a := 0; a < len(onebitkenneytiles); a++ {

					destrec := rl.NewRectangle(x, y, 32, 32)
					origin := rl.NewVector2(0, 0)

					rl.DrawTexturePro(imgs, onebitkenneytiles[a], destrec, origin, 0.0, rl.White)

					if rl.CheckCollisionPointRec(mousescreen, destrec) {
						rl.DrawRectangleRec(destrec, rl.Fade(brightorange(), fadeblink))
						if clickpause == 0 {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if labelchangeimgon {
									objs[activobjnum].labelchangeimg = onebitkenneytiles[a]
									labelchangeimgon = false
								} else if labelchangeimgonobj2 {
									objs[activobjnum].labelchangeimgobj2 = onebitkenneytiles[a]
									labelchangeimgonobj2 = false
								} else {
									objs[activobjnum].img = onebitkenneytiles[a]
									if objs[activobjnum].complex {
										for a := 0; a < len(objs[activobjnum].objsin); a++ {
											objs[activobjnum].objsin[a].img = objs[activobjnum].img
										}
									}
								}
								tilemenuoff()
							}
						}
					}
					x += 40
					if x+40 >= scrwf32 {
						x = 10
						y += 40
					}
				}
				for a := 0; a < len(onebitvar2tiles); a++ {

					destrec := rl.NewRectangle(x, y, 32, 32)
					origin := rl.NewVector2(0, 0)

					rl.DrawTexturePro(imgs, onebitvar2tiles[a], destrec, origin, 0.0, rl.White)
					if rl.CheckCollisionPointRec(mousescreen, destrec) {
						rl.DrawRectangleRec(destrec, rl.Fade(brightorange(), fadeblink))
						if clickpause == 0 {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if labelchangeimgon {
									objs[activobjnum].labelchangeimg = onebitvar2tiles[a]
									labelchangeimgon = false
								} else if labelchangeimgonobj2 {
									objs[activobjnum].labelchangeimgobj2 = onebitvar2tiles[a]
									labelchangeimgonobj2 = false
								} else {
									objs[activobjnum].img = onebitvar2tiles[a]
									if objs[activobjnum].complex {
										for a := 0; a < len(objs[activobjnum].objsin); a++ {
											objs[activobjnum].objsin[a].img = objs[activobjnum].img
										}
									}
								}
								tilemenuoff()
							}
						}
					}
					x += 40
					if x+40 >= scrwf32 {
						x = 10
						y += 40
					}
				}

				dmenurightnav(y)

			case "1 bit tiles":

				for a := 0; a < len(onebit16pxtiles); a++ {

					destrec := rl.NewRectangle(x, y, 32, 32)
					origin := rl.NewVector2(0, 0)

					rl.DrawTexturePro(imgs, onebit16pxtiles[a], destrec, origin, 0.0, rl.White)
					if rl.CheckCollisionPointRec(mousev2, destrec) {
						rl.DrawRectangleRec(destrec, rl.Fade(brightorange(), fadeblink))
						if clickpause == 0 {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if labelchangeimgon {
									objs[activobjnum].labelchangeimg = onebit16pxtiles[a]
									labelchangeimgon = false
								} else if labelchangeimgonobj2 {
									objs[activobjnum].labelchangeimgobj2 = onebit16pxtiles[a]
									labelchangeimgonobj2 = false
								} else {
									objs[activobjnum].img = onebit16pxtiles[a]
									if objs[activobjnum].complex {
										for a := 0; a < len(objs[activobjnum].objsin); a++ {
											objs[activobjnum].objsin[a].img = objs[activobjnum].img
										}
									}
								}
								tilemenuoff()
							}
						}
					}

					x += 40

					if x+40 >= scrwf32 {
						x = 10
						y += 40
					}
				}

			}

		}

	}

	rl.EndMode2D()

	//buttonmainmenus
	rl.DrawRectangle(0, 0, scrw, 40, darkred())

	x = float32(10)
	y = float32(10)

	rec3 := blankrec

	for a := 0; a < len(tilemenuonoff); a++ {

		x, y, rec3 = buttonmainmenu(tilemenuonoff[a].name, x, y, tilemenuonoff[a].onoff)
		if rl.CheckCollisionPointRec(mousev2, rec3) {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				if !tilemenuonoff[a].onoff {
					tilemenuonoff[a].onoff = true
					switchtilemenu(a)
				}
			}
		}
	}

	//closewin
	v2 := rl.NewVector2(scrwf32-closewinimg.Width*2, closewinimg.Height)
	rec2 := rl.NewRectangle(v2.X, v2.Y, closewinimg.Width, closewinimg.Height)
	if rl.CheckCollisionPointRec(mousev2, rec2) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			tileselecton = false
			menufocus = false
		}
		rl.DrawTextureRec(imgs, closewinimg, v2, brightred())
	} else {
		rl.DrawTextureRec(imgs, closewinimg, v2, rl.White)
	}

}
func dobjs() { //MARK: dobjs

	for a := 0; a < len(objs); a++ {

		//movement
		if playon {
			if objs[a].path_move {
				if len(objs[a].path) > 1 {
					moveobjpath(a)
				}
			} else {
				if objs[a].random_direction {
					upobjrandomdirection(a)
				}
				if objs[a].direction_x != 0 || objs[a].direction_y != 0 {
					moveobj(a)
				}
			}
		}
		if objs[a].onscreen {
			//objsin
			if objs[a].complex {

				for b := 0; b < len(objs[a].objsin); b++ {

					//no color
					if objs[a].objsin[b].outline_color == blankcolor && objs[a].objsin[b].fill_color1 == blankcolor {
						rl.DrawRectangleRec(objs[a].objsin[b].rec, rl.Fade(rl.Green, 0.1))
						rl.DrawText("no color", objs[a].objsin[b].rec.ToInt32().X+txts, objs[a].objsin[b].rec.ToInt32().Y+txts, txts, rl.White)
					}
					//img
					if objs[a].objsin[b].img != blankrec {
						destrec := rl.NewRectangle(objs[a].objsin[b].rec.X+objs[a].objsin[b].rec.Width/2, objs[a].objsin[b].rec.Y+objs[a].objsin[b].rec.Height/2, objs[a].objsin[b].rec.Width, objs[a].objsin[b].rec.Height)
						origin := rl.NewVector2(objs[a].objsin[b].rec.Width/2, objs[a].objsin[b].rec.Height/2)

						//shadow
						if objs[a].objsin[b].shadow {
							shadowrec := rl.NewRectangle(destrec.X+objs[a].objsin[b].shadow_x, destrec.Y+objs[a].objsin[b].shadow_y, destrec.Width, destrec.Height)
							rl.DrawTexturePro(imgs, objs[a].objsin[b].img, shadowrec, origin, objs[a].objsin[b].img_rotation, rl.Fade(objs[a].objsin[b].shadow_color, objs[a].objsin[b].shadow_fade))
						}

						if objs[a].objsin[b].fill_color1 == blankcolor {
							rl.DrawTexturePro(imgs, objs[a].objsin[b].img, destrec, origin, objs[a].objsin[b].img_rotation, rl.Fade(rl.White, objs[a].objsin[b].fade))

						} else {
							rl.DrawTexturePro(imgs, objs[a].objsin[b].img, destrec, origin, objs[a].objsin[b].img_rotation, rl.Fade(objs[a].objsin[b].fill_color1, objs[a].objsin[b].fade))
						}

						//shadow
						if objs[a].objsin[b].shadow {
							shadowrec := rl.NewRectangle(destrec.X+objs[a].objsin[b].shadow_x, destrec.Y+objs[a].objsin[b].shadow_y, destrec.Width, destrec.Height)
							rl.DrawTexturePro(imgs, objs[a].objsin[b].img, shadowrec, origin, objs[a].objsin[b].img_rotation, rl.Fade(objs[a].objsin[b].shadow_color, objs[a].objsin[b].shadow_fade))
						}

						//ghosting img
						if objs[a].objsin[b].ghosting {

							destrec.X += rFloat32(-objs[a].objsin[b].ghosting_x, objs[a].objsin[b].ghosting_x+1)
							destrec.Y += rFloat32(-objs[a].objsin[b].ghosting_y, objs[a].objsin[b].ghosting_y+1)

							rl.DrawTexturePro(imgs, objs[a].objsin[b].img, destrec, origin, objs[a].objsin[b].img_rotation, rl.Fade(objs[a].objsin[b].ghosting_color, objs[a].objsin[b].ghosting_fade))

						}

					} else {

						//shadow
						if objs[a].objsin[b].shadow {
							shadowrec := rl.NewRectangle(objs[a].objsin[b].rec.X+objs[a].objsin[b].shadow_x, objs[a].objsin[b].rec.Y+objs[a].objsin[b].shadow_y, objs[a].objsin[b].rec.Width, objs[a].objsin[b].rec.Height)
							rl.DrawRectangleRec(shadowrec, rl.Fade(objs[a].objsin[b].shadow_color, objs[a].objsin[b].shadow_fade))
						}

						if objs[a].objsin[b].gradient_h {
							rl.DrawRectangleGradientH(objs[a].objsin[b].rec.ToInt32().X, objs[a].objsin[b].rec.ToInt32().Y, objs[a].objsin[b].rec.ToInt32().Width, objs[a].objsin[b].rec.ToInt32().Height, objs[a].objsin[b].fill_color1, objs[a].objsin[b].fill_color2)
						} else if objs[a].gradient_v {
							rl.DrawRectangleGradientV(objs[a].objsin[b].rec.ToInt32().X, objs[a].objsin[b].rec.ToInt32().Y, objs[a].objsin[b].rec.ToInt32().Width, objs[a].objsin[b].rec.ToInt32().Height, objs[a].objsin[b].fill_color1, objs[a].objsin[b].fill_color2)
						} else {
							rl.DrawRectangleRec(objs[a].objsin[b].rec, rl.Fade(objs[a].objsin[b].fill_color1, objs[a].objsin[b].fade))

							//ghosting solid fill
							if objs[a].objsin[b].ghosting {
								ghostx := rFloat32(-objs[a].objsin[b].ghosting_x, objs[a].objsin[b].ghosting_x+1)
								ghosty := rFloat32(-objs[a].objsin[b].ghosting_y, objs[a].objsin[b].ghosting_y+1)

								ghostrec := rl.NewRectangle(objs[a].objsin[b].rec.X+ghostx, objs[a].objsin[b].rec.Y+ghosty, objs[a].objsin[b].rec.Width, objs[a].objsin[b].rec.Height)

								rl.DrawRectangleRec(ghostrec, rl.Fade(objs[a].objsin[b].ghosting_color, objs[a].objsin[b].ghosting_fade))
							}
						}

					}

					//objsin outline
					if settings.outlineson {
						rl.DrawRectangleLinesEx(objs[a].objsin[b].rec, objs[a].objsin[b].outline_w, objs[a].objsin[b].outline_color)
					}
				}

			} else { //obj

				if objs[a].shape == 0 { // square
					//no color
					if objs[a].outline_color == blankcolor && objs[a].fill_color1 == blankcolor {
						rl.DrawRectangleRec(objs[a].rec, rl.Fade(rl.Green, 0.1))
						rl.DrawText("no color", objs[a].rec.ToInt32().X+txts, objs[a].rec.ToInt32().Y+txts, txts, rl.White)
					}
					//img
					if objs[a].img != blankrec {

						destrec := rl.NewRectangle(objs[a].rec.X+objs[a].rec.Width/2, objs[a].rec.Y+objs[a].rec.Height/2, objs[a].rec.Width, objs[a].rec.Height)
						origin := rl.NewVector2(objs[a].rec.Width/2, objs[a].rec.Height/2)

						//shadow
						if objs[a].shadow {
							shadowrec := rl.NewRectangle(destrec.X+objs[a].shadow_x, destrec.Y+objs[a].shadow_y, destrec.Width, destrec.Height)

							rl.DrawTexturePro(imgs, objs[a].img, shadowrec, origin, objs[a].img_rotation, rl.Fade(objs[a].shadow_color, objs[a].shadow_fade))
						}

						//img
						if objs[a].fill_color1 == blankcolor {
							rl.DrawTexturePro(imgs, objs[a].img, destrec, origin, objs[a].img_rotation, rl.Fade(rl.White, objs[a].fade))
						} else {
							rl.DrawTexturePro(imgs, objs[a].img, destrec, origin, objs[a].img_rotation, rl.Fade(objs[a].fill_color1, objs[a].fade))
						}

						//ghosting img
						if objs[a].ghosting {
							destrec.X += rFloat32(-objs[a].ghosting_x, objs[a].ghosting_x+1)
							destrec.Y += rFloat32(-objs[a].ghosting_y, objs[a].ghosting_y+1)

							if objs[a].fill_color1 == blankcolor {
								rl.DrawTexturePro(imgs, objs[a].img, destrec, origin, objs[a].img_rotation, rl.Fade(objs[a].ghosting_color, objs[a].ghosting_fade))
							} else {
								rl.DrawTexturePro(imgs, objs[a].img, destrec, origin, objs[a].img_rotation, rl.Fade(objs[a].ghosting_color, objs[a].ghosting_fade))
							}
						}

					} else {

						//shadow
						if objs[a].shadow {
							shadowrec := rl.NewRectangle(objs[a].rec.X+objs[a].shadow_x, objs[a].rec.Y+objs[a].shadow_y, objs[a].rec.Width, objs[a].rec.Height)
							rl.DrawRectangleRec(shadowrec, rl.Fade(objs[a].shadow_color, objs[a].shadow_fade))
						}
						//fill
						if objs[a].outline_only {
							rl.DrawRectangleLinesEx(objs[a].rec, objs[a].outline_w, objs[a].outline_color)
						} else {
							if objs[a].gradient_h {
								rl.DrawRectangleGradientH(objs[a].rec.ToInt32().X, objs[a].rec.ToInt32().Y, objs[a].rec.ToInt32().Width, objs[a].rec.ToInt32().Height, objs[a].fill_color1, objs[a].fill_color2)
							} else if objs[a].gradient_v {
								rl.DrawRectangleGradientV(objs[a].rec.ToInt32().X, objs[a].rec.ToInt32().Y, objs[a].rec.ToInt32().Width, objs[a].rec.ToInt32().Height, objs[a].fill_color1, objs[a].fill_color2)
							} else {
								rl.DrawRectangleRec(objs[a].rec, rl.Fade(objs[a].fill_color1, objs[a].fade))
							}
						}
					}

					//obj outline
					if settings.outlineson {
						rl.DrawRectangleLinesEx(objs[a].rec, objs[a].outline_w, objs[a].outline_color)
					}
				} else {

					if objs[a].shape == 1 {
						//shadow
						if objs[a].shadow {
							v2 := rl.NewVector2(objs[a].circv2.X+objs[a].shadow_x, objs[a].circv2.Y+objs[a].shadow_y)
							rl.DrawCircleV(v2, objs[a].circrad, rl.Fade(objs[a].shadow_color, objs[a].shadow_fade))
						}
						//shape
						if objs[a].outline_only {
							rl.DrawCircleLines(int32(objs[a].circv2.X), int32(objs[a].circv2.Y), objs[a].circrad, objs[a].outline_color)
						} else {
							if objs[a].fill_color1 == blankcolor {
								rl.DrawCircleV(objs[a].circv2, objs[a].circrad, brightred())
							} else {
								rl.DrawCircleV(objs[a].circv2, objs[a].circrad, objs[a].fill_color1)
							}
							rl.DrawCircleLines(int32(objs[a].circv2.X), int32(objs[a].circv2.Y), objs[a].circrad, objs[a].outline_color)
						}
					} else {
						//shadow
						if objs[a].shadow {
							v2 := rl.NewVector2(objs[a].circv2.X+objs[a].shadow_x, objs[a].circv2.Y+objs[a].shadow_y)
							rl.DrawPoly(v2, int32(objs[a].shape+1), objs[a].circrad, objs[a].rotation, rl.Fade(objs[a].shadow_color, objs[a].shadow_fade))
						}
						//shape
						if objs[a].outline_only {
							rl.DrawPolyLines(objs[a].circv2, int32(objs[a].shape+1), objs[a].circrad, objs[a].rotation, objs[a].outline_color)
						} else {
							if objs[a].fill_color1 == blankcolor {
								rl.DrawPoly(objs[a].circv2, int32(objs[a].shape+1), objs[a].circrad, objs[a].rotation, brightred())
							} else {
								rl.DrawPoly(objs[a].circv2, int32(objs[a].shape+1), objs[a].circrad, objs[a].rotation, objs[a].fill_color1)
							}
							rl.DrawPolyLines(objs[a].circv2, int32(objs[a].shape+1), objs[a].circrad, objs[a].rotation, objs[a].outline_color)
						}

					}

				}
			}

			//objs movement path
			if settings.showpathson {
				if len(objs[a].path) > 0 {

					for b := 0; b < len(objs[a].path); b++ {
						rl.DrawCircleV(objs[a].path[b], 4, rl.SkyBlue)
						if b > 0 {
							rl.DrawLineV(objs[a].path[b-1], objs[a].path[b], rl.SkyBlue)
						}

					}
				}
			}
			//select rec
			if a == activobjnum {
				if objs[a].shape == 0 {
					rec2 := rl.NewRectangle(objs[a].rec.X-4, objs[a].rec.Y-4, objs[a].rec.Width+8, objs[a].rec.Height+8)
					if onoff15 {
						rec2.X -= 4
						rec2.Y -= 4
						rec2.Width += 8
						rec2.Height += 8
					}
					rl.DrawRectangleLinesEx(rec2, 1.0, brightred())
				} else {
					circrad := objs[a].circrad + 8
					if onoff15 {
						circrad += 8
					}

					rl.DrawCircleLines(int32(objs[a].circv2.X), int32(objs[a].circv2.Y), circrad, brightred())

				}
			}
		}
	}

	//draw create path
	if createpathon {
		if len(objs[activobjnum].path) > 0 {

			for a := 0; a < len(objs[activobjnum].path); a++ {
				rl.DrawCircleV(objs[activobjnum].path[a], 4, rl.SkyBlue)
				if a > 0 {
					rl.DrawLineV(objs[activobjnum].path[a-1], objs[activobjnum].path[a], rl.SkyBlue)
				}
			}
		}
	}

}

func duiobjs() { //MARK: duiobjs

	for a := 0; a < len(uiobjs); a++ {

		if settings.outlineson {
			if uiobjs[a].shape == 0 {
				rl.DrawRectangleLinesEx(uiobjs[a].rec, 1.0, brightorange())
			} else {
				rl.DrawCircleLines(int32(uiobjs[a].circv2.X), int32(uiobjs[a].circv2.Y), uiobjs[a].circrad, brightorange())
			}
		}

		switch uiobjs[a].shape {
		case 0:
			//shadow
			if uiobjs[a].shadow {
				rec := rl.NewRectangle(uiobjs[a].rec.X+uiobjs[a].shadow_x, uiobjs[a].rec.Y+uiobjs[a].shadow_y, uiobjs[a].rec.Width, uiobjs[a].rec.Height)
				rl.DrawRectangleRec(rec, rl.Fade(uiobjs[a].shadow_color, uiobjs[a].shadow_fade))
			}
			//shape
			if uiobjs[a].fill_color1 == blankcolor {
				rl.DrawRectangleRec(uiobjs[a].rec, brightred())
			} else {
				rl.DrawRectangleRec(uiobjs[a].rec, uiobjs[a].fill_color1)
			}
		case 1:
			//shadow
			if uiobjs[a].shadow {
				v2 := rl.NewVector2(uiobjs[a].circv2.X+uiobjs[a].shadow_x, uiobjs[a].circv2.Y+uiobjs[a].shadow_y)
				rl.DrawCircleV(v2, uiobjs[a].circrad, rl.Fade(uiobjs[a].shadow_color, uiobjs[a].shadow_fade))
			}
			//shape
			if uiobjs[a].fill_color1 == blankcolor {
				rl.DrawCircleV(uiobjs[a].circv2, uiobjs[a].circrad, brightred())
			} else {
				rl.DrawCircleV(uiobjs[a].circv2, uiobjs[a].circrad, uiobjs[a].fill_color1)
			}

		case 2:
			//shadow
			if uiobjs[a].shadow {
				v2 := rl.NewVector2(uiobjs[a].circv2.X+uiobjs[a].shadow_x, uiobjs[a].circv2.Y+uiobjs[a].shadow_y)
				rl.DrawPoly(v2, 3, uiobjs[a].circrad, uiobjs[a].rotation, rl.Fade(uiobjs[a].shadow_color, uiobjs[a].shadow_fade))
			}
			//shape
			if uiobjs[a].fill_color1 == blankcolor {
				rl.DrawPoly(uiobjs[a].circv2, 3, uiobjs[a].circrad, uiobjs[a].rotation, brightred())
			} else {
				rl.DrawPoly(uiobjs[a].circv2, 3, uiobjs[a].circrad, uiobjs[a].rotation, uiobjs[a].fill_color1)
			}
		case 3:
			//shadow
			if uiobjs[a].shadow {
				v2 := rl.NewVector2(uiobjs[a].circv2.X+uiobjs[a].shadow_x, uiobjs[a].circv2.Y+uiobjs[a].shadow_y)
				rl.DrawPoly(v2, 4, uiobjs[a].circrad, uiobjs[a].rotation, rl.Fade(uiobjs[a].shadow_color, uiobjs[a].shadow_fade))
			}
			//shape
			if uiobjs[a].fill_color1 == blankcolor {
				rl.DrawPoly(uiobjs[a].circv2, 4, uiobjs[a].circrad, uiobjs[a].rotation, brightred())
			} else {
				rl.DrawPoly(uiobjs[a].circv2, 4, uiobjs[a].circrad, uiobjs[a].rotation, uiobjs[a].fill_color1)
			}
		case 4:
			//shadow
			if uiobjs[a].shadow {
				v2 := rl.NewVector2(uiobjs[a].circv2.X+uiobjs[a].shadow_x, uiobjs[a].circv2.Y+uiobjs[a].shadow_y)
				rl.DrawPoly(v2, 5, uiobjs[a].circrad, uiobjs[a].rotation, rl.Fade(uiobjs[a].shadow_color, uiobjs[a].shadow_fade))
			}
			//shape
			if uiobjs[a].fill_color1 == blankcolor {
				rl.DrawPoly(uiobjs[a].circv2, 5, uiobjs[a].circrad, uiobjs[a].rotation, brightred())
			} else {
				rl.DrawPoly(uiobjs[a].circv2, 5, uiobjs[a].circrad, uiobjs[a].rotation, uiobjs[a].fill_color1)
			}
		case 5:
			//shadow
			if uiobjs[a].shadow {
				v2 := rl.NewVector2(uiobjs[a].circv2.X+uiobjs[a].shadow_x, uiobjs[a].circv2.Y+uiobjs[a].shadow_y)
				rl.DrawPoly(v2, 6, uiobjs[a].circrad, uiobjs[a].rotation, rl.Fade(uiobjs[a].shadow_color, uiobjs[a].shadow_fade))
			}
			//shape
			if uiobjs[a].fill_color1 == blankcolor {
				rl.DrawPoly(uiobjs[a].circv2, 6, uiobjs[a].circrad, uiobjs[a].rotation, brightred())
			} else {
				rl.DrawPoly(uiobjs[a].circv2, 6, uiobjs[a].circrad, uiobjs[a].rotation, uiobjs[a].fill_color1)
			}
		case 6:
			//shadow
			if uiobjs[a].shadow {
				v2 := rl.NewVector2(uiobjs[a].circv2.X+uiobjs[a].shadow_x, uiobjs[a].circv2.Y+uiobjs[a].shadow_y)
				rl.DrawPoly(v2, 7, uiobjs[a].circrad, uiobjs[a].rotation, rl.Fade(uiobjs[a].shadow_color, uiobjs[a].shadow_fade))
			}
			//shape
			if uiobjs[a].fill_color1 == blankcolor {
				rl.DrawPoly(uiobjs[a].circv2, 7, uiobjs[a].circrad, uiobjs[a].rotation, brightred())
			} else {
				rl.DrawPoly(uiobjs[a].circv2, 7, uiobjs[a].circrad, uiobjs[a].rotation, uiobjs[a].fill_color1)
			}
		case 7:
			//shadow
			if uiobjs[a].shadow {
				v2 := rl.NewVector2(uiobjs[a].circv2.X+uiobjs[a].shadow_x, uiobjs[a].circv2.Y+uiobjs[a].shadow_y)
				rl.DrawPoly(v2, 8, uiobjs[a].circrad, uiobjs[a].rotation, rl.Fade(uiobjs[a].shadow_color, uiobjs[a].shadow_fade))
			}
			//shape
			if uiobjs[a].fill_color1 == blankcolor {
				rl.DrawPoly(uiobjs[a].circv2, 8, uiobjs[a].circrad, uiobjs[a].rotation, brightred())
			} else {
				rl.DrawPoly(uiobjs[a].circv2, 8, uiobjs[a].circrad, uiobjs[a].rotation, uiobjs[a].fill_color1)
			}
		case 8:
			//shadow
			if uiobjs[a].shadow {
				v2 := rl.NewVector2(uiobjs[a].circv2.X+uiobjs[a].shadow_x, uiobjs[a].circv2.Y+uiobjs[a].shadow_y)
				rl.DrawPoly(v2, 9, uiobjs[a].circrad, uiobjs[a].rotation, rl.Fade(uiobjs[a].shadow_color, uiobjs[a].shadow_fade))
			}
			//shape
			if uiobjs[a].fill_color1 == blankcolor {
				rl.DrawPoly(uiobjs[a].circv2, 9, uiobjs[a].circrad, uiobjs[a].rotation, brightred())
			} else {
				rl.DrawPoly(uiobjs[a].circv2, 9, uiobjs[a].circrad, uiobjs[a].rotation, uiobjs[a].fill_color1)
			}
		case 9:
			//shadow
			if uiobjs[a].shadow {
				v2 := rl.NewVector2(uiobjs[a].circv2.X+uiobjs[a].shadow_x, uiobjs[a].circv2.Y+uiobjs[a].shadow_y)
				rl.DrawPoly(v2, 10, uiobjs[a].circrad, uiobjs[a].rotation, rl.Fade(uiobjs[a].shadow_color, uiobjs[a].shadow_fade))
			}
			//shape
			if uiobjs[a].fill_color1 == blankcolor {
				rl.DrawPoly(uiobjs[a].circv2, 10, uiobjs[a].circrad, uiobjs[a].rotation, brightred())
			} else {
				rl.DrawPoly(uiobjs[a].circv2, 10, uiobjs[a].circrad, uiobjs[a].rotation, uiobjs[a].fill_color1)
			}
		}

		//selectrec
		if a == activuiobjnum {

			if uiobjs[a].shape == 0 {
				rec2 := rl.NewRectangle(uiobjs[a].rec.X-2, uiobjs[a].rec.Y-2, uiobjs[a].rec.Width+4, uiobjs[a].rec.Height+4)
				if onoff15 {
					rec2.X -= 4
					rec2.Y -= 4
					rec2.Width += 8
					rec2.Height += 8
				}
				rl.DrawRectangleLinesEx(rec2, 1.0, brightred())
			} else {

				circrad := uiobjs[a].circrad + 8
				if onoff15 {
					circrad += 8
				}

				rl.DrawCircleLines(int32(uiobjs[a].circv2.X), int32(uiobjs[a].circv2.Y), circrad, brightred())
			}

		}

	}

}
func dsettings() { //MARK: dsettings

	settingsnames := reflect.TypeOf(settings)
	fieldnames := make([]string, settingsnames.NumField())

	for i := range fieldnames {
		fieldnames[i] = settingsnames.Field(i).Name
	}

	x := txts * 2
	y := txts * 2
	txtspace := int32(150)

	for _, txt := range fieldnames {

		if txt == "gridon" {
			break
		}
		mouseselrec := rl.NewRectangle(float32(x-txts), float32(y-2), float32((txtspace+20)+(txts*2)), float32(txts+4))
		if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
			rl.DrawRectangleRec(mouseselrec, rl.Fade(brightred(), 0.2))
		}
		rl.DrawText(txt, x, y, txts, rl.White)

		switch txt {
		case "level_x_left":
			rl.DrawText(fmt.Sprint(settings.level_x_left), x+txtspace, y, txts, rl.White)
			if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

				}
			}
		case "level_y_top":
			rl.DrawText(fmt.Sprint(settings.level_y_top), x+txtspace, y, txts, rl.White)
			if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

				}
			}
		case "level_x_right":
			rl.DrawText(fmt.Sprint(settings.level_x_right), x+txtspace, y, txts, rl.White)
			if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

				}
			}
		case "level_y_bottom":
			rl.DrawText(fmt.Sprint(settings.level_y_bottom), x+txtspace, y, txts, rl.White)
			if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

				}
			}
		case "screen_width_multiplier":
			rl.DrawText(fmt.Sprint(settings.screen_width_multiplier), x+txtspace, y, txts, rl.White)
			if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

				}
			}
		case "screen_height_multiplier":
			rl.DrawText(fmt.Sprint(settings.screen_height_multiplier), x+txtspace, y, txts, rl.White)
			if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

				}
			}
		case "camera_x_change":
			rl.DrawText(fmt.Sprint(settings.camera_x_change), x+txtspace, y, txts, rl.White)
			if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					gettxton = true
					gettxtcharlimit = 5
					gettxtname = "camera_x_change"
					gettxttype = 2
				}
			}
		case "camera_y_change":
			rl.DrawText(fmt.Sprint(settings.camera_y_change), x+txtspace, y, txts, rl.White)
			if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					gettxton = true
					gettxtcharlimit = 5
					gettxtname = "camera_y_change"
					gettxttype = 2
				}
			}
		case "border_top":
			rl.DrawText(fmt.Sprint(settings.border_top), x+txtspace, y, txts, rl.White)
			if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					gettxton = true
					gettxtcharlimit = 5
					gettxtname = "border_top"
					gettxttype = 2
				}
			}
		case "border_left":
			rl.DrawText(fmt.Sprint(settings.border_left), x+txtspace, y, txts, rl.White)
			if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					gettxton = true
					gettxtcharlimit = 5
					gettxtname = "border_left"
					gettxttype = 2
				}
			}
		case "border_right":
			rl.DrawText(fmt.Sprint(settings.border_right), x+txtspace, y, txts, rl.White)
			if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					gettxton = true
					gettxtcharlimit = 5
					gettxtname = "border_right"
					gettxttype = 2
				}
			}
		case "border_bottom":
			rl.DrawText(fmt.Sprint(settings.border_bottom), x+txtspace, y, txts, rl.White)
			if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					gettxton = true
					gettxtcharlimit = 5
					gettxtname = "border_bottom"
					gettxttype = 2
				}
			}

		}

		y += txts + (txts / 2)

	}

	//closewin
	v2 := rl.NewVector2(scrwf32-closewinimg.Width*2, closewinimg.Height)
	rec2 := rl.NewRectangle(v2.X, v2.Y, closewinimg.Width, closewinimg.Height)
	if rl.CheckCollisionPointRec(mousev2, rec2) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			settingson = false
			menufocus = false
		}
		rl.DrawTextureRec(imgs, closewinimg, v2, brightred())
	} else {
		rl.DrawTextureRec(imgs, closewinimg, v2, rl.White)
	}

}
func dcolorpal() { //MARK: dcolorpal

	rl.DrawRectangleRec(colorpalbackrec, rl.Fade(darkred(), 0.4))

	colorpalrec.X = scrwf32 - (menu.rec.Width - 5)
	colorpalrec.Y = menu.rec.Y + float32(txts*4)

	txtlen := rl.MeasureText("random colors", txts)
	rl.DrawText("random colors", int32((colorpalrec.X+(colorpalrec.Width/2))-float32(txtlen/2)), colorpalrec.ToInt32().Y-(txts*2), txts, rl.White)

	if menu.lr {

	}

	recw := colorpalrec.Width / 8
	x := colorpalrec.X
	y := colorpalrec.Y

	if !colorscreated {
		makecolors(0)
	}

	rl.DrawRectangleRec(colorpalrec, rl.Black)

	//refresh colours
	v2 := rl.NewVector2(colorpalrec.X+colorpalrec.Width-(refreshimg.Width+4), colorpalrec.Y+colorpalrec.Height-(refreshimg.Height+4))
	refreshrec := rl.NewRectangle(v2.X, v2.Y, refreshimg.Width, refreshimg.Height)
	if rl.CheckCollisionPointRec(mousev2, refreshrec) {
		helptxt("refreshcolor")
		rl.DrawTextureRec(imgs, refreshimg, v2, brightred())
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			makecolors(1)
		}
	} else {
		rl.DrawTextureRec(imgs, refreshimg, v2, rl.White)
	}
	v2.X -= recw - 4
	// clear colours
	blankcolorrec := rl.NewRectangle(v2.X, v2.Y, closewinimg.Width, closewinimg.Height)
	if rl.CheckCollisionPointRec(mousev2, blankcolorrec) {
		helptxt("clearcolor")
		rl.DrawTextureRec(imgs, closewinimg, v2, brightred())
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			switch changecolorsel {
			case 15:
				objs[activobjnum].labelchangeshadowcolorobj2 = blankcolor
				changecolorsel = 0
			case 14:
				objs[activobjnum].labelchangeghostcolorobj2 = blankcolor
				changecolorsel = 0
			case 13:
				objs[activobjnum].labelchangefill2obj2 = blankcolor
				changecolorsel = 0
			case 12:
				objs[activobjnum].labelchangefill1obj2 = blankcolor
				changecolorsel = 0
			case 11:
				objs[activobjnum].labelchangeshadowcolor = blankcolor
				changecolorsel = 0
			case 10:
				objs[activobjnum].labelchangeghostcolor = blankcolor
				changecolorsel = 0
			case 9:
				objs[activobjnum].labelchangefill2 = blankcolor
				changecolorsel = 0
			case 8:
				objs[activobjnum].labelchangefill1 = blankcolor
				changecolorsel = 0

			case 6:
				objs[activobjnum].shadow_color = blankcolor
				changecolorsel = 0
			case 5:
				uiobjs[activuiobjnum].shadow_color = blankcolor
				changecolorsel = 0
			case 4:
				objs[activobjnum].ghosting_color = blankcolor
				if objs[activobjnum].complex {
					for a := 0; a < len(objs[activobjnum].objsin); a++ {
						objs[activobjnum].objsin[a].ghosting_color = blankcolor
					}
				}
				changecolorsel = 0
			case 1:
				objs[activobjnum].outline_color = blankcolor
				if objs[activobjnum].complex {
					for a := 0; a < len(objs[activobjnum].objsin); a++ {
						objs[activobjnum].objsin[a].outline_color = blankcolor
					}
				}
				changecolorsel = 0
			case 2:
				objs[activobjnum].fill_color1 = blankcolor
				if objs[activobjnum].complex {
					for a := 0; a < len(objs[activobjnum].objsin); a++ {
						objs[activobjnum].objsin[a].fill_color1 = blankcolor
					}
				}
				changecolorsel = 0
			case 3:
				objs[activobjnum].fill_color2 = blankcolor
				if objs[activobjnum].complex {
					for a := 0; a < len(objs[activobjnum].objsin); a++ {
						objs[activobjnum].objsin[a].fill_color2 = blankcolor
					}
				}
				changecolorsel = 0
			}
		}
	} else {
		rl.DrawTextureRec(imgs, closewinimg, v2, rl.White)
	}

	// random pallette
	for a := 0; a < len(colorpalrand); a++ {

		rec := rl.NewRectangle(x+1, y+1, recw-1, recw-1)
		rl.DrawRectangleRec(rec, colorpalrand[a].color)

		if rl.CheckCollisionPointRec(mousev2, rec) {
			rl.DrawRectangleRec(rec, rl.Fade(brightorange(), fadeblink))
			selcolor = colorpalrand[a]
			helptxt("color")
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

				switch changecolorsel {
				case 15:
					objs[activobjnum].labelchangeshadowcolorobj2 = colorpalrand[a].color
					changecolorsel = 0
				case 14:
					objs[activobjnum].labelchangeghostcolorobj2 = colorpalrand[a].color
					changecolorsel = 0
				case 13:
					objs[activobjnum].labelchangefill2obj2 = colorpalrand[a].color
					changecolorsel = 0
				case 12:
					objs[activobjnum].labelchangefill1obj2 = colorpalrand[a].color
					changecolorsel = 0
				case 11:
					objs[activobjnum].labelchangeshadowcolor = colorpalrand[a].color
					changecolorsel = 0
				case 10:
					objs[activobjnum].labelchangeghostcolor = colorpalrand[a].color
					changecolorsel = 0
				case 9:
					objs[activobjnum].labelchangefill2 = colorpalrand[a].color
					changecolorsel = 0
				case 8:
					objs[activobjnum].labelchangefill1 = colorpalrand[a].color
					changecolorsel = 0

				case 6:
					objs[activobjnum].shadow_color = colorpalrand[a].color
					changecolorsel = 0
				case 5:
					uiobjs[activuiobjnum].shadow_color = colorpalrand[a].color
					changecolorsel = 0
				case 4:
					objs[activobjnum].ghosting_color = colorpalrand[a].color
					if objs[activobjnum].complex {
						for a := 0; a < len(objs[activobjnum].objsin); a++ {
							objs[activobjnum].objsin[a].ghosting_color = objs[activobjnum].ghosting_color
						}
					}
					changecolorsel = 0
				case 1:
					objs[activobjnum].outline_color = colorpalrand[a].color
					if objs[activobjnum].complex {
						for a := 0; a < len(objs[activobjnum].objsin); a++ {
							objs[activobjnum].objsin[a].outline_color = objs[activobjnum].outline_color
						}
					}
					changecolorsel = 0
				case 2:
					objs[activobjnum].fill_color1 = colorpalrand[a].color
					objs[activobjnum].ghosting_color = objs[activobjnum].fill_color1
					if objs[activobjnum].complex {
						for a := 0; a < len(objs[activobjnum].objsin); a++ {
							objs[activobjnum].objsin[a].fill_color1 = objs[activobjnum].fill_color1
							objs[activobjnum].objsin[a].ghosting_color = objs[activobjnum].fill_color1
						}
					}
					changecolorsel = 0
				case 3:
					objs[activobjnum].fill_color2 = colorpalrand[a].color
					if objs[activobjnum].complex {
						for a := 0; a < len(objs[activobjnum].objsin); a++ {
							objs[activobjnum].objsin[a].fill_color2 = objs[activobjnum].fill_color2
						}
					}
					changecolorsel = 0
				}

				if !settings.colorpallockon {
					colorpalon = false
				}

			}
			if rl.IsMouseButtonPressed(rl.MouseRightButton) {
				colorpaluser[colorusercurrentnum] = colorpalrand[a]
				colorusercurrentnum++
				if colorusercurrentnum == len(colorpaluser) {
					colorusercurrentnum = 0
				}
			}
		}

		x += recw
		if x+recw > colorpalrec.X+colorpalrec.Width {
			x = colorpalrec.X
			y += recw
		}
	}

	xtxt := int32(colorpalrec.X + (colorpalrec.Width / 2))
	ytxt := int32(colorpalrec.Y + colorpalrec.Height + float32(txts*2))

	// standard pallette
	txtlen = rl.MeasureText("standard colors", txts)
	rl.DrawText("standard colors", xtxt-(txtlen/2), ytxt, txts, rl.White)

	x = colorpalrec.X
	y = float32(ytxt + (txts * 2))

	standcolorrec := rl.NewRectangle(x, y, colorpalbackrec.Width, recw*3)
	rl.DrawRectangleRec(standcolorrec, rl.Black)

	for a := 0; a < len(colorpalstand); a++ {

		rec := rl.NewRectangle(x+1, y+1, recw-1, recw-1)
		rl.DrawRectangleRec(rec, colorpalstand[a])

		if rl.CheckCollisionPointRec(mousev2, rec) {
			rl.DrawRectangleRec(rec, rl.Fade(brightorange(), fadeblink))

			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

				switch changecolorsel {
				case 6:
					objs[activobjnum].shadow_color = colorpalstand[a]
					changecolorsel = 0
				case 5:
					uiobjs[activuiobjnum].shadow_color = colorpalstand[a]
					changecolorsel = 0
				case 4:
					objs[activobjnum].ghosting_color = colorpalstand[a]
					if objs[activobjnum].complex {
						for a := 0; a < len(objs[activobjnum].objsin); a++ {
							objs[activobjnum].objsin[a].ghosting_color = objs[activobjnum].ghosting_color
						}
					}
					changecolorsel = 0
				case 1:
					objs[activobjnum].outline_color = colorpalstand[a]
					if objs[activobjnum].complex {
						for a := 0; a < len(objs[activobjnum].objsin); a++ {
							objs[activobjnum].objsin[a].outline_color = objs[activobjnum].outline_color
						}
					}
					changecolorsel = 0
				case 2:
					objs[activobjnum].fill_color1 = colorpalstand[a]
					objs[activobjnum].ghosting_color = objs[activobjnum].fill_color1
					if objs[activobjnum].complex {
						for a := 0; a < len(objs[activobjnum].objsin); a++ {
							objs[activobjnum].objsin[a].fill_color1 = objs[activobjnum].fill_color1
							objs[activobjnum].objsin[a].ghosting_color = objs[activobjnum].fill_color1
						}
					}
					changecolorsel = 0
				case 3:
					objs[activobjnum].fill_color2 = colorpalstand[a]
					if objs[activobjnum].complex {
						for a := 0; a < len(objs[activobjnum].objsin); a++ {
							objs[activobjnum].objsin[a].fill_color2 = objs[activobjnum].fill_color2
						}
					}
					changecolorsel = 0
				}

				if !settings.colorpallockon {
					colorpalon = false
				}

			}
		}

		x += recw
		if x+recw > colorpalrec.X+colorpalrec.Width {
			x = colorpalrec.X
			y += recw
		}
	}

	x = colorpalrec.X
	y += float32(txts * 4)

	xtxt = int32(x + (colorpalrec.Width / 2))
	ytxt = int32(y) - txts*2

	// user pallette
	txtlen = rl.MeasureText("user colors", txts)
	rl.DrawText("user colors", xtxt-(txtlen/2), ytxt, txts, rl.White)
	colorpaluserrec := rl.NewRectangle(x, y, colorpalrec.Width, colorpalrec.Height)
	rl.DrawRectangleRec(colorpaluserrec, rl.Black)

	for a := 0; a < len(colorpaluser); a++ {

		rec := rl.NewRectangle(x+1, y+1, recw-1, recw-1)
		if colorpaluser[a].color == blankcolor {
			rl.DrawRectangleRec(rec, rl.Fade(brightred(), 0.4))
			rl.DrawRectangleLinesEx(rec, 1.0, brightorange())
		} else {
			rl.DrawRectangleRec(rec, colorpaluser[a].color)
		}

		if rl.CheckCollisionPointRec(mousev2, rec) {
			rl.DrawRectangleRec(rec, rl.Fade(brightorange(), fadeblink))
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

				editcoloron = true

			}
		}

		x += recw
		if x+recw > colorpalrec.X+colorpalrec.Width {
			x = colorpalrec.X
			y += recw
		}
	}

}
func dfxmenu() { //MARK: dfxmenu

	x := int32(100)
	y := int32(100)

	rec := buttonline("scanlines", x, y)
	if rl.CheckCollisionPointRec(mousev2, rec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			upfxdemoonoff("scan")
		}
	}

	x += rec.ToInt32().Width + txts

	rec = buttonline("pixel noise", x, y)
	if rl.CheckCollisionPointRec(mousev2, rec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			upfxdemoonoff("pixel")
		}
	}

	x += rec.ToInt32().Width + txts

	rec = buttonline("ghosting", x, y)
	if rl.CheckCollisionPointRec(mousev2, rec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			upfxdemoonoff("ghost")
		}
	}

	//demo settings info box
	demosettingsrec := rl.NewRectangle(float32(x), float32(y+100), scrwf32/3, scrhf32/3)
	rl.DrawRectangleRec(demosettingsrec, rl.Black)

	txtx := demosettingsrec.X + 20
	txty := demosettingsrec.Y + 20

	if fxdemoscanlinesmenu {
		txthere("scanlines settings", false, txtm, txtx, txty, demosettingsrec.Width)
	} else if fxdemopixelnoisemenu {
		txthere("pixel noise settings", false, txtm, txtx, txty, demosettingsrec.Width)
	} else if fxdemoghost {
		txthere("ghosting settings", false, txtm, txtx, txty, demosettingsrec.Width)
	}

	//right side demo rec

	rl.DrawRectangleRec(fxdemorec, rl.Black)
	fxdemorec = rl.NewRectangle(scrwf32/2+40, 40, (scrwf32/2)-80, scrhf32-80)
	txthere("demo", false, txtm, fxdemorec.X+10, fxdemorec.Y+10, fxdemorec.Width)

	for a := 0; a < len(fxdemoobjs); a++ {

		if fxdemoobjs[a].imgon {

			destrec := rl.NewRectangle(float32(fxdemoobjs[a].x), float32(fxdemoobjs[a].y), fxdemoobjs[a].width, fxdemoobjs[a].width)
			origin := rl.NewVector2(fxdemoobjs[a].width/2, fxdemoobjs[a].width/2)

			rl.DrawTexturePro(imgs, fxdemoobjs[a].img, destrec, origin, fxdemoobjs[a].rotation, rl.Fade(fxdemoobjs[a].color, fxdemoobjs[a].fade))
		} else {
			switch fxdemoobjs[a].shape {

			case 2:
				rl.DrawCircle(fxdemoobjs[a].x, fxdemoobjs[a].y, fxdemoobjs[a].radius, rl.Fade(fxdemoobjs[a].color, fxdemoobjs[a].fade))
			case 3, 4, 5, 6, 7:
				v2 := rl.NewVector2(float32(fxdemoobjs[a].x), float32(fxdemoobjs[a].y))
				rl.DrawPoly(v2, fxdemoobjs[a].shape, fxdemoobjs[a].radius, fxdemoobjs[a].rotation, rl.Fade(fxdemoobjs[a].color, fxdemoobjs[a].fade))
			}
		}

		if fxdemoobjs[a].rotates {
			fxdemoobjs[a].rotation += fxdemoobjs[a].rotateamount
		}

	}

	if fxdemoscanlinesmenu {
		for _, checkfx := range defaultfx {
			if checkfx.num == 0 {
				for a := 0; a < len(checkfx.y); a++ {
					rl.DrawRectangle(fxdemorec.ToInt32().X, checkfx.y[a], fxdemorec.ToInt32().X+fxdemorec.ToInt32().Width, checkfx.height, rl.Fade(checkfx.color, checkfx.fade))
					checkfx.y[a] += checkfx.spacing
					if checkfx.y[a] > scrh {
						checkfx.y[a] = checkfx.spacing
					}
				}
			}
		}
	} else if fxdemopixelnoisemenu {

		for _, checkfx := range defaultfx {
			if checkfx.num == 1 {
				for a := 0; a < len(checkfx.yf32); a++ {

					width := rInt32(checkfx.min, checkfx.max)
					rl.DrawRectangle(int32(checkfx.xf32[a]+(scrwf32/2)), int32(checkfx.yf32[a]), width, width, rl.Fade(checkfx.color, checkfx.fade))

					checkfx.xf32[a] = rFloat32(0, scrwf32)
					checkfx.yf32[a] = rFloat32(0, scrhf32)
				}

			}
		}

	}

	//closewin
	v2 := rl.NewVector2(scrwf32-closewinimg.Width*2, closewinimg.Height)
	rec2 := rl.NewRectangle(v2.X, v2.Y, closewinimg.Width, closewinimg.Height)
	if rl.CheckCollisionPointRec(mousev2, rec2) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			fxmenuon = false
			menufocus = false
		}
		rl.DrawTextureRec(imgs, closewinimg, v2, brightred())
	} else {
		rl.DrawTextureRec(imgs, closewinimg, v2, rl.White)
	}
}

//MARK: EVENTS LABELS EVENTS LABELS EVENTS LABELS EVENNTS LABELS EVENTS LABELS EVENTS LABELS
func saveevent(num int) { //MARK: saveevent

	switch objs[activobjnum].events[num].event_type {
	case 0:
		if objs[activobjnum].events[num].name == "" {
			objs[activobjnum].events[num].name = "collision_boundary"
		}
	case 1:
		if objs[activobjnum].events[num].name == "" {
			objs[activobjnum].events[num].name = "collision_obj"
		}
	}

	neweventon = false
	addevent = false
	addingevent = false

}
func savelabeleventacitvobj() { //MARK: savelabeleventacitvobj

	objs[activobjnum].labeleventson = true

	newlabeleventdetails := labeleventdetails{}

	newlabeleventdetails.name = labeleventnametxt

	newlabeleventdetails.event = labeleventslist[labeleventnum]
	newlabeleventdetails.labelobj2 = labelsalllist[labelsalllistnum]
	newlabeleventdetails.action = labelactionslist[labelactionnum]
	newlabeleventdetails.actionobj2 = labelactionslistobj2[labelactionnumobj2]

	if objs[activobjnum].labelevent2 {
		newlabeleventdetails.action2 = labelactionslist[labelactionnum2]
	}
	if objs[activobjnum].labelevent3 {
		newlabeleventdetails.action3 = labelactionslist[labelactionnum3]
	}

	if objs[activobjnum].labeleventobj2 {
		newlabeleventdetails.action2obj2 = labelactionslistobj2[labelactionnum2obj2]
	}
	if objs[activobjnum].labelevent2obj2 {
		newlabeleventdetails.action3obj2 = labelactionslistobj2[labelactionnum3obj2]
	}

	if labelspawnon {
		newlabeleventdetails.spawnmore = true
		newlabeleventdetails.spawnobj = usrsaveobjs[usrsaveobjnum]
		newlabeleventdetails.spawnnum = labelspawnnum
	}
	if labelspawnon2 {
		newlabeleventdetails.spawnmore2 = true
		newlabeleventdetails.spawnobj2 = usrsaveobjs[usrsaveobjnum2]
		newlabeleventdetails.spawnnum2 = labelspawnnum2
	}
	if labelspawnon3 {
		newlabeleventdetails.spawnmore3 = true
		newlabeleventdetails.spawnobj3 = usrsaveobjs[usrsaveobjnum3]
		newlabeleventdetails.spawnnum3 = labelspawnnum3
	}
	if labelspawnon4 {
		newlabeleventdetails.spawnmore = true
		newlabeleventdetails.spawnobj4 = usrsaveobjs[usrsaveobjnum4]
		newlabeleventdetails.spawnnum4 = labelspawnnum4
	}

	objs[activobjnum].labelevents = append(objs[activobjnum].labelevents, newlabeleventdetails)

	labeleventon = false
	menufocus = false

	resetlabelevent()

}
func resetlabelevent() { //MARK: resetlabelevent

	labelspawnon, labelspawnon2, labelspawnon3, labelspawnon4 = false, false, false, false

	labeleventnum, labelsalllistnum, labelactionnum, labelactionnumobj2, usrsaveobjnum, usrsaveobjnum2, usrsaveobjnum3, usrsaveobjnum4 = 0, 0, 0, 0, 0, 0, 0, 0

	labelspawnnum, labelspawnnum2, labelspawnnum3, labelspawnnum4 = 1, 1, 1, 1

	deflabeleventnamechanged = deflabeleventname
}

//MARK: MENUS MENUS MENUS MENUS MENUS MENUS MENUS MENUS MENUS MENUS MENUS MENUS MENUS MENUS
func mainmenu() { //MARK: mainmenu

	//menu items
	menux := int32(menu.rec.Width)
	menuy := int32(menu.rec.Y) + txts
	x, y, rec := mainmenuitem("obj+", menux, menuy, newobjon)
	if rl.CheckCollisionPointRec(mousev2, rec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			upmainmenuinfo("objmenuon")
			if newobjon {
				newobjon = false
			} else {
				newobjon = true
			}
		}
	}
	x, y, rec = mainmenuitem("ui +", x, y, newuiobjon)
	if rl.CheckCollisionPointRec(mousev2, rec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			upmainmenuinfo("uiobjmenuon")
			if newuiobjon {
				newuiobjon = false
			} else {
				newuiobjon = true
			}

		}
	}

	x, y, rec = mainmenuitem("text+", x, y, newtxtobj)
	if rl.CheckCollisionPointRec(mousev2, rec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			upmainmenuinfo("textmenuon")
			if newtxtobj {
				newtxtobj = false
			} else {
				newtxtobj = true
			}
		}
	}

	x, y, rec = mainmenuitem("timer+", x, y, newuiobjon)
	if rl.CheckCollisionPointRec(mousev2, rec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			upmainmenuinfo("timermenuon")
			if newtimerobj {
				newtimerobj = false
			} else {
				newtimerobj = true
			}
		}
	}
	x, y, rec = mainmenuitem("fx+", x, y, newuiobjon)
	if rl.CheckCollisionPointRec(mousev2, rec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			makefxdemo()
			fxmenuon = true
			menufocus = true
		}
	}

	if uiobjmenuon {
		//ui obj info
		if activuiobjnum != blanknum {
			y += txts * 2
			x = menu.rec.ToInt32().X + 5

			uiobjfieldnames := reflect.TypeOf(uiobjs[activuiobjnum])
			fieldnames2 := make([]string, uiobjfieldnames.NumField())
			fieldnames := make([]string, 0)

			for i := range fieldnames2 {
				if uiobjfieldnames.Field(i).Name == "rec" {
					break
				}
				fieldnames = append(fieldnames, uiobjfieldnames.Field(i).Name)
			}

			height := float32((txts + (txts / 2)) * int32(len(fieldnames)))
			height += float32(txts * 2)

			infoboxrec := rl.NewRectangle(float32(x), float32(y), menu.rec.Width-10, height)
			rl.DrawRectangleRec(infoboxrec, rl.Fade(rl.Black, 0.7))

			x += txts
			y += txts

			for _, txt := range fieldnames {
				mouseselrec := rl.NewRectangle(infoboxrec.X, float32(y-2), infoboxrec.Width, float32(txts+4))
				if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
					rl.DrawRectangleRec(mouseselrec, rl.Fade(brightred(), 0.2))
				}
				rl.DrawText(txt, x, y, txts, rl.White)

				switch txt {
				case "shadow_color":
					if uiobjs[activuiobjnum].shadow_color == blankcolor {
						txtlen := rl.MeasureText("choose color", txts)
						rl.DrawText("choose color", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					}
					rl.DrawRectangle(int32(menu.rec.X+menu.rec.Width-(togglewid+float32(txts))), y, int32(togglewid), int32(togglewid), uiobjs[activuiobjnum].shadow_color)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							changecolorsel = 5
							colorpalon = true
						}
					}

				case "shadow_fade":
					txt := fmt.Sprint(uiobjs[activuiobjnum].shadow_fade)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 1
							gettxtname = "uiobj_shadow_fade"
							gettxttype = 2
						}
					}
				case "shadow_y":
					txt := fmt.Sprint(uiobjs[activuiobjnum].shadow_y)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 4
							gettxtname = "uiobj_shadow_y"
							gettxttype = 5
						}
					}
				case "shadow_x":
					txt := fmt.Sprint(uiobjs[activuiobjnum].shadow_x)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 4
							gettxtname = "uiobj_shadow_x"
							gettxttype = 5
						}
					}
				case "center_y":

					txt := fmt.Sprint(uiobjs[activuiobjnum].circv2.Y)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 3
							gettxtname = "uiobj_center_y"
							gettxttype = 2
						}
					}
				case "center_x":
					txt := fmt.Sprint(uiobjs[activuiobjnum].circv2.X)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 3
							gettxtname = "uiobj_center_x"
							gettxttype = 2
						}
					}
				case "topleft_x":
					txt := fmt.Sprint(uiobjs[activuiobjnum].rec.X)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 3
							gettxtname = "uiobj_topleft_x"
							gettxttype = 2
						}
					}
				case "topleft_y":
					txt := fmt.Sprint(uiobjs[activuiobjnum].rec.Y)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 3
							gettxtname = "uiobj_topleft_y"
							gettxttype = 2
						}
					}
				case "polygon_circ_radius":
					num := uiobjs[activuiobjnum].rec.Width / 2
					txt := fmt.Sprint(num)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 3
							gettxtname = "uiobj_polygon_circ_radius"
							gettxttype = 2
						}
					}
				case "rec_height":
					txt := fmt.Sprint(uiobjs[activuiobjnum].rec.Height)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 3
							gettxtname = "uiobj_rech"
							gettxttype = 2
						}
					}
				case "rec_width":
					txt := fmt.Sprint(uiobjs[activuiobjnum].rec.Width)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 3
							gettxtname = "uiobj_recw"
							gettxttype = 2
						}
					}
				case "rotation":
					txt := fmt.Sprint(uiobjs[activuiobjnum].rotation)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) && uiobjs[activuiobjnum].shape != 0 && uiobjs[activuiobjnum].shape != 1 {
							gettxton = true
							gettxtcharlimit = 3
							gettxtname = "uiobj_rotation"
							gettxttype = 2
						} else if uiobjs[activuiobjnum].shape == 1 {
							helptxt("uiobj_circ_rotation")
						} else if uiobjs[activuiobjnum].shape == 0 {
							helptxt("uiobj_rec_rotation")
						}
					}

				case "shadow":
					rec = toggle("", int32(menu.rec.X)+menu.rec.ToInt32().Width-30, y, uiobjs[activuiobjnum].shadow)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							if uiobjs[activuiobjnum].shadow {
								uiobjs[activuiobjnum].shadow = false
							} else {
								uiobjs[activuiobjnum].shadow = true
							}

						}
					}
				case "name":
					if uiobjs[activuiobjnum].name == "" {
						txtlen := rl.MeasureText("_________________", txts)
						rl.DrawText("_________________", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen-(txts/2), y, txts, rl.White)
					} else {
						txtlen := rl.MeasureText(uiobjs[activuiobjnum].name, txts)
						rl.DrawText(uiobjs[activuiobjnum].name, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					}
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 20
							gettxtname = "uiobjname"
							gettxttype = 3
						}
					}

				case "shape":
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							uiobjs[activuiobjnum].shape++
							if uiobjs[activuiobjnum].shape == 10 {
								uiobjs[activuiobjnum].shape = 0
							}
						}
					}
					switch uiobjs[activuiobjnum].shape {
					case 0:
						txtlen := rl.MeasureText("rectangle", txts)
						rl.DrawText("rectangle", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					case 1:
						txtlen := rl.MeasureText("circle", txts)
						rl.DrawText("circle", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					case 2:
						txtlen := rl.MeasureText("triangle", txts)
						rl.DrawText("triangle", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					case 3:
						txtlen := rl.MeasureText("square polygon", txts)
						rl.DrawText("square polygon", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					case 4:
						txtlen := rl.MeasureText("pentagon", txts)
						rl.DrawText("pentagon", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					case 5:
						txtlen := rl.MeasureText("hexagon", txts)
						rl.DrawText("hexagon", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					case 6:
						txtlen := rl.MeasureText("septagon", txts)
						rl.DrawText("septagon", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					case 7:
						txtlen := rl.MeasureText("octagon", txts)
						rl.DrawText("octagon", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					case 8:
						txtlen := rl.MeasureText("nonagon", txts)
						rl.DrawText("nonagon", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					case 9:
						txtlen := rl.MeasureText("decagon", txts)
						rl.DrawText("decagon", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					}

				}

				y += txts
				y += txts / 2
			}

		}
	}
	if objmenuon && !neweventon && !addobjcontrols && !addobjpath && !addlabelon { // obj info

		if activobjnum != blanknum {
			y += txts * 2
			x = menu.rec.ToInt32().X + 5

			objfieldnames := reflect.TypeOf(objs[activobjnum])
			fieldnames2 := make([]string, objfieldnames.NumField())
			fieldnames := make([]string, 0)

			for i := range fieldnames2 {
				if objfieldnames.Field(i).Name == "rec" {
					break
				}
				fieldnames = append(fieldnames, objfieldnames.Field(i).Name)
			}

			height := float32((txts + (txts / 2)) * int32(len(fieldnames)))
			height += float32(txts * 2)

			infoboxrec := rl.NewRectangle(float32(x), float32(y), menu.rec.Width-10, height)
			rl.DrawRectangleRec(infoboxrec, rl.Fade(rl.Black, 0.7))

			x += txts
			y += txts

			for _, txt := range fieldnames {

				mouseselrec := rl.NewRectangle(infoboxrec.X, float32(y-2), infoboxrec.Width, float32(txts+4))
				if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
					rl.DrawRectangleRec(mouseselrec, rl.Fade(brightred(), 0.2))
				}

				switch txt {
				case "hp":
					txt := strconv.Itoa(objs[activobjnum].hp)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 3
							gettxtname = "hp_activobj"
							gettxttype = 2
						}
					}
				case "path_move":
					rec = toggle("", int32(menu.rec.X)+menu.rec.ToInt32().Width-30, y, objs[activobjnum].path_move)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							if objs[activobjnum].path_move {
								objs[activobjnum].path_move = false
							} else {
								objs[activobjnum].path_move = true
								if objs[activobjnum].direction_x == 0 {
									objs[activobjnum].direction_x = 1
								}
								if objs[activobjnum].direction_y == 0 {
									objs[activobjnum].direction_y = 1
								}
							}
						}
					}
				case "rotate_speed":
					txt := fmt.Sprint(objs[activobjnum].rotate_speed)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 3
							gettxtname = "rotate_speed"
							gettxttype = 2
						}
					}
				case "rotate_timer":
					txt := fmt.Sprint(objs[activobjnum].rotate_timer)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 3
							gettxtname = "rotate_timer"
							gettxttype = 2
						}
					}
				case "rotate_rand":
					rec = toggle("", int32(menu.rec.X)+menu.rec.ToInt32().Width-30, y, objs[activobjnum].rotate_rand)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							if objs[activobjnum].rotate_rand {
								objs[activobjnum].rotate_rand = false
							} else {
								objs[activobjnum].rotate_rand = true
								if objs[activobjnum].rotates == false {
									objs[activobjnum].rotates = true
								}
								if objs[activobjnum].rotate_speed == 0 {
									objs[activobjnum].rotate_speed = settings.rotate_speed
									objs[activobjnum].orig_rotatespeed = objs[activobjnum].rotate_speed
								}
								if objs[activobjnum].rotate_timer == 0 {
									objs[activobjnum].rotate_timer = settings.rotate_timer
									objs[activobjnum].orig_rotatetimer = objs[activobjnum].rotate_timer
								}
							}
						}
					}
				case "rotate_lr":
					if objs[activobjnum].rotate_lr {
						txtlen := rl.MeasureText("left", txts)
						rl.DrawText("left", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					} else {
						txtlen := rl.MeasureText("right", txts)
						rl.DrawText("right", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					}
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							if objs[activobjnum].rotate_lr {
								objs[activobjnum].rotate_lr = false
							} else {
								objs[activobjnum].rotate_lr = true
							}
						}
					}
				case "rotates":
					rec = toggle("", int32(menu.rec.X)+menu.rec.ToInt32().Width-30, y, objs[activobjnum].rotates)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							if objs[activobjnum].rotates {
								objs[activobjnum].rotates = false
							} else {
								objs[activobjnum].rotates = true
								if objs[activobjnum].rotate_speed == 0 {
									objs[activobjnum].rotate_speed = 1
								}
							}
						}
					}
				case "rotation":
					txt := fmt.Sprint(objs[activobjnum].rotation)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 5
							gettxtname = "obj_rotation"
							gettxttype = 4
						}
					}
				case "outline_only":
					rec = toggle("", int32(menu.rec.X)+menu.rec.ToInt32().Width-30, y, objs[activobjnum].outline_only)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							if objs[activobjnum].outline_only {
								objs[activobjnum].outline_only = false
							} else {
								objs[activobjnum].outline_only = true
							}
						}
					}
				case "foreground_obj":
					rec = toggle("", int32(menu.rec.X)+menu.rec.ToInt32().Width-30, y, objs[activobjnum].foreground_obj)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							if objs[activobjnum].foreground_obj {
								objs[activobjnum].background_obj = false
								objs[activobjnum].foreground_obj = false
								objs[activobjnum].middleground_obj = true
							} else {
								objs[activobjnum].background_obj = false
								objs[activobjnum].foreground_obj = true
								objs[activobjnum].middleground_obj = false
							}
						}
					}
				case "background_obj":
					rec = toggle("", int32(menu.rec.X)+menu.rec.ToInt32().Width-30, y, objs[activobjnum].background_obj)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							if objs[activobjnum].background_obj {
								objs[activobjnum].background_obj = false
								objs[activobjnum].foreground_obj = false
								objs[activobjnum].middleground_obj = true
							} else {
								objs[activobjnum].background_obj = true
								objs[activobjnum].foreground_obj = false
								objs[activobjnum].middleground_obj = false
							}
						}
					}
				case "middleground_obj":
					rec = toggle("", int32(menu.rec.X)+menu.rec.ToInt32().Width-30, y, objs[activobjnum].middleground_obj)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							if objs[activobjnum].middleground_obj {
								objs[activobjnum].background_obj = true
								objs[activobjnum].foreground_obj = false
								objs[activobjnum].middleground_obj = false

							} else {
								objs[activobjnum].background_obj = false
								objs[activobjnum].foreground_obj = false
								objs[activobjnum].middleground_obj = true
							}
						}
					}
				case "shape":
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							objs[activobjnum].shape++
							if objs[activobjnum].shape == 10 {
								objs[activobjnum].shape = 0
							}
						}
					}
					switch objs[activobjnum].shape {
					case 0:
						txtlen := rl.MeasureText("rectangle", txts)
						rl.DrawText("rectangle", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					case 1:
						txtlen := rl.MeasureText("circle", txts)
						rl.DrawText("circle", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					case 2:
						txtlen := rl.MeasureText("triangle", txts)
						rl.DrawText("triangle", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					case 3:
						txtlen := rl.MeasureText("square polygon", txts)
						rl.DrawText("square polygon", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					case 4:
						txtlen := rl.MeasureText("pentagon", txts)
						rl.DrawText("pentagon", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					case 5:
						txtlen := rl.MeasureText("hexagon", txts)
						rl.DrawText("hexagon", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					case 6:
						txtlen := rl.MeasureText("septagon", txts)
						rl.DrawText("septagon", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					case 7:
						txtlen := rl.MeasureText("octagon", txts)
						rl.DrawText("octagon", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					case 8:
						txtlen := rl.MeasureText("nonagon", txts)
						rl.DrawText("nonagon", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					case 9:
						txtlen := rl.MeasureText("decagon", txts)
						rl.DrawText("decagon", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					}
				case "rand_direc_time":
					txt := fmt.Sprint(objs[activobjnum].rand_direc_time)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {

					}
				case "rand_direc_timer_min":
					txt := fmt.Sprint(objs[activobjnum].rand_direc_timer_min)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 3
							gettxtname = "rand_direc_timer_min"
							gettxttype = 2
						}
					}
				case "rand_direc_timer_max":
					txt := fmt.Sprint(objs[activobjnum].rand_direc_timer_max)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 3
							gettxtname = "rand_direc_timer_max"
							gettxttype = 2
						}
					}
				case "rand_direc_max_y":
					txt := fmt.Sprint(objs[activobjnum].rand_direc_max_y)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 3
							gettxtname = "rand_direc_max_y"
							gettxttype = 2
						}
					}
				case "rand_direc_max_x":
					txt := fmt.Sprint(objs[activobjnum].rand_direc_max_x)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 3
							gettxtname = "rand_direc_max_x"
							gettxttype = 2
						}
					}
				case "random_direction":
					rec = toggle("", int32(menu.rec.X)+menu.rec.ToInt32().Width-30, y, objs[activobjnum].random_direction)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							if objs[activobjnum].random_direction {
								objs[activobjnum].random_direction = false
							} else {
								objs[activobjnum].random_direction = true
								objs[activobjnum].rand_direc_time = 0
							}
						}
					}
				case "shadow":
					rec = toggle("", int32(menu.rec.X)+menu.rec.ToInt32().Width-30, y, objs[activobjnum].shadow)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							if objs[activobjnum].shadow {
								objs[activobjnum].shadow = false
							} else {
								objs[activobjnum].shadow = true
							}
						}
					}
				case "shadow_color":
					if objs[activobjnum].shadow_color == blankcolor {
						txtlen := rl.MeasureText("choose color", txts)
						rl.DrawText("choose color", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					}
					rl.DrawRectangle(int32(menu.rec.X+menu.rec.Width-(togglewid+float32(txts))), y, int32(togglewid), int32(togglewid), objs[activobjnum].shadow_color)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							changecolorsel = 6
							colorpalon = true
						}
					}

				case "shadow_fade":
					txt := fmt.Sprint(objs[activobjnum].shadow_fade)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 1
							gettxtname = "shadow_fade"
							gettxttype = 2
						}
					}
				case "shadow_y":
					txt := fmt.Sprint(objs[activobjnum].shadow_y)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 4
							gettxtname = "shadow_y"
							gettxttype = 5
						}
					}
				case "shadow_x":
					txt := fmt.Sprint(objs[activobjnum].shadow_x)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 4
							gettxtname = "shadow_x"
							gettxttype = 5
						}
					}

				case "fade":
					txt := fmt.Sprint(objs[activobjnum].fade)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 1
							gettxtname = "fade"
							gettxttype = 2
						}
					}
				case "ghosting_color":
					if objs[activobjnum].ghosting_color == blankcolor {
						txtlen := rl.MeasureText("choose color", txts)
						rl.DrawText("choose color", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					}
					rl.DrawRectangle(int32(menu.rec.X+menu.rec.Width-(togglewid+float32(txts))), y, int32(togglewid), int32(togglewid), objs[activobjnum].ghosting_color)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							changecolorsel = 4
							colorpalon = true
						}
					}

				case "ghosting_fade":
					txt := fmt.Sprint(objs[activobjnum].ghosting_fade)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 1
							gettxtname = "ghosting_fade"
							gettxttype = 2
						}
					}

				case "ghosting_y":
					txt := fmt.Sprint(objs[activobjnum].ghosting_y)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 3
							gettxtname = "ghosting_y"
							gettxttype = 2
						}
					}

				case "ghosting_x":
					txt := fmt.Sprint(objs[activobjnum].ghosting_x)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 3
							gettxtname = "ghosting_x"
							gettxttype = 2
						}
					}

				case "ghosting":

					rec = toggle("", int32(menu.rec.X)+menu.rec.ToInt32().Width-30, y, objs[activobjnum].ghosting)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							if objs[activobjnum].ghosting {
								objs[activobjnum].ghosting = false
							} else {
								objs[activobjnum].ghosting = true
							}
							if objs[activobjnum].complex {
								for a := 0; a < len(objs[activobjnum].objsin); a++ {
									objs[activobjnum].objsin[a].ghosting = objs[activobjnum].ghosting
									objs[activobjnum].objsin[a].ghosting_color = objs[activobjnum].fill_color1
								}
							}
						}
					}

				case "tile_h":
					txt := ""
					if !objs[activobjnum].complex {
						txt = fmt.Sprint(objs[activobjnum].rec.Height)
					} else {
						txt = fmt.Sprint(objs[activobjnum].tile_h)
					}
					txtlen := rl.MeasureText(txt, txts)

					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) && objs[activobjnum].complex {
							gettxton = true
							gettxtcharlimit = 5
							gettxtname = "tile_h"
							gettxttype = 2
						}
					}
				case "tile_w":
					txt := ""
					if !objs[activobjnum].complex {
						txt = fmt.Sprint(objs[activobjnum].rec.Width)
					} else {
						txt = fmt.Sprint(objs[activobjnum].tile_w)
					}
					txtlen := rl.MeasureText(txt, txts)

					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) && objs[activobjnum].complex {
							gettxton = true
							gettxtcharlimit = 5
							gettxtname = "tile_w"
							gettxttype = 2
						}
					}
				case "orig_direcx":
					txt := fmt.Sprint(objs[activobjnum].orig_direcx)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)

				case "orig_direcy":
					txt := fmt.Sprint(objs[activobjnum].orig_direcy)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)

				case "direction_x":
					txt := fmt.Sprint(objs[activobjnum].direction_x)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 5
							gettxtname = "direction_x"
							gettxttype = 5
						}
					}
				case "direction_y":
					txt := fmt.Sprint(objs[activobjnum].direction_y)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 5
							gettxtname = "direction_y"
							gettxttype = 5
						}
					}
				case "img_rotate_speed":
					txt := fmt.Sprint(objs[activobjnum].img_rotate_speed)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 5
							gettxtname = "img_rotate_speed"
							gettxttype = 4
							if objs[activobjnum].complex {
								for a := 0; a < len(objs[activobjnum].objsin); a++ {
									objs[activobjnum].objsin[a].img_rotate_speed = objs[activobjnum].img_rotate_speed
								}
							}
						}
					}
				case "img_rotate_lr":
					if objs[activobjnum].img_rotate_lr {
						txtlen := rl.MeasureText("left", txts)
						rl.DrawText("left", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					} else {
						txtlen := rl.MeasureText("right", txts)
						rl.DrawText("right", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					}
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							if objs[activobjnum].img_rotate_lr {
								objs[activobjnum].img_rotate_lr = false
							} else {
								objs[activobjnum].img_rotate_lr = true
							}
							if objs[activobjnum].complex {
								for a := 0; a < len(objs[activobjnum].objsin); a++ {
									objs[activobjnum].objsin[a].img_rotate_lr = objs[activobjnum].img_rotate_lr
								}
							}
						}
					}
				case "img_rotates":
					rec = toggle("", int32(menu.rec.X)+menu.rec.ToInt32().Width-30, y, objs[activobjnum].img_rotates)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							if objs[activobjnum].img_rotates {
								objs[activobjnum].img_rotates = false
							} else {
								objs[activobjnum].img_rotates = true
							}
							if objs[activobjnum].complex {
								for a := 0; a < len(objs[activobjnum].objsin); a++ {
									objs[activobjnum].objsin[a].img_rotates = objs[activobjnum].img_rotates
								}
							}
						}
					}
				case "img_rotation":
					txt := fmt.Sprint(objs[activobjnum].img_rotation)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 5
							gettxtname = "img_rotation"
							gettxttype = 4
						}
					}
				case "img":
					txtlen := rl.MeasureText("choose img", txts)
					rl.DrawText("choose img", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							tileselecton = true
							menufocus = true
						}
					}
				case "outline_w":
					txt := fmt.Sprint(objs[activobjnum].outline_w)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 5
							gettxtname = "outline_w"
							gettxttype = 4
						}
					}
				case "topleft_x":
					txt := fmt.Sprint(objs[activobjnum].rec.X)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 5
							gettxtname = "topleft_x"
							gettxttype = 4
						}
					}
				case "topleft_y":
					txt := fmt.Sprint(objs[activobjnum].rec.Y)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 5
							gettxtname = "topleft_y"
							gettxttype = 4
						}
					}
				case "height":
					txt := fmt.Sprint(objs[activobjnum].rec.Height)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 5
							gettxtname = "obj_h"
							gettxttype = 4
						}
					}
				case "width":
					txt := fmt.Sprint(objs[activobjnum].rec.Width)
					txtlen := rl.MeasureText(txt, txts)
					rl.DrawText(txt, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 5
							gettxtname = "obj_w"
							gettxttype = 4
						}
					}
				case "gradient_h":
					rec = toggle("", int32(menu.rec.X)+menu.rec.ToInt32().Width-30, y, objs[activobjnum].gradient_h)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							if objs[activobjnum].gradient_h {
								objs[activobjnum].gradient_h = false
								if objs[activobjnum].complex {
									for a := 0; a < len(objs[activobjnum].objsin); a++ {
										objs[activobjnum].objsin[a].gradient_h = false
									}
								}
							} else {
								objs[activobjnum].gradient_h = true
								objs[activobjnum].gradient_v = false
								for a := 0; a < len(objs[activobjnum].objsin); a++ {
									objs[activobjnum].objsin[a].gradient_h = true
									objs[activobjnum].objsin[a].gradient_v = false
								}
							}
						}
					}
				case "gradient_v":
					rec = toggle("", int32(menu.rec.X)+menu.rec.ToInt32().Width-30, y, objs[activobjnum].gradient_v)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							if objs[activobjnum].gradient_v {
								objs[activobjnum].gradient_v = false
								if objs[activobjnum].complex {
									for a := 0; a < len(objs[activobjnum].objsin); a++ {
										objs[activobjnum].objsin[a].gradient_v = false
									}
								}
							} else {
								objs[activobjnum].gradient_v = true
								objs[activobjnum].gradient_h = false
								for a := 0; a < len(objs[activobjnum].objsin); a++ {
									objs[activobjnum].objsin[a].gradient_h = false
									objs[activobjnum].objsin[a].gradient_v = true
								}
							}
						}
					}
				case "fill_color1":
					if objs[activobjnum].fill_color1 == blankcolor {
						txtlen := rl.MeasureText("choose color", txts)
						rl.DrawText("choose color", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					}
					rl.DrawRectangle(int32(menu.rec.X+menu.rec.Width-(togglewid+float32(txts))), y, int32(togglewid), int32(togglewid), objs[activobjnum].fill_color1)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							changecolorsel = 2
							colorpalon = true
						}
					}
				case "fill_color2":
					if objs[activobjnum].fill_color2 == blankcolor {
						txtlen := rl.MeasureText("choose color", txts)
						rl.DrawText("choose color", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					}
					rl.DrawRectangle(int32(menu.rec.X+menu.rec.Width-(togglewid+float32(txts))), y, int32(togglewid), int32(togglewid), objs[activobjnum].fill_color2)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							changecolorsel = 3
							colorpalon = true
						}
					}
				case "outline_color":
					if objs[activobjnum].outline_color == blankcolor {
						txtlen := rl.MeasureText("choose color", txts)
						rl.DrawText("choose color", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, brightorange())
					}
					rl.DrawRectangle(int32(menu.rec.X+menu.rec.Width-(togglewid+float32(txts))), y, int32(togglewid), int32(togglewid), objs[activobjnum].outline_color)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							changecolorsel = 1
							colorpalon = true
						}
					}
				case "name":
					if objs[activobjnum].name == "" {
						txtlen := rl.MeasureText("_________________", txts)
						rl.DrawText("_________________", (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen-(txts/2), y, txts, rl.White)
					} else {
						txtlen := rl.MeasureText(objs[activobjnum].name, txts)
						rl.DrawText(objs[activobjnum].name, (infoboxrec.ToInt32().X+infoboxrec.ToInt32().Width)-5-txtlen, y, txts, rl.White)
					}
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							gettxton = true
							gettxtcharlimit = 20
							gettxtname = "objname"
							gettxttype = 3
						}
					}
				case "complex":
					rec = toggle("", int32(menu.rec.X)+menu.rec.ToInt32().Width-30, y, objs[activobjnum].complex)
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							if objs[activobjnum].complex {
								objs[activobjnum].complex = false
								clearobjsin()
							} else {
								objs[activobjnum].complex = true
								makeobjsin()
							}
						}
					}
				}

				rl.DrawText(txt, x, y, txts, rl.White)
				y += txts
				y += txts / 2

			}
		}
	} else if objmenuon && !neweventon && addobjcontrols && !addobjpath && !addlabelon { // add controls
		y += txts * 2
		x = menu.rec.ToInt32().X + 5

		addkeysrec := rl.NewRectangle(float32(x), float32(y), menu.rec.Width-10, float32(txts*2))
		if addevent || rl.CheckCollisionPointRec(mousev2, addkeysrec) {
			rl.DrawRectangleRec(addkeysrec, brightred())
			rl.DrawRectangleLinesEx(addkeysrec, 1.0, rl.White)

		} else {
			rl.DrawRectangleRec(addkeysrec, rl.Black)
		}

		if rl.CheckCollisionPointRec(mousev2, addkeysrec) {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				createcontrols(activobjnum, 0)

			}
		}

		txtlen := rl.MeasureText("keys +", txts)
		rl.DrawText("keys +", int32(menu.rec.X+menu.rec.Width/2)-txtlen/2, y+txts/2, txts, rl.White)

	} else if objmenuon && neweventon && !addobjcontrols && !addobjpath && !addlabelon { // add events
		y += txts * 2
		x = menu.rec.ToInt32().X + 5

		//add event rec
		neweventrec := rl.NewRectangle(float32(x), float32(y), menu.rec.Width-10, float32(txts*2))
		if addevent || rl.CheckCollisionPointRec(mousev2, neweventrec) {
			rl.DrawRectangleRec(neweventrec, brightred())
			rl.DrawRectangleLinesEx(neweventrec, 1.0, rl.White)

		} else {
			rl.DrawRectangleRec(neweventrec, rl.Black)
		}

		if rl.CheckCollisionPointRec(mousev2, neweventrec) {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				if addevent {
					addevent = false
					addingevent = false
				} else {
					addevent = true
					newevent := event{}
					objs[activobjnum].events = append(objs[activobjnum].events, newevent)
				}
			}
		}

		txtlen := rl.MeasureText("new event +", txts)
		rl.DrawText("new event +", int32(menu.rec.X+menu.rec.Width/2)-txtlen/2, y+txts/2, txts, rl.White)

		//obj events list
		if !addevent {
			if len(objs[activobjnum].events) > 0 {

				y = int32(neweventrec.Y + neweventrec.Height + float32(txts))

				height := float32((txts + (txts / 2)) * int32(len(objs[activobjnum].events)))
				height += float32(txts * 2)

				infoboxrec := rl.NewRectangle(float32(x), float32(y), menu.rec.Width-10, height)
				rl.DrawRectangleRec(infoboxrec, rl.Fade(rl.Black, 0.7))

				x += txts
				y += txts

				for a := 0; a < len(objs[activobjnum].events); a++ {
					mouseselrec := rl.NewRectangle(infoboxrec.X, float32(y-2), infoboxrec.Width, float32(txts+4))
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						rl.DrawRectangleRec(mouseselrec, rl.Fade(brightred(), 0.2))
					}
					rl.DrawText(objs[activobjnum].events[a].name, x, y, txts, rl.White)
					y += txts
					y += txts / 2
				}
			}
		}

		//add events list
		y += txts * 3
		if addevent && !addingevent {
			height := float32((txts + (txts / 2)) * int32(len(eventslist)))
			height += float32(txts * 2)

			infoboxrec := rl.NewRectangle(float32(x), float32(y), menu.rec.Width-10, height)
			rl.DrawRectangleRec(infoboxrec, rl.Fade(rl.Black, 0.7))

			x += txts
			y += txts

			for a := 0; a < len(eventslist); a++ {
				mouseselrec := rl.NewRectangle(infoboxrec.X, float32(y-2), infoboxrec.Width, float32(txts+4))
				if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
					rl.DrawRectangleRec(mouseselrec, rl.Fade(brightred(), 0.2))
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						switcheventsmenu(a)
						addingevent = true
					}
				}

				rl.DrawText(eventslist[a], x, y, txts, rl.White)
				y += txts
				y += txts / 2
			}
		} else if addevent && addingevent {

			num := 0
			for a := 0; a < len(eventsonoff); a++ {
				if eventsonoff[a] {
					num = a
					break
				}
			}

			currenteventnum := len(objs[activobjnum].events) - 1

			switch eventslist[num] {
			case "collision_obj", "collision_boundary":

				if eventslist[num] == "collision_obj" {
					objs[activobjnum].events[currenteventnum].event_type = 1
				}

				txtlen := rl.MeasureText("on collision with another obj:", txts)

				rl.DrawText("on collision with another obj:", (menu.rec.ToInt32().X + (menu.rec.ToInt32().Width/2 - (txtlen / 2))), y, txts, rl.White)

				y += txts * 2
				x = menu.rec.ToInt32().X + 5

				objeventfieldnames := reflect.TypeOf(objs[activobjnum].events[currenteventnum])
				fieldnames2 := make([]string, objeventfieldnames.NumField())
				fieldnames := make([]string, 0)

				for i := range fieldnames2 {
					if objeventfieldnames.Field(i).Name == "event_type" {
						break
					}
					fieldnames = append(fieldnames, objeventfieldnames.Field(i).Name)
				}

				height := float32((txts + (txts / 2)) * int32(len(fieldnames)))
				height += float32(txts * 2)

				infoboxrec := rl.NewRectangle(float32(x), float32(y), menu.rec.Width-10, height)
				rl.DrawRectangleRec(infoboxrec, rl.Fade(rl.Black, 0.7))

				//save event rec
				saveeventrec := rl.NewRectangle(infoboxrec.X, infoboxrec.Y+infoboxrec.Height+float32(txts), menu.rec.Width-10, float32(txts*2))

				if rl.CheckCollisionPointRec(mousev2, saveeventrec) {
					rl.DrawRectangleRec(saveeventrec, brightred())
					rl.DrawRectangleLinesEx(saveeventrec, 1.0, rl.White)
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						saveevent(currenteventnum)
					}

				} else {
					rl.DrawRectangleRec(saveeventrec, rl.Black)
				}
				txtlen = rl.MeasureText("save event", txts)
				rl.DrawText("save event", int32(menu.rec.X+menu.rec.Width/2)-txtlen/2, saveeventrec.ToInt32().Y+(txts/2), txts, rl.White)

				x += txts
				y += txts

				for _, txt := range fieldnames {
					mouseselrec := rl.NewRectangle(infoboxrec.X, float32(y-2), infoboxrec.Width, float32(txts+4))
					if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
						rl.DrawRectangleRec(mouseselrec, rl.Fade(brightred(), 0.2))
					}
					rl.DrawText(txt, x, y, txts, rl.White)

					switch txt {
					case "bounce":
						rec = toggle("", int32(menu.rec.X)+menu.rec.ToInt32().Width-30, y, objs[activobjnum].events[currenteventnum].bounce)
						if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].events[currenteventnum].bounce {
									objs[activobjnum].events[currenteventnum].bounce = false
								} else {
									objs[activobjnum].events[currenteventnum].bounce = true
									objs[activobjnum].events[currenteventnum].bounce_random = false
									objs[activobjnum].events[currenteventnum].stop_moving = false
								}
							}
						}
					case "bounce_random":
						rec = toggle("", int32(menu.rec.X)+menu.rec.ToInt32().Width-30, y, objs[activobjnum].events[currenteventnum].bounce_random)
						if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].events[currenteventnum].bounce_random {
									objs[activobjnum].events[currenteventnum].bounce_random = false
								} else {
									objs[activobjnum].events[currenteventnum].bounce_random = true
									objs[activobjnum].events[currenteventnum].bounce = false
									objs[activobjnum].events[currenteventnum].stop_moving = false
								}
							}
						}
					case "stop_moving":
						rec = toggle("", int32(menu.rec.X)+menu.rec.ToInt32().Width-30, y, objs[activobjnum].events[currenteventnum].stop_moving)
						if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].events[currenteventnum].stop_moving {
									objs[activobjnum].events[currenteventnum].stop_moving = false
								} else {
									objs[activobjnum].events[currenteventnum].stop_moving = true
									objs[activobjnum].events[currenteventnum].bounce_random = false
									objs[activobjnum].events[currenteventnum].bounce = false
								}
							}
						}
					case "explode":
						rec = toggle("", int32(menu.rec.X)+menu.rec.ToInt32().Width-30, y, objs[activobjnum].events[currenteventnum].explode)
						if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].events[currenteventnum].explode {
									objs[activobjnum].events[currenteventnum].explode = false
								} else {
									objs[activobjnum].events[currenteventnum].explode = true
								}
							}
						}
					case "destroy":
						rec = toggle("", int32(menu.rec.X)+menu.rec.ToInt32().Width-30, y, objs[activobjnum].events[currenteventnum].destroy)
						if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].events[currenteventnum].destroy {
									objs[activobjnum].events[currenteventnum].destroy = false
								} else {
									objs[activobjnum].events[currenteventnum].destroy = true
									objs[activobjnum].events[currenteventnum].explode = false
									objs[activobjnum].events[currenteventnum].bounce_random = false
									objs[activobjnum].events[currenteventnum].bounce = false
									objs[activobjnum].events[currenteventnum].stop_moving = false
									objs[activobjnum].events[currenteventnum].invisible = false
									objs[activobjnum].events[currenteventnum].visible = false
								}
							}
						}
					case "visible":
						rec = toggle("", int32(menu.rec.X)+menu.rec.ToInt32().Width-30, y, objs[activobjnum].events[currenteventnum].visible)
						if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].events[currenteventnum].visible {
									objs[activobjnum].events[currenteventnum].visible = false
								} else {
									objs[activobjnum].events[currenteventnum].visible = true
									objs[activobjnum].events[currenteventnum].invisible = false
								}
							}
						}
					case "invisible":
						rec = toggle("", int32(menu.rec.X)+menu.rec.ToInt32().Width-30, y, objs[activobjnum].events[currenteventnum].invisible)
						if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].events[currenteventnum].invisible {
									objs[activobjnum].events[currenteventnum].invisible = false
								} else {
									objs[activobjnum].events[currenteventnum].invisible = true
									objs[activobjnum].events[currenteventnum].visible = false
								}
							}
						}

					}

					y += txts
					y += txts / 2
				}

			}

		}
	} else if objmenuon && addobjpath && !addlabelon {

		y += txts * 2
		x = menu.rec.ToInt32().X + 5

		//new path rec
		newpathrec := rl.NewRectangle(float32(x), float32(y), menu.rec.Width-10, float32(txts*2))
		if addevent || rl.CheckCollisionPointRec(mousev2, newpathrec) {
			rl.DrawRectangleRec(newpathrec, brightred())
			rl.DrawRectangleLinesEx(newpathrec, 1.0, rl.White)

		} else {
			rl.DrawRectangleRec(newpathrec, rl.Black)
		}

		if rl.CheckCollisionPointRec(mousev2, newpathrec) {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

				if createpathon {
					createpathon = false
				} else {
					createpathon = true
					objs[activobjnum].path_move = true
					if len(objs[activobjnum].path) == 0 {
						objs[activobjnum].path = append(objs[activobjnum].path, objs[activobjnum].circv2)
					}
				}

			}
		}

		if createpathon {
			txtlen := rl.MeasureText("end path", txts)
			rl.DrawText("end path", int32(menu.rec.X+menu.rec.Width/2)-txtlen/2, y+txts/2, txts, rl.White)
		} else {
			txtlen := rl.MeasureText("create path", txts)
			rl.DrawText("create path", int32(menu.rec.X+menu.rec.Width/2)-txtlen/2, y+txts/2, txts, rl.White)
		}

		//delete path rec
		if len(objs[activobjnum].path) > 0 {
			y += txts * 2
			x = menu.rec.ToInt32().X + 5

			newpathrec = rl.NewRectangle(float32(x), float32(y), menu.rec.Width-10, float32(txts*2))
			if addevent || rl.CheckCollisionPointRec(mousev2, newpathrec) {
				rl.DrawRectangleRec(newpathrec, brightred())
				rl.DrawRectangleLinesEx(newpathrec, 1.0, rl.White)

			} else {
				rl.DrawRectangleRec(newpathrec, rl.Black)
			}

			if rl.CheckCollisionPointRec(mousev2, newpathrec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

				}
			}

			txtlen := rl.MeasureText("delete path", txts)
			rl.DrawText("delete path", int32(menu.rec.X+menu.rec.Width/2)-txtlen/2, y+txts/2, txts, rl.White)

		}

	} else if objmenuon && addlabelon {

		y += txts * 2
		x = menu.rec.ToInt32().X + 5

		if addinglabel {

			gettxton = true
			gettxtcharlimit = 20
			gettxtname = "obj_label"
			gettxttype = 3

		} else {

			//new label rec
			newlabelrec := rl.NewRectangle(float32(x), float32(y), menu.rec.Width-10, float32(txts*2))
			if addevent || rl.CheckCollisionPointRec(mousev2, newlabelrec) {
				rl.DrawRectangleRec(newlabelrec, brightred())
				rl.DrawRectangleLinesEx(newlabelrec, 1.0, rl.White)

			} else {
				rl.DrawRectangleRec(newlabelrec, rl.Black)
			}

			if rl.CheckCollisionPointRec(mousev2, newlabelrec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					addinglabel = true
				}
			}

			txtlen := rl.MeasureText("new label +", txts)
			rl.DrawText("new label +", int32(menu.rec.X+menu.rec.Width/2)-txtlen/2, y+txts/2, txts, rl.White)

			y += newlabelrec.ToInt32().Height + txts

			//new label rec
			existinglabelrec := rl.NewRectangle(float32(x), float32(y), menu.rec.Width-10, float32(txts*2))
			if addevent || rl.CheckCollisionPointRec(mousev2, existinglabelrec) {
				rl.DrawRectangleRec(existinglabelrec, brightred())
				rl.DrawRectangleLinesEx(existinglabelrec, 1.0, rl.White)

			} else {
				rl.DrawRectangleRec(existinglabelrec, rl.Black)
			}

			if rl.CheckCollisionPointRec(mousev2, existinglabelrec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					addexistinglabelon = true

				}
			}

			txtlen = rl.MeasureText("existing label +", txts)
			rl.DrawText("existing label +", int32(menu.rec.X+menu.rec.Width/2)-txtlen/2, y+txts/2, txts, rl.White)

			y += newlabelrec.ToInt32().Height + txts

			//new label rec
			newlabeleventrec := rl.NewRectangle(float32(x), float32(y), menu.rec.Width-10, float32(txts*2))
			if addevent || rl.CheckCollisionPointRec(mousev2, newlabeleventrec) {
				rl.DrawRectangleRec(newlabeleventrec, brightred())
				rl.DrawRectangleLinesEx(newlabeleventrec, 1.0, rl.White)

			} else {
				rl.DrawRectangleRec(newlabeleventrec, rl.Black)
			}

			if rl.CheckCollisionPointRec(mousev2, newlabeleventrec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					labeleventon = true
				}
			}

			txtlen = rl.MeasureText("label event +", txts)
			rl.DrawText("label event +", int32(menu.rec.X+menu.rec.Width/2)-txtlen/2, y+txts/2, txts, rl.White)

			y += newlabelrec.ToInt32().Height + txts

			//new label rec
			labelsallrec := rl.NewRectangle(float32(x), float32(y), menu.rec.Width-10, float32(txts*2))

			if rl.CheckCollisionPointRec(mousev2, labelsallrec) {
				rl.DrawRectangleRec(labelsallrec, brightred())
				rl.DrawRectangleLinesEx(labelsallrec, 1.0, rl.White)

			} else {
				rl.DrawRectangleRec(labelsallrec, rl.Black)
			}

			if rl.CheckCollisionPointRec(mousev2, labelsallrec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					viewlabelsliston = true
				}
			}

			txtlen = rl.MeasureText("all labels list", txts)
			rl.DrawText("all labels list", int32(menu.rec.X+menu.rec.Width/2)-txtlen/2, y+txts/2, txts, rl.White)

		}
		// list activ obj labels
		if len(objs[activobjnum].labels) > 0 {

			y += txts * 3
			txtlen := rl.MeasureText("obj labels", txts)
			rl.DrawText("obj labels", int32(menu.rec.X+menu.rec.Width/2)-txtlen/2, y+txts/2, txts, rl.White)
			y += txts * 2

			height := float32((txts + (txts / 2)) * int32(len(objs[activobjnum].labels)))
			height += float32(txts * 2)

			infoboxrec := rl.NewRectangle(float32(x), float32(y), menu.rec.Width-10, height)
			rl.DrawRectangleRec(infoboxrec, rl.Fade(rl.Black, 0.7))

			x += txts
			y += txts

			for i, labelname := range objs[activobjnum].labels {
				mouseselrec := rl.NewRectangle(infoboxrec.X, float32(y-2), infoboxrec.Width, float32(txts+4))
				if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
					rl.DrawRectangleRec(mouseselrec, rl.Fade(brightred(), 0.2))
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						editlabelnum = i
						editlabelactivobjon = true
					}
				}
				rl.DrawText(labelname, x, y, txts, rl.White)

				y += txts
				y += txts / 2

			}

		}
		x = menu.rec.ToInt32().X + 5
		// list activ obj label events
		if len(objs[activobjnum].labelevents) > 0 {

			y += txts * 3
			txtlen := rl.MeasureText("obj label events", txts)
			rl.DrawText("obj label events", int32(menu.rec.X+menu.rec.Width/2)-txtlen/2, y+txts/2, txts, rl.White)
			y += txts * 2

			height := float32((txts + (txts / 2)) * int32(len(objs[activobjnum].labels)))
			height += float32(txts * 2)

			infoboxrec := rl.NewRectangle(float32(x), float32(y), menu.rec.Width-10, height)
			rl.DrawRectangleRec(infoboxrec, rl.Fade(rl.Black, 0.7))

			x += txts
			y += txts

			for _, checklabelevent := range objs[activobjnum].labelevents {
				mouseselrec := rl.NewRectangle(infoboxrec.X, float32(y-2), infoboxrec.Width, float32(txts+4))
				if rl.CheckCollisionPointRec(mousev2, mouseselrec) {
					rl.DrawRectangleRec(mouseselrec, rl.Fade(brightred(), 0.2))
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

					}
				}
				rl.DrawText(checklabelevent.name, x, y, txts, rl.White)

				y += txts
				y += txts / 2

			}

		}

	}

}
func menugeneric(name string) { //MARK: menugeneric
	menufocus = true
	gettxtreturn = ""
	//back rec
	rl.DrawRectangle(0, 0, scrw, scrh, rl.Fade(rl.Black, 0.6))
	recwin := rl.NewRectangle(cntrscr.X-400, cntrscr.Y-200, 800, 400)
	//resize window for diffent menus
	if name == "labelevent" {
		recwin.Y -= 150
		recwin.Height += 300
	}
	rl.DrawRectangleRec(recwin, rl.Fade(darkred(), 0.7))

	//close win
	v2 := rl.NewVector2(recwin.X+recwin.Width-(closewinimg.Width*2), recwin.Y+closewinimg.Height)
	rec2 := rl.NewRectangle(v2.X, v2.Y, closewinimg.Width, closewinimg.Height)
	if rl.CheckCollisionPointRec(mousev2, rec2) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			editlabelactivobjon = false
			editlabelallon = false
			labeleventon = false
			menufocus = false
		}
		rl.DrawTextureRec(imgs, closewinimg, v2, brightred())
	} else {
		rl.DrawTextureRec(imgs, closewinimg, v2, rl.White)
	}

	switch name {

	case "labelevent":
		txthere("current obj label event", true, txtm, recwin.X, recwin.Y+20, recwin.Width)

		x := int32(recwin.X + 50)
		y := int32(recwin.Y + 80)

		txtlen := rl.MeasureText("if current obj", txts)
		rl.DrawText("if current obj", x, y, txts, rl.White)
		x += txtlen + txts

		//text back rec
		backrec := rl.NewRectangle(float32(x), float32(y-(txts/2)), 100, float32(txts*2))
		rl.DrawRectangleRec(backrec, rl.Black)
		//list down arrow
		v2 := rl.NewVector2((backrec.X+backrec.Width)-(darrowimg.Width+4), backrec.Y+float32(txts/2))
		downlistrec := rl.NewRectangle(v2.X-4, backrec.Y, darrowimg.Width+8, backrec.Height)
		if rl.CheckCollisionPointRec(mousev2, downlistrec) {
			rl.DrawRectangleRec(downlistrec, brightred())
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				labeleventnum++
				if labeleventnum > len(labeleventslist)-1 {
					labeleventnum = 0
				}
			}
		} else {
			rl.DrawRectangleRec(downlistrec, brightorange())
		}
		rl.DrawTextureRec(imgs, darrowimg, v2, rl.White)

		//label events list txt
		x += 5
		txt := labeleventslist[labeleventnum]
		rl.DrawText(txt, x, y, txts, rl.White)

		x = int32(backrec.X+backrec.Width) + txts

		txtlen = rl.MeasureText("objs with label", txts)
		rl.DrawText("objs with label", x, y, txts, rl.White)
		x += txtlen + txts

		//text back rec
		backrec = rl.NewRectangle(float32(x), float32(y-(txts/2)), 100, float32(txts*2))
		rl.DrawRectangleRec(backrec, rl.Black)
		//list down arrow
		v2 = rl.NewVector2((backrec.X+backrec.Width)-(darrowimg.Width+4), backrec.Y+float32(txts/2))
		downlistrec = rl.NewRectangle(v2.X-4, backrec.Y, darrowimg.Width+8, backrec.Height)
		if rl.CheckCollisionPointRec(mousev2, downlistrec) {
			rl.DrawRectangleRec(downlistrec, brightred())
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				labelsalllistnum++
				if labelsalllistnum > len(labelsalllist)-1 {
					labelsalllistnum = 0
				}
			}
		} else {
			rl.DrawRectangleRec(downlistrec, brightorange())
		}
		rl.DrawTextureRec(imgs, darrowimg, v2, rl.White)

		//labels all list txt
		x += 5
		txt = labelsalllist[labelsalllistnum]
		rl.DrawText(txt, x, y, txts, rl.White)

		x = int32(backrec.X+backrec.Width) + txts
		txtlen = rl.MeasureText("then current obj", txts)
		rl.DrawText("then current obj", x, y, txts, rl.White)
		x += txtlen + txts

		//text back rec
		backrec = rl.NewRectangle(float32(x), float32(y-(txts/2)), 130, float32(txts*2))
		rl.DrawRectangleRec(backrec, rl.Black)
		//list down arrow
		v2 = rl.NewVector2((backrec.X+backrec.Width)-(darrowimg.Width+4), backrec.Y+float32(txts/2))
		downlistrec = rl.NewRectangle(v2.X-4, backrec.Y, darrowimg.Width+8, backrec.Height)
		if rl.CheckCollisionPointRec(mousev2, downlistrec) {
			rl.DrawRectangleRec(downlistrec, brightred())
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				labelactionnum++
				if labelactionnum > len(labelactionslist)-1 {
					labelactionnum = 0
				}
			}
		} else {
			rl.DrawRectangleRec(downlistrec, brightorange())
		}
		rl.DrawTextureRec(imgs, darrowimg, v2, rl.White)

		//MARK: label events
		//label action list txt
		x += 5
		txt = labelactionslist[labelactionnum]
		rl.DrawText(txt, x, y, txts, rl.White)

		//obj event input
		if labelactionslist[labelactionnum] != "does nothing" {
			x = int32(recwin.X + 50)
			y += txts * 3

			switch labelactionslist[labelactionnum] {
			case "changes hp":
				txtlen = rl.MeasureText("change amount + or -", txts)
				rl.DrawText("change amount + or -", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt := fmt.Sprint(objs[activobjnum].labelchangehp)
				rec := buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 3
						gettxtname = "labelevent_labelchangehp"
						gettxttype = 5
					}
				}

				x += rec.ToInt32().Width + txts

				buttontxt = "="
				if objs[activobjnum].labelchangehpplus {
					buttontxt = "+ ="
				}
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						if objs[activobjnum].labelchangehpplus {
							objs[activobjnum].labelchangehpplus = false
						} else {
							objs[activobjnum].labelchangehpplus = true
						}
					}
				}

				x += rec.ToInt32().Width + txts

			case "changes position":

				txtlen = rl.MeasureText("new x", txts)
				rl.DrawText("new x", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt := fmt.Sprint(objs[activobjnum].labelchangeposx)
				rec := buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 3
						gettxtname = "labelevent_labelchangeposx"
						gettxttype = 5
					}
				}

				x += rec.ToInt32().Width + txts

				buttontxt = "="
				if objs[activobjnum].labelchangeposxplus {
					buttontxt = "+ ="
				}
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						if objs[activobjnum].labelchangeposxplus {
							objs[activobjnum].labelchangeposxplus = false
						} else {
							objs[activobjnum].labelchangeposxplus = true
						}
					}
				}

				x += rec.ToInt32().Width + txts

				txtlen = rl.MeasureText("new y", txts)
				rl.DrawText("new y", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt = fmt.Sprint(objs[activobjnum].labelchangeposy)
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 3
						gettxtname = "labelevent_labelchangeposy"
						gettxttype = 5
					}
				}

				x += rec.ToInt32().Width + txts

				buttontxt = "="
				if objs[activobjnum].labelchangeposyplus {
					buttontxt = "+ ="
				}
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						if objs[activobjnum].labelchangeposyplus {
							objs[activobjnum].labelchangeposyplus = false
						} else {
							objs[activobjnum].labelchangeposyplus = true
						}
					}
				}

				x += rec.ToInt32().Width + txts

			case "changes layer":

				txtlen = rl.MeasureText("moves to layer ?", txts)
				rl.DrawText("moves to layer ?", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt := "middleground"
				if objs[activobjnum].labelchangelayer == 1 {
					buttontxt = "foreground"
				} else if objs[activobjnum].labelchangelayer == -1 {
					buttontxt = "background"
				}

				rec := buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						objs[activobjnum].labelchangelayer++
						if objs[activobjnum].labelchangelayer > 1 {
							objs[activobjnum].labelchangelayer = -1
						}
					}
				}
				x += rec.ToInt32().Width + txts

				switch objs[activobjnum].labelchangelayer {
				case 0:
					if objs[activobjnum].middleground_obj {

						rl.DrawText("* current layer", x, y, txts, rl.White)
						txtlen = rl.MeasureText("* current layer", txts)
						x += txtlen + (txts * 2)
					}
				case -1:
					if objs[activobjnum].background_obj {

						rl.DrawText("* current layer", x, y, txts, rl.White)
						txtlen = rl.MeasureText("* current layer", txts)
						x += txtlen + (txts * 2)
					}
				case 1:
					if objs[activobjnum].foreground_obj {

						rl.DrawText("* current layer", x, y, txts, rl.White)
						txtlen = rl.MeasureText("* current layer", txts)
						x += txtlen + (txts * 2)
					}
				}

			case "changes shadow":

				txtlen = rl.MeasureText("shadow", txts)
				rl.DrawText("shadow", x, y, txts, rl.White)
				x += txtlen + txts

				onofftxt := "off"
				if objs[activobjnum].labelchangeshadow {
					onofftxt = "on"
				}
				rec := buttonorange(onofftxt, x, y-(txts/2))

				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						if objs[activobjnum].labelchangeshadow {
							objs[activobjnum].labelchangeshadow = false
						} else {
							objs[activobjnum].labelchangeshadow = true
						}
					}
				}

				x += rec.ToInt32().Width + txts

				txtlen = rl.MeasureText("x", txts)
				rl.DrawText("x", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt := fmt.Sprint(objs[activobjnum].labelchangeshadowx)
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 3
						gettxtname = "labelevent_changeshadowx"
						gettxttype = 5
					}
				}

				x += rec.ToInt32().Width + txts

				txtlen = rl.MeasureText("y", txts)
				rl.DrawText("y", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt = fmt.Sprint(objs[activobjnum].labelchangeshadowy)
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 3
						gettxtname = "labelevent_changeshadowy"
						gettxttype = 5
					}
				}

				x += rec.ToInt32().Width + txts

				txtlen = rl.MeasureText("fade", txts)
				rl.DrawText("fade", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt = fmt.Sprint(objs[activobjnum].labelchangeshadowfade)
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 1
						gettxtname = "labelevent_changeshadowfade"
						gettxttype = 2

					}
				}

				x += rec.ToInt32().Width + txts

				txtlen = rl.MeasureText("change color", txts)
				rl.DrawText("change color", x, y, txts, rl.White)
				x += txtlen + txts

				rec = rl.NewRectangle(float32(x), float32(y-5), 50, float32(txts+10))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					rl.DrawRectangleRec(rec, brightred())
					rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						changecolorsel = 11
						colorpalon = true
					}
				} else {
					rl.DrawRectangleRec(rec, objs[activobjnum].labelchangeshadowcolor)
				}

				x += rec.ToInt32().Width + txts

			case "changes ghosting":

				txtlen = rl.MeasureText("ghosting", txts)
				rl.DrawText("ghosting", x, y, txts, rl.White)
				x += txtlen + txts

				onofftxt := "off"
				if objs[activobjnum].labelchangeghosting {
					onofftxt = "on"
				}
				rec := buttonorange(onofftxt, x, y-(txts/2))

				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						if objs[activobjnum].labelchangeghosting {
							objs[activobjnum].labelchangeghosting = false
						} else {
							objs[activobjnum].labelchangeghosting = true
						}
					}
				}

				x += rec.ToInt32().Width + txts

				txtlen = rl.MeasureText("x", txts)
				rl.DrawText("x", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt := fmt.Sprint(objs[activobjnum].labelchangeghostx)
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 3
						gettxtname = "labelevent_changeghostx"
						gettxttype = 5
					}
				}

				x += rec.ToInt32().Width + txts

				txtlen = rl.MeasureText("y", txts)
				rl.DrawText("y", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt = fmt.Sprint(objs[activobjnum].labelchangeghosty)
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 3
						gettxtname = "labelevent_changeghosty"
						gettxttype = 5
					}
				}

				x += rec.ToInt32().Width + txts

				txtlen = rl.MeasureText("fade", txts)
				rl.DrawText("fade", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt = fmt.Sprint(objs[activobjnum].labelchangeghostfade)
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 1
						gettxtname = "labelevent_changeghostfade"
						gettxttype = 2

					}
				}

				x += rec.ToInt32().Width + txts

				txtlen = rl.MeasureText("change color", txts)
				rl.DrawText("change color", x, y, txts, rl.White)
				x += txtlen + txts

				rec = rl.NewRectangle(float32(x), float32(y-5), 50, float32(txts+10))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					rl.DrawRectangleRec(rec, brightred())
					rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						changecolorsel = 10
						colorpalon = true
					}
				} else {
					rl.DrawRectangleRec(rec, objs[activobjnum].labelchangeghostcolor)
				}

				x += rec.ToInt32().Width + txts

			case "changes fill":
				txtlen = rl.MeasureText("new fill color 1", txts)
				rl.DrawText("new fill color 1", x, y, txts, rl.White)
				x += txtlen + txts

				rec := rl.NewRectangle(float32(x), float32(y-5), 50, float32(txts+10))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					rl.DrawRectangleRec(rec, brightred())
					rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						changecolorsel = 8
						colorpalon = true
					}
				} else {
					rl.DrawRectangleRec(rec, objs[activobjnum].labelchangefill1)
				}

				x += rec.ToInt32().Width + txts
				txtlen = rl.MeasureText("gradient fill ?", txts)
				rl.DrawText("gradient fill ?", x, y, txts, rl.White)
				x += txtlen + txts

				objs[activobjnum].labelchangefillgradienton = tickbox(x, y, objs[activobjnum].labelchangefillgradienton)

				x += txts * 3

				if objs[activobjnum].labelchangefillgradienton {

					txtlen = rl.MeasureText("new fill color 2", txts)
					rl.DrawText("new fill color 2", x, y, txts, rl.White)
					x += txtlen + txts

					rec = rl.NewRectangle(float32(x), float32(y-5), 50, float32(txts+10))
					if rl.CheckCollisionPointRec(mousev2, rec) {
						rl.DrawRectangleRec(rec, brightred())
						rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							changecolorsel = 9
							colorpalon = true
						}
					} else {
						rl.DrawRectangleRec(rec, objs[activobjnum].labelchangefill2)
					}
					x += rec.ToInt32().Width + txts

					gradtxt := "horizontal gradient"
					if objs[activobjnum].labelchangefillgradienthv {
						gradtxt = "vertical gradient"
					}

					rec := buttonorange(gradtxt, x, y-(txts/2))

					if rl.CheckCollisionPointRec(mousev2, rec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							if objs[activobjnum].labelchangefillgradienthv {
								objs[activobjnum].labelchangefillgradienthv = false
							} else {
								objs[activobjnum].labelchangefillgradienthv = true
							}
						}
					}
					x += rec.ToInt32().Width + txts
				}
			case "changes image":

				txtlen = rl.MeasureText("new image", txts)
				rl.DrawText("new image", x, y, txts, rl.White)
				x += txtlen + txts

				rec := buttonorange("select image", x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if clickpause == 0 {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							labelchangeimgon = true
							clickpause = fps / 3
						}
					}
				}

				x += rec.ToInt32().Width + txts

			case "changes size":

				txtlen = rl.MeasureText("new width", txts)
				rl.DrawText("new width", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt := fmt.Sprint(objs[activobjnum].labelchangewidth)
				rec := buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 3
						gettxtname = "labelevent_changewidth"
						gettxttype = 5
					}
				}

				x += rec.ToInt32().Width + txts

				buttontxt = "="
				if objs[activobjnum].labelchangewidthplus {
					buttontxt = "+ ="
				}
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						if objs[activobjnum].labelchangewidthplus {
							objs[activobjnum].labelchangewidthplus = false
						} else {
							objs[activobjnum].labelchangewidthplus = true
						}
					}
				}

				x += rec.ToInt32().Width + txts

				txtlen = rl.MeasureText("new height", txts)
				rl.DrawText("new height", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt = fmt.Sprint(objs[activobjnum].labelchangeheight)
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 3
						gettxtname = "labelevent_changeheight"
						gettxttype = 5
					}
				}
				x += rec.ToInt32().Width + txts

				buttontxt = "="
				if objs[activobjnum].labelchangeheightplus {
					buttontxt = "+ ="
				}
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						if objs[activobjnum].labelchangeheightplus {
							objs[activobjnum].labelchangeheightplus = false
						} else {
							objs[activobjnum].labelchangeheightplus = true
						}
					}
				}

				x += rec.ToInt32().Width + txts

			case "changes rotation":
				txtlen = rl.MeasureText("new angle", txts)
				rl.DrawText("new angle", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt := fmt.Sprint(objs[activobjnum].labelchangerotation)
				rec := buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 3
						gettxtname = "labelevent_changerotation"
						gettxttype = 5
					}
				}
				x += rec.ToInt32().Width + txts

				buttontxt = "="
				if objs[activobjnum].labelchangerotationplus {
					buttontxt = "+ ="
				}
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						if objs[activobjnum].labelchangerotationplus {
							objs[activobjnum].labelchangerotationplus = false
						} else {
							objs[activobjnum].labelchangerotationplus = true
						}
					}
				}

				x += rec.ToInt32().Width + txts

			case "changes direction":

				txtlen = rl.MeasureText("new direction x", txts)
				rl.DrawText("new direction x", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt := fmt.Sprint(objs[activobjnum].labelchangedirecx)
				rec := buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 3
						gettxtname = "labelevent_changedirecx"
						gettxttype = 5
					}
				}

				x += rec.ToInt32().Width + txts

				buttontxt = "="
				if objs[activobjnum].labelchangedirecxplus {
					buttontxt = "+ ="
				}
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						if objs[activobjnum].labelchangedirecxplus {
							objs[activobjnum].labelchangedirecxplus = false
						} else {
							objs[activobjnum].labelchangedirecxplus = true
						}
					}
				}

				x += rec.ToInt32().Width + txts

				txtlen = rl.MeasureText("new direction y", txts)
				rl.DrawText("new direction y", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt = fmt.Sprint(objs[activobjnum].labelchangedirecy)
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 3
						gettxtname = "labelevent_changedirecy"
						gettxttype = 5
					}
				}

				x += rec.ToInt32().Width + txts

				buttontxt = "="
				if objs[activobjnum].labelchangedirecyplus {
					buttontxt = "+ ="
				}
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						if objs[activobjnum].labelchangedirecyplus {
							objs[activobjnum].labelchangedirecyplus = false
						} else {
							objs[activobjnum].labelchangedirecyplus = true
						}
					}
				}

				x += rec.ToInt32().Width + txts

			}

			// add activ obj event 2
			txtlen = rl.MeasureText("more current obj events ?", txts)
			rl.DrawText("more current obj events ?", x, y, txts, rl.White)
			x += txtlen + txts

			objs[activobjnum].labelevent2 = tickbox(x, y, objs[activobjnum].labelevent2)

			if objs[activobjnum].labelevent2 {
				x = int32(recwin.X + 50)
				y += txts * 3
				txtlen = rl.MeasureText("and current obj", txts)
				rl.DrawText("and current obj", x, y, txts, rl.White)
				x += txtlen + txts

				//text back rec
				backrec = rl.NewRectangle(float32(x), float32(y-(txts/2)), 130, float32(txts*2))
				rl.DrawRectangleRec(backrec, rl.Black)
				//list down arrow
				v2 = rl.NewVector2((backrec.X+backrec.Width)-(darrowimg.Width+4), backrec.Y+float32(txts/2))
				downlistrec = rl.NewRectangle(v2.X-4, backrec.Y, darrowimg.Width+8, backrec.Height)
				if rl.CheckCollisionPointRec(mousev2, downlistrec) {
					rl.DrawRectangleRec(downlistrec, brightred())
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						labelactionnum2++
						if labelactionnum2 > len(labelactionslist)-1 {
							labelactionnum2 = 0
						}
					}
				} else {
					rl.DrawRectangleRec(downlistrec, brightorange())
				}
				rl.DrawTextureRec(imgs, darrowimg, v2, rl.White)

				//label action 2 list txt
				x += 5
				txt = labelactionslist[labelactionnum2]
				rl.DrawText(txt, x, y, txts, rl.White)

				if labelactionslist[labelactionnum2] != "does nothing" {
					x = int32(recwin.X + 50)
					y += txts * 3

					switch labelactionslist[labelactionnum2] {
					case "changes hp":
						txtlen = rl.MeasureText("change amount + or -", txts)
						rl.DrawText("change amount + or -", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt := fmt.Sprint(objs[activobjnum].labelchangehp)
						rec := buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 3
								gettxtname = "labelevent_labelchangehp"
								gettxttype = 5
							}
						}

						x += rec.ToInt32().Width + txts

						buttontxt = "="
						if objs[activobjnum].labelchangehpplus {
							buttontxt = "+ ="
						}
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].labelchangehpplus {
									objs[activobjnum].labelchangehpplus = false
								} else {
									objs[activobjnum].labelchangehpplus = true
								}
							}
						}

						x += rec.ToInt32().Width + txts

					case "changes position":

						txtlen = rl.MeasureText("new x", txts)
						rl.DrawText("new x", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt := fmt.Sprint(objs[activobjnum].labelchangeposx)
						rec := buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 3
								gettxtname = "labelevent_labelchangeposx"
								gettxttype = 5
							}
						}

						x += rec.ToInt32().Width + txts

						buttontxt = "="
						if objs[activobjnum].labelchangeposxplus {
							buttontxt = "+ ="
						}
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].labelchangeposxplus {
									objs[activobjnum].labelchangeposxplus = false
								} else {
									objs[activobjnum].labelchangeposxplus = true
								}
							}
						}

						x += rec.ToInt32().Width + txts

						txtlen = rl.MeasureText("new y", txts)
						rl.DrawText("new y", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt = fmt.Sprint(objs[activobjnum].labelchangeposy)
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 3
								gettxtname = "labelevent_labelchangeposy"
								gettxttype = 5
							}
						}

						x += rec.ToInt32().Width + txts

						buttontxt = "="
						if objs[activobjnum].labelchangeposyplus {
							buttontxt = "+ ="
						}
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].labelchangeposyplus {
									objs[activobjnum].labelchangeposyplus = false
								} else {
									objs[activobjnum].labelchangeposyplus = true
								}
							}
						}

						x += rec.ToInt32().Width + txts

					case "changes layer":

						txtlen = rl.MeasureText("moves to layer ?", txts)
						rl.DrawText("moves to layer ?", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt := "middleground"
						if objs[activobjnum].labelchangelayer == 1 {
							buttontxt = "foreground"
						} else if objs[activobjnum].labelchangelayer == -1 {
							buttontxt = "background"
						}

						rec := buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								objs[activobjnum].labelchangelayer++
								if objs[activobjnum].labelchangelayer > 1 {
									objs[activobjnum].labelchangelayer = -1
								}
							}
						}
						x += rec.ToInt32().Width + txts

						switch objs[activobjnum].labelchangelayer {
						case 0:
							if objs[activobjnum].middleground_obj {

								rl.DrawText("* current layer", x, y, txts, rl.White)
								txtlen = rl.MeasureText("* current layer", txts)
								x += txtlen + (txts * 2)
							}
						case -1:
							if objs[activobjnum].background_obj {

								rl.DrawText("* current layer", x, y, txts, rl.White)
								txtlen = rl.MeasureText("* current layer", txts)
								x += txtlen + (txts * 2)
							}
						case 1:
							if objs[activobjnum].foreground_obj {

								rl.DrawText("* current layer", x, y, txts, rl.White)
								txtlen = rl.MeasureText("* current layer", txts)
								x += txtlen + (txts * 2)
							}
						}

					case "changes shadow":

						txtlen = rl.MeasureText("shadow", txts)
						rl.DrawText("shadow", x, y, txts, rl.White)
						x += txtlen + txts

						onofftxt := "off"
						if objs[activobjnum].labelchangeshadow {
							onofftxt = "on"
						}
						rec := buttonorange(onofftxt, x, y-(txts/2))

						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].labelchangeshadow {
									objs[activobjnum].labelchangeshadow = false
								} else {
									objs[activobjnum].labelchangeshadow = true
								}
							}
						}

						x += rec.ToInt32().Width + txts

						txtlen = rl.MeasureText("x", txts)
						rl.DrawText("x", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt := fmt.Sprint(objs[activobjnum].labelchangeshadowx)
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 3
								gettxtname = "labelevent_changeshadowx"
								gettxttype = 5
							}
						}

						x += rec.ToInt32().Width + txts

						txtlen = rl.MeasureText("y", txts)
						rl.DrawText("y", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt = fmt.Sprint(objs[activobjnum].labelchangeshadowy)
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 3
								gettxtname = "labelevent_changeshadowy"
								gettxttype = 5
							}
						}

						x += rec.ToInt32().Width + txts

						txtlen = rl.MeasureText("fade", txts)
						rl.DrawText("fade", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt = fmt.Sprint(objs[activobjnum].labelchangeshadowfade)
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 1
								gettxtname = "labelevent_changeshadowfade"
								gettxttype = 2

							}
						}

						x += rec.ToInt32().Width + txts

						txtlen = rl.MeasureText("change color", txts)
						rl.DrawText("change color", x, y, txts, rl.White)
						x += txtlen + txts

						rec = rl.NewRectangle(float32(x), float32(y-5), 50, float32(txts+10))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							rl.DrawRectangleRec(rec, brightred())
							rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								changecolorsel = 11
								colorpalon = true
							}
						} else {
							rl.DrawRectangleRec(rec, objs[activobjnum].labelchangeshadowcolor)
						}

						x += rec.ToInt32().Width + txts

					case "changes ghosting":

						txtlen = rl.MeasureText("ghosting", txts)
						rl.DrawText("ghosting", x, y, txts, rl.White)
						x += txtlen + txts

						onofftxt := "off"
						if objs[activobjnum].labelchangeghosting {
							onofftxt = "on"
						}
						rec := buttonorange(onofftxt, x, y-(txts/2))

						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].labelchangeghosting {
									objs[activobjnum].labelchangeghosting = false
								} else {
									objs[activobjnum].labelchangeghosting = true
								}
							}
						}

						x += rec.ToInt32().Width + txts

						txtlen = rl.MeasureText("x", txts)
						rl.DrawText("x", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt := fmt.Sprint(objs[activobjnum].labelchangeghostx)
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 3
								gettxtname = "labelevent_changeghostx"
								gettxttype = 5
							}
						}

						x += rec.ToInt32().Width + txts

						txtlen = rl.MeasureText("y", txts)
						rl.DrawText("y", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt = fmt.Sprint(objs[activobjnum].labelchangeghosty)
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 3
								gettxtname = "labelevent_changeghosty"
								gettxttype = 5
							}
						}

						x += rec.ToInt32().Width + txts

						txtlen = rl.MeasureText("fade", txts)
						rl.DrawText("fade", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt = fmt.Sprint(objs[activobjnum].labelchangeghostfade)
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 1
								gettxtname = "labelevent_changeghostfade"
								gettxttype = 2

							}
						}

						x += rec.ToInt32().Width + txts

						txtlen = rl.MeasureText("change color", txts)
						rl.DrawText("change color", x, y, txts, rl.White)
						x += txtlen + txts

						rec = rl.NewRectangle(float32(x), float32(y-5), 50, float32(txts+10))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							rl.DrawRectangleRec(rec, brightred())
							rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								changecolorsel = 10
								colorpalon = true
							}
						} else {
							rl.DrawRectangleRec(rec, objs[activobjnum].labelchangeghostcolor)
						}

						x += rec.ToInt32().Width + txts

					case "changes fill":
						txtlen = rl.MeasureText("new fill color 1", txts)
						rl.DrawText("new fill color 1", x, y, txts, rl.White)
						x += txtlen + txts

						rec := rl.NewRectangle(float32(x), float32(y-5), 50, float32(txts+10))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							rl.DrawRectangleRec(rec, brightred())
							rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								changecolorsel = 8
								colorpalon = true
							}
						} else {
							rl.DrawRectangleRec(rec, objs[activobjnum].labelchangefill1)
						}

						x += rec.ToInt32().Width + txts
						txtlen = rl.MeasureText("gradient fill ?", txts)
						rl.DrawText("gradient fill ?", x, y, txts, rl.White)
						x += txtlen + txts

						objs[activobjnum].labelchangefillgradienton = tickbox(x, y, objs[activobjnum].labelchangefillgradienton)

						x += txts * 3

						if objs[activobjnum].labelchangefillgradienton {

							txtlen = rl.MeasureText("new fill color 2", txts)
							rl.DrawText("new fill color 2", x, y, txts, rl.White)
							x += txtlen + txts

							rec = rl.NewRectangle(float32(x), float32(y-5), 50, float32(txts+10))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								rl.DrawRectangleRec(rec, brightred())
								rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									changecolorsel = 9
									colorpalon = true
								}
							} else {
								rl.DrawRectangleRec(rec, objs[activobjnum].labelchangefill2)
							}
							x += rec.ToInt32().Width + txts

							gradtxt := "horizontal gradient"
							if objs[activobjnum].labelchangefillgradienthv {
								gradtxt = "vertical gradient"
							}

							rec := buttonorange(gradtxt, x, y-(txts/2))

							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									if objs[activobjnum].labelchangefillgradienthv {
										objs[activobjnum].labelchangefillgradienthv = false
									} else {
										objs[activobjnum].labelchangefillgradienthv = true
									}
								}
							}
							x += rec.ToInt32().Width + txts
						}
					case "changes image":

						txtlen = rl.MeasureText("new image", txts)
						rl.DrawText("new image", x, y, txts, rl.White)
						x += txtlen + txts

						rec := buttonorange("select image", x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if clickpause == 0 {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									labelchangeimgon = true
									clickpause = fps / 3
								}
							}
						}

						x += rec.ToInt32().Width + txts

					case "changes size":

						txtlen = rl.MeasureText("new width", txts)
						rl.DrawText("new width", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt := fmt.Sprint(objs[activobjnum].labelchangewidth)
						rec := buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 3
								gettxtname = "labelevent_changewidth"
								gettxttype = 5
							}
						}

						x += rec.ToInt32().Width + txts

						buttontxt = "="
						if objs[activobjnum].labelchangewidthplus {
							buttontxt = "+ ="
						}
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].labelchangewidthplus {
									objs[activobjnum].labelchangewidthplus = false
								} else {
									objs[activobjnum].labelchangewidthplus = true
								}
							}
						}

						x += rec.ToInt32().Width + txts

						txtlen = rl.MeasureText("new height", txts)
						rl.DrawText("new height", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt = fmt.Sprint(objs[activobjnum].labelchangeheight)
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 3
								gettxtname = "labelevent_changeheight"
								gettxttype = 5
							}
						}
						x += rec.ToInt32().Width + txts

						buttontxt = "="
						if objs[activobjnum].labelchangeheightplus {
							buttontxt = "+ ="
						}
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].labelchangeheightplus {
									objs[activobjnum].labelchangeheightplus = false
								} else {
									objs[activobjnum].labelchangeheightplus = true
								}
							}
						}

						x += rec.ToInt32().Width + txts

					case "changes rotation":
						txtlen = rl.MeasureText("new angle", txts)
						rl.DrawText("new angle", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt := fmt.Sprint(objs[activobjnum].labelchangerotation)
						rec := buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 3
								gettxtname = "labelevent_changerotation"
								gettxttype = 5
							}
						}
						x += rec.ToInt32().Width + txts

						buttontxt = "="
						if objs[activobjnum].labelchangerotationplus {
							buttontxt = "+ ="
						}
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].labelchangerotationplus {
									objs[activobjnum].labelchangerotationplus = false
								} else {
									objs[activobjnum].labelchangerotationplus = true
								}
							}
						}

						x += rec.ToInt32().Width + txts

					case "changes direction":

						txtlen = rl.MeasureText("new direction x", txts)
						rl.DrawText("new direction x", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt := fmt.Sprint(objs[activobjnum].labelchangedirecx)
						rec := buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 3
								gettxtname = "labelevent_changedirecx"
								gettxttype = 5
							}
						}

						x += rec.ToInt32().Width + txts

						buttontxt = "="
						if objs[activobjnum].labelchangedirecxplus {
							buttontxt = "+ ="
						}
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].labelchangedirecxplus {
									objs[activobjnum].labelchangedirecxplus = false
								} else {
									objs[activobjnum].labelchangedirecxplus = true
								}
							}
						}

						x += rec.ToInt32().Width + txts

						txtlen = rl.MeasureText("new direction y", txts)
						rl.DrawText("new direction y", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt = fmt.Sprint(objs[activobjnum].labelchangedirecy)
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 3
								gettxtname = "labelevent_changedirecy"
								gettxttype = 5
							}
						}

						x += rec.ToInt32().Width + txts

						buttontxt = "="
						if objs[activobjnum].labelchangedirecyplus {
							buttontxt = "+ ="
						}
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].labelchangedirecyplus {
									objs[activobjnum].labelchangedirecyplus = false
								} else {
									objs[activobjnum].labelchangedirecyplus = true
								}
							}
						}

						x += rec.ToInt32().Width + txts

					}

					txtlen = rl.MeasureText("more current obj events ?", txts)
					rl.DrawText("more current obj events ?", x, y, txts, rl.White)
					x += txtlen + txts

					objs[activobjnum].labelevent3 = tickbox(x, y, objs[activobjnum].labelevent3)
				}
				// add activ obj event 3

				if objs[activobjnum].labelevent3 {
					x = int32(recwin.X + 50)
					y += txts * 3
					txtlen = rl.MeasureText("and current obj", txts)
					rl.DrawText("and current obj", x, y, txts, rl.White)
					x += txtlen + txts

					//text back rec
					backrec = rl.NewRectangle(float32(x), float32(y-(txts/2)), 130, float32(txts*2))
					rl.DrawRectangleRec(backrec, rl.Black)
					//list down arrow
					v2 = rl.NewVector2((backrec.X+backrec.Width)-(darrowimg.Width+4), backrec.Y+float32(txts/2))
					downlistrec = rl.NewRectangle(v2.X-4, backrec.Y, darrowimg.Width+8, backrec.Height)
					if rl.CheckCollisionPointRec(mousev2, downlistrec) {
						rl.DrawRectangleRec(downlistrec, brightred())
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							labelactionnum3++
							if labelactionnum3 > len(labelactionslist)-1 {
								labelactionnum3 = 0
							}
						}
					} else {
						rl.DrawRectangleRec(downlistrec, brightorange())
					}
					rl.DrawTextureRec(imgs, darrowimg, v2, rl.White)

					//label action 3 list txt
					x += 5
					txt = labelactionslist[labelactionnum3]
					rl.DrawText(txt, x, y, txts, rl.White)

					if labelactionslist[labelactionnum3] != "does nothing" {
						x = int32(recwin.X + 50)
						y += txts * 3

						switch labelactionslist[labelactionnum3] {
						case "changes hp":
							txtlen = rl.MeasureText("change amount + or -", txts)
							rl.DrawText("change amount + or -", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt := fmt.Sprint(objs[activobjnum].labelchangehp)
							rec := buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 3
									gettxtname = "labelevent_labelchangehp"
									gettxttype = 5
								}
							}

							x += rec.ToInt32().Width + txts

							buttontxt = "="
							if objs[activobjnum].labelchangehpplus {
								buttontxt = "+ ="
							}
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									if objs[activobjnum].labelchangehpplus {
										objs[activobjnum].labelchangehpplus = false
									} else {
										objs[activobjnum].labelchangehpplus = true
									}
								}
							}

							x += rec.ToInt32().Width + txts

						case "changes position":

							txtlen = rl.MeasureText("new x", txts)
							rl.DrawText("new x", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt := fmt.Sprint(objs[activobjnum].labelchangeposx)
							rec := buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 3
									gettxtname = "labelevent_labelchangeposx"
									gettxttype = 5
								}
							}

							x += rec.ToInt32().Width + txts

							buttontxt = "="
							if objs[activobjnum].labelchangeposxplus {
								buttontxt = "+ ="
							}
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									if objs[activobjnum].labelchangeposxplus {
										objs[activobjnum].labelchangeposxplus = false
									} else {
										objs[activobjnum].labelchangeposxplus = true
									}
								}
							}

							x += rec.ToInt32().Width + txts

							txtlen = rl.MeasureText("new y", txts)
							rl.DrawText("new y", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt = fmt.Sprint(objs[activobjnum].labelchangeposy)
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 3
									gettxtname = "labelevent_labelchangeposy"
									gettxttype = 5
								}
							}

							x += rec.ToInt32().Width + txts

							buttontxt = "="
							if objs[activobjnum].labelchangeposyplus {
								buttontxt = "+ ="
							}
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									if objs[activobjnum].labelchangeposyplus {
										objs[activobjnum].labelchangeposyplus = false
									} else {
										objs[activobjnum].labelchangeposyplus = true
									}
								}
							}

							x += rec.ToInt32().Width + txts

						case "changes layer":

							txtlen = rl.MeasureText("moves to layer ?", txts)
							rl.DrawText("moves to layer ?", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt := "middleground"
							if objs[activobjnum].labelchangelayer == 1 {
								buttontxt = "foreground"
							} else if objs[activobjnum].labelchangelayer == -1 {
								buttontxt = "background"
							}

							rec := buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									objs[activobjnum].labelchangelayer++
									if objs[activobjnum].labelchangelayer > 1 {
										objs[activobjnum].labelchangelayer = -1
									}
								}
							}
							x += rec.ToInt32().Width + txts

							switch objs[activobjnum].labelchangelayer {
							case 0:
								if objs[activobjnum].middleground_obj {

									rl.DrawText("* current layer", x, y, txts, rl.White)
									txtlen = rl.MeasureText("* current layer", txts)
									x += txtlen + (txts * 2)
								}
							case -1:
								if objs[activobjnum].background_obj {

									rl.DrawText("* current layer", x, y, txts, rl.White)
									txtlen = rl.MeasureText("* current layer", txts)
									x += txtlen + (txts * 2)
								}
							case 1:
								if objs[activobjnum].foreground_obj {

									rl.DrawText("* current layer", x, y, txts, rl.White)
									txtlen = rl.MeasureText("* current layer", txts)
									x += txtlen + (txts * 2)
								}
							}

						case "changes shadow":

							txtlen = rl.MeasureText("shadow", txts)
							rl.DrawText("shadow", x, y, txts, rl.White)
							x += txtlen + txts

							onofftxt := "off"
							if objs[activobjnum].labelchangeshadow {
								onofftxt = "on"
							}
							rec := buttonorange(onofftxt, x, y-(txts/2))

							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									if objs[activobjnum].labelchangeshadow {
										objs[activobjnum].labelchangeshadow = false
									} else {
										objs[activobjnum].labelchangeshadow = true
									}
								}
							}

							x += rec.ToInt32().Width + txts

							txtlen = rl.MeasureText("x", txts)
							rl.DrawText("x", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt := fmt.Sprint(objs[activobjnum].labelchangeshadowx)
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 3
									gettxtname = "labelevent_changeshadowx"
									gettxttype = 5
								}
							}

							x += rec.ToInt32().Width + txts

							txtlen = rl.MeasureText("y", txts)
							rl.DrawText("y", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt = fmt.Sprint(objs[activobjnum].labelchangeshadowy)
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 3
									gettxtname = "labelevent_changeshadowy"
									gettxttype = 5
								}
							}

							x += rec.ToInt32().Width + txts

							txtlen = rl.MeasureText("fade", txts)
							rl.DrawText("fade", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt = fmt.Sprint(objs[activobjnum].labelchangeshadowfade)
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 1
									gettxtname = "labelevent_changeshadowfade"
									gettxttype = 2

								}
							}

							x += rec.ToInt32().Width + txts

							txtlen = rl.MeasureText("change color", txts)
							rl.DrawText("change color", x, y, txts, rl.White)
							x += txtlen + txts

							rec = rl.NewRectangle(float32(x), float32(y-5), 50, float32(txts+10))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								rl.DrawRectangleRec(rec, brightred())
								rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									changecolorsel = 11
									colorpalon = true
								}
							} else {
								rl.DrawRectangleRec(rec, objs[activobjnum].labelchangeshadowcolor)
							}

							x += rec.ToInt32().Width + txts

						case "changes ghosting":

							txtlen = rl.MeasureText("ghosting", txts)
							rl.DrawText("ghosting", x, y, txts, rl.White)
							x += txtlen + txts

							onofftxt := "off"
							if objs[activobjnum].labelchangeghosting {
								onofftxt = "on"
							}
							rec := buttonorange(onofftxt, x, y-(txts/2))

							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									if objs[activobjnum].labelchangeghosting {
										objs[activobjnum].labelchangeghosting = false
									} else {
										objs[activobjnum].labelchangeghosting = true
									}
								}
							}

							x += rec.ToInt32().Width + txts

							txtlen = rl.MeasureText("x", txts)
							rl.DrawText("x", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt := fmt.Sprint(objs[activobjnum].labelchangeghostx)
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 3
									gettxtname = "labelevent_changeghostx"
									gettxttype = 5
								}
							}

							x += rec.ToInt32().Width + txts

							txtlen = rl.MeasureText("y", txts)
							rl.DrawText("y", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt = fmt.Sprint(objs[activobjnum].labelchangeghosty)
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 3
									gettxtname = "labelevent_changeghosty"
									gettxttype = 5
								}
							}

							x += rec.ToInt32().Width + txts

							txtlen = rl.MeasureText("fade", txts)
							rl.DrawText("fade", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt = fmt.Sprint(objs[activobjnum].labelchangeghostfade)
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 1
									gettxtname = "labelevent_changeghostfade"
									gettxttype = 2

								}
							}

							x += rec.ToInt32().Width + txts

							txtlen = rl.MeasureText("change color", txts)
							rl.DrawText("change color", x, y, txts, rl.White)
							x += txtlen + txts

							rec = rl.NewRectangle(float32(x), float32(y-5), 50, float32(txts+10))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								rl.DrawRectangleRec(rec, brightred())
								rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									changecolorsel = 10
									colorpalon = true
								}
							} else {
								rl.DrawRectangleRec(rec, objs[activobjnum].labelchangeghostcolor)
							}

							x += rec.ToInt32().Width + txts

						case "changes fill":
							txtlen = rl.MeasureText("new fill color 1", txts)
							rl.DrawText("new fill color 1", x, y, txts, rl.White)
							x += txtlen + txts

							rec := rl.NewRectangle(float32(x), float32(y-5), 50, float32(txts+10))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								rl.DrawRectangleRec(rec, brightred())
								rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									changecolorsel = 8
									colorpalon = true
								}
							} else {
								rl.DrawRectangleRec(rec, objs[activobjnum].labelchangefill1)
							}

							x += rec.ToInt32().Width + txts
							txtlen = rl.MeasureText("gradient fill ?", txts)
							rl.DrawText("gradient fill ?", x, y, txts, rl.White)
							x += txtlen + txts

							objs[activobjnum].labelchangefillgradienton = tickbox(x, y, objs[activobjnum].labelchangefillgradienton)

							x += txts * 3

							if objs[activobjnum].labelchangefillgradienton {

								txtlen = rl.MeasureText("new fill color 2", txts)
								rl.DrawText("new fill color 2", x, y, txts, rl.White)
								x += txtlen + txts

								rec = rl.NewRectangle(float32(x), float32(y-5), 50, float32(txts+10))
								if rl.CheckCollisionPointRec(mousev2, rec) {
									rl.DrawRectangleRec(rec, brightred())
									rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
									if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
										changecolorsel = 9
										colorpalon = true
									}
								} else {
									rl.DrawRectangleRec(rec, objs[activobjnum].labelchangefill2)
								}
								x += rec.ToInt32().Width + txts

								gradtxt := "horizontal gradient"
								if objs[activobjnum].labelchangefillgradienthv {
									gradtxt = "vertical gradient"
								}

								rec := buttonorange(gradtxt, x, y-(txts/2))

								if rl.CheckCollisionPointRec(mousev2, rec) {
									if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
										if objs[activobjnum].labelchangefillgradienthv {
											objs[activobjnum].labelchangefillgradienthv = false
										} else {
											objs[activobjnum].labelchangefillgradienthv = true
										}
									}
								}
								x += rec.ToInt32().Width + txts
							}
						case "changes image":

							txtlen = rl.MeasureText("new image", txts)
							rl.DrawText("new image", x, y, txts, rl.White)
							x += txtlen + txts

							rec := buttonorange("select image", x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if clickpause == 0 {
									if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
										labelchangeimgon = true
										clickpause = fps / 3
									}
								}
							}

							x += rec.ToInt32().Width + txts

						case "changes size":

							txtlen = rl.MeasureText("new width", txts)
							rl.DrawText("new width", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt := fmt.Sprint(objs[activobjnum].labelchangewidth)
							rec := buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 3
									gettxtname = "labelevent_changewidth"
									gettxttype = 5
								}
							}

							x += rec.ToInt32().Width + txts

							buttontxt = "="
							if objs[activobjnum].labelchangewidthplus {
								buttontxt = "+ ="
							}
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									if objs[activobjnum].labelchangewidthplus {
										objs[activobjnum].labelchangewidthplus = false
									} else {
										objs[activobjnum].labelchangewidthplus = true
									}
								}
							}

							x += rec.ToInt32().Width + txts

							txtlen = rl.MeasureText("new height", txts)
							rl.DrawText("new height", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt = fmt.Sprint(objs[activobjnum].labelchangeheight)
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 3
									gettxtname = "labelevent_changeheight"
									gettxttype = 5
								}
							}
							x += rec.ToInt32().Width + txts

							buttontxt = "="
							if objs[activobjnum].labelchangeheightplus {
								buttontxt = "+ ="
							}
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									if objs[activobjnum].labelchangeheightplus {
										objs[activobjnum].labelchangeheightplus = false
									} else {
										objs[activobjnum].labelchangeheightplus = true
									}
								}
							}

							x += rec.ToInt32().Width + txts

						case "changes rotation":
							txtlen = rl.MeasureText("new angle", txts)
							rl.DrawText("new angle", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt := fmt.Sprint(objs[activobjnum].labelchangerotation)
							rec := buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 3
									gettxtname = "labelevent_changerotation"
									gettxttype = 5
								}
							}
							x += rec.ToInt32().Width + txts

							buttontxt = "="
							if objs[activobjnum].labelchangerotationplus {
								buttontxt = "+ ="
							}
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									if objs[activobjnum].labelchangerotationplus {
										objs[activobjnum].labelchangerotationplus = false
									} else {
										objs[activobjnum].labelchangerotationplus = true
									}
								}
							}

							x += rec.ToInt32().Width + txts

						case "changes direction":

							txtlen = rl.MeasureText("new direction x", txts)
							rl.DrawText("new direction x", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt := fmt.Sprint(objs[activobjnum].labelchangedirecx)
							rec := buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 3
									gettxtname = "labelevent_changedirecx"
									gettxttype = 5
								}
							}

							x += rec.ToInt32().Width + txts

							buttontxt = "="
							if objs[activobjnum].labelchangedirecxplus {
								buttontxt = "+ ="
							}
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									if objs[activobjnum].labelchangedirecxplus {
										objs[activobjnum].labelchangedirecxplus = false
									} else {
										objs[activobjnum].labelchangedirecxplus = true
									}
								}
							}

							x += rec.ToInt32().Width + txts

							txtlen = rl.MeasureText("new direction y", txts)
							rl.DrawText("new direction y", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt = fmt.Sprint(objs[activobjnum].labelchangedirecy)
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 3
									gettxtname = "labelevent_changedirecy"
									gettxttype = 5
								}
							}

							x += rec.ToInt32().Width + txts

							buttontxt = "="
							if objs[activobjnum].labelchangedirecyplus {
								buttontxt = "+ ="
							}
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									if objs[activobjnum].labelchangedirecyplus {
										objs[activobjnum].labelchangedirecyplus = false
									} else {
										objs[activobjnum].labelchangedirecyplus = true
									}
								}
							}

							x += rec.ToInt32().Width + txts

						}
					}
				}
			}

		}
		x = int32(recwin.X + 50)
		y += txts * 3

		txtlen = rl.MeasureText("and obj 2", txts)
		rl.DrawText("and obj 2", x, y, txts, rl.White)
		x += txtlen + txts

		//text back rec
		backrec = rl.NewRectangle(float32(x), float32(y-(txts/2)), 130, float32(txts*2))
		rl.DrawRectangleRec(backrec, rl.Black)
		//list down arrow
		v2 = rl.NewVector2((backrec.X+backrec.Width)-(darrowimg.Width+4), backrec.Y+float32(txts/2))
		downlistrec = rl.NewRectangle(v2.X-4, backrec.Y, darrowimg.Width+8, backrec.Height)
		if rl.CheckCollisionPointRec(mousev2, downlistrec) {
			rl.DrawRectangleRec(downlistrec, brightred())
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				labelactionnumobj2++
				if labelactionnumobj2 > len(labelactionslistobj2)-1 {
					labelactionnumobj2 = 0
				}
			}
		} else {
			rl.DrawRectangleRec(downlistrec, brightorange())
		}
		rl.DrawTextureRec(imgs, darrowimg, v2, rl.White)

		//label obj2 action list txt
		x += 5
		txt = labelactionslistobj2[labelactionnumobj2]
		rl.DrawText(txt, x, y, txts, rl.White)

		x = int32(backrec.X+backrec.Width) + txts

		if labelactionslistobj2[labelactionnumobj2] != "does nothing" {
			x = int32(recwin.X + 50)
			y += txts * 3

			switch labelactionslistobj2[labelactionnumobj2] {
			case "changes hp":
				txtlen = rl.MeasureText("change amount + or -", txts)
				rl.DrawText("change amount + or -", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt := fmt.Sprint(objs[activobjnum].labelchangehpobj2)
				rec := buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 3
						gettxtname = "labelevent_labelchangehpobj2"
						gettxttype = 5
					}
				}

				x += rec.ToInt32().Width + txts

				buttontxt = "="
				if objs[activobjnum].labelchangehpplusobj2 {
					buttontxt = "+ ="
				}
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						if objs[activobjnum].labelchangehpplusobj2 {
							objs[activobjnum].labelchangehpplusobj2 = false
						} else {
							objs[activobjnum].labelchangehpplusobj2 = true
						}
					}
				}

				x += rec.ToInt32().Width + txts

			case "changes position":

				txtlen = rl.MeasureText("new x", txts)
				rl.DrawText("new x", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt := fmt.Sprint(objs[activobjnum].labelchangeposxobj2)
				rec := buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 3
						gettxtname = "labelevent_labelchangeposxobj2"
						gettxttype = 5
					}
				}

				x += rec.ToInt32().Width + txts

				buttontxt = "="
				if objs[activobjnum].labelchangeposxplusobj2 {
					buttontxt = "+ ="
				}
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						if objs[activobjnum].labelchangeposxplusobj2 {
							objs[activobjnum].labelchangeposxplusobj2 = false
						} else {
							objs[activobjnum].labelchangeposxplusobj2 = true
						}
					}
				}

				x += rec.ToInt32().Width + txts

				txtlen = rl.MeasureText("new y", txts)
				rl.DrawText("new y", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt = fmt.Sprint(objs[activobjnum].labelchangeposyobj2)
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 3
						gettxtname = "labelevent_labelchangeposyobj2"
						gettxttype = 5
					}
				}

				x += rec.ToInt32().Width + txts

				buttontxt = "="
				if objs[activobjnum].labelchangeposyplusobj2 {
					buttontxt = "+ ="
				}
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						if objs[activobjnum].labelchangeposyplusobj2 {
							objs[activobjnum].labelchangeposyplusobj2 = false
						} else {
							objs[activobjnum].labelchangeposyplusobj2 = true
						}
					}
				}

				x += rec.ToInt32().Width + txts

			case "changes layer":

				txtlen = rl.MeasureText("moves to layer ?", txts)
				rl.DrawText("moves to layer ?", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt := "middleground"
				if objs[activobjnum].labelchangelayerobj2 == 1 {
					buttontxt = "foreground"
				} else if objs[activobjnum].labelchangelayerobj2 == -1 {
					buttontxt = "background"
				}

				rec := buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						objs[activobjnum].labelchangelayerobj2++
						if objs[activobjnum].labelchangelayerobj2 > 1 {
							objs[activobjnum].labelchangelayerobj2 = -1
						}
					}
				}
				x += rec.ToInt32().Width + txts

				switch objs[activobjnum].labelchangelayerobj2 {
				case 0:
					if objs[activobjnum].middleground_obj {

						rl.DrawText("* obj 1 layer", x, y, txts, rl.White)
						txtlen = rl.MeasureText("* obj 1 layer", txts)
						x += txtlen + (txts * 2)
					}
				case -1:
					if objs[activobjnum].background_obj {

						rl.DrawText("* obj 1 layer", x, y, txts, rl.White)
						txtlen = rl.MeasureText("* obj 1 layer", txts)
						x += txtlen + (txts * 2)
					}
				case 1:
					if objs[activobjnum].foreground_obj {

						rl.DrawText("* obj 1 layer", x, y, txts, rl.White)
						txtlen = rl.MeasureText("* obj 1 layer", txts)
						x += txtlen + (txts * 2)
					}
				}

			case "changes shadow":

				txtlen = rl.MeasureText("shadow", txts)
				rl.DrawText("shadow", x, y, txts, rl.White)
				x += txtlen + txts

				onofftxt := "off"
				if objs[activobjnum].labelchangeshadowobj2 {
					onofftxt = "on"
				}
				rec := buttonorange(onofftxt, x, y-(txts/2))

				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						if objs[activobjnum].labelchangeshadowobj2 {
							objs[activobjnum].labelchangeshadowobj2 = false
						} else {
							objs[activobjnum].labelchangeshadowobj2 = true
						}
					}
				}

				x += rec.ToInt32().Width + txts

				txtlen = rl.MeasureText("x", txts)
				rl.DrawText("x", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt := fmt.Sprint(objs[activobjnum].labelchangeshadowxobj2)
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 3
						gettxtname = "labelevent_changeshadowxobj2"
						gettxttype = 5
					}
				}

				x += rec.ToInt32().Width + txts

				txtlen = rl.MeasureText("y", txts)
				rl.DrawText("y", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt = fmt.Sprint(objs[activobjnum].labelchangeshadowyobj2)
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 3
						gettxtname = "labelevent_changeshadowyobj2"
						gettxttype = 5
					}
				}

				x += rec.ToInt32().Width + txts

				txtlen = rl.MeasureText("fade", txts)
				rl.DrawText("fade", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt = fmt.Sprint(objs[activobjnum].labelchangeshadowfadeobj2)
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 1
						gettxtname = "labelevent_changeshadowfadeobj2"
						gettxttype = 2

					}
				}

				x += rec.ToInt32().Width + txts

				txtlen = rl.MeasureText("change color", txts)
				rl.DrawText("change color", x, y, txts, rl.White)
				x += txtlen + txts

				rec = rl.NewRectangle(float32(x), float32(y-5), 50, float32(txts+10))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					rl.DrawRectangleRec(rec, brightred())
					rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						changecolorsel = 15
						colorpalon = true
					}
				} else {
					rl.DrawRectangleRec(rec, objs[activobjnum].labelchangeshadowcolorobj2)
				}

				x += rec.ToInt32().Width + txts

			case "changes ghosting":

				txtlen = rl.MeasureText("ghosting", txts)
				rl.DrawText("ghosting", x, y, txts, rl.White)
				x += txtlen + txts

				onofftxt := "off"
				if objs[activobjnum].labelchangeghostingobj2 {
					onofftxt = "on"
				}
				rec := buttonorange(onofftxt, x, y-(txts/2))

				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						if objs[activobjnum].labelchangeghostingobj2 {
							objs[activobjnum].labelchangeghostingobj2 = false
						} else {
							objs[activobjnum].labelchangeghostingobj2 = true
						}
					}
				}

				x += rec.ToInt32().Width + txts

				txtlen = rl.MeasureText("x", txts)
				rl.DrawText("x", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt := fmt.Sprint(objs[activobjnum].labelchangeghostxobj2)
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 3
						gettxtname = "labelevent_changeghostxobj2"
						gettxttype = 5
					}
				}

				x += rec.ToInt32().Width + txts

				txtlen = rl.MeasureText("y", txts)
				rl.DrawText("y", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt = fmt.Sprint(objs[activobjnum].labelchangeghostyobj2)
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 3
						gettxtname = "labelevent_changeghostyobj2"
						gettxttype = 5
					}
				}

				x += rec.ToInt32().Width + txts

				txtlen = rl.MeasureText("fade", txts)
				rl.DrawText("fade", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt = fmt.Sprint(objs[activobjnum].labelchangeghostfadeobj2)
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 1
						gettxtname = "labelevent_changeghostfadeobj2"
						gettxttype = 2

					}
				}

				x += rec.ToInt32().Width + txts

				txtlen = rl.MeasureText("change color", txts)
				rl.DrawText("change color", x, y, txts, rl.White)
				x += txtlen + txts

				rec = rl.NewRectangle(float32(x), float32(y-5), 50, float32(txts+10))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					rl.DrawRectangleRec(rec, brightred())
					rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						changecolorsel = 14
						colorpalon = true
					}
				} else {
					rl.DrawRectangleRec(rec, objs[activobjnum].labelchangeghostcolorobj2)
				}

				x += rec.ToInt32().Width + txts

			case "changes fill":
				txtlen = rl.MeasureText("new fill color 1", txts)
				rl.DrawText("new fill color 1", x, y, txts, rl.White)
				x += txtlen + txts

				rec := rl.NewRectangle(float32(x), float32(y-5), 50, float32(txts+10))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					rl.DrawRectangleRec(rec, brightred())
					rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						changecolorsel = 12
						colorpalon = true
					}
				} else {
					rl.DrawRectangleRec(rec, objs[activobjnum].labelchangefill1obj2)
				}

				x += rec.ToInt32().Width + txts
				txtlen = rl.MeasureText("gradient fill ?", txts)
				rl.DrawText("gradient fill ?", x, y, txts, rl.White)
				x += txtlen + txts

				objs[activobjnum].labelchangefillgradientonobj2 = tickbox(x, y, objs[activobjnum].labelchangefillgradientonobj2)

				x += txts * 3

				if objs[activobjnum].labelchangefillgradientonobj2 {

					txtlen = rl.MeasureText("new fill color 2", txts)
					rl.DrawText("new fill color 2", x, y, txts, rl.White)
					x += txtlen + txts

					rec = rl.NewRectangle(float32(x), float32(y-5), 50, float32(txts+10))
					if rl.CheckCollisionPointRec(mousev2, rec) {
						rl.DrawRectangleRec(rec, brightred())
						rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							changecolorsel = 13
							colorpalon = true
						}
					} else {
						rl.DrawRectangleRec(rec, objs[activobjnum].labelchangefill2obj2)
					}
					x += rec.ToInt32().Width + txts

					gradtxt := "horizontal gradient"
					if objs[activobjnum].labelchangefillgradienthvobj2 {
						gradtxt = "vertical gradient"
					}

					rec := buttonorange(gradtxt, x, y-(txts/2))

					if rl.CheckCollisionPointRec(mousev2, rec) {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							if objs[activobjnum].labelchangefillgradienthvobj2 {
								objs[activobjnum].labelchangefillgradienthvobj2 = false
							} else {
								objs[activobjnum].labelchangefillgradienthvobj2 = true
							}
						}
					}
					x += rec.ToInt32().Width + txts
				}
			case "changes image":

				txtlen = rl.MeasureText("new image", txts)
				rl.DrawText("new image", x, y, txts, rl.White)
				x += txtlen + txts

				rec := buttonorange("select image", x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						labelchangeimgonobj2 = true
					}
				}

				x += rec.ToInt32().Width + txts

			case "changes size":

				txtlen = rl.MeasureText("new width", txts)
				rl.DrawText("new width", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt := fmt.Sprint(objs[activobjnum].labelchangewidthobj2)
				rec := buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 3
						gettxtname = "labelevent_changewidthobj2"
						gettxttype = 5
					}
				}

				x += rec.ToInt32().Width + txts

				buttontxt = "="
				if objs[activobjnum].labelchangewidthplusobj2 {
					buttontxt = "+ ="
				}
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						if objs[activobjnum].labelchangewidthplusobj2 {
							objs[activobjnum].labelchangewidthplusobj2 = false
						} else {
							objs[activobjnum].labelchangewidthplusobj2 = true
						}
					}
				}

				x += rec.ToInt32().Width + txts

				txtlen = rl.MeasureText("new height", txts)
				rl.DrawText("new height", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt = fmt.Sprint(objs[activobjnum].labelchangeheightobj2)
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 3
						gettxtname = "labelevent_changeheightobj2"
						gettxttype = 5
					}
				}
				x += rec.ToInt32().Width + txts

				buttontxt = "="
				if objs[activobjnum].labelchangeheightplusobj2 {
					buttontxt = "+ ="
				}
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						if objs[activobjnum].labelchangeheightplusobj2 {
							objs[activobjnum].labelchangeheightplusobj2 = false
						} else {
							objs[activobjnum].labelchangeheightplusobj2 = true
						}
					}
				}

				x += rec.ToInt32().Width + txts

			case "changes rotation":
				txtlen = rl.MeasureText("new angle", txts)
				rl.DrawText("new angle", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt := fmt.Sprint(objs[activobjnum].labelchangerotationobj2)
				rec := buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 3
						gettxtname = "labelevent_changerotationobj2"
						gettxttype = 5
					}
				}
				x += rec.ToInt32().Width + txts

				buttontxt = "="
				if objs[activobjnum].labelchangerotationplusobj2 {
					buttontxt = "+ ="
				}
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						if objs[activobjnum].labelchangerotationplusobj2 {
							objs[activobjnum].labelchangerotationplusobj2 = false
						} else {
							objs[activobjnum].labelchangerotationplusobj2 = true
						}
					}
				}

				x += rec.ToInt32().Width + txts

			case "changes direction":

				txtlen = rl.MeasureText("new direction x", txts)
				rl.DrawText("new direction x", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt := fmt.Sprint(objs[activobjnum].labelchangedirecxobj2)
				rec := buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 3
						gettxtname = "labelevent_changedirecxobj2"
						gettxttype = 5
					}
				}

				x += rec.ToInt32().Width + txts

				buttontxt = "="
				if objs[activobjnum].labelchangedirecxplusobj2 {
					buttontxt = "+ ="
				}
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						if objs[activobjnum].labelchangedirecxplusobj2 {
							objs[activobjnum].labelchangedirecxplusobj2 = false
						} else {
							objs[activobjnum].labelchangedirecxplusobj2 = true
						}
					}
				}

				x += rec.ToInt32().Width + txts

				txtlen = rl.MeasureText("new direction y", txts)
				rl.DrawText("new direction y", x, y, txts, rl.White)
				x += txtlen + txts

				buttontxt = fmt.Sprint(objs[activobjnum].labelchangedirecyobj2)
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						gettxton = true
						gettxtcharlimit = 3
						gettxtname = "labelevent_changedirecyobj2"
						gettxttype = 5
					}
				}

				x += rec.ToInt32().Width + txts

				buttontxt = "="
				if objs[activobjnum].labelchangedirecyplusobj2 {
					buttontxt = "+ ="
				}
				rec = buttonorange(buttontxt, x, y-(txts/2))
				if rl.CheckCollisionPointRec(mousev2, rec) {
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						if objs[activobjnum].labelchangedirecyplusobj2 {
							objs[activobjnum].labelchangedirecyplusobj2 = false
						} else {
							objs[activobjnum].labelchangedirecyplusobj2 = true
						}
					}
				}

				x += rec.ToInt32().Width + txts
			}

			// add activ obj event 2
			txtlen = rl.MeasureText("more obj 2 events ?", txts)
			rl.DrawText("more obj 2 events ?", x, y, txts, rl.White)
			x += txtlen + txts

			objs[activobjnum].labeleventobj2 = tickbox(x, y, objs[activobjnum].labeleventobj2)

			if objs[activobjnum].labeleventobj2 {
				x = int32(recwin.X + 50)
				y += txts * 3
				txtlen = rl.MeasureText("and obj 2", txts)
				rl.DrawText("and obj 2", x, y, txts, rl.White)
				x += txtlen + txts

				//text back rec
				backrec = rl.NewRectangle(float32(x), float32(y-(txts/2)), 130, float32(txts*2))
				rl.DrawRectangleRec(backrec, rl.Black)
				//list down arrow
				v2 = rl.NewVector2((backrec.X+backrec.Width)-(darrowimg.Width+4), backrec.Y+float32(txts/2))
				downlistrec = rl.NewRectangle(v2.X-4, backrec.Y, darrowimg.Width+8, backrec.Height)
				if rl.CheckCollisionPointRec(mousev2, downlistrec) {
					rl.DrawRectangleRec(downlistrec, brightred())
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						labelactionnum2obj2++
						if labelactionnum2obj2 > len(labelactionslistobj2)-1 {
							labelactionnum2obj2 = 0
						}
					}
				} else {
					rl.DrawRectangleRec(downlistrec, brightorange())
				}
				rl.DrawTextureRec(imgs, darrowimg, v2, rl.White)

				//label action 2 list txt
				x += 5
				txt = labelactionslistobj2[labelactionnum2obj2]
				rl.DrawText(txt, x, y, txts, rl.White)

				if labelactionslistobj2[labelactionnum2obj2] != "does nothing" {
					x = int32(recwin.X + 50)
					y += txts * 3

					switch labelactionslistobj2[labelactionnum2obj2] {
					case "changes hp":
						txtlen = rl.MeasureText("change amount + or -", txts)
						rl.DrawText("change amount + or -", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt := fmt.Sprint(objs[activobjnum].labelchangehpobj2)
						rec := buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 3
								gettxtname = "labelevent_labelchangehpobj2"
								gettxttype = 5
							}
						}

						x += rec.ToInt32().Width + txts

						buttontxt = "="
						if objs[activobjnum].labelchangehpplusobj2 {
							buttontxt = "+ ="
						}
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].labelchangehpplusobj2 {
									objs[activobjnum].labelchangehpplusobj2 = false
								} else {
									objs[activobjnum].labelchangehpplusobj2 = true
								}
							}
						}

						x += rec.ToInt32().Width + txts

					case "changes position":

						txtlen = rl.MeasureText("new x", txts)
						rl.DrawText("new x", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt := fmt.Sprint(objs[activobjnum].labelchangeposxobj2)
						rec := buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 3
								gettxtname = "labelevent_labelchangeposxobj2"
								gettxttype = 5
							}
						}

						x += rec.ToInt32().Width + txts

						buttontxt = "="
						if objs[activobjnum].labelchangeposxplusobj2 {
							buttontxt = "+ ="
						}
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].labelchangeposxplusobj2 {
									objs[activobjnum].labelchangeposxplusobj2 = false
								} else {
									objs[activobjnum].labelchangeposxplusobj2 = true
								}
							}
						}

						x += rec.ToInt32().Width + txts

						txtlen = rl.MeasureText("new y", txts)
						rl.DrawText("new y", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt = fmt.Sprint(objs[activobjnum].labelchangeposyobj2)
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 3
								gettxtname = "labelevent_labelchangeposyobj2"
								gettxttype = 5
							}
						}

						x += rec.ToInt32().Width + txts

						buttontxt = "="
						if objs[activobjnum].labelchangeposyplusobj2 {
							buttontxt = "+ ="
						}
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].labelchangeposyplusobj2 {
									objs[activobjnum].labelchangeposyplusobj2 = false
								} else {
									objs[activobjnum].labelchangeposyplusobj2 = true
								}
							}
						}

						x += rec.ToInt32().Width + txts

					case "changes layer":

						txtlen = rl.MeasureText("moves to layer ?", txts)
						rl.DrawText("moves to layer ?", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt := "middleground"
						if objs[activobjnum].labelchangelayerobj2 == 1 {
							buttontxt = "foreground"
						} else if objs[activobjnum].labelchangelayerobj2 == -1 {
							buttontxt = "background"
						}

						rec := buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								objs[activobjnum].labelchangelayerobj2++
								if objs[activobjnum].labelchangelayerobj2 > 1 {
									objs[activobjnum].labelchangelayerobj2 = -1
								}
							}
						}
						x += rec.ToInt32().Width + txts

						switch objs[activobjnum].labelchangelayerobj2 {
						case 0:
							if objs[activobjnum].middleground_obj {

								rl.DrawText("* obj 1 layer", x, y, txts, rl.White)
								txtlen = rl.MeasureText("* obj 1 layer", txts)
								x += txtlen + (txts * 2)
							}
						case -1:
							if objs[activobjnum].background_obj {

								rl.DrawText("* obj 1 layer", x, y, txts, rl.White)
								txtlen = rl.MeasureText("* obj 1 layer", txts)
								x += txtlen + (txts * 2)
							}
						case 1:
							if objs[activobjnum].foreground_obj {

								rl.DrawText("* obj 1 layer", x, y, txts, rl.White)
								txtlen = rl.MeasureText("* obj 1 layer", txts)
								x += txtlen + (txts * 2)
							}
						}

					case "changes shadow":

						txtlen = rl.MeasureText("shadow", txts)
						rl.DrawText("shadow", x, y, txts, rl.White)
						x += txtlen + txts

						onofftxt := "off"
						if objs[activobjnum].labelchangeshadowobj2 {
							onofftxt = "on"
						}
						rec := buttonorange(onofftxt, x, y-(txts/2))

						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].labelchangeshadowobj2 {
									objs[activobjnum].labelchangeshadowobj2 = false
								} else {
									objs[activobjnum].labelchangeshadowobj2 = true
								}
							}
						}

						x += rec.ToInt32().Width + txts

						txtlen = rl.MeasureText("x", txts)
						rl.DrawText("x", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt := fmt.Sprint(objs[activobjnum].labelchangeshadowxobj2)
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 3
								gettxtname = "labelevent_changeshadowxobj2"
								gettxttype = 5
							}
						}

						x += rec.ToInt32().Width + txts

						txtlen = rl.MeasureText("y", txts)
						rl.DrawText("y", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt = fmt.Sprint(objs[activobjnum].labelchangeshadowyobj2)
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 3
								gettxtname = "labelevent_changeshadowyobj2"
								gettxttype = 5
							}
						}

						x += rec.ToInt32().Width + txts

						txtlen = rl.MeasureText("fade", txts)
						rl.DrawText("fade", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt = fmt.Sprint(objs[activobjnum].labelchangeshadowfadeobj2)
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 1
								gettxtname = "labelevent_changeshadowfadeobj2"
								gettxttype = 2

							}
						}

						x += rec.ToInt32().Width + txts

						txtlen = rl.MeasureText("change color", txts)
						rl.DrawText("change color", x, y, txts, rl.White)
						x += txtlen + txts

						rec = rl.NewRectangle(float32(x), float32(y-5), 50, float32(txts+10))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							rl.DrawRectangleRec(rec, brightred())
							rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								changecolorsel = 15
								colorpalon = true
							}
						} else {
							rl.DrawRectangleRec(rec, objs[activobjnum].labelchangeshadowcolorobj2)
						}

						x += rec.ToInt32().Width + txts

					case "changes ghosting":

						txtlen = rl.MeasureText("ghosting", txts)
						rl.DrawText("ghosting", x, y, txts, rl.White)
						x += txtlen + txts

						onofftxt := "off"
						if objs[activobjnum].labelchangeghostingobj2 {
							onofftxt = "on"
						}
						rec := buttonorange(onofftxt, x, y-(txts/2))

						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].labelchangeghostingobj2 {
									objs[activobjnum].labelchangeghostingobj2 = false
								} else {
									objs[activobjnum].labelchangeghostingobj2 = true
								}
							}
						}

						x += rec.ToInt32().Width + txts

						txtlen = rl.MeasureText("x", txts)
						rl.DrawText("x", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt := fmt.Sprint(objs[activobjnum].labelchangeghostxobj2)
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 3
								gettxtname = "labelevent_changeghostxobj2"
								gettxttype = 5
							}
						}

						x += rec.ToInt32().Width + txts

						txtlen = rl.MeasureText("y", txts)
						rl.DrawText("y", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt = fmt.Sprint(objs[activobjnum].labelchangeghostyobj2)
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 3
								gettxtname = "labelevent_changeghostyobj2"
								gettxttype = 5
							}
						}

						x += rec.ToInt32().Width + txts

						txtlen = rl.MeasureText("fade", txts)
						rl.DrawText("fade", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt = fmt.Sprint(objs[activobjnum].labelchangeghostfadeobj2)
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 1
								gettxtname = "labelevent_changeghostfadeobj2"
								gettxttype = 2

							}
						}

						x += rec.ToInt32().Width + txts

						txtlen = rl.MeasureText("change color", txts)
						rl.DrawText("change color", x, y, txts, rl.White)
						x += txtlen + txts

						rec = rl.NewRectangle(float32(x), float32(y-5), 50, float32(txts+10))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							rl.DrawRectangleRec(rec, brightred())
							rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								changecolorsel = 14
								colorpalon = true
							}
						} else {
							rl.DrawRectangleRec(rec, objs[activobjnum].labelchangeghostcolorobj2)
						}

						x += rec.ToInt32().Width + txts

					case "changes fill":
						txtlen = rl.MeasureText("new fill color 1", txts)
						rl.DrawText("new fill color 1", x, y, txts, rl.White)
						x += txtlen + txts

						rec := rl.NewRectangle(float32(x), float32(y-5), 50, float32(txts+10))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							rl.DrawRectangleRec(rec, brightred())
							rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								changecolorsel = 12
								colorpalon = true
							}
						} else {
							rl.DrawRectangleRec(rec, objs[activobjnum].labelchangefill1obj2)
						}

						x += rec.ToInt32().Width + txts
						txtlen = rl.MeasureText("gradient fill ?", txts)
						rl.DrawText("gradient fill ?", x, y, txts, rl.White)
						x += txtlen + txts

						objs[activobjnum].labelchangefillgradientonobj2 = tickbox(x, y, objs[activobjnum].labelchangefillgradientonobj2)

						x += txts * 3

						if objs[activobjnum].labelchangefillgradientonobj2 {

							txtlen = rl.MeasureText("new fill color 2", txts)
							rl.DrawText("new fill color 2", x, y, txts, rl.White)
							x += txtlen + txts

							rec = rl.NewRectangle(float32(x), float32(y-5), 50, float32(txts+10))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								rl.DrawRectangleRec(rec, brightred())
								rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									changecolorsel = 13
									colorpalon = true
								}
							} else {
								rl.DrawRectangleRec(rec, objs[activobjnum].labelchangefill2obj2)
							}
							x += rec.ToInt32().Width + txts

							gradtxt := "horizontal gradient"
							if objs[activobjnum].labelchangefillgradienthvobj2 {
								gradtxt = "vertical gradient"
							}

							rec := buttonorange(gradtxt, x, y-(txts/2))

							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									if objs[activobjnum].labelchangefillgradienthvobj2 {
										objs[activobjnum].labelchangefillgradienthvobj2 = false
									} else {
										objs[activobjnum].labelchangefillgradienthvobj2 = true
									}
								}
							}
							x += rec.ToInt32().Width + txts
						}
					case "changes image":

						txtlen = rl.MeasureText("new image", txts)
						rl.DrawText("new image", x, y, txts, rl.White)
						x += txtlen + txts

						rec := buttonorange("select image", x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								labelchangeimgonobj2 = true
							}
						}

						x += rec.ToInt32().Width + txts

					case "changes size":

						txtlen = rl.MeasureText("new width", txts)
						rl.DrawText("new width", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt := fmt.Sprint(objs[activobjnum].labelchangewidthobj2)
						rec := buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 3
								gettxtname = "labelevent_changewidthobj2"
								gettxttype = 5
							}
						}

						x += rec.ToInt32().Width + txts

						buttontxt = "="
						if objs[activobjnum].labelchangewidthplusobj2 {
							buttontxt = "+ ="
						}
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].labelchangewidthplusobj2 {
									objs[activobjnum].labelchangewidthplusobj2 = false
								} else {
									objs[activobjnum].labelchangewidthplusobj2 = true
								}
							}
						}

						x += rec.ToInt32().Width + txts

						txtlen = rl.MeasureText("new height", txts)
						rl.DrawText("new height", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt = fmt.Sprint(objs[activobjnum].labelchangeheightobj2)
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 3
								gettxtname = "labelevent_changeheightobj2"
								gettxttype = 5
							}
						}
						x += rec.ToInt32().Width + txts

						buttontxt = "="
						if objs[activobjnum].labelchangeheightplusobj2 {
							buttontxt = "+ ="
						}
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].labelchangeheightplusobj2 {
									objs[activobjnum].labelchangeheightplusobj2 = false
								} else {
									objs[activobjnum].labelchangeheightplusobj2 = true
								}
							}
						}

						x += rec.ToInt32().Width + txts

					case "changes rotation":
						txtlen = rl.MeasureText("new angle", txts)
						rl.DrawText("new angle", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt := fmt.Sprint(objs[activobjnum].labelchangerotationobj2)
						rec := buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 3
								gettxtname = "labelevent_changerotationobj2"
								gettxttype = 5
							}
						}
						x += rec.ToInt32().Width + txts

						buttontxt = "="
						if objs[activobjnum].labelchangerotationplusobj2 {
							buttontxt = "+ ="
						}
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].labelchangerotationplusobj2 {
									objs[activobjnum].labelchangerotationplusobj2 = false
								} else {
									objs[activobjnum].labelchangerotationplusobj2 = true
								}
							}
						}

						x += rec.ToInt32().Width + txts

					case "changes direction":

						txtlen = rl.MeasureText("new direction x", txts)
						rl.DrawText("new direction x", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt := fmt.Sprint(objs[activobjnum].labelchangedirecxobj2)
						rec := buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 3
								gettxtname = "labelevent_changedirecxobj2"
								gettxttype = 5
							}
						}

						x += rec.ToInt32().Width + txts

						buttontxt = "="
						if objs[activobjnum].labelchangedirecxplusobj2 {
							buttontxt = "+ ="
						}
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].labelchangedirecxplusobj2 {
									objs[activobjnum].labelchangedirecxplusobj2 = false
								} else {
									objs[activobjnum].labelchangedirecxplusobj2 = true
								}
							}
						}

						x += rec.ToInt32().Width + txts

						txtlen = rl.MeasureText("new direction y", txts)
						rl.DrawText("new direction y", x, y, txts, rl.White)
						x += txtlen + txts

						buttontxt = fmt.Sprint(objs[activobjnum].labelchangedirecyobj2)
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								gettxton = true
								gettxtcharlimit = 3
								gettxtname = "labelevent_changedirecyobj2"
								gettxttype = 5
							}
						}

						x += rec.ToInt32().Width + txts

						buttontxt = "="
						if objs[activobjnum].labelchangedirecyplusobj2 {
							buttontxt = "+ ="
						}
						rec = buttonorange(buttontxt, x, y-(txts/2))
						if rl.CheckCollisionPointRec(mousev2, rec) {
							if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
								if objs[activobjnum].labelchangedirecyplusobj2 {
									objs[activobjnum].labelchangedirecyplusobj2 = false
								} else {
									objs[activobjnum].labelchangedirecyplusobj2 = true
								}
							}
						}

						x += rec.ToInt32().Width + txts
					}

					txtlen = rl.MeasureText("more obj 2 events ?", txts)
					rl.DrawText("more obj 2 events ?", x, y, txts, rl.White)
					x += txtlen + txts

					objs[activobjnum].labelevent2obj2 = tickbox(x, y, objs[activobjnum].labelevent2obj2)
				}
				// add activ obj event 3

				if objs[activobjnum].labelevent2obj2 {
					x = int32(recwin.X + 50)
					y += txts * 3
					txtlen = rl.MeasureText("and obj 2", txts)
					rl.DrawText("and obj 2", x, y, txts, rl.White)
					x += txtlen + txts

					//text back rec
					backrec = rl.NewRectangle(float32(x), float32(y-(txts/2)), 130, float32(txts*2))
					rl.DrawRectangleRec(backrec, rl.Black)
					//list down arrow
					v2 = rl.NewVector2((backrec.X+backrec.Width)-(darrowimg.Width+4), backrec.Y+float32(txts/2))
					downlistrec = rl.NewRectangle(v2.X-4, backrec.Y, darrowimg.Width+8, backrec.Height)
					if rl.CheckCollisionPointRec(mousev2, downlistrec) {
						rl.DrawRectangleRec(downlistrec, brightred())
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							labelactionnum3obj2++
							if labelactionnum3obj2 > len(labelactionslistobj2)-1 {
								labelactionnum3obj2 = 0
							}
						}
					} else {
						rl.DrawRectangleRec(downlistrec, brightorange())
					}
					rl.DrawTextureRec(imgs, darrowimg, v2, rl.White)

					//label action 3 list txt
					x += 5
					txt = labelactionslistobj2[labelactionnum3obj2]
					rl.DrawText(txt, x, y, txts, rl.White)

					if labelactionslistobj2[labelactionnum3obj2] != "does nothing" {
						x = int32(recwin.X + 50)
						y += txts * 3

						switch labelactionslistobj2[labelactionnum3obj2] {
						case "changes hp":
							txtlen = rl.MeasureText("change amount + or -", txts)
							rl.DrawText("change amount + or -", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt := fmt.Sprint(objs[activobjnum].labelchangehpobj2)
							rec := buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 3
									gettxtname = "labelevent_labelchangehpobj2"
									gettxttype = 5
								}
							}

							x += rec.ToInt32().Width + txts

							buttontxt = "="
							if objs[activobjnum].labelchangehpplusobj2 {
								buttontxt = "+ ="
							}
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									if objs[activobjnum].labelchangehpplusobj2 {
										objs[activobjnum].labelchangehpplusobj2 = false
									} else {
										objs[activobjnum].labelchangehpplusobj2 = true
									}
								}
							}

							x += rec.ToInt32().Width + txts

						case "changes position":

							txtlen = rl.MeasureText("new x", txts)
							rl.DrawText("new x", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt := fmt.Sprint(objs[activobjnum].labelchangeposxobj2)
							rec := buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 3
									gettxtname = "labelevent_labelchangeposxobj2"
									gettxttype = 5
								}
							}

							x += rec.ToInt32().Width + txts

							buttontxt = "="
							if objs[activobjnum].labelchangeposxplusobj2 {
								buttontxt = "+ ="
							}
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									if objs[activobjnum].labelchangeposxplusobj2 {
										objs[activobjnum].labelchangeposxplusobj2 = false
									} else {
										objs[activobjnum].labelchangeposxplusobj2 = true
									}
								}
							}

							x += rec.ToInt32().Width + txts

							txtlen = rl.MeasureText("new y", txts)
							rl.DrawText("new y", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt = fmt.Sprint(objs[activobjnum].labelchangeposyobj2)
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 3
									gettxtname = "labelevent_labelchangeposyobj2"
									gettxttype = 5
								}
							}

							x += rec.ToInt32().Width + txts

							buttontxt = "="
							if objs[activobjnum].labelchangeposyplusobj2 {
								buttontxt = "+ ="
							}
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									if objs[activobjnum].labelchangeposyplusobj2 {
										objs[activobjnum].labelchangeposyplusobj2 = false
									} else {
										objs[activobjnum].labelchangeposyplusobj2 = true
									}
								}
							}

							x += rec.ToInt32().Width + txts

						case "changes layer":

							txtlen = rl.MeasureText("moves to layer ?", txts)
							rl.DrawText("moves to layer ?", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt := "middleground"
							if objs[activobjnum].labelchangelayerobj2 == 1 {
								buttontxt = "foreground"
							} else if objs[activobjnum].labelchangelayerobj2 == -1 {
								buttontxt = "background"
							}

							rec := buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									objs[activobjnum].labelchangelayerobj2++
									if objs[activobjnum].labelchangelayerobj2 > 1 {
										objs[activobjnum].labelchangelayerobj2 = -1
									}
								}
							}
							x += rec.ToInt32().Width + txts

							switch objs[activobjnum].labelchangelayerobj2 {
							case 0:
								if objs[activobjnum].middleground_obj {

									rl.DrawText("* obj 1 layer", x, y, txts, rl.White)
									txtlen = rl.MeasureText("* obj 1 layer", txts)
									x += txtlen + (txts * 2)
								}
							case -1:
								if objs[activobjnum].background_obj {

									rl.DrawText("* obj 1 layer", x, y, txts, rl.White)
									txtlen = rl.MeasureText("* obj 1 layer", txts)
									x += txtlen + (txts * 2)
								}
							case 1:
								if objs[activobjnum].foreground_obj {

									rl.DrawText("* obj 1 layer", x, y, txts, rl.White)
									txtlen = rl.MeasureText("* obj 1 layer", txts)
									x += txtlen + (txts * 2)
								}
							}

						case "changes shadow":

							txtlen = rl.MeasureText("shadow", txts)
							rl.DrawText("shadow", x, y, txts, rl.White)
							x += txtlen + txts

							onofftxt := "off"
							if objs[activobjnum].labelchangeshadowobj2 {
								onofftxt = "on"
							}
							rec := buttonorange(onofftxt, x, y-(txts/2))

							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									if objs[activobjnum].labelchangeshadowobj2 {
										objs[activobjnum].labelchangeshadowobj2 = false
									} else {
										objs[activobjnum].labelchangeshadowobj2 = true
									}
								}
							}

							x += rec.ToInt32().Width + txts

							txtlen = rl.MeasureText("x", txts)
							rl.DrawText("x", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt := fmt.Sprint(objs[activobjnum].labelchangeshadowxobj2)
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 3
									gettxtname = "labelevent_changeshadowxobj2"
									gettxttype = 5
								}
							}

							x += rec.ToInt32().Width + txts

							txtlen = rl.MeasureText("y", txts)
							rl.DrawText("y", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt = fmt.Sprint(objs[activobjnum].labelchangeshadowyobj2)
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 3
									gettxtname = "labelevent_changeshadowyobj2"
									gettxttype = 5
								}
							}

							x += rec.ToInt32().Width + txts

							txtlen = rl.MeasureText("fade", txts)
							rl.DrawText("fade", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt = fmt.Sprint(objs[activobjnum].labelchangeshadowfadeobj2)
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 1
									gettxtname = "labelevent_changeshadowfadeobj2"
									gettxttype = 2

								}
							}

							x += rec.ToInt32().Width + txts

							txtlen = rl.MeasureText("change color", txts)
							rl.DrawText("change color", x, y, txts, rl.White)
							x += txtlen + txts

							rec = rl.NewRectangle(float32(x), float32(y-5), 50, float32(txts+10))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								rl.DrawRectangleRec(rec, brightred())
								rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									changecolorsel = 15
									colorpalon = true
								}
							} else {
								rl.DrawRectangleRec(rec, objs[activobjnum].labelchangeshadowcolorobj2)
							}

							x += rec.ToInt32().Width + txts

						case "changes ghosting":

							txtlen = rl.MeasureText("ghosting", txts)
							rl.DrawText("ghosting", x, y, txts, rl.White)
							x += txtlen + txts

							onofftxt := "off"
							if objs[activobjnum].labelchangeghostingobj2 {
								onofftxt = "on"
							}
							rec := buttonorange(onofftxt, x, y-(txts/2))

							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									if objs[activobjnum].labelchangeghostingobj2 {
										objs[activobjnum].labelchangeghostingobj2 = false
									} else {
										objs[activobjnum].labelchangeghostingobj2 = true
									}
								}
							}

							x += rec.ToInt32().Width + txts

							txtlen = rl.MeasureText("x", txts)
							rl.DrawText("x", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt := fmt.Sprint(objs[activobjnum].labelchangeghostxobj2)
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 3
									gettxtname = "labelevent_changeghostxobj2"
									gettxttype = 5
								}
							}

							x += rec.ToInt32().Width + txts

							txtlen = rl.MeasureText("y", txts)
							rl.DrawText("y", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt = fmt.Sprint(objs[activobjnum].labelchangeghostyobj2)
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 3
									gettxtname = "labelevent_changeghostyobj2"
									gettxttype = 5
								}
							}

							x += rec.ToInt32().Width + txts

							txtlen = rl.MeasureText("fade", txts)
							rl.DrawText("fade", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt = fmt.Sprint(objs[activobjnum].labelchangeghostfadeobj2)
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 1
									gettxtname = "labelevent_changeghostfadeobj2"
									gettxttype = 2

								}
							}

							x += rec.ToInt32().Width + txts

							txtlen = rl.MeasureText("change color", txts)
							rl.DrawText("change color", x, y, txts, rl.White)
							x += txtlen + txts

							rec = rl.NewRectangle(float32(x), float32(y-5), 50, float32(txts+10))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								rl.DrawRectangleRec(rec, brightred())
								rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									changecolorsel = 14
									colorpalon = true
								}
							} else {
								rl.DrawRectangleRec(rec, objs[activobjnum].labelchangeghostcolorobj2)
							}

							x += rec.ToInt32().Width + txts

						case "changes fill":
							txtlen = rl.MeasureText("new fill color 1", txts)
							rl.DrawText("new fill color 1", x, y, txts, rl.White)
							x += txtlen + txts

							rec := rl.NewRectangle(float32(x), float32(y-5), 50, float32(txts+10))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								rl.DrawRectangleRec(rec, brightred())
								rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									changecolorsel = 12
									colorpalon = true
								}
							} else {
								rl.DrawRectangleRec(rec, objs[activobjnum].labelchangefill1obj2)
							}

							x += rec.ToInt32().Width + txts
							txtlen = rl.MeasureText("gradient fill ?", txts)
							rl.DrawText("gradient fill ?", x, y, txts, rl.White)
							x += txtlen + txts

							objs[activobjnum].labelchangefillgradientonobj2 = tickbox(x, y, objs[activobjnum].labelchangefillgradientonobj2)

							x += txts * 3

							if objs[activobjnum].labelchangefillgradientonobj2 {

								txtlen = rl.MeasureText("new fill color 2", txts)
								rl.DrawText("new fill color 2", x, y, txts, rl.White)
								x += txtlen + txts

								rec = rl.NewRectangle(float32(x), float32(y-5), 50, float32(txts+10))
								if rl.CheckCollisionPointRec(mousev2, rec) {
									rl.DrawRectangleRec(rec, brightred())
									rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
									if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
										changecolorsel = 13
										colorpalon = true
									}
								} else {
									rl.DrawRectangleRec(rec, objs[activobjnum].labelchangefill2obj2)
								}
								x += rec.ToInt32().Width + txts

								gradtxt := "horizontal gradient"
								if objs[activobjnum].labelchangefillgradienthvobj2 {
									gradtxt = "vertical gradient"
								}

								rec := buttonorange(gradtxt, x, y-(txts/2))

								if rl.CheckCollisionPointRec(mousev2, rec) {
									if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
										if objs[activobjnum].labelchangefillgradienthvobj2 {
											objs[activobjnum].labelchangefillgradienthvobj2 = false
										} else {
											objs[activobjnum].labelchangefillgradienthvobj2 = true
										}
									}
								}
								x += rec.ToInt32().Width + txts
							}
						case "changes image":

							txtlen = rl.MeasureText("new image", txts)
							rl.DrawText("new image", x, y, txts, rl.White)
							x += txtlen + txts

							rec := buttonorange("select image", x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									labelchangeimgonobj2 = true
								}
							}

							x += rec.ToInt32().Width + txts

						case "changes size":

							txtlen = rl.MeasureText("new width", txts)
							rl.DrawText("new width", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt := fmt.Sprint(objs[activobjnum].labelchangewidthobj2)
							rec := buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 3
									gettxtname = "labelevent_changewidthobj2"
									gettxttype = 5
								}
							}

							x += rec.ToInt32().Width + txts

							buttontxt = "="
							if objs[activobjnum].labelchangewidthplusobj2 {
								buttontxt = "+ ="
							}
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									if objs[activobjnum].labelchangewidthplusobj2 {
										objs[activobjnum].labelchangewidthplusobj2 = false
									} else {
										objs[activobjnum].labelchangewidthplusobj2 = true
									}
								}
							}

							x += rec.ToInt32().Width + txts

							txtlen = rl.MeasureText("new height", txts)
							rl.DrawText("new height", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt = fmt.Sprint(objs[activobjnum].labelchangeheightobj2)
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 3
									gettxtname = "labelevent_changeheightobj2"
									gettxttype = 5
								}
							}
							x += rec.ToInt32().Width + txts

							buttontxt = "="
							if objs[activobjnum].labelchangeheightplusobj2 {
								buttontxt = "+ ="
							}
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									if objs[activobjnum].labelchangeheightplusobj2 {
										objs[activobjnum].labelchangeheightplusobj2 = false
									} else {
										objs[activobjnum].labelchangeheightplusobj2 = true
									}
								}
							}

							x += rec.ToInt32().Width + txts

						case "changes rotation":
							txtlen = rl.MeasureText("new angle", txts)
							rl.DrawText("new angle", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt := fmt.Sprint(objs[activobjnum].labelchangerotationobj2)
							rec := buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 3
									gettxtname = "labelevent_changerotationobj2"
									gettxttype = 5
								}
							}
							x += rec.ToInt32().Width + txts

							buttontxt = "="
							if objs[activobjnum].labelchangerotationplusobj2 {
								buttontxt = "+ ="
							}
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									if objs[activobjnum].labelchangerotationplusobj2 {
										objs[activobjnum].labelchangerotationplusobj2 = false
									} else {
										objs[activobjnum].labelchangerotationplusobj2 = true
									}
								}
							}

							x += rec.ToInt32().Width + txts

						case "changes direction":

							txtlen = rl.MeasureText("new direction x", txts)
							rl.DrawText("new direction x", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt := fmt.Sprint(objs[activobjnum].labelchangedirecxobj2)
							rec := buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 3
									gettxtname = "labelevent_changedirecxobj2"
									gettxttype = 5
								}
							}

							x += rec.ToInt32().Width + txts

							buttontxt = "="
							if objs[activobjnum].labelchangedirecxplusobj2 {
								buttontxt = "+ ="
							}
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									if objs[activobjnum].labelchangedirecxplusobj2 {
										objs[activobjnum].labelchangedirecxplusobj2 = false
									} else {
										objs[activobjnum].labelchangedirecxplusobj2 = true
									}
								}
							}

							x += rec.ToInt32().Width + txts

							txtlen = rl.MeasureText("new direction y", txts)
							rl.DrawText("new direction y", x, y, txts, rl.White)
							x += txtlen + txts

							buttontxt = fmt.Sprint(objs[activobjnum].labelchangedirecyobj2)
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									gettxton = true
									gettxtcharlimit = 3
									gettxtname = "labelevent_changedirecyobj2"
									gettxttype = 5
								}
							}

							x += rec.ToInt32().Width + txts

							buttontxt = "="
							if objs[activobjnum].labelchangedirecyplusobj2 {
								buttontxt = "+ ="
							}
							rec = buttonorange(buttontxt, x, y-(txts/2))
							if rl.CheckCollisionPointRec(mousev2, rec) {
								if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
									if objs[activobjnum].labelchangedirecyplusobj2 {
										objs[activobjnum].labelchangedirecyplusobj2 = false
									} else {
										objs[activobjnum].labelchangedirecyplusobj2 = true
									}
								}
							}

							x += rec.ToInt32().Width + txts
						}
					}
				}
			}
		}

		x = int32(recwin.X + 50)
		y += txts * 4

		txtlen = rl.MeasureText("spawn more objs ? ", txts)
		rl.DrawText("spawn more objs ? ", x, y, txts, rl.White)
		x += txtlen + txts
		labelspawnon = tickbox(x, y, labelspawnon)

		if labelspawnon {
			x = int32(recwin.X + 50)
			y += txts * 3

			txtlen = rl.MeasureText("and spawns other objs", txts)
			rl.DrawText("and spawns other objs", x, y, txts, rl.White)
			x += txtlen + txts

			//text back rec
			backrec = rl.NewRectangle(float32(x), float32(y-(txts/2)), 130, float32(txts*2))
			rl.DrawRectangleRec(backrec, rl.Black)
			//list down arrow
			if len(usrsaveobjs) > 0 {
				v2 = rl.NewVector2((backrec.X+backrec.Width)-(darrowimg.Width+4), backrec.Y+float32(txts/2))
				downlistrec = rl.NewRectangle(v2.X-4, backrec.Y, darrowimg.Width+8, backrec.Height)
				if rl.CheckCollisionPointRec(mousev2, downlistrec) {
					rl.DrawRectangleRec(downlistrec, brightred())
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						usrsaveobjnum++
						if usrsaveobjnum > len(usrsaveobjs)-1 {
							usrsaveobjnum = 0
						}
					}
				} else {
					rl.DrawRectangleRec(downlistrec, brightorange())
				}
				rl.DrawTextureRec(imgs, darrowimg, v2, rl.White)

				//usr save objs list txt
				x += 5
				txt = usrsaveobjs[usrsaveobjnum].name
				rl.DrawText(txt, x, y, txts, rl.White)
			} else {
				x += 5
				rl.DrawText("no objs saved", x, y, txts, rl.White)
			}

			x = int32(backrec.X+backrec.Width) + txts
			txtlen = rl.MeasureText("number to spawn", txts)
			rl.DrawText("number to spawn", x, y, txts, rl.White)
			x += txtlen + txts

			labelspawnnum, x = numbox(x, y, 1, 10, labelspawnnum)
			x += txts

			txtlen = rl.MeasureText("spawn more objs ? ", txts)
			rl.DrawText("spawn more objs ? ", x, y, txts, rl.White)
			x += txtlen + txts

			labelspawnon2 = tickbox(x, y, labelspawnon2)

			x += txts * 4

		}

		if labelspawnon2 {
			labelspawnon = true
			x = int32(recwin.X + 50)
			y += txts * 3

			txtlen = rl.MeasureText("and spawns other objs", txts)
			rl.DrawText("and spawns other objs", x, y, txts, rl.White)
			x += txtlen + txts

			//text back rec
			backrec = rl.NewRectangle(float32(x), float32(y-(txts/2)), 130, float32(txts*2))
			rl.DrawRectangleRec(backrec, rl.Black)
			//list down arrow
			if len(usrsaveobjs) > 0 {
				v2 = rl.NewVector2((backrec.X+backrec.Width)-(darrowimg.Width+4), backrec.Y+float32(txts/2))
				downlistrec = rl.NewRectangle(v2.X-4, backrec.Y, darrowimg.Width+8, backrec.Height)
				if rl.CheckCollisionPointRec(mousev2, downlistrec) {
					rl.DrawRectangleRec(downlistrec, brightred())
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						usrsaveobjnum2++
						if usrsaveobjnum2 > len(usrsaveobjs)-1 {
							usrsaveobjnum2 = 0
						}
					}
				} else {
					rl.DrawRectangleRec(downlistrec, brightorange())
				}
				rl.DrawTextureRec(imgs, darrowimg, v2, rl.White)

				//usr save objs list txt
				x += 5
				txt = usrsaveobjs[usrsaveobjnum2].name
				rl.DrawText(txt, x, y, txts, rl.White)
			} else {
				x += 5
				rl.DrawText("no objs saved", x, y, txts, rl.White)
			}

			x = int32(backrec.X+backrec.Width) + txts
			txtlen = rl.MeasureText("number to spawn", txts)
			rl.DrawText("number to spawn", x, y, txts, rl.White)
			x += txtlen + txts

			labelspawnnum2, x = numbox(x, y, 1, 10, labelspawnnum2)
			x += txts

			txtlen = rl.MeasureText("spawn more objs ? ", txts)
			rl.DrawText("spawn more objs ? ", x, y, txts, rl.White)
			x += txtlen + txts

			labelspawnon3 = tickbox(x, y, labelspawnon3)
		}
		if labelspawnon3 {
			labelspawnon2 = true
			x = int32(recwin.X + 50)
			y += txts * 3

			txtlen = rl.MeasureText("and spawns other objs", txts)
			rl.DrawText("and spawns other objs", x, y, txts, rl.White)
			x += txtlen + txts

			//text back rec
			backrec = rl.NewRectangle(float32(x), float32(y-(txts/2)), 130, float32(txts*2))
			rl.DrawRectangleRec(backrec, rl.Black)
			//list down arrow
			if len(usrsaveobjs) > 0 {
				v2 = rl.NewVector2((backrec.X+backrec.Width)-(darrowimg.Width+4), backrec.Y+float32(txts/2))
				downlistrec = rl.NewRectangle(v2.X-4, backrec.Y, darrowimg.Width+8, backrec.Height)
				if rl.CheckCollisionPointRec(mousev2, downlistrec) {
					rl.DrawRectangleRec(downlistrec, brightred())
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						usrsaveobjnum3++
						if usrsaveobjnum3 > len(usrsaveobjs)-1 {
							usrsaveobjnum3 = 0
						}
					}
				} else {
					rl.DrawRectangleRec(downlistrec, brightorange())
				}
				rl.DrawTextureRec(imgs, darrowimg, v2, rl.White)

				//usr save objs list txt
				x += 5
				txt = usrsaveobjs[usrsaveobjnum3].name
				rl.DrawText(txt, x, y, txts, rl.White)
			} else {
				x += 5
				rl.DrawText("no objs saved", x, y, txts, rl.White)
			}

			x = int32(backrec.X+backrec.Width) + txts
			txtlen = rl.MeasureText("number to spawn", txts)
			rl.DrawText("number to spawn", x, y, txts, rl.White)
			x += txtlen + txts

			labelspawnnum3, x = numbox(x, y, 1, 10, labelspawnnum3)
			x += txts

			txtlen = rl.MeasureText("spawn more objs ? ", txts)
			rl.DrawText("spawn more objs ? ", x, y, txts, rl.White)
			x += txtlen + txts

			labelspawnon4 = tickbox(x, y, labelspawnon4)

		}
		if labelspawnon4 {

			x = int32(recwin.X + 50)
			y += txts * 3

			txtlen = rl.MeasureText("and spawns other objs", txts)
			rl.DrawText("and spawns other objs", x, y, txts, rl.White)
			x += txtlen + txts

			//text back rec
			backrec = rl.NewRectangle(float32(x), float32(y-(txts/2)), 130, float32(txts*2))
			rl.DrawRectangleRec(backrec, rl.Black)
			//list down arrow
			if len(usrsaveobjs) > 0 {
				v2 = rl.NewVector2((backrec.X+backrec.Width)-(darrowimg.Width+4), backrec.Y+float32(txts/2))
				downlistrec = rl.NewRectangle(v2.X-4, backrec.Y, darrowimg.Width+8, backrec.Height)
				if rl.CheckCollisionPointRec(mousev2, downlistrec) {
					rl.DrawRectangleRec(downlistrec, brightred())
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						usrsaveobjnum4++
						if usrsaveobjnum4 > len(usrsaveobjs)-1 {
							usrsaveobjnum4 = 0
						}
					}
				} else {
					rl.DrawRectangleRec(downlistrec, brightorange())
				}
				rl.DrawTextureRec(imgs, darrowimg, v2, rl.White)

				//usr save objs list txt
				x += 5
				txt = usrsaveobjs[usrsaveobjnum4].name
				rl.DrawText(txt, x, y, txts, rl.White)
			} else {
				x += 5
				rl.DrawText("no objs saved", x, y, txts, rl.White)
			}

			x = int32(backrec.X+backrec.Width) + txts
			txtlen = rl.MeasureText("number to spawn", txts)
			rl.DrawText("number to spawn", x, y, txts, rl.White)
			x += txtlen + txts

			labelspawnnum4, x = numbox(x, y, 1, 10, labelspawnnum4)

		}

		x = int32(recwin.X + 50)
		y += txts * 4

		//label event name
		txtlen = rl.MeasureText("label name : ", txts)
		rl.DrawText("label name : ", x, y, txts, rl.White)
		x += txtlen + txts

		labeleventnametxt = deflabeleventnamechanged

		if deflabeleventnamechanged == deflabeleventname {
			labeleventnametxt = deflabeleventname + " " + strconv.Itoa(len(objs[activobjnum].labelevents)+1)
		}

		rec := buttonorange(labeleventnametxt, x, y-(txts/2))
		if rl.CheckCollisionPointRec(mousev2, rec) {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				gettxton = true
				gettxtcharlimit = 20
				gettxtname = "labelevent_rename"
				gettxttype = 3
			}
		}

		y = int32(recwin.Y + recwin.Height - 50)
		x = int32(recwin.X + recwin.Width - 100)

		//save rec
		saverec := buttonline("save", x, y)

		if rl.CheckCollisionPointRec(mousev2, saverec) {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				savelabeleventacitvobj()
			}

		}

	case "editlabelall":

		txthere("edit label", true, txtm, recwin.X, recwin.Y+20, recwin.Width)

		xtab := int32(recwin.X + 50)
		ytab := int32(recwin.Y + 100)

		rec := buttonline("delete", xtab, ytab)
		if rl.CheckCollisionPointRec(mousev2, rec) {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				labelsalllist = remstring(labelsalllist, editlabelnum)
				editlabelallon = false
			}
		}
		xtab += rec.ToInt32().Width + txts

		rec = buttonline("rename", xtab, ytab)
		if rl.CheckCollisionPointRec(mousev2, rec) {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				editlabelallon = false
				gettxton = true
				gettxtcharlimit = 20
				gettxtname = "labelall_rename"
				gettxttype = 3

			}
		}

	case "editlabelactivobj":

		txthere("edit label", true, txtm, recwin.X, recwin.Y+20, recwin.Width)

		xtab := int32(recwin.X + 50)
		ytab := int32(recwin.Y + 100)

		rec := buttonline("delete", xtab, ytab)
		if rl.CheckCollisionPointRec(mousev2, rec) {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				objs[activobjnum].labels = remstring(objs[activobjnum].labels, editlabelnum)
				editlabelactivobjon = false
				menufocus = false
			}
		}
		xtab += rec.ToInt32().Width + txts

		rec = buttonline("rename", xtab, ytab)
		if rl.CheckCollisionPointRec(mousev2, rec) {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				editlabelactivobjon = false
				menufocus = false
				gettxton = true
				gettxtcharlimit = 20
				gettxtname = "obj_label_rename"
				gettxttype = 3

			}
		}

	}

}
func listgeneric(name string) { //MARK: listgeneric
	menufocus = true
	gettxtreturn = ""

	//back recs
	rl.DrawRectangle(0, 0, scrw, scrh, rl.Fade(rl.Black, 0.4))
	backrec := rl.NewRectangle(scrwf32/2-200, 0, 400, scrhf32)
	rl.DrawRectangleRec(backrec, rl.Black)

	//close win
	v2 := rl.NewVector2(backrec.X+backrec.Width-(closewinimg.Width*2), backrec.Y+closewinimg.Height)
	rec2 := rl.NewRectangle(v2.X, v2.Y, closewinimg.Width, closewinimg.Height)
	if rl.CheckCollisionPointRec(mousev2, rec2) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			viewlabelsliston, addexistinglabelon = false, false
			menufocus = false
		}
		rl.DrawTextureRec(imgs, closewinimg, v2, brightred())
	} else {
		rl.DrawTextureRec(imgs, closewinimg, v2, rl.White)
	}

	x := backrec.X + 10
	y := backrec.Y + 50
	//	x32 := int32(x)
	//	y32 := int32(y)

	switch name {

	case "addexistinglabel":
		txthere("add existing label to activ obj", true, txtm, x, y, backrec.Width)
		y += float32(txtm * 2)
		if len(labelsalllist) > 0 {
			for a := 0; a < len(labelsalllist); a++ {
				selectrec := rl.NewRectangle(backrec.X, y-4, backrec.Width, float32(txts+8))
				if rl.CheckCollisionPointRec(mousev2, selectrec) {
					rl.DrawRectangleRec(selectrec, rl.Fade(darkred(), 0.4))
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						objs[activobjnum].labels = append(objs[activobjnum].labels, labelsalllist[a])
					}
				}
				txthere(labelsalllist[a], true, txts, x, y, backrec.Width)
				y += float32(txts + (txts / 2))
			}

		} else {
			txthere("there are no saved labels", true, txts, x, y, backrec.Width)
		}
	case "labelslist":
		txthere("all labels", true, txtm, x, y, backrec.Width)
		y += float32(txtm * 2)
		if len(labelsalllist) > 0 {
			for a := 0; a < len(labelsalllist); a++ {
				selectrec := rl.NewRectangle(backrec.X, y-4, backrec.Width, float32(txts+8))
				if rl.CheckCollisionPointRec(mousev2, selectrec) {
					rl.DrawRectangleRec(selectrec, rl.Fade(darkred(), 0.4))
					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						editlabelallon = true
						editlabelnum = a
					}
				}
				txthere(labelsalllist[a], true, txts, x, y, backrec.Width)
				y += float32(txts + (txts / 2))
			}

		} else {
			txthere("there are no saved labels", true, txts, x, y, backrec.Width)
		}
	}

}
func switcheventsmenu(num int) { //MARK: switcheventsmenu

	for a := 0; a < len(eventsonoff); a++ {
		eventsonoff[a] = false
	}
	if num != blanknum {
		eventsonoff[num] = true
	}

}
func infobarmenu() { //MARK: infobarmenu

	x := txts
	y := int32(infobar.rec.Y + 5)

	//play restart icons
	restartv2 := rl.NewVector2(float32(x), float32(y+(txts/2)))
	checkrec := rl.NewRectangle(restartv2.X, restartv2.Y, refreshimg.Width, refreshimg.Height)
	if rl.CheckCollisionPointRec(mousev2, checkrec) {
		rl.DrawTextureRec(imgs, refreshimg, restartv2, brightred())
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

		}
	} else {
		rl.DrawTextureRec(imgs, refreshimg, restartv2, rl.White)
	}
	x += refreshimg.ToInt32().Width + txts
	playarrowv2 := rl.NewVector2(float32(x), float32(y+(txts/2)))
	checkrec = rl.NewRectangle(playarrowv2.X, playarrowv2.Y, playimg.Width, playimg.Height)
	if rl.CheckCollisionPointRec(mousev2, checkrec) {
		rl.DrawTextureRec(imgs, playimg, playarrowv2, brightred())
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if playon {
				playon = false
			} else {
				playon = true
			}
		}
	} else {
		if playon {
			rl.DrawTextureRec(imgs, playimg, playarrowv2, brightorange())
		} else {
			rl.DrawTextureRec(imgs, playimg, playarrowv2, rl.White)
		}
	}

	x += playimg.ToInt32().Width * 2
	xorig := x

	//objs info
	if objmenuon {
		//obj cnt
		txt := strconv.Itoa(len(objs))
		txtlen := rl.MeasureText(txt, txts)
		rl.DrawText(txt, x, y, txts, rl.White)
		x += txtlen + (txts / 2)
		rl.DrawText("obj cnt", x, y, txts, rl.White)

		x = xorig
		y += txts
		y += txts / 2

		//activ obj num
		txt = strconv.Itoa(activobjnum)
		if activobjnum == blanknum {
			txt = " "
		}
		txtlen = rl.MeasureText(txt, txts)
		rl.DrawText(txt, x, y, txts, rl.White)
		x += txtlen + (txts / 2)
		rl.DrawText("activ obj num", x, y, txts, rl.White)
		txtlen = rl.MeasureText("activ obj num", txts)

		//obj tools
		x = menu.rec.ToInt32().Width
		y = int32(infobar.rec.Y+infobar.rec.Height/2) - txts/2

		xf32 := float32(x)
		yf32 := float32(y)
		yf32 -= float32(txts / 2)

		rec := blankrec
		if activobjnum != blanknum {
			xf32, yf32, rec = buttonmainmenu("delete", xf32, yf32, blankbool)
			if rl.CheckCollisionPointRec(mousev2, rec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					objs = remobjs(objs, activobjnum)
					activobjnum = blanknum
				}
			}
			xf32, yf32, rec = buttonmainmenu("copy", xf32, yf32, copyobjon)
			if rl.CheckCollisionPointRec(mousev2, rec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					if copyobjon {
						copyobjon = false
					} else {
						copyobjon = true
					}
				}
			}
			xf32, yf32, rec = buttonmainmenu("events", xf32, yf32, neweventon)
			if rl.CheckCollisionPointRec(mousev2, rec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					if neweventon {
						neweventon = false
					} else {
						neweventon = true
					}
				}
			}
			xf32, yf32, rec = buttonmainmenu("controls", xf32, yf32, addobjcontrols)
			if rl.CheckCollisionPointRec(mousev2, rec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					if addobjcontrols {
						addobjcontrols = false
					} else {
						addobjcontrols = true
					}
				}
			}
			xf32, yf32, rec = buttonmainmenu("path", xf32, yf32, addobjpath)
			if rl.CheckCollisionPointRec(mousev2, rec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					if addobjpath {
						addobjpath = false
					} else {
						addobjpath = true
					}
				}
			}
			xf32, yf32, rec = buttonmainmenu("labels", xf32, yf32, addlabelon)
			if rl.CheckCollisionPointRec(mousev2, rec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					if addlabelon {
						addlabelon = false
					} else {
						addlabelon = true
					}
				}
			}
		}

	}

	//time settings icons
	inforicons()
}
func inforicons() { //MARK: inforicons

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
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			settingson = true
			menufocus = true
		}
	} else {
		rl.DrawTextureRec(imgs, settingsimg, v2, rl.White)
	}

	//toggles
	txtlen = rl.MeasureText("grid", txts)
	y := int32(infobar.rec.Y + 5)
	x -= float32(txtlen * 2)

	rec = toggle("grid", int32(x), y, settings.gridon)
	if rl.CheckCollisionPointRec(mousev2, rec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if clickpause == 0 {
				if settings.gridon {
					settings.gridon = false
					clickpause = fps / 3
				} else {
					settings.gridon = true
					clickpause = fps / 3
				}
			}
		}
	}
	y += txts * 2

	rec = toggle("snap", int32(x), y, settings.snapon)
	if rl.CheckCollisionPointRec(mousev2, rec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if clickpause == 0 {
				if settings.snapon {
					settings.snapon = false
					clickpause = fps / 3
				} else {
					settings.snapon = true
					clickpause = fps / 3
				}
			}
		}
	}

	txtlen = rl.MeasureText("outlines", txts)
	y = int32(infobar.rec.Y + 5)
	x -= float32(txtlen + txts)

	rec = toggle("outlines", int32(x), y, settings.outlineson)
	if rl.CheckCollisionPointRec(mousev2, rec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if clickpause == 0 {
				if settings.outlineson {
					settings.outlineson = false
					clickpause = fps / 3
				} else {
					settings.outlineson = true
					clickpause = fps / 3
				}
			}
		}
	}

	y += txts * 2

	rec = toggle("ruler", int32(x), y, settings.ruleron)
	if rl.CheckCollisionPointRec(mousev2, rec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if clickpause == 0 {
				if settings.ruleron {
					settings.ruleron = false
					clickpause = fps / 3
				} else {
					settings.ruleron = true
					clickpause = fps / 3
				}
			}
		}
	}

	txtlen = rl.MeasureText("colors", txts)
	y = int32(infobar.rec.Y + 5)
	x -= float32(txtlen*2 + (txts / 2))

	rec = toggle("colors", int32(x), y, settings.colorpallockon)
	if rl.CheckCollisionPointRec(mousev2, rec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if clickpause == 0 {
				if settings.colorpallockon {
					settings.colorpallockon = false
					colorpalon = false
					clickpause = fps / 3
				} else {
					settings.colorpallockon = true
					colorpalon = true
					clickpause = fps / 3
				}
			}
		}
	}
	y += txts * 2
	rec = toggle("paths", int32(x), y, settings.showpathson)
	if rl.CheckCollisionPointRec(mousev2, rec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if clickpause == 0 {
				if settings.showpathson {
					settings.showpathson = false
					clickpause = fps / 3
				} else {
					settings.showpathson = true
					clickpause = fps / 3
				}
			}
		}
	}

	txtlen = rl.MeasureText("animate", txts)
	y = int32(infobar.rec.Y + 5)
	x -= float32(txtlen + ((txts * 2) + txts/2))

	rec = toggle("animate", int32(x), y, settings.animate)
	if rl.CheckCollisionPointRec(mousev2, rec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if clickpause == 0 {
				if settings.animate {
					settings.animate = false
					clickpause = fps / 3
				} else {
					settings.animate = true
					clickpause = fps / 3
				}
			}
		}
	}

	if menu.lr {
		//left menu icon
		x = menu.rec.X + float32(txts/2)
		y := menu.rec.Y
		y -= larrowimg.Height
		y -= float32(txts)
		y += menu.rec.Height
		v2 = rl.NewVector2(x, y)
		rec = rl.NewRectangle(v2.X, v2.Y, larrowimg.Width, larrowimg.Height)
		if rl.CheckCollisionPointRec(mousev2, rec) {
			rl.DrawTextureRec(imgs, larrowimg, v2, brightred())
			helptxt("larrow")
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				menu.lr = false
				upmenu()
			}
		} else {
			rl.DrawTextureRec(imgs, larrowimg, v2, rl.White)
		}
	} else {
		//right menu icon
		x = menu.rec.Width - float32(txts/2)
		x -= rarrowimg.Width
		y := menu.rec.Y
		y -= rarrowimg.Height
		y -= float32(txts)
		y += menu.rec.Height
		v2 = rl.NewVector2(x, y)
		rec = rl.NewRectangle(v2.X, v2.Y, rarrowimg.Width, rarrowimg.Height)
		if rl.CheckCollisionPointRec(mousev2, rec) {
			rl.DrawTextureRec(imgs, rarrowimg, v2, brightred())
			helptxt("rarrow")
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				menu.lr = true
				upmenu()
			}
		} else {
			rl.DrawTextureRec(imgs, rarrowimg, v2, rl.White)
		}
	}

}
func mainmenuitem(itemname string, x, y int32, onoff bool) (x2, y2 int32, rec2 rl.Rectangle) { //MARK: mainmenuitem

	txtlen := rl.MeasureText(itemname, txts)
	rec := rl.NewRectangle(0, 0, 0, 0)
	if menu.lr {
		if x == int32(menu.rec.Width) {
			x = menu.rec.ToInt32().X
		}
		x += txts / 2
		if x+(txtlen+(txts*2)) > int32(scrwf32) {
			x = menu.rec.ToInt32().X
			x += txts / 2
			y += txts * 2
			y += txts / 2
		}

		rec = rl.NewRectangle(float32(x), float32(y-(txts/2)), float32(txtlen+(txts*2)), float32(txts*2))
		if rl.CheckCollisionPointRec(mousev2, rec) || onoff {
			rl.DrawRectangleRec(rec, brightred())
			rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
			if rl.CheckCollisionPointRec(mousev2, rec) {
				helptxt(itemname)
			}
		} else {
			rl.DrawRectangleRec(rec, rl.Black)
		}
		x += txts
		rl.DrawText(itemname, x, y, txts, rl.White)
		x += txtlen + txts
	} else {
		if x == int32(menu.rec.Width) {
			x -= txtlen + txts
			x -= txts / 2
		} else {
			x -= txtlen + txts
			x -= txts
		}
		if x-(txts) <= 0 {
			x = int32(menu.rec.Width)
			x -= txtlen + txts
			x -= txts / 2
			y += txts * 2
			y += txts / 2
		}
		rec = rl.NewRectangle(float32(x-txts), float32(y-(txts/2)), float32(txtlen+(txts*2)), float32(txts*2))
		if rl.CheckCollisionPointRec(mousev2, rec) || onoff {
			rl.DrawRectangleRec(rec, brightred())
			rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
			if rl.CheckCollisionPointRec(mousev2, rec) {
				helptxt(itemname)
			}
		} else {
			rl.DrawRectangleRec(rec, rl.Black)
		}
		rl.DrawText(itemname, x, y, txts, rl.White)

		x -= 4
	}

	return x, y, rec

}
func toggle(name string, x, y int32, onoff bool) (rec2 rl.Rectangle) { //MARK: toggle

	txtlen := rl.MeasureText(name, txts)
	x -= txtlen

	rec := rl.NewRectangle(float32(x), float32(y), float32(txtlen+txts)+togglewid, togglewid)
	if rl.CheckCollisionPointRec(mousev2, rec) {
		rl.DrawText(name, x, y, txts, brightred())
	} else {
		rl.DrawText(name, x, y, txts, rl.White)
	}

	rl.DrawRectangle(x+txtlen+txts, y, int32(togglewid), int32(togglewid), rl.Black)
	rl.DrawRectangleLines(x+txtlen+txts, y, int32(togglewid), int32(togglewid), rl.White)
	if onoff {
		rl.DrawRectangle(x+txtlen+txts, y, int32(togglewid), int32(togglewid), rl.White)
		rl.DrawRectangle(x+txtlen+txts+2, y+2, int32(togglewid)-4, int32(togglewid)-4, brightred())
	}

	return rec

}
func upmainmenuinfo(name string) { //MARK: upmainmenuinfo

	switch name {

	case "uiobjmenuon":
		uiobjmenuon = true
		objmenuon = false
	case "objmenuon":
		objmenuon = true
		uiobjmenuon = false

	}

}
func buttonmainmenu(name string, x, y float32, onoff bool) (x2, y2 float32, rec2 rl.Rectangle) { //MARK: buttonmainmenu

	txtlen := rl.MeasureText(name, txts)

	rec := rl.NewRectangle(x, y, float32(txtlen+(txts*2)), float32(txts*2))

	if rl.CheckCollisionPointRec(mousev2, rec) || onoff {
		rl.DrawRectangleRec(rec, brightred())
		rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
		if rl.CheckCollisionPointRec(mousev2, rec) {
			helptxt(name)
		}
	} else {
		rl.DrawRectangleRec(rec, rl.Black)
	}
	rl.DrawText(name, int32(x)+txts, int32(y)+txts/2, txts, rl.White)

	x += float32(txtlen + (txts * 2) + txts/2)

	if x >= scrwf32-100 {
		x = 10
		y += float32(txts * 4)
	}

	return x, y, rec
}
func buttonline(name string, x, y int32) rl.Rectangle { //MARK: buttonline

	txtlen := rl.MeasureText(name, txts)
	txtlen += txts * 2
	rec := rl.NewRectangle(float32(x), float32(y), float32(txtlen), float32(txts*2))
	if rl.CheckCollisionPointRec(mousev2, rec) {
		rl.DrawRectangleRec(rec, brightred())
	}
	rl.DrawRectangleLines(x, y, txtlen, txts*2, rl.White)
	rl.DrawText(name, x+txts, y+(txts/2), txts, rl.White)

	return rec

}
func buttonorange(name string, x, y int32) rl.Rectangle { //MARK: buttonorange

	txtlen := rl.MeasureText(name, txts)
	txtlen += txts * 2

	rec := rl.NewRectangle(float32(x), float32(y), float32(txtlen), float32(txts*2))
	if rl.CheckCollisionPointRec(mousev2, rec) {
		rl.DrawRectangleRec(rec, brightred())
		rl.DrawRectangleLinesEx(rec, 1.0, rl.White)
		rl.DrawText(name, x+txts, y+(txts/2), txts, rl.White)
	} else {
		rl.DrawRectangleRec(rec, brightorange())
		rl.DrawRectangleLinesEx(rec, 1.0, rl.Black)
		rl.DrawText(name, x+txts, y+(txts/2), txts, rl.Black)
	}

	return rec
}
func switchtilemenu(num int) { //MARK: switchtilemenu

	for a := 0; a < len(tilemenuonoff); a++ {
		tilemenuonoff[a].onoff = false
	}
	tilemenuonoff[num].onoff = true
}
func tilemenuoff() { //MARK: tilemenuoff

	tileselecton = false
	camtileselect.Target.Y = 0
	menufocus = false

}
func tickbox(x, y int32, switchvalue bool) bool { //MARK: tickbox

	rec := rl.NewRectangle(float32(x-2), float32(y-(txts/2)), tickimg.Width+4, tickimg.Height+4)
	v2 := rl.NewVector2(float32(x), rec.Y+2)

	if switchvalue {
		rl.DrawRectangleRec(rec, brightred())
		if rl.CheckCollisionPointRec(mousev2, rec) {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				switchvalue = false
			}
		}
	} else {
		rl.DrawRectangleRec(rec, brightred())
		if rl.CheckCollisionPointRec(mousev2, rec) {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				switchvalue = true
			}
		} else {
			rl.DrawRectangleRec(rec, brightorange())
		}
	}

	rl.DrawTextureRec(imgs, tickimg, v2, rl.White)

	return switchvalue
}
func numbox(x, y int32, min, max, numtochange int) (int, int32) { //MARK: numbox

	leftrec := rl.NewRectangle(float32(x), float32(y-(txts/2)), float32(txts*2), float32(txts*2))
	midrec := rl.NewRectangle(leftrec.X+leftrec.Width, leftrec.Y, leftrec.Width*2, leftrec.Height)
	rightrec := rl.NewRectangle(midrec.X+midrec.Width, leftrec.Y, leftrec.Width, leftrec.Height)

	if rl.CheckCollisionPointRec(mousev2, leftrec) {
		rl.DrawRectangleRec(leftrec, brightred())
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			numtochange--
			if numtochange < min {
				numtochange = max
			}
		}
	} else {
		rl.DrawRectangleRec(leftrec, brightorange())
	}

	rl.DrawRectangleRec(midrec, rl.Black)

	if rl.CheckCollisionPointRec(mousev2, rightrec) {
		rl.DrawRectangleRec(rightrec, brightred())
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			numtochange++
			if numtochange > max {
				numtochange = min
			}
		}
	} else {
		rl.DrawRectangleRec(rightrec, brightorange())
	}

	txtlen := rl.MeasureText(" - ", txtm)
	rl.DrawText(" - ", leftrec.ToInt32().X+(leftrec.ToInt32().Width/2)-(txtlen/2)-1, y-(txts/2)+1, txtm, rl.Black)
	rl.DrawText(" - ", leftrec.ToInt32().X+(leftrec.ToInt32().Width/2)-(txtlen/2), y-(txts/2), txtm, rl.White)

	txtlen = rl.MeasureText(" + ", txtm)
	rl.DrawText(" + ", rightrec.ToInt32().X+(rightrec.ToInt32().Width/2)-(txtlen/2)-1, y-(txts/2)+1, txtm, rl.Black)
	rl.DrawText(" + ", rightrec.ToInt32().X+(rightrec.ToInt32().Width/2)-(txtlen/2), y-(txts/2), txtm, rl.White)

	numtxt := strconv.Itoa(numtochange)
	txtlen = rl.MeasureText(numtxt, txts)
	rl.DrawText(numtxt, midrec.ToInt32().X+(midrec.ToInt32().Width/2)-txtlen/2, y, txts, rl.White)

	return numtochange, int32(rightrec.ToInt32().X + (rightrec.ToInt32().Width))
}

//MARK: TEXT TEXT TEXT TEXT TEXT TEXT TEXT TEXT TEXT TEXT TEXT TEXT TEXT TEXT TEXT TEXT TEXT

func returntxt() { //MARK: returntxt

	switch gettxtname {
	case "hp_activobj":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].hp = int(value)
	case "labelevent_labelchangehpobj2":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].labelchangehpobj2 = int(value)
	case "labelevent_labelchangeposyobj2":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].labelchangeposyobj2 = float32(value)
	case "labelevent_labelchangeposxobj2":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].labelchangeposxobj2 = float32(value)
	case "labelevent_changeshadowfadeobj2":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		if value < 10 {
			value = value / 10
		}
		if value > 9 {
			value = 1.0
		}
		objs[activobjnum].labelchangeshadowfadeobj2 = float32(value)
		if objs[activobjnum].complex {
			for a := 0; a < len(objs[activobjnum].objsin); a++ {
				objs[activobjnum].objsin[a].labelchangeshadowfadeobj2 = objs[activobjnum].labelchangeshadowfadeobj2
			}
		}
	case "labelevent_changeshadowxobj2":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].labelchangeshadowxobj2 = float32(value)
		if objs[activobjnum].complex {
			for a := 0; a < len(objs[activobjnum].objsin); a++ {
				objs[activobjnum].objsin[a].labelchangeshadowxobj2 = objs[activobjnum].labelchangeshadowxobj2
			}
		}
	case "labelevent_changeshadowyobj2":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].labelchangeshadowyobj2 = float32(value)
		if objs[activobjnum].complex {
			for a := 0; a < len(objs[activobjnum].objsin); a++ {
				objs[activobjnum].objsin[a].labelchangeshadowyobj2 = objs[activobjnum].labelchangeshadowyobj2
			}
		}
	case "labelevent_changeghostfadeobj2":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		if value < 10 {
			value = value / 10
		}
		if value > 9 {
			value = 1.0
		}
		objs[activobjnum].labelchangeghostfadeobj2 = float32(value)
		if objs[activobjnum].complex {
			for a := 0; a < len(objs[activobjnum].objsin); a++ {
				objs[activobjnum].objsin[a].labelchangeghostfadeobj2 = objs[activobjnum].labelchangeghostfadeobj2
			}
		}
	case "labelevent_changeghostyobj2":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].labelchangeghostyobj2 = float32(value)
		if objs[activobjnum].complex {
			for a := 0; a < len(objs[activobjnum].objsin); a++ {
				objs[activobjnum].objsin[a].labelchangeghostyobj2 = objs[activobjnum].labelchangeghostyobj2
			}
		}
	case "labelevent_changeghostxobj2":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].labelchangeghostxobj2 = float32(value)
		if objs[activobjnum].complex {
			for a := 0; a < len(objs[activobjnum].objsin); a++ {
				objs[activobjnum].objsin[a].labelchangeghostxobj2 = objs[activobjnum].labelchangeghostxobj2
			}
		}
	case "labelevent_changeheightobj2":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].labelchangeheightobj2 = float32(value)
	case "labelevent_changewidthobj2":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].labelchangewidthobj2 = float32(value)
	case "labelevent_changerotationobj2":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].labelchangerotationobj2 = float32(value)
	case "labelevent_changedirecyobj2":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].labelchangedirecyobj2 = float32(value)
	case "labelevent_changedirecxobj2":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].labelchangedirecxobj2 = float32(value)
	case "labelevent_labelchangehp":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].labelchangehp = int(value)
	case "labelevent_labelchangeposx":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].labelchangeposx = float32(value)
	case "labelevent_labelchangeposy":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].labelchangeposy = float32(value)
	case "labelevent_changeshadowy":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].labelchangeshadowy = float32(value)
		if objs[activobjnum].complex {
			for a := 0; a < len(objs[activobjnum].objsin); a++ {
				objs[activobjnum].objsin[a].labelchangeshadowy = objs[activobjnum].labelchangeshadowy
			}
		}
	case "labelevent_changeshadowx":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].labelchangeshadowx = float32(value)
		if objs[activobjnum].complex {
			for a := 0; a < len(objs[activobjnum].objsin); a++ {
				objs[activobjnum].objsin[a].labelchangeshadowx = objs[activobjnum].labelchangeshadowx
			}
		}

	case "labelevent_changeshadowfade":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		if value < 10 {
			value = value / 10
		}
		if value > 9 {
			value = 1.0
		}
		objs[activobjnum].labelchangeshadowfade = float32(value)
		if objs[activobjnum].complex {
			for a := 0; a < len(objs[activobjnum].objsin); a++ {
				objs[activobjnum].objsin[a].labelchangeshadowfade = objs[activobjnum].labelchangeshadowfade
			}
		}
	case "labelevent_changeghosty":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].labelchangeghosty = float32(value)
		if objs[activobjnum].complex {
			for a := 0; a < len(objs[activobjnum].objsin); a++ {
				objs[activobjnum].objsin[a].labelchangeghosty = objs[activobjnum].labelchangeghosty
			}
		}
	case "labelevent_changeghostx":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].labelchangeghostx = float32(value)
		if objs[activobjnum].complex {
			for a := 0; a < len(objs[activobjnum].objsin); a++ {
				objs[activobjnum].objsin[a].labelchangeghostx = objs[activobjnum].labelchangeghostx
			}
		}

	case "labelevent_changeghostfade":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		if value < 10 {
			value = value / 10
		}
		if value > 9 {
			value = 1.0
		}
		objs[activobjnum].labelchangeghostfade = float32(value)
		if objs[activobjnum].complex {
			for a := 0; a < len(objs[activobjnum].objsin); a++ {
				objs[activobjnum].objsin[a].labelchangeghostfade = objs[activobjnum].labelchangeghostfade
			}
		}
	case "labelevent_changeheight":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].labelchangeheight = float32(value)
	case "labelevent_changewidth":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].labelchangewidth = float32(value)
	case "labelevent_changerotation":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].labelchangerotation = float32(value)
	case "labelevent_changedirecy":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].labelchangedirecy = float32(value)
	case "labelevent_changedirecx":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].labelchangedirecx = float32(value)
	case "labelevent_rename":
		if len(objs[activobjnum].labelevents) > 0 {
			samename := false
			for {
				for _, checklabelevent := range objs[activobjnum].labelevents {
					if checklabelevent.name == gettxtreturn {
						gettxtreturn = gettxtreturn + strconv.Itoa(rInt(100, 1000))
						samename = true
					}
				}
				if !samename {
					break
				}
			}
		}
		deflabeleventnamechanged = gettxtreturn
	case "labelall_rename":
		for i, checkobj := range objs {
			if len(checkobj.labels) > 0 {
				for a := 0; a < len(checkobj.labels); a++ {
					if checkobj.labels[a] == labelsalllist[editlabelnum] {
						objs[i].labels[a] = gettxtreturn
					}
				}
			}
		}
		labelsalllist[editlabelnum] = gettxtreturn
	case "obj_label_rename":
		oldlabelname := objs[activobjnum].labels[editlabelnum]
		for i, labelname := range labelsalllist {
			if labelname == oldlabelname {
				labelsalllist[i] = gettxtreturn
			}
		}
		objs[activobjnum].labels[editlabelnum] = gettxtreturn
	case "obj_label":
		objs[activobjnum].labels = append(objs[activobjnum].labels, gettxtreturn)
		samelabel := true
		for _, labelname := range labelsalllist {
			if labelname == gettxtreturn {
				samelabel = false
			}
		}
		if samelabel {
			labelsalllist = append(labelsalllist, gettxtreturn)
		}
		addinglabel = false
	case "rotate_speed":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		if value < 0 {
			value = math.Abs(value)
		}

		objs[activobjnum].rotate_speed = float32(value)
		objs[activobjnum].orig_rotatespeed = objs[activobjnum].rotate_speed
	case "rotate_timer":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}

		objs[activobjnum].rotate_timer = float32(value) * float32(fps)
		objs[activobjnum].orig_rotatetimer = objs[activobjnum].rotate_timer
	case "obj_rotation":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}

		objs[activobjnum].rotation = float32(value)
	case "rand_direc_timer_min":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		if objs[activobjnum].rand_direc_timer_min > objs[activobjnum].rand_direc_timer_max {
			objs[activobjnum].rand_direc_timer_min = objs[activobjnum].rand_direc_timer_max - 1
		}
		if objs[activobjnum].rand_direc_timer_min < 0 {
			objs[activobjnum].rand_direc_timer_min = 0
		}
		objs[activobjnum].rand_direc_timer_min = float32(value)
	case "rand_direc_timer_max":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		if objs[activobjnum].rand_direc_timer_min > objs[activobjnum].rand_direc_timer_max {
			objs[activobjnum].rand_direc_timer_max = objs[activobjnum].rand_direc_timer_min + 1
		}
		if objs[activobjnum].rand_direc_timer_max < 0 {
			objs[activobjnum].rand_direc_timer_max = 0
		}
		objs[activobjnum].rand_direc_timer_max = float32(value)
	case "rand_direc_max_y":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].rand_direc_max_x = float32(value)
	case "rand_direc_max_x":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].rand_direc_max_x = float32(value)
	case "addscreen_width":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		if value == 0 {
			break
		}
		if value > 99 {
			value = 99
		}

		if settings.screen_width_multiplier+float32(value) > 100 {
			value = float64((settings.screen_width_multiplier + float32(value)) - 100)
		}
		switch addscreendirec {
		case 2:
			settings.level_x_right += float32(value) * scrwf32
			settings.border_right = settings.level_x_right
		case 4:
			settings.level_x_left -= float32(value) * scrwf32
			settings.border_left = settings.level_x_left
		}
		settings.screen_width_multiplier += float32(value)
		if settings.screen_width_multiplier > 100 {
			settings.screen_width_multiplier = 100
		}
	case "addscreen_height":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		if value == 0 {
			break
		}
		if value > 99 {
			value = 99
		}

		if settings.screen_height_multiplier+float32(value) > 100 {
			value = float64((settings.screen_height_multiplier + float32(value)) - 100)
		}

		switch addscreendirec {
		case 1:
			settings.level_y_top -= float32(value) * scrhf32
			settings.border_top = settings.level_y_top
		case 3:
			settings.level_y_bottom += float32(value) * scrhf32
			settings.border_bottom = settings.level_y_bottom
		}

		settings.screen_height_multiplier += float32(value)
		if settings.screen_height_multiplier > 100 {
			settings.screen_height_multiplier = 100
		}

	case "shadow_fade":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		if value < 10 {
			value = value / 10
		}
		if value > 9 {
			value = 1.0
		}
		objs[activobjnum].shadow_fade = float32(value)
		if objs[activobjnum].complex {
			for a := 0; a < len(objs[activobjnum].objsin); a++ {
				objs[activobjnum].objsin[a].shadow_fade = objs[activobjnum].shadow_fade
			}
		}
	case "shadow_x":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}

		objs[activobjnum].shadow_x = float32(value)

		if objs[activobjnum].complex {
			for a := 0; a < len(objs[activobjnum].objsin); a++ {
				objs[activobjnum].objsin[a].shadow_x = objs[activobjnum].shadow_x
			}
		}
	case "shadow_y":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}

		objs[activobjnum].shadow_y = float32(value)
		if objs[activobjnum].complex {
			for a := 0; a < len(objs[activobjnum].objsin); a++ {
				objs[activobjnum].objsin[a].shadow_y = objs[activobjnum].shadow_y
			}
		}
	case "fade":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		if value < 10 {
			value = value / 10
		}
		if value > 9 {
			value = 1.0
		}
		objs[activobjnum].fade = float32(value)
		if objs[activobjnum].complex {
			for a := 0; a < len(objs[activobjnum].objsin); a++ {
				objs[activobjnum].objsin[a].fade = objs[activobjnum].fade
			}
		}

	case "uiobj_shadow_fade":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		if value < 10 {
			value = value / 10
		}
		if value > 9 {
			value = 1.0
		}
		uiobjs[activuiobjnum].shadow_fade = float32(value)
	case "uiobj_shadow_x":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}

		uiobjs[activuiobjnum].shadow_x = float32(value)
	case "uiobj_shadow_y":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}

		uiobjs[activuiobjnum].shadow_y = float32(value)

	case "uiobj_center_x":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}

		uiobjs[activuiobjnum].circv2.X = float32(value)
		uiobjs[activuiobjnum].rec.X = uiobjs[activuiobjnum].circv2.X - (uiobjs[activuiobjnum].rec.Width / 2)
	case "uiobj_center_y":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}

		uiobjs[activuiobjnum].circv2.Y = float32(value)
		uiobjs[activuiobjnum].rec.Y = uiobjs[activuiobjnum].circv2.Y - (uiobjs[activuiobjnum].rec.Height / 2)

	case "uiobj_topleft_y":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}

		uiobjs[activuiobjnum].rec.Y = float32(value)
		uiobjs[activuiobjnum].circv2.Y = uiobjs[activuiobjnum].rec.Y + (uiobjs[activuiobjnum].rec.Height / 2)
	case "uiobj_topleft_x":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}

		uiobjs[activuiobjnum].rec.X = float32(value)
		uiobjs[activuiobjnum].circv2.X = uiobjs[activuiobjnum].rec.X + (uiobjs[activuiobjnum].rec.Width / 2)
	case "uiobj_polygon_circ_radius":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		if value < 1 {
			value = 1
		}
		uiobjs[activuiobjnum].circrad = float32(value)
		uiobjs[activuiobjnum].rec.Width = float32(value * 2)
	case "uiobj_rech":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		if value < 2 {
			value = 2
		}
		uiobjs[activuiobjnum].rec.Height = float32(value)

	case "uiobj_recw":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		if value < 2 {
			value = 2
		}
		uiobjs[activuiobjnum].rec.Width = float32(value)
		uiobjs[activuiobjnum].circrad = uiobjs[activuiobjnum].rec.Width / 2

	case "uiobj_rotation":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		for {
			if value > 360 {
				value = value - 360
			} else {
				break
			}
		}

		uiobjs[activuiobjnum].rotation = float32(value)

		if uiobjs[activuiobjnum].complex {
			for a := 0; a < len(uiobjs[activuiobjnum].uiobjsin); a++ {
				uiobjs[activuiobjnum].uiobjsin[a].rotation = uiobjs[activuiobjnum].rotation
			}
		}

	case "uiobjname":
		uiobjs[activuiobjnum].name = gettxtreturn

	case "ghosting_fade":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		if value < 10 {
			value = value / 10
		}
		if value > 9 {
			value = 1.0
		}
		objs[activobjnum].ghosting_fade = float32(value)

		if objs[activobjnum].complex {
			for a := 0; a < len(objs[activobjnum].objsin); a++ {
				objs[activobjnum].objsin[a].ghosting_fade = objs[activobjnum].ghosting_fade
			}
		}
	case "ghosting_y":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		if value < 1 {
			value = 1
		}
		objs[activobjnum].ghosting_y = float32(value)

		if objs[activobjnum].complex {
			for a := 0; a < len(objs[activobjnum].objsin); a++ {
				objs[activobjnum].objsin[a].ghosting_y = objs[activobjnum].ghosting_y
			}
		}
	case "ghosting_x":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		if value < 1 {
			value = 1
		}
		objs[activobjnum].ghosting_x = float32(value)

		if objs[activobjnum].complex {
			for a := 0; a < len(objs[activobjnum].objsin); a++ {
				objs[activobjnum].objsin[a].ghosting_x = objs[activobjnum].ghosting_x
			}
		}

	case "tile_h":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		if value < 2 {
			value = 2
		}
		objs[activobjnum].tile_h = float32(value)

		newheig := objs[activobjnum].tile_h * float32(objs[activobjnum].tilenumh)

		objs[activobjnum].rec.Height = newheig

		for a := 0; a < len(objs[activobjnum].objsin); a++ {
			objs[activobjnum].objsin[a].tile_h = objs[activobjnum].tile_h
			objs[activobjnum].objsin[a].rec.Height = objs[activobjnum].tile_h
		}

		x := objs[activobjnum].rec.X
		y := objs[activobjnum].rec.Y
		for a := 0; a < len(objs[activobjnum].objsin); a++ {
			objs[activobjnum].objsin[a].rec.Y = y

			if objs[activobjnum].tile_w == 0 {
				x += settings.tilew
			} else {
				x += objs[activobjnum].tile_w
			}

			if x >= objs[activobjnum].rec.X+objs[activobjnum].rec.Width {
				x = objs[activobjnum].rec.X
				if objs[activobjnum].tile_h == 0 {
					y += settings.tileh
				} else {
					y += objs[activobjnum].tile_h
				}
			}
		}
	case "tile_w":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		if value < 2 {
			value = 2
		}
		objs[activobjnum].tile_w = float32(value)

		newwid := objs[activobjnum].tile_w * float32(objs[activobjnum].tilenumw)

		objs[activobjnum].rec.Width = newwid

		for a := 0; a < len(objs[activobjnum].objsin); a++ {
			objs[activobjnum].objsin[a].tile_w = objs[activobjnum].tile_w
			objs[activobjnum].objsin[a].rec.Width = objs[activobjnum].tile_w
		}

		x := objs[activobjnum].rec.X
		y := objs[activobjnum].rec.Y
		for a := 0; a < len(objs[activobjnum].objsin); a++ {
			objs[activobjnum].objsin[a].rec.X = x

			if objs[activobjnum].tile_w == 0 {
				x += settings.tilew
			} else {
				x += objs[activobjnum].tile_w
			}

			if x >= objs[activobjnum].rec.X+objs[activobjnum].rec.Width {
				x = objs[activobjnum].rec.X
				if objs[activobjnum].tile_h == 0 {
					y += settings.tileh
				} else {
					y += objs[activobjnum].tile_h
				}
			}
		}

	case "outline_w":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		if value < 0.5 {
			value = 0.5
		}
		objs[activobjnum].outline_w = float32(value)
		if objs[activobjnum].complex {
			for a := 0; a < len(objs[activobjnum].objsin); a++ {
				objs[activobjnum].objsin[a].outline_w = objs[activobjnum].outline_w
			}
		}
	case "direction_x":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].direction_x = float32(value)
		objs[activobjnum].orig_direcx = objs[activobjnum].direction_x
	case "direction_y":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}

		objs[activobjnum].direction_y = float32(value)
		objs[activobjnum].orig_direcy = objs[activobjnum].direction_y
	case "img_rotate_speed":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].img_rotate_speed = float32(value)
	case "img_rotation":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].img_rotation = float32(value)
		if objs[activobjnum].complex {
			for a := 0; a < len(objs[activobjnum].objsin); a++ {
				objs[activobjnum].objsin[a].img_rotation = objs[activobjnum].img_rotation
			}
		}

	case "obj_w":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		if value < 1 {
			value = 1
		}
		objs[activobjnum].rec.Width = float32(value)
	case "obj_h":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		if value < 1 {
			value = 1
		}
		objs[activobjnum].rec.Height = float32(value)
	case "topleft_y":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].rec.Y = float32(value)
		if objs[activobjnum].complex {
			x := objs[activobjnum].rec.X
			y := objs[activobjnum].rec.Y
			for a := 0; a < len(objs[activobjnum].objsin); a++ {

				objs[activobjnum].objsin[a].rec.Y = y
				if objs[activobjnum].tile_w == 0 {
					x += settings.tilew
				} else {
					x += objs[activobjnum].tile_w
				}
				if x >= objs[activobjnum].rec.X+objs[activobjnum].rec.Width {
					x = objs[activobjnum].rec.X
					if objs[activobjnum].tile_h == 0 {
						y += settings.tileh
					} else {
						y += objs[activobjnum].tile_h
					}
				}
			}
		}
	case "topleft_x":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		objs[activobjnum].rec.X = float32(value)
		if objs[activobjnum].complex {
			x := objs[activobjnum].rec.X
			y := objs[activobjnum].rec.Y
			for a := 0; a < len(objs[activobjnum].objsin); a++ {
				objs[activobjnum].objsin[a].rec.X = x

				if objs[activobjnum].tile_w == 0 {
					x += settings.tilew
				} else {
					x += objs[activobjnum].tile_w
				}
				if x >= objs[activobjnum].rec.X+objs[activobjnum].rec.Width {
					x = objs[activobjnum].rec.X
					if objs[activobjnum].tile_h == 0 {
						y += settings.tileh
					} else {
						y += objs[activobjnum].tile_h
					}
				}
			}
		}
	case "objname":
		objs[activobjnum].name = gettxtreturn
	case "border_top":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		settings.border_top = float32(value)
	case "border_left":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		settings.border_left = float32(value)
	case "border_right":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		settings.border_right = float32(value)
	case "border_bottom":
		value, err := strconv.ParseFloat(gettxtreturn, 32)
		if err != nil {

		}
		settings.border_bottom = float32(value)

	}

	cleargettxt()
	menufocus = false
	gettxton = false
}

func gettxt() { //MARK: gettxt
	menufocus = true
	gettxtreturn = ""
	//back rec
	rl.DrawRectangle(0, 0, scrw, scrh, rl.Fade(rl.Black, 0.6))
	rec := rl.NewRectangle(cntrscr.X-400, cntrscr.Y-200, 800, 400)
	rl.DrawRectangleRec(rec, rl.Fade(darkred(), 0.7))

	//close win
	v2 := rl.NewVector2(rec.X+rec.Width-(closewinimg.Width*2), rec.Y+closewinimg.Height)
	rec2 := rl.NewRectangle(v2.X, v2.Y, closewinimg.Width, closewinimg.Height)
	if rl.CheckCollisionPointRec(mousev2, rec2) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			gettxton = false
			menufocus = false
		}
		rl.DrawTextureRec(imgs, closewinimg, v2, brightred())
	} else {
		rl.DrawTextureRec(imgs, closewinimg, v2, rl.White)
	}
	inprec := rl.NewRectangle(rec.X+50, rec.Y+rec.Height-200, rec.Width-100, 40)
	rl.DrawRectangleRec(inprec, rl.Fade(rl.Black, 0.7))

	if keybcount > 0 {
		v2 = rl.NewVector2(inprec.X+inprec.Width-(tickimg.Width*2), inprec.Y+14)
		tickrec := rl.NewRectangle(inprec.X+inprec.Width-(tickimg.Width*2), inprec.Y, 30, 40)
		if rl.CheckCollisionPointRec(mousev2, tickrec) {
			rl.DrawTextureRec(imgs, tickimg, v2, brightred())
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				for a := 0; a < len(keybinput); a++ {
					if keybinput[a] != "" {
						gettxtreturn = gettxtreturn + keybinput[a]
					}
				}
				returntxt()
			}
		} else {
			rl.DrawTextureRec(imgs, tickimg, v2, rl.White)
		}

		if rl.IsKeyPressed(rl.KeyEnter) || rl.IsKeyPressed(rl.KeyKpEnter) {
			for a := 0; a < len(keybinput); a++ {
				if keybinput[a] != "" {
					gettxtreturn = gettxtreturn + keybinput[a]
				}
			}
			returntxt()
		}
	}
	txtlen := rl.MeasureText("enter text below", txtm)

	rl.DrawText("enter text below", inprec.ToInt32().X+(inprec.ToInt32().Width/2)-(txtlen/2), inprec.ToInt32().Y-txtm*2, txtm, rl.White)

	//help text
	if gettxtname == "ghosting_fade" || gettxtname == "fade" || gettxtname == "shadow" {
		txtlen = rl.MeasureText("7 equals 0.7 fade and 77 equals 0.77 fade", txts)
		rl.DrawText("7 equals 0.7 fade and 77 equals 0.77 fade", inprec.ToInt32().X+(inprec.ToInt32().Width/2)-(txtlen/2), inprec.ToInt32().Y+txtm*3, txts, rl.White)
	}

	if gettxtname == "addscreen" {
		txtlen = rl.MeasureText("mutliplier for screens in this direction - max sreens = 100 X 100", txts)
		rl.DrawText("mutliplier for screens in this direction - max sreens = 100 X 100", inprec.ToInt32().X+(inprec.ToInt32().Width/2)-(txtlen/2), inprec.ToInt32().Y+txtm*3, txts, rl.White)

	}

	if gettxttype == 1 || gettxttype == 3 {
		if rl.IsKeyPressed(rl.KeyA) {
			keybinput[keybcount] = "a"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyB) {
			keybinput[keybcount] = "b"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyC) {
			keybinput[keybcount] = "c"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyD) {
			keybinput[keybcount] = "d"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyE) {
			keybinput[keybcount] = "e"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyF) {
			keybinput[keybcount] = "f"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyG) {
			keybinput[keybcount] = "g"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyH) {
			keybinput[keybcount] = "h"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyI) {
			keybinput[keybcount] = "i"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyJ) {
			keybinput[keybcount] = "j"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyK) {
			keybinput[keybcount] = "k"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyL) {
			keybinput[keybcount] = "l"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyM) {
			keybinput[keybcount] = "m"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyN) {
			keybinput[keybcount] = "n"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyO) {
			keybinput[keybcount] = "o"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyP) {
			keybinput[keybcount] = "p"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyQ) {
			keybinput[keybcount] = "q"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyR) {
			keybinput[keybcount] = "r"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyS) {
			keybinput[keybcount] = "s"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyT) {
			keybinput[keybcount] = "t"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyU) {
			keybinput[keybcount] = "u"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyV) {
			keybinput[keybcount] = "v"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyW) {
			keybinput[keybcount] = "w"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyX) {
			keybinput[keybcount] = "x"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyY) {
			keybinput[keybcount] = "y"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyZ) {
			keybinput[keybcount] = "z"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeySpace) {
			keybinput[keybcount] = " "
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyMinus) || rl.IsKeyPressed(rl.KeyKpSubtract) {
			keybinput[keybcount] = "-"
			keybcount++
		}
	}
	if gettxttype == 2 || gettxttype == 3 || gettxttype == 4 || gettxttype == 5 {
		if rl.IsKeyPressed(rl.KeyOne) || rl.IsKeyPressed(rl.KeyKp1) {
			keybinput[keybcount] = "1"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyTwo) || rl.IsKeyPressed(rl.KeyKp2) {
			keybinput[keybcount] = "2"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyThree) || rl.IsKeyPressed(rl.KeyKp3) {
			keybinput[keybcount] = "3"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyFour) || rl.IsKeyPressed(rl.KeyKp4) {
			keybinput[keybcount] = "4"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyFive) || rl.IsKeyPressed(rl.KeyKp5) {
			keybinput[keybcount] = "5"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeySix) || rl.IsKeyPressed(rl.KeyKp6) {
			keybinput[keybcount] = "6"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeySeven) || rl.IsKeyPressed(rl.KeyKp7) {
			keybinput[keybcount] = "7"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyEight) || rl.IsKeyPressed(rl.KeyKp8) {
			keybinput[keybcount] = "8"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyNine) || rl.IsKeyPressed(rl.KeyKp9) {
			keybinput[keybcount] = "9"
			keybcount++
		}
		if rl.IsKeyPressed(rl.KeyZero) || rl.IsKeyPressed(rl.KeyKp0) {
			keybinput[keybcount] = "0"
			keybcount++
		}
	}
	if gettxttype == 4 || gettxttype == 5 {
		if rl.IsKeyPressed(rl.KeyPeriod) || rl.IsKeyPressed(rl.KeyKpDecimal) {
			keybinput[keybcount] = "."
			keybcount++
		}
	}
	if gettxttype == 5 {
		if rl.IsKeyPressed(rl.KeyMinus) || rl.IsKeyPressed(rl.KeyKpSubtract) {
			keybinput[keybcount] = "-"
			keybcount++
		}
	}

	if rl.IsKeyPressed(rl.KeyBackspace) {

		if keybcount == gettxtcharlimit {
			keybinput[keybcount] = ""
		}
		if keybcount > 0 {
			keybcount--
			keybinput[keybcount] = ""
		}
		if keybinput[keybcount] != "" {
			keybinput[keybcount] = ""
		}
	}

	if keybcount > gettxtcharlimit {
		keybcount = gettxtcharlimit
		charlimiton = true
	} else if keybcount < gettxtcharlimit {
		charlimiton = false
	}
	if charlimiton {
		txtlen := rl.MeasureText("character limit", txtm)
		rl.DrawText("character limit", inprec.ToInt32().X+(inprec.ToInt32().Width/2)-(txtlen/2), inprec.ToInt32().Y+80, txtm, rl.White)
	}

	x := int32(inprec.X + 20)

	for a := 0; a < len(keybinput); a++ {
		if keybinput[a] != "" {
			rl.DrawText(keybinput[a], x, inprec.ToInt32().Y+10, txtm, rl.White)
			x += txtm - 6
			if keybinput[a] == "i" || keybinput[a] == "." {
				x -= 8
			}
			if keybinput[a] == "l" || keybinput[a] == "j" || keybinput[a] == "1" {
				x -= 6
			}
			if keybinput[a] == "f" || keybinput[a] == "t" {
				x -= 6
			}
		}
	}

}
func cleargettxt() { //MARK: cleargettxt
	for a := 0; a < len(keybinput); a++ {
		keybinput[a] = ""
	}
	keybcount = 0
	gettxtcharlimit = 0
	gettxtname = ""

}
func helptxt(name string) { //MARK: helptxt
	switch name {
	case "uiobj_rec_rotation":
		rl.DrawText("cannot rotate a rectangle - change shape to square polygon", menu.rec.ToInt32().X+menu.rec.ToInt32().Width+10, int32(mousev2.Y), txts, rl.White)
	case "uiobj_circ_rotation":
		rl.DrawText("cannot rotate a circle", menu.rec.ToInt32().X+menu.rec.ToInt32().Width+10, int32(mousev2.Y), txts, rl.White)
	case "clearcolor":
		rl.DrawText("clear color", menu.rec.ToInt32().X+menu.rec.ToInt32().Width+10, int32(mousev2.Y), txts, rl.White)
	case "refreshcolor":
		rl.DrawText("randomize colors", menu.rec.ToInt32().X+menu.rec.ToInt32().Width+10, int32(mousev2.Y), txts, rl.White)

	case "color":
		rtxt := strconv.Itoa(int(selcolor.r))
		gtxt := strconv.Itoa(int(selcolor.g))
		btxt := strconv.Itoa(int(selcolor.b))
		ftxt := strconv.Itoa(int(selcolor.fade))

		txtlen := rl.MeasureText(rtxt, txts)
		rl.DrawText(rtxt, menu.rec.ToInt32().X+menu.rec.ToInt32().Width+10, int32(mousev2.Y), txts, rl.White)
		rl.DrawText("red", menu.rec.ToInt32().X+menu.rec.ToInt32().Width+10+txtlen+txts, int32(mousev2.Y), txts, rl.White)
		txtlen = rl.MeasureText(gtxt, txts)
		rl.DrawText(gtxt, menu.rec.ToInt32().X+menu.rec.ToInt32().Width+10, int32(mousev2.Y)+txts, txts, rl.White)
		rl.DrawText("green", menu.rec.ToInt32().X+menu.rec.ToInt32().Width+10+txtlen+txts, int32(mousev2.Y)+txts, txts, rl.White)
		txtlen = rl.MeasureText(btxt, txts)
		rl.DrawText(btxt, menu.rec.ToInt32().X+menu.rec.ToInt32().Width+10, int32(mousev2.Y)+txts*2, txts, rl.White)
		rl.DrawText("blue", menu.rec.ToInt32().X+menu.rec.ToInt32().Width+10+txtlen+txts, int32(mousev2.Y)+txts*2, txts, rl.White)
		txtlen = rl.MeasureText(ftxt, txts)
		rl.DrawText(ftxt, menu.rec.ToInt32().X+menu.rec.ToInt32().Width+10, int32(mousev2.Y)+txts*3, txts, rl.White)
		rl.DrawText("fade", menu.rec.ToInt32().X+menu.rec.ToInt32().Width+10+txtlen+txts, int32(mousev2.Y)+txts*3, txts, rl.White)

	case "obj +":
		if menu.lr {

		} else {
			rl.DrawText("new object", int32(mousev2.X+float32(txts*4)), int32(mousev2.Y), txts, rl.White)
		}
	case "rarrow":
		rl.DrawText("menu right", int32(mousev2.X+float32(txts*4)), int32(mousev2.Y), txts, rl.White)
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
func txthere(txt string, cntr bool, size int32, x, y, wid float32) {

	x32 := int32(x)
	y32 := int32(y)
	wid32 := int32(wid)

	txtlen := rl.MeasureText(txt, size)

	if cntr {

		rl.DrawText(txt, (x32+(wid32/2))-txtlen/2, y32, size, rl.White)

	} else {
		rl.DrawText(txt, x32, y32, size, rl.White)
	}

}

//MARK: OBJ ACTIONS OBJ ACTIONS OBJ ACTIONS OBJ ACTIONS OBJ ACTIONS OBJ ACTIONS OBJ ACTIONS
func remobjs(s []obj, index int) []obj { //MARK: remobjs
	return append(s[:index], s[index+1:]...)
}
func remstring(s []string, index int) []string { //MARK: remobjs
	return append(s[:index], s[index+1:]...)
}
func copyobj(num int) { //MARK: copyobj

	newobj := objs[activobjnum]

	newobj.rec.X = startv2mouse.X
	newobj.rec.Y = startv2mouse.Y
	newobj.circv2.X = startv2mouse.X + newobj.circrad
	newobj.circv2.Y = startv2mouse.Y + newobj.circrad

	objs = append(objs, newobj)

	copyobjon = false
}

func addv2path() { //MARK: addv2path

	objs[activobjnum].path = append(objs[activobjnum].path, startv2mouseworld)

}

//MARK: MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE
func moveobjpath(objnum int) { //MARK: moveobjpath

	if objs[objnum].reversepath {
		if objs[objnum].currentpathv2 > 0 {
			if objs[objnum].circv2 != objs[objnum].path[objs[objnum].currentpathv2-1] {

				if objs[objnum].circv2.X > objs[objnum].path[objs[objnum].currentpathv2-1].X {
					objs[objnum].circv2.X -= objs[objnum].direction_x
				} else if objs[objnum].circv2.X < objs[objnum].path[objs[objnum].currentpathv2-1].X {
					objs[objnum].circv2.X += objs[objnum].direction_x
				}

				if objs[objnum].circv2.Y > objs[objnum].path[objs[objnum].currentpathv2-1].Y {
					objs[objnum].circv2.Y -= objs[objnum].direction_y
				} else if objs[objnum].circv2.Y < objs[objnum].path[objs[objnum].currentpathv2-1].Y {
					objs[objnum].circv2.Y += objs[objnum].direction_y
				}

				objs[objnum].rec.X = objs[objnum].circv2.X - (objs[objnum].rec.Width / 2)
				objs[objnum].rec.Y = objs[objnum].circv2.Y - (objs[objnum].rec.Height / 2)

				if objs[objnum].complex {
					x := objs[objnum].rec.X
					y := objs[objnum].rec.Y
					for a := 0; a < len(objs[objnum].objsin); a++ {
						objs[objnum].objsin[a].rec.X = x
						objs[objnum].objsin[a].rec.Y = y
						if objs[objnum].tile_w == 0 {
							x += settings.tilew
						} else {
							x += objs[objnum].tile_w
						}
						if x >= objs[objnum].rec.X+objs[objnum].rec.Width {
							x = objs[objnum].rec.X
							if objs[objnum].tile_h == 0 {
								y += settings.tileh
							} else {
								y += objs[objnum].tile_h
							}
						}
					}
				}

			} else {
				objs[objnum].currentpathv2--
			}

		} else {
			objs[objnum].reversepath = false
		}
	} else {
		if objs[objnum].currentpathv2 < len(objs[objnum].path)-1 {
			if objs[objnum].circv2 != objs[objnum].path[objs[objnum].currentpathv2+1] {

				if objs[objnum].circv2.X > objs[objnum].path[objs[objnum].currentpathv2+1].X {
					objs[objnum].circv2.X -= objs[objnum].direction_x
				} else if objs[objnum].circv2.X < objs[objnum].path[objs[objnum].currentpathv2+1].X {
					objs[objnum].circv2.X += objs[objnum].direction_x
				}

				if objs[objnum].circv2.Y > objs[objnum].path[objs[objnum].currentpathv2+1].Y {
					objs[objnum].circv2.Y -= objs[objnum].direction_y
				} else if objs[objnum].circv2.Y < objs[objnum].path[objs[objnum].currentpathv2+1].Y {
					objs[objnum].circv2.Y += objs[objnum].direction_y
				}

				objs[objnum].rec.X = objs[objnum].circv2.X - (objs[objnum].rec.Width / 2)
				objs[objnum].rec.Y = objs[objnum].circv2.Y - (objs[objnum].rec.Height / 2)

				if objs[objnum].complex {
					x := objs[objnum].rec.X
					y := objs[objnum].rec.Y
					for a := 0; a < len(objs[objnum].objsin); a++ {
						objs[objnum].objsin[a].rec.X = x
						objs[objnum].objsin[a].rec.Y = y
						if objs[objnum].tile_w == 0 {
							x += settings.tilew
						} else {
							x += objs[objnum].tile_w
						}
						if x >= objs[objnum].rec.X+objs[objnum].rec.Width {
							x = objs[objnum].rec.X
							if objs[objnum].tile_h == 0 {
								y += settings.tileh
							} else {
								y += objs[objnum].tile_h
							}
						}
					}
				}

			} else {
				objs[objnum].currentpathv2++
			}
		} else {
			objs[objnum].reversepath = true
		}
	}
}
func moveobj(objnum int) { //MARK: moveobj

	objobjcollisions := false
	for a := 0; a < len(objs[objnum].events); a++ {
		if objs[objnum].events[a].event_type == 1 {
			objobjcollisions = true
		}
	}

	if objobjcollisions {
		checkobjobjcollision(objnum)
	}

	objs[objnum].collisrec = rl.NewRectangle(objs[objnum].rec.X+(objs[objnum].direction_x), objs[objnum].rec.Y+(objs[objnum].direction_y), objs[objnum].rec.Width, objs[objnum].rec.Height)

	if objs[objnum].collisrec.X+objs[objnum].direction_x > settings.border_left && objs[objnum].collisrec.X+objs[objnum].rec.Width+objs[objnum].direction_x < settings.border_right {
		objs[objnum].rec.X += objs[objnum].direction_x
		objs[objnum].circv2.X += objs[objnum].direction_x
		if objs[objnum].complex {
			x := objs[objnum].rec.X
			y := objs[objnum].rec.Y
			for a := 0; a < len(objs[objnum].objsin); a++ {
				objs[objnum].objsin[a].rec.X = x
				if objs[objnum].tile_w == 0 {
					x += settings.tilew
				} else {
					x += objs[objnum].tile_w
				}
				if x >= objs[objnum].rec.X+objs[objnum].rec.Width {
					x = objs[objnum].rec.X
					if objs[objnum].tile_h == 0 {
						y += settings.tileh
					} else {
						y += objs[objnum].tile_h
					}
				}
			}
		}
	} else {
		checkobjboundarycollision(objnum)
	}

	if objs[objnum].collisrec.Y+objs[objnum].rec.Height+objs[objnum].direction_y < settings.border_bottom && objs[objnum].collisrec.Y+objs[objnum].direction_y > settings.border_top {
		objs[objnum].rec.Y += objs[objnum].direction_y
		objs[objnum].circv2.Y += objs[objnum].direction_y
		if objs[objnum].complex {
			x := objs[objnum].rec.X
			y := objs[objnum].rec.Y
			for a := 0; a < len(objs[objnum].objsin); a++ {
				objs[objnum].objsin[a].rec.Y = y
				if objs[objnum].tile_w == 0 {
					x += settings.tilew
				} else {
					x += objs[objnum].tile_w
				}
				if x >= objs[objnum].rec.X+objs[objnum].rec.Width {
					x = objs[objnum].rec.X
					if objs[objnum].tile_h == 0 {
						y += settings.tileh
					} else {
						y += objs[objnum].tile_h
					}
				}
			}
		}
	} else {
		checkobjboundarycollision(objnum)
	}

}

func checkobjobjcollision(objnum int) { //MARK: checkobjobjcollision

	objs[objnum].collisrec = rl.NewRectangle(objs[objnum].rec.X+(objs[objnum].direction_x*2), objs[objnum].rec.Y+(objs[objnum].direction_y*2), objs[objnum].rec.Width, objs[objnum].rec.Height)

	for num, checkobj := range objs {
		if rl.CheckCollisionRecs(objs[objnum].collisrec, checkobj.collisrec) && objnum != num {

			if len(objs[objnum].labelevents) > 0 {

				for i, checklabelevent := range objs[objnum].labelevents {

					if checklabelevent.event == "collides with" {
						upobjlabelevents(objnum, num, i, "collides with")
					}

				}
			}

			for eventnum, objevents := range objs[objnum].events {
				if objevents.event_type == 1 {
					if objs[objnum].events[eventnum].bounce || objs[objnum].events[eventnum].bounce_random {
						if objs[objnum].direction_x > 0 {
							if objs[objnum].events[eventnum].bounce {
								bounceobj(objnum, 4)
							} else if objs[objnum].events[eventnum].bounce_random {
								bounceobjrandom(objnum, 4)
							}
						} else {
							if objs[objnum].events[eventnum].bounce {
								bounceobj(objnum, 2)
							} else if objs[objnum].events[eventnum].bounce_random {
								bounceobjrandom(objnum, 2)
							}
						}
						if objs[objnum].direction_y > 0 {
							if objs[objnum].events[eventnum].bounce {
								bounceobj(objnum, 1)
							} else if objs[objnum].events[eventnum].bounce_random {
								bounceobjrandom(objnum, 1)
							}
						} else {
							if objs[objnum].events[eventnum].bounce {
								bounceobj(objnum, 3)
							} else if objs[objnum].events[eventnum].bounce_random {
								bounceobjrandom(objnum, 3)
							}
						}
					}
				}
			}

			if len(checkobj.events) > 0 {
				for eventnum, objevents := range checkobj.events {
					if objevents.event_type == 1 {
						if checkobj.events[eventnum].bounce || checkobj.events[eventnum].bounce_random {
							if objs[num].direction_x > 0 {
								if checkobj.events[eventnum].bounce {
									bounceobj(num, 4)
								} else if checkobj.events[eventnum].bounce_random {
									bounceobjrandom(num, 4)
								}
							} else {
								if checkobj.events[eventnum].bounce {
									bounceobj(num, 2)
								} else if checkobj.events[eventnum].bounce_random {
									bounceobjrandom(num, 2)
								}
							}
							if objs[num].direction_y > 0 {
								if checkobj.events[eventnum].bounce {
									bounceobj(num, 1)
								} else if checkobj.events[eventnum].bounce_random {
									bounceobjrandom(num, 1)
								}
							} else {
								if checkobj.events[eventnum].bounce {
									bounceobj(num, 3)
								} else if checkobj.events[eventnum].bounce_random {
									bounceobjrandom(num, 3)
								}
							}
						}
					}
				}
			}

		}
	}

}
func checkobjboundarycollision(objnum int) { //MARK: checkobjboundarycollision

	for a := 0; a < len(objs[objnum].events); a++ {

		if objs[objnum].events[a].event_type == 0 { // boundary collisions
			objs[objnum].collisrec = rl.NewRectangle(objs[objnum].rec.X+objs[objnum].direction_x, objs[objnum].rec.Y+objs[objnum].direction_y, objs[objnum].rec.Width, objs[objnum].rec.Height)

			if objs[objnum].collisrec.X+objs[objnum].direction_x <= settings.border_left {
				if objs[objnum].events[a].bounce {
					bounceobj(objnum, 2)
				} else if objs[objnum].events[a].bounce_random {
					bounceobjrandom(objnum, 2)
				} else if objs[objnum].events[a].stop_moving {
					stopobj(objnum)
				}
			} else if objs[objnum].collisrec.X+objs[objnum].rec.Width+objs[objnum].direction_x >= settings.border_right {
				if objs[objnum].events[a].bounce {
					bounceobj(objnum, 4)
				} else if objs[objnum].events[a].bounce_random {
					bounceobjrandom(objnum, 4)
				} else if objs[objnum].events[a].stop_moving {
					stopobj(objnum)
				}
			}

			if objs[objnum].collisrec.Y+objs[objnum].rec.Height+objs[objnum].direction_y >= settings.border_bottom {
				if objs[objnum].events[a].bounce {
					bounceobj(objnum, 1)
				} else if objs[objnum].events[a].bounce_random {
					bounceobjrandom(objnum, 1)
				} else if objs[objnum].events[a].stop_moving {
					stopobj(objnum)
				}
			} else if objs[objnum].collisrec.Y+objs[objnum].direction_y <= settings.border_top {
				if objs[objnum].events[a].bounce {
					bounceobj(objnum, 3)
				} else if objs[objnum].events[a].bounce_random {
					bounceobjrandom(objnum, 3)
				} else if objs[objnum].events[a].stop_moving {
					stopobj(objnum)
				}
			}

		}
	}

}
func bounceobj(objnum, direction int) { //MARK: bounceobj

	switch direction {
	case 1:
		objs[objnum].direction_y = -objs[objnum].direction_y
	case 2:
		objs[objnum].direction_x = float32(math.Abs(float64(objs[objnum].direction_x)))
	case 3:
		objs[objnum].direction_y = float32(math.Abs(float64(objs[objnum].direction_y)))
	case 4:
		objs[objnum].direction_x = -objs[objnum].direction_x
	}

}
func bounceobjrandom(objnum, direction int) { //MARK: bounceobjrandom

	switch direction {
	case 1:
		value2 := float32(math.Abs(float64(objs[objnum].orig_direcy))) + rFloat32(1, 4)
		newvalue := rFloat32(1, value2)
		if newvalue < 1 && objs[objnum].orig_direcy > 1 {
			newvalue += 1
		}
		objs[objnum].direction_y = -newvalue
		if objs[objnum].direction_x != 0 {
			objs[objnum].direction_x = rFloat32(-objs[objnum].orig_direcx, objs[objnum].orig_direcx+1)
		}
	case 2:
		value2 := float32(math.Abs(float64(objs[objnum].orig_direcx))) + rFloat32(1, 4)
		newvalue := rFloat32(1, value2)
		if newvalue < 1 && objs[objnum].orig_direcx > 1 {
			newvalue += 1
		}
		objs[objnum].direction_x = newvalue
		if objs[objnum].direction_y != 0 {
			objs[objnum].direction_y = rFloat32(-objs[objnum].orig_direcy, objs[objnum].orig_direcy+1)
		}
	case 3:
		value2 := float32(math.Abs(float64(objs[objnum].orig_direcy))) + rFloat32(1, 4)
		newvalue := rFloat32(1, value2)
		if newvalue < 1 && objs[objnum].orig_direcy > 1 {
			newvalue += 1
		}
		objs[objnum].direction_y = newvalue
		if objs[objnum].direction_x != 0 {
			objs[objnum].direction_x = rFloat32(-objs[objnum].orig_direcx, objs[objnum].orig_direcx+1)
		}
	case 4:
		value2 := float32(math.Abs(float64(objs[objnum].orig_direcx))) + rFloat32(1, 4)
		newvalue := rFloat32(1, value2)
		if newvalue < 1 && objs[objnum].orig_direcx > 1 {
			newvalue += 1
		}
		objs[objnum].direction_x = -newvalue
		if objs[objnum].direction_y != 0 {
			objs[objnum].direction_y = rFloat32(-objs[objnum].orig_direcy, objs[objnum].orig_direcy+1)
		}
	}

}
func stopobj(objnum int) { //MARK: stopobj

	objs[objnum].direction_x = 0
	objs[objnum].direction_y = 0

	objs[objnum].orig_direcx = 0
	objs[objnum].orig_direcy = 0

}

//MARK: UPDATE UPDATE UPDATE UPDATE UPDATE UPDATE UPDATE UPDATE UPDATE UPDATE UPDATE UPDATE
func upobjs() { //MARK: upobjs

	for a := 0; a < len(objs); a++ {

		if settings.animate {
			if objs[a].complex {
				for b := 0; b < len(objs[a].objsin); b++ {
					//img rotation
					if objs[a].objsin[b].img_rotates {
						if objs[a].objsin[b].img_rotate_lr {
							objs[a].objsin[b].img_rotation -= objs[a].objsin[b].img_rotate_speed
						} else {
							objs[a].objsin[b].img_rotation += objs[a].objsin[b].img_rotate_speed
						}
						if objs[a].objsin[b].img_rotation > 360 {
							objs[a].objsin[b].img_rotation = 0
						}
						if objs[a].objsin[b].img_rotation < -360 {
							objs[a].objsin[b].img_rotation = 0
						}
					}
				}
			}
			//img rotation
			if objs[a].img_rotates {
				if objs[a].img_rotate_lr {
					objs[a].img_rotation -= objs[a].img_rotate_speed
				} else {
					objs[a].img_rotation += objs[a].img_rotate_speed
				}
				if objs[a].img_rotation > 360 {
					objs[a].img_rotation = 0
				}
				if objs[a].img_rotation < -360 {
					objs[a].img_rotation = 0
				}
			}

			//polygons
			if objs[a].shape > 1 {
				if objs[a].rotates {
					if objs[a].rotate_lr {
						objs[a].rotation -= objs[a].rotate_speed
					} else {
						objs[a].rotation += objs[a].rotate_speed
					}
				}
			}

			if objs[a].random_direction {
				if objs[a].rand_direc_time > 0 {
					objs[a].rand_direc_time--
					if objs[a].rand_direc_time < 0 {
						objs[a].rand_direc_time = 0
					}
				}
			}
			if objs[a].rotates {
				if objs[a].rotate_timer > 0 {
					objs[a].rotate_timer--
					if objs[a].rotate_timer <= 0 {
						objs[a].rotate_timer = 0
						if objs[a].rotate_rand {
							objs[a].rotate_lr = flipcoin()
							objs[a].rotate_timer = rFloat32(0, objs[a].orig_rotatetimer+1)
							objs[a].rotate_speed = rFloat32(0, objs[a].orig_rotatespeed+1)
						} else {
							if objs[a].rotate_lr {
								objs[a].rotate_lr = false
							} else {
								objs[a].rotate_lr = true
							}
							objs[a].rotate_timer = objs[a].orig_rotatetimer
						}
					}
				}
			}
		}

	}

}
func upobjlabelevents(objnum, obj2num, labeleventnum int, event string) { //MARK: upobjlabelevents

	switch event {

	case "collides with":
		switch objs[objnum].labelevents[labeleventnum].action {
		case "changes direction":
			if objs[objnum].labelchangedirecxplus {
				objs[objnum].direction_x += objs[objnum].labelchangedirecx
			} else {
				objs[objnum].direction_x = objs[objnum].labelchangedirecx
			}
			if objs[objnum].labelchangedirecyplus {
				objs[objnum].direction_y += objs[objnum].labelchangedirecy
			} else {
				objs[objnum].direction_y = objs[objnum].labelchangedirecy
			}
		case "changes rotation":
			if objs[objnum].labelchangerotationplus {
				objs[objnum].rotation += objs[objnum].labelchangerotation
			} else {
				objs[objnum].rotation = objs[objnum].labelchangerotation
			}

		case "changes size":
			if objs[objnum].labelchangewidthplus && objs[objnum].labelchangeheightplus {
				objs[objnum].rec = rl.NewRectangle(objs[objnum].topleft_x, objs[objnum].topleft_y, objs[objnum].width+objs[objnum].labelchangewidth, objs[objnum].height+objs[objnum].labelchangeheight)
			} else if objs[objnum].labelchangewidthplus && !objs[objnum].labelchangeheightplus {
				objs[objnum].rec = rl.NewRectangle(objs[objnum].topleft_x, objs[objnum].topleft_y, objs[objnum].width+objs[objnum].labelchangewidth, objs[objnum].labelchangeheight)
			} else if !objs[objnum].labelchangewidthplus && objs[objnum].labelchangeheightplus {
				objs[objnum].rec = rl.NewRectangle(objs[objnum].topleft_x, objs[objnum].topleft_y, objs[objnum].labelchangewidth, objs[objnum].height+objs[objnum].labelchangeheight)
			} else if !objs[objnum].labelchangewidthplus && !objs[objnum].labelchangeheightplus {
				objs[objnum].rec = rl.NewRectangle(objs[objnum].topleft_x, objs[objnum].topleft_y, objs[objnum].labelchangewidth, objs[objnum].labelchangeheight)
			}

			objs[objnum].width = objs[objnum].rec.Width
			objs[objnum].height = objs[objnum].rec.Height
		case "changes image":
			objs[objnum].img = objs[objnum].labelchangeimg
		case "changes fill":
			objs[objnum].fill_color1 = objs[objnum].labelchangefill1
			if objs[objnum].labelchangefillgradienton {
				objs[objnum].fill_color2 = objs[objnum].labelchangefill2
				if objs[objnum].labelchangefillgradienton {
					if objs[objnum].labelchangefillgradienthv {
						objs[objnum].gradient_v = true
					} else {
						objs[objnum].gradient_h = true
					}
				}
			}
		case "changes ghosting":
			if objs[objnum].labelchangeghosting {
				objs[objnum].ghosting = true
				objs[objnum].ghosting_x = objs[objnum].labelchangeghostx
				objs[objnum].ghosting_y = objs[objnum].labelchangeghosty
				objs[objnum].ghosting_fade = objs[objnum].labelchangeghostfade
				objs[objnum].ghosting_color = objs[objnum].labelchangeghostcolor
			} else {
				objs[objnum].ghosting = false
			}
		case "changes shadow":
			if objs[objnum].labelchangeshadow {
				objs[objnum].shadow = true
				objs[objnum].shadow_x = objs[objnum].labelchangeshadowx
				objs[objnum].shadow_y = objs[objnum].labelchangeshadowy
				objs[objnum].shadow_fade = objs[objnum].labelchangeshadowfade
				objs[objnum].shadow_color = objs[objnum].labelchangeshadowcolor
			} else {
				objs[objnum].shadow = false
			}
		case "changes layer":
			if objs[objnum].labelchangelayer == 0 {
				objs[objnum].middleground_obj = true
				objs[objnum].background_obj = false
				objs[objnum].foreground_obj = false
			}
			if objs[objnum].labelchangelayer == -1 {
				objs[objnum].middleground_obj = false
				objs[objnum].background_obj = true
				objs[objnum].foreground_obj = false
			}
			if objs[objnum].labelchangelayer == 1 {
				objs[objnum].middleground_obj = false
				objs[objnum].background_obj = false
				objs[objnum].foreground_obj = true
			}
		case "changes position":
			if objs[objnum].labelchangeposxplus {
				objs[objnum].rec.X += objs[objnum].labelchangeposx
			} else {
				objs[objnum].rec.X = objs[objnum].labelchangeposx
			}
			if objs[objnum].labelchangeposyplus {
				objs[objnum].rec.Y += objs[objnum].labelchangeposy
			} else {
				objs[objnum].rec.Y = objs[objnum].labelchangeposy
			}
		case "changes hp":
			if objs[objnum].labelchangehpplus {
				objs[objnum].hp += objs[objnum].labelchangehp
			} else {
				objs[objnum].hp = objs[objnum].labelchangehp
			}
		}

		if objs[objnum].labelevents[labeleventnum].actionobj2 != "does nothing" {
			switch objs[objnum].labelevents[labeleventnum].actionobj2 {
			case "changes hp":
				if objs[objnum].labelchangehpplusobj2 {
					objs[obj2num].hp += objs[objnum].labelchangehpobj2
				} else {
					objs[obj2num].hp = objs[objnum].labelchangehpobj2
				}
			case "changes position":
				if objs[objnum].labelchangeposxplusobj2 {
					objs[obj2num].rec.X += objs[objnum].labelchangeposxobj2
				} else {
					objs[obj2num].rec.X = objs[objnum].labelchangeposxobj2
				}
				if objs[objnum].labelchangeposyplusobj2 {
					objs[obj2num].rec.Y += objs[objnum].labelchangeposyobj2
				} else {
					objs[obj2num].rec.Y = objs[objnum].labelchangeposyobj2
				}
			case "changes layer":
				if objs[objnum].labelchangelayerobj2 == -1 {
					objs[obj2num].background_obj = true
					objs[obj2num].middleground_obj = false
					objs[obj2num].foreground_obj = false
				} else if objs[objnum].labelchangelayerobj2 == 0 {
					objs[obj2num].background_obj = false
					objs[obj2num].middleground_obj = true
					objs[obj2num].foreground_obj = false
				} else if objs[objnum].labelchangelayerobj2 == 1 {
					objs[obj2num].background_obj = false
					objs[obj2num].middleground_obj = false
					objs[obj2num].foreground_obj = true
				}
			case "changes shadow":
				if objs[objnum].labelchangeghostingobj2 {
					objs[obj2num].shadow_color = objs[objnum].labelchangeshadowcolorobj2
					objs[obj2num].shadow_fade = objs[objnum].labelchangeshadowfadeobj2
					objs[obj2num].shadow_x = objs[objnum].labelchangeshadowxobj2
					objs[obj2num].shadow_y = objs[objnum].labelchangeshadowyobj2
				} else {
					objs[obj2num].shadow = false
				}
			case "changes ghosting":
				if objs[objnum].labelchangeghostingobj2 {
					objs[obj2num].ghosting_color = objs[objnum].labelchangeghostcolorobj2
					objs[obj2num].ghosting_fade = objs[objnum].labelchangeghostfadeobj2
					objs[obj2num].ghosting_x = objs[objnum].labelchangeghostxobj2
					objs[obj2num].ghosting_y = objs[objnum].labelchangeghostyobj2
				} else {
					objs[obj2num].ghosting = false
				}
			case "changes fill":
				objs[obj2num].fill_color1 = objs[objnum].labelchangefill1obj2
				if objs[objnum].labelchangefillgradientonobj2 {
					objs[obj2num].fill_color2 = objs[objnum].labelchangefill2obj2
					if objs[objnum].labelchangefillgradienthvobj2 {
						objs[obj2num].gradient_v = true
						objs[obj2num].gradient_h = false
					} else {
						objs[obj2num].gradient_v = false
						objs[obj2num].gradient_h = true
					}
				}
			case "changes image":
				objs[obj2num].img = objs[objnum].labelchangeimgobj2

			case "changes direction":
				if objs[objnum].labelchangedirecxplusobj2 {
					objs[obj2num].direction_x += objs[objnum].labelchangedirecxobj2
				} else {
					objs[obj2num].direction_x = objs[objnum].labelchangedirecxobj2
				}
				if objs[objnum].labelchangedirecyplusobj2 {
					objs[obj2num].direction_y += objs[objnum].labelchangedirecyobj2
				} else {
					objs[obj2num].direction_y = objs[objnum].labelchangedirecyobj2
				}

			case "changes rotation":
				if objs[objnum].labelchangerotationplusobj2 {
					objs[obj2num].rotation += objs[objnum].labelchangerotationobj2
				} else {
					objs[obj2num].rotation = objs[objnum].labelchangerotationobj2
				}

			case "changes size":
				if objs[objnum].labelchangewidthplusobj2 && objs[objnum].labelchangeheightplusobj2 {
					objs[obj2num].rec = rl.NewRectangle(objs[obj2num].topleft_x, objs[obj2num].topleft_y, objs[obj2num].width+objs[objnum].labelchangewidthobj2, objs[obj2num].height+objs[objnum].labelchangeheightobj2)
				} else if objs[objnum].labelchangewidthplusobj2 && !objs[objnum].labelchangeheightplusobj2 {
					objs[obj2num].rec = rl.NewRectangle(objs[obj2num].topleft_x, objs[obj2num].topleft_y, objs[obj2num].width+objs[objnum].labelchangewidthobj2, objs[objnum].labelchangeheightobj2)
				} else if !objs[objnum].labelchangewidthplusobj2 && objs[objnum].labelchangeheightplusobj2 {
					objs[obj2num].rec = rl.NewRectangle(objs[obj2num].topleft_x, objs[obj2num].topleft_y, objs[objnum].labelchangewidthobj2, objs[obj2num].height+objs[objnum].labelchangeheightobj2)
				} else if !objs[obj2num].labelchangewidthplusobj2 && !objs[obj2num].labelchangeheightplusobj2 {
					objs[obj2num].rec = rl.NewRectangle(objs[obj2num].topleft_x, objs[obj2num].topleft_y, objs[objnum].labelchangewidthobj2, objs[objnum].labelchangeheightobj2)
				}

				objs[objnum].width = objs[objnum].rec.Width
				objs[objnum].height = objs[objnum].rec.Height
			}

		}

		if objs[objnum].labelevents[labeleventnum].action2 != "does nothing" {
			switch objs[objnum].labelevents[labeleventnum].action2 {
			case "changes direction":
				if objs[objnum].labelchangedirecxplus {
					objs[objnum].direction_x += objs[objnum].labelchangedirecx
				} else {
					objs[objnum].direction_x = objs[objnum].labelchangedirecx
				}
				if objs[objnum].labelchangedirecyplus {
					objs[objnum].direction_y += objs[objnum].labelchangedirecy
				} else {
					objs[objnum].direction_y = objs[objnum].labelchangedirecy
				}
			case "changes rotation":
				if objs[objnum].labelchangerotationplus {
					objs[objnum].rotation += objs[objnum].labelchangerotation
				} else {
					objs[objnum].rotation = objs[objnum].labelchangerotation
				}

			case "changes size":
				if objs[objnum].labelchangewidthplus && objs[objnum].labelchangeheightplus {
					objs[objnum].rec = rl.NewRectangle(objs[objnum].topleft_x, objs[objnum].topleft_y, objs[objnum].width+objs[objnum].labelchangewidth, objs[objnum].height+objs[objnum].labelchangeheight)
				} else if objs[objnum].labelchangewidthplus && !objs[objnum].labelchangeheightplus {
					objs[objnum].rec = rl.NewRectangle(objs[objnum].topleft_x, objs[objnum].topleft_y, objs[objnum].width+objs[objnum].labelchangewidth, objs[objnum].labelchangeheight)
				} else if !objs[objnum].labelchangewidthplus && objs[objnum].labelchangeheightplus {
					objs[objnum].rec = rl.NewRectangle(objs[objnum].topleft_x, objs[objnum].topleft_y, objs[objnum].labelchangewidth, objs[objnum].height+objs[objnum].labelchangeheight)
				} else if !objs[objnum].labelchangewidthplus && !objs[objnum].labelchangeheightplus {
					objs[objnum].rec = rl.NewRectangle(objs[objnum].topleft_x, objs[objnum].topleft_y, objs[objnum].labelchangewidth, objs[objnum].labelchangeheight)
				}

				objs[objnum].width = objs[objnum].rec.Width
				objs[objnum].height = objs[objnum].rec.Height
			case "changes image":
				objs[objnum].img = objs[objnum].labelchangeimg
			case "changes fill":
				objs[objnum].fill_color1 = objs[objnum].labelchangefill1
				if objs[objnum].labelchangefillgradienton {
					objs[objnum].fill_color2 = objs[objnum].labelchangefill2
					if objs[objnum].labelchangefillgradienton {
						if objs[objnum].labelchangefillgradienthv {
							objs[objnum].gradient_v = true
						} else {
							objs[objnum].gradient_h = true
						}
					}
				}
			case "changes ghosting":
				if objs[objnum].labelchangeghosting {
					objs[objnum].ghosting = true
					objs[objnum].ghosting_x = objs[objnum].labelchangeghostx
					objs[objnum].ghosting_y = objs[objnum].labelchangeghosty
					objs[objnum].ghosting_fade = objs[objnum].labelchangeghostfade
					objs[objnum].ghosting_color = objs[objnum].labelchangeghostcolor
				} else {
					objs[objnum].ghosting = false
				}
			case "changes shadow":
				if objs[objnum].labelchangeshadow {
					objs[objnum].shadow = true
					objs[objnum].shadow_x = objs[objnum].labelchangeshadowx
					objs[objnum].shadow_y = objs[objnum].labelchangeshadowy
					objs[objnum].shadow_fade = objs[objnum].labelchangeshadowfade
					objs[objnum].shadow_color = objs[objnum].labelchangeshadowcolor
				} else {
					objs[objnum].shadow = false
				}
			case "changes layer":
				if objs[objnum].labelchangelayer == 0 {
					objs[objnum].middleground_obj = true
					objs[objnum].background_obj = false
					objs[objnum].foreground_obj = false
				}
				if objs[objnum].labelchangelayer == -1 {
					objs[objnum].middleground_obj = false
					objs[objnum].background_obj = true
					objs[objnum].foreground_obj = false
				}
				if objs[objnum].labelchangelayer == 1 {
					objs[objnum].middleground_obj = false
					objs[objnum].background_obj = false
					objs[objnum].foreground_obj = true
				}
			case "changes position":
				if objs[objnum].labelchangeposxplus {
					objs[objnum].rec.X += objs[objnum].labelchangeposx
				} else {
					objs[objnum].rec.X = objs[objnum].labelchangeposx
				}
				if objs[objnum].labelchangeposyplus {
					objs[objnum].rec.Y += objs[objnum].labelchangeposy
				} else {
					objs[objnum].rec.Y = objs[objnum].labelchangeposy
				}
			case "changes hp":
				if objs[objnum].labelchangehpplus {
					objs[objnum].hp += objs[objnum].labelchangehp
				} else {
					objs[objnum].hp = objs[objnum].labelchangehp
				}
			}

		}

		if objs[objnum].labelevents[labeleventnum].action2obj2 != "does nothing" {
			switch objs[objnum].labelevents[labeleventnum].actionobj2 {
			case "changes hp":
				if objs[objnum].labelchangehpplusobj2 {
					objs[obj2num].hp += objs[objnum].labelchangehpobj2
				} else {
					objs[obj2num].hp = objs[objnum].labelchangehpobj2
				}
			case "changes position":
				if objs[objnum].labelchangeposxplusobj2 {
					objs[obj2num].rec.X += objs[objnum].labelchangeposxobj2
				} else {
					objs[obj2num].rec.X = objs[objnum].labelchangeposxobj2
				}
				if objs[objnum].labelchangeposyplusobj2 {
					objs[obj2num].rec.Y += objs[objnum].labelchangeposyobj2
				} else {
					objs[obj2num].rec.Y = objs[objnum].labelchangeposyobj2
				}
			case "changes layer":
				if objs[objnum].labelchangelayerobj2 == -1 {
					objs[obj2num].background_obj = true
					objs[obj2num].middleground_obj = false
					objs[obj2num].foreground_obj = false
				} else if objs[objnum].labelchangelayerobj2 == 0 {
					objs[obj2num].background_obj = false
					objs[obj2num].middleground_obj = true
					objs[obj2num].foreground_obj = false
				} else if objs[objnum].labelchangelayerobj2 == 1 {
					objs[obj2num].background_obj = false
					objs[obj2num].middleground_obj = false
					objs[obj2num].foreground_obj = true
				}
			case "changes shadow":
				if objs[objnum].labelchangeghostingobj2 {
					objs[obj2num].shadow_color = objs[objnum].labelchangeshadowcolorobj2
					objs[obj2num].shadow_fade = objs[objnum].labelchangeshadowfadeobj2
					objs[obj2num].shadow_x = objs[objnum].labelchangeshadowxobj2
					objs[obj2num].shadow_y = objs[objnum].labelchangeshadowyobj2
				} else {
					objs[obj2num].shadow = false
				}
			case "changes ghosting":
				if objs[objnum].labelchangeghostingobj2 {
					objs[obj2num].ghosting_color = objs[objnum].labelchangeghostcolorobj2
					objs[obj2num].ghosting_fade = objs[objnum].labelchangeghostfadeobj2
					objs[obj2num].ghosting_x = objs[objnum].labelchangeghostxobj2
					objs[obj2num].ghosting_y = objs[objnum].labelchangeghostyobj2
				} else {
					objs[obj2num].ghosting = false
				}
			case "changes fill":
				objs[obj2num].fill_color1 = objs[objnum].labelchangefill1obj2
				if objs[objnum].labelchangefillgradientonobj2 {
					objs[obj2num].fill_color2 = objs[objnum].labelchangefill2obj2
					if objs[objnum].labelchangefillgradienthvobj2 {
						objs[obj2num].gradient_v = true
						objs[obj2num].gradient_h = false
					} else {
						objs[obj2num].gradient_v = false
						objs[obj2num].gradient_h = true
					}
				}
			case "changes image":
				objs[obj2num].img = objs[objnum].labelchangeimgobj2

			case "changes direction":
				if objs[objnum].labelchangedirecxplusobj2 {
					objs[obj2num].direction_x += objs[objnum].labelchangedirecxobj2
				} else {
					objs[obj2num].direction_x = objs[objnum].labelchangedirecxobj2
				}
				if objs[objnum].labelchangedirecyplusobj2 {
					objs[obj2num].direction_y += objs[objnum].labelchangedirecyobj2
				} else {
					objs[obj2num].direction_y = objs[objnum].labelchangedirecyobj2
				}

			case "changes rotation":
				if objs[objnum].labelchangerotationplusobj2 {
					objs[obj2num].rotation += objs[objnum].labelchangerotationobj2
				} else {
					objs[obj2num].rotation = objs[objnum].labelchangerotationobj2
				}

			case "changes size":
				if objs[objnum].labelchangewidthplusobj2 && objs[objnum].labelchangeheightplusobj2 {
					objs[obj2num].rec = rl.NewRectangle(objs[obj2num].topleft_x, objs[obj2num].topleft_y, objs[obj2num].width+objs[objnum].labelchangewidthobj2, objs[obj2num].height+objs[objnum].labelchangeheightobj2)
				} else if objs[objnum].labelchangewidthplusobj2 && !objs[objnum].labelchangeheightplusobj2 {
					objs[obj2num].rec = rl.NewRectangle(objs[obj2num].topleft_x, objs[obj2num].topleft_y, objs[obj2num].width+objs[objnum].labelchangewidthobj2, objs[objnum].labelchangeheightobj2)
				} else if !objs[objnum].labelchangewidthplusobj2 && objs[objnum].labelchangeheightplusobj2 {
					objs[obj2num].rec = rl.NewRectangle(objs[obj2num].topleft_x, objs[obj2num].topleft_y, objs[objnum].labelchangewidthobj2, objs[obj2num].height+objs[objnum].labelchangeheightobj2)
				} else if !objs[obj2num].labelchangewidthplusobj2 && !objs[obj2num].labelchangeheightplusobj2 {
					objs[obj2num].rec = rl.NewRectangle(objs[obj2num].topleft_x, objs[obj2num].topleft_y, objs[objnum].labelchangewidthobj2, objs[objnum].labelchangeheightobj2)
				}

				objs[objnum].width = objs[objnum].rec.Width
				objs[objnum].height = objs[objnum].rec.Height
			}

		}
		if objs[objnum].labelevents[labeleventnum].action3 != "does nothing" {
			switch objs[objnum].labelevents[labeleventnum].action2 {
			case "changes direction":
				if objs[objnum].labelchangedirecxplus {
					objs[objnum].direction_x += objs[objnum].labelchangedirecx
				} else {
					objs[objnum].direction_x = objs[objnum].labelchangedirecx
				}
				if objs[objnum].labelchangedirecyplus {
					objs[objnum].direction_y += objs[objnum].labelchangedirecy
				} else {
					objs[objnum].direction_y = objs[objnum].labelchangedirecy
				}
			case "changes rotation":
				if objs[objnum].labelchangerotationplus {
					objs[objnum].rotation += objs[objnum].labelchangerotation
				} else {
					objs[objnum].rotation = objs[objnum].labelchangerotation
				}

			case "changes size":
				if objs[objnum].labelchangewidthplus && objs[objnum].labelchangeheightplus {
					objs[objnum].rec = rl.NewRectangle(objs[objnum].topleft_x, objs[objnum].topleft_y, objs[objnum].width+objs[objnum].labelchangewidth, objs[objnum].height+objs[objnum].labelchangeheight)
				} else if objs[objnum].labelchangewidthplus && !objs[objnum].labelchangeheightplus {
					objs[objnum].rec = rl.NewRectangle(objs[objnum].topleft_x, objs[objnum].topleft_y, objs[objnum].width+objs[objnum].labelchangewidth, objs[objnum].labelchangeheight)
				} else if !objs[objnum].labelchangewidthplus && objs[objnum].labelchangeheightplus {
					objs[objnum].rec = rl.NewRectangle(objs[objnum].topleft_x, objs[objnum].topleft_y, objs[objnum].labelchangewidth, objs[objnum].height+objs[objnum].labelchangeheight)
				} else if !objs[objnum].labelchangewidthplus && !objs[objnum].labelchangeheightplus {
					objs[objnum].rec = rl.NewRectangle(objs[objnum].topleft_x, objs[objnum].topleft_y, objs[objnum].labelchangewidth, objs[objnum].labelchangeheight)
				}

				objs[objnum].width = objs[objnum].rec.Width
				objs[objnum].height = objs[objnum].rec.Height
			case "changes image":
				objs[objnum].img = objs[objnum].labelchangeimg
			case "changes fill":
				objs[objnum].fill_color1 = objs[objnum].labelchangefill1
				if objs[objnum].labelchangefillgradienton {
					objs[objnum].fill_color2 = objs[objnum].labelchangefill2
					if objs[objnum].labelchangefillgradienton {
						if objs[objnum].labelchangefillgradienthv {
							objs[objnum].gradient_v = true
						} else {
							objs[objnum].gradient_h = true
						}
					}
				}
			case "changes ghosting":
				if objs[objnum].labelchangeghosting {
					objs[objnum].ghosting = true
					objs[objnum].ghosting_x = objs[objnum].labelchangeghostx
					objs[objnum].ghosting_y = objs[objnum].labelchangeghosty
					objs[objnum].ghosting_fade = objs[objnum].labelchangeghostfade
					objs[objnum].ghosting_color = objs[objnum].labelchangeghostcolor
				} else {
					objs[objnum].ghosting = false
				}
			case "changes shadow":
				if objs[objnum].labelchangeshadow {
					objs[objnum].shadow = true
					objs[objnum].shadow_x = objs[objnum].labelchangeshadowx
					objs[objnum].shadow_y = objs[objnum].labelchangeshadowy
					objs[objnum].shadow_fade = objs[objnum].labelchangeshadowfade
					objs[objnum].shadow_color = objs[objnum].labelchangeshadowcolor
				} else {
					objs[objnum].shadow = false
				}
			case "changes layer":
				if objs[objnum].labelchangelayer == 0 {
					objs[objnum].middleground_obj = true
					objs[objnum].background_obj = false
					objs[objnum].foreground_obj = false
				}
				if objs[objnum].labelchangelayer == -1 {
					objs[objnum].middleground_obj = false
					objs[objnum].background_obj = true
					objs[objnum].foreground_obj = false
				}
				if objs[objnum].labelchangelayer == 1 {
					objs[objnum].middleground_obj = false
					objs[objnum].background_obj = false
					objs[objnum].foreground_obj = true
				}
			case "changes position":
				if objs[objnum].labelchangeposxplus {
					objs[objnum].rec.X += objs[objnum].labelchangeposx
				} else {
					objs[objnum].rec.X = objs[objnum].labelchangeposx
				}
				if objs[objnum].labelchangeposyplus {
					objs[objnum].rec.Y += objs[objnum].labelchangeposy
				} else {
					objs[objnum].rec.Y = objs[objnum].labelchangeposy
				}
			case "changes hp":
				if objs[objnum].labelchangehpplus {
					objs[objnum].hp += objs[objnum].labelchangehp
				} else {
					objs[objnum].hp = objs[objnum].labelchangehp
				}
			}

		}
		if objs[objnum].labelevents[labeleventnum].action3obj2 != "does nothing" {
			switch objs[objnum].labelevents[labeleventnum].actionobj2 {
			case "changes hp":
				if objs[objnum].labelchangehpplusobj2 {
					objs[obj2num].hp += objs[objnum].labelchangehpobj2
				} else {
					objs[obj2num].hp = objs[objnum].labelchangehpobj2
				}
			case "changes position":
				if objs[objnum].labelchangeposxplusobj2 {
					objs[obj2num].rec.X += objs[objnum].labelchangeposxobj2
				} else {
					objs[obj2num].rec.X = objs[objnum].labelchangeposxobj2
				}
				if objs[objnum].labelchangeposyplusobj2 {
					objs[obj2num].rec.Y += objs[objnum].labelchangeposyobj2
				} else {
					objs[obj2num].rec.Y = objs[objnum].labelchangeposyobj2
				}
			case "changes layer":
				if objs[objnum].labelchangelayerobj2 == -1 {
					objs[obj2num].background_obj = true
					objs[obj2num].middleground_obj = false
					objs[obj2num].foreground_obj = false
				} else if objs[objnum].labelchangelayerobj2 == 0 {
					objs[obj2num].background_obj = false
					objs[obj2num].middleground_obj = true
					objs[obj2num].foreground_obj = false
				} else if objs[objnum].labelchangelayerobj2 == 1 {
					objs[obj2num].background_obj = false
					objs[obj2num].middleground_obj = false
					objs[obj2num].foreground_obj = true
				}
			case "changes shadow":
				if objs[objnum].labelchangeghostingobj2 {
					objs[obj2num].shadow_color = objs[objnum].labelchangeshadowcolorobj2
					objs[obj2num].shadow_fade = objs[objnum].labelchangeshadowfadeobj2
					objs[obj2num].shadow_x = objs[objnum].labelchangeshadowxobj2
					objs[obj2num].shadow_y = objs[objnum].labelchangeshadowyobj2
				} else {
					objs[obj2num].shadow = false
				}
			case "changes ghosting":
				if objs[objnum].labelchangeghostingobj2 {
					objs[obj2num].ghosting_color = objs[objnum].labelchangeghostcolorobj2
					objs[obj2num].ghosting_fade = objs[objnum].labelchangeghostfadeobj2
					objs[obj2num].ghosting_x = objs[objnum].labelchangeghostxobj2
					objs[obj2num].ghosting_y = objs[objnum].labelchangeghostyobj2
				} else {
					objs[obj2num].ghosting = false
				}
			case "changes fill":
				objs[obj2num].fill_color1 = objs[objnum].labelchangefill1obj2
				if objs[objnum].labelchangefillgradientonobj2 {
					objs[obj2num].fill_color2 = objs[objnum].labelchangefill2obj2
					if objs[objnum].labelchangefillgradienthvobj2 {
						objs[obj2num].gradient_v = true
						objs[obj2num].gradient_h = false
					} else {
						objs[obj2num].gradient_v = false
						objs[obj2num].gradient_h = true
					}
				}
			case "changes image":
				objs[obj2num].img = objs[objnum].labelchangeimgobj2

			case "changes direction":
				if objs[objnum].labelchangedirecxplusobj2 {
					objs[obj2num].direction_x += objs[objnum].labelchangedirecxobj2
				} else {
					objs[obj2num].direction_x = objs[objnum].labelchangedirecxobj2
				}
				if objs[objnum].labelchangedirecyplusobj2 {
					objs[obj2num].direction_y += objs[objnum].labelchangedirecyobj2
				} else {
					objs[obj2num].direction_y = objs[objnum].labelchangedirecyobj2
				}

			case "changes rotation":
				if objs[objnum].labelchangerotationplusobj2 {
					objs[obj2num].rotation += objs[objnum].labelchangerotationobj2
				} else {
					objs[obj2num].rotation = objs[objnum].labelchangerotationobj2
				}

			case "changes size":
				if objs[objnum].labelchangewidthplusobj2 && objs[objnum].labelchangeheightplusobj2 {
					objs[obj2num].rec = rl.NewRectangle(objs[obj2num].topleft_x, objs[obj2num].topleft_y, objs[obj2num].width+objs[objnum].labelchangewidthobj2, objs[obj2num].height+objs[objnum].labelchangeheightobj2)
				} else if objs[objnum].labelchangewidthplusobj2 && !objs[objnum].labelchangeheightplusobj2 {
					objs[obj2num].rec = rl.NewRectangle(objs[obj2num].topleft_x, objs[obj2num].topleft_y, objs[obj2num].width+objs[objnum].labelchangewidthobj2, objs[objnum].labelchangeheightobj2)
				} else if !objs[objnum].labelchangewidthplusobj2 && objs[objnum].labelchangeheightplusobj2 {
					objs[obj2num].rec = rl.NewRectangle(objs[obj2num].topleft_x, objs[obj2num].topleft_y, objs[objnum].labelchangewidthobj2, objs[obj2num].height+objs[objnum].labelchangeheightobj2)
				} else if !objs[obj2num].labelchangewidthplusobj2 && !objs[obj2num].labelchangeheightplusobj2 {
					objs[obj2num].rec = rl.NewRectangle(objs[obj2num].topleft_x, objs[obj2num].topleft_y, objs[objnum].labelchangewidthobj2, objs[objnum].labelchangeheightobj2)
				}

				objs[objnum].width = objs[objnum].rec.Width
				objs[objnum].height = objs[objnum].rec.Height
			}

		}
	}

}
func up() { //MARK: up

	inp()
	timers()
	upobjs()

}
func upmenu() { //MARK: upmenu
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
			menu.rec = rl.NewRectangle(scrwf32-menu.wid, infobar.rec.Height+1, menu.wid, scrhf32-(infobar.rec.Height+1))
		} else {
			menu.rec = rl.NewRectangle(0, infobar.rec.Height+1, menu.wid, scrhf32-(infobar.rec.Height+1))
		}
	}
}
func upobjrandomdirection(numobj int) { //MARK: upobjrandomdirection

	if objs[numobj].rand_direc_timer_min == 0 && objs[numobj].rand_direc_timer_max == 0 {
		objs[numobj].rand_direc_timer_min = 0
		objs[numobj].rand_direc_timer_max = settings.random_direc_time_max
	}

	if objs[numobj].rand_direc_max_x == 0 && objs[numobj].rand_direc_max_y == 0 {
		objs[numobj].rand_direc_max_x = settings.random_direc_xy_minmax
		objs[numobj].rand_direc_max_y = settings.random_direc_xy_minmax
		objs[numobj].rand_direc_time = rFloat32(objs[numobj].rand_direc_timer_min, objs[numobj].rand_direc_timer_max+1) * float32(fps)
	}

	if objs[numobj].rand_direc_time == 0 {
		objs[numobj].direction_x = rFloat32(-objs[numobj].rand_direc_max_x, objs[numobj].rand_direc_max_x+1)
		objs[numobj].orig_direcx = objs[numobj].direction_x
		objs[numobj].direction_y = rFloat32(-objs[numobj].rand_direc_max_y, objs[numobj].rand_direc_max_y+1)
		objs[numobj].orig_direcy = objs[numobj].direction_y

		objs[numobj].rand_direc_time = rFloat32(-objs[numobj].rand_direc_timer_min, objs[numobj].rand_direc_timer_max+1) * float32(fps)
	}

}
func upfxdemoonoff(name string) { //MARK: upfxdemoonoff
	switch name {
	case "ghost":
		fxdemoghostmenu = true
		fxdemoshadowmenu = false
		fxdemoscanlinesmenu = false
		fxdemopixelnoisemenu = false
	case "shadow":
		fxdemoghostmenu = false
		fxdemoshadowmenu = true
		fxdemoscanlinesmenu = false
		fxdemopixelnoisemenu = false
	case "scan":
		fxdemoghostmenu = false
		fxdemoshadowmenu = false
		fxdemoscanlinesmenu = true
		fxdemopixelnoisemenu = false
	case "pixel":
		fxdemoghostmenu = false
		fxdemoshadowmenu = false
		fxdemoscanlinesmenu = false
		fxdemopixelnoisemenu = true

	}
}

//MARK: MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE

func makeeventslist() { //MARK: makeeventslist

	eventslist = append(eventslist, "collision_boundary")
	eventsonoff = append(eventsonoff, false)
	eventslist = append(eventslist, "collision_obj")
	eventsonoff = append(eventsonoff, false)

}
func makenewuiobj() { //MARK: makenewuiobj

	newuiobj := uiobj{}
	newuiobj.rec = selrec
	newuiobj.outline_color = settings.outline_color

	newuiobj.shadow_color = rl.DarkGray
	newuiobj.shadow_fade = 0.4
	newuiobj.shadow_x = -10
	newuiobj.shadow_y = 10

	newuiobj.circv2 = rl.NewVector2(newuiobj.rec.X+newuiobj.rec.Width/2, newuiobj.rec.Y+newuiobj.rec.Height/2)
	newuiobj.circrad = newuiobj.rec.Width / 2

	uiobjs = append(uiobjs, newuiobj)
	activuiobjnum = len(uiobjs) - 1
}
func makeobjsin() { //MARK: makeobjsin

	objs[activobjnum].objsin = make([]obj, 0)

	numw := false

	x := objs[activobjnum].rec.X
	y := objs[activobjnum].rec.Y

	for {

		newobj := obj{}

		newobj = objs[activobjnum]

		if objs[activobjnum].tile_h == 0 || objs[activobjnum].tile_w == 0 {
			newobj.rec = rl.NewRectangle(x, y, settings.tilew, settings.tileh)
		} else {
			newobj.rec = rl.NewRectangle(x, y, objs[activobjnum].tile_w, objs[activobjnum].tile_h)
		}

		newobj.name = ""
		newobj.complex = false

		objs[activobjnum].objsin = append(objs[activobjnum].objsin, newobj)

		if objs[activobjnum].tile_w == 0 {
			x += settings.tilew
		} else {
			x += objs[activobjnum].tile_w
		}
		if !numw {
			objs[activobjnum].tilenumw++
		}

		if x >= objs[activobjnum].rec.X+objs[activobjnum].rec.Width {
			numw = true
			x = objs[activobjnum].rec.X
			if objs[activobjnum].tile_h == 0 {
				y += settings.tileh
			} else {
				y += objs[activobjnum].tile_h
			}
			objs[activobjnum].tilenumh++
			if y >= objs[activobjnum].rec.Y+objs[activobjnum].rec.Height {
				break
			}

		}

	}

}
func clearobjsin() { //MARK: clearobjsin

	count := len(objs[activobjnum].objsin) - 1
	for {
		objs[activobjnum].objsin = remobjs(objs[activobjnum].objsin, count)
		count--
		if count < 0 {

			break
		}

	}

}
func makeobj() { //MARK: makeobj

	newobj := obj{}
	newobj.rec = newobjrec
	newobj.topleft_x = newobj.rec.X
	newobj.topleft_y = newobj.rec.Y
	newobj.width = newobj.rec.Width
	newobj.height = newobj.rec.Height
	newobj.collisrec = newobjrec
	newobj.onscreen = true
	newobj.outline_color = settings.outline_color
	newobj.img = blankrec
	newobj.tile_w = settings.tilew
	newobj.tile_h = settings.tileh
	newobj.ghosting_x = settings.ghosting_x
	newobj.ghosting_y = settings.ghosting_y
	newobj.ghosting_fade = settings.ghosting_fade
	newobj.outline_w = settings.outline_w
	newobj.shadow_color = rl.DarkGray
	newobj.shadow_x = -10
	newobj.shadow_y = 10
	newobj.shadow_fade = 0.4
	newobj.fade = 1.0
	newobj.labelchangedirecx = 10
	newobj.labelchangedirecy = 10
	newobj.labelchangefill1 = randomcolor()
	newobj.labelchangefill2 = randomcolor()
	newobj.labelchangeghostcolor = randomcolor()
	newobj.labelchangeshadowcolor = randomcolor()
	newobj.labelchangerotation = 45
	switch currentlayer {
	case 0:
		newobj.middleground_obj = true
		newobj.foreground_obj = false
		newobj.background_obj = false
	case -1:
		newobj.middleground_obj = false
		newobj.foreground_obj = false
		newobj.background_obj = true
	case 1:
		newobj.middleground_obj = false
		newobj.foreground_obj = true
		newobj.background_obj = false
	}

	newobj.circv2 = rl.NewVector2(newobj.rec.X+newobj.rec.Width/2, newobj.rec.Y+newobj.rec.Height/2)
	newobj.circrad = newobj.rec.Width / 2

	objs = append(objs, newobj)
	activobjnum = len(objs) - 1

}

func makemenus() { //MARK: makemenus
	infobar.wid = 40
	menu.wid = 200
	infobar.rec = rl.NewRectangle(0, 0, scrwf32, infobar.wid)
	menu.rec = rl.NewRectangle(0, infobar.rec.Height+1, menu.wid, scrhf32-(infobar.rec.Height+1))
}
func makecolors(num int) { //MARK: makecolors

	switch num {
	case 0:

		//random colors
		colorpalrand = make([]color, 0)
		recw := colorpalrec.Width / 8
		x := colorpalrec.X
		y := colorpalrec.Y

		length := 0
		hcount := 0
		for {
			x += recw
			length++
			if x+recw > colorpalrec.X+colorpalrec.Width {
				x = colorpalrec.X
				y += recw
				hcount++
			}
			if y+recw > colorpalrec.Y+colorpalrec.Height {
				break
			}
		}
		length -= 2

		colorpalrec.Height = float32(hcount) * recw
		colorpalrec.Height += 2

		for a := 0; a < length; a++ {
			newcolor := color{}
			newcolor.r = uint8(rInt(0, 256))
			newcolor.g = uint8(rInt(0, 256))
			newcolor.b = uint8(rInt(0, 256))
			newcolor.fade = uint8(255)
			newcolor.color = rl.NewColor(newcolor.r, newcolor.g, newcolor.b, newcolor.fade)
			colorpalrand = append(colorpalrand, newcolor)

		}

		//standard colors
		colorpalstand = make([]rl.Color, 24)
		colorpalstand[0] = rl.White
		colorpalstand[1] = rl.Yellow
		colorpalstand[2] = rl.Gold
		colorpalstand[3] = rl.Orange
		colorpalstand[4] = rl.Magenta
		colorpalstand[5] = rl.Pink
		colorpalstand[6] = rl.Red
		colorpalstand[7] = rl.Maroon
		colorpalstand[8] = rl.Green
		colorpalstand[9] = rl.Lime
		colorpalstand[10] = rl.DarkGreen
		colorpalstand[11] = rl.SkyBlue
		colorpalstand[12] = rl.Blue
		colorpalstand[13] = rl.DarkBlue
		colorpalstand[14] = rl.Purple
		colorpalstand[15] = rl.Violet
		colorpalstand[16] = rl.DarkPurple
		colorpalstand[17] = rl.Beige
		colorpalstand[18] = rl.Brown
		colorpalstand[19] = rl.DarkBrown
		colorpalstand[20] = rl.LightGray
		colorpalstand[21] = rl.Gray
		colorpalstand[22] = rl.DarkGray
		colorpalstand[23] = rl.Black

		//user colors
		colorpaluser = make([]color, 96)
		for a := 0; a < len(colorpaluser); a++ {
			colorpaluser[a].color = blankcolor
		}
		colorscreated = true

	case 1:
		//random colors
		colorpalrand = make([]color, 0)
		recw := colorpalrec.Width / 8
		x := colorpalrec.X
		y := colorpalrec.Y

		length := 0
		hcount := 0
		for {
			x += recw
			length++
			if x+recw > colorpalrec.X+colorpalrec.Width {
				x = colorpalrec.X
				y += recw
				hcount++
			}
			if y+recw > colorpalrec.Y+colorpalrec.Height {
				break
			}
		}
		length -= 2

		colorpalrec.Height = float32(hcount) * recw
		colorpalrec.Height += 2

		for a := 0; a < length; a++ {
			newcolor := color{}
			newcolor.r = uint8(rInt(0, 256))
			newcolor.g = uint8(rInt(0, 256))
			newcolor.b = uint8(rInt(0, 256))
			newcolor.fade = uint8(255)
			newcolor.color = rl.NewColor(newcolor.r, newcolor.g, newcolor.b, newcolor.fade)
			colorpalrand = append(colorpalrand, newcolor)

		}
	}

}
func makeimgs() { //MARK: makeimgs

	//1 bit 16px tiles
	x := float32(0)
	y := float32(280)

	for {
		onebit16pxtiles = append(onebit16pxtiles, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x >= 192 {
			x = 0
			y += 16
		}
		if y >= 600 {
			break
		}
	}

	//1 bit kenney
	x = 0
	y = 600

	for {
		onebitkenneytiles = append(onebitkenneytiles, rl.NewRectangle(x, y, 16, 16))
		x += 17
		if x >= 815 {
			x = 0
			y += 17
		}
		if y >= 972 {
			break
		}
	}

	//1 bit various
	x = 7
	y = 985

	for {
		onebitvar2tiles = append(onebitvar2tiles, rl.NewRectangle(x, y, 18, 18))
		x += 24
		if x >= 504 {
			x = 7
			y += 24
		}
		if y >= 1482 {
			break
		}
	}

}
func maketileselectmenu() { //MARK: maketileselectmenu

	newtilemenu := menulist{}
	newtilemenu.name = "1 bit various"
	newtilemenu.onoff = true
	tilemenuonoff = append(tilemenuonoff, newtilemenu)

	newtilemenu = menulist{}
	newtilemenu.name = "1 bit tiles"
	tilemenuonoff = append(tilemenuonoff, newtilemenu)

}
func makesettings() { //MARK: makesettings

	//settings.colorpallockon = true
	//	settings.gridon = true

	settings.animate = true
	settings.snapon = true
	settings.random_direc_time_max = 3
	settings.random_direc_xy_minmax = 12
	settings.screen_height_multiplier = 1
	settings.screen_width_multiplier = 1
	settings.level_x_left = 0
	settings.level_y_top = 0
	settings.level_x_right = settings.screen_width_multiplier * scrwf32
	settings.level_y_bottom = settings.screen_height_multiplier * scrhf32
	settings.outlineson = true
	settings.ghosting_fade = 0.4
	settings.ghosting_x = 12
	settings.ghosting_y = 12
	settings.tilew = 32
	settings.tileh = 32
	settings.border_top = 0
	settings.border_bottom = scrhf32
	settings.border_left = 0
	settings.border_right = scrwf32
	settings.outline_color = rl.Green
	settings.outline_w = 1.0
	settings.camera_y_change = settings.tileh * 2
	settings.camera_x_change = settings.tileh * 2
	settings.rotate_speed = 10
	settings.rotate_timer = 5

}
func makedefkeys() { //MARK: makedefkeys
	ukey = rl.KeyKp8
	dkey = rl.KeyKp2
	lkey = rl.KeyKp4
	rkey = rl.KeyKp6
	ulkey = rl.KeyKp7
	urkey = rl.KeyKp9
	dlkey = rl.KeyKp1
	drkey = rl.KeyKp3

	up_change_def = 10
	down_change_def = 10
	left_change_def = 10
	right_change_def = 10
	upleft_change_def = 10
	upright_change_def = 10
	downleft_change_def = 10
	downright_change_def = 10
}
func makelabels() { //MARK: makelabels

	labeleventslist = append(labeleventslist, "collides with")
	labeleventslist = append(labeleventslist, "is near")
	labeleventslist = append(labeleventslist, "label event 3")
	labeleventslist = append(labeleventslist, "label event 4")

	labelsalllist = append(labelsalllist, "testlabel 1")
	labelsalllist = append(labelsalllist, "testlabel 2")
	labelsalllist = append(labelsalllist, "testlabel 3")
	labelsalllist = append(labelsalllist, "testlabel 4")

	labelactionslist = append(labelactionslist, "does nothing")
	labelactionslist = append(labelactionslist, "changes direction")
	labelactionslist = append(labelactionslist, "changes rotation")
	labelactionslist = append(labelactionslist, "changes size")
	labelactionslist = append(labelactionslist, "changes image")
	labelactionslist = append(labelactionslist, "changes fill")
	labelactionslist = append(labelactionslist, "changes ghosting")
	labelactionslist = append(labelactionslist, "changes shadow")
	labelactionslist = append(labelactionslist, "changes layer")
	labelactionslist = append(labelactionslist, "changes position")
	labelactionslist = append(labelactionslist, "changes hp")

	labelactionslistobj2 = append(labelactionslistobj2, "does nothing")
	labelactionslistobj2 = append(labelactionslistobj2, "changes direction")
	labelactionslistobj2 = append(labelactionslistobj2, "changes rotation")
	labelactionslistobj2 = append(labelactionslistobj2, "changes size")
	labelactionslistobj2 = append(labelactionslistobj2, "changes image")
	labelactionslistobj2 = append(labelactionslistobj2, "changes fill")
	labelactionslistobj2 = append(labelactionslistobj2, "changes ghosting")
	labelactionslistobj2 = append(labelactionslistobj2, "changes shadow")
	labelactionslistobj2 = append(labelactionslistobj2, "changes layer")
	labelactionslistobj2 = append(labelactionslistobj2, "changes position")
	labelactionslistobj2 = append(labelactionslistobj2, "changes hp")

	newobj := obj{}
	newobj.name = "testobj 1"
	usrsaveobjs = append(usrsaveobjs, newobj)
	newobj = obj{}
	newobj.name = "testobj 2"
	usrsaveobjs = append(usrsaveobjs, newobj)
	newobj = obj{}
	newobj.name = "testobj 3"
	usrsaveobjs = append(usrsaveobjs, newobj)

}
func makefxdemo() { //MARK: makefxdemo

	num := rInt(4, 9)

	fxdemoobjs = nil

	for a := 0; a < num; a++ {
		newfx := fxdemoobj{}

		if rolldice() > 4 {
			newfx.imgon = true
			newfx.img = onebitkenneytiles[rInt(0, len(onebitkenneytiles)-1)]
			multi := rInt(5, 11)
			newfx.width = newfx.img.Width * float32(multi)
		} else {
			newfx.shape = rInt32(2, 8)
			newfx.radius = rFloat32(50, 120)
		}
		if newfx.imgon {
			newfx.x = rInt32(int(fxdemorec.X), int(fxdemorec.X+fxdemorec.Width-newfx.width))
			newfx.y = rInt32(int(fxdemorec.Y), int(fxdemorec.Y+fxdemorec.Height-newfx.width))
		} else {
			newfx.x = rInt32(int(fxdemorec.X), int(fxdemorec.X+fxdemorec.Width-(newfx.radius*2)))
			newfx.y = rInt32(int(fxdemorec.Y), int(fxdemorec.Y+fxdemorec.Height-(newfx.radius*2)))
		}
		newfx.color = randomcolor()
		newfx.rotates = flipcoin()
		newfx.rotation = rFloat32(0, 360)
		newfx.rotateamount = rFloat32(1, 7)
		newfx.fade = 1.0
		fxdemoobjs = append(fxdemoobjs, newfx)
	}

}
func makedefaultfx() { //MARK: makedefaultfx

	//scan 0
	newfx := fx{}
	newfx.color = rl.Green
	newfx.height = 1
	newfx.spacing = 3
	newfx.fade = 0.4

	scry := int32(0)
	for {
		newfx.y = append(newfx.y, scry)
		scry += newfx.height + newfx.spacing
		if scry > scrh {
			break
		}
	}

	//pixel noise 1
	defaultfx = append(defaultfx, newfx)

	newfx = fx{}
	newfx.num = 1
	newfx.color = rl.Green
	newfx.fade = 1.0
	newfx.min = 1
	newfx.max = 5

	num := 100
	for {

		newfx.xf32 = append(newfx.xf32, rFloat32(0, scrwf32))
		newfx.yf32 = append(newfx.yf32, rFloat32(0, scrhf32))
		num--
		if num == 0 {
			break
		}
	}
	defaultfx = append(defaultfx, newfx)

	//ghosting 2
	newfx = fx{}
	newfx.x2 = -10
	newfx.y2 = 10
	newfx.color = rl.DarkBlue
	newfx.fade = 0.5

	defaultfx = append(defaultfx, newfx)

}

//MARK: INPUT INPUT INPUT INPUT INPUT INPUT INPUT INPUT INPUT INPUT INPUT INPUT INPUT INPUT INPUT
func createcontrols(objnum, inptype int) { //MARK: createcontrols

	objs[objnum].controlson = true

	switch inptype {
	case 0:

		newcontrols := controlinp{}
		newcontrols.inptype = 0
		newcontrols.dkey = dkey
		newcontrols.lkey = lkey
		newcontrols.ukey = ukey
		newcontrols.rkey = rkey

		newcontrols.down = down_change_def
		newcontrols.up = up_change_def
		newcontrols.right = right_change_def
		newcontrols.left = left_change_def

		objs[objnum].controls = append(objs[objnum].controls, newcontrols)

	}

}
func inp() { //MARK: inp

	if activobjnum != blanknum {
		if objs[activobjnum].controlson {
			for a := 0; a < len(objs[activobjnum].controls); a++ {
				if objs[activobjnum].controls[a].inptype == 0 {
					if rl.IsKeyDown(objs[activobjnum].controls[a].ukey) {
						objs[activobjnum].rec.Y -= objs[activobjnum].controls[a].up
					}
					if rl.IsKeyDown(objs[activobjnum].controls[a].dkey) {
						objs[activobjnum].rec.Y += objs[activobjnum].controls[a].down
					}
					if rl.IsKeyDown(objs[activobjnum].controls[a].rkey) {
						objs[activobjnum].rec.X += objs[activobjnum].controls[a].right
					}
					if rl.IsKeyDown(objs[activobjnum].controls[a].lkey) {
						objs[activobjnum].rec.X -= objs[activobjnum].controls[a].left
					}
				}
			}
		}
	}

	if !inmenu && !menufocus {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

			selrec = blankrec
			newobjrec = blankrec
			startv2mouse = mousev2
			startv2mouseworld = rl.GetScreenToWorld2D(mousev2, camera)
			if settings.snapon {
				startv2mouse = mousepointgridscr
				startv2mouseworld = mousepointgridworld
			}
			//create path
			if createpathon {
				addv2path()
			}
			//clear active obj
			if !checkmouseobj() && !copyobjon && !createpathon {
				activobjnum = blanknum
				activuiobjnum = blanknum
				if objmenuon {
					objmenuon = false
				}
				if uiobjmenuon {
					uiobjmenuon = false
				}
			}
			//place copy obj
			if copyobjon {
				copyobj(activobjnum)
			}
		} else if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
			if newobjon {
				makeobj()
				newobjon = false
			}
			if newuiobjon {
				makenewuiobj()
				newuiobjon = false
			}
		}
		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			if settings.snapon {
				if mousev2.X > startv2mouse.X && mousev2.Y > startv2mouse.Y {
					selrec = rl.NewRectangle(startv2mouse.X, startv2mouse.Y, mousepointgridscr.X-startv2mouse.X, mousepointgridscr.Y-startv2mouse.Y)
					newobjrec = rl.NewRectangle(startv2mouseworld.X, startv2mouseworld.Y, mousepointgridworld.X-startv2mouseworld.X, mousepointgridworld.Y-startv2mouseworld.Y)
				} else if mousev2.X < startv2mouse.X && mousev2.Y > startv2mouse.Y {
					selrec = rl.NewRectangle(mousepointgridscr.X, startv2mouse.Y, startv2mouse.X-mousepointgridscr.X, mousepointgridscr.Y-startv2mouse.Y)
					newobjrec = rl.NewRectangle(mousepointgridworld.X, startv2mouseworld.Y, startv2mouseworld.X-mousepointgridworld.X, mousepointgridworld.Y-startv2mouseworld.Y)
				} else if mousev2.X > startv2mouse.X && mousev2.Y < startv2mouse.Y {
					selrec = rl.NewRectangle(startv2mouse.X, mousepointgridscr.Y, mousepointgridscr.X-startv2mouse.X, startv2mouse.Y-mousepointgridscr.Y)
					newobjrec = rl.NewRectangle(startv2mouseworld.X, mousepointgridworld.Y, mousepointgridworld.X-startv2mouseworld.X, startv2mouseworld.Y-mousepointgridworld.Y)
				} else if mousev2.X < startv2mouse.X && mousev2.Y < startv2mouse.Y {
					selrec = rl.NewRectangle(mousepointgridscr.X, mousepointgridscr.Y, startv2mouse.X-mousepointgridscr.X, startv2mouse.Y-mousepointgridscr.Y)
					newobjrec = rl.NewRectangle(mousepointgridworld.X, mousepointgridworld.Y, startv2mouseworld.X-mousepointgridworld.X, startv2mouseworld.Y-mousepointgridworld.Y)
				}
			} else {
				if mousev2world.X > startv2mouseworld.X && mousev2world.Y > startv2mouseworld.Y {
					selrec = rl.NewRectangle(startv2mouse.X, startv2mouse.Y, mousev2.X-startv2mouse.X, mousev2.Y-startv2mouse.Y)
					newobjrec = rl.NewRectangle(startv2mouseworld.X, startv2mouseworld.Y, mousev2world.X-startv2mouseworld.X, mousev2world.Y-startv2mouseworld.Y)
				} else if mousev2world.X < startv2mouseworld.X && mousev2world.Y > startv2mouseworld.Y {
					selrec = rl.NewRectangle(mousev2.X, startv2mouse.Y, startv2mouse.X-mousev2.X, mousev2.Y-startv2mouse.Y)
					newobjrec = rl.NewRectangle(mousev2world.X, startv2mouseworld.Y, startv2mouseworld.X-mousev2world.X, mousev2world.Y-startv2mouseworld.Y)
				} else if mousev2world.X > startv2mouseworld.X && mousev2world.Y < startv2mouseworld.Y {
					selrec = rl.NewRectangle(startv2mouse.X, mousev2.Y, mousev2.X-startv2mouse.X, startv2mouse.Y-mousev2.Y)
					newobjrec = rl.NewRectangle(startv2mouseworld.X, mousev2world.Y, mousev2world.X-startv2mouseworld.X, startv2mouseworld.Y-mousev2world.Y)
				} else if mousev2world.X < startv2mouseworld.X && mousev2world.Y < startv2mouseworld.Y {
					selrec = rl.NewRectangle(mousev2.X, mousev2.Y, startv2mouse.X-mousev2.X, startv2mouse.Y-mousev2.Y)
					newobjrec = rl.NewRectangle(mousev2world.X, mousev2world.Y, startv2mouseworld.X-mousev2world.X, startv2mouseworld.Y-mousev2world.Y)
				}
			}
		}

	}

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
func checkmouseobj() bool { //MARK: checkmouseobj

	foundobj := false
	for a := 0; a < len(objs); a++ {
		if objs[a].onscreen {

			if rl.CheckCollisionPointRec(mousev2world, objs[a].rec) {
				activobjnum = a
				upmainmenuinfo("objmenuon")
				if activuiobjnum != blanknum {
					activuiobjnum = blanknum
				}
				foundobj = true
			}

		}
	}
	for a := 0; a < len(uiobjs); a++ {
		if uiobjs[a].shape == 0 {
			if rl.CheckCollisionPointRec(mousev2world, uiobjs[a].rec) {
				activuiobjnum = a
				upmainmenuinfo("uiobjmenuon")
				if activobjnum != blanknum {
					activobjnum = blanknum
				}
				foundobj = true
			}
		} else {
			if rl.CheckCollisionPointCircle(mousev2world, uiobjs[a].circv2, uiobjs[a].circrad) {
				activuiobjnum = a
				upmainmenuinfo("uiobjmenuon")
				if activobjnum != blanknum {
					activobjnum = blanknum
				}
				foundobj = true
			}
		}
	}

	return foundobj

}
func findmousepoint() { //MARK: findmousepoint

	x := camera.Target.X
	y := camera.Target.Y

	for {

		checkrec := rl.NewRectangle(x, y, settings.tilew, settings.tileh)

		if rl.CheckCollisionPointRec(mousev2world, checkrec) {
			mousepointgridworld = rl.NewVector2(x, y)
			mousepointgridscr = rl.GetWorldToScreen2D(mousepointgridworld, camera)
		}

		x += settings.tilew
		if x > camera.Target.X+scrwf32 {
			x = camera.Target.X
			y += settings.tileh
		}

		if y > camera.Target.Y+scrhf32 {
			break
		}

	}

	if settings.gridon {

		//debug mouse to camera world
		txt := fmt.Sprint(mousev2.Y)
		txtlen := rl.MeasureText(txt+" mouse y screen  ", txts)
		txt2 := fmt.Sprint(mousev2.X)
		txtlen2 := rl.MeasureText(txt2+" mouse x screen", txts)
		rl.DrawRectangle(int32(mousev2.X+10), int32(mousev2.Y+10), txtlen+txtlen2+20, 20, rl.White)
		rl.DrawText(txt+" mouse y screen  "+txt2+" mouse x screen", int32(mousev2.X+15), int32(mousev2.Y+15), txts, rl.Black)

		txt = fmt.Sprint(mousev2world.Y)
		txtlen = rl.MeasureText(txt+" mouse y world  ", txts)
		txt2 = fmt.Sprint(mousev2.X)
		txtlen2 = rl.MeasureText(txt2+" mouse x world", txts)
		rl.DrawRectangle(int32(mousev2.X+10), int32(mousev2.Y+30), txtlen+txtlen2+20, 20, rl.White)
		rl.DrawText(txt+" mouse y world  "+txt2+" mouse x world", int32(mousev2.X+15), int32(mousev2.Y+35), txts, rl.Black)
	}

}

//MARK:  CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE
func initial() { //MARK: initial

	fxdemorec = rl.NewRectangle(scrwf32/2+40, 40, (scrwf32/2)-80, scrhf32-80)
	makefxdemo()

	fxmenuon = true

	makesettings()
	makemenus()
	maketileselectmenu()
	makeeventslist()
	makedefkeys()
	makelabels()
	makedefaultfx()

	colorpalrec = rl.NewRectangle(scrwf32-menu.rec.Width+(float32(txts/2)), menu.rec.Y+float32(txts), menu.rec.Width-10, (menu.rec.Height/3)-50)
	rightnavrec = rl.NewRectangle(scrwf32-30, 50, 26, scrhf32-60)
	colorpalbackrec = rl.NewRectangle(scrwf32-menu.rec.Width, menu.rec.Y, menu.rec.Width, menu.rec.Height)
}
func timers() { //MARK: timers

	if clickpause != 0 {
		clickpause--
		if clickpause < 0 {
			clickpause = 0
		}
	}

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
		if fadeblink > 0.4 {
			fadeblink -= 0.05
		} else {
			fadeblinkon = false
		}
	} else {
		if fadeblink < 0.9 {
			fadeblink += 0.05
		} else {
			fadeblinkon = true
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
	makeimgs()
	scr(1)
	initial()
	rl.SetTargetFPS(fps)

	for !rl.WindowShouldClose() {
		frames++
		frames32++
		mousev2 = rl.GetMousePosition()
		mousev2world = rl.GetScreenToWorld2D(mousev2, camera)
		if settings.snapon {
			findmousepoint()
			if mousepointgridscr != blankv2 {
				if !inmenu {
					rl.DrawCircleV(mousepointgridscr, 4, brightorange())
				}
			}
		}
		if rl.CheckCollisionPointRec(mousev2, menu.rec) || rl.CheckCollisionPointRec(mousev2, infobar.rec) {
			inmenu = true
		} else {
			inmenu = false
			if colorpalon {
				if rl.CheckCollisionPointRec(mousev2, colorpalrec) {
					inmenu = true
				} else {
					inmenu = false
				}
			}

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
		if copyobjon {
			rl.DrawRectangleLines(int32(mousev2.X), int32(mousev2.Y), objs[activobjnum].rec.ToInt32().Width, objs[activobjnum].rec.ToInt32().Height, brightorange())
		}
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
	}

	scrhf32 = float32(scrh)
	scrwf32 = float32(scrw)
	scrhint = int(scrh)
	scrwint = int(scrw)

	cntrscr = rl.NewVector2(scrwf32/2, scrhf32/2)

	camera.Zoom = 1.0
	camtileselect.Zoom = 1.0

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
func brightorange() rl.Color {
	color := rl.NewColor(253, 95, 0, 255)
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

// MARK: random numbers
func rInt(min, max int) int {
	return rand.Intn(max-min) + min
}
func rInt32(min, max int) int32 {
	i := rand.Intn(max-min) + min
	return int32(i)
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
