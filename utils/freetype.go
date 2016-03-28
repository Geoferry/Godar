package utils

import (
	"github.com/golang/freetype"
	"image"
	"image/color"
	"io/ioutil"
	"log"
	"strconv"
)

func DrawString(x, y int, content string, rgba *image.RGBA, fontColor color.RGBA) {
	fontFile, ok := Config.GetSetting("fontFile")
	fontDPI, ok1 := Config.GetSetting("fontDPI")
	fontSize, ok2 := Config.GetSetting("fontSize")
	if !ok || !ok1 || !ok2 {
		log.Println("Font configuration not set yet")
	}

	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		log.Println(err)
		return
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	c := freetype.NewContext()

	DPI, err := strconv.Atoi(fontDPI)
	Size, err1 := strconv.Atoi(fontSize)
	if err != nil {
		log.Println(err)
	}
	if err1 != nil {
		log.Println(err1)
	}

	c.SetDPI(float64(DPI))
	c.SetFont(font)
	c.SetFontSize(float64(Size))
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	uniform := image.NewUniform(fontColor)
	c.SetSrc(uniform)

	position := freetype.Pt(x, y)
	_, err = c.DrawString(content, position)
	if err != nil {
		log.Println(err)
		return
	}
}
