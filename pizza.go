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

	}
