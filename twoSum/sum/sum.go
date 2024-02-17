package sum

import (
	"fmt"
	"strings"
)

func Graph(graph map[string][]string) bool {

	var queue []string
	queue = append(queue, graph["you"]...)

	var searched []string

	for len(queue) != 0 {
		person := queue[0]
			if !searchedPerson(person, searched){
				if personIsSeller(person) {
					fmt.Printf("%s is a mango seller!\n", person)
					return true
				} else {
					queue = queue[1:]
					queue = append(queue, graph[person]...)
				}
			}else{
				searched = append(searched, person)
			}
			fmt.Println(searched)
		}
	return false
}

func searchedPerson(name string, sear []string) bool{
	for i := 0; i < len(sear); i++ {
		if name == sear[i]{
			return false
		}		
	}
	return true
}

func personIsSeller(name string) bool {
	return name[len(name)-1] == 'm'
}

func IntegerRom(num int) string {
	var res []string
	var b strings.Builder

	roman := [4][]string{
		[]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		[]string{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		[]string{"C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		[]string{"M", "MM", "MMM"},
	}

	if num/1000 > 0 {
		n := num / 1000
		num %= 1000
		res = append(res, roman[3][n-1])
	}
	if num/100 > 0 {
		n := num / 100
		num %= 100
		res = append(res, roman[2][n-1])
	}
	if num/10 > 0 {
		n := num / 10
		num %= 10
		res = append(res, roman[1][n-1])
	}
	if num%10 > 0 {
		n := num % 10
		res = append(res, roman[0][n-1])
	}
	for i := 0; i < len(res); i++ {
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

	Highlow(high, low+1)
}

func UpdateName(name *string) {
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
		} else {
			res = height[j] * (j - i)
			j--
		}

		if res > answ {
			answ = res
		}
	}
	return answ
}
