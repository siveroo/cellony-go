package comp

import (
	"cellony/game/menu"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type PositionData struct {
	X, Y float64
}

type VelocityData struct {
	X, Y float64
}

type Activity uint8

const (
	Wandering Activity = iota
	Delivering
	Attacking
	Fleeing
	Mining
)

type CellActivityData struct {
	Activity Activity
}

type SpeedData struct {
	Speed float64
}

type ClockData struct {
	Time     float64
	Cooldown float64
}

type SpriteData struct {
	Sprite  *ebiten.Image
	Z       int32
	Scale   float64
	Opacity float64
}

type VerticesData struct {
	Vertices []ebiten.Vertex
}

type IndicesData struct {
	Indices []uint16
}

const (
	DirtyMask uint8 = 1 << iota
	WallMask        // if set, this tile is a wall, else it's a food
	MarkedMask
)

type GridData struct {
	Grid [][]float32 // doubles down as grid health
	Mask [][]uint8
}

type ImageData struct {
	Img *ebiten.Image
}

type ColorData struct {
	R uint8
	G uint8
	B uint8
}

type HiveData struct {
	SpawnCooldown  float64
	SpawnCount     int
	SpawnCountdown float64
}

type ParentData struct {
	Id donburi.Entity
}

type PheromoneData struct {
	Age          float64
	Intensity    float64
	MaxIntensity float64
	HiveID       donburi.Entity
	Activity     Activity
}

type CellClass uint8

const (
	Wanderer CellClass = iota
	Gatherer
	Soldier
)

type CellData struct {
	Class           CellClass
	Health          float64
	PheromoneChance float64
}

type HUDData struct {
	Menu *menu.Menu
}

// Components
var Position = donburi.NewComponentType[PositionData]()
var Velocity = donburi.NewComponentType[VelocityData]()
var CellActivity = donburi.NewComponentType[CellActivityData]()
var Speed = donburi.NewComponentType[SpeedData]()
var Sprite = donburi.NewComponentType[SpriteData]()
var Color = donburi.NewComponentType[ColorData]()
var Vertices = donburi.NewComponentType[VerticesData]()
var Indices = donburi.NewComponentType[IndicesData]()
var Grid = donburi.NewComponentType[GridData]()
var Image = donburi.NewComponentType[ImageData]()
var Hive = donburi.NewComponentType[HiveData]()
var Parent = donburi.NewComponentType[ParentData]()
var Pheromone = donburi.NewComponentType[PheromoneData]()
var Cell = donburi.NewComponentType[CellData]()
var HUD = donburi.NewComponentType[HUDData]()
