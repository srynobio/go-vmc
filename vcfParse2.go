package main

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/brentp/vcfgo"
	"github.com/brentp/xopen"
	"log"
	"os"
	"time"
)

var version = "v1.0.0"

/*
	Define all struts used in
	VMC model software.
*/

var Id string
var DateTime = time.Now()

type Identifier struct {
	accession string
	namespace string
}

type Interval struct {
	start uint32
	end   uint32
}

type Location struct {
	id          string
	interval    int
	sequence_id string
}

type Allele struct {
	id          string
	location_id string
	state       string
}

type Genotype struct {
	id            string
	haplotype_ids []string
	completedness int
}

type Haplotype struct {
	id            string
	allele_id     []string
	completedness int
}

type VMCID struct {
	VMCidentifierID string
	SequenceID      string
	LocationID      string
	AlleleID        string
	HaplotypeID     string
	GenotypeID      string
}

func main() {

	fmt.Println(DateTime)

	fh, err := xopen.Ropen(os.Args[1])
	eCheck(err)
	defer fh.Close()

	rdr, err := vcfgo.NewReader(fh, false)
	eCheck(err)
	defer rdr.Close()

	var vmc VMCID
	var allele Allele
	var interval Interval

	// set DummySeqID
	vmc.SequenceID = "VMC:GS_Ya6Rs7DHhDeg7YaOSg1EoNi3U_nQ9SvO"

	for {
		variant := rdr.Read()
		if variant == nil {
			break
		}

		// set values from variant.
		altAllele := variant.Alt()
		start := variant.Start() - 1
		end := variant.End() + 1

		if len(altAllele) > 1 {
			log.Panicln("multiallelic variants found, please pre-run vt decomposes.")
		}

		// set non VMC struct values.
		allele.state = altAllele[0]
		allele.location_id = vmc.LocationID

		interval.start = start
		interval.end = end

		/* call SQL to get:
		1 - namespace & accession
		2 - VMC sequence_id
		*/

		// set VMC struct
		vmc.LocationID = VMCLocationID(vmc, interval)
		vmc.AlleleID = VMCAlleleID(vmc, allele)

		fmt.Println(allele)
	}
}

// ------------------------- //
// VMC functions
// ------------------------- //

func VMCLocationID(v VMCID, i Interval) string {

	seqID := v.SequenceID
	intervalString := fmt.Sprint(i.start) + ":" + fmt.Sprint(i.end)

	location := "<Location:<Identifier:" + seqID + ">:<Interval:" + intervalString + ">>"
	DigestLocation := VMCDigestId([]byte(location), 24)
	return "VMC:GL_" + DigestLocation
}

// ------------------------- //

func VMCAlleleID(v VMCID, a Allele) string {

	vmcLocation := v.LocationID
	state := a.state

	allele := "<Allele:<Identifier:" + vmcLocation + ">:" + state + ">"
	DigestAllele := VMCDigestId([]byte(allele), 24)
	return "VMC:GA_" + DigestAllele
}

// ------------------------- //

func VMCDigestId(bv []byte, truncate int) string {
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
