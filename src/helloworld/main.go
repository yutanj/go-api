// package main

// import "fmt"

// func main() {
// 	for i := 0; i < 100; i++ {
// 		switch {
// 		case i%3 == 0 && i%5 == 0:
// 			fmt.Println("fizzbuzz")
// 		case i%3 == 0:
// 			fmt.Println("fizz")
// 		case i%5 == 0:
// 			fmt.Println("buzz")
// 		default:
// 			fmt.Println(i)
// 		}
// 	}
// }

// package main

// import "fmt"

// func findprimes(number int) bool {
// 	for i := 2; i < number; i++ {
//         if number % i == 0 {
// 			return false
//         }
//     }

// 	if number != 1 {
// 		return true
// 	} else {
// 	    return false
// 	}
// }

// func main() {
//     fmt.Println("Prime numbers less than 20:")

//     for number := 1; number < 20; number++ {
//         if findprimes(number) {
//             fmt.Printf("%v ", number)
//         }
//     }
// }

// package main

// import "fmt"

// func fibonacci(n int) []int {
// 	if n < 2 {
// 		return make([]int, 0)
// 	}

// 	nums := make([]int, n)
// 	nums[0], nums[1] = 1, 1

// 	for i := 2; i < n; i++ {
// 		nums[i] = nums[i-1] + nums[i-2]
// 	}

// 	return nums
// }

// func main() {
// 	var num int

// 	fmt.Print("What's the Fibonacci sequence you want? ")
// 	fmt.Scanln(&num)
// 	fmt.Println("The Fibonacci sequence is:", fibonacci(num))
// }

// package main

// import (
// 	"fmt"
// )

// func romanToArabic(numeral string) int {
// 	romanMap := map[rune]int{
// 		'M': 1000,
// 		'D': 500,
// 		'C': 100,
// 		'L': 50,
// 		'X': 10,
// 		'V': 5,
// 		'I': 1,
// 	}

// 	arabicVals := make([]int, len(numeral)+1)

// 	for index, digit := range numeral {
// 		if val, present := romanMap[digit]; present {
//             println(val)
//             println(present)
// 			arabicVals[index] = val
// 		} else {
// 			fmt.Printf("Error: The roman numeral %s has a bad digit: %c\n", numeral, digit)
// 			return 0
// 		}
// 	}
//     println(arabicVals)

// 	total := 0

// 	for index := 0; index < len(numeral); index++ {
// 		if arabicVals[index] < arabicVals[index+1] {
// 			arabicVals[index] = -arabicVals[index]
// 		}
// 		total += arabicVals[index]
// 	}

// 	return total
// }

// func main() {
// 	fmt.Println("MCLX is", romanToArabic("MCLX"))
// 	//fmt.Println("MCMXCIX is ", romanToArabic("MCMXCIX"))
// 	//fmt.Println("MCMZ is", romanToArabic("MCMZ"))
// }

// package main

// import (
//     "log"
// )

// func main() {
//     log.Print("Hey, I'm a log!")
// }

// package main

// import "fmt"

// type triangle struct {
//     size int
// }

// type square struct {
//     size int
// }

// func (t triangle) perimeter() int {
//     return t.size * 3
// }

// func (s square) perimeter() int {
//     return s.size * 4
// }

// func main() {
//     t := triangle{3}
//     s := square{4}
//     fmt.Println("Perimeter (triangle):", t.perimeter())
//     fmt.Println("Perimeter (square):", s.perimeter())
// }

// package main

// import (
//     //"bytes"
//     "encoding/json"
//     "fmt"
// )

// var xjapanJson string = `
// {
//     "members": [
//         { "name": "Toshl",   "instrument": "vocal"  },
//         { "name": "PATA",    "instrument": "guitar" },
//         { "name": "HEATH",   "instrument": "bass"   },
//         { "name": "SUGIZO",  "instrument": "guitar" },
//         { "name": "YOSHIKI", "instrument": "drums"  }
//     ],
//     "songs": ["紅", "Silent Jealousy"]
// }
// `

// type Musician struct {
//     Name string `json:"name"`
//     Instrument string `json:"instrument"`
// }

// type XJapan struct {
//     Members []Musician `json:"members"`
//     Songs  []string `json:"songs"`
// }

// func main() {
//     // Unmarshal結果の格納先である構造体のポインターを取得
//     xJapan := new(XJapan)

//     // JSON文字列をバイト列にキャスト
//     jsonBytes := []byte(xjapanJson)
//     println(jsonBytes)

//     // xJapanにバイト列を格納する
//     if err := json.Unmarshal(jsonBytes, xJapan); err != nil {
//         fmt.Println(err)
//         return
//     }
//     for _, members := range xJapan.Members {
//         fmt.Printf("NAME: %-7s INSTRUMENT: %s\n", members.Name, members.Instrument)
//     }
// /*
//     出力結果
//     NAME: Toshl   INSTRUMENT: vocal
//     NAME: PATA    INSTRUMENT: guitar
//     NAME: HEATH   INSTRUMENT: bass
//     NAME: SUGIZO  INSTRUMENT: guitar
//     NAME: YOSHIKI INSTRUMENT: drums
// */

//     for _, song := range xJapan.Songs {
//         fmt.Println(song)
//     }
// /*
//     出力結果
//     紅
//     Silent Jealousy
// */
// }

package main

import (
    "fmt"
    "log"
    "net/http"
)

type dollars float32

func (d dollars) String() string {
    return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    for item, price := range db {
        fmt.Fprintf(w, "%s: %s\n", item, price)
    }
}

func main() {
    db := database{"Go T-Shirt": 25, "Go Jacket": 55}
    log.Fatal(http.ListenAndServe("localhost:8000", db))
}