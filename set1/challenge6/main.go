package main

import (
	"fmt" 
	"bufio"
    "log"
    "os"
    "path/filepath"
)

func main() {
	
	string1 := "this is a test"
	string2 := "wokka wokka!!!"
	editDistance := computeEditDistance(string1, string2)
	fmt.Println("Edit distance is: ", editDistance)
}

func readInputStringsFromFile() []string{
    
	result := make([]string, 0)

    absPath, _ := filepath.Abs("./challenge6.txt")
    file, err := os.Open(absPath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        result = append(result, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return result
}

func determineKeySize() int {
	return 0;
}

func computeEditDistance(str1, str2 string) int {
	distance := 0

	str1Array := []byte(str1)
	str2Array := []byte(str2)

	for i, _ := range str1Array {
		xorArray := str1Array[i] ^ str2Array[i]
		distance = distance + getPopcount(uint8(xorArray))
	}

	return distance
}

func getPopcount(bite uint8) int {
	popcount := 0
	for i:=0; i<=7; i++ {

		//fmt.Println("iterating through popcount with ", i)
		if( ( (bite >> uint8(i)) & 1) == 1) {
			popcount++
			//fmt.Println("just upped popcount")
		}
	}
	return popcount
}
