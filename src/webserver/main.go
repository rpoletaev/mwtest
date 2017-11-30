package main

import (
	"encoding/json"
	"net/http"

	"github.com/rpoletaev/mwtest/src/types"

	"github.com/rpoletaev/mwtest/src/middleware"
)

func main() {
	http.Handle("/",
		middleware.Autch.Authorization(
			middleware.Compress.Compression(
				// 	middleware.Convertor.Convertation(
				http.HandlerFunc(MainHandlerFunc),
			),
		),
	// 	),
	)

	middleware.Autch.SetBasicAuth(`myname`, `pass123`)
	http.ListenAndServe(":3000", nil)
}

func MainHandlerFunc(wr http.ResponseWriter, rq *http.Request) {
	var err error
	var rsp *types.Response
	var buf []byte

	rsp = &types.Response{
		Name:          "Гугл",
		Default:       true,
		Url:           `https://www.goole.com`,
		GmailLogo:     `https://cdn.worldvectorlogo.com/logos/gmail.svg`,
		HangoutslLogo: `https://cdn.worldvectorlogo.com/logos/google-hangouts.svg`,
		DriveLogo:     `https://cdn.worldvectorlogo.com/logos/google-drive.svg`,
	}
	if buf, err = json.Marshal(rsp); err != nil {
		wr.WriteHeader(http.StatusInternalServerError)
		return
	}
	wr.Header().Add("Content-Type", "application/json")
	wr.WriteHeader(http.StatusOK)
	wr.Write(buf)
}
