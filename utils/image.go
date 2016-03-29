package utils

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"strconv"
)

/*********	Start 	DrawCircle ********/

func DrawCircle(centerX, centerY, radius, thick int, c color.RGBA, rgba *image.RGBA) {
	minX, maxX := centerX-radius-thick, centerX+radius+thick
	minY, maxY := centerY-radius-thick, centerY+radius+thick
	for x := minX; x < maxX; x++ {
		for y := minY; y < maxY; y++ {
			if possibleCirclePoint(centerX, centerY, x, y, radius, thick) {
				rgba.Set(x, y, c)
			}
		}
	}
}

func possibleCirclePoint(cx, cy, x, y, r, thick int) bool {
	dx, dy := cx-x, cy-y
	if (dx*dx+dy*dy > (r+thick)*(r+thick)) || (dx*dx+dy*dy < (r-thick)*(r-thick)) {
		return false
	}
	return true
}

/*********  Start   DrawLine   ********/
func DrawLine(x0, y0, x1, y1, thick int, c color.RGBA, rgba *image.RGBA) {
	dx := math.Abs(float64(x1 - x0))
	dy := math.Abs(float64(y1 - y0))
	sx, sy := 1, 1
	if x0 >= x1 {
		sx = -1
	}
	if y0 >= y1 {
		sy = -1
	}
	err := dx - dy
	for {
		rgba.Set(x0, y0, c)
		for i := 1; i <= thick; i++ {
			rgba.Set(x0+i, y0, c)
			rgba.Set(x0, y0+i, c)
		}
		if x0 == x1 && y0 == y1 {
			return
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

/*
	*********  Start   DrawDataLineByData   ********
	The first parameter is a map
	In this map, string is the subject's title on each vertex
*/
func DrawDataLineByData(data map[string]int, thick int, rgba *image.RGBA, lineColor, fontColor color.RGBA) {
	/*
		1st layer -- 1000
		2nd layer -- 500
		3rd layer -- 100
	*/
	centerX, centerY := rgba.Bounds().Max.X/2, rgba.Bounds().Max.Y/2
	radius, radians := GetRadiusRadians(rgba, len(data))

	cc, n := 0, len(data)
	x0, x1, fx, y0, y1, fy := 0, 0, 0, 0, 0, 0

	for key, v := range data {
		tmpRadians := radians * float64(cc)
		x0, y0 = x1, y1
		var tmp float64
		if v <= 100 {
			tmp = float64(v) / float64(100) / 3
		} else if v > 100 && v <= 500 {
			tmp = float64(1.0/3.0) + float64(v-100)/float64(400)/3
		} else if v > 500 && v <= 1000 {
			tmp = float64(2.0/3.0) + float64(v-500)/float64(500)/3
		} else {
			tmp = 1.0
		}

		x1 = ToInt(math.Sin(tmpRadians)*tmp*float64(radius)) + centerX
		y1 = ToInt(math.Cos(tmpRadians)*tmp*float64(radius)) + centerY

		if cc == 0 {
			fx, fy = x1, y1
		} else if cc < n-1 {
			DrawLine(x0, y0, x1, y1, thick, lineColor, rgba)
		} else if cc == n-1 {
			DrawLine(x0, y0, x1, y1, thick, lineColor, rgba)
			DrawLine(fx, fy, x1, y1, thick, lineColor, rgba)
		}
		//DrawString
		tx, ty := calFontPosition(x1, y1, centerX, centerY)
		DrawString(tx, ty, key, rgba, fontColor)
		cc++
	}
}

/*
	*********  Start   DrawDataLineByPercentage   ********
	The first parameter is a map
	In this map, string is the subject's title on each vertex
*/
func DrawDataLineByPercentage(data map[string]int, thick int, rgba *image.RGBA, lineColor, fontColor color.RGBA) {
	var sum float64
	for _, v := range data {
		sum += float64(v)
	}
	/*
		1st layer -- 100%
		2nd layer -- 40%
		3rd layer -- 5%
	*/
	centerX, centerY := rgba.Bounds().Max.X/2, rgba.Bounds().Max.Y/2
	radius, radians := GetRadiusRadians(rgba, len(data))

	cc, n := 0, len(data)
	x0, x1, fx, y0, y1, fy := 0, 0, 0, 0, 0, 0

	for key, v := range data {
		fmt.Println(v, float64(v)/sum, getVertexPerByPer(v, sum))

		tmpRadians := radians * float64(cc)
		x0, y0 = x1, y1
		per := getVertexPerByPer(v, sum)
		// per := float64(v) / sum

		// if per <= 0.05 {
		// 	per = per / 0.05 / 3
		// } else if per > 0.05 && per <= 0.4 {
		// 	per = float64(1.0/3.0) + float64((per-0.05)/0.35/3)
		// } else if per > 0.4 && per <= 1 {
		// 	per = float64(2.0/3.0) + float64((per-0.4)/0.6/3)
		// }

		x1 = ToInt(math.Sin(tmpRadians)*per*float64(radius)) + centerX
		y1 = ToInt(math.Cos(tmpRadians)*per*float64(radius)) + centerY

		if cc == 0 {
			fx, fy = x1, y1
		} else if cc < n-1 {
			DrawLine(x0, y0, x1, y1, thick, lineColor, rgba)
		} else if cc == n-1 {
			DrawLine(x0, y0, x1, y1, thick, lineColor, rgba)
			DrawLine(fx, fy, x1, y1, thick, lineColor, rgba)
		}
		//DrawString
		tx, ty := calFontPosition(x1, y1, centerX, centerY)
		DrawString(tx, ty, key, rgba, fontColor)
		cc++
	}
}

/*
	This function only used for generate point position.

	Get Radius according to the img's X & Y axis.

	Get Radians according to N and radius.
*/
func GetRadiusRadians(img *image.RGBA, n int) (radius int, radians float64) {
	tmpMin := getMinimum(img.Bounds().Max.X, img.Bounds().Max.Y)
	if tmpMin/2 <= 50 {
		radius = tmpMin / 2
	} else {
		radius = tmpMin / 2 * 4 / 5
	}
	radians = math.Pi * 2.0 / float64(n)
	return
}

/**************		Private Function	*******************/
// func getVertexPerByVal() {
// 	key, ok := Config.GetSetting("equal_division")
// 	layers, ok2 := Config.GetSetting("layers")
// 	if !ok {
// 		fmt.Println("equal_division not set in config.conf")
// 		//return 0.0
// 	}
// 	if !ok2 {
// 		fmt.Println("layers not set in config.conf")
// 		//return 0.0
// 	}

// 	if key == "0" {

// 	} else {

// 	}
// }

func getVertexPerByPer(v int, sum float64) float64 {
	equal, ok := Config.GetSetting("equal_division")
	fmt.Println("e", equal)
	if !ok {
		return 0.0
	}

	l, ok := Config.GetSetting("layers")
	if !ok {
		return 0.0
	}

	m, ok := Config.GetSetting("min_percent")
	if !ok {
		return 0.0
	}

	layers, err := strconv.Atoi(l)
	if err != nil {
		return 0.0
	}

	min_percent, err := strconv.ParseFloat(m, 64)
	if err != nil {
		return 0.0
	}

	mm := make(map[int]float64)
	mm[0] = 0.0
	for i := 1; i <= layers; i++ {
		if i == 1 {
			mm[i] = min_percent
			continue
		}
		if i == layers {
			mm[i] = 1
			continue
		}
		mm[i] = subFunc(i-1) / subFunc(layers)
	}
	// for i := 1; i <= layers; i++ {
	// 	fmt.Print(mm[i], "\t")
	// }
	// fmt.Println()

	if equal == "0" {
		fmt.Println(0)
		per := float64(v) / sum
		ans := 0.0
		for i := 1; i <= layers; i++ {
			if per >= mm[i-1] && per <= mm[i] {
				ans = float64(i-1)/float64(layers) + (per-mm[i-1])/(mm[i]-mm[i-1])/float64(layers)
				break
			}
		}
		return ans
	} else {
		fmt.Println(1)
		return float64(v) / sum
	}
}

func subFunc(layers int) float64 {
	return math.Sqrt(float64(layers)) + float64(layers)*1.5
}

func getMinimum(tmp ...int) int {
	if len(tmp) == 0 {
		return 0
	}

	if len(tmp) == 1 {
		return tmp[0]
	}

	min := tmp[0]
	for i := range tmp {
		if min > tmp[i] {
			min = tmp[i]
		}
	}
	return min
}

func calFontPosition(x0, y0, centerX, centerY int) (x, y int) {
	offset := 50
	x, y = x0, y0
	if x0 < centerX {
		x = x0 - offset
	} else if x0 > centerX {
		x = x0 + offset
	}

	if y0 < centerY {
		y = y0 - offset
	} else if y0 > centerY {
		y = y + offset

	}
	return
}
