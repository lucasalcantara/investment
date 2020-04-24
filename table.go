package investment

import (
	"fmt"
	"strings"
)

func buildRatioTable(ratios []Ratio) string {
	table := `
	<table border=1>	
		<tr>
			<td>VALUATION</td>
			%s
			<td>A Number that is _______ is better</td>
		</tr>
		<tr>
			<td>Price to Earning (P/E)</td>
			%s
			<td>Lower</td>
		</tr>
		<tr>
			<td>PEG</td>
			%s
			<td>Below 1 generally</td>
		</tr>
		<tr>
			<td>Price To Sales (P/S)</td>
			%s
			<td>Lower</td>
		</tr>
		<tr>
			<td>Price to Book (P/B)</td>
			%s
			<td>Lower</td>
		</tr>
		<tr>
			<td>Dividend Yield</td>
			%s
			<td>Higher</td>
		</tr>
		<tr>
			<td>Dividend Payout</td>
			%s
			<td>Lower (Careful > 70)</td>
		</tr>
		<tr>
		</tr>
		<tr>
			<td><b>PROFITABILITY</b></td>
		</tr>
		<tr>
			<td>Return on assets (ROA)</td>
			%s
			<td>Higher</td>
		</tr>
		<tr>
			<td>Return on Equity (ROE)</td>
			%s
			<td>Higher</td>
		</tr>
		<tr>
			<td>Return on Investment (ROI)</td>
			%s
			<td>Higher</td>
		</tr>
		<tr>
			<td>Profit Margin</td>
			%s
			<td>Higher</td>
		</tr>
		<tr>
		</tr>
		<tr>
			<td><b>LIQUIDITY</b></td>
		</tr>
		<tr>
			<td>Current Ratio</td>
			%s
			<td>Greater than 1.0</td>
		</tr>
		<tr>
			<td>Quick Ratio</td>
			%s
			<td>Greater than 1.0</td>
		</tr>
		<tr>
		</tr>
		<tr>
			<td><b>DEBT</b></td>
		</tr>
		<tr>
			<td>Debit to Equity</td>
			%s
			<td>Lower is better</td>
		</tr>
		<tr>
			<td>Interest Coverage</td>
			%s
			<td>Lower than 1 is bad</td>
		</tr>
		<tr>
		</tr>
		<tr>
			<td><b>EFFICIENCY (OPERATING)</b></td>
		</tr>
		<tr>
			<td>Asset Turnover</td>
			%s
			<td>Higher</td>
		</tr>
		<tr>
			<td>Inventory Turnover</td>
			%s
			<td>Higher</td>
		</tr>
	</table>
`

	table = fmt.Sprintf(table, buildTickers(ratios),
		buildPEs(ratios), buildPEGs(ratios), buildPSs(ratios), buildPBs(ratios), buildDYs(ratios), buildDPs(ratios),
		buildROAs(ratios), buildROEs(ratios), buildROIs(ratios), buildPMs(ratios),
		buildCRs(ratios), buildQRs(ratios), buildDEs(ratios), buildDPs(ratios), buildDPs(ratios), buildDPs(ratios))

	return table
}

func buildTickers(ratios []Ratio) string {
	print := strings.Builder{}
	for _, r := range ratios {
		print.WriteString("<td>" + r.Ticker + "</td>")
	}

	return print.String()
}

func buildPEs(ratios []Ratio) string {
	print := strings.Builder{}
	for _, r := range ratios {
		print.WriteString(fmt.Sprintf("<td>%.2f</td>", r.Valuation.PE))
	}

	return print.String()
}

func buildPEGs(ratios []Ratio) string {
	print := strings.Builder{}
	for _, r := range ratios {
		print.WriteString(fmt.Sprintf("<td>%.2f</td>", r.Valuation.PEG))
	}

	return print.String()
}

func buildPSs(ratios []Ratio) string {
	print := strings.Builder{}
	for _, r := range ratios {
		print.WriteString(fmt.Sprintf("<td>%.2f</td>", r.Valuation.PS))
	}

	return print.String()
}

func buildPBs(ratios []Ratio) string {
	print := strings.Builder{}
	for _, r := range ratios {
		print.WriteString(fmt.Sprintf("<td>%.2f</td>", r.Valuation.PB))
	}

	return print.String()
}

func buildDYs(ratios []Ratio) string {
	print := strings.Builder{}
	for _, r := range ratios {
		print.WriteString(fmt.Sprintf("<td>%.2f</td>", r.Valuation.DY))
	}

	return print.String()
}

func buildDPs(ratios []Ratio) string {
	print := strings.Builder{}
	for range ratios {
		print.WriteString(fmt.Sprintf("<td></td>"))
	}

	return print.String()
}

func buildROAs(ratios []Ratio) string {
	print := strings.Builder{}
	for _, r := range ratios {
		print.WriteString(fmt.Sprintf("<td>%.2f</td>", r.Profitability.ROA))
	}

	return print.String()
}

func buildROEs(ratios []Ratio) string {
	print := strings.Builder{}
	for _, r := range ratios {
		print.WriteString(fmt.Sprintf("<td>%.2f</td>", r.Profitability.ROE))
	}

	return print.String()
}

func buildROIs(ratios []Ratio) string {
	print := strings.Builder{}
	for _, r := range ratios {
		print.WriteString(fmt.Sprintf("<td>%.2f</td>", r.Profitability.ROI))
	}

	return print.String()
}

func buildPMs(ratios []Ratio) string {
	print := strings.Builder{}
	for _, r := range ratios {
		print.WriteString(fmt.Sprintf("<td>%.2f</td>", r.Profitability.PM))
	}

	return print.String()
}

func buildCRs(ratios []Ratio) string {
	print := strings.Builder{}
	for _, r := range ratios {
		print.WriteString(fmt.Sprintf("<td>%.2f</td>", r.Liquidity.CR))
	}

	return print.String()
}

func buildQRs(ratios []Ratio) string {
	print := strings.Builder{}
	for _, r := range ratios {
		print.WriteString(fmt.Sprintf("<td>%.2f</td>", r.Liquidity.QR))
	}

	return print.String()
}

func buildDEs(ratios []Ratio) string {
	print := strings.Builder{}
	for _, r := range ratios {
		print.WriteString(fmt.Sprintf("<td>%.2f</td>", r.Debt.DE))
	}

	return print.String()
}
