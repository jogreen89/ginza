// parse.go
//
// (c) 2016 zubernetes
package main

import (
	"fmt"
	//"image/gif"
	"io/ioutil"
	//"io"
	"log"
	"os"	
	"strings"
	"time"
)

// Program constants.
const CURRENT_SYSTEM_OS string = "DEFAULT"
const OPEN_DIRECTORY string = "DIRECTORY"
const OPEN_FILE string = "FILE"
const PROGRESS_BAR_SPEED_OPEN_DIRECTORY int64 = 10
const PROGRESS_BAR_SPEED_OPEN_FILE int64 = 20
const SYSTEM_OS_WINDOWS string = "WINDOWS"
const SYSTEM_OS_MACOSX string = "MAC OS X"
const TERMINAL_WIDTH_80 int = 80
const TERMINAL_WIDTH_160 int = 160

// Checks for failures.
func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

func checkVowels(b []byte) {
	fmt.Println("Here are the VOWELS found in your file -- >\n"+
				"*******************************************")
	fmt.Print("a -- ")
	fmt.Println(strings.Count(string(b), "a"))
	fmt.Print("e -- ")
	fmt.Println(strings.Count(string(b), "e"))
	fmt.Print("i -- ")
	fmt.Println(strings.Count(string(b), "i"))
	fmt.Print("o -- ")
	fmt.Println(strings.Count(string(b), "o"))
	fmt.Print("u -- ")
	fmt.Println(strings.Count(string(b), "u"))

}

func resume() bool {
	var (
		s = " "
		b = true
	)

	fmt.Println("Would you like to continue?")
	_, err := fmt.Scanf("%v", &s)
	check(err)

	if (s == "y" || s == "Y" || s == "yes") {
		fmt.Println("Thank you for using this program.")	
	} else {
		displayGoodbye()
		b = false
	}

	return b
}

func displayOptions() {
	fmt.Println("1. Count how many vowels (a, e, i, o, u) appear (UTF-8).\n"+
				"2. Count instances of a target word (UTF-8).\n"+
				"3. Decode a .gif file (.gif).")
}

func displayTitle() {
	fmt.Println("Hello and thank you for using my program")
	fmt.Println("The current time is ", time.Now())
	fmt.Println("Would you like to open the input file? [y/n]")
}

func displayGoodbye() {
	fmt.Println("Thank you for using my program! Goodbye.")
}

func findWord(b []byte) {
	s := " "
	fmt.Println("What word are you looking for?")
	_, err := fmt.Scanf("%v", &s)
	check(err)
	x := strings.Count(string(b), s)
	fmt.Print("There are ")
	fmt.Print(x)
	fmt.Print(" instances of \""+s+"\" in the file.\n")
}

func ioHandler(s string) string {
	if (s == SYSTEM_OS_MACOSX) {
		_, err := fmt.Scanf("%v", &s)
		check(err)
	} else {
		_, err := fmt.Scanf("%f\n", &s)
		check(err)
	}
	return s
}

// func Decode(r io.Reader) (image.Image, error)
func gifDecoder(s string) {
	reader, err := os.Open(s)
	check(err)
	defer reader.Close()
	/*
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}*/
}

func main () {
	var (
		s string = ""
		current_os = "DEFAULT"
		b bool = true
	)

	displayTitle()

	current_os = SYSTEM_OS_WINDOWS
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if (pair[1]=="iTerm.app") {
			current_os = SYSTEM_OS_MACOSX
		} 
	}

	s = ioHandler(current_os)

	b = startProgram(s)

	if (b) {
		for (b) {
			s = openDirectory()
			openLocalFile(s)
			b = resume()
		}
	} else {
		fmt.Println("Exiting.")
	}
}

func openDirectory() string {
	s := " "

	progressBarTerm80(OPEN_DIRECTORY)
	fmt.Println("Which file would you like to scan for data?")

	data, err := ioutil.ReadDir(".")
	check(err)

	directoryFiles := len(data)
	fmt.Println(data[0].Name())
	i := 0
	for i < directoryFiles {
		fmt.Println(data[i].Name())
		i += 1
	}

	fmt.Println("Which file would you like to open?")

	_, err = fmt.Scanf("%v", &s)
	check(err)

	return s
}

func openLocalFile(inFile string) {
	s := " "

	data, err := ioutil.ReadFile(inFile)

	check(err)

	fmt.Println("What would you like to parse?")

	displayOptions()

	_, err = fmt.Scanf("%v", &s)

	switch {
	case "1" == s:
		progressBarTerm80(OPEN_FILE)
		checkVowels(data)
	case "2" == s:
		progressBarTerm80(OPEN_FILE)
		findWord(data)
	case "3" == s:
		progressBarTerm80(OPEN_FILE)
		fmt.Println("Opened a gif.")
	}
}

func progressBarTerm80(s string) {
	if (s == OPEN_DIRECTORY) {
		fmt.Println("** OPENING LOCAL DIRECTORY **")
	} else {
		fmt.Println("** READING YOUR INPUT FILE **")
	}
	sum := 0
	for sum < TERMINAL_WIDTH_80 { 
		sum += 1
		fmt.Print("*")
		if (s == OPEN_DIRECTORY) {
			time.Sleep(time.Duration(PROGRESS_BAR_SPEED_OPEN_DIRECTORY)*time.Millisecond)
		} else {
			time.Sleep(time.Duration(PROGRESS_BAR_SPEED_OPEN_FILE)*time.Millisecond)
		}
	}
	fmt.Println()
}

func startProgram(s string) bool {
	var decision bool
	if (s == "y") {
		decision = true
	} else {
		decision = false
	}
	return decision
}
