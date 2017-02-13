	

package main

import (
	"fmt" 
    "log"
    "os"
    "path/filepath"
    "io"
    "bytes"
    "encoding/base64"
    "crypto/aes"
)



func main() {
	inputString := readInputStringFromFile()
	inputByteArray, err := base64.StdEncoding.DecodeString(inputString)

	if(err != nil) {
		fmt.Println("Error decoding base64 string: ", err)
	}

	cipher := []byte("YELLOW SUBMARINE")

	result := string(decryptAes128Ecb(inputByteArray, cipher))

	fmt.Println("Decrypted string is: ", result)

}

func decryptAes128Ecb(data, key []byte) []byte {
    cipher, _ := aes.NewCipher([]byte(key))
    decrypted := make([]byte, len(data))
    size := 16

    for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
        cipher.Decrypt(decrypted[bs:be], data[bs:be])
    }

    return decrypted
}

func readInputStringFromFile() string {
    
    absPath, _ := filepath.Abs("./challenge1-7.txt")
	buf := bytes.NewBuffer(nil)
  	f, err := os.Open(absPath)
  	    if err != nil {
        log.Fatal(err)
    }
  	io.Copy(buf, f)           
  	f.Close()
  	
  	return string(buf.Bytes())
}