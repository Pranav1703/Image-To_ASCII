package cmd

import (
	"flag"
	"fmt"
	"log"

)

const USAGE = `
---------- Flags -----------
-image    = path the of the input image
-width    = set width to resize the image. Default value is 0 and image will be unchanged
-height   = set height to resize the image. Default value is 0 and image will be unchanged		
-inverted = 

`

func ParseCmds() (string,int,int,bool){
	var imagePath *string = flag.String("image","","Image path")
	var inverted *bool = flag.Bool("inverted",false,"Inverted Mode")
	var width *int = flag.Int("width",0,"Set the width of the image")
	var height *int = flag.Int("height",0,"Set the height of the image")

	flag.Usage = func() {fmt.Print(USAGE)}

	flag.Parse()
	if *imagePath ==""{
		log.Fatal("Image path not specified. Terminating program")
	}

	// return *imagePath,*inverted
	return *imagePath,*width,*height,*inverted
}