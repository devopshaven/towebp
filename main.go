package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/chai2010/webp"
)

func main() {
	// var buf bytes.Buffer
	// var width, height int
	// var data []byte
	// var err error

	wd, _ := os.Getwd()

	filepath.Walk(wd, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		switch strings.ToLower(filepath.Ext(path)) {
		case ".jpg", ".jpeg":
			// fmt.Println("JPEG file found: " + path)
			EncodeWebp(DecodeJPEG(path), path)
		case ".png":
			EncodeWebp(DecodePNG(path), path)
		}

		// fmt.Println(path, info.Size())
		return nil
	})
}

func DecodePNG(fp string) image.Image {
	f, _ := os.Open(fp)

	// Decode webp
	m, err := png.Decode(f)
	if err != nil {
		panic(err)
	}

	return m
}

func DecodeJPEG(fp string) image.Image {
	f, _ := os.Open(fp)

	// Decode webp
	m, err := jpeg.Decode(f)
	if err != nil {
		panic(err)
	}

	return m
}

func EncodeWebp(img image.Image, filename string) {

	outFileName := filename[:strings.LastIndex(filename, ".")]

	var buf bytes.Buffer

	// Encode lossless webp
	if err := webp.Encode(&buf, img, &webp.Options{Lossless: true}); err != nil {
		log.Println(err)
	}
	if err := ioutil.WriteFile(outFileName+".webp", buf.Bytes(), 0666); err != nil {
		log.Println(err)
	}

	fmt.Printf("Save %s.webp ok\n", outFileName)
}

func getmeta() {
	// // Load file data
	// if data, err = ioutil.ReadFile("./testdata/1_webp_ll.webp"); err != nil {
	// 	log.Println(err)
	// }

	// // GetInfo
	// if width, height, _, err = webp.GetInfo(data); err != nil {
	// 	log.Println(err)
	// }
	// fmt.Printf("width = %d, height = %d\n", width, height)

	// // GetMetadata
	// if metadata, err := webp.GetMetadata(data, "ICCP"); err != nil {
	// 	fmt.Printf("Metadata: err = %v\n", err)
	// } else {
	// 	fmt.Printf("Metadata: %s\n", string(metadata))
	// }
}
