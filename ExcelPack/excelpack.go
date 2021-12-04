package ExcelPack

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

type Date struct {
	day  int
	mon  int
	year int
}

type CarDetails struct {
	CarName   string
	CarAmount float64
}

var f *excelize.File
var err error

func GenerateTotalAmount(sheetList map[string]int) map[string]float64 {
	sum := 0.0
	m := map[string]float64{}
	//var cd []CarDetails
	for i, _ := range sheetList {
		cols, err := f.GetCols(i)
		if err != nil {
			m[i] = 0
			return m
		} else {
			for _, col := range cols[6:7] {
				for _, colCell := range col[6:] {
					a, _ := strconv.Atoi(colCell)
					sum = sum + float64(a)

				}
			}

			m[i] = sum

			sum = 0.0
		}

	}
	//	fmt.Println(cd)
	//	fmt.Println("Sheet Count", sheetCount, "total kharcha", total)
	return m
}

func GetSheetList() map[string]int {
	m := map[string]int{}
	f, err = excelize.OpenFile("simple4.xlsx")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Ronald")

	for _, r := range f.GetSheetList() {

		m[r] = 0

	}
	return m

}

func GenerateIndex(sheetname string, from string, to string, cols [][]string) (int, int) {
	fmt.Println("I  am generating index")
	var sindex, eindex int
	layout := "2006-01-02"
	fromParse, _ := time.Parse(layout, from)
	toParse, _ := time.Parse(layout, to)
	fromDate := fromParse.Format("01-02-06")
	toDate := toParse.Format("01-02-06")
	fmt.Println(fromDate, toDate)
	for _, col := range cols[3:4] {

		for i, colCell := range col[6:] {
			//fmt.Println(sheetname, i, colCell)
			if fromDate >= colCell {
				fmt.Println(fromDate, colCell, fromDate >= colCell)
				sindex = 256
			} else {
				sindex = i + 5
				fmt.Println(sindex)
				break
			}
		}

		//fmt.Println(sheetname, sindex, eindex)

	}
	for _, col1 := range cols[3:4] {

		for i, colCell := range col1[6:] {
			//fmt.Println(sheetname, i, colCell)
			if toDate >= colCell {
				fmt.Println(toDate, colCell, toDate >= colCell)
				eindex = 256
			} else {
				eindex = i + 6

				fmt.Println(sindex)
				break
			}
		}

	}
	fmt.Println(sheetname, sindex, eindex)
	return sindex, eindex
}
func GenerateTotalAmountWithDate(sheetList map[string]int, from string, to string) map[string]float64 {
	sum := 0.0
	fmt.Println("I  am generating list with date")
	m := map[string]float64{}
	for i, _ := range sheetList {
		cols, _ := f.GetCols(i)
		start, end := GenerateIndex(i, from, to, cols)
		// i want starting and end index
		fmt.Println(start, end)
		for _, col := range cols[6:7] {

			for _, colCell := range col[start:end] {
				a, _ := strconv.Atoi(colCell)
				sum = sum + float64(a)

			}
		}
		//cd = append(cd, CarDetails{CarName: i, CarAmount: sum})
		m[i] = sum
		fmt.Println("car name:", i, "from:", from, "to:", to, "Total Amount:", sum)
		sum = 0.0

	}
	fmt.Println("done")
	//fmt.Println(cd)
	//	fmt.Println("Sheet Count", sheetCount, "total kharcha", total)
	return m
}
