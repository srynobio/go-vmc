package main

import (
	"github.com/alexflint/go-arg"
	"github.com/brentp/vcfgo"
	"github.com/brentp/xopen"
	"github.com/srynobio/go-vmc/vmc"
	"os"
)

func main() {

	var args struct {
		VCF    string `arg:"required",  "VCF file to annotate with VMC digest ids"`
		Fasta  string `arg: "Reference fasta file used to create VCF file."`
		Output string `arg:"required", Name for output VCF file."`
	}
	arg.MustParse(&args)

	fh, err := xopen.Ropen(args.VCF)
	if err != nil {
		panic("VCF file not given or could not be opened.")
	}
	defer fh.Close()

	// VCF reader
	rdr, err := vcfgo.NewReader(fh, false)
	if err != nil {
		panic("Could not read given VCF file.")
	}
	defer rdr.Close()

	// create the writer
	output, err := os.Create(args.Output)
	if err != nil {
		panic(err)
	}
	defer output.Close()

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

		record := vmc.VMCMarshal(variant, "VMC")
		variant.Info().Set("VMCGSID", vmc.SequenceID(record))
		variant.Info().Set("VMCGLID", vmc.LocationID(record))
		variant.Info().Set("VMCGAID", vmc.AlleleID(record))

		writer.WriteVariant(variant)
	}
}

// ------------------------- //
