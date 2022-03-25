package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func readCasefile(filepath string) []string {
	var filecontent []string
	file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	var size = stat.Size()
	fmt.Println("file size=", size)

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				return filecontent
			} else {
				fmt.Println("Read file error!", err)
				return filecontent
			}
		} else {
			text := line
			filecontent = append(filecontent, text)
			return filecontent
		}
	}
}

func TestIsLeapYear(t *testing.T) {
	testCase := readCasefile("leapYearCase.txt")
	for _, s := range testCase {
		year, _ := strconv.Atoi(s)
		if ans := IsLeapYear(year); ans != true {
			t.Errorf("Expected be true, but got %t", ans)
		}
	}
}

func TestIsValidDate(t *testing.T) {
	validtCase := readCasefile("validDate.txt")

	for _, s := range validtCase {
		var month, day, year int
		date := strings.Split(s, "/")
		month, _ = strconv.Atoi(date[0])
		day, _ = strconv.Atoi(date[1])
		year, _ = strconv.Atoi(date[2])
		if ans := IsValidDate(month, day, year); ans == false {
			t.Errorf("Issue Date:%d, %d, %d", month, day, year)
			t.Errorf("Expected be true, but got %t", ans)
		}
	}
	invalidtCase := readCasefile("invalidDate.txt")
	for _, s := range invalidtCase {
		var month, day, year int
		date := strings.Split(s, "/")
		month, _ = strconv.Atoi(date[0])
		day, _ = strconv.Atoi(date[1])
		year, _ = strconv.Atoi(date[2])
		if ans := IsValidDate(month, day, year); ans == true {
			t.Errorf("Issue Date:%d, %d, %d", month, day, year)
			t.Errorf("Expected be false, but got %t", ans)
		}
	}
}
