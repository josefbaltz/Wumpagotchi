package main

import (
	"fmt"
	"image"
	"image/color"
	"net/http"

	"gopkg.in/go-playground/colors.v1"

	"github.com/disintegration/imaging"
)

// LeafedWumpus Leafifies a Wumpus
// This function requires 3 Arguments and returns an image.Image object
// BaseImageURL is a URL containing an image that will be used for the base layer
// SleepingLeaf is a boolean that determines which Leaf image to use, if false it will use the default leaf, if true it will use a modified image that is shifted down
// The intention of the shifted leaf is for use with the Sleeping Wumpus image so the leaf isn't floating above their head
// UserWumpus is the User's Wumpus, This function pulls the color from the Wumpus and Recolors the leaf accordingly
func LeafedWumpus(BaseImageURL string, SleepingLeaf bool, UserWumpus Wumpus) (WumpusImage image.Image) {
	var baseImage image.Image
	var leafImage image.Image
	var err error
	// Get the provided URL and decode it as an image.Image for use later
	baseURL, err := http.Get(BaseImageURL)
	if err != nil {
		fmt.Println("ERROR RETRIEVING IMAGE" + err.Error())
		return
	}
	defer baseURL.Body.Close()
	baseImage, _, err = image.Decode(baseURL.Body)
	if err != nil {
		fmt.Println(err)
	}

	if SleepingLeaf == false {
		// Get the Leaf URL and Decode it as an image.Image for use later
		leafURL, err := http.Get("https://orangeflare.me/imagehosting/Wumpagotchi/Leaf.png")
		if err != nil {
			fmt.Println("ERROR RETRIEVING IMAGE" + err.Error())
			return
		}
		defer baseURL.Body.Close()
		leafImage, _, err = image.Decode(leafURL.Body)
		if err != nil {
			fmt.Println(err)
		}
		if err != nil {
			fmt.Println(err)
		}
		// Take the Base 10 Color value stored in the User's Wumpus and format it as a Base 16 and then pad it to 6 characters
		// Additionally Add a Hashtag to the front so we can then convert it to an RGBA Object
		pc, err := colors.Parse("#" + fmt.Sprintf("%06x", UserWumpus.Color))
		if err != nil {
			fmt.Println("#" + fmt.Sprintf("%06x", UserWumpus.Color))
			fmt.Println(err)
			return
		}
		WumpusColors := pc.ToRGBA()
		WumpusColor := color.NRGBA{
			R: WumpusColors.R,
			G: WumpusColors.G,
			B: WumpusColors.B,
			A: 255,
		}
		LeafColor := color.NRGBA{
			R: 124,
			G: 176,
			B: 81,
			A: 255,
		}
		// Since the Leaf iamge is a pre-determined image and we know EXACTLY where the pixels need to be replaced
		// instead of scanning the entire image we scan only the area where we want pixels to be replaced with the new color
		for x := 618; x < 821; x++ {
			for y := 168; y < 323; y++ {
				r, g, b, a := leafImage.At(x, y).RGBA()
				PixelColor := color.NRGBA{
					R: uint8(r),
					G: uint8(g),
					B: uint8(b),
					A: uint8(a),
				}
				if PixelColor == LeafColor {
					leafImage.(*image.NRGBA).SetNRGBA(x, y, WumpusColor)
				}
			}
		}
	} else {
		// All of the code in this else statement is the same as above besides the singular line below
		// The only thing that changes is a different URL
		leafURL, err := http.Get("https://orangeflare.me/imagehosting/Wumpagotchi/AsleepLeaf.png")
		if err != nil {
			fmt.Println("ERROR RETRIEVING IMAGE" + err.Error())
			return
		}
		defer baseURL.Body.Close()
		leafImage, _, err = image.Decode(leafURL.Body)
		if err != nil {
			fmt.Println(err)
		}
		pc, err := colors.Parse("#" + fmt.Sprintf("%06x", UserWumpus.Color))
		if err != nil {
			fmt.Println("#" + fmt.Sprintf("%06x", UserWumpus.Color))
			fmt.Println(err)
			return
		}
		WumpusColors := pc.ToRGBA()
		WumpusColor := color.NRGBA{
			R: WumpusColors.R,
			G: WumpusColors.G,
			B: WumpusColors.B,
			A: 255,
		}
		LeafColor := color.NRGBA{
			R: 124,
			G: 176,
			B: 81,
			A: 255,
		}
		for x := 618; x < 821; x++ {
			for y := 312; y < 467; y++ {
				r, g, b, a := leafImage.At(x, y).RGBA()
				PixelColor := color.NRGBA{
					R: uint8(r),
					G: uint8(g),
					B: uint8(b),
					A: uint8(a),
				}
				if PixelColor == LeafColor {
					leafImage.(*image.NRGBA).SetNRGBA(x, y, WumpusColor)
				}
			}
		}
	}

	// Finally we Overlay the recolored leaf image over the base image provided at the start and then return this as an image.Image
	return imaging.Overlay(baseImage, leafImage, image.Pt(0, 0), 1.0)
}
