package main

import "fmt"
import "encoding/hex"
import "strings"

func main() {
	
	//Challenge 3: Determine the random character which was XORd with the inputstring to decode it.
	inputstring := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	inputByteArray, _ := hex.DecodeString(inputstring)

	characters := []string {"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", 
	"k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P",
	"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	
	dictionary := []string {"the", "be", "to", "of", "and", "a", "in", "that", 
	"have", "I", "it", "for", "not", "on", "with", "he", 
	"as", "you", "do", "at", "this", "but", "his", "by", "from", "like"}

	highestWC := 0;
	mostLikelyChar := "none"

	for _, c := range characters {

		testArray := make([]byte, len(inputByteArray))
		copy(testArray, inputByteArray)
		
		for k:=0; k<len(testArray); k++ {
			testArray[k] = (testArray[k] ^ []byte(c)[0]);
		}

    	fmt.Println(string(testArray))
		
		localWC := 0;
		for j, _ := range dictionary {
        	if(strings.Contains(string(testArray), dictionary[j])) {
        		localWC++
        	}	

    	}

    	if(localWC > highestWC) {
    		highestWC = localWC
    		mostLikelyChar = c
    	}

    }

    fmt.Println(mostLikelyChar)
}




