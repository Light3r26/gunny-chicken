package main
//TODO 

import (
	"log"
	"math"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	FPS                     = 60.0
	DELTA_TIME              = 1.0 / FPS
	N_SEGMENTI_WIDTH        = 10
	N_SEGMENTI_HEIGHT       = 5
	TILE_PIXEL              = 64.0
	BORDER_DISTANCE         = TILE_PIXEL / 2
	TILES_SPEED_COEFFICIENT = TILE_PIXEL * DELTA_TIME
	CHICKEN_SPEED           = 0.75 * TILES_SPEED_COEFFICIENT
	WALL_CREATION_SPEED     = 5 * TILE_PIXEL
	WALL_DESTRUCTION        = 20 * TILE_PIXEL

	LOGICAL_SCREEN_WIDTH  = TILE_PIXEL*N_SEGMENTI_WIDTH + BORDER_DISTANCE*2
	LOGICAL_SCREEN_HEIGHT = TILE_PIXEL*N_SEGMENTI_HEIGHT + BORDER_DISTANCE*4
	HORIZONTAL            = 1
	VERTICAL              = 0


)

var (
	gunny_chicken_sprite *ebiten.Image
	bg_sprite            *ebiten.Image
	wall_sprite          *ebiten.Image
	CHICKEN_HITBOX       = HitBox{26, 26}
	WALL_HITBOX          = HitBox{8, 8}
)

type Vector struct {
	X float64
	Y float64
}
type HitBox struct {
	Width  float64
	Height float64
}
type Wall struct {
	CollisionEntity
	NSegments   int
	Finished    bool
	Orientation int
}
type Chicken struct {
	CollisionEntity
	Orientation Vector
}
type CollisionEntity struct {
	Vector
	HitBox
}
func (v Vector) test() float64{
	return v.X + v.Y
}

func (w Wall) DrawWall(screen *ebiten.Image) {
	OpWall := &ebiten.DrawImageOptions{}
	OpWall.GeoM.Scale(TILE_PIXEL/8,1) //temp
	OpWall.GeoM.Translate(w.X, w.Y)
	var ShiftVector Vector
	if w.Orientation == VERTICAL {
		OpWall.GeoM.Rotate(math.Pi / 2)
		ShiftVector.Y = TILE_PIXEL

	} else {
		ShiftVector.X = TILE_PIXEL
	}
	for _ = range w.NSegments {
		screen.DrawImage(wall_sprite, OpWall)
		OpWall.GeoM.Translate(ShiftVector.X, ShiftVector.Y)
	}

}
func (c CollisionEntity) GetVertices() []Vector {
	positions := make([]Vector, 4)
	for i := range 4 {
		positions = append(
			positions, Vector{
				c.X + c.Width*float64(i&2),
				c.Y + c.Height*float64(i&1),
			})

	}
	return positions
}
func (cs CollisionEntity) CheckCollision(ct CollisionEntity) bool {
	ctV := cs.GetVertices()
	for _, point := range ctV {
		if point.X > cs.X && point.Y > cs.Y && point.X < (cs.X+cs.Width) && point.Y < (cs.Y+cs.Height) {
			return true
		}
	}
	return false
}

func load_sprites() {
	var err error
	bg_sprite, _, err = ebitenutil.NewImageFromFile("./sprites/bg.png")
	if err != nil {
		log.Fatal(err)
	}
	gunny_chicken_sprite, _, err = ebitenutil.NewImageFromFile("./sprites/gunny-chicken.png")
	if err != nil {
		log.Fatal(err)
	}
	wall_sprite, _, err = ebitenutil.NewImageFromFile("./sprites/wall.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	Level int
	Timer int
	Points int
	HorizontalWall Wall
	VerticalWall Wall
	Chickens []Chicken
	Cursors Vector
}
func NewChiken() *Chicken{
	return &Chicken{
	}



}
func InitGame() *Game{
	return &Game{
		Level: 1,
		Timer: 60 * FPS,

	}
	
}

func (g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	wall1 := Wall{
		NSegments:   N_SEGMENTI_WIDTH,
		Orientation: HORIZONTAL,

		CollisionEntity: CollisionEntity{

			Vector: Vector{BORDER_DISTANCE,BORDER_DISTANCE*2},
			HitBox: HitBox{N_SEGMENTI_WIDTH*TILE_PIXEL,8},

		},
	}

	bgOp := &ebiten.DrawImageOptions{}
	bgOp.GeoM.Translate(0, 0)
	bgOp.GeoM.Scale(LOGICAL_SCREEN_WIDTH, LOGICAL_SCREEN_HEIGHT)
	chickenOp := &ebiten.DrawImageOptions{}
	chickenOp.GeoM.Translate(0, 0)
	screen.DrawImage(bg_sprite, bgOp)
	wall1.DrawWall(screen)
	screen.DrawImage(gunny_chicken_sprite, chickenOp)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return LOGICAL_SCREEN_WIDTH, LOGICAL_SCREEN_HEIGHT
}

func main() {
	load_sprites()
	ebiten.SetWindowSize(1400, 700)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
