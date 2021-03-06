package main

import (
	"fmt"
	"github.com/Geoferry/Godar/models"
	"github.com/Geoferry/Godar/utils"
	"image"
	"image/color"
	"image/png"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		p := &models.Ngon{}
		path, _ := utils.NewImagePath()
		file, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		rgba := image.NewRGBA(utils.SetSize(1920, 1080))
		data := make(map[string]int)
		data["S1"] = 3500
		data["S2"] = 800
		data["S3"] = 480
		data["S4"] = 620
		data["S5"] = 300
		data["S6"] = 170
		data["S7"] = 670
		p.New(len(data), rgba)
		p.FillLayer(1, rgba, color.RGBA{80, 180, 240, 255})
		p.FillLayer(2, rgba, color.RGBA{40, 120, 160, 255})
		p.FillLayer(3, rgba, color.RGBA{20, 60, 80, 255})
		//p.FillLayer(4, rgba, color.RGBA{10, 30, 40, 255})
		p.DrawNgonLine(2, rgba, color.RGBA{0, 0, 0, 255})
		utils.DrawDataLineByPercentage(data, 2, rgba, color.RGBA{190, 60, 80, 255}, color.RGBA{176, 25, 120, 255})
		utils.DrawString(100, 200, "按百分比画", rgba, color.RGBA{176, 25, 120, 255})
		png.Encode(file, rgba)
		wg.Done()
	}()

	/*****************************************/
	go func() {
		p2 := &models.Ngon{}
		path2, _ := utils.NewImagePath()
		file2, err2 := os.Create(path2)
		if err2 != nil {
			fmt.Println(err2)
		}
		defer file2.Close()

		rgba2 := image.NewRGBA(utils.SetSize(1920, 1080))
		data2 := make(map[string]int)
		data2["S1"] = 3500
		data2["S2"] = 800
		data2["S3"] = 480
		data2["S4"] = 620
		data2["S5"] = 300
		data2["S6"] = 170
		p2.New(len(data2), rgba2)
		p2.FillLayer(1, rgba2, color.RGBA{80, 180, 240, 255})
		p2.FillLayer(2, rgba2, color.RGBA{40, 120, 160, 255})
		p2.FillLayer(3, rgba2, color.RGBA{20, 60, 80, 255})
		// p2.FillLayer(4, rgba2, color.RGBA{10, 30, 40, 255})
		p2.DrawNgonLine(2, rgba2, color.RGBA{0, 0, 0, 255})
		utils.DrawDataLineByData(data2, 2, rgba2, color.RGBA{190, 60, 80, 255}, color.RGBA{176, 25, 120, 255})
		utils.DrawString(100, 200, "按数值画", rgba2, color.RGBA{176, 25, 120, 255})
		png.Encode(file2, rgba2)
		wg.Done()
	}()
	/*****************************************/
	go func() {
		c1 := &models.Circle{}
		path3, _ := utils.NewImagePath()
		file3, err3 := os.Create(path3)
		if err3 != nil {
			fmt.Println(err3)
		}
		defer file3.Close()

		rgba3 := image.NewRGBA(utils.SetSize(1920, 1080))
		data3 := make(map[string]int)
		data3["S1"] = 3500
		data3["S2"] = 800
		data3["S3"] = 480
		data3["S4"] = 620
		data3["S5"] = 300
		data3["S6"] = 170
		c1.New(len(data3), rgba3)
		c1.FillLayer(1, rgba3, color.RGBA{80, 180, 240, 255})
		c1.FillLayer(2, rgba3, color.RGBA{40, 120, 160, 255})
		c1.FillLayer(3, rgba3, color.RGBA{20, 60, 80, 255})
		// p2.FillLayer(4, rgba2, color.RGBA{10, 30, 40, 255})
		c1.DrawCurve(2, rgba3, color.RGBA{0, 0, 0, 255})
		utils.DrawDataLineByData(data3, 2, rgba3, color.RGBA{190, 60, 80, 255}, color.RGBA{176, 25, 120, 255})
		utils.DrawString(100, 200, "按数值画", rgba3, color.RGBA{176, 25, 120, 255})
		png.Encode(file3, rgba3)
		wg.Done()
	}()
	wg.Wait()
}
