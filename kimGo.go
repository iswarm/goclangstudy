package main

import "fmt"
import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/hex"
    "io"
)


func main(){
    key,_ := hex.DecodeString("6368616e676520746869732070617373")
    fmt.Printf("%d\n", aes.BlockSize)
    str := "exampleplaintext"
    fmt.Printf("%d\n", len(str))
    plaintext := []byte("exampleplaintextexampleplaintextexampleplaintext")
    if len(plaintext)%aes.BlockSize != 0 {
        panic("plaintext is not a multiple of the block size")
    }

    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }
    ciphertext := make([]byte, aes.BlockSize+len(plaintext))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        panic(err)
    }
    mode := cipher.NewCBCEncrypter(block, iv)
    mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
    fmt.Printf("%x\n",ciphertext)
}


