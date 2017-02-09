package main

import (
	"fmt" 
	"encoding/hex"
)

func main() {
	
	inputString := "Burning 'em, if you ain't quick and nimble I go crazy when I hear a cymbal"
	xorKey := "ICE"
	
	inputByteArray:=[]byte(inputString)
	updatedArray := make([]byte, len(inputByteArray))
	
	for i:=0; i<len(inputByteArray); i++ {
		updatedArray[i] = (inputByteArray[i] ^ []byte(xorKey)[i%len(xorKey)])	
    }

	resultString := hex.EncodeToString(updatedArray)   
	fmt.Println(resultString)
}





