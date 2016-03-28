package utils

import (
	//"fmt"
	"errors"
	"image"
	"math/rand"
	"os"
	"strconv"
	"time"
)

/*
	Create a new image

	The first parameter(A int array) indicates the maximum value of X-axis & Y-axis

	This function could create PNG & JPEG file,
	you can determine it with a specific imageType string,
	you can ignore it to use default type -- PNG.

	If you enter a wrong imageType string or not supported imageType,
	it would be change to default imageType.

	This function returns a string -- new picture's path.
*/
func NewImagePath() (string, error) {
	prePath, ok := Config.GetSetting("pathPrefix")
	if !ok {
		return "", errors.New("pathPrefix not set yet")
	}
	//Create output folder if not exist
	if ok, err := pathExists(prePath); ok {
		//fmt.Println("Folder already exists!")
	} else if err == nil {
		//fmt.Println("Folder not exist! Create a new folder!")
		os.Mkdir(prePath, 0777)
	} else {
		return "", err
	}
	//Set path
	path := getPath(prePath)

	return path, nil
}

/*
	Set maximum X & Y

	If not require size, it will set default size dx = 2560, dy = 2560

	If there's only one parameter, it will set X and let Y use default value

	If there're more than one parameter, only use the previous two parameter

	This function return an image.Rectangle
*/
func SetSize(size ...int) (img image.Rectangle) {
	tx, ok := Config.GetSetting("defaultX")
	ty, ok1 := Config.GetSetting("defaultY")
	if !ok1 || !ok {
		return
	}

	dx, err := strconv.Atoi(tx)
	dy, err1 := strconv.Atoi(ty)
	if err != nil || err1 != nil {

	}

	if len(size) == 0 {
		img.Max.X = dx
		img.Max.Y = dy
	}

	if len(size) == 1 {
		img.Max.X = size[0]
		img.Max.Y = dy
	}

	if len(size) >= 2 {
		img.Max.X = size[0]
		img.Max.Y = size[1]
	}
	return
}

/******************Private Function*******************/

/*
	Generate prefix string with time
*/
func genTime() string {
	return time.Now().Format("20060102_150405")
}

func getPath(prePath string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	path := prePath + genTime() + "_" + strconv.Itoa(r.Intn(10)*r.Intn(20)+r.Intn(30)) + ".png"
	return path
}

/*
	Check if the path is exist

	If bool returns true, then the folder or file is exist.

	If bool return false && error returns nil, then the folder or file is not exist.

	If error not returns nil, then we're not sure the folder or file is exist.
*/
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
