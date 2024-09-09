package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	// Define variables
	var records [][]string

	// Define basic flags used by this program
	filename := flag.String("file", "quiz.csv", "The name of the CSV quiz file!.")
	timeout := flag.Int("timeout", 60, "The quiz total duration")	
	flag.Parse()

	// Show values used during the execution
	fmt.Println("The file used for this quiz is:", *filename)
	fmt.Println("The timeout for this quiz is:", *timeout)
	fmt.Println("-----------------------------------")
	fmt.Println()
	
	records = parse_csv_file(filename)
	mytotal := start_quiz(records, *timeout)
	fmt.Printf("Your total score is: %v/%v\n", mytotal, len(records))
}

func start_quiz (records[][] string, dur int) (total int){
	var first, second, result, user_result string
	total = 0
	timeoutchan := make(chan int)
	responsechan := make(chan string)
	
	go timeoutcounter(dur, timeoutchan)
	
	for key := range(records) {
		result = records[key][2]

		go func ()  {
			first = records[key][0]
			second = records[key][1]	
			fmt.Printf("What's the result of %v + %v ?\n",first, second)
			fmt.Scan(&user_result)
			responsechan <- user_result
		}()

		select {
		case <- timeoutchan:
			return total
		case answer := <- responsechan:
			if answer == result {
				total += 1
			}
		}
	}

	return total
}

func parse_csv_file(filename *string) [][]string {
	//Get the current working directory
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	
	fullpath := pwd + "/" + *filename

	file, err := os.Open(fullpath)
		if err != nil {
		fmt.Println(err)
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	
	if err != nil {
		fmt.Println(err)
	}
	
	defer file.Close()
	return records
}

func timeoutcounter(dur int, timeoutchan chan int) {
	time.Sleep(time.Second * time.Duration(dur))	
	timeoutchan <- 1
	defer close(timeoutchan)
}