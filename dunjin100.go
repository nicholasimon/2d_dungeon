package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	rl "github.com/lachee/raylib-goplus/raylib"
)

var (

	// tile extras
	imgextra1    = rl.NewRectangle(47, 298, 16, 16)
	imgextra2    = rl.NewRectangle(63, 298, 16, 16)
	imgextra3    = rl.NewRectangle(79, 298, 16, 16)
	imgextra4    = rl.NewRectangle(96, 299, 16, 16)
	imgextra5    = rl.NewRectangle(112, 298, 16, 16)
	imgextra6    = rl.NewRectangle(129, 298, 16, 16)
	imgextra7    = rl.NewRectangle(147, 299, 16, 16)
	imgextra8    = rl.NewRectangle(164, 300, 16, 16)
	imgextra9    = rl.NewRectangle(180, 300, 16, 16)
	imgextra10   = rl.NewRectangle(198, 299, 16, 16)
	layereextras = make([]blok, mapa)
	// enemies
	imgenemy1    = rl.NewRectangle(3, 374, 15, 18)
	imgenemy2    = rl.NewRectangle(26, 375, 15, 17)
	imgenemy3    = rl.NewRectangle(50, 375, 16, 15)
	imgenemy4    = rl.NewRectangle(74, 374, 15, 18)
	imgenemy5    = rl.NewRectangle(96, 374, 17, 16)
	imgenemy6    = rl.NewRectangle(119, 378, 18, 11)
	imgenemy7    = rl.NewRectangle(143, 374, 16, 16)
	imgenemy8    = rl.NewRectangle(169, 375, 17, 16)
	imgenemy9    = rl.NewRectangle(193, 375, 16, 16)
	imgenemy10   = rl.NewRectangle(216, 375, 16, 17)
	imgenemy11   = rl.NewRectangle(239, 374, 18, 17)
	imgenemy12   = rl.NewRectangle(262, 376, 16, 16)
	imgenemyskul = rl.NewRectangle(5, 401, 17, 20)
	layerenemies = make([]blokenemy, mapa)
	// stats
	statson bool
	// fx
	nighton, fxon  bool
	imgnightshadow = rl.NewRectangle(520, 360, 120, 120)
	// loot
	lootboxbloknumber, inventorycount int
	inventory                         = make([]blokloot, 10)

	magicitems   = make([]blokloot, 101)
	weaponsitems = make([]blokloot, 101)
	armoritems   = make([]blokloot, 101)
	potionsitems = make([]blokloot, 101)
	secretitems  = make([]blokloot, 101)

	animatelootboxon, openlootboxon bool

	imgcoin = rl.NewRectangle(350, 76, 14, 16)

	imglootbox = rl.NewRectangle(2, 182, 17, 16)

	imgpotion1 = rl.NewRectangle(0, 76, 16, 16)
	imgpotion2 = rl.NewRectangle(16, 75, 16, 16)
	imgpotion3 = rl.NewRectangle(32, 76, 16, 16)
	imgpotion4 = rl.NewRectangle(47, 75, 16, 16)
	imgpotion5 = rl.NewRectangle(63, 74, 16, 17)
	imgpotion6 = rl.NewRectangle(78, 74, 16, 17)
	imgpotion7 = rl.NewRectangle(94, 75, 16, 16)
	imgpotion8 = rl.NewRectangle(111, 75, 16, 16)

	imgweapon1 = rl.NewRectangle(0, 94, 16, 16)
	imgweapon2 = rl.NewRectangle(16, 94, 16, 16)
	imgweapon3 = rl.NewRectangle(32, 94, 16, 16)
	imgweapon4 = rl.NewRectangle(48, 94, 16, 16)
	imgweapon5 = rl.NewRectangle(64, 94, 16, 16)
	imgweapon6 = rl.NewRectangle(80, 95, 16, 16)
	imgweapon7 = rl.NewRectangle(96, 95, 16, 16)
	imgweapon8 = rl.NewRectangle(112, 95, 16, 16)

	imgmagic1 = rl.NewRectangle(0, 115, 16, 16)
	imgmagic2 = rl.NewRectangle(17, 115, 16, 16)
	imgmagic3 = rl.NewRectangle(34, 115, 16, 16)
	imgmagic4 = rl.NewRectangle(52, 116, 18, 17)
	imgmagic5 = rl.NewRectangle(74, 115, 16, 17)
	imgmagic6 = rl.NewRectangle(93, 116, 16, 16)
	imgmagic7 = rl.NewRectangle(110, 116, 16, 16)
	imgmagic8 = rl.NewRectangle(126, 116, 16, 16)

	imgarmor1 = rl.NewRectangle(2, 136, 16, 16)
	imgarmor2 = rl.NewRectangle(18, 136, 16, 16)
	imgarmor3 = rl.NewRectangle(34, 136, 16, 16)
	imgarmor4 = rl.NewRectangle(50, 136, 16, 16)
	imgarmor5 = rl.NewRectangle(66, 136, 16, 16)
	imgarmor6 = rl.NewRectangle(82, 136, 16, 16)
	imgarmor7 = rl.NewRectangle(98, 136, 16, 16)
	imgarmor8 = rl.NewRectangle(114, 137, 17, 16)

	imgsecret1 = rl.NewRectangle(0, 159, 16, 16)
	imgsecret2 = rl.NewRectangle(17, 160, 16, 16)
	imgsecret3 = rl.NewRectangle(34, 158, 16, 16)
	imgsecret4 = rl.NewRectangle(53, 158, 16, 16)
	imgsecret5 = rl.NewRectangle(0, 159, 16, 16)
	imgsecret6 = rl.NewRectangle(17, 160, 16, 16)
	imgsecret7 = rl.NewRectangle(34, 158, 16, 16)
	imgsecret8 = rl.NewRectangle(53, 158, 16, 16)

	// mouse
	mouserecl                       = 34
	mouserecx, mouserecy, mouseblok int
	mouseon, mousegridon            bool
	// level
	secretrevealed, levelswitch                                                            bool
	level                                                                                  = 100
	currentwall                                                                            rl.Rectangle
	currentfloor                                                                           rl.Rectangle
	roomlength, roomlength2, roomarea, roomlengthholder, roomlength2holder, roomareaholder int
	// colors
	randomfade1, randomfade2 float32
	colorizeon               bool
	// layout
	centerblok = 350
	// player
	atk, def, per               = 1, 1, 1
	playermoves, playermovesmax = 10, 10
	playerhp, playerhpmax       = 3, 3
	playerblok                  = blokplayer{}
	player, playermovestotal    int
	imgplayer1                  = rl.NewRectangle(0, 32, 16, 16)
	imghp                       = rl.NewRectangle(0, 48, 14, 12)
	// core
	mousepos                                          rl.Vector2
	monw                                              = 1280
	monh                                              = 720
	gridw                                             = monw / 32
	gridh                                             = monh / 32
	grida                                             = gridw * gridh
	draww                                             = gridw
	drawh                                             = gridh
	drawa                                             = draww * drawh
	mapw                                              = 37
	maph                                              = 20
	mapa                                              = mapw * maph
	halfw                                             = 16
	halfh                                             = 9
	boundarytop                                       = 1
	boundarybottom                                    = maph - 1
	boundaryleft                                      = 1
	boundaryright                                     = mapw - 1
	layer1                                            = make([]blok, mapa)
	layer2                                            = make([]blok, mapa)
	fps                                               = 30
	drawblock, drawblocknext, framecount              int
	debugon, gridon, centerlineson                    bool
	onoff2, onoff3, onoff6, onoff10, onoff15, onoff30 bool
	imgs                                              rl.Texture2D
	camera, camera2x, camera4x, camera8x              rl.Camera2D
	// tiles
	tilestairs1       = rl.NewRectangle(0, 60, 16, 16)
	tilegrounddetail1 = rl.NewRectangle(0, 205, 16, 16)
	tilegrounddetail2 = rl.NewRectangle(16, 205, 16, 16)
	tilegrounddetail3 = rl.NewRectangle(32, 205, 16, 16)
	tilegrounddetail4 = rl.NewRectangle(48, 205, 16, 16)
	tilegrounddetail5 = rl.NewRectangle(65, 205, 16, 16)
	tilegrounddetail6 = rl.NewRectangle(81, 205, 16, 16)
	tilegrounddetail7 = rl.NewRectangle(97, 205, 16, 16)
	tilegrounddetail8 = rl.NewRectangle(114, 205, 16, 16)
	tilegrounddetail9 = rl.NewRectangle(131, 205, 16, 16)
	tileground1       = rl.NewRectangle(0, 0, 16, 16)
	tileground2       = rl.NewRectangle(16, 0, 16, 16)
	tileground3       = rl.NewRectangle(32, 0, 16, 16)
	tileground4       = rl.NewRectangle(48, 0, 16, 16)
	tileground5       = rl.NewRectangle(64, 0, 16, 16)
	tileground6       = rl.NewRectangle(80, 0, 16, 16)
	tileground7       = rl.NewRectangle(96, 0, 16, 16)
	tileground8       = rl.NewRectangle(112, 0, 16, 16)
	tileground9       = rl.NewRectangle(128, 0, 16, 16)
	tileground10      = rl.NewRectangle(144, 0, 16, 16)
	tileground11      = rl.NewRectangle(160, 0, 16, 16)
	tileground12      = rl.NewRectangle(176, 0, 16, 16)
	tileground13      = rl.NewRectangle(192, 0, 16, 16)
	tileground14      = rl.NewRectangle(208, 0, 16, 16)
	tileground15      = rl.NewRectangle(224, 0, 16, 16)
	tileground16      = rl.NewRectangle(240, 0, 16, 16)
	tileground17      = rl.NewRectangle(256, 0, 16, 16)
	tileground18      = rl.NewRectangle(272, 0, 16, 16)
	tileground19      = rl.NewRectangle(288, 0, 16, 16)

	tilewall1  = rl.NewRectangle(0, 16, 16, 16)
	tilewall2  = rl.NewRectangle(16, 16, 16, 16)
	tilewall3  = rl.NewRectangle(32, 16, 16, 16)
	tilewall4  = rl.NewRectangle(48, 16, 16, 16)
	tilewall5  = rl.NewRectangle(64, 16, 16, 16)
	tilewall6  = rl.NewRectangle(80, 16, 16, 16)
	tilewall7  = rl.NewRectangle(96, 16, 16, 16)
	tilewall8  = rl.NewRectangle(112, 16, 16, 16)
	tilewall9  = rl.NewRectangle(128, 16, 16, 16)
	tilewall10 = rl.NewRectangle(144, 16, 16, 16)
	tilewall11 = rl.NewRectangle(160, 16, 16, 16)
	tilewall12 = rl.NewRectangle(176, 16, 16, 16)
	tilewall13 = rl.NewRectangle(192, 16, 16, 16)
	tilewall14 = rl.NewRectangle(208, 16, 16, 16)
	tilewall15 = rl.NewRectangle(224, 16, 16, 16)
	tilewall16 = rl.NewRectangle(240, 16, 16, 16)
	tilewall17 = rl.NewRectangle(256, 16, 16, 16)
	tilewall18 = rl.NewRectangle(272, 16, 16, 16)
	tilewall19 = rl.NewRectangle(288, 16, 16, 16)
	tilewall20 = rl.NewRectangle(304, 16, 16, 16)
	tilewall21 = rl.NewRectangle(320, 16, 16, 16)
	tilewall22 = rl.NewRectangle(336, 16, 16, 16)
	tilewall23 = rl.NewRectangle(352, 16, 16, 16)
	tilewall24 = rl.NewRectangle(368, 16, 16, 16)
	tilewall25 = rl.NewRectangle(384, 16, 16, 16)
	tilewall26 = rl.NewRectangle(400, 16, 16, 16)
	tilewall27 = rl.NewRectangle(416, 16, 16, 16)
	tilewall28 = rl.NewRectangle(432, 16, 16, 16)
	tilewall29 = rl.NewRectangle(448, 16, 16, 16)
	tilewall30 = rl.NewRectangle(464, 16, 16, 16)
	tilewall31 = rl.NewRectangle(480, 16, 16, 16)
	tilewall32 = rl.NewRectangle(496, 16, 16, 16)
	tilewall33 = rl.NewRectangle(512, 16, 16, 16)
	tilewall34 = rl.NewRectangle(528, 16, 16, 16)
)

