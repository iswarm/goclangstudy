package main

import "fmt"
import "os"
import "log"
import "bytes"
import "path/filepath"
//import "github.com/ethereum/go-ethereum/log"
import "github.com/urfave/cli"


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

    logger.Print("Hello, log file!\n")

    fmt.Print(&buf)
    fmt.Print("filepath.Base(os.Args[0]) = \n", filepath.Base(os.Args[0]))
    app := cli.NewApp()
    app.Flags = []cli.Flag{
        cli.StringFlag{
            Name: "lang",
            Value:"english",
            Usage: "language for the greeting",
        },
    }

    app.Action = func(c *cli.Context) error {
        name :="Nefertiti"
        if c.NArg() >0 {
            name = c.Args().Get(0)            
        }
        if c.String("lang") == "spanish" {
            fmt.Println("Hola", name)
        } else {
            fmt.Println("Hello", name)
        }
        return nil
    }
    app.Run(os.Args)
}


