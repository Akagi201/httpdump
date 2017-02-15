package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/Akagi201/light"
	log "github.com/Sirupsen/logrus"
	"github.com/dre1080/recover"
	flags "github.com/jessevdk/go-flags"
	"github.com/pressly/lg"
	"github.com/rs/cors"
)

var opts struct {
	ListenAddr string `long:"listen" default:"0.0.0.0:8327" description:"HTTP address to listen at"`
	LogLevel   string `long:"log_level" default:"info" description:"log level"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	dumpReq, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Fatal(err)
	}
	log.Debugf("DUMP request: %v", string(dumpReq))

	w.Header().Set("Httpdump-Author", "Akagi201")

	r.Write(w)
}

func main() {
	parser := flags.NewParser(&opts, flags.HelpFlag|flags.PassDoubleDash|flags.IgnoreUnknown)

	_, err := parser.Parse()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}

	if level, err := log.ParseLevel(strings.ToLower(opts.LogLevel)); err != nil {
		log.SetLevel(level)
	}

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger := log.New()

	root := light.New()
	root.Use(lg.RequestLogger(logger))
	recovery := recover.New(&recover.Options{
		Log: log.Print,
	})
	handleCORS := cors.Default().Handler
	root.Use(handleCORS)
	root.Use(recovery)

	root.HandleAll("/*path", http.HandlerFunc(handler))

	log.Infof("httpdump listening at: %v", opts.ListenAddr)
	http.ListenAndServe(opts.ListenAddr, root)
}
