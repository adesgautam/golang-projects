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
	timeLimit := flag.Int("limit", 30, "Maximum time limit for the quiz in seconds")
	flag.Parse()



	data, err := ioutil.ReadFile(*csvFile)
	if err != nil {
		fmt.Println("Failed to read file 'problems.csv'")
		os.Exit(1)
	}

	d := strings.Split(string(data), "\n")

	totscore := 0
	
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	problemloop:
	for i, q := range d {
		qu := strings.Split(q, ",")
		ques := strings.TrimSpace(qu[0])
		ans := strings.TrimSpace(qu[1])
		fmt.Println("Q",i+1,":", ques)

		ansCh := make(chan string)
		go func(){
			var text string
			fmt.Scanf("%s\n", &text)
			ansCh <- text
		}()

		select {
			case <-timer.C:
				fmt.Println()
				break problemloop
			case answer := <-ansCh:
				if answer == ans {
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