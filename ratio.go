package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"math"
	"strconv"
	"strings"
	"sync"
)

const (
	finvizQuerySelector = ".table-dark-row-cp > td > a"
	valuationUrl        = "https://finviz.com/screener.ashx?v=121&t=%s"
	financialUrl        = "https://finviz.com/screener.ashx?v=161&t=%s"
)

type Ratio struct {
	Ticker        string
	Valuation     Valuation
	Profitability Profitability
	Liquidity     Liquidity
	Debt          Debt
	Efficiency    Efficiency
}

type Valuation struct {
	PE  float64 `title:"Price to Earning (P/E)"`
	PEG float64 `title:"Price/Earnings-to-Growth (PEG)"`
	PS  float64 `title:"Price To Sales (P/S)"`
	PB  float64 `title:"Price to Book (P/B)"`
	DY  float64 `title:"Dividend Yield"`
	DP  float64 `title:"Dividend Payout"`
}

type Profitability struct {
	ROA float64 `title:"Return on assets (ROA)"`
	ROE float64 `title:"Return on Equity (ROE)"`
	ROI float64 `title:"Return on Investment (ROI)"`
	PM  float64 `title:"Profit Margin"`
}

type Liquidity struct {
	CR float64 `title:"Currently ratio"`
	QR float64 `title:"Quick ratio"`
}

type Debt struct {
	DE float64 `title:"Debit to Equity"`
	IC float64 `title:"Interest Coverage"`
}

type Efficiency struct {
	AT float64 `title:"Assert turnover"`
	IT float64 `title:"Inventory turnover"`
}

type finvizValuationRatio struct {
	PE  float64 // Price to Earning (P/E)
	PEG float64
	PS  float64 // Price To Sales (P/S)
	PB  float64 // Price to Book (P/B)
}

type finvizFinancialRatio struct {
	DY    float64 // Dividend yield
	ROA   float64 // Return on assets (ROA)
	ROE   float64 // Return on Equity (ROE)
	ROI   float64 // Return on Investment (ROI)
	PM    float64 // Profit Margin
	CR    float64 // Current ratio
	QR    float64 // Quick ratio
	DE    float64 // Debit to equity
	Price float64 // Price
}

func FetchRatios(tickers ...string) []Ratio {
	ratios := make([]Ratio, len(tickers), len(tickers))

	wg := sync.WaitGroup{}

	for i, t := range tickers {
		wg.Add(1)

		go func(index int, ticker string) {
			defer wg.Done()
			valuation := fetchValuation(ticker)
			financial := fetchFinancial(ticker)

			ratios[index] = Ratio{
				Ticker: ticker,
				Valuation: Valuation{
					PE:  valuation.PE,
					PEG: valuation.PEG,
					PS:  valuation.PS,
					PB:  valuation.PB,
					DY:  financial.DY,
				},
				Profitability: Profitability{
					ROA: financial.ROA,
					ROE: financial.ROE,
					ROI: financial.ROI,
					PM:  financial.PM,
				},
				Liquidity: Liquidity{
					CR: financial.CR,
					QR: financial.QR,
				},
				Debt: Debt{
					DE: financial.DE,
				},
			}
		}(i, t)
	}

	wg.Wait()

	return ratios
}

func fetchValuation(ticker string) finvizValuationRatio {
	ratioPositions := map[int]float64{
		3: 0,
		5: 0,
		6: 0,
		7: 0,
	}

	collectRatios(valuationUrl, ticker, ratioPositions)

	return finvizValuationRatio{
		PE:  ratioPositions[3],
		PEG: ratioPositions[5],
		PS:  ratioPositions[6],
		PB:  ratioPositions[7],
	}
}

func collectRatios(url, ticker string, ratioPositions map[int]float64) {
	c := colly.NewCollector()

	c.OnHTML(finvizQuerySelector, func(e *colly.HTMLElement) {
		if _, ok := ratioPositions[e.Index]; ok {
			v, err := strconv.ParseFloat(strings.Trim(e.Text, "%"), 32)
			if err != nil {
				v = 0
			}

			ratioPositions[e.Index] = math.Round(v*100) / 100
		}
	})

	c.Visit(fmt.Sprintf(url, ticker))
}

func fetchFinancial(ticker string) finvizFinancialRatio {
	ratioPositions := map[int]float64{
		3:  0,
		4:  0,
		5:  0,
		6:  0,
		7:  0,
		8:  0,
		10: 0,
		13: 0,
		15: 0,
	}

	collectRatios(financialUrl, ticker, ratioPositions)

	return finvizFinancialRatio{
		DY:    ratioPositions[3],
		ROA:   ratioPositions[4],
		ROE:   ratioPositions[5],
		ROI:   ratioPositions[6],
		CR:    ratioPositions[7],
		QR:    ratioPositions[8],
		DE:    ratioPositions[10],
		PM:    ratioPositions[13],
		Price: ratioPositions[15],
	}
}
