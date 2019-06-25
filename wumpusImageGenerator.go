package main

import (
	"image"
	"net/http"

	"github.com/disintegration/imaging"
)

//LeafedWumpus Leafifies a Wumpus
func LeafedWumpus(BaseImageURL string, SleepingLeaf bool) {
	var baseImage image.Image
	var leafImage image.Image
	baseURL, _ := http.Get(BaseImageURL)
	defer baseURL.Body.Close()
	baseImage, _, _ = image.Decode(baseURL.Body)

	if SleepingLeaf == false {
		leafURL, _ := http.Get("https://orangeflare.me/imagehosting/Wumpagotchi/Leaf.png")
		defer baseURL.Body.Close()
		leafImage, _, _ = image.Decode(leafURL.Body)
	} else {
		leafURL, _ := http.Get("https://orangeflare.me/imagehosting/Wumpagotchi/AsleepLeaf.png")
		defer baseURL.Body.Close()
		leafImage, _, _ = image.Decode(leafURL.Body)
	}

	imaging.PasteCenter(baseImage, leafImage)
}
