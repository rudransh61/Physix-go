package broadphase

import (
	"github.com/rudransh61/Physix-go/pkg/vector"
)

// SpatialHash represents a spatial hash for broad-phase collision detection.
type SpatialHash struct {
	cellSize float64
	cells    map[int]map[int][]interface{}
	width    float64
	height   float64
}

// NewSpatialHash creates a new SpatialHash with the given cell size, width, and height.
func NewSpatialHash(cellSize, width, height float64) *SpatialHash {
	return &SpatialHash{
		cellSize: cellSize,
		cells:    make(map[int]map[int][]interface{}),
		width:    width,
		height:   height,
	}
}

// Clear clears the spatial hash.
func (sh *SpatialHash) Clear() {
	sh.cells = make(map[int]map[int][]interface{})
}

// Add adds an object to the spatial hash.
func (sh *SpatialHash) Add(obj interface{}, position vector.Vector) {
	cellX := int(position.X / sh.cellSize)
	cellY := int(position.Y / sh.cellSize)
	if sh.cells[cellX] == nil {
		sh.cells[cellX] = make(map[int][]interface{})
	}
	if sh.cells[cellX][cellY] == nil {
		sh.cells[cellX][cellY] = make([]interface{}, 0)
	}
	sh.cells[cellX][cellY] = append(sh.cells[cellX][cellY], obj)
}

// Query returns nearby objects around the given position from the spatial hash.
func (sh *SpatialHash) Query(position vector.Vector) []interface{} {
	cellX := int(position.X / sh.cellSize)
	cellY := int(position.Y / sh.cellSize)
	nearbyObjects := make([]interface{}, 0)
	for x := cellX - 1; x <= cellX+1; x++ {
		for y := cellY - 1; y <= cellY+1; y++ {
			if sh.cells[x] != nil && sh.cells[x][y] != nil {
				nearbyObjects = append(nearbyObjects, sh.cells[x][y]...)
			}
		}
	}
	return nearbyObjects
}
