package main

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/brentp/vcfgo"
	"github.com/brentp/xopen"
	"log"
	"os"
)

type VMC struct {
	Chromosome      string
	Interval        string
	State           string
	VMCidentifierID string
	SequenceID      string
	LocationID      string
	AlleleID        string
	HaplotypeID     string
	GenotypeID      string
}

func main() {
	fh, err := xopen.Ropen(os.Args[1])
	eCheck(err)
	defer fh.Close()

	rdr, err := vcfgo.NewReader(fh, false)
	eCheck(err)
	defer rdr.Close()

	var vb VMC

	// set DummySeqID
	vb.SequenceID = "VMC:GS_Ya6Rs7DHhDeg7YaOSg1EoNi3U_nQ9SvO"

	for {
		variant := rdr.Read()
		if variant == nil {
			break
		}

		// Check for alternate alleles.
		altAllele := variant.Alt()
		if len(altAllele) > 1 {
			log.Panicln("multiallelic variants found, please pre-run vt decomposes.")
		}

		// Define VMC location
		vb.Chromosome = variant.Chrom()
		vb.Interval = fmt.Sprint(variant.Start()-1) + ":" + fmt.Sprint(variant.End()+1)
		vb.State = altAllele[0]

		vb.LocationID = LocationID(vb)
		vb.AlleleID = AlleleID(vb)

		fmt.Println(vb)
	}
}

// ------------------------- //
// VMC functions
// ------------------------- //

func LocationID(class VMC) string {

	seqID := class.SequenceID
	interval := class.Interval

	location := "<Location:<Identifier:" + seqID + ">:<Interval:" + interval + ">>"
	DigestLocation := DigestId([]byte(location), 24)
	return "VMC:GL_" + DigestLocation
}

// ------------------------- //

func AlleleID(class VMC) string {

	vmcLocation := class.LocationID
	state := class.State

	allele := "<Allele:<Identifier:" + vmcLocation + ">:" + state + ">"
	DigestAllele := DigestId([]byte(allele), 24)
	return "VMC:GA_" + DigestAllele
}

// ------------------------- //

func DigestId(bv []byte, truncate int) string {
	hasher := sha512.New()
	hasher.Write(bv)

	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil)[:truncate])
	return sha
}

// ------------------------- //

func eCheck(p error) {
	if p != nil {
		panic(p)
	}
}

// ------------------------- //
