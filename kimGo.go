package main

import "fmt"
import "os"
import "log"
import "bytes"
//import "github.com/ethereum/go-ethereum/log"
//import "github.com/urfave/cli.v1"


func main(){
    file, err := os.Open("/Users/mac/goclangstudy/README.md")
    if err != nil {
        log.Fatal(err)
    }
    
        data:=make([]byte, 1000)
        count,err:=file.Read(data)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("read %d bytes: %q\n", count, data[:count])

    var (
        buf bytes.Buffer
        logger = log.New(&buf, "logger: ", log.Lshortfile)
    )

    logger.Print("Hello, log file!")

    fmt.Print(&buf)
   
}


