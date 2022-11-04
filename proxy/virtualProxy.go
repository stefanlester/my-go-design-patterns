package main

import "fmt"

type Image interface {
	Draw()
}

type Bitmap struct {
	filename string
}

func (b *Bitmap) Draw() {
	fmt.Println("Drawing image", b.filename)
}

func NewBitmap(filename string) *Bitmap {
	fmt.Println("Loading image from", filename)
	return &Bitmap{filename: filename}
}

func DrawImage(image Image) {
	fmt.Println("About to draw the image")
	image.Draw()
	fmt.Println("Done drawing the image")
}
