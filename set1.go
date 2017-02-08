package main

import "fmt"
import "encoding/hex"
import "strings"

func main() {
	
	//Challenge 3: Determine the random character which was XORd with the inputstring to decode it.
	inputstring := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	inputByteArray, _ := hex.DecodeString(inputstring)

	dictionary := []string {"the", "be", "to", "of", "and", "a", "in", "that", 
	"have", "I", "it", "for", "not", "on", "with", "he", 
	"as", "you", "do", "at", "this", "but", "his", "by", "from"}

	highestWC := 0;
	mostLikelyChar := -1

	for i:=0; i < 16; i++ {
		

		testArray := make([]byte, len(inputByteArray))
		copy(testArray, inputByteArray)
		testArray[0] = testArray[0] ^ byte(i);
		testString := string(testArray)
		fmt.Println(testString)

		localWC := 0;
		for j, _ := range dictionary {
        	if(strings.Contains(testString, dictionary[j])) {
        		localWC++
        	}	

    	}

    	if(localWC > highestWC) {
    		highestWC = localWC
    		mostLikelyChar = i
    	}
    }

    fmt.Println(mostLikelyChar)
}




