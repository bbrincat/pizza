package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
)

func main(){

	file , err := os.Open("inputs/example.in")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rows, cols , minToppings , maxArea int

	if scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		rows , err = strconv.Atoi(fields[0])
		cols , err = strconv.Atoi(fields[1])
		minToppings , err = strconv.Atoi(fields[2])
		maxArea , err = strconv.Atoi(fields[3])

		fmt.Println()
	}

	pizza := make([][]string, rows)

	var i = 0
	for scanner.Scan() {
		col := strings.Split(scanner.Text(), "")
		pizza[i] = col
		i = i +1
	}

	fmt.Println(rows,cols,minToppings, maxArea)
	for _, r := range pizza {
		for _, c := range r {
			fmt.Printf("%s ",c)
		}
		fmt.Printf("\n")
	}

	evalutionFn := getEvaluationFn(pizza, minToppings, maxArea)


	s := pizzaSlice{coord{0,0}, coord{2,1}}
	fmt.Println(evalutionFn(s))

}
	func getEvaluationFn(pizza [][]string, L,H int) (func (slice pizzaSlice) (ok bool, score int)){
		return  func (slice pizzaSlice) (ok bool, score int){
			mushroom, tomato, total := evalSlice(pizza, slice)
			if mushroom >= L && tomato >= L && total <=H {
				return true, total
			}
			return false, total
		}
	}


	func evalSlice( pizza [][]string, s pizzaSlice) (mushroom, tomato, total int ){
		for i := s.start.row; i <= s.end.row ; i ++ {
			for j := s.start.col; j <= s.end.col; j++{

				if pizza[i][j] == "M"{
					mushroom++
					}

				if pizza[i][j] == "T"{
					tomato++
				}
				total ++
			}

			}
			return
	}

	type coord struct {
		row, col int
	}

	type pizzaSlice struct {
		start, end coord
	}
