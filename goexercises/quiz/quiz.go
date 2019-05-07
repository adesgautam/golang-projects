package main

import (
		"fmt"
		"io/ioutil"
		"strings"
		//"bufio"
		"os"
		"flag"
		"time"
		)


func main(){

	csvFile := flag.String("csv", "problems.csv", "a csv file in the format question,answer")
	t := flag.Int("time", 30, "Maximum time for the quiz")
	flag.Parse()

	data, err := ioutil.ReadFile(*csvFile)
	if err != nil {
		fmt.Println("Failed to read file 'problems.csv'")
		os.Exit(1)
	}

	d := strings.Split(string(data), "\n")
	//fmt.Println(d)

	totscore := 0
	text := ""
	before := time.Now()

	for i, q := range d {
		now := time.Now()
		diff := now.Sub(before).Seconds()
		if diff > float64(*t) {
			break
		} else {
			qu := strings.Split(q, ",")
			ques := strings.TrimSpace(qu[0])
			ans := strings.TrimSpace(qu[1])

			fmt.Println("Q",i+1,":", ques)
			fmt.Scanf("%s\n", &text)

	        if text == ans {
	        	totscore++
	        	fmt.Println("Correct !")
	        } else {
	        	fmt.Println("Incorrect !")
	        }
	    }
	}

	fmt.Println("Total questions:", len(d))
	fmt.Println("Number question you answered correcty:", totscore)
}