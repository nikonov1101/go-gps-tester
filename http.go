package main

import (
	"github.com/Unknwon/macaron"
	"fmt"
	)

type jsonResponse struct {
	Code 	 int 		`json:"code"`
	Message  string 	`json:"message"`
	Extended interface{} `json:"extended"` 
}

func HttpIndex(ctx *macaron.Context) {
	ctx.HTML(200, "index")
    return
}

func HttpPutGps(ctx *macaron.Context) {
	lat := ctx.QueryFloat64("lat")
	lon := ctx.QueryFloat64("lon")
	srcAddr := ctx.RemoteAddr();

	fmt.Printf("%s lat = %.f2; lon = %.f2", srcAddr, lat, lon)
	// todo: rewrite using chan
	go saveGps(lat, lon)

	ctx.JSON(200, &jsonResponse{
		Code: 0,
		Message: "OK",
		Extended: fmt.Sprintf("lat=%v; lon=%v;", lat, lon),
	})
}