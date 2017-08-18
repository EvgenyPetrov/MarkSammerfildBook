package main

import (
    "fmt"
    "os"
)

type Font struct {
    family      string
    size        int
}

func New(family string, size int) (*Font, error) {
    if family == "" || family == " " {
        fmt.Errorf("Cannot create font with nil family")
    }
    ValidSize(&size)
    return &Font{family, size} , nil
}
func ValidSize(size *int) {
    if *size > 144 {
        *size = 144
    } else if *size < 5 {
        *size = 5
    }
}
func (font *Font) SetFamily(family string) error {
    if family == "" || family == " " {
        return fmt.Errorf("Cannot create font with nil family")
    } else {
        font.family = family
        return nil
    }
}
func (font *Font) SetSize(size int) {
    ValidSize(&size)
    font.size = size
}
func (font *Font) GetSize() int {
    return font.size
}
func (font *Font) GetFamily() string {
    return font.family
}
func (font *Font) String() string {
    return fmt.Sprintf("{font-family:\t\"%s\";\tfont-size:\t%dpt;}", font.GetFamily(), font.GetSize())
}

func main() {
    titleFont, err := New("serif", 11)
    if err != nil {
        fmt.Println(err)
        os.Exit(0)
    }
    titleFont.SetFamily("Helvetica")
    titleFont.SetSize(20)
    fmt.Println(titleFont)
}
