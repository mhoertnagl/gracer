package main

import (
	"fmt"

	"github.com/mhoertnagl/gracer/canvas"
)

func main() {
	v := canvas.NewCanvas(300, 250)
	v.Set(100, 100, canvas.NewColor(1, 1, 1))
	v.WriteToFile("out.jpg")
	fmt.Println("Done")
}
