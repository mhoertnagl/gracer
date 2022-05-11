package main

import (
	"log"
	"math"
	"time"

	"github.com/mhoertnagl/gracer/alg"
	"github.com/mhoertnagl/gracer/canvas"
	"github.com/mhoertnagl/gracer/obj"
	"github.com/mhoertnagl/gracer/render"
)

func main() {
	defer timeTrack(time.Now(), "Ray Trace")

	pm1 := render.NewSolidPattern(canvas.NewColor(0.6, 0.6, 0.6))
	pm2 := render.NewSolidPattern(canvas.NewColor(0.8, 0.8, 0.8))
	pm5 := render.NewCheckers3DPattern(pm1, pm2)
	// pm5.Transform = alg.Scaling(0.5, 0.5, 0.5)

	floor := render.NewPlane()
	floor.Material.Pattern = pm5
	// floor.Material.Reflective = 0.1
	// floor.Material.Specular = 0.1
	// floor.Material.Shininess = 10

	wall := render.NewPlane()
	wall.Material.Pattern = pm5
	wall.Transform = alg.
		Translation(0, 0, -5).
		MultMat(alg.RotationY(math.Pi / 2)).
		MultMat(alg.RotationX(math.Pi / 2))

	parser := obj.NewParser()
	parser.ParseFile("./models/teapot_n_glass.obj")
	teapot := parser.Root
	for _, v := range teapot.Kids {
		v.GetMaterial().Ambient = 0.2
		v.GetMaterial().Reflective = 0.1
	}
	teapot.Transform = alg.
		Translation(10, 0, 0)
		//.
		//MultMat(alg.RotationY(-math.Pi / 2)) //.MultMat(alg.Scaling(0.1, 0.1, 0.1))
	// wall.Material.Reflective = 0.1
	// wall.Material.Specular = 0.1
	// wall.Material.Shininess = 10
	// floor.Material.ReceiveShadow = false

	// leftWall := render.NewPlane()
	// leftWall.Transform = alg.
	// 	Translation(0, 0, -5).
	// 	MultMat(alg.RotationY(-math.Pi / 4)).
	// 	MultMat(alg.RotationX(math.Pi / 2))
	// // leftWall.Material.Pattern = pm5
	// // leftWall.Material.Reflective = 0.1
	// leftWall.Material.Color = canvas.White
	// leftWall.Material.Specular = 0.1
	// leftWall.Material.Shininess = 10
	// // leftWall.Material.ReceiveShadow = false

	// rightWall := render.NewPlane()
	// rightWall.Transform = alg.
	// 	Translation(0, 0, -5).
	// 	MultMat(alg.RotationY(math.Pi / 4)).
	// 	MultMat(alg.RotationX(math.Pi / 2))
	// rightWall.Material.Color = canvas.White
	// // rightWall.Material.Pattern = pm5
	// // rightWall.Material.Reflective = 0.1
	// rightWall.Material.Specular = 0.1
	// rightWall.Material.Shininess = 10
	// // rightWall.Material.ReceiveShadow = false

	// // p1 := render.NewSolidPattern(canvas.White)
	// // p2 := render.NewSolidPattern(canvas.NewColor(0.5, 0.5, 0.5))
	// // p3 := render.NewRingPattern(p1, p2)
	// // p3.Transform = alg.Scaling(0.1, 0.1, 0.1)

	// // middle := render.NewSphere()
	// middle := render.NewCube()
	// middle.Transform = alg.
	// 	Translation(0, 1, 0).
	// 	MultMat(alg.Scaling(0.5, 1, 0.5)).
	// 	MultMat(alg.RotationY(math.Pi / 4))
	// // middle.Material.Pattern = p3
	// middle.Material.Color = canvas.NewColor(0.1, 0.1, 0.1)
	// middle.Material.Ambient = 0.2
	// middle.Material.Diffuse = 0.1
	// middle.Material.Specular = 1.0
	// middle.Material.Shininess = 300
	// middle.Material.Reflective = 0.9
	// middle.Material.Transparency = 0.9
	// middle.Material.RefractiveIndex = 1.52
	// // middle.Material.ReceiveShadow = false

	// bubble := render.NewSphere()
	// bubble.Transform = alg.Translation(0, 1, 0).MultMat(alg.Scaling(0.3, 0.3, 0.3))
	// // middle.Material.Pattern = p3
	// bubble.Material.Color = canvas.NewColor(0.1, 0.1, 0.1)
	// bubble.Material.Ambient = 0.0
	// bubble.Material.Diffuse = 0.0
	// bubble.Material.Specular = 0.0 //1.0
	// // bubble.Material.Shininess = 300
	// bubble.Material.Reflective = 1
	// bubble.Material.Transparency = 1
	// bubble.Material.RefractiveIndex = 1.00029
	// bubble.Material.ReceiveShadow = false

	// right := render.NewCube()
	// right.Transform = alg.
	// 	Translation(1.5, 0.5, 0.5).
	// 	MultMat(alg.Scaling(0.5, 0.5, 0.5))

	// right.Material.Color = canvas.White
	// // right.Material.Ambient = 0.3
	// right.Material.Diffuse = 1
	// // right.Material.Specular = 0.3
	// // right.Material.Shininess = 100
	// right.Material.Reflective = 0.1
	// // right.Material.Transparency = 0.0
	// right.Material.RefractiveIndex = 1 //.52

	// left := render.NewCube()
	// left.Transform = alg.
	// 	Translation(-1.5, 0.33, -0.75).
	// 	MultMat(alg.Scaling(0.33, 0.33, 0.33))
	// left.Material.Color = canvas.NewColor(0, 1, 0)
	// left.Material.Diffuse = 0.7
	// left.Material.Specular = 0.3
	// left.Material.Reflective = 0.1

	// cyl := render.NewCylinder()
	// cyl.Transform = alg.
	// 	Translation(1.5, 0, -1).
	// 	MultMat(alg.Scaling(0.33, 2, 0.33))
	// cyl.Material.Color = canvas.NewColor(1, 0, 0)
	// cyl.Material.Diffuse = 0.7
	// cyl.Material.Specular = 0.3
	// cyl.Material.Reflective = 0.1

	// group := render.NewGroup()
	// group.AddKid(middle)
	// group.AddKid(bubble)
	// group.AddKid(right)
	// group.AddKid(left)
	// group.AddKid(cyl)
	// group.Transform = alg.Id4.RotateY(math.Pi/4).Translate(0.3, 0, 1.5)

	// light := render.NewPointLight(alg.NewPoint(-50, 150, 150), canvas.White)
	light := render.NewPointLight(alg.NewPoint(10, 20, -10), canvas.White)

	world := render.NewWorld()
	world.MaxBounces = 5
	world.AddLight(light)
	// world.AddLight(light2)
	world.AddObject(floor)
	world.AddObject(wall)
	world.AddObject(teapot)
	// world.AddObject(leftWall)
	// world.AddObject(rightWall)
	// world.AddObject(group)
	// world.AddObject(middle)
	// world.AddObject(bubble)
	// world.AddObject(right)
	// world.AddObject(left)
	// world.AddObject(cyl)

	// camera := render.NewCamera(300, 300, math.Pi/3)
	// camera := render.NewCamera(600, 600, math.Pi/3)
	camera := render.NewCamera(1200, 1200, math.Pi/3)
	camera.Transform = render.ViewTransform(
		alg.NewPoint(20, 4, 0),
		alg.NewPoint(0, 0, 0),
		alg.NewVector3(0, 1, 0),
		// alg.NewPoint(2.5, 2.5, 3),
		// alg.NewPoint(0, 0, 0),
		// alg.NewVector3(0, 1, 0),
		// alg.NewPoint(0, 9, 3),
		// alg.NewPoint(0, 7, 0),
		// alg.NewVector3(0, 1, 0),
		// alg.NewPoint(0, 11, 0),
		// alg.NewPoint(0, 0, 0),
		// alg.NewVector3(0, 0, 1),
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
