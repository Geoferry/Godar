package models

import (
	"fmt"
	"github.com/Geoferry/Godar/utils"
	"image"
	"image/color"
	"math"
	"strconv"
)

/*
	Create a new circular radar chart

	we need to get N points and N edges
*/
func (cir *Circle) New(n int, img *image.RGBA) {
	cir.n = n
	cir.centerX = img.Bounds().Max.X / 2
	cir.centerY = img.Bounds().Max.Y / 2

	cir.layerPoint = make(map[*point]bool)

	tmpMin := getMinimum(img.Bounds().Max.X, img.Bounds().Max.Y)
	if tmpMin/2 <= 50 {
		cir.radius = tmpMin / 2
	} else {
		cir.radius = tmpMin / 2 * 4 / 5
	}
	radians := math.Pi * 2.0 / float64(n)

	for i := 0; i < n; i++ {
		tmpRadians := radians * float64(i)
		tmpX := utils.ToInt(math.Sin(tmpRadians)*float64(cir.radius)) + cir.centerX
		tmpY := utils.ToInt(math.Cos(tmpRadians)*float64(cir.radius)) + cir.centerY
		tmpPoint := &point{}
		tmpPoint.x = tmpX
		tmpPoint.y = tmpY
		cir.layerPoint[tmpPoint] = true
	}
}

func (cir *Circle) DrawCurve(thick int, rgba *image.RGBA, c color.RGBA) {
	l, ok := utils.Config.GetSetting("layers")
	if !ok {
		return
	}
	layers, err := strconv.Atoi(l)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 1; i <= layers; i++ {
		utils.DrawCircle(cir.centerX, cir.centerY, cir.radius*i/layers, thick, c, rgba)
	}

	for p, _ := range cir.layerPoint {
		utils.DrawLine(cir.centerX, cir.centerY, p.x, p.y, thick, c, rgba)
	}
}

func (cir *Circle) FillLayer(layer int, rgba *image.RGBA, c color.RGBA) {
	l, ok := utils.Config.GetSetting("layers")
	if !ok {
		return
	}
	layers, err := strconv.Atoi(l)
	if err != nil {
		fmt.Println(err)
		return
	}

	//tt := (layers - layer + 1) / layers
	minX, maxX := cir.centerX-cir.radius*layers/(layers-layer+1), cir.centerX+cir.radius*layers/(layers-layer+1)
	minY, maxY := cir.centerY-cir.radius*layers/(layers-layer+1), cir.centerY+cir.radius*layers/(layers-layer+1)
	if layer == layers {
		for x := minX; x < maxX; x++ {
			for y := minY; y < maxY; y++ {
				dx, dy := x-cir.centerX, y-cir.centerY
				if dx*dx+dy*dy < cir.radius*cir.radius/layers/layers {
					rgba.Set(x, y, c)
				}
			}
		}
		return
	}
	for x := minX; x < maxX; x++ {
		for y := minY; y < maxY; y++ {
			dx, dy := x-cir.centerX, y-cir.centerY
			if (dx*dx+dy*dy < cir.radius*cir.radius*(layers-layer+1)*(layers-layer+1)/layers/layers) &&
				(dx*dx+dy*dy > cir.radius*cir.radius*(layers-layer)*(layers-layer)/layers/layers) {
				rgba.Set(x, y, c)
			}
		}
	}
}
