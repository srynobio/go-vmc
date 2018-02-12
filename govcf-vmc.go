package main

import (
	"os"
	"strings"

	"github.com/alexflint/go-arg"
	"github.com/brentp/vcfgo"
	"github.com/brentp/xopen"
	"github.com/srynobio/go-vmc/vmc"
)

func main() {

	var args struct {
		VCF       string `arg:"required,help:VCF file to annotate with VMC digest ids. [Required]"`
		FastaFile string `arg:"required,help:Reference fasta file used to create VMC sequence database."`
		DATABASE  string `arg:"required,help:Prebuild VMC sequence database."`
	}
	arg.MustParse(&args)
	outFile := strings.Replace(args.VCF, "vcf", "vmc.vcf", -1)

	fh, err := xopen.Ropen(args.VCF)
	if err != nil {
		panic("VCF file not given or could not be opened.")
	}
	defer fh.Close()

	// create the writer
	output, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	// VCF reader
	rdr, err := vcfgo.NewReader(fh, false)
	if err != nil {
		panic("Could not read given VCF file.")
	}
	defer rdr.Close()

	// Add VMC INFO to the header.
	rdr.AddInfoToHeader("VMCGSID", "1", "String", "VMC Sequence identifier")
	rdr.AddInfoToHeader("VMCGLID", "1", "String", "VMC Location identifier")
	rdr.AddInfoToHeader("VMCGAID", "1", "String", "VMC Allele identifier")

	//create the new writer
	writer, err := vcfgo.NewWriter(output, rdr.Header)
	if err != nil {
		panic(err)
	}

	for {
		variant := rdr.Read()
		if variant == nil {
			break
		}

		// Check for alternate alleles	.
		altAllele := variant.Alt()
		if len(altAllele) > 1 {
			panic("multiallelic variant found, please pre-run with vt.")
		}

		record := vmc.VMCMarshal(variant, args.FastaFile, args.DATABASE, "VMC")
		variant.Info().Set("VMCGSID", vmc.SequenceID(record))
		variant.Info().Set("VMCGLID", vmc.LocationID(record))
		variant.Info().Set("VMCGAID", vmc.AlleleID(record))

		writer.WriteVariant(variant)
	}
}
