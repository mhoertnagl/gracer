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

	pm1 := render.NewSolidPattern(canvas.NewColor(0.6, 0.6, 0.6))
	pm2 := render.NewSolidPattern(canvas.NewColor(0.8, 0.8, 0.8))
	pm5 := render.NewCheckers3DPattern(pm1, pm2)

	floor := render.NewPlane()
	floor.Material.Pattern = pm5

	wall := render.NewPlane()
	wall.Material.Pattern = pm5
	wall.Transform = alg.
		Translation(0, 0, -5).
		MultMat(alg.RotationY(math.Pi / 2)).
		MultMat(alg.RotationX(math.Pi / 2))

	// parser := obj.NewParser()
	// parser.ParseFile("./assets/models/teapot_n_glass.obj")
	// teapot := parser.Root
	// for _, v := range teapot.Kids {
	// 	v.GetMaterial().Ambient = 0.2
	// 	v.GetMaterial().Reflective = 0.1
	// }
	// teapot.Transform = alg.
	// 	Translation(10, 0, 0)

	cube := render.NewCube()
	cube.GetMaterial().Ambient = 0.2
	cube.GetMaterial().Reflective = 0.1
	cube.Transform = alg.Id4.
		MultMat(alg.Translation(10, 1, 0)) //.
		// MultMat(alg.RotationY(math.Pi / 4))
	sphere := render.NewSphere()
	sphere.GetMaterial().Ambient = 0.15
	sphere.GetMaterial().Reflective = 0.1
	// sphere.Transform = alg.Scaling(1.41, 1.41, 1.41)
	sphere.Transform = alg.Id4.
		MultMat(alg.Translation(10, 1, 0)).
		MultMat(alg.Scaling(1.41, 1.41, 1.41))
	csg1 := render.NewIntersect(cube, sphere)
	// csg1.Transform = alg.Id4.
	// 	MultMat(alg.Translation(10, 1, 0)).
	// 	MultMat(alg.RotationY(math.Pi / 4))
	cylx := render.NewCylinder()
	cylx.Transform = alg.Id4.
		MultMat(alg.Translation(10, 1, -2)).
		MultMat(alg.Scaling(0.5, 0.5, 4)).
		MultMat(alg.RotationX(math.Pi / 2)) //.
		// MultMat(alg.RotationZ(math.Pi / 4))
	cyly := render.NewCylinder()
	cyly.Transform = alg.Id4.
		MultMat(alg.Translation(10, -1, 0)).
		MultMat(alg.Scaling(0.5, 4, 0.5))
	cylz := render.NewCylinder()
	cylz.Transform = alg.Id4.
		MultMat(alg.Translation(12, 1, 0)).
		MultMat(alg.Scaling(4, 0.5, 0.5)).
		MultMat(alg.RotationZ(math.Pi / 2))
	csg2 := render.NewUnion(cylx, cyly)
	csg3 := render.NewUnion(csg2, cylz)
	// csg3.Transform = alg.Id4.
	// 	MultMat(alg.Translation(10, 1, 0)).
	// 	MultMat(alg.RotationY(math.Pi / 4))
	csg4 := render.NewDifference(csg1, csg3)
	// csg4.Transform = alg.Id4.
	// MultMat(alg.RotationY(math.Pi / 4))
	//MultMat(alg.RotationY(math.Pi / 4))
	// cylx.Transform = alg.Id4.
	// 	MultMat(alg.Translation(10, 1, 0)).
	// 	MultMat(alg.Scaling(0.5, 0.5, 2)).
	// 	MultMat(alg.RotationX(math.Pi / 2))
	light := render.NewPointLight(alg.NewPoint(10, 20, -10), canvas.White)

	world := render.NewWorld()
	world.MaxBounces = 5
	world.AddLight(light)
	world.AddObject(floor)
	world.AddObject(wall)
	// world.AddObject(csg1)
	world.AddObject(csg4)
	// world.AddObject(teapot)

	//camera := render.NewCamera(300, 300, math.Pi/3)
	//camera := render.NewCamera(600, 600, math.Pi/3)
	camera := render.NewCamera(1200, 1200, math.Pi/3)
	camera.Transform = render.ViewTransform(
		// alg.NewPoint(20, 4, 0),
		// alg.NewPoint(0, 0, 0),
		// alg.NewVector3(0, 1, 0),
		alg.NewPoint(16, 4, 5),
		alg.NewPoint(10, 1, 0),
		alg.NewVector3(0, 1, 0),
	)

	canvas := world.Render(camera)
	log.Printf("Number of Rays: %d", int(world.NumRays))
	log.Printf("Hit Rate: %f", world.NumHits/world.NumRays)
	canvas.WriteToFile("out.jpg")
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
