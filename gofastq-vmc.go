package main

import (
	"crypto/sha512"
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/shenwei356/bio/seqio/fastx"
	"io"
	"os"
)

const goUsage = `



Required options:

	-fasta	   : Fasta file to used to create VMC based ID.

	-nameSpace :	Namespace accession authority 
					Examples: NCBI, ENsembl, UCSC

Additional options:

`

func main() {

	fqPtr := flag.String("fasta", "", "Original fasta file to create VMC IDs")
	namespacePtr := flag.String("nameSpace", "", "Accessioning authority")
	flag.Parse()

	if *fqPtr == "" && *namespacePtr == "" {
		fmt.Println(goUsage)
		os.Exit(1)
	}

	// Incoming fastq file.
	reader, err := fastx.NewDefaultReader(*fqPtr)
	eCheck(err)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		eCheck(err)

		//accession := string(record.Name)
		accession := string(record.ID)
		digestID := VMCDigestId(record.Seq.Seq, 24)
		compactURI := *namespacePtr + ":" + accession

		fmt.Println("What will be added to db:")
		fmt.Println("Namespace: ", *namespacePtr)
		fmt.Println("Accession: ", accession)
		fmt.Println("SeqID: ", digestID)
		fmt.Println("CURIE: ", compactURI)
		fmt.Println("VMC_SeqID: ", "VMC:GS_"+digestID)

	}
}

// ------------------------ //

func VMCDigestId(bv []byte, truncate int) string {
	hasher := sha512.New()
	hasher.Write(bv)

	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil)[:truncate])
	return sha
}

// -------------------------------------------- //

func eCheck(p error) {
	if p != nil {
		panic(p)
	}
}

// -------------------------------------------- //

/*
   filePtr := flag.String("bam", "", "lossless BAM file to validate")
   cpus := flag.Int("cpus", 0, "Number of cpus to allow co-processing.")
   auxTag := flag.String("tag", "OQ", "Required tag all reads must contain.")
   flag.Parse()
*/
