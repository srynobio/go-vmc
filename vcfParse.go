package main

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/brentp/vcfgo"
	"github.com/brentp/xopen"
	"log"
	"os"
	//	"strings"
)

func main() {
	fh, err := xopen.Ropen(os.Args[1])
	eCheck(err)
	defer fh.Close()

	rdr, err := vcfgo.NewReader(fh, false)
	eCheck(err)

	for {
		variant := rdr.Read()
		if variant == nil {
			break
		}
		//interStart := variant.Start() - 1
		//interEnd := variant.End() + 1
		altAllele := variant.Alt()

		if len(altAllele) > 1 {
			log.Panicln("vt....Alternative allele found!")
		}

		fmt.Println(altAllele)

	}

}

// ------------------------- //
func eCheck(p error) {
	if p != nil {
		panic(p)
	}
}

// ------------------------- //

func DigestId(bv []byte, truncate int) string {
	hasher := sha512.New()
	hasher.Write(bv)

	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil)[:truncate])
	return sha
}

// ------------------------- //
