package main

/* Reference info:
ftp://ftp.ncbi.nlm.nih.gov/genomes/refseq/vertebrate_mammalian/Homo_sapiens/all_assembly_versions/
*/

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/shenwei356/bio/seqio/fastx"
	"io"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	// open fastq file
	reader, err := fastx.NewDefaultReader(os.Args[1])
	eCheck(err)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		eCheck(err)

		namespaceAccession := "NCBI:" + string(record.ID)
		fmt.Println(namespaceAccession)

		matched, err := regexp.Match("NC_", record.ID)
		eCheck(err)

		if matched == true {
			sequenceRecord := record.Seq
			input := []byte(sequenceRecord.Seq)
			fmt.Println(digest(input, 24))
		} else {
			fmt.Println("Fasta record does not contain NC_")
			fmt.Println(string(record.ID))
			continue
		}
	}
}

// ------------------------ //

func digest(bv []byte, truncate int) string {
	hasher := sha512.New()
	hasher.Write(bv)

	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil)[:truncate])
	vmcID := "VMC:GS_" + sha
	return vmcID
}

// -------------------------------------------- //

func eCheck(p error) {
	if p != nil {
		panic(p)
	}
}

// -------------------------------------------- //
