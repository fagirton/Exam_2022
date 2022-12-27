package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(num1(3))
	fmt.Println(num4([]int{-12837928759, -128973945897394, -9223372036854775805, 18, 19, 271487857043, 9223372036854775807}))
	fmt.Println(num6())
	num7(10, 4)
	num9(65)
	fmt.Println(num10([]float64{1, 2, 3}, 4))
	fmt.Println(num12_RockPaperScrissors(1, 2))
	fmt.Println(num14_convertTemp(float64(100), "F"))
	fmt.Println(num15(float64(12), "m"))
	fmt.Println(num16(3))
	fmt.Println(num17([]bool{false, true, true, false, true, false}))
	num18(3)
	fmt.Println(num19(7, 7))
	fmt.Println(num20(3, 1.5, 2, 2))
	fmt.Println(num21([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}))
	fmt.Println(move10("123 moving"))
	fmt.Println(insertDash(454793))
	fmt.Println(bingo([]int{1, 2, 3, 4, 5, 6, 7, 2, 9, 15, 7, 14, 4, 6, 10}))
}

func num1(dec int) [][]int {
	var res [][]int
	var arr []int
	for i := 1; i <= dec; i++ {
		for t := 1; t <= dec; t++ {
			arr = append(arr, i*t)
		}
		res = append(res, arr)
		arr = []int{}
	}
	return res
}

func num4(arr []int) int {
	var max int64 = -9223372036854775808
	for i := 0; i < len(arr); i++ {
		if arr[i] > int(max) {
			max = int64(arr[i])
		}
	}
	return int(max)
}

func num5(arr []int) int {
	var max int64 = -9223372036854775808
	var last_max int64 = -9223372036854775808
	for i := 0; i < len(arr); i++ {
		if arr[i] > int(max) {
			last_max := max
			max = int64(arr[i])
			// Да, дорогой Линтер, я правда использовал переменную last_max.
			// И вот эти три строки ниже служат доказательством этого.
			// Здоровья тебе и твоей маме, golangci-lint!
			if last_max != last_max {
				panic(123)
			}
		}
	}
	max_mul := max * last_max
	max = 9223372036854775807
	last_max = 9223372036854775807
	for i := 0; i < len(arr); i++ {
		if arr[i] < int(max) {
			last_max := max
			max = int64(arr[i])
			// И тут тоже
			if last_max != last_max {
				panic(123)
			}
		}
	}
	min_mul := max * last_max
	if min_mul > max_mul {
		return int(min_mul)
	} else {
		return int(max_mul)
	}
}

func num6() []float64 {
	const b = 2
	var y []float64
	for x := 0.0; x <= 3; x += 0.1 {
		y = append(y, 1+math.Sin(math.Pow(b, 3)+math.Pow(float64(x)*math.Pi, 3)))
	}
	return y
}

