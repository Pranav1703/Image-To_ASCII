package main

import (
	"ImageToAscii/cmd"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"github.com/nfnt/resize"
	"log"
	"fmt"
)

// var Ascii2 = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. "
var Ascii []rune = []rune{'.',':','░','▒','▓','█'}

func main(){

	imagePath,width,height,inverted := cmd.ParseCmds()

	file,err := os.Open(imagePath)
	if err!=nil{
		log.Fatalln(err)

	}
	defer file.Close()
	
	image.RegisterFormat("jpeg","jpeg",jpeg.Decode,jpeg.DecodeConfig)

	img,_,err := image.Decode(file)

	if err!=nil{
		log.Fatalln("Couldnt decode file: ",err)
	}


	outFile, err := os.Create("output.txt")
    if err != nil {
        fmt.Println("couldnt create outfile: ",err)
    }
    defer outFile.Close()

	//resize img

	newImg := imageResize(img,width,height)

	b := newImg.Bounds()

	//new image with bounds b in which each pixel has value (r,g,b,a) = (0,0,0,0) 
	imgSet :=image.NewRGBA(b) 


	for y := b.Min.Y;y<b.Max.Y;y++{

		var row string = "" 

		for x := b.Min.X;x<b.Max.X;x++ {
			
			currentPixel := newImg.At(x,y)
			r, g, b, _ := currentPixel.RGBA()

			// lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			lum := 0.2126 * float64(r) + 0.7152 *float64(g) + 0.0722*float64(b)

			invertedLum := 65535 - lum

			// creates a new grey pixel 
			greyPixel := color.Gray{uint8(lum / 256)}
			

			// sets the pixel to newGreyPixel at (x,y) in imgSet
			imgSet.Set(x,y,greyPixel)

			// uses ascii2 series of characters
			// var i int = int(lum / 3.65)
			// fmt.Print(string(densityAscii2[i]))
			
			// uses ascii series of characters

			if inverted {
				invertedGreyPixel := color.Gray{uint8(invertedLum/256)}
				var index int = int(float64(invertedGreyPixel.Y) / 42.6)
				row += string(Ascii[index])
			}else{
				var index int = int(float64(greyPixel.Y) / 42.6)
				row += string(Ascii[index])
			}

		}

		fmt.Println(row)

		// _,err := outFile.WriteString(row)
		// if err!=nil{
		// 	log.Fatalln(err)
		// }
		// fmt.Println("no of lines written to output.txt?",lines)

	} 

}

func imageResize(img image.Image,width int , height int)(image.Image){
	return resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
}



/*

formulas to convert rgba to greyscale 

1.Y = 0.299 ∙ Red + 0.587 ∙ Green + 0.114 ∙ Blue
2.Y = (0.2126 * R + 0.7152 * G + 0.0722 * B)

algorithm(might have to change)
1.resize the image to desirable width and height while maintaining the aspect ratio
2.convert the image to grey scale
3.convert the each grey scale pixel to ascii character
*/
