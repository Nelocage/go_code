package ebiten2dgame

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Game struct {
}

func (g *Game) Update() error {
	return nil //nil在go中就代表着错误
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "hello world")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func GameMain() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("hello world")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
	//go游戏开发的相关库
	//https://ebiten-zh.vercel.app/tour/hello_world.html#%E7%BC%96%E5%86%99%E4%BB%A3%E7%A0%81
}
