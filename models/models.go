package models

//Point
type point struct {
	x, y int
}

//Edge from (x0, y0) to (x1, y1)
type edge struct {
	x0, x1, y0, y1 int
}

/*
	N regular polygon

	Vertex is a map for every useful vertex we need
*/
type Ngon struct {
	nedges  int
	centerX int
	centerY int
	edges   map[int]*edge
	vertex  map[int]map[int]int //number - x_axis - y_axis
}

type Circle struct {
	n          int
	centerX    int
	centerY    int
	radius     int
	layerPoint map[*point]bool //points on each layer
}

/**************		Private Function	*******************/

//This Function created for Ngon's FillLayer method
func getX(y int, e1 *edge, e2 *edge) (ok bool, x0, x1 int) {
	if (e1.x0 == e1.x1) || (e2.x0 == e2.x1) {
		return true, e1.x0, e2.x0
	}
	//Line kx + c = y
	dy, dx := e1.y0-e1.y1, e1.x0-e1.x1
	c := e1.y0 - dy*e1.x0/dx
	x0 = (y - c) * dx / dy

	dy, dx = e2.y0-e2.y1, e2.x0-e2.x1
	c = e2.y0 - dy*e2.x0/dx
	x1 = (y - c) * dx / dy
	return true, x0, x1
}

func getMaximum(tmp ...int) int {
	if len(tmp) == 0 {
		return 0
	}

	if len(tmp) == 1 {
		return tmp[0]
	}

	max := tmp[0]
	for i := range tmp {
		if max < tmp[i] {
			max = tmp[i]
		}
	}
	return max
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
