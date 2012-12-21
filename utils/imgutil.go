package ccgo

import (
	"fmt"
	"image"
	//	"math"
	//	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

const EXT_JPEG = ".jpg"
const EXT_PNG = ".png"

var tilex, tiley int
var zonex, zoney int
var ext string = EXT_JPEG

func ExportSplitedImage(tokens []string) {
	if len(tokens) != 4 {
		fmt.Println("Usage: export tile_width tile_height source_path")
		return
	}

	tilex, _ = strconv.Atoi(tokens[1])
	tiley, _ = strconv.Atoi(tokens[2])
	dir := tokens[3]

	err := filepath.Walk(dir, func(filePath string, fileInfo os.FileInfo, err error) error {
		if fileInfo == nil {
			return err
		}
		if fileInfo.IsDir() {
			return nil
		}

		ext = path.Ext(fileInfo.Name())
		newDir := dir + "/" + strings.Split(fileInfo.Name(), ".")[0]
		fmt.Println(newDir)

		switch ext {
		case EXT_JPEG, EXT_PNG:
			parseFile(filePath, newDir)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}

func imageEncode(file *os.File, img image.Image) (err error) {
	if ext == EXT_JPEG {
		err = jpeg.Encode(file, img, &jpeg.Options{90})
	} else {
		err = png.Encode(file, img)
	}
	return
}

func imageDecode(file *os.File) (img image.Image, err error) {
	if ext == EXT_JPEG {
		img, err = jpeg.Decode(file)
	} else {
		img, err = png.Decode(file)
	}
	return
}

func parseFile(filePath string, dir string) {

	if !IsDir(dir) {
		os.Mkdir(dir, os.ModeDir)
	}

	f1, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	m1, err := imageDecode(f1) // jpeg.Decode(f1)
	if err != nil {
		panic(err)
	}
	bounds := m1.Bounds()

	width := bounds.Dx()
	height := bounds.Dy()

	zonex = width / tilex
	zoney = height / tiley
	if width%tilex > 0 {
		zonex += 1
	}
	if height%tiley > 0 {
		zoney += 1
	}

	for i := 0; i < zonex; i++ {
		for j := 0; j < zoney; j++ {
			zeroRect := image.Rectangle{image.Point{0, 0}, image.Point{tilex, tiley}}
			m := image.NewRGBA(zeroRect)
			//			white := color.RGBA{255, 255, 255, 255}
			rect := image.Rect(0, 0, tilex+(i*tilex), tiley+(j*tiley))
			pt := image.Pt(i*tilex, j*tiley)
			//			rect := image.Rectangle{image.Point{i * tilex, j * tiley}, image.Point{tilex + (i * tilex), tiley + (j * tiley)}}
			//			draw.Draw(m, zeroRect, &image.Uniform{white}, image.ZP, draw.Src)
			draw.Draw(m, rect, m1, pt, draw.Src)
			// draw.Draw(m, image.Rect(100, 200, 300, 600), m2, image.Pt(250, 60), draw.Src)

			key := fmt.Sprintf("%d_%d%s", i, j, ext)
			exportImg(dir+"/"+key, m)
		}
	}

	m1 = nil
}

func exportImg(fileName string, img image.Image) {
	filePath := fileName
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = imageEncode(f, img) // jpeg.Encode(f, img, &jpeg.Options{90})
	if err != nil {
		panic(err)
	}

	img = nil
	fmt.Printf("generate " + filePath + " ok\n")
}
