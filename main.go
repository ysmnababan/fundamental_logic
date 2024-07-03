package main

import (
	"log"
	"strconv"
	"sync"
)

func Merge(c ...chan string) <-chan string {
	// TODO: implement function
	var wg sync.WaitGroup
	out := make(chan string, 1000)
	for _, val := range c {
		wg.Add(1)
		// log.Println(c[i])
		go func(sc chan string) {
			defer wg.Done()
			log.Println("go routine")
			for msg := range sc {
				out <- msg
				// log.Println(msg)
			}

		}(val)
		// for msg := range c[i] {
		// 	out <- msg
		// }
	}
	wg.Wait()
	out <- "oke"
	close(out)
	return out
}

func consume(c <-chan string) []string {
	var msgs []string
	for m := range c {
		msgs = append(msgs, m)
	}
	return msgs
}
func main() {
	// a := []string{"NORTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}
	a := []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}
	// a := []string{"SOUTH", "NORTH", "WEST"}
	// log.Println(Opposite("SOUTH","EAST"))
	// log.Println(ParseDigit(515))
	// // a := []string{"SOUTH", "NORTH", "WEST"}
	// a := ParseDigit(530)
	// // b := []string{"SOUTH", "WEST"}
	// b := ParseDigit(531)
	// log.Println(unorderedEqual(a, b))
	// log.Println(NextSmaller(1234567908))
	
	log.Println(DirReduc(a))
}

func NextSmaller(n int) int {
	// Your code here
	newN := ParseDigit(n)

	if len(newN) < 2 {
		return -1
	}
	for i := n - 1; i > 9; i-- {
		if unorderedEqual(newN, ParseDigit(i)) {
			log.Println(i)
			return i
		}
	}
	return -1
}

func unorderedEqual(first, second []string) bool {
	log.Println(first, second)
	if len(first) != len(second) {
		return false
	}
	exists1 := make(map[string]int)
	for _, value := range first {
		if _, ok := exists1[value]; !ok {
			exists1[value] = 1
		} else {
			exists1[value]++
		}
	}

	exists2 := make(map[string]int)
	for _, value := range second {
		if _, ok := exists2[value]; !ok {
			exists2[value] = 1
		} else {
			exists2[value]++
		}
	}
	log.Println(exists1, exists2)
	if len(exists1) != len(exists2) {
		return false
	}
	for key, _ := range exists1 {
		if exists1[key] != exists2[key] {
			return false
		}
	}
	return true
}

func ParseDigit(n int) []string {
	res := []string{}
	numStr := strconv.Itoa(n)
	for _, digit := range numStr {
		// val, _ := digit
		res = append(res, string(digit))
	}
	return res
}

func DirReduc(arr []string) []string {
	// your code

	stage1 := optimaze(arr)
	log.Println(stage1)
	log.Println("optimeze", optimaze(stage1))
	if len(stage1) < 2 {
		return stage1
	}

	change := true
	for change {
		if len(optimaze(stage1)) == len(stage1) {
			// log.Println()
			change = false
		} else {
			stage1 = optimaze(stage1)
		}
	}

	return stage1
}

func optimaze(arr []string) []string {
	if len(arr) < 2 {
		return arr
	}
	temp := []string{}
	for i := 0; i < len(arr)-1; i++ {
		log.Println(i, temp)
		if Opposite(arr[i], arr[i+1]) {
			log.Println("opposite")
			i += 1
		} else {
			// log.Println("here")
			temp = append(temp, arr[i])
		}
		if i == len(arr)-2 {
			temp = append(temp, arr[i+1])
		}
	}
	return temp
}

func Opposite(a, b string) bool {
	var opp bool = false
	// log.Println(a, b)
	switch a {
	case "NORTH":
		// log.Println("here")
		if b == "SOUTH" {
			opp = true
		}
	case "SOUTH":
		if b == "NORTH" {
			opp = true
		}
	case "EAST":
		if b == "WEST" {
			opp = true
		}
	case "WEST":
		if b == "EAST" {
			opp = true
		}
	}
	// log.Println(opp)
	return opp
}

func SumDigPow(a, b uint64) []uint64 {
	res := []uint64{}
	for i := a; i <= b; i++ {
		if IsDigiPow(i) {
			log.Println("here")
			res = append(res, i)
		}
	}
	return res
}

func IsDigiPow(a uint64) bool {
	digit := []uint64{}
	tempA := a
	for a != 0 {
		digit = append(digit, a%10)
		a = a / 10
	}
	// log.Println(digit)
	var sum uint64 = 0
	for i := 0; i < len(digit); i++ {
		sum += pow(digit[i], len(digit)-i)
	}
	log.Println(sum)
	if sum == tempA {
		return true
	}
	return false
}

func pow(a uint64, b int) uint64 {
	var sum uint64
	sum = 1
	for i := 1; i <= int(b); i++ {
		sum = sum * a
	}
	return sum
}
