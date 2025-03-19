package main

import (
	"math"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

// Vector3 represents a 3D vector.
type Vector3 struct {
	X, Y, Z float64
}

// Add returns the addition of two vectors.
func (v Vector3) Add(other Vector3) Vector3 {
	return Vector3{v.X + other.X, v.Y + other.Y, v.Z + other.Z}
}

// Sub returns the subtraction of two vectors.
func (v Vector3) Sub(other Vector3) Vector3 {
	return Vector3{v.X - other.X, v.Y - other.Y, v.Z - other.Z}
}

// Scale returns the vector scaled by a scalar value.
func (v Vector3) Scale(s float64) Vector3 {
	return Vector3{v.X * s, v.Y * s, v.Z * s}
}

// Length returns the length (magnitude) of the vector.
func (v Vector3) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Normalize returns the normalized vector (unit vector).
func (v Vector3) Normalize() Vector3 {
	length := v.Length()
	return Vector3{v.X / length, v.Y / length, v.Z / length}
}

// Dot returns the dot product of two vectors.
func (v Vector3) Dot(other Vector3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

// Cross returns the cross product of two vectors.
func (v Vector3) Cross(other Vector3) Vector3 {
	return Vector3{
		v.Y*other.Z - v.Z*other.Y,
		v.Z*other.X - v.X*other.Z,
		v.X*other.Y - v.Y*other.X,
	}
}

// Matrix represents a 3x3 matrix.
type Matrix struct {
	Values [3][3]float64
}

// IdentityMatrix returns the identity matrix.
func IdentityMatrix() Matrix {
	return Matrix{
		Values: [3][3]float64{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		},
	}
}

// RotateX returns a rotation matrix around the x-axis by angle in radians.
func RotateX(angle float64) Matrix {
	sin := math.Sin(angle)
	cos := math.Cos(angle)
	return Matrix{
		Values: [3][3]float64{
			{1, 0, 0},
			{0, cos, -sin},
			{0, sin, cos},
		},
	}
}

// RotateY returns a rotation matrix around the y-axis by angle in radians.
func RotateY(angle float64) Matrix {
	sin := math.Sin(angle)
	cos := math.Cos(angle)
	return Matrix{
		Values: [3][3]float64{
			{cos, 0, sin},
			{0, 1, 0},
			{-sin, 0, cos},
		},
	}
}

// MulVec multiplies the matrix by a vector.
func (m Matrix) MulVec(v Vector3) Vector3 {
	return Vector3{
		X: m.Values[0][0]*v.X + m.Values[0][1]*v.Y + m.Values[0][2]*v.Z,
		Y: m.Values[1][0]*v.X + m.Values[1][1]*v.Y + m.Values[1][2]*v.Z,
		Z: m.Values[2][0]*v.X + m.Values[2][1]*v.Y + m.Values[2][2]*v.Z,
	}
}

func (m Matrix) Mul(other Matrix) Matrix {
	var result Matrix
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result.Values[i][j] = m.Values[i][0]*other.Values[0][j] +
				m.Values[i][1]*other.Values[1][j] +
				m.Values[i][2]*other.Values[2][j]
		}
	}
	return result
}


// Camera represents a simple 3D camera.
type Camera struct {
	Position    Vector3
	Direction   Vector3
	Up          Vector3
	FOV         float64 // Field of view in degrees
	Speed       float64 // Movement speed of the camera
	Pitch, Yaw  float64 // Pitch and yaw angles in degrees
	RotationMat Matrix  // Rotation matrix for camera orientation
}

// NewCamera creates a new camera with default parameters.
func NewCamera() *Camera {
	return &Camera{
		Position:    Vector3{0, 0, 0},
		Direction:   Vector3{0, 0, -1},
		Up:          Vector3{0, 1, 0},
		FOV:         60.0, // Default FOV of 60 degrees
		Speed:       0.1,  // Default speed for camera movement
		Pitch:       0.0,  // Initial pitch angle
		Yaw:         0.0,  // Initial yaw angle
		RotationMat: IdentityMatrix(),
	}
}

// GetRay generates a ray from the camera through the specified screen coordinates.
func (c *Camera) GetRay(x, y int) Vector3 {
	// Assume the screen is normalized to [-1, 1] in both x and y directions
	aspectRatio := float64(ScreenWidth) / float64(ScreenHeight)
	angle := math.Tan(math.Pi * c.FOV / 360.0)
	screenX := (2.0*float64(x)/float64(ScreenWidth) - 1.0) * angle * aspectRatio
	screenY := (1.0 - 2.0*float64(y)/float64(ScreenHeight)) * angle

	screenPos := Vector3{screenX, screenY, -1} // -1 is the distance to the screen plane

	// Rotate screenPos based on camera's pitch and yaw
	screenPos = c.RotationMat.MulVec(screenPos)

	rayDir := screenPos.Sub(c.Position).Normalize()

	return rayDir
}

