package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func main() {
	tickers := os.Args[1:]
	ratios := FetchRatios(tickers...)
	for _, r := range ratios {
		fmt.Println(r.Ticker,
			"valuation:", r.Valuation,
			"profitability:", r.Profitability,
			"liquidity:", r.Liquidity)
	}

	go func() {
		time.Sleep(time.Second)
		exec.Command("open", "http://localhost:8080").Run()
	}()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, buildRatioTable(ratios))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
