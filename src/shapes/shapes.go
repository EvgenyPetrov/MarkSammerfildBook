package main

import (
    "image/color"
    "image/draw"
    "image"
    "log"
    "runtime"
    "path/filepath"
    "os"
    "strings"
    "image/jpeg"
    "image/png"
    "fmt"
)

func main() {
    FilledImage(150, 5000, color.Black)
}

func FilledImage(width, heigth int, fill color.Color) draw.Image {
    if fill == nil {
        fill = color.Black
    }
    width = saneLength(width)
    heigth = saneLength(heigth)
    img := image.NewRGBA(image.Rect(0, 0, width, heigth))
    draw.Draw(img, img.Bounds(), &image.Uniform{fill}, image.ZP, draw.Src)
    return img
}

var saneLength, saneRadius, saneSides func(int) int

func init()  {
    saneLength = makeBoundedIntFunc(1, 4096)
    saneRadius = makeBoundedIntFunc(1, 1024)
    saneSides = makeBoundedIntFunc(3, 60)
}

func makeBoundedIntFunc(min, max int) func(int) int {
    return func(x int) int {
        valid := x
        switch {
        case x < min:
            valid = min
        case x > max:
            valid = max
        }
        if valid != x {
            log.Printf("%s(): replaced %d with %d\n", caller(1), x, valid)
        }
        return valid
    }
}

func caller(steps int) string {
    name := "?"
    if pc, _, _, ok := runtime.Caller(steps + 1); ok {
        name = filepath.Base(runtime.FuncForPC(pc).Name())
    }
    return name
}

func DrowShapes(img draw.Image, x, y int, shapes ...Shaper) error {
    for _, shape := range shapes {
        if err := shape.Draw(img, x, y); err != nil {
            return err
        }
    }
    return nil
}

func SaveImage(img image.Image, filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    switch strings.ToLower(filepath.Ext(filename)) {
    case ".jpg", ".jpeg":
        return jpeg.Encode(file, img, nil)
    case ".png":
        return png.Encode(file, img)
    }
    return fmt.Errorf("shapes.SaveImage(): '%s' has unrecognized suffix", filename)
}
