//packages usage samples
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
	path := ""
	wg.Add(4)
	/*******	package Utils usage sample -- DrawLine		*******/
	go func() {
		path, _ = utils.NewImagePath()
		f1, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer f1.Close()

		r1 := image.NewRGBA(utils.SetSize())
		utils.DrawLine(20, 150, 120, 250, 0, color.RGBA{0, 0, 0, 255}, r1)
		png.Encode(f1, r1)
		fmt.Println("DrawLine Sample finished")
		wg.Done()
	}()

	/*******	package Utils usage sample -- DrawCircle 	*******/
	go func() {
		path, _ = utils.NewImagePath()
		f2, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer f2.Close()

		r2 := image.NewRGBA(utils.SetSize())
		utils.DrawCircle(128, 128, 40, 3, color.RGBA{0, 0, 0, 255}, r2)
		png.Encode(f2, r2)
		fmt.Println("DrawCircle Sample finished")
		wg.Done()
	}()

	/*******	package Models usage sample -- Ngon 	*******/
	go func() {
		path, _ = utils.NewImagePath()
		f3, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer f3.Close()

		r3 := image.NewRGBA(utils.SetSize(2560, 1920))
		p := models.Ngon{}
		p.New(6, r3)
		p.FillLayer(1, r3, color.RGBA{80, 180, 240, 255})
		p.FillLayer(2, r3, color.RGBA{40, 120, 160, 255})
		p.FillLayer(3, r3, color.RGBA{20, 60, 80, 255})
		p.DrawNgonLine(2, r3, color.RGBA{0, 0, 0, 255})
		png.Encode(f3, r3)
		fmt.Println("N regular polygon radar chart sample finished")
		wg.Done()
	}()

	/*******	package Models usage sample -- Circle 	*******/
	go func() {
		path, _ = utils.NewImagePath()
		f4, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer f4.Close()

		r4 := image.NewRGBA(utils.SetSize(2560, 1920))
		cir := models.Circle{}
		cir.New(5, r4)
		cir.FillLayer(1, r4, color.RGBA{160, 150, 190, 255})
		cir.FillLayer(2, r4, color.RGBA{140, 110, 150, 255})
		cir.FillLayer(3, r4, color.RGBA{120, 70, 230, 255})
		//cir.FillLayer(4, r4, color.RGBA{60, 35, 115, 255})
		cir.DrawCurve(3, r4, color.RGBA{0, 0, 0, 255})
		png.Encode(f4, r4)
		fmt.Println("Circular radar chart sample finished")
		wg.Done()
	}()
	wg.Wait()
}
