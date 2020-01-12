package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"flag"
	"github.com/chai2010/webp"
	"image/jpeg"
)

func main() {

	name := flag.String("path", "", "Enter the jpeg image you want to convert into webp")
	q := flag.Int("quality", 40, "Enter the quality measure you want to use. The bigger the more accurate image")

	out := flag.String("out", "output.webp", "Location to output your image. Default output.webp")

	flag.Parse()

	var buf bytes.Buffer
	// var width, height int
	var data []byte
	var err error

	// Load file data
	if data, err = ioutil.ReadFile(*name); err != nil {
		panic(err)
	}

	m, err := jpeg.Decode(bytes.NewReader(data))
	if err != nil {
		fmt.Printf("the error is: %#v", err)
	}
	// Encode lossless webp
	if err = webp.Encode(&buf, m, &webp.Options{
		Lossless: false,
		Quality:  float32(*q),
		Exact:    false,
	}); err != nil {
		log.Println(err)
	}
	if err = ioutil.WriteFile(*out, buf.Bytes(), 0666); err != nil {
		log.Println(err)
	}

	fmt.Println("Save output.webp ok")
}
