package main

import (
	_ "embed"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var spvEndpoint = "https://spv.ncp.nathanferns.xyz/"

func Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	track := p.ByName("track")

	if track == "" {
		http.Error(w, "missing track!", http.StatusBadRequest)
		return
	}

	cover := spvEndpoint + track + "/cover/large"
	payload := RunPayload{
		Size: []int{330, 330},
		Files: []File{
			{
				Name: "main.ql",
				Code: ipodKittyQl,
			},
		},
		Assets: []interface{}{
			ImageAsset{
				Name: "kitty",
				Url:  "https://nathanielfernandes.b-cdn.net/silly/kitty.png",
			},
			ImageAsset{
				Name: "cover",
				Url:  cover,
			},
		},
	}

	canvasUrl, err := GetCanvasUrl(&payload)

	if err != nil {
		http.Error(w, "error generating url", http.StatusInternalServerError)
		return
	}

	fmt.Println("gen: ", cover)

	http.Redirect(w, r, canvasUrl, http.StatusFound)
}

//go:embed ipodkitty.ql
var ipodKittyQl string

func main() {
	router := httprouter.New()
	router.GET("/:track", Get)

	fmt.Println("listening on :80")
	if err := http.ListenAndServe("0.0.0.0:80", router); err != nil {
		panic(err)
	}
}
