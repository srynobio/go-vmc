package vmc

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/brentp/vcfgo"
	"time"
)

const Version = "v1.0.0"

// Define the VMC struct

type VMCID struct {
	Version    string
	Id         string
	DateTime   time.Time
	Identifier struct {
		accession string
		namespace string
	}
	Interval struct {
		start uint32
		end   uint32
	}
	Location struct {
		id          string
		interval    int
		sequence_id string
	}
	Allele struct {
		id          string
		location_id string
		state       string
	}
	Genotype struct {
		id            string
		haplotype_ids []string
		completedness int
	}
	Haplotype struct {
		id            string
		allele_id     []string
		completedness int
	}
	//SequenceID  string
	//LocationID  string
	//AlleleID    string
	HaplotypeID string
	GenotypeID  string
}

// ------------------------- //
// VMC functions
// ------------------------- //

// TODO:
// method to build or get seq_id from file or db.

func CreateVMC(v *vcfgo.Variant) *VMCID {

	vmc := VMCID{}
	vmc.Version = Version
	vmc.DateTime = time.Now()
	vmc.Identifier.namespace = "VMC"

	//set Dummy Seq ID
	vmc.Location.sequence_id = vmc.Identifier.namespace + ":GS_Ya6Rs7DHhDeg7YaOSg1EoNi3U_nQ9SvO"

	vmc.Interval.start = v.Start() - 1
	vmc.Interval.end = v.End() - 1

	vmcLocationID(vmc)

	vmc.Allele.state = v.Alt()[0]
	vmc.Allele.location_id = vmc.LocationID
	vmc.alleleid = vmcalleleid(vmc)

	var gogo = &vmc
	return gogo

}

// ------------------------- //

func vmcLocationID(v VMCID) {

	seqID := v.Location.sequence_id
	intervalString := fmt.Sprint(v.Interval.start) + ":" + fmt.Sprint(v.Interval.end)

	location := "<Location:<Identifier:" + seqID + ">:<Interval:" + intervalString + ">>"
	DigestLocation := vmcDigestId([]byte(location), 24)
	id := v.Identifier.namespace + ":GL_" + DigestLocation

	v.Location.id = id
}

// ------------------------- //

func vmcAlleleID(v VMCID) {

	vmcLocation := v.Location.id
	state := v.Allele.state

	allele := "<Allele:<Identifier:" + vmcLocation + ">:" + state + ">"
	DigestAllele := vmcDigestId([]byte(allele), 24)
	id := v.Identifier.namespace + ":GA_" + DigestAllele

	v.Allele.id = id

	v.Allele.id = id

}

// ------------------------- //

func vmcDigestId(bv []byte, truncate int) string {
	hasher := sha512.New()
	hasher.Write(bv)

	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil)[:truncate])
	return sha
}

// ------------------------- //
