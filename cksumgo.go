package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
)

var (
	p        = fmt.Printf
	isForAll = flag.Bool("A", false, "calculates md5,sha1,sha256,sha512 for one")
	topt     = flag.String("t", "md5", "default is md5. You can use md5,sha1,sha256,sha512.")
	// when '-t sha1', *topt is set to "sha1"
)

func getCheckSum(hashname string, file *os.File) (string, string, string) {
	var h hash.Hash
	switch hashname {
	case "sha1":
		h = sha1.New()
	case "sha256":
		h = sha256.New()
	case "sha512":
		h = sha512.New()
	default:
		h = md5.New()
		hashname = "md5"
	}
	if _, err := io.Copy(h, file); err != nil {
		log.Fatal(err)
	}
	cksum := fmt.Sprintf("%x", h.Sum(nil))
	return cksum, hashname, file.Name()
}

func printForOne(t string, file *os.File) {

	h, ty, fn := getCheckSum(t, file)
	p("%s -- [%s] %s\n", fn, ty, h)
}

func printForAll(file *os.File) {

	p("%s -- \n", file.Name())

	for _, t := range []string{"md5", "sha1", "sha256", "sha512"} {
		h, ty, _ := getCheckSum(t, file)
		p("[%s] %s\n", ty, h)

	}

}

func main() {
	flag.Parse()

	for _, af := range flag.Args() {
		file, err := os.Open(af)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		if *isForAll == true {
			printForAll(file)
		} else {
			printForOne(*topt, file)
		}
	}
}
