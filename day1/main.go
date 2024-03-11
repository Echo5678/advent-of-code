package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
    file, err := os.Open("./input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

	var total = 0;
    
    for scanner.Scan() {
		var mySlice []int = make([]int, 0)
		for t, v := range scanner.Text() {
			if(47 < v && v < 58) {
				fmt.Print(t, "\n")
				mySlice = append(mySlice, int(v));
			}
		}

		if(len(mySlice) == 1) {
			temp1 := int(rune(mySlice[0])) - 48
			temp := strconv.Itoa(temp1) + strconv.Itoa(temp1);

			i, err := strconv.Atoi(temp)

			if(err != nil) {
				panic(err)
			}

			total += i
		} else {
			temp1 := int(rune(mySlice[0]) - 48);
			temp2 := int (rune(mySlice[len(mySlice) - 1])) - 48;
			
			temp := strconv.Itoa(temp1) + strconv.Itoa(temp2);

			i, err := strconv.Atoi(temp);

			if(err != nil) {
				panic(err);
			}

			total += i
		}
		
    }

	fmt.Println("Total: ", total);

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	
}