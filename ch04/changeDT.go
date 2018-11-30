package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide one text file to process!")
		os.Exit(1)
	}

	filename := arguments[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file %s\n", err)
		os.Exit(1)
	}
	defer f.Close()

	notAMatch := 0
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s\n")
		}

		// r1 := regexp.MustCompile(".*\\[(\\d\\d\\/\\w+/\\d\\d\\d\\d:\\d\\d:\\d\\d:\\d\\d.*)\\] .*")
		r1 := regexp.MustCompile("\\d\\d\\d\\d/\\d\\d/\\d\\d\\s\\d\\d:\\d\\d:\\d\\d")
		if r1.MatchString(line) {
			// fmt.Println("MATCHED")
			match := r1.FindStringSubmatch(line)
			// fmt.Println(match[0])
			d1, err := time.Parse("2006/01/02 15:04:05", match[0])
			if err == nil {
				newFormat := d1.Format(time.Stamp)
				fmt.Print(strings.Replace(line, match[0], newFormat, 1))
			} else {
				notAMatch++
			}
			continue
		}

		r2 := regexp.MustCompile(".*\\[(\\w+\\-\\d\\d-\\d\\d:\\d\\d:\\d\\d:\\d\\d.*)\\] .*")
		if r2.MatchString(line) {
			match := r2.FindStringSubmatch(line)
			d1, err := time.Parse("Jan-02-06:15:04:05 -0700", match[1])
			if err == nil {
				newFormat := d1.Format(time.Stamp)
				fmt.Print(strings.Replace(line, match[1], newFormat, 1))
			} else {
				notAMatch++
			}
			continue
		}
	}
	fmt.Println(notAMatch, "lines didn't match!")

}
