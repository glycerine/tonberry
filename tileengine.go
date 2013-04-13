package tonberry

import (
	"encoding/json"
	"github.com/zeroshade/Go-SDL/sdl"
	"image"
	"io/ioutil"
)

var (
	maps map[string]mapInfo
)

type mapInfo struct {
	SpriteFile string     `json:"file"`
	THeight    int        `json:"tile_height"`
	TWidth     int        `json:"tile_width"`
	LHeight    uint16     `json:"level_height"`
	LWidth     uint16     `json:"level_width"`
	SpriteNum  int        `json:"tile_sprites"`
	Map        []int      `json:"map"`
	Clips      []sdl.Rect `json:"clips"`
	Tiles      []Tile
	SpriteSurf *sdl.Surface
}

type Tile struct {
	Box  image.Rectangle
	Type int
}

func InitTileMap(fn, mapName string) {
	file, e := ioutil.ReadFile(fn)
	if e != nil {
		panic("Bad Tile File")
	}

	if maps == nil {
		maps = make(map[string]mapInfo)
	}

	var m mapInfo
	err := json.Unmarshal(file, &m)
	if err != nil {
		panic(err)
	}

	x, y := 0, 0

	m.Tiles = make([]Tile, len(m.Map))
	for t := range m.Tiles {
		tileType := m.Map[t]
		if tileType >= 0 && tileType < m.SpriteNum {
			m.Tiles[t] = Tile{
				Box:  image.Rect(x, y, x+m.TWidth, y+m.THeight),
				Type: tileType,
			}
		}

		x += m.TWidth

		if uint16(x) >= m.LWidth {
			x = 0
			y += m.THeight
		}
	}

	m.SpriteSurf = load_image(m.SpriteFile)
	maps[mapName] = m
}

func DrawMap(name string, c Camera, sc Screen) {
	if m, ok := maps[name]; ok {
		m.show(c, sc)
	}
}

func MapHeight(name string) int {
	return int(maps[name].LHeight)
}

func MapWidth(name string) int {
	return int(maps[name].LWidth)
}

func MapTiles(name string) []Tile {
	return maps[name].Tiles
}

func (mi *mapInfo) show(c Camera, sc Screen) {
	cam := sdl.GoRectFromRect(c.Rect)

	for _, t := range mi.Tiles {
		if cam.Overlaps(t.Box) {
			apply_surface_clip(int16(t.Box.Min.X-cam.Min.X),
				int16(t.Box.Min.Y-cam.Min.Y), mi.SpriteSurf,
				sc.Surface, &mi.Clips[t.Type])
		}
	}
}
