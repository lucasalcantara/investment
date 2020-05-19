package main

import (
	"fmt"
	"reflect"
	"strings"
)

func buildRatioTable(ratios []Ratio) string {
	table := strings.Builder{}
	table.WriteString(`
	<table border=1>	
		<tr>
			<td></td>
`)

	for _, r := range ratios {
		table.WriteString("<td>" + r.Ticker + "</td>")
	}

	table.WriteString("</tr>")

	// mapping the ratios attributes
	rows := make(map[string]map[string][]string)
	ratiosIterator(ratios, func(t reflect.Type, v reflect.Value, i int, j int) {
		n := t.Field(i).Name
		f := t.Field(i).Type.Field(j)
		ti := f.Tag.Get("title")

		if _, ok := rows[n]; !ok {
			rows[n] = make(map[string][]string)
		}

		rows[n][ti] = append(rows[n][ti], fmt.Sprintf("%.2f", v.Field(i).Field(j).Float()))
	})

	// appending the ratio values to table
	// as it uses the struct Ratio to know all the attributes to be appended just an array with one empty Ratio
	// is necessary to be passed.
	ratiosIterator([]Ratio{{}}, func(t reflect.Type, v reflect.Value, i int, j int) {
		n := t.Field(i).Name
		if j == 0 {
			table.WriteString("<tr><td><b>" + n + "</b></td></tr>")
		}

		tag := t.Field(i).Type.Field(j).Tag.Get("title")
		table.WriteString("<tr><td>" + tag + "</td>")
		for _, v := range rows[n][tag] {
			table.WriteString("<td>" + v + "</td>")
		}
		table.WriteString("</tr>")

		table.WriteString("<tr></tr>")
	})

	table.WriteString("</table>")
	return table.String()
}

func ratiosIterator(ratios []Ratio, f func(reflect.Type, reflect.Value, int, int)) {
	for _, r := range ratios {
		t := reflect.TypeOf(r)
		v := reflect.ValueOf(r)
		for i := 0; i < t.NumField(); i++ {
			if t.Field(i).Type.Kind() == reflect.Struct {
				for j := 0; j < t.Field(i).Type.NumField(); j++ {
					f(t, v, i, j)
				}
			}
		}
	}
}
