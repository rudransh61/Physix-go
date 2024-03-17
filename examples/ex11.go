package main

import (
	"physix/pkg/polygon"
	"physix/pkg/vector"
	"physix/internal/collision"
	"fmt"
	// "physix/internal/physics"
)

func main() {
	// Define two polygons
	polygon1 := polygon.Polygon{
		Vertices: []vector.Vector{{X: 100, Y: 100}, {X: 200, Y: 100}, {X: 200, Y: 200}, {X: 100, Y: 200}},
	}
	polygon2 := polygon.Polygon{
		Vertices: []vector.Vector{{X: 150, Y: 150}, {X: 250, Y: 150}, {X: 250, Y: 250}, {X: 150, Y: 250}},
	}

	//\(\(250,500\),\(250,300\),\(250,250\),\(200,250\))
	polygon3 := polygon.Polygon{
		Vertices: []vector.Vector{{X: 250, Y: 500}, {X: 250, Y: 300}, {X: 250, Y: 250}, {X: 200, Y: 250}},
	}
	// Check if the polygons are colliding
	if collision.PolygonCollision(polygon1, polygon2) {
		fmt.Println("Polygons 1 2 are colliding")
	} else {
		fmt.Println("Polygons 1 2 are not colliding")
	}
	// Check if the polygons are colliding
	if collision.PolygonCollision(polygon1, polygon3) {
		fmt.Println("Polygons 1 3 are colliding")
	} else {
		fmt.Println("Polygons 1 3 are not colliding")
	}
	// Check if the polygons are colliding
	if collision.PolygonCollision(polygon2, polygon3) {
		fmt.Println("Polygons 2 3 are colliding")
	} else {
		fmt.Println("Polygons 2 3 are not colliding")
	}
}