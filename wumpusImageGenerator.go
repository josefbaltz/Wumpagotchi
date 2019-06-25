package main

import (
	"fmt"
	"image"
	"image/color"
	"net/http"
	"strconv"

	"gopkg.in/go-playground/colors.v1"

	"github.com/disintegration/imaging"
)

//LeafedWumpus Leafifies a Wumpus
func LeafedWumpus(BaseImageURL string, SleepingLeaf bool, UserWumpus Wumpus) (WumpusImage image.Image) {
	var baseImage image.Image
	var leafImage image.Image
	var err error
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

		pc, _ := colors.Parse("#" + strconv.FormatInt(int64(UserWumpus.Color), 16))
		WumpusColors := pc.ToRGBA()
		WumpusColor := color.NRGBA{
			R: WumpusColors.R,
			G: WumpusColors.G,
			B: WumpusColors.B,
			A: 255,
		}
		for x := 618; x < 821; x++ {
			for y := 168; y < 323; y++ {
				leafImage.(*image.NRGBA).SetNRGBA(x, y, WumpusColor)
			}
		}
	} else {
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
	}

	return imaging.Overlay(baseImage, leafImage, image.Pt(0, 0), 1.0)
}
