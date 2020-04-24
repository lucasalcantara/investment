package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func main() {
	tickers := []string{}
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


	response, err := http.Get("https://finbox.com/NYSE:FL/explorer/interest_coverage")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", string(contents))
	}
}
