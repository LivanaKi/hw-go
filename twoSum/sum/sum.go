package sum

import (
	"fmt"
	"strings"

	"github.com/gammazero/deque"
)

func Graph(){
	var graph = map[string][]string{
		"you":    []string{"alice", "bob", "claire"},
		"bob":    []string{"anuj", "peggy"},
		"alice":  []string{"you", "peggy", "alice"},
		"claire": []string{"jonny", "thom", "bob"},
		"anuj":   []string{},
		"peggy":  []string{"you"},
		"thom":   []string{},
		"jonny":  []string{"peggy", "peggy"},
		"jim":    []string{"thom"},
	}
	
	deque.New[string]()
}


func IntegerRom (num int) string{
	var res []string
	var b strings.Builder
	
	roman := [4][]string{
		[]string{ "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		[]string{ "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		[]string{ "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		[]string{ "M", "MM", "MMM"},
	}

	if num/1000 > 0{
		n := num/1000
		num %= 1000
		res = append(res, roman[3][n-1])
	}
	if num/100 > 0{
		n := num/100
		num %= 100
		res = append(res, roman[2][n-1])
	}
	if num/10 > 0{
		n := num/10
		num %= 10
		res = append(res, roman[1][n-1])
	}
	if num%10 > 0{
		n := num%10
		res = append(res, roman[0][n-1])
	}
	for i := 0; i < len(res); i++{
		b.WriteString(res[i])
	}

	return b.String()
}


func Findprimes(number int) bool {
	for i := 2; i < number; i++ {
        if number%i == 0 {
			return false
        }
    }
	if number > 1 {
		return true
	} else {
	    return false
	}
}

func Highlow(high int, low int) {
    if high < low {
        fmt.Println("Panic!")
        panic("highlow() low greater than high")
    }
    defer fmt.Println("Deferred: highlow(", high, ",", low, ")")
    fmt.Println("Call: highlow(", high, ",", low, ")")

    Highlow(high, low + 1)
}


func UpdateName(name *string){
	*name = "David"
}


func MaxArea(height []int) int {
	i := 0
	j := len(height) - 1
	var res int
	var answ int

	for j >= i {

		if height[i] <= height[j] {
			res = height[i] * (j - i)
			i++
		}else{
			res = height[j]*(j - i)
			j--
		}

		if res > answ{
			answ = res
		}
	}
	return answ
}