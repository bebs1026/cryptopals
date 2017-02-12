package main

import (
	"fmt" 
    "log"
    "os"
    "path/filepath"
    "strings"
    "io"
    "bytes"
    "encoding/base64"
)



func main() {
	
	//string1 := "this is a test"
	//string2 := "wokka wokka!!!"
	//editDistance := computeEditDistance(string1, string2)

	inputString := readInputStringFromFile()
	inputByteArray, err := base64.StdEncoding.DecodeString(inputString)

	if(err != nil) {
		fmt.Println("Error decoding base64 string: ", err)
	}

	keySize := calculateKeySizeFromString(inputByteArray)

	fmt.Println("Key size is: ", keySize)

	cipher := ""
	m := make(map[int][]byte)

	for i,j := range inputByteArray {
		curArray := m[(i%keySize)]
		m[(i%keySize)] = append(curArray,j)
	}

	for i:=0;i<len(m);i++ {
		cipher += findSingleCharXORCiper(m[i])
	}

	fmt.Println("Attempting to decode with cipher: ", cipher)
	translatedString := applyMulitCharCipher (cipher, inputByteArray)
	fmt.Println("Translated string is: ", string(translatedString))
}

func calculateKeySizeFromString(inputByteArray []byte) int {
	lowestEditDistance := 10000000000000.00
	keysize := 0;

	for i:=2; i<=40; i++ {

		slice1 := inputByteArray[:i]
		slice2 := inputByteArray[i:(i*2)]
		slice3 := inputByteArray[(i*2):(i*3)]
		slice4 := inputByteArray[(i*3):(i*4)]
		slice5 := inputByteArray[(i*4):(i*5)]
	
		oneToTwoDistanceUnadjusted := (computeEditDistance(string(slice1), string(slice2)) )
		oneToTwoEditDistance := (float64(oneToTwoDistanceUnadjusted) / float64(i))
		oneToThreeDistanceUnadjusted := (computeEditDistance(string(slice1), string(slice3)) )
		oneToThreeEditDistance := (float64(oneToThreeDistanceUnadjusted) / float64(i))
		oneToFourDistanceUnadjusted := (computeEditDistance(string(slice1), string(slice4)) )
		oneToFourEditDistance := (float64(oneToFourDistanceUnadjusted) / float64(i))
		oneToFiveDistanceUnadjusted := (computeEditDistance(string(slice1), string(slice5)) )
		oneToFiveEditDistance := (float64(oneToFiveDistanceUnadjusted) / float64(i))
		twoToThreeDistanceUnadjusted := (computeEditDistance(string(slice2), string(slice3)) )
		twoToThreeEditDistance := (float64(twoToThreeDistanceUnadjusted) / float64(i))
		threeToFourDistanceUnadjusted := (computeEditDistance(string(slice3), string(slice4)) )
		threeToFourEditDistance := (float64(threeToFourDistanceUnadjusted) / float64(i))
		fourToFiveDistanceUnadjusted := (computeEditDistance(string(slice4), string(slice5)) )
		fourToFiveEditDistance := (float64(fourToFiveDistanceUnadjusted) / float64(i))

		editDistance := ((oneToTwoEditDistance+oneToThreeEditDistance+oneToFourEditDistance+oneToFiveEditDistance+
			twoToThreeEditDistance+threeToFourEditDistance+fourToFiveEditDistance)/7.0)
		
		if(editDistance < lowestEditDistance) {
			lowestEditDistance = editDistance
			keysize = i
		}
	}

	return keysize
}

func readInputStringFromFile() string {
    
    absPath, _ := filepath.Abs("./challenge6.txt")
	buf := bytes.NewBuffer(nil)
  	f, err := os.Open(absPath)
  	    if err != nil {
        log.Fatal(err)
    }
  	io.Copy(buf, f)           
  	f.Close()
  	
  	return string(buf.Bytes())
}

func findSingleCharXORCiper(inputByteArray []byte) string {
	
	upperDictionaryLetters := []string {"E","T","A","O","I","N"," ","S","R","H","L","D","C","U","M",
	 "F","W","Y","P","V","B","G","K","Q","J","X","Z"}

	upperLetterMap := make(map[string]int)
	for i,j := range upperDictionaryLetters {
		upperLetterMap[j] = (len(upperDictionaryLetters)-i)
	}

	highestWC := 0;
	mostLikelyChar := "none"

	for i:=0; i<128; i++ {

		testArray := make([]byte, len(inputByteArray))
		copy(testArray, inputByteArray)
		
		for k:=0; k<len(testArray); k++ {
			testArray[k] = (testArray[k] ^ byte(i));
		}

		localWC := 0;
		
		for i, j := range upperDictionaryLetters {
        	upperCaseString := strings.ToUpper(string(testArray))
	        if(strings.Contains(upperCaseString, upperDictionaryLetters[i])) {
        		localWC=localWC+upperLetterMap[j]
        	}	
        }

    	if(localWC > highestWC) {
    		highestWC = localWC
    		mostLikelyChar = string(i)
    	}

    }
 	
   	return mostLikelyChar
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
		if( ( (bite >> uint8(i)) & 1) == 1) {
			popcount++
		}
	}
	return popcount
}

func applyMulitCharCipher(cipher string, text []byte) string {

	updatedArray := make([]byte, len(text))
	
	for i:=0; i<len(text); i++ {
		updatedArray[i] = (text[i] ^ []byte(cipher)[i%len(cipher)])	
    }

	resultString := string(updatedArray)   
	return resultString
}
