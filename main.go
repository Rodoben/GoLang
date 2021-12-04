package main

import (
	"AAdarsh/ExcelPack"
	"fmt"
	"html/template"
	"net/http"
)

type Date struct {
	day  int
	mon  int
	year int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/listdate", listdate)
	http.ListenAndServe(":8020", nil)
}

func index(res http.ResponseWriter, req *http.Request) {

	sl := ExcelPack.GetSheetList()
	list := ExcelPack.GenerateTotalAmount(sl)
	tpl.ExecuteTemplate(res, "index3.gohtml", list)
}

func listdate(res http.ResponseWriter, req *http.Request) {
	from := req.FormValue("From")
	to := req.FormValue("To")
	fmt.Println(from, to)

	s := ExcelPack.GetSheetList()

	list := ExcelPack.GenerateTotalAmountWithDate(s, from, to)
	tpl.ExecuteTemplate(res, "index3.gohtml", list)
}