/*

screen 1280X720 -  37 bloks wide X 20 height = 740 bloks (mapa)



*/
type blokenemy struct {
	movetype, movecounter int
	activ, leftright      bool
	color                 rl.Color
	img                   rl.Rectangle
	fade, rotation        float32
	desc                  string
}
type blokloot struct {
	activ                                      bool
	blocknumber, stat1, stat2, stat3, itemtype int
	color                                      rl.Color
	img                                        rl.Rectangle
	fade                                       float32
	desc                                       string
}
type blokplayer struct {
	color rl.Color
	img   rl.Rectangle
	fade  float32
	x, y  int
	desc  string
}
type blok struct {
	color                                                              rl.Color
	img                                                                rl.Rectangle
	secret, stairs, solid, drawdir, explode, activ, rotating, destruct bool
	velocity, direction                                                int
	fade                                                               float32
	desc                                                               string
}

func raylib() { // MARK: raylib
	rl.InitWindow(monw, monh, "dunjin100")
	rl.SetExitKey(rl.KeyEnd)          // key to end the game and close window
	imgs = rl.LoadTexture("imgs.png") // load images
	rl.SetTargetFPS(fps)
	//rl.HideCursor()
	//rl.ToggleFullscreen()
	for !rl.WindowShouldClose() {

		framecount++
		mousepos = rl.GetMousePosition()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)
		drawlayers()
		if fxon {
			drawfx()
		}
		rl.EndMode2D()

		drawnocamera()
		if statson {
			drawstats()
		}
		update()

		rl.EndDrawing()

	}
	rl.CloseWindow()
}
func update() { // MARK: update

	timers()
	input()

	if levelswitch {
		clearmaps()
		resetnewlevel()
		nextlevel()
		levelswitch = false
	}

	if mouseon {
		getmouseblock()
	}

	if debugon {
		debug()
	}
}
func drawlayers() { // MARK:  drawlayers

	x := 32
	y := 28
	count := 0
	countblok := 0
	for a := 0; a < len(layer1); a++ {
		//draw enemies
		if layerenemies[a].activ {
			v2 := rl.NewVector2(8, 8)
			rec := rl.NewRectangle(float32(x+8), float32(y+8), 16, 16)
			if colorizeon {
				rl.DrawTexturePro(imgs, layerenemies[a].img, rec, v2, layerenemies[a].rotation, rl.Fade(layerenemies[a].color, layerenemies[a].fade))
				//	rl.DrawTextureRec(imgs, layerenemies[a].img, v2, rl.Fade(layerenemies[a].color, layerenemies[a].fade))
			} else {
				rl.DrawTexturePro(imgs, layerenemies[a].img, rec, v2, layerenemies[a].rotation, rl.Fade(rl.White, layerenemies[a].fade))
				//rl.DrawTextureRec(imgs, layerenemies[a].img, v2, rl.Fade(rl.White, layerenemies[a].fade))
			}
		}
		// draw loot boxes
		if layer2[a].activ {
			v2 := rl.NewVector2(float32(x), float32(y))

			if colorizeon {
				if !secretrevealed && layer2[a].img == imglootbox {
				} else {
					rl.DrawTextureRec(imgs, layer2[a].img, v2, rl.Fade(layer2[a].color, layer2[a].fade))
				}
			} else {
				if !secretrevealed && layer2[a].img == imglootbox {
				} else {
					rl.DrawTextureRec(imgs, layer2[a].img, v2, rl.Fade(rl.White, layer2[a].fade))
				}
			}
			// check player proximity
			if layer2[a].img == imglootbox {
				if player+1 == a || player-1 == a || (player+1)-mapw == a || (player-1)-mapw == a || player-mapw == a || (player+1)+mapw == a || (player-1)+mapw == a || player+mapw == a {
					openlootbox(a)
					animatelootboxon = true
					lootboxbloknumber = a
				}
			}
		}
		// draw items
		if magicitems[level].activ {
			if magicitems[level].blocknumber == countblok {
				v2 := rl.NewVector2(float32(x), float32(y))
				if colorizeon {
					rl.DrawTextureRec(imgs, magicitems[level].img, v2, rl.Fade(magicitems[level].color, magicitems[level].fade-randomfade2))
				} else {
					rl.DrawTextureRec(imgs, magicitems[level].img, v2, rl.Fade(rl.White, magicitems[level].fade))
				}
			}
			if magicitems[level].blocknumber == player {
				pickupitem(magicitems[level].blocknumber, "magic")
			}
		}
		if weaponsitems[level].activ {
			if weaponsitems[level].blocknumber == countblok {
				v2 := rl.NewVector2(float32(x), float32(y))
				if colorizeon {
					rl.DrawTextureRec(imgs, weaponsitems[level].img, v2, rl.Fade(weaponsitems[level].color, weaponsitems[level].fade-randomfade2))
				} else {
					rl.DrawTextureRec(imgs, weaponsitems[level].img, v2, rl.Fade(rl.White, weaponsitems[level].fade))
				}
			}
			if weaponsitems[level].blocknumber == player {
				pickupitem(weaponsitems[level].blocknumber, "weapon")
			}
		}
		if armoritems[level].activ {
			if armoritems[level].blocknumber == countblok {
				v2 := rl.NewVector2(float32(x), float32(y))
				if colorizeon {
					rl.DrawTextureRec(imgs, armoritems[level].img, v2, rl.Fade(armoritems[level].color, armoritems[level].fade-randomfade2))
				} else {
					rl.DrawTextureRec(imgs, armoritems[level].img, v2, rl.Fade(rl.White, armoritems[level].fade))
				}
			}
			if armoritems[level].blocknumber == player {
				pickupitem(armoritems[level].blocknumber, "armor")
			}
		}
		if potionsitems[level].activ {
			if potionsitems[level].blocknumber == countblok {
				v2 := rl.NewVector2(float32(x), float32(y))
				if colorizeon {
					rl.DrawTextureRec(imgs, potionsitems[level].img, v2, rl.Fade(potionsitems[level].color, potionsitems[level].fade-randomfade2))
				} else {
					rl.DrawTextureRec(imgs, potionsitems[level].img, v2, rl.Fade(rl.White, potionsitems[level].fade))
				}
			}
			if potionsitems[level].blocknumber == player {
				pickupitem(potionsitems[level].blocknumber, "potion")
			}
		}
		if secretitems[level].activ {
			if secretitems[level].blocknumber == countblok {
				v2 := rl.NewVector2(float32(x), float32(y))
				if colorizeon {
					rl.DrawTextureRec(imgs, secretitems[level].img, v2, rl.Fade(secretitems[level].color, secretitems[level].fade-randomfade2))
				} else {
					rl.DrawTextureRec(imgs, secretitems[level].img, v2, rl.Fade(rl.White, secretitems[level].fade))
				}
			}
			if secretitems[level].blocknumber == player {
				pickupitem(secretitems[level].blocknumber, "secret")
			}
		}
		// draw extras
		if layereextras[a].activ {
			v2 := rl.NewVector2(float32(x), float32(y))
			if colorizeon {
				rl.DrawTextureRec(imgs, layereextras[a].img, v2, rl.Fade(layereextras[a].color, layereextras[a].fade))
			} else {
				rl.DrawTextureRec(imgs, layereextras[a].img, v2, rl.Fade(rl.White, layereextras[a].fade))
			}
		}
		// draw tiles
		if layer1[a].activ {
			v2 := rl.NewVector2(float32(x), float32(y))
			if colorizeon {
				rl.DrawTextureRec(imgs, layer1[a].img, v2, rl.Fade(layer1[a].color, layer1[a].fade))
			} else {
				rl.DrawTextureRec(imgs, layer1[a].img, v2, rl.Fade(rl.White, layer1[a].fade))
			}
			// reveal secret room
			if layer1[a].secret {
				if player+1 == a || player-1 == a || (player+1)-mapw == a || (player-1)-mapw == a || player-mapw == a || (player+1)+mapw == a || (player-1)+mapw == a || player+mapw == a {
					if !secretrevealed {
						layer1[a].solid = false
						revealsecret()
					}
				}
			}
		}

		// draw player
		if player == countblok {
			v2 := rl.NewVector2(float32(x), float32(y))
			if onoff6 {
				v2.Y--
			}
			playerblok.x = int(v2.X)
			playerblok.y = int(v2.Y)
			if colorizeon {
				rl.DrawTextureRec(imgs, playerblok.img, v2, rl.Fade(playerblok.color, playerblok.fade))
			} else {
				rl.DrawTextureRec(imgs, playerblok.img, v2, rl.Fade(rl.White, playerblok.fade))
			}

		}
		// end level
		if layer1[player].stairs {
			levelswitch = true
		}

		x += 16
		count++
		countblok++
		if count == mapw {
			count = 0
			x = 32
			y += 16
		}

	}

	// draw layer 2 / night
	x = 32
	y = 28
	count = 0
	countblok = 0
	for a := 0; a < len(layer1); a++ {

		if player == countblok {
			v2 := rl.NewVector2(float32(x), float32(y))

			// night
			if nighton {
				v2 = rl.NewVector2(float32(x-52), float32(y-52))
				rl.DrawTextureRec(imgs, imgnightshadow, v2, rl.White)
				rl.DrawRectangle(0, 0, monw, y-52, rl.Fade(rl.Black, 0.8))
				rl.DrawRectangle(0, y+68, monw, monh, rl.Fade(rl.Black, 0.8))
				rl.DrawRectangle(0, y-52, x-52, 120, rl.Fade(rl.Black, 0.8))
				rl.DrawRectangle(x+68, y-52, monw, 120, rl.Fade(rl.Black, 0.8))
			}

		}

		x += 16
		count++
		countblok++
		if count == mapw {
			count = 0
			x = 32
			y += 16
		}

	}

	// menubars
	drawmenubars()
	// stats

	// grid
	if gridon {
		x := 16
		for a := 0; a < gridw; a++ {
			rl.DrawLine(x, 0, x, monh/2, rl.Fade(rl.Green, 0.4))
			x += 16
		}
		y := 12
		for a := 0; a < gridh; a++ {
			rl.DrawLine(0, y, monw/2, y, rl.Fade(rl.Green, 0.4))
			y += 16
		}

	}
}
func drawnocamera() { // MARK: drawnocamera

	// menubars
	drawmenubarsnocamera()

	x := 64
	y := 56
	count := 0
	countblok := 0
	for a := 0; a < mapa; a++ {

		// mouse blok highlight
		if mouseblok == a && mousegridon {
			if colorizeon {
				rl.DrawRectangle(x, y, 32, 32, rl.Fade(randomcolor(), 0.4))
			} else {
				rl.DrawRectangleLines((x-2)-mouserecx, (y-2)-mouserecy, mouserecl, mouserecl, rl.White)

			}
		}

		// gridnumbers
		if gridon {
			countbloktext := strconv.Itoa(countblok)
			rl.DrawText(countbloktext, x+6, y+10, 10, rl.Fade(rl.Green, 0.5))
		}
		x += 32
		count++
		countblok++
		if count == mapw {
			count = 0
			y += 32
			x = 64
		}

	}

	// draw center lines
	if centerlineson {
		rl.DrawLine(monw/2, 0, monw/2, monh, rl.Green)
		rl.DrawLine(0, monh/2, monw, monh/2, rl.Green)
	}
}
func drawfx() { // MARK: drawfx

	// scan lines
	y := 0
	for a := 0; a < monh/2; a++ {
		rl.DrawLine(0, y, monw, y, rl.Fade(rl.Black, 0.3))
		y += 2
	}
	// pixel noise
	for a := 0; a < 100; a++ {
		rl.DrawRectangle(rInt(10, monw-10), rInt(10, monh-10), 1, 1, rl.Black)
	}

}
func drawmenubarsnocamera() { // MARK: drawmenubarsnocamera

	// blok text
	x := (monw / 3) * 2
	y := 16
	if mousegridon && layer1[mouseblok].activ && mouseblok != player && magicitems[level].blocknumber != mouseblok && weaponsitems[level].blocknumber != mouseblok && armoritems[level].blocknumber != mouseblok && secretitems[level].blocknumber != mouseblok && potionsitems[level].blocknumber != mouseblok {
		rl.DrawText(layer1[mouseblok].desc, x, y, 20, rl.White)
	} else if mousegridon && magicitems[level].blocknumber == mouseblok {
		rl.DrawText(magicitems[level].desc, x, y, 20, rl.White)
	} else if mousegridon && weaponsitems[level].blocknumber == mouseblok {
		rl.DrawText(weaponsitems[level].desc, x, y, 20, rl.White)
	} else if mousegridon && armoritems[level].blocknumber == mouseblok {
		rl.DrawText(armoritems[level].desc, x, y, 20, rl.White)
	} else if mousegridon && secretitems[level].blocknumber == mouseblok {
		rl.DrawText(secretitems[level].desc, x, y, 20, rl.White)
	} else if mousegridon && potionsitems[level].blocknumber == mouseblok {
		rl.DrawText(potionsitems[level].desc, x, y, 20, rl.White)
	} else if mouseblok == player {
		rl.DrawText(playerblok.desc, x, y, 20, rl.White)
	}
	// stats
	x = 12
	y = 450

	levtext := strconv.Itoa(level)
	atktext := strconv.Itoa(atk)
	deftext := strconv.Itoa(def)
	pertext := strconv.Itoa(per)
	movestext := strconv.Itoa(playermoves)

	if playermoves < 0 {
		playermoves = 0
	}

	if level == 100 {
		rl.DrawText(levtext, x+7, y-20, 20, randomcolor())
		rl.DrawText(levtext, x+6, y-21, 20, rl.Black)
		rl.DrawText(levtext, x+5, y-22, 20, rl.White)
		rl.DrawText("LEV", x+2, y+2, 20, randomcolor())
		rl.DrawText("LEV", x+1, y+1, 20, rl.Black)
		rl.DrawText("LEV", x, y, 20, rl.White)
	} else if level < 100 && level > 10 {
		rl.DrawText(levtext, x+10, y-20, 20, randomcolor())
		rl.DrawText(levtext, x+9, y-21, 20, rl.Black)
		rl.DrawText(levtext, x+8, y-22, 20, rl.White)
		rl.DrawText("LEV", x+2, y+2, 20, randomcolor())
		rl.DrawText("LEV", x+1, y+1, 20, rl.Black)
		rl.DrawText("LEV", x, y, 20, rl.White)
	} else {
		rl.DrawText(levtext, x+17, y-20, 20, randomcolor())
		rl.DrawText(levtext, x+16, y-21, 20, rl.Black)
		rl.DrawText(levtext, x+15, y-22, 20, rl.White)
		rl.DrawText("LEV", x+2, y+2, 20, randomcolor())
		rl.DrawText("LEV", x+1, y+1, 20, rl.Black)
		rl.DrawText("LEV", x, y, 20, rl.White)
	}
	y += 48
	if playermoves > 9 {
		rl.DrawText(movestext, x+12, y-20, 20, randomcolor())
		rl.DrawText(movestext, x+11, y-21, 20, rl.Black)
		rl.DrawText(movestext, x+10, y-22, 20, rl.White)
		rl.DrawText("MOV", x+2, y+2, 20, randomcolor())
		rl.DrawText("MOV", x+1, y+1, 20, rl.Black)
		rl.DrawText("MOV", x, y, 20, rl.White)
	} else {
		rl.DrawText(movestext, x+17, y-20, 20, randomcolor())
		rl.DrawText(movestext, x+16, y-21, 20, rl.Black)
		rl.DrawText(movestext, x+15, y-22, 20, rl.White)
		rl.DrawText("MOV", x+2, y+2, 20, randomcolor())
		rl.DrawText("MOV", x+1, y+1, 20, rl.Black)
		rl.DrawText("MOV", x, y, 20, rl.White)
	}
	y += 48
	rl.DrawText(atktext, x+17, y-20, 20, randomcolor())
	rl.DrawText(atktext, x+16, y-21, 20, rl.Black)
	rl.DrawText(atktext, x+15, y-22, 20, rl.White)
	rl.DrawText("ATK", x+2, y+2, 20, randomcolor())
	rl.DrawText("ATK", x+1, y+1, 20, rl.Black)
	rl.DrawText("ATK", x, y, 20, rl.White)
	y += 48
	rl.DrawText(deftext, x+17, y-20, 20, randomcolor())
	rl.DrawText(deftext, x+16, y-21, 20, rl.Black)
	rl.DrawText(deftext, x+15, y-22, 20, rl.White)
	rl.DrawText("DEF", x+2, y+2, 20, randomcolor())
	rl.DrawText("DEF", x+1, y+1, 20, rl.Black)
	rl.DrawText("DEF", x, y, 20, rl.White)
	y += 48
	rl.DrawText(pertext, x+17, y-20, 20, randomcolor())
	rl.DrawText(pertext, x+16, y-21, 20, rl.Black)
	rl.DrawText(pertext, x+15, y-22, 20, rl.White)
	rl.DrawText("PER", x+2, y+2, 20, randomcolor())
	rl.DrawText("PER", x+1, y+1, 20, rl.Black)
	rl.DrawText("PER", x, y, 20, rl.White)

}
func drawmenubars() { // MARK: drawmenubars

	x := 8
	y := 6
	// hp
	playerhp = 2
	for a := 0; a < playerhpmax; a++ {
		v2 := rl.NewVector2(float32(x), float32(y))
		rl.DrawTextureRec(imgs, imghp, v2, rl.White)
		x += 16
	}
	x = 8
	for a := 0; a < playerhp; a++ {
		v2 := rl.NewVector2(float32(x), float32(y))
		rl.DrawTextureRec(imgs, imghp, v2, rl.Red)
		x += 16
	}
	// inventory
	x = 8
	y = 32
	for a := 0; a < 5; a++ {
		rl.DrawRectangleLines(x, y, 16, 16, rl.White)
		y += 20
	}
	y = 32
	for a := 0; a < 10; a++ {
		if inventory[a].activ {
			v2 := rl.NewVector2(float32(x), float32(y))
			if colorizeon {
				rl.DrawTextureRec(imgs, inventory[a].img, v2, rl.Fade(inventory[a].color, inventory[a].fade))
			} else {
				rl.DrawTextureRec(imgs, inventory[a].img, v2, rl.White)
			}
		}
		y += 20
	}
	// power ups
	x = 110
	y = 6
	for a := 0; a < 5; a++ {
		rl.DrawRectangleLines(x, y, 16, 16, rl.White)
		x += 20
	}
}
func drawstats() { // MARK: drawstats

	rl.DrawRectangle(monw/2-300, 0, 600, monh, rl.Fade(rl.Black, 0.9))

	leftmrg := monw/2 - 280

	rl.BeginMode2D(camera8x)

	v2 := rl.NewVector2(float32(((monw/2-300)/8)+10), 10)
	if onoff6 {
		v2.Y -= 2
	}
	rl.DrawTextureRec(imgs, playerblok.img, v2, rl.White)
	rl.EndMode2D()

	playermovestotaltext := strconv.Itoa(playermovestotal)

	rl.DrawText("total moves", leftmrg+2, 252, 40, randomcolor())
	rl.DrawText("total moves", leftmrg+1, 251, 40, rl.Black)
	rl.DrawText("total moves", leftmrg, 250, 40, rl.White)
	rl.DrawText(playermovestotaltext, monw/2+2, 252, 40, randomcolor())
	rl.DrawText(playermovestotaltext, monw/2+1, 251, 40, rl.Black)
	rl.DrawText(playermovestotaltext, monw/2, 250, 40, rl.White)

}
func pickupitem(blocknumber int, itemtype string) { // MARK: pickupitem

	switch itemtype {
	case "weapon":
		if inventory[0].activ {
			holder := weaponsitems[level]
			weaponsitems[level] = inventory[0]
			weaponsitems[level].blocknumber = findemptyblokdropinventory(blocknumber)
			inventory[0] = holder
		} else {
			inventory[0] = weaponsitems[level]
			weaponsitems[level] = blokloot{}
		}
	case "armor":
		if inventory[1].activ {
			holder := armoritems[level]
			armoritems[level] = inventory[1]
			armoritems[level].blocknumber = findemptyblokdropinventory(blocknumber)
			inventory[1] = holder
		} else {
			inventory[1] = armoritems[level]
			armoritems[level] = blokloot{}
		}
	case "secret":
		if inventory[2].activ {
			holder := secretitems[level]
			secretitems[level] = inventory[2]
			secretitems[level].blocknumber = findemptyblokdropinventory(blocknumber)
			inventory[2] = holder
		} else {
			inventory[2] = secretitems[level]
			secretitems[level] = blokloot{}
		}
	case "magic":
		if inventory[3].activ {
			holder := magicitems[level]
			magicitems[level] = inventory[3]
			magicitems[level].blocknumber = findemptyblokdropinventory(blocknumber)
			inventory[3] = holder
		} else {
			inventory[3] = magicitems[level]
			magicitems[level] = blokloot{}
		}
	case "potion":
		if inventory[4].activ {
			holder := potionsitems[level]
			potionsitems[level] = inventory[4]
			potionsitems[level].blocknumber = findemptyblokdropinventory(blocknumber)
			inventory[4] = holder
		} else {
			inventory[4] = potionsitems[level]
			potionsitems[level] = blokloot{}
		}
	}

}
func findemptyblokdropinventory(blocknumber int) int { // MARK: findemptyblokdropinventory

	blok1 := (blocknumber - 1) - mapw
	blok2 := (blocknumber) - mapw
	blok3 := (blocknumber + 1) - mapw
	blok4 := (blocknumber - 1)
	blok5 := (blocknumber + 1)
	blok6 := (blocknumber + 1) + mapw
	blok7 := (blocknumber - 1) + mapw
	blok8 := (blocknumber) + mapw

	selected := blok5
	end := false
	for {

		choose := rInt(1, 9)

		switch choose {
		case 1:
			if layer1[blok1].activ && !layer1[blok1].solid && !layer1[blok1].stairs {
				if magicitems[level].blocknumber != blok1 && weaponsitems[level].blocknumber != blok1 && armoritems[level].blocknumber != blok1 && secretitems[level].blocknumber != blok1 && potionsitems[level].blocknumber != blok1 {
					selected = blok1
					end = true
				}
			}
		case 2:
			if layer1[blok2].activ && !layer1[blok2].solid && !layer1[blok1].stairs {
				if magicitems[level].blocknumber != blok2 && weaponsitems[level].blocknumber != blok2 && armoritems[level].blocknumber != blok2 && secretitems[level].blocknumber != blok2 && potionsitems[level].blocknumber != blok2 {
					selected = blok2
					end = true
				}
			}
		case 3:
			if layer1[blok3].activ && !layer1[blok3].solid && !layer1[blok1].stairs {
				if magicitems[level].blocknumber != blok3 && weaponsitems[level].blocknumber != blok3 && armoritems[level].blocknumber != blok3 && secretitems[level].blocknumber != blok3 && potionsitems[level].blocknumber != blok3 {
					selected = blok3
					end = true
				}
			}
		case 4:
			if layer1[blok4].activ && !layer1[blok4].solid && !layer1[blok1].stairs {
				if magicitems[level].blocknumber != blok4 && weaponsitems[level].blocknumber != blok4 && armoritems[level].blocknumber != blok4 && secretitems[level].blocknumber != blok4 && potionsitems[level].blocknumber != blok4 {
					selected = blok4
					end = true
				}
			}
		case 5:
			if layer1[blok5].activ && !layer1[blok5].solid && !layer1[blok1].stairs {
				if magicitems[level].blocknumber != blok5 && weaponsitems[level].blocknumber != blok5 && armoritems[level].blocknumber != blok5 && secretitems[level].blocknumber != blok5 && potionsitems[level].blocknumber != blok5 {
					selected = blok5
					end = true
				}
			}
		case 6:
			if layer1[blok6].activ && !layer1[blok6].solid && !layer1[blok1].stairs {
				if magicitems[level].blocknumber != blok6 && weaponsitems[level].blocknumber != blok6 && armoritems[level].blocknumber != blok6 && secretitems[level].blocknumber != blok6 && potionsitems[level].blocknumber != blok6 {
					selected = blok6
					end = true
				}
			}
		case 7:
			if layer1[blok7].activ && !layer1[blok7].solid && !layer1[blok1].stairs {
				if magicitems[level].blocknumber != blok7 && weaponsitems[level].blocknumber != blok7 && armoritems[level].blocknumber != blok7 && secretitems[level].blocknumber != blok7 && potionsitems[level].blocknumber != blok7 {
					selected = blok7
					end = true
				}
			}
		case 8:
			if layer1[blok8].activ && !layer1[blok8].solid && !layer1[blok1].stairs {
				if magicitems[level].blocknumber != blok8 && weaponsitems[level].blocknumber != blok8 && armoritems[level].blocknumber != blok8 && secretitems[level].blocknumber != blok8 && potionsitems[level].blocknumber != blok8 {
					selected = blok8
					end = true
				}
			}

		}
		if end {
			break
		}

	}
	return selected
}
func findfloorblock() int { // MARK: findfloorblock
	block := centerblok
	for {
		block = rInt(mapw+1, (mapa - (mapw + 1)))
		if layer1[block].activ && !layer1[block].solid {
			break
		}
	}
	return block
}
func nextlevel() { // MARK: nextlevel
	level--
	createdunjin()
}
func createdunjin() {
	currentfloor = choosefloorblok()
	currentwall = choosewallblok()

	/*
	   0 square
	   1 rectangle
	   2 triangle up
	   3 triangle down
	   4 triangle left
	   5 triangle right


	*/
	// fill map - create function different wall types / colors
	for a := 0; a < mapa; a++ {

		layer1[a].activ = true
		layer1[a].solid = true
		layer1[a].img = currentwall
		layer1[a].color = randomorange()
		layer1[a].fade = 1.0
		layer1[a].desc = "wall text here"

	}

	// create default square
	roomlength = 13
	roomlength2 = 13
	roomblok := centerblok
	roomblok -= 5
	roomblok -= mapw * 5
	createroom(roomblok, 0)
	createstairs()
	createsecretroom()

}
func createstairs() { // MARK: createstairs

	stairsblock := centerblok

	for {
		stairsblock = rInt(mapw+1, (mapa - (mapw + 1)))
		if !layer1[stairsblock].solid && layer1[stairsblock].activ {
			break
		}
	}

	layer1[stairsblock].stairs = true
	layer1[stairsblock].img = tilestairs1
	layer1[stairsblock].fade = 1.0
	layer1[stairsblock].color = randomcolor()
	layer1[stairsblock].desc = "next level upstairs"

	placeplayer(stairsblock)
	placeloot()
	placeenemies()
}
func createroom(roomblok, shape int) { // MARK: createroom

	count := 0

	switch shape {
	case 0: // square
		roomarea = roomlength * roomlength
		for a := 0; a < roomarea; a++ {
			fillfloor(roomblok)
			roomblok++
			count++
			if count == roomlength {
				count = 0
				roomblok -= roomlength
				roomblok += mapw
			}
		}

	}

}
func createsecretroom() { // MARK: createsecretroom

	// choose room location
	roomblock := centerblok
	choose := rInt(1, 4)
	passagedirection := 0

	switch choose {
	case 1:
		roomblock = mapw + 1
		roomblock += rInt(0, 4)
		roomblock += rInt(0, 4) * mapw
		passagedirection = 5
	case 2:
		roomblock -= (halfw)
		roomblock += rInt(0, 4)
		roomblock += rInt(-2, 3) * mapw
		passagedirection = 4
	case 3:
		roomblock -= (halfw)
		roomblock += rInt(0, 4)
		roomblock += rInt(3, 6) * mapw
		passagedirection = 3
	}

	length := rInt(3, 5)
	width := rInt(3, 5)
	area := length * width
	count := 0
	passageblock := roomblock
	// add up down extra blocks
	h, v := getblokhv(roomblock)
	if h > 1 && h < maph-1 {
		if flipcoin() {
			fillfloor((roomblock + length/2) - mapw)
			fillsecret((roomblock + length/2) - mapw)
		}
		if flipcoin() {
			fillfloor((roomblock + length/2) + (mapw * width))
			fillsecret((roomblock + length/2) + (mapw * width))
		}
	}
	if v > 1 {

	}
	lootboxblok := roomblock + (length / 2)
	lootboxblok += (width / 2) * mapw
	// make secret room
	for a := 0; a < area; a++ {
		fillfloor(roomblock)
		fillsecret(roomblock)
		roomblock++
		count++
		if count == length {
			if rolldice() == 6 {
				fillfloor(roomblock)
				fillsecret(roomblock)
			}
			count = 0
			roomblock -= length
			roomblock += mapw
		}
	}

	passageblock += length
	passageblock += (width / 2) * mapw

	switch passagedirection {
	case 4:
		for {
			passagelength := rInt(3, 7)
			for a := 0; a < passagelength; a++ {
				fillfloor(passageblock)
				fillsecret(passageblock)
				passageblock++
				if !layer1[passageblock].solid || layer1[passageblock].stairs {
					findpassagewall(passageblock, 4)
					break
				}
			}
			if !layer1[passageblock].solid || layer1[passageblock].stairs {
				break
			}
			passagelength = rInt(1, 3)
			if flipcoin() {
				for a := 0; a < passagelength; a++ {
					fillfloor(passageblock)
					fillsecret(passageblock)
					passageblock += mapw
					if !layer1[passageblock].solid || layer1[passageblock].stairs {
						findpassagewall(passageblock, 4)
						break
					}
				}
			} else {
				for a := 0; a < passagelength; a++ {
					fillfloor(passageblock)
					fillsecret(passageblock)
					passageblock -= mapw
					if !layer1[passageblock].solid || layer1[passageblock].stairs {
						findpassagewall(passageblock, 4)
						break
					}
				}
			}
			if !layer1[passageblock].solid || layer1[passageblock].stairs {
				break
			}
		}
	case 5:
		for {
			passagelength := rInt(3, 7)
			for a := 0; a < passagelength; a++ {
				fillfloor(passageblock)
				fillsecret(passageblock)
				passageblock++
				if !layer1[passageblock].solid || layer1[passageblock].stairs {
					findpassagewall(passageblock, 4)
					break
				}
			}
			if !layer1[passageblock].solid || layer1[passageblock].stairs {
				break
			}
			passagelength = rInt(3, 7)
			for a := 0; a < passagelength; a++ {
				fillfloor(passageblock)
				fillsecret(passageblock)
				passageblock += mapw
				if !layer1[passageblock].solid || layer1[passageblock].stairs {
					findpassagewall(passageblock, 4)
					break
				}
			}
			if !layer1[passageblock].solid || layer1[passageblock].stairs {
				break
			}
		}
	case 3:
		for {
			passagelength := rInt(3, 7)
			for a := 0; a < passagelength; a++ {
				fillfloor(passageblock)
				fillsecret(passageblock)
				passageblock++
				if !layer1[passageblock].solid || layer1[passageblock].stairs {
					findpassagewall(passageblock, 4)
					break
				}
			}
			if !layer1[passageblock].solid || layer1[passageblock].stairs {
				break
			}
			passagelength = rInt(3, 7)
			for a := 0; a < passagelength; a++ {
				fillfloor(passageblock)
				fillsecret(passageblock)
				passageblock -= mapw
				if !layer1[passageblock].solid || layer1[passageblock].stairs {
					findpassagewall(passageblock, 4)
					break
				}
			}
			if !layer1[passageblock].solid || layer1[passageblock].stairs {
				break
			}
		}
	}

	layer2[lootboxblok].img = imglootbox
	layer2[lootboxblok].activ = true
	layer2[lootboxblok].solid = true
	layer2[lootboxblok].fade = 1.0
	layer2[lootboxblok].color = randomcolor()

}
func revealsecret() { // MARK: revealsecret

	for a := 0; a < len(layer1); a++ {

		if layer1[a].secret {

			layer1[a].img = currentfloor
			layer1[a].fade = 0.2
			layer1[a].secret = false
			layer1[a].color = randomgreen()

		}

	}

	secretrevealed = true
}
func createloot() { // MARK: createloot

	/*
	 0 magic
	 1 weapons
	 2 armor
	 3 potions
	 4 secret
	*/
	for a := 0; a < 101; a++ {

		magicitems[a].activ = true
		weaponsitems[a].activ = true
		armoritems[a].activ = true
		potionsitems[a].activ = true
		secretitems[a].activ = true

		magicitems[a].color = randomcolor()
		weaponsitems[a].color = randomcolor()
		armoritems[a].color = randomcolor()
		potionsitems[a].color = randomcolor()
		secretitems[a].color = randomcolor()

		magicitems[a].img = chooseimg("magic")
		weaponsitems[a].img = chooseimg("weapon")
		armoritems[a].img = chooseimg("armor")
		potionsitems[a].img = chooseimg("potion")
		secretitems[a].img = chooseimg("secret")

		magicitems[a].desc = "magic item"
		weaponsitems[a].desc = "weapon item"
		armoritems[a].desc = "armor item"
		potionsitems[a].desc = "potion item"
		secretitems[a].desc = "secret item"

		magicitems[a].fade = 1.0
		weaponsitems[a].fade = 1.0
		armoritems[a].fade = 1.0
		potionsitems[a].fade = 1.0
		secretitems[a].fade = 1.0

		magicitems[a].stat1 = rolldice()
		weaponsitems[a].stat1 = rolldice()
		armoritems[a].stat1 = rolldice()
		potionsitems[a].stat1 = rolldice()
		secretitems[a].stat1 = rolldice()

		magicitems[a].stat2 = rolldice()
		weaponsitems[a].stat2 = rolldice()
		armoritems[a].stat2 = rolldice()
		potionsitems[a].stat2 = rolldice()
		secretitems[a].stat2 = rolldice()

		magicitems[a].stat3 = rolldice()
		weaponsitems[a].stat3 = rolldice()
		armoritems[a].stat3 = rolldice()
		potionsitems[a].stat3 = rolldice()
		secretitems[a].stat3 = rolldice()

	}

}
func createnemy() blokenemy { // MARK: createnemy
	newenemy := blokenemy{}
	if level > 80 {
		newenemy.activ = true
		choose := rInt(1, 4)
		if choose == 1 {
			newenemy.img = imgenemy1
		} else if choose == 2 {
			newenemy.img = imgenemy2
		} else if choose == 3 {
			newenemy.img = imgenemy3
		}
		newenemy.movetype = rInt(1, 4)
		newenemy.color = randomcolor()
		newenemy.fade = 1.0
		newenemy.rotation = 45.0
		newenemy.leftright = flipcoin()
		newenemy.desc = "enemy text here"
	} else if level < 81 && level > 60 {
		newenemy.activ = true
		choose := rInt(1, 4)
		if choose == 1 {
			newenemy.img = imgenemy4
		} else if choose == 2 {
			newenemy.img = imgenemy5
		} else if choose == 3 {
			newenemy.img = imgenemy6
		}
		newenemy.color = randomcolor()
		newenemy.fade = 1.0
		newenemy.rotation = 45.0
		newenemy.leftright = flipcoin()
		newenemy.desc = "enemy text here"
	} else if level < 61 && level > 40 {
		newenemy.activ = true
		choose := rInt(1, 4)
		if choose == 1 {
			newenemy.img = imgenemy7
		} else if choose == 2 {
			newenemy.img = imgenemy8
		} else if choose == 3 {
			newenemy.img = imgenemy9
		}
		newenemy.color = randomcolor()
		newenemy.fade = 1.0
		newenemy.rotation = 45.0
		newenemy.leftright = flipcoin()
		newenemy.desc = "enemy text here"
	} else if level < 41 && level > 20 {
		newenemy.activ = true
		choose := rInt(1, 4)
		if choose == 1 {
			newenemy.img = imgenemy10
		} else if choose == 2 {
			newenemy.img = imgenemy11
		} else if choose == 3 {
			newenemy.img = imgenemy12
		}
		newenemy.color = randomcolor()
		newenemy.fade = 1.0
		newenemy.rotation = 45.0
		newenemy.leftright = flipcoin()
		newenemy.desc = "enemy text here"
	} else if level < 21 && level > 0 {
		newenemy.activ = true
		choose := rInt(1, 4)
		if choose == 1 {
			newenemy.img = imgenemy11
		} else if choose == 2 {
			newenemy.img = imgenemy12
		} else if choose == 3 {
			newenemy.img = imgenemyskul
		}
		newenemy.color = randomcolor()
		newenemy.fade = 1.0
		newenemy.rotation = 45.0
		newenemy.leftright = flipcoin()
		newenemy.desc = "enemy text here"
	}

	return newenemy

}
func chooseimg(itemtype string) rl.Rectangle { // MARK: chooseimg

	choose := rInt(1, 9)
	chooseimg := rl.NewRectangle(0, 0, 0, 0)

	switch itemtype {
	case "armor":
		switch choose {
		case 1:
			chooseimg = imgarmor1
		case 2:
			chooseimg = imgarmor2
		case 3:
			chooseimg = imgarmor3
		case 4:
			chooseimg = imgarmor4
		case 5:
			chooseimg = imgarmor5
		case 6:
			chooseimg = imgarmor6
		case 7:
			chooseimg = imgarmor7
		case 8:
			chooseimg = imgarmor8
		}
	case "weapon":
		switch choose {
		case 1:
			chooseimg = imgweapon1
		case 2:
			chooseimg = imgweapon2
		case 3:
			chooseimg = imgweapon3
		case 4:
			chooseimg = imgweapon4
		case 5:
			chooseimg = imgweapon5
		case 6:
			chooseimg = imgweapon6
		case 7:
			chooseimg = imgweapon7
		case 8:
			chooseimg = imgweapon8
		}
	case "magic":
		switch choose {
		case 1:
			chooseimg = imgmagic1
		case 2:
			chooseimg = imgmagic2
		case 3:
			chooseimg = imgmagic3
		case 4:
			chooseimg = imgmagic4
		case 5:
			chooseimg = imgmagic5
		case 6:
			chooseimg = imgmagic6
		case 7:
			chooseimg = imgmagic7
		case 8:
			chooseimg = imgmagic8
		}
	case "potion":
		switch choose {
		case 1:
			chooseimg = imgpotion1
		case 2:
			chooseimg = imgpotion2
		case 3:
			chooseimg = imgpotion3
		case 4:
			chooseimg = imgpotion4
		case 5:
			chooseimg = imgpotion5
		case 6:
			chooseimg = imgpotion6
		case 7:
			chooseimg = imgpotion7
		case 8:
			chooseimg = imgpotion8
		}
	case "secret":
		switch choose {
		case 1:
			chooseimg = imgsecret1
		case 2:
			chooseimg = imgsecret2
		case 3:
			chooseimg = imgsecret3
		case 4:
			chooseimg = imgsecret4
		case 5:
			chooseimg = imgsecret5
		case 6:
			chooseimg = imgsecret6
		case 7:
			chooseimg = imgsecret7
		case 8:
			chooseimg = imgsecret8
		}
	}
	return chooseimg
}
func fillwall(roomblok int) { // MARK: fillwall
	layer1[roomblok].activ = true
	layer1[roomblok].solid = true
	layer1[roomblok].img = currentwall
	layer1[roomblok].color = randomorange()
	layer1[roomblok].fade = 1.0
	layer1[roomblok].desc = "wall text here"
}
func fillsecret(roomblok int) {
	layer1[roomblok].secret = true
	layer1[roomblok].img = currentwall
	layer1[roomblok].color = randomorange()
	layer1[roomblok].fade = 1.0
}
func fillsecretwall(roomblok int) { // MARK: fillsecretwall
	layer1[roomblok].activ = true
	layer1[roomblok].solid = true
	layer1[roomblok].secret = true
	layer1[roomblok].img = currentwall
	layer1[roomblok].color = randombluedark()
	layer1[roomblok].fade = 1.0
	layer1[roomblok].desc = "wall text here"
}
func fillfloor(roomblok int) { // MARK: fillfloor

	layer1[roomblok].activ = true
	layer1[roomblok].solid = false
	layer1[roomblok].img = currentfloor
	if rolldice()+rolldice() == 12 {
		choose := rInt(1, 10)
		switch choose {
		case 1:
			layer1[roomblok].img = tilegrounddetail1
		case 2:
			layer1[roomblok].img = tilegrounddetail2
		case 3:
			layer1[roomblok].img = tilegrounddetail3
		case 4:
			layer1[roomblok].img = tilegrounddetail4
		case 5:
			layer1[roomblok].img = tilegrounddetail5
		case 6:
			layer1[roomblok].img = tilegrounddetail6
		case 7:
			layer1[roomblok].img = tilegrounddetail7
		case 8:
			layer1[roomblok].img = tilegrounddetail8
		case 9:
			layer1[roomblok].img = tilegrounddetail9
		}
	}
	layer1[roomblok].color = randomgreen()
	layer1[roomblok].fade = 0.2
	layer1[roomblok].desc = "floor text here"

	if rolldice()+rolldice()+rolldice() == 18 {
		layereextras[roomblok].activ = true
		choose := rInt(1, 11)
		switch choose {
		case 1:
			layereextras[roomblok].img = imgextra1
		case 2:
			layereextras[roomblok].img = imgextra2
		case 3:
			layereextras[roomblok].img = imgextra3
		case 4:
			layereextras[roomblok].img = imgextra4
		case 5:
			layereextras[roomblok].img = imgextra5
		case 6:
			layereextras[roomblok].img = imgextra6
		case 7:
			layereextras[roomblok].img = imgextra7
		case 8:
			layereextras[roomblok].img = imgextra8
		case 9:
			layereextras[roomblok].img = imgextra9
		case 10:
			layereextras[roomblok].img = imgextra10
		}
		layereextras[roomblok].color = rl.LightGray
		layereextras[roomblok].fade = 0.4
		layereextras[roomblok].desc = "extras text here"

	}
}
func openlootbox(blocknumber int) { // MARK: openlootbox

}
func findpassagewall(roomblock, direction int) {

	check := roomblock
	count := 0
	switch direction {
	case 4:
		for {
			check--
			if !layer1[check].solid && layer1[check-mapw].solid && layer1[check+mapw].solid {
				fillsecretwall(check)
				count = 10
			}
			count++
			if count >= 10 {
				break
			}
		}
	}
}
func placeenemies() { // MARK: placeloot
	for {
		enemyblock := rInt(mapw+1, (mapa - (mapw + 1)))
		if !layer1[enemyblock].solid && layer1[enemyblock].activ && enemyblock != player && !layer1[enemyblock].stairs {
			layerenemies[enemyblock] = createnemy()
			break
		}
	}
}
func placeloot() { // MARK: placeloot

	// place magic
	for {
		lootblock := rInt(mapw+1, (mapa - (mapw + 1)))
		if !layer1[lootblock].solid && layer1[lootblock].activ && lootblock != player && !layer1[lootblock].stairs {
			magicitems[level].blocknumber = lootblock
			break
		}
	}
	// place weapons
	for {
		lootblock := rInt(mapw+1, (mapa - (mapw + 1)))
		if !layer1[lootblock].solid && layer1[lootblock].activ && lootblock != player && !layer1[lootblock].stairs && lootblock != magicitems[level].blocknumber {
			weaponsitems[level].blocknumber = lootblock
			break
		}
	}
	// place armor
	for {
		lootblock := rInt(mapw+1, (mapa - (mapw + 1)))
		if !layer1[lootblock].solid && layer1[lootblock].activ && lootblock != player && !layer1[lootblock].stairs && lootblock != magicitems[level].blocknumber && lootblock != weaponsitems[level].blocknumber {
			armoritems[level].blocknumber = lootblock
			break
		}
	}
	// place potions
	for {
		lootblock := rInt(mapw+1, (mapa - (mapw + 1)))
		if !layer1[lootblock].solid && layer1[lootblock].activ && lootblock != player && !layer1[lootblock].stairs && lootblock != magicitems[level].blocknumber && lootblock != weaponsitems[level].blocknumber && lootblock != armoritems[level].blocknumber {
			potionsitems[level].blocknumber = lootblock
			break
		}
	}
	// place secrets
	for {
		lootblock := rInt(mapw+1, (mapa - (mapw + 1)))
		if !layer1[lootblock].solid && layer1[lootblock].activ && lootblock != player && !layer1[lootblock].stairs && lootblock != magicitems[level].blocknumber && lootblock != weaponsitems[level].blocknumber && lootblock != armoritems[level].blocknumber && lootblock != potionsitems[level].blocknumber {
			secretitems[level].blocknumber = lootblock
			break
		}
	}

}
func placeplayer(stairsblock int) { // MARK: placeplayer
	for {
		player = rInt(mapw+1, (mapa - (mapw + 1)))
		if !layer1[player].solid && !layer1[player].stairs {

			h, v := getblokhv(player)
			h2, v2 := getblokhv(stairsblock)

			hdiff := int(math.Abs(float64(h2 - h)))
			vdiff := int(math.Abs(float64(v2 - v)))

			if hdiff < playermovesmax || vdiff < playermovesmax {
				break
			}

		}
	}
}
func moveplayer(direction int) { // MARK: moveplayer

	playerhasmoved := false

	switch direction {
	case 1:
		if !layer1[(player-1)+mapw].solid && layer1[(player-1)+mapw].activ {
			player--
			player += mapw
			playermoves--
			playerhasmoved = true
		}
	case 2:
		if !layer1[(player)+mapw].solid && layer1[(player)+mapw].activ {
			player += mapw
			playermoves--
			playerhasmoved = true
		}
	case 3:
		if !layer1[(player+1)+mapw].solid && layer1[(player+1)+mapw].activ {
			player++
			player += mapw
			playermoves--
			playerhasmoved = true
		}
	case 4:
		if !layer1[(player-1)].solid && layer1[(player-1)].activ {
			player--
			playermoves--
			playerhasmoved = true
		}
	case 6:
		if !layer1[(player+1)].solid && layer1[(player+1)].activ {
			player++
			playermoves--
			playerhasmoved = true
		}
	case 7:
		if !layer1[(player-1)-mapw].solid && layer1[(player-1)-mapw].activ {
			player--
			player -= mapw
			playermoves--
			playerhasmoved = true
		}
	case 8:
		if !layer1[(player)-mapw].solid && layer1[(player)-mapw].activ {
			player -= mapw
			playermoves--
			playerhasmoved = true
		}
	case 9:
		if !layer1[(player+1)-mapw].solid && layer1[(player+1)-mapw].activ {
			player++
			player -= mapw
			playermoves--
			playerhasmoved = true
		}
	}

	if playerhasmoved {
		movenemies()
		playermovestotal++
	}

}
func movenemies() { // MARK: movenemies

	// left up
	for a := 0; a < len(layerenemies); a++ {
		if layerenemies[a].activ {
			enemyh, enemyv := getblokhv(a)
			playerh, playerv := getblokhv(player)
			switch layerenemies[a].movetype {
			case 1:
				if layerenemies[a].movecounter != 4 {
					layerenemies[a].movecounter++
				} else {
					if math.Abs(float64(playerv-enemyv)) > math.Abs(float64(playerh-enemyh)) {
						if playerv <= enemyv {
							enemyleft(a)
						}
					} else {
						if playerh <= enemyh {
							enemyup(a)
						}
					}
					layerenemies[a].movecounter = 0
				}
			case 2:
				if layerenemies[a].movecounter != 4 {
					layerenemies[a].movecounter++
				} else {
					if math.Abs(float64(playerv-enemyv)) > math.Abs(float64(playerh-enemyh)) {
						if playerv <= enemyv {
							enemyleft(a)
						}
					} else {

						if playerh <= enemyh {
							enemyup(a)
						}
					}
					layerenemies[a].movecounter = 0
				}
			case 3:
				if layerenemies[a].movecounter != 4 {
					layerenemies[a].movecounter++
				} else {
					if math.Abs(float64(playerv-enemyv)) > math.Abs(float64(playerh-enemyh)) {
						if playerv <= enemyv {
							enemyleft(a)
						}
					} else {
						if playerh <= enemyh {
							enemyup(a)
						}
					}
				}
				layerenemies[a].movecounter = 0
			}
		}
	}
	// right down
	for a := len(layerenemies) - 1; a > 0; a-- {
		if layerenemies[a].activ {
			enemyh, enemyv := getblokhv(a)
			playerh, playerv := getblokhv(player)
			switch layerenemies[a].movetype {
			case 1:
				if layerenemies[a].movecounter != 4 {
					layerenemies[a].movecounter++
				} else {
					if math.Abs(float64(playerv-enemyv)) > math.Abs(float64(playerh-enemyh)) {
						if playerv > enemyv {
							enemyright(a)
						}
					} else {
						if playerh > enemyh {
							enemydown(a)
						}
					}
					layerenemies[a].movecounter = 0
				}
			case 2:
				if layerenemies[a].movecounter != 4 {
					layerenemies[a].movecounter++
				} else {
					if math.Abs(float64(playerv-enemyv)) > math.Abs(float64(playerh-enemyh)) {
						if playerv > enemyv {
							enemyright(a)
						}
					} else {
						if playerh > enemyh {
							enemydown(a)
						}
					}
					layerenemies[a].movecounter = 0
				}
			case 3:
				if layerenemies[a].movecounter != 4 {
					layerenemies[a].movecounter++
				} else {
					if math.Abs(float64(playerv-enemyv)) > math.Abs(float64(playerh-enemyh)) {
						if playerv > enemyv {
							enemyright(a)
						}
					} else {
						if playerh > enemyh {
							enemydown(a)
						}
					}
				}
				layerenemies[a].movecounter = 0
			}
		}
	}
}
func enemyleft(blocknumber int) {
	if !layer1[blocknumber-1].solid {
		layerenemies[blocknumber-1] = layerenemies[blocknumber]
		layerenemies[blocknumber] = blokenemy{}
	}
}
func enemyright(blocknumber int) {
	if !layer1[blocknumber+1].solid {
		layerenemies[blocknumber+1] = layerenemies[blocknumber]
		layerenemies[blocknumber] = blokenemy{}
	}

}
func enemyup(blocknumber int) {
	if !layer1[blocknumber-mapw].solid {
		layerenemies[blocknumber-mapw] = layerenemies[blocknumber]
		layerenemies[blocknumber] = blokenemy{}
	}

}
func enemydown(blocknumber int) {
	if !layer1[blocknumber+mapw].solid {
		layerenemies[blocknumber+mapw] = layerenemies[blocknumber]
		layerenemies[blocknumber] = blokenemy{}
	}

}
func animateenemies() { // MARK: animateenemies
	for a := 0; a < len(layerenemies); a++ {
		if layerenemies[a].activ {
			if rolldice()+rolldice() == 12 {
				layerenemies[a].leftright = flipcoin()
			}
			if layerenemies[a].leftright {
				layerenemies[a].rotation += rFloat32(3, 11)
			} else {
				layerenemies[a].rotation -= rFloat32(3, 11)
			}
		}
	}
}
func timers() { // MARK: timers

	// onoff6
	if onoff6 {
		if animatelootboxon {
			if imglootbox.X == 2 {
				imglootbox.X = 26
				layer2[lootboxbloknumber].img = imglootbox
			} else if imglootbox.X == 26 {
				imglootbox.X = 50
				layer2[lootboxbloknumber].img = imglootbox
			}
		}

	}

	// onoff3
	if onoff3 {

		animateenemies()

		randomfade2 = 0.2
		randomfade1 = rF32(0.2, 0.6)
		mouserecl += 2
		mouserecx++
		mouserecy++
		if mouserecl > 40 {
			mouserecl = 34
			mouserecx = 0
			mouserecy = 0
		}

	} else {
		randomfade2 = 0.0
	}

	if framecount%2 == 0 {
		if onoff2 {
			onoff2 = false
		} else {
			onoff2 = true
		}
	}
	if framecount%3 == 0 {
		if onoff3 {
			onoff3 = false
		} else {
			onoff3 = true
		}
	}
	if framecount%6 == 0 {
		if onoff6 {
			onoff6 = false
		} else {
			onoff6 = true
		}
	}
	if framecount%10 == 0 {
		if onoff10 {
			onoff10 = false
		} else {
			onoff10 = true
		}
	}
	if framecount%15 == 0 {
		if onoff15 {
			onoff15 = false
		} else {
			onoff15 = true
		}
	}
	if framecount%30 == 0 {
		if onoff30 {
			onoff30 = false
		} else {
			onoff30 = true
		}
	}

}
func input() { // MARK: input
	if rl.IsKeyPressed(rl.KeyF3) {
		if nighton {
			nighton = false
		} else {
			nighton = true
		}
	}
	if rl.IsKeyPressed(rl.KeyF2) {
		if statson {
			statson = false
		} else {
			statson = true
		}
	}
	if rl.IsKeyPressed(rl.KeyF1) {
		if fxon {
			fxon = false
		} else {
			fxon = true
		}
	}

	if rl.IsKeyPressed(rl.KeyKp1) {
		moveplayer(1)
	}
	if rl.IsKeyPressed(rl.KeyKp2) {
		moveplayer(2)
	}
	if rl.IsKeyPressed(rl.KeyKp3) {
		moveplayer(3)
	}
	if rl.IsKeyPressed(rl.KeyKp4) {
		moveplayer(4)
	}
	if rl.IsKeyPressed(rl.KeyKp6) {
		moveplayer(6)
	}
	if rl.IsKeyPressed(rl.KeyKp7) {
		moveplayer(7)
	}
	if rl.IsKeyPressed(rl.KeyKp8) {
		moveplayer(8)
	}
	if rl.IsKeyPressed(rl.KeyKp9) {
		moveplayer(9)
	}

	if rl.IsKeyPressed(rl.KeyKp0) {
		if gridon {
			gridon = false
		} else {
			gridon = true
		}
	}
	if rl.IsKeyPressed(rl.KeyKpDivide) {
		if colorizeon {
			colorizeon = false
		} else {
			colorizeon = true
		}
	}
	if rl.IsKeyPressed(rl.KeyKpMultiply) {
		if centerlineson {
			centerlineson = false
		} else {
			centerlineson = true
		}
	}

	if rl.IsKeyPressed(rl.KeyKpDecimal) {
		if debugon {
			debugon = false
		} else {
			debugon = true
		}
	}

	if rl.IsKeyPressed(rl.KeyKpAdd) {
		if camera.Zoom == 1.0 {
			camera.Zoom = 2.0
		} else if camera.Zoom == 2.0 {
			camera.Zoom = 3.0
		} else if camera.Zoom == 3.0 {
			camera.Zoom = 4.0
		}
	}
	if rl.IsKeyPressed(rl.KeyKpSubtract) {
		if camera.Zoom == 2.0 {
			camera.Zoom = 1.0
		} else if camera.Zoom == 3.0 {
			camera.Zoom = 2.0
		} else if camera.Zoom == 4.0 {
			camera.Zoom = 3.0
		}
	}

}
func clearmaps() { // MARK: clearmaps

	for a := 0; a < len(layer1); a++ {
		layer1[a] = blok{}
		layer2[a] = blok{}
		layerenemies[a] = blokenemy{}
		layereextras[a] = blok{}
	}
}
func resetnewlevel() { // MARK: resetnewlevel

	imglootbox.X = 2
	animatelootboxon = false
	playermoves = playermovesmax
	secretrevealed = false

}
func choosewallblok() rl.Rectangle { // MARK: choosewallblok

	wallblock := rl.NewRectangle(0, 0, 0, 0)

	choose := rInt(1, 20)

	switch choose {
	case 1:
		wallblock = tilewall1
	case 2:
		wallblock = tilewall2
	case 3:
		wallblock = tilewall3
	case 4:
		wallblock = tilewall4
	case 5:
		wallblock = tilewall5
	case 6:
		wallblock = tilewall6
	case 7:
		wallblock = tilewall7
	case 8:
		wallblock = tilewall8
	case 9:
		wallblock = tilewall9
	case 10:
		wallblock = tilewall10
	case 11:
		wallblock = tilewall11
	case 12:
		wallblock = tilewall12
	case 13:
		wallblock = tilewall13
	case 14:
		wallblock = tilewall14
	case 15:
		wallblock = tilewall15
	case 16:
		wallblock = tilewall16
	case 17:
		wallblock = tilewall17
	case 18:
		wallblock = tilewall18
	case 19:
		wallblock = tilewall19
	case 20:
		wallblock = tilewall20
	case 21:
		wallblock = tilewall21
	case 22:
		wallblock = tilewall22
	case 23:
		wallblock = tilewall23
	case 24:
		wallblock = tilewall24
	case 25:
		wallblock = tilewall25
	case 26:
		wallblock = tilewall26
	case 27:
		wallblock = tilewall27
	case 28:
		wallblock = tilewall28
	case 29:
		wallblock = tilewall29
	case 30:
		wallblock = tilewall30
	case 31:
		wallblock = tilewall31
	case 32:
		wallblock = tilewall32
	case 33:
		wallblock = tilewall33
	case 34:
		wallblock = tilewall34
	}

	return wallblock
}
func choosefloorblok() rl.Rectangle { // MARK: choosefloorblok

	floorblok := rl.NewRectangle(0, 0, 0, 0)

	choose := rInt(1, 20)

	switch choose {
	case 1:
		floorblok = tileground1
	case 2:
		floorblok = tileground2
	case 3:
		floorblok = tileground3
	case 4:
		floorblok = tileground4
	case 5:
		floorblok = tileground5
	case 6:
		floorblok = tileground6
	case 7:
		floorblok = tileground7
	case 8:
		floorblok = tileground8
	case 9:
		floorblok = tileground9
	case 10:
		floorblok = tileground10
	case 11:
		floorblok = tileground11
	case 12:
		floorblok = tileground12
	case 13:
		floorblok = tileground13
	case 14:
		floorblok = tileground14
	case 15:
		floorblok = tileground15
	case 16:
		floorblok = tileground16
	case 17:
		floorblok = tileground17
	case 18:
		floorblok = tileground18
	case 19:
		floorblok = tileground19

	}

	return floorblok
}
func endblockhv(block int) bool { // MARK: endblockhv
	end := false
	h, v := block/mapw, block%mapw
	if h < 2 || h > boundaryright || v < 2 || v > boundarybottom {
		end = true
	}
	return end
}
func getblokhv(block int) (int, int) { // MARK: getblokhv
	h, v := block/mapw, block%mapw
	return h, v
}
func getmouseblock() { // MARK: getmouseblock

	if mousepos.X >= 64 && mousepos.X <= 1248 && mousepos.Y >= 56 && mousepos.Y <= 696 {
		mousegridon = true
		//	ychange := 32
		count := 0
		countblok := 0
		x := float32(64)
		y := float32(56)
		for a := 0; a < mapa; a++ {

			if mousepos.X >= x && mousepos.X <= x+32 && mousepos.Y >= y && mousepos.Y <= y+32 {
				mouseblok = countblok
			}

			x += 32
			count++
			countblok++
			if count == mapw {
				count = 0
				x = 64
				y += 32
			}

		}

	} else {
		mousegridon = false
	}

}
func test() { // MARK: test

	player = centerblok
	playerblok.img = imgplayer1
	playerblok.color = randomcolor()
	playerblok.fade = 1.0
	playerblok.desc = "player text here"

	createdunjin()
	createloot()

}
func debug() { // MARK: debug
	rl.DrawRectangle(monw-300, 0, 500, monw, rl.Fade(rl.Black, 0.8))
	rl.DrawFPS(monw-290, monh-100)

	gridatext := strconv.Itoa(grida)
	mousextext := fmt.Sprintf("%g", mousepos.X)
	mouseytext := fmt.Sprintf("%g", mousepos.Y)
	mapwtext := strconv.Itoa(mapw)
	maphtext := strconv.Itoa(maph)
	mapatext := strconv.Itoa(mapa)
	mousebloktext := strconv.Itoa(mouseblok)
	playertext := strconv.Itoa(player)

	rl.DrawText(gridatext, monw-290, 10, 10, rl.White)
	rl.DrawText("grida", monw-150, 10, 10, rl.White)
	rl.DrawText(mousextext, monw-290, 20, 10, rl.White)
	rl.DrawText("mouseX", monw-150, 20, 10, rl.White)
	rl.DrawText(mouseytext, monw-290, 30, 10, rl.White)
	rl.DrawText("mouseY", monw-150, 30, 10, rl.White)
	rl.DrawText(mapwtext, monw-290, 40, 10, rl.White)
	rl.DrawText("mapw", monw-150, 40, 10, rl.White)
	rl.DrawText(maphtext, monw-290, 50, 10, rl.White)
	rl.DrawText("maph", monw-150, 50, 10, rl.White)
	rl.DrawText(mapatext, monw-290, 60, 10, rl.White)
	rl.DrawText("mapa", monw-150, 60, 10, rl.White)
	rl.DrawText(mousebloktext, monw-290, 70, 10, rl.White)
	rl.DrawText("mouseblok", monw-150, 70, 10, rl.White)
	rl.DrawText(playertext, monw-290, 80, 10, rl.White)
	rl.DrawText("player", monw-150, 80, 10, rl.White)

}
func setinitialvalues() { // MARK: initialize

	mouseon = true

	test()

}
func main() { // MARK: main
	rand.Seed(time.Now().UnixNano()) // random numbers
	rl.SetTraceLogLevel(rl.LogError) // hides info window
	rl.InitWindow(monw, monh, "setscreen")
	setscreen()
	rl.CloseWindow()
	setinitialvalues()
	raylib()
}
func setscreen() { // MARK: setscreen

	rl.SetWindowSize(monw, monh)

	camera.Zoom = 2.0
	camera.Target.X = 0
	camera.Target.Y = 0

	camera2x.Zoom = 2.0
	camera4x.Zoom = 4.0
	camera8x.Zoom = 8.0
}

// random colors https://www.rapidtables.com/web/color/RGB_Color.html
func randomgrey() rl.Color {
	color := rl.NewColor(uint8(rInt(105, 192)), uint8(rInt(105, 192)), uint8(rInt(105, 192)), 255)
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

// random numbers
func rF32(min, max float32) float32 {
	return (rand.Float32() * (max - min)) + min

}
func rInt(min, max int) int {
	return rand.Intn(max-min) + min
}
func rInt32(min, max int) int32 {
	a := int32(rand.Intn(max-min) + min)
	return a
}
func rFloat32(min, max int) float32 {
	a := float32(rand.Intn(max-min) + min)
	return a
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
