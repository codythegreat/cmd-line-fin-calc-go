// a command line finance calculator in go
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/alpeb/go-finance/fin"
)

func main() {
	rate, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		panic(err)
	}
	periods, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}
	pmt, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		panic(err)
	}
	pv, err := strconv.ParseFloat(os.Args[4], 64)
	if err != nil {
		panic(err)
	}
	pmtType, err := strconv.ParseInt(os.Args[5], 10, 8)
	if err != nil {
		panic(err)
	}
	if pmtType == 0 {
		res, err := fin.FutureValue(rate, periods, pmt, pv, fin.PayBegin)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Rate: %v, Periods: %v, PMT: %v, PV: %v\nResult: %v\n", rate, periods, pmt, pv, res)
	} else if pmtType == 1 {
		res, err := fin.FutureValue(rate, periods, pmt, pv, fin.PayEnd)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Rate: %v, Periods: %v, PMT: %v, PV: %v\nResult: %v\n", rate, periods, pmt, pv, res)
	}
}
