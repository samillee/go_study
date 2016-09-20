// -----------------------------------------------
// https://github.com/GoesToEleven/GolangTraining
// 소스에서 참고할 만한 것 추출
package main

import (
	"bufio"
	"crypto/md5"
	"encoding/csv" // csvReader := csv.NewReader(f), record, err := csvReader.Read()
	"fmt"
	"hash/fnv"  // super fast hash
	"io"        // io.WriteString(dst, str)
	"io/ioutil" // bs, err := ioutil.ReadAll(f), err := ioutil.WriteFile("hey.txt", []byte(myStr), 0777)
	"log"
	"net/http"
	"os"            // os.Create(), os.Open()
	"path/filepath" // filepath.Walk - recursive
	"strconv"
	"strings"
)

// 06_constants\
const (
	a = iota // 0
	b        // 1
	c        // 2
)

type ByteSize float64

// 엄청 좋은 예제
const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println("https://github.com/GoesToEleven/GolangTraining")

	n := 42
	fmt.Println(n)

	fmt.Printf("%v \t %T \t %d \t %b \t %X \t %#X \t %q \n", n, n, n, n, n, n, n)
	fmt.Println("n's memory address - ", &n)

	/*
		for i := n; i < 122; i++ {
			fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
		}
	*/

	// 05_blank-identifier\02_http-get_example\01_with-error-checking\
	fmt.Println()
	//res, err := http.Get("http://www.mcleods.com/")
	res, err := http.Get("http://www.showit.co.kr")
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(res.Header)
	//fmt.Printf("%v\n", res.Header)
	fmt.Println()
	for k, v := range res.Header {
		fmt.Printf("%s \t %v \n", k, v)
	}

	fmt.Println()

	page, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close() // Body를 읽으면 반드시 Close 해줘야 함
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%v", page)
	//fmt.Printf("%v", len(page))
	fmt.Println("page len(page):", len(page))

	// 10_for-loop\
	fmt.Println()
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d, \t", i)
		if i >= 10 {
			break
		}
	}

	/*
		// 10_for-loop\07_rune-loop_UTF8\
		fmt.Println()
		for i := 250; i <= 340; i++ {
			//fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
			fmt.Printf("%v - %v - %v", i, string(i), []byte(string(i)))
			if (i+1)%5 == 0 {
				fmt.Println()
			}
		}

		fmt.Println()

		for i := 50; i <= 140; i++ {
			fmt.Printf("%v - %v - %v", i, string(i), []byte(string(i)))
			if (i+1)%5 == 0 {
				fmt.Println()
			}
		}
	*/

	// 11_switch-statements
	fmt.Println()
	fmt.Println()
	fmt.Println("11_switch-statements")
	switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println("Wassup Tim, or, err, Jenny")
	case "Marcus", "Medhi":
		fmt.Println("Both of your names start with M")
	case "Julian", "Sushant":
		fmt.Println("Wassup Julian / Sushant")
	}

	myFriendsName := "Mike"
	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Wassup my friend with name of length 2")
	case myFriendsName == "Marcus", myFriendsName == "Medhi":
		fmt.Println("Your name is either Marcus or Medhi")
	case myFriendsName == "Julian":
		fmt.Println("Wassup Julian")
	default:
		fmt.Println("nothing matched; this is the default")
	}

	SwitchOnType(7)
	SwitchOnType("Samil Lee")

	// 12_if_else-if_else
	fmt.Println()
	fmt.Println("if i%3 == 0")
	for i := 0; i <= 10; i++ {
		if i%3 == 0 {
			fmt.Printf("%d \t", i)
		}
	}

	// 27_package-os ... 31_package-ioutil
	cp_simple("src.txt", "dst.txt")
	writeString(os.Args[1], "Hello Mike")
	cp("src.txt", "dst.txt")

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(info)
		return nil
	})

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path, info.Name(), info.Size(), info.Mode(), info.IsDir())
		return nil
	})

}

func SwitchOnType(x interface{}) {
	switch x.(type) { // this is an assert; asserting, "x is of this type"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	//case contact:
	//	fmt.Println("contact")
	default:
		fmt.Println("unknown")

	}
}

func writeString(srcName, text string) error {
	//dst, err := os.Create(os.Args[1])
	dst, err := os.Create(srcName)
	if err != nil {
		log.Fatalf("error creating destination file:%v ", err)
		return err
	}
	defer dst.Close()

	// 한줄로도 가능
	//dst.WriteString(text)

	/*
		_, err = io.WriteString(dst, str)
		//	bs := []byte(text)
		//	_, err = dst.Write(bs)
		if err != nil {
			log.Fatalln("error writing to file: ", err.Error())
		}
	*/

	rdr := strings.NewReader(text)
	io.Copy(dst, rdr)

	return nil
}

