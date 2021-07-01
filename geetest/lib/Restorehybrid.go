package lib

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
)
import "github.com/Mas0nShi/goConsole/console"

type rGBA struct {
	R, G, B, _ uint32
}
func (c rGBA) toInt() int {
	return int((c.R >> 8) << 16 + (c.G >> 8) << 8 + c.B >> 8)
}
func RefuseImage(byteImg []byte) (image.Image,error) {
	img, err := jpeg.Decode(bytes.NewBuffer(byteImg))
	if err != nil {
		console.Error(err)
		return nil, err
	}
	bounds := img.Bounds()
	var sliceImg *image.YCbCr
	newImg := image.NewRGBA(bounds)
	pank1 := []int{157,145,265,277,181,169,241,253,109,97,289,301,85,73,25,37,13,1,121,133,61,49,217,229,205,193}
	pank2 := []int{145,157,277,265,169,181,253,241,97,109,301,289,73,85,37,25,1,13,133,121,49,61,229,217,193,205}
	step := 0
	for _, xPos := range pank1 {
		sliceImg = img.(*image.YCbCr).SubImage(image.Rect(xPos, 80, xPos + 10, 160)).(*image.YCbCr)
		draw.Draw(newImg,newImg.Bounds().Add(image.Pt(step, 0)), sliceImg, sliceImg.Bounds().Min, draw.Src)
		step += 10
	}
	step = 0
	for _, xPos := range pank2 {
		sliceImg = img.(*image.YCbCr).SubImage(image.Rect(xPos, 0, xPos + 10, 80)).(*image.YCbCr)
		draw.Draw(newImg,newImg.Bounds().Add(image.Pt(step, 80)), sliceImg, sliceImg.Bounds().Min, draw.Src)
		step += 10
	}
	return newImg, nil
}
func CompareOcr(gap *image.Image, src *image.Image) (int,error) {
	bounds := (*src).Bounds()
	find := false
	xPos := 0
	for x := 0; x < bounds.Dx(); x++ {
		i := 0
		for y := 0; y < bounds.Dy(); y++ {
			var t rGBA
			t.R,t.G,t.B,_ = (*src).At(x, y).RGBA()
			srcPixel := t.toInt()
			t.R,t.G,t.B,_ = (*gap).At(x, y).RGBA()
			gapPixel := t.toInt()
			if srcPixel != gapPixel && srcPixel - gapPixel > 2850135 {i++}
			if i >= 3 {
				find = true
				xPos = x
				break
			}
		}
		if find {
			break
		}
		i = 0
	}
	return xPos ,nil
}
func TestDrawLine(img image.Image, x int) image.Image {
	newImg := image.NewRGBA(img.Bounds())
	draw.Draw(newImg,newImg.Bounds(), img, img.Bounds().Min, draw.Src)
	for i := 0; i < newImg.Bounds().Dy(); i++ {
		newImg.Set(x, i, color.RGBA{R: 220, G: 20, B: 60})
	}
	return newImg
}