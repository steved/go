package main

import (
	"flag"
	"log"

	"github.com/steved/go/context"
	"github.com/steved/go/web"
)

var version string

func getVersion() string {
	if version == "" {
		return "none"
	}
	return version
}

func main() {
	flagData := flag.String("data", "data",
		"The location to use for the data store")
	flagAddr := flag.String("addr", ":8067",
		"The address that the HTTP server will bind")
	flagAdmin := flag.Bool("admin", false,
		"If allowing admin level requests")
	flag.Parse()

	ctx, err := context.Open(*flagData)
	if err != nil {
		log.Panic(err)
	}
	defer ctx.Close()

	log.Panic(web.ListenAndServe(*flagAddr, *flagAdmin, getVersion(), ctx))
}
