package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/brentp/vcfgo"
	"github.com/brentp/xopen"
	//"github.com/shenwei356/bio/seqio/fastx"
	"github.com/srynobio/go-vmc/vmc"
	//"log"
	//	"os"
	"runtime"
	//"time"
)

func main() {

	var args struct {
		VCF       string `help: "VCF file to annotate with VMC digest ids"`
		Fasta     string `help: "Reference fasta file used to create VCF file."`
		CPUS      int    `arg: "-cpus limit the number of available cpus."`
		NameSpace string `arg: "-nameSpace: Accessioning authority used to build identifier. Default: VMC"`
	}
	args.CPUS = 0
	args.NameSpace = "VMC"
	arg.MustParse(&args)

	// Set GOMAXPROCESS
	runtime.GOMAXPROCS(args.CPUS)

	fh, err := xopen.Ropen(args.VCF)
	if err != nil {
		panic("VCF file not given.")
	}
	defer fh.Close()

	rdr, err := vcfgo.NewReader(fh, false)
	if err != nil {
		panic("Could not read record from given VCF file.")
	}
	defer rdr.Close()

	// set up vmc record.
	record := vmc.Initialize()

	// Add VMC INFO to the header.
	rdr.AddInfoToHeader("VMCGSID", "1", "String", "VMC Sequence identifier")
	rdr.AddInfoToHeader("VMCGLID", "1", "String", "VMC Location identifier")
	rdr.AddInfoToHeader("VMCGAID", "1", "String", "VMC Allele identifier")
	//	rdr.AddInfoToHeader("VMC Generation date", "1", "String", generated_at)

	for {
		variant := rdr.Read()
		if variant == nil {
			break
		}

		// Check for alternate alleles.
		altAllele := variant.Alt()
		if len(altAllele) > 1 {
			panic("multiallelic variants found, please pre-run vt decomposes.")
		}

		location := record.DigestLocation(variant, args.NameSpace)
		allele := record.DigestAllele(location, variant, args.NameSpace)

	}
}

// ------------------------- //
