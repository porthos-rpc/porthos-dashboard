package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/facebookgo/httpdown"
	"github.com/porthos-rpc/porthos-dashboard/handlers"
	"github.com/porthos-rpc/porthos-dashboard/metrics"
)

func defaultValue(a, b string) string {
	if len(a) == 0 {
		return b
	}

	return a
}

func main() {
	bindAddress := flag.String("bind", defaultValue(os.Getenv("BIND_ADDRESS"), ":3000"), "Bind Address.")
	brokerURL := flag.String("broker", defaultValue(os.Getenv("BROKER_URL"), "amqp://"), "Broker URL.")

	flag.Parse()

	metrics.NewCollector(*brokerURL)

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", handlers.IndexHandler)

	server := &http.Server{
		Addr:           *bindAddress,
		Handler:        serveMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	hd := &httpdown.HTTP{
		StopTimeout: 1 * time.Second,
		KillTimeout: 1 * time.Second,
	}

	fmt.Printf("Listening to %s\n", *bindAddress)
	fmt.Println("Hit CTRL-C to exit...")

	if err := httpdown.ListenAndServe(server, hd); err != nil {
		panic(err)
	}
}
