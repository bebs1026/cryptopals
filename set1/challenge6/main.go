package main

import (
	"fmt" 
	"encoding/hex"
)

func main() {
	
	string1 := "this is a test"
	string2 := "wokka wokka!!!"
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

	for i, e := range str1Array {
		popCount := 0
		xorArray := make([]byte, len(str1Array))
		xorArray = str1Array ^ str2Array
	}

	return distance
}

func getPopcount() int {
	popcount := 0

	return popcount
}