// RotateCamera rotates the camera based on pitch and yaw angles.
// RotateCamera rotates the camera based on pitch and yaw angles.
func (c *Camera) RotateCamera(pitchDelta, yawDelta float64) {
	c.Pitch += pitchDelta
	c.Yaw += yawDelta

	// Limit pitch angle to avoid flipping upside down
	if c.Pitch > 89.0 {
		c.Pitch = 89.0
	}
	if c.Pitch < -89.0 {
		c.Pitch = -89.0
	}

	// Calculate new direction vector based on pitch and yaw
	radPitch := math.Pi * c.Pitch / 180.0
	radYaw := math.Pi * c.Yaw / 180.0
	dir := Vector3{
		X: math.Cos(radPitch) * math.Sin(radYaw),
		Y: math.Sin(radPitch),
		Z: -math.Cos(radPitch) * math.Cos(radYaw),
	}
	c.Direction = dir.Normalize()

	// Calculate new up vector based on pitch and yaw
	cross := Vector3{X: 0, Y: 1, Z: 0}.Cross(c.Direction).Normalize()
	c.Up = cross.Cross(c.Direction).Normalize()

	// Calculate rotation matrix for camera orientation
	c.RotationMat = RotateY(radYaw).Mul(RotateX(radPitch))
}


// Update updates the game logic.
func (c *Camera) Update() {
	// Handle rotation based on arrow key presses
	rotationSpeed := 1.0
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		c.RotateCamera(0, -rotationSpeed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		c.RotateCamera(0, rotationSpeed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		c.RotateCamera(-rotationSpeed, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyU) {
		c.RotateCamera(rotationSpeed, 0)
	}

	// Handle movement based on WASD key presses
	movementSpeed := 0.3
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		c.Position = c.Position.Add(c.Direction.Scale(movementSpeed))
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		c.Position = c.Position.Sub(c.Direction.Scale(movementSpeed))
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		c.Position = c.Position.Sub(c.Up.Cross(c.Direction).Normalize().Scale(movementSpeed))
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		c.Position = c.Position.Add(c.Up.Cross(c.Direction).Normalize().Scale(movementSpeed))
	}
}

// Ray represents a ray with an origin and direction.
type Ray struct {
	Origin, Direction Vector3
}

// PointAtParameter returns a point along the ray at a given parameter t.
func (r Ray) PointAtParameter(t float64) Vector3 {
	return r.Origin.Add(r.Direction.Scale(t))
}

// Sphere represents a sphere in 3D space.
type Sphere struct {
	Center       Vector3
	Radius       float64
	Color        color.RGBA
	Reflectivity float64 // Reflectivity factor (0.0 - 1.0)
	Emission     float64 // Emission factor (0.0 - 1.0)
}

// HitResult represents the result of a ray-sphere intersection.
type HitResult struct {
	Hit          bool
	Point        Vector3
	Normal       Vector3
	Distance     float64
	Sphere       Sphere // Sphere that was hit
	InsideSphere bool   // Whether the hit point is inside the sphere
}
// Hit checks if the ray intersects with the sphere and returns the hit result.
func (s Sphere) Hit(ray Ray, tMin, tMax float64) HitResult {
	oc := ray.Origin.Sub(s.Center)
	a := ray.Direction.Dot(ray.Direction)
	b := oc.Dot(ray.Direction)
	c := oc.Dot(oc) - s.Radius*s.Radius
	discriminant := b*b - a*c

	hitResult := HitResult{Hit: false}

	if discriminant > 0 {
		temp := (-b - math.Sqrt(discriminant)) / a
		if temp < tMax && temp > tMin {
			hitResult.Hit = true
			hitResult.Distance = temp
			hitResult.Point = ray.PointAtParameter(hitResult.Distance)
			hitResult.Normal = hitResult.Point.Sub(s.Center).Normalize()
			hitResult.Sphere = s
			hitResult.InsideSphere = hitResult.Normal.Dot(ray.Direction) > 0
			return hitResult
		}
		temp = (-b + math.Sqrt(discriminant)) / a
		if temp < tMax && temp > tMin {
			hitResult.Hit = true
			hitResult.Distance = temp
			hitResult.Point = ray.PointAtParameter(hitResult.Distance)
			hitResult.Normal = hitResult.Point.Sub(s.Center).Normalize()
			hitResult.Sphere = s
			hitResult.InsideSphere = hitResult.Normal.Dot(ray.Direction) > 0
			return hitResult
		}
	}

	return hitResult
}


// Game represents the game state.
type Game struct {
	camera     *Camera
	spheres    []Sphere
	maxBounces int // Maximum number of reflections
}

// Update updates the game logic.
func (g *Game) Update() error {
	g.camera.Update()
	return nil
}

// Draw draws the game.
func (g *Game) Draw(screen *ebiten.Image) {
	for y := 0; y < ScreenHeight; y++ {
		for x := 0; x < ScreenWidth; x++ {
			color := g.traceRay(g.camera.Position, g.camera.GetRay(x, y), 0)
			screen.Set(x, y, color)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

// traceRay traces a ray and calculates the resulting color recursively up to maxBounces reflections.
func (g *Game) traceRay(origin Vector3, rayDirection Vector3, depth int) color.RGBA {
	if depth >= g.maxBounces {
		return color.RGBA{0, 0, 0, 255} // Return black if maximum depth reached
	}

	closestHit := HitResult{Hit: false, Distance: math.Inf(1)} // Initialize with no hit

	// Check for intersection with spheres
	for _, sphere := range g.spheres {
		hit := sphere.Hit(Ray{Origin: origin, Direction: rayDirection}, 0.001, closestHit.Distance)
		if hit.Hit {
			closestHit = hit
		}
	}

	if closestHit.Hit {
		// Calculate color based on hit properties
		if closestHit.Sphere.Reflectivity > 0 {
			// Calculate reflection ray
			reflectedDirection := rayDirection.Sub(closestHit.Normal.Scale(2 * rayDirection.Dot(closestHit.Normal)))
			reflectedColor := g.traceRay(closestHit.Point, reflectedDirection, depth+1)

			// Mix reflection color with sphere color based on reflectivity
			color := color.RGBA{
				R: uint8(float64(closestHit.Sphere.Color.R) * (1.0 - closestHit.Sphere.Reflectivity)),
				G: uint8(float64(closestHit.Sphere.Color.G) * (1.0 - closestHit.Sphere.Reflectivity)),
				B: uint8(float64(closestHit.Sphere.Color.B) * (1.0 - closestHit.Sphere.Reflectivity)),
				A: 255,
			}

			color.R += uint8(float64(reflectedColor.R) * closestHit.Sphere.Reflectivity)
			color.G += uint8(float64(reflectedColor.G) * closestHit.Sphere.Reflectivity)
			color.B += uint8(float64(reflectedColor.B) * closestHit.Sphere.Reflectivity)

			return color
		} else {
			// No reflection, calculate surface color directly
			normal := closestHit.Normal
			color := color.RGBA{
				R: uint8(float64(closestHit.Sphere.Color.R) * (0.5 * (normal.X + 1))),
				G: uint8(float64(closestHit.Sphere.Color.G) * (0.5 * (normal.Y + 1))),
				B: uint8(float64(closestHit.Sphere.Color.B) * (0.5 * (normal.Z + 1))),
				A: 255,
			}
			return color
		}
	} else {
		// No intersection, return background color
		return color.RGBA{0, 0, 0, 255}
	}
}

func main() {
	// Set up the window
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("3D Raytracer with Ebiten")

	// Create camera
	camera := NewCamera()

	// Create spheres
	spheres := []Sphere{
		{Center: Vector3{-2, 0, -5}, Radius: 1.0, Color: color.RGBA{255, 0, 0, 255}, Reflectivity: 0.3, Emission: 0.9},    // Red sphere with reflection and emission
		{Center: Vector3{0, 0, -5}, Radius: 1.0, Color: color.RGBA{0, 255, 0, 255}, Reflectivity: 0.3},                     // Green sphere without reflection
		{Center: Vector3{2, 0, -5}, Radius: 1.0, Color: color.RGBA{0, 0, 255, 255}, Reflectivity: 0.4, Emission: 0.5},    // Blue sphere with high reflection and emission
		{Center: Vector3{-1, 1, -3}, Radius: 0.5, Color: color.RGBA{255, 255, 0, 255}, Reflectivity: 0.5},                  // Yellow sphere without emission
		{Center: Vector3{1, 1, -3}, Radius: 0.5, Color: color.RGBA{255, 0, 255, 255}, Reflectivity: 0.6, Emission: 0.3},   // Magenta sphere with reflection and emission
		{Center: Vector3{0, -10001, -5}, Radius: 10000.0, Color: color.RGBA{255, 255, 255, 255}, Reflectivity: 0.2, Emission: 0.9}, // Ground plane with emission
	}
	

	// Create game
	game := &Game{
		camera:     camera,
		spheres:    spheres,
		maxBounces: 3, // Limit reflections to 3 bounces
	}

	// Run the game
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

