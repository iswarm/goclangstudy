package main

import "fmt"
import "os"
import "log"
import "bytes"
import "path/filepath"
//import "github.com/ethereum/go-ethereum/log"
import "github.com/urfave/cli"
//import "github.com/ethereum/go-ethereum/ethdb"

import (
    //"github.com/ethereum/go-ethereum/log"
	//"github.com/ethereum/go-ethereum/metrics"
	"github.com/syndtr/goleveldb/leveldb"
	//"github.com/syndtr/goleveldb/leveldb/errors"
	//"github.com/syndtr/goleveldb/leveldb/filter"
	//"github.com/syndtr/goleveldb/leveldb/iterator"
	//"github.com/syndtr/goleveldb/leveldb/opt"
	//"github.com/syndtr/goleveldb/leveldb/util"
)


func main(){
    
    app := cli.NewApp()
    app.Flags = []cli.Flag{
        cli.StringFlag{
            Name: "lang",
            Value:"english",
            Usage: "language for the greeting",
        },
        cli.StringFlag{
            Name:"levelDb",
            Value: "",
            Usage:"learning to call the api of level-db",
        },
        cli.StringFlag{
            Name:"openFile",
            Value:"",
            Usage:"open the default file",
        },
    }

    app.Action = func(c *cli.Context) error {
        name :="Nefertiti"
        if c.NArg() >0 {
            name = c.Args().Get(0)            
        }
        if c.String("lang") == "spanish" {
            fmt.Println("Hola", name)
        } else if c.String("openFile") == "openFile"{
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
        } else if c.String("levelDb") == "levelDb" {
            //db *leveldb.DB
            db,err := leveldb.OpenFile("/Users/mac/goclangstudy", nil)
            
            defer db.Close()
            if err != nil {
                fmt.Print("Can not open database file")
            }
            err = db.Put([]byte("productCode"), []byte("01010112132"), nil)
            if err != nil {
                fmt.Print("Can not put database file")
            }
            data, err := db.Get([]byte("productCode"),nil)
            if err != nil {
                fmt.Print("Can not get database file")
            }
            fmt.Print("data = \n", string(data[:]) )
            err = db.Delete([]byte("productCode"), nil)
            if err != nil {
                fmt.Print("Can not delete database file")
            }
        }
        return nil
    }
    app.Run(os.Args)
}