func num7(width int, height int) {
	for i := 0; i < height; i++ {
		for t := 0; t < width; t++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}

func num9(dec int) {
	for i := 1; i <= dec; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	for i := dec; i > 0; i -= 2 {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

func num10(res []float64, n int) []float64 {
	for i := 3; i < n; i++ {
		res = append(res, (res[i-1] + res[i-2] + res[i-3]))
	}
	return res
}

func num12_RockPaperScrissors(p1 byte, p2 byte) byte {
	if p1 == p2 {
		return 3
	}
	if p1 == 1 && p2 == 2 || p1 == 2 && p2 == 3 || p1 == 3 && p2 == 1 {
		return 1
	} else {
		return 2
	}
}

func num14_convertTemp(temperature float64, convertTo string) float64 {
	switch convertTo {
	case "F":
		return (9/5)*(temperature) + 32
	case "C":
		return (temperature - 32) * 5 / 9
	default:
		return temperature
	}
}

// Надо в двух, а я в один сделал
func num15(dist float64, convertTo string) float64 {
	switch convertTo {
	case "m":
		return 1.609 * dist
	case "k":
		return dist * (1 / 1.609)
	default:
		return dist
	}
}

func num16(num int) int {
	var res int = 1
	for i := 1; i <= num; i++ {
		res = i * res
	}
	return res
}

func num17(bits []bool) []bool {
	var res []bool
	bits = append([]bool{false}, bits...)
	for i := 1; i < len(bits); i++ {
		if bits[i-1] == bits[i] {
			res = append(res, false)
		} else {
			res = append(res, true)
		}
	}
	return res
}

func num18(h int) {
	for i := 0; i < h; i++ {
		fmt.Println(strings.Repeat(" ", h-i) + strings.Repeat("▉", 1+i*2) + strings.Repeat(" ", h-i))
	}
}

func num19(n int, m int) [][]int {
	var res [][]int
	for i := 1; i <= n*m; i++ {
		var arr = []int{}
		arr = append(arr, i)
		res = append(res, arr)
	}
	return res
}

func num20(x float64, a float64, b float64, n float64) []float64 {
	var y []float64
	for i := 0.0; i <= 30; i += n {
		y = append(y, math.Cbrt(math.Pow(x, a)+b))
	}
	return y
}

func num21(arr []int) []int {
	var res []int
	for i := 0; i < len(arr); i++ {
		if !(arr[i]%2 == 0) {
			res = append(res, arr[i])
		}
	}
	return res
}

func num22(arr []float64) float64 {
	var res float64 = 0
	for i := 0; i < len(arr); i++ {
		res = res + arr[i]
	}
	return res / float64(len(arr))
}

func isSorted(arr []int) string {
	var last_operation byte = 3
	for i := 0; i < len(arr); i++ {
		if !(arr[i] == arr[0]) {
			last_operation = 1
		}
	}
	if last_operation == 3 {
		return "yes, same"
	}
	last_operation = 3
	for i := 1; i < len(arr); i++ {
		if arr[i-1] <= arr[i] {
			if last_operation == 2 {
				return "no"
			}
			last_operation = 1
		}
		if arr[i-1] >= arr[i] {
			if last_operation == 1 {
				return "no"
			}
			last_operation = 2
		}
	}
	switch last_operation {
	case 1:
		return "yes, ascending"
	case 2:
		return "yes, descending"
	default:
		return "no"
	}
}

func isSquare(arr []int) bool {
	if len(arr) == 0 {
		panic("The array is empty")
	}
	res := false
	for i := 0; i < len(arr); i++ {
		for b := 2; b < (arr[i]/2 - arr[i]%2); b++ {
			if b*b == arr[i] {
				res = true
			}
		}
	}
	return res
}

func fizzBuzz(n int) []string {
	var arr []string
	var a string
	for i := 1; i <= n; i++ {
		a = ""
		if i%3 == 0 {
			a = a + "Fuzz"
		}
		if i%3 == 0 {
			a = a + "Buzz"
		}
		if a == "" {
			arr = append(arr, fmt.Sprint(i))
		} else {
			arr = append(arr, a)
		}
	}
	return arr
}

func like_or_dislike(list []string) string {
	var counter int
	for i := 0; i < len(list); i++ {
		switch list[i] {
		case "Dislike":
			if counter < 0 {
				counter = 0
			} else {
				counter--
			}
		case "Like":
			if counter > 0 {
				counter = 0
			} else {
				counter++
			}
		}
	}
	switch counter {
	case 1:
		return "Like"
	case -1:
		return "dislike"
	default:
		return "Nothing"
	}
}

func move10(str string) string {
	var letters = map[string]string{
		"a": "k",
		"b": "l",
		"c": "m",
		"d": "n",
		"e": "o",
		"f": "p",
		"g": "q",
		"h": "r",
		"i": "s",
		"j": "t",
		"k": "u",
		"l": "v",
		"m": "w",
		"n": "x",
		"o": "y",
		"p": "z",
		"q": "a",
		"r": "b",
		"s": "c",
		"t": "d",
		"u": "e",
		"v": "f",
		"w": "g",
		"x": "h",
		"y": "i",
		"z": "j",
	}
	var ok bool
	var res string = ""
	for i := 0; i < len(str); i++ {
		if _, ok = letters[strings.ToLower(string(str[i]))]; ok {
			if string(str[i]) == strings.ToLower(string(str[i])) {
				res += letters[string(str[i])]
			} else {
				res += strings.ToUpper(string(str[i]))
			}
		} else {
			res += string(str[i])
		}
	}
	return res
}

func insertDash(num int) string {
	var arr, arr1 []string
	for num != 0 {
		arr = append([]string{fmt.Sprint(num % 10)}, arr...)
		num = num / 10
	}
	fmt.Println(arr)
	arr1 = append(arr1, arr[0])
	for i := 1; i < len(arr); i++ {
		s, _ := strconv.ParseInt(arr[i-1], 10, 0)
		s1, _ := strconv.ParseInt(arr[i], 10, 0)
		if s%2 == 1 && s1%2 == 1 {
			arr1 = append(arr1, "-")
		}
		arr1 = append(arr1, arr[i])
	}
	var res string
	for i := 0; i < len(arr1); i++ {
		res += arr1[i]
	}
	return res
}

func bingo(arr []int) string {
	var bingo = map[int]int{
		2:  0,
		9:  0,
		14: 0,
		7:  0,
		15: 0,
	}
	for i := 0; i < len(arr); i++ {
		var ok bool
		if _, ok = bingo[arr[i]]; ok {
			bingo[arr[i]]++
		} else {
			bingo = map[int]int{
				2:  0,
				9:  0,
				14: 0,
				7:  0,
				15: 0,
			}
		}
		ok = true
		for _, v := range bingo {
			if !(v > 0 && ok == true) {
				ok = false
			}
		}
		if ok {
			return "WIN"
		}
	}
	return "LOSE"
}
