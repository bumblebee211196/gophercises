package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type problem struct {
	Question 	string
	Answer 		string
}

func solution() {
	// Read CLI args
	csvFile := flag.String("csv", "problems.csv", "a csv file in the formal of \"question,answer\"");
	shuffle := flag.Bool("shuffle", false, "shuffle the problems");
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds");
	
	flag.Parse();

	// Read file
	file, err := os.Open(*csvFile);
	if err != nil {
		exit(fmt.Sprintf("Unable to open csv file %s", *csvFile));
	}
	defer file.Close();
	records, err := csv.NewReader(file).ReadAll();
	if err != nil {
		fmt.Printf("Unable to parse csv file. Error: %v", err)
	}
	problems := parseLines(records);
	
	// Shuffle the problems
	if *shuffle {
		problems = shuffleLines(problems);
	}

	// Set timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second);

	correct, total := 0, len(problems);
	// Ask question
	for _, p := range problems {
		fmt.Printf("What is %s ?\n", p.Question);
		answerChannel := make(chan string);
		go func() {
			var answer string;
			fmt.Scanf("%s", &answer);
			answerChannel <- answer;
		}();

		select {
			case <-timer.C:
				exit(fmt.Sprintf("\nYou scored %d out of %d!", correct, total));
			case answer := <-answerChannel:
				if answer == p.Answer {
					correct++;
				}	
		}

	}
	fmt.Printf("\nYou scored %d out of %d!\n", correct, total);
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines));
	for i, line := range lines {
		problems[i] = problem{
			line[0], 
			strings.TrimSpace(line[1])};
	}
	return problems;
}

func shuffleLines(lines []problem) []problem {
	rand.Seed(time.Now().UnixNano());
	rand.Shuffle(len(lines), func(i, j int) { lines[i], lines[j] = lines[j], lines[i] });
	return lines;
}

func exit(msg string) {
	fmt.Println(msg);
	os.Exit(1);
}