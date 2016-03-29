/*
	This file is for N regular polygon structure
*/
package models

import (
	"github.com/Geoferry/Godar/utils"
	//"fmt"
	"fmt"
	"image"
	"image/color"
	"math"
	"strconv"
)

/*
	Create a new N regular polygon radar chart

	In order to get a radar chart,

	the number of all the vertex except center vertex is layers * N

	the number of all the edges is (layers + 1) * N
*/
func (p *Ngon) New(nedge int, img *image.RGBA) {
	l, ok := utils.Config.GetSetting("layers")
	if !ok {
		return
	}
	layers, err := strconv.Atoi(l)
	if err != nil {
		fmt.Println(err)
		return
	}

	p.nedges = nedge
	p.centerX = img.Bounds().Max.X / 2
	p.centerY = img.Bounds().Max.Y / 2
	/*
		Get all the layers * N vertex we need
	*/
	p.vertex = make(map[int]map[int]int)
	radius, radians := utils.GetRadiusRadians(img, nedge)

	for j := 0; j < layers; j++ {
		for i := 0; i < nedge; i++ {
			tmpRadians := radians * float64(i)
			tmpX := utils.ToInt(math.Sin(tmpRadians)*float64(radius*(layers-j)/layers)) + p.centerX
			tmpY := utils.ToInt(math.Cos(tmpRadians)*float64(radius*(layers-j)/layers)) + p.centerY
			p.vertex[j*nedge+i] = make(map[int]int)
			p.vertex[j*nedge+i][tmpX] = tmpY
		}
	}

	/*
		Add all the edges by two step

		Step1: add the edges between center vertex to outer vertex

		Step2: add the edges between adjacent vertex layer by layer

		Note: At "Step 2", the tmpVerx and tmpVerY could only contain 1000 elements,
		if there happens any panic, you can change it below.
	*/

	countEdge := 0
	p.edges = make(map[int]*edge)
	//Step 1:
	for i := 0; i < nedge; i++ {
		for tmpX, tmpY := range p.vertex[i] {
			tmpEdge := &edge{}
			tmpEdge.x0 = p.centerX
			tmpEdge.y0 = p.centerY
			tmpEdge.x1 = tmpX
			tmpEdge.y1 = tmpY
			p.edges[countEdge] = tmpEdge
			countEdge++
		}
	}
	//Step 2:
	tmpVerX := make([]int, 1000)
	tmpVerY := make([]int, 1000)
	for i := 0; i < len(p.vertex); i++ {
		for tmpX, tmpY := range p.vertex[i] {
			tmpVerX[i] = tmpX
			tmpVerY[i] = tmpY
		}
	}
	tmpVerX = tmpVerX[:layers*nedge]
	tmpVerY = tmpVerY[:layers*nedge]
	//fmt.Println(len(tmpVerX), cap(tmpVerX))
	for i := 0; i < len(tmpVerX); i++ {
		tmpEdge := &edge{}
		tmpEdge.x0 = tmpVerX[i]
		tmpEdge.y0 = tmpVerY[i]
		if (i+1)%nedge == 0 {
			tmpEdge.x1 = tmpVerX[i-nedge+1]
			tmpEdge.y1 = tmpVerY[i-nedge+1]
		} else {
			tmpEdge.x1 = tmpVerX[i+1]
			tmpEdge.y1 = tmpVerY[i+1]
		}
		p.edges[countEdge] = tmpEdge
		countEdge++
	}
}

func (p *Ngon) DrawNgonLine(thick int, rgba *image.RGBA, c color.RGBA) {
	for _, key := range p.edges {
		utils.DrawLine(key.x0, key.y0, key.x1, key.y1, thick, c, rgba)
	}
}

func (p *Ngon) FillLayer(layer int, rgba *image.RGBA, c color.RGBA) {
	if p.nedges < 3 {
		return
	}

	l, ok := utils.Config.GetSetting("layers")
	if !ok {
		return
	}
	layers, err := strconv.Atoi(l)
	if err != nil {
		fmt.Println(err)
		return
	}
	if layer > layers {
		return
	}

	for i := 0; i < p.nedges/2; i++ {
		e1, e2 := p.edges[layer*p.nedges+i], p.edges[layer*p.nedges+p.nedges-i-1]

		maxY := getMaximum(e1.y0, e1.y1, e2.y0, e2.y1)
		minY := getMinimum(e1.y0, e1.y1, e2.y0, e2.y1)

		for y := minY; y <= maxY; y++ {
			//fmt.Println("y: ", y)
			ok, x0, x1 := getX(y, e1, e2)
			if !ok {
				continue
			}

			dx := 1
			if x0 > x1 {
				dx = -1
			}
			//fmt.Println("x0: ", x0, "\tx1: ", x1)
			for x := x0; x != x1; x += dx {
				rgba.Set(x, y, c)
			}
		}
	}
}
