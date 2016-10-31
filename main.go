package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/Akagi201/light"
	"github.com/dre1080/recover"
	"github.com/gohttp/logger"
	flags "github.com/jessevdk/go-flags"
	"github.com/rs/cors"
)

var opts struct {
	Host string `long:"host" default:"0.0.0.0" description:"ip to bind to"`
	Port uint16 `long:"port" default:"2222" description:"port to bind to"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	dumpReq, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DUMP request: %v", string(dumpReq))

	w.Header().Set("Httpdump-Author", "Akagi201")

	r.Write(w)
	//io.Copy(w, r.Body)
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		if !strings.Contains(err.Error(), "Usage") {
			log.Printf("error: %v\n", err.Error())
			os.Exit(1)
		} else {
			// log.Printf("%v\n", err.Error())
			os.Exit(0)
		}
	}

	app := light.New()
	app.Use(logger.New())
	recovery := recover.New(&recover.Options{
		Log: log.Print,
	})
	handleCORS := cors.Default().Handler
	app.Use(handleCORS)
	app.Use(recovery)

	app.Post("/", http.HandlerFunc(handler))

	log.Printf("httpdump listening at: %v:%v", opts.Host, opts.Port)
	app.Listen(fmt.Sprintf("%v:%d", opts.Host, opts.Port))
}
