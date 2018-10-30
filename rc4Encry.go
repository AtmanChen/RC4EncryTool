package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
)

var (
	h      bool   /* help command */
	r      bool   /* scan folder recursivly and encode */
	s      bool   /* encode single file only */
	dir    string /* folder directory */
	jsPath string /* filepath of js file */
)

type rc4EncryToolArgs struct {
	h        bool
	r        bool
	s        bool
	filePath string
}

func main() {
	flag.Parse()
	var args rc4EncryToolArgs
	flag.BoolVar(&args.h, "h", false, "this help")
	flag.BoolVar(&args.r, "r", false, "scan folders recursivly")
	flag.BoolVar(&args.s, "s", false, "encode single js file")
	args.filePath = flag.Args()[0]
	args.encodeFile()
	flag.Usage = displayUsage
}

func displayUsage() {
	fmt.Fprintf(os.Stderr, "RC4EncryTool version: RC4Encry/1.0.1 Usage: RC4Encry [-hrs] [-r folder directory] [-s js file path]")
	flag.PrintDefaults()
}

func (a *rc4EncryToolArgs) encodeFile() {

}

func encodeJSFile(jsPath string) {
	fmt.Println(path.Dir(jsPath))
	ext := path.Ext(jsPath)
	fmt.Println(ext)
	if ext != ".js" {
		fmt.Println("Not a js file.")
		return
	}
	file, err := os.Open(jsPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(file)
}
