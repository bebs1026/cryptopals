package main

import (
	"fmt" 
	"encoding/hex"
	"strings"
    "bufio"
    "log"
    "os"
    "path/filepath"
)

func main() {
	
	inputStringArray := readInputStringsFromFile()
	
	dictionary := []string {"the", "be", "to", "of", "and", "in", "that", 
	"have", "it", "for", "not", "on", "with", "he", 
	"as", "you", "do", "at", "this", "but", "his", "by", "from", "they",
	"we", "say", "her", "she", "or", "an", "will", "my", "one", "all", "would",
	"there", "their", "what", "so", "up", "out", "if", "about", "who", "get",
	"which", "go", "me", "when", "make", "can", "like", "time", "no", "just",
	"him", "know", "take", "people", "into", "year", "your", "good", "some",
	"could", "them", "see", "other", "than", "then", "now", "look", "only",
	"come", "its", "over", "think", "also", "back", "after", "use", "two",
	"how", "our", "work", "first", "well", "way", "even", "new", "want",
	"because", "any", "these", "give", "day", "most", "us"}

	highestWC := 0;
	mostLikelyChar := "none"
	decodedString := "not found yet"

	for _, inputString := range inputStringArray {
		
		inputByteArray, _ := hex.DecodeString(inputString)
		for i:=0; i<128; i++ {

			testArray := make([]byte, len(inputByteArray))
			copy(testArray, inputByteArray)
			
			for k:=0; k<len(testArray); k++ {
				testArray[k] = (testArray[k] ^ byte(i));
			}

			localWC := 0;
			for j, _ := range dictionary {
	        	if(strings.Contains(string(testArray), dictionary[j])) {
	        		localWC++
	        	}	

	    	}

	    	if(localWC > highestWC) {
	    		highestWC = localWC
	    		mostLikelyChar = string(i)
	    		decodedString = string(testArray)
	    	}

	    }
	}

    fmt.Println(highestWC)
    fmt.Println(mostLikelyChar)
   	fmt.Println(decodedString)
}

func readInputStringsFromFile() []string{
    
	result := make([]string, 0)

    absPath, _ := filepath.Abs("./file.txt")
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




