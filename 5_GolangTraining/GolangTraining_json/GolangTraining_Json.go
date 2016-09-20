// -----------------------------------------------
// https://github.com/GoesToEleven/GolangTraining
// 소스에서 Json 관련 참고할 만한 것 추출

package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	jsonData := `
	{
	"name": "Todd McLeod",
	"age": 44
	}
	`
	var obj map[string]interface{}

	//type Anything interface{}
	//var obj map[string]Anything

	err := json.Unmarshal([]byte(jsonData), &obj)
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)

	jsonData2 := `
	[100, 200, 300.5, 400.1234]
	`
	var obj2 []float64

	err2 := json.Unmarshal([]byte(jsonData2), &obj2)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println()
	fmt.Println(obj2)

	bs, err := json.Marshal([]int{1, 2, 3, 4})
	fmt.Println()
	fmt.Println("bs, err := json.Marshal([]int{1, 2, 3, 4}) string(bs):", string(bs), err)

	readJsonFile("data.json")

	readJsonFile2("data.json")

	csvFile2Json("table.csv", "output.json")

	toJson_ex1()
}

func readJsonFile(fn string) {
	//f, err := os.Open("data.json")
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var obj map[string]interface{}
	err = json.NewDecoder(f).Decode(&obj)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println(obj)
}

type StockData struct {
	Returns []float64 `json:"returns"`
}

func readJsonFile2(fn string) {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var obj StockData
	err = json.NewDecoder(f).Decode(&obj)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println(obj)
}

type StockData2 struct {
	Date     string
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   float64
	AdjClose float64
}

func toFloat(str string) float64 {
	v, _ := strconv.ParseFloat(str, 64)
	return v
}

func csvFile2Json(fn, fnOut string) {

	//open file
	//src, err := os.Open("table.csv")
	src, err := os.Open(fn)
	if err != nil {
		log.Fatalln("couldn't open file", err.Error())
	}
	defer src.Close()

	// reader for csv file
	rdr := csv.NewReader(src)

	// read csv file
	data, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln("couldn't readall", err.Error())
	}

	// convert to JSON
	listOfStockData := []StockData2{}
	// put data into struct
	for _, row := range data {
		sd := StockData2{
			Date:     row[0],
			Open:     toFloat(row[1]),
			High:     toFloat(row[2]),
			Low:      toFloat(row[3]),
			Close:    toFloat(row[4]),
			Volume:   toFloat(row[5]),
			AdjClose: toFloat(row[6]),
		}
		listOfStockData = append(listOfStockData, sd)
	}

	//convert to JSON
	dst, err := os.Create(fnOut)
	if err != nil {
		log.Fatalln("couldn't open file", err.Error())
	}
	defer dst.Close()

	err = json.NewEncoder(dst).Encode(listOfStockData)

	//b, err := json.Marshal(data)

	if err != nil {
		log.Fatalln("couldn't marshall", err.Error())
	}

	fmt.Println()
	fmt.Println("Check the json output file:", fnOut)
	// show
	// os.Stdout.Write(b)
	// fmt.Println(b)
}

type Article struct {
	Name  string
	draft bool
}

func toJson_ex1() {
	myArticle := Article{
		Name:  "Once And Then Again",
		draft: false,
	}

	data, err := json.Marshal(myArticle)

	if err != nil {
		log.Fatalln("couldn't marshall", err.Error())
	}

	fmt.Println()
	fmt.Println(string(data))
	//os.Stdout.Write(data)
}
