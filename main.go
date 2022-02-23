package main

import (
	"fmt"

	"github.com/mhoertnagl/gracer/alg"
	"github.com/mhoertnagl/gracer/canvas"
	"github.com/mhoertnagl/gracer/render"
)

func main() {
	ray_origin := alg.NewPoint(0, 0, -5)
	wall_z := 10.0
	wall_size := 7.0
	canvas_pixels := 300
	light_pos := alg.NewPoint(-10, 10, -10)
	light_color := canvas.NewColor(1, 1, 1)
	light := render.NewPointLight(light_pos, light_color)
	shape := render.NewSphere()
	shape.Material.Color = canvas.NewColor(1, 0.2, 1)
	pixel_size := wall_size / float64(canvas_pixels)
	half := wall_size / 2
	v := canvas.NewCanvas(canvas_pixels, canvas_pixels)
	for y := 0; y < canvas_pixels; y++ {
		world_y := half - pixel_size*float64(y)
		for x := 0; x < canvas_pixels; x++ {
			world_x := -half + pixel_size*float64(x)
			position := alg.NewPoint(world_x, world_y, wall_z)
			r := render.NewRay(ray_origin, position.Sub(ray_origin).Norm())
			xs := shape.Intersect(r)
			if xs.Hit() != nil {
				hit := xs.Hit()
				pnt := r.Position(hit.Distance)
				eye := r.Direction.Neg()
				normal := hit.Object.NormalAt(pnt)
				color := light.Lighting(hit.Object.GetMaterial(), pnt, eye, normal)
				v.Set(x, y, color)
			}
		}
	}
	v.WriteToFile("out.jpg")
	fmt.Println("Done")
}