func readAll(src string) string {
	//f, err := os.Open(os.Args[1])
	f, err := os.Open(src)
	if err != nil {
		log.Fatalln("my program broke", err.Error())
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("my program broke again")
	}

	str := string(bs)

	return str
}

func cp_simple(srcName, dstName string) error {
	src, err := os.Open(srcName)
	if err != nil {
		panic(err)
		return fmt.Errorf("%v", err)
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		//panic(err)
		return fmt.Errorf("%v", err)
	}
	defer dst.Close()

	bs := make([]byte, 5)
	src.Read(bs)
	dst.Write(bs)

	return nil
}

// buffered
func cp(srcName, dstName string) error {
	src, err := os.Open(srcName)
	if err != nil {
		return fmt.Errorf("error opening source file: %v", err)
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return fmt.Errorf("error creating destination file:%v ", err)
	}
	defer dst.Close()

	br := bufio.NewReader(src)

	_, err = io.Copy(dst, br)
	if err != nil {
		return fmt.Errorf("error writing to destination file: %v ", err)
	}

	return nil
}

func scan_ex(fn string) {
	src, err := os.Open(fn)
	if err != nil {
		log.Printf("error opening source file: %v", err)
	}
	defer src.Close()

	scanner := bufio.NewScanner(src)
	//scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		line := scanner.Text()
		//word := scanner.Text()
		if len(line) > 0 {
			fmt.Println(">>>", strings.ToUpper(line[0:1])+line[1:], "\n")
			//fmt.Print(strings.ToUpper(line[0:1])+line[1:], " ")
		}
	}
}

type state struct {
	id               int
	name             string
	abbreviation     string
	censusRegionName string
}

func parseState(columns map[string]int, record []string) (*state, error) {
	id, err := strconv.Atoi(record[columns["id"]])
	name := record[columns["name"]]
	abbreviation := record[columns["abbreviation"]]
	censusRegionName := record[columns["census_region_name"]]
	if err != nil {
		return nil, err
	}
	return &state{
		id:               id,
		name:             name,
		abbreviation:     abbreviation,
		censusRegionName: censusRegionName,
	}, nil
}

func csvReadWriteToHtml() {
	// #1 open a file
	f, err := os.Open("../state_table.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// #2 parse a csv file
	csvReader := csv.NewReader(f)
	columns := make(map[string]int)

	stateLookup := map[string]*state{}

	for rowCount := 0; ; rowCount++ {
		record, err := csvReader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}

		if rowCount == 0 {
			for idx, column := range record {
				columns[column] = idx
			}
		} else {
			// #3 do stuff for each row
			state, err := parseState(columns, record)
			if err != nil {
				log.Fatalln(err)
			}
			// #4 add each row to stateLookup map
			stateLookup[state.abbreviation] = state
		}
	}

	// #5 lookup the state
	if len(os.Args) < 2 {
		log.Fatalln("expected state abbreviation")
	}
	abbreviation := os.Args[1]
	state, ok := stateLookup[abbreviation]
	if !ok {
		log.Fatalln("invalid state abbreviation")
	}

	fmt.Println(`
<html>
    <head></head>
    <body>
      <table>
        <tr>
          <th>Abbreviation</th>
          <th>Name</th>
        </tr>`)

	fmt.Println(`
        <tr>
          <td>` + state.abbreviation + `</td>
          <td>` + state.name + `</td>
        </tr>
    `)

	fmt.Println(`
      </table>
    </body>
</html>
    `)
}

/*
at terminal:
go install

at terminal:
programName <state abbreviation> > index.html
*/

func md5file(fileName string) []byte {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	h := md5.New()
	io.Copy(h, f)
	//fmt.Printf("The hash (sum) is: %x\n", h.Sum(nil))
	return h.Sum(nil)
}

func fnvfile(fileName string) []byte {
	//f, err := os.Open(os.Args[1])
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	h := fnv.New64()
	io.Copy(h, f)

	//fmt.Println(h.Sum64())
	//fmt.Printf("The hash (sum) is: %x\n", h.Sum64())
	retrun h.Sum64()
}

func timediff(from, to string) {
	//from, to := os.Args[1], os.Args[2]

	fromTime, _ := time.Parse("2006-01-01_this-does-not-compile", from)
	toTime, _ := time.Parse("2006-01-01_this-does-not-compile", to)

	dur := toTime.Sub(fromTime)
	fmt.Println("elapsed days:", int(dur/(time.Hour*24)))

}
