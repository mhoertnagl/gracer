package main

import (
	"log"
	"math"
	"time"

	"github.com/mhoertnagl/gracer/alg"
	"github.com/mhoertnagl/gracer/canvas"
	"github.com/mhoertnagl/gracer/render"
)

func main() {
	defer timeTrack(time.Now(), "Ray Trace")

	// wallMaterial := render.NewMaterial()
	// wallMaterial.Color = canvas.NewColor(1, 0.9, 0.9)
	// wallMaterial.Specular = 0

	pm1 := render.NewSolidPattern(canvas.White)
	pm2 := render.NewSolidPattern(canvas.NewColor(0.5, 0.5, 0.5))
	pm3 := render.NewStripePattern(pm1, pm2)
	pm4 := render.NewStripePattern(pm1, pm2)
	pm4.Transform = alg.RotationY(math.Pi / 2)
	pm5 := render.NewBlendedPattern(pm3, pm4)

	// pattern.Transform = alg.Scaling(2, 1, 1)

	floor := render.NewPlane()
	// floor.Transform = alg.RotationY(math.Pi / 5)
	floor.Material.Pattern = pm5

	// floor := render.NewSphere()
	// floor.Transform = alg.Scaling(10, 0.01, 10)
	// floor.Material = wallMaterial

	// leftWall := render.NewSphere()
	// leftWall.Transform = alg.
	// 	Translation(0, 0, 5).
	// 	MultMat(alg.RotationY(-math.Pi / 4)).
	// 	MultMat(alg.RotationX(math.Pi / 2)).
	// 	MultMat(alg.Scaling(10, 0.01, 10))
	// leftWall.Material = wallMaterial

	// rightWall := render.NewSphere()
	// rightWall.Transform = alg.
	// 	Translation(0, 0, 5).
	// 	MultMat(alg.RotationY(math.Pi / 4)).
	// 	MultMat(alg.RotationX(math.Pi / 2)).
	// 	MultMat(alg.Scaling(10, 0.01, 10))
	// rightWall.Material = wallMaterial

	p1 := render.NewSolidPattern(canvas.White)
	p2 := render.NewSolidPattern(canvas.NewColor(0.5, 0.5, 0.5))
	p3 := render.NewRingPattern(p1, p2)
	p3.Transform = alg.Scaling(0.1, 0.1, 0.1)

	middle := render.NewSphere()
	middle.Transform = alg.Translation(-0.5, 1, 0.5)
	middle.Material.Pattern = p3
	middle.Material.Color = canvas.NewColor(0.1, 1, 0.5)
	middle.Material.Diffuse = 0.7
	middle.Material.Specular = 0.3

	right := render.NewSphere()
	right.Transform = alg.
		Translation(1.5, 0.5, -0.5).
		MultMat(alg.Scaling(0.5, 0.5, 0.5))
	right.Material.Color = canvas.NewColor(0.5, 1, 0.1)
	right.Material.Diffuse = 0.7
	right.Material.Specular = 0.3
	right.Material.Shininess = 10

	left := render.NewSphere()
	left.Transform = alg.
		Translation(-1.5, 0.33, -0.75).
		MultMat(alg.Scaling(0.33, 0.33, 0.33))
	left.Material.Color = canvas.NewColor(1, 0.8, 0.1)
	left.Material.Diffuse = 0.7
	left.Material.Specular = 0.3

	light := render.NewPointLight(alg.NewPoint(-10, 10, -10), canvas.White)
	// light2 := render.NewPointLight(alg.NewPoint(10, 10, 10), canvas.White)

	world := render.NewWorld()
	world.AddLight(light)
	// world.AddLight(light2)
	world.AddObject(floor)
	// world.AddObject(leftWall)
	// world.AddObject(rightWall)
	world.AddObject(middle)
	world.AddObject(right)
	world.AddObject(left)

	camera := render.NewCamera(300, 150, math.Pi/3)
	camera.Transform = render.ViewTransform(
		// alg.NewPoint(0, 1.5, -5),
		alg.NewPoint(0, 3, -5),
		alg.NewPoint(0, 1, 0),
		alg.NewVector3(0, 1, 0),
	)

	canvas := world.Render(camera)
	canvas.WriteToFile("out.jpg")
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
