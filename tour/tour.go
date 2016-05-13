package main

import (
	"fmt"
	"image"
	"io"
	"math"
	"math/rand"
	"runtime"
	"strings"
	"sync"
	"time"
)

const Pi = 3.14
const (
	Big   = 1 << 100
	Small = Big >> 99
)

var myMap map[string]Vertex2

func main() {
	const World = "삼일"

	for i := 0; i < 3; i++ {
		defer fmt.Println("defer for:", i)
	}

	t := time.Now()
	fmt.Printf("The time - time.Now() is %%s %s\n", t)
	fmt.Printf("The time - time.Now() is %%d %d\n", t)
	fmt.Printf("The time - time.Now() is %%v %v\n", t)
	fmt.Printf("100 is %%T %T\n", 100)
	fmt.Printf("The time - time.Now() is %%T %T\n", t)
	fmt.Printf("The time - time.Now() is %%t %t\n", t)

	fmt.Println("")

	fmt.Println("rand.Intn(100) without seed giving the same number:", rand.Intn(100))

	r := rand.New(rand.NewSource(99))
	fmt.Println("r.Intn(100) with fixed seed 99 giving the same number: ", r.Intn(100))

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println("r.Intn(100) with non-fixed seed time.Now().UnixNano() giving each different numbers: ", r.Intn(100))

	fmt.Println("add(231, 231) = ", add(231, 231))
	a, b := "2310", "이삼일"
	j, k := swap(a, b)
	fmt.Printf("a = %s, b = %s : swap(a, b) returns c = %s, d = %s\n", a, b, j, k)

	c, python, java := 1, false, "no!"
	fmt.Printf("c, python, java := 1, false, \"no!\" returns %v, %v, %v\n", c, python, java)

	n, m := split(231)
	fmt.Printf("n = %v, m = %v : split(231)\n", n, m)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println("Type conversion : float64(x*x + y*y)", x, y, z)

	fmt.Println("Constants", World, Pi)

	fmt.Println("")

	sum := 0
	for i := 0; i < 231; i++ {
		sum += i
	}
	fmt.Println("for i := 0; i < 231; i++ { sum += i } =", sum)

	sum2 := 1
	for sum2 < 231 {
		sum2 += sum2
	}
	// for sum2 < 231 { sum2 += sum2 }
	fmt.Println("for ; sum2 < 231; { sum2 += sum2 } =", sum2)

	fmt.Println("sqrt(-231) =", sqrt(-231))

	// switch example
	fmt.Printf("switch os := runtime.GOOS; os {...}: ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println(os, ":", "OS X")
	case "linux":
		fmt.Println(os, ":", "Linux")
	default:
		fmt.Println(os)
	}

	fmt.Printf("switch t := time.Now() {...} : ")
	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("Good morning!", t)
	case t.Hour() < 17:
		fmt.Println("Good afternoon.", t)
	default:
		fmt.Println("Good evening.", t)
	}

	// struct and pointer example
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println("struct Vertex{1e9, 2}", v)

	// slice example
	primes := [6]int{2, 3, 5, 7, 11, 13}

	//var slc []int = primes[1:3]
	slc := primes[1:]
	fmt.Println("slice : primes[1:] :", slc)

	slc2 := []struct {
		i int
		b bool
	}{
		{2, true}, {3, false}, {5, true}, {7, true}, {11, false}, {13, true},
	}
	fmt.Println("slice : []struct { ... } {...}:", slc2)

	// map example
	myMap = make(map[string]Vertex2)
	myMap["Allif"] = Vertex2{40.68433, -74.39967}

	fmt.Println("myMap[\"Allif\"]", myMap["Allif"])

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts { // map도 range 사용 가능
		fmt.Printf("%v: %v\n", name, ip)
	}

	fmt.Println("")

	// interface example
	var ia Abser
	rf := MyFloat(-math.Sqrt2)
	rv := Vertex2{3, 4}

	ia = rf  // a MyFloat implements Abser
	ia = &rv // a *Vertex implements Abser
	//ia = rv  // a Vertex does not implement Abser

	fmt.Println("Interface example : ", ia.Abs())

	// io Reader example
	rr := strings.NewReader("Hello, Reader!")
	rb := make([]byte, 8)
	for {
		n, err := rr.Read(rb)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, rb)
		fmt.Printf("b[:n] = %q\n", rb[:n])
		if err == io.EOF {
			break
		}
	}

	// image example
	im := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println("image.NewRGBA(image.Rect(0, 0, 100, 100)):", im.Bounds(), im.At(0, 0))

	// channel and select example
	ch := make(chan int)
	quitCh := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("go routine with chanel and select: ", i, <-ch)
		}
		quitCh <- 0
	}()
	fibonacci(ch, quitCh)

	// mutex example
	mutexc := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go mutexc.Inc("somekey")
	}

	time.Sleep(time.Second)

	fmt.Println("sync.Mutex", mutexc.Value("somekey"))

	// defer example
	fmt.Println("Counting done in defer")

}

func add(x int, y int) int              { return x + y }
func swap(x, y string) (string, string) { return y, x }
func split(sum int) (x, y int)          { x = sum * 4 / 9; y = sum - x; return }
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func printSlice(s []int) { fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) }

func typeSwitchDo(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y int
}

type Vertex2 struct {
	X, Y float64
}

func (v *Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}
func Sqrt(x float64) (float64, error) {
	return 0, nil
}

// go concurrency select
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
			//default: // nonblocking
			//	fmt.Println("...")
		}
	}
}

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}
