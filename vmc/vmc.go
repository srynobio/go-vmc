package vmc

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/brentp/vcfgo"
	//"github.com/shenwei356/bio/seqio/fastx"
	"time"
)

const Version = "v1.0.0"

type VMC struct {
	Id         string
	Identifier struct {
		accession string
		namespace string
	}
	Interval struct {
		start uint64
		end   uint64
	}
	Location struct {
		id          string
		interval    string
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
	Meta struct {
		generated_at time.Time
		vmc_version  string
	}
}

// ------------------------- //
// VMC functions
// ------------------------- //

func init() {

}

// TODO:
// method to build or get seq_id from file or db.

func Initialize() *VMC {
	vmc := VMC{}

	vmc.Meta.generated_at = time.Time
	vmc.Meta.vmc_version = Version

	return &vmc
}

// ------------------------------------------------------ //

/*
//func (v *VMC) VMCBuild(vcf *vcfgo.Variant, namespace string) *VMC {
func VMCBuild(vcf *vcfgo.Variant, namespace string) *VMC {

	v := VMC{}
	v.digestLocation(vcf, namespace)
	v.digestAllele(vcf, namespace)
	return &v
}
*/
// ------------------------------------------------------ //

func (v *VMC) DigestLocation(vcf *vcfgo.Variant, namespace string) *VMC {
	//func VMCLocation(v *VMC, vcf *vcfgo.Variant, namespace string) {
	seqID := "Ya6Rs7DHhDeg7YaOSg1EoNi3U_nQ9SvO"

	intervalString := fmt.Sprint(uint64(vcf.Start())) + ":" + fmt.Sprint(uint64(vcf.End()))
	location := "<Location:<Identifier:" + seqID + ">:<Interval:" + intervalString + ">>"
	DigestLocation := Digest([]byte(location), 24)

	v.Location.interval = intervalString
	v.Location.sequence_id = seqID
	v.Location.id = DigestLocation

	if namespace == "VMC" {
		identifier := namespace + ":GL_" + DigestLocation
		v.Location.id = identifier

	} else {

		identifier := namespace + ":" + DigestLocation
		v.Location.id = identifier
	}
	return v
}

// ------------------------------------------------------ //

func (v *VMC) DigestAllele(vcf *vcfgo.Variant, namespace string) *VMC {
	//func vmcAllele(v *VMC, vcf *vcfgo.Variant, namespace string) {

	state := fmt.Sprint(vcf.Alt())

	allele := "<Allele:<Identifier:" + v.Location.id + ">:" + state + ">"
	DigestAllele := Digest([]byte(allele), 24)

	v.Allele.location_id = v.Location.id
	v.Allele.state = state

	if namespace == "VMC" {
		identifier := namespace + ":GA_" + DigestAllele
		v.Allele.id = identifier
	} else {

		identifier := namespace + ":" + DigestAllele
		v.Allele.id = identifier
	}
	return v
}

// ------------------------------------------------------ //

func (v *VMC) AlleleID() {

}

// ------------------------------------------------------ //

func Digest(bv []byte, truncate int) string {
	hasher := sha512.New()
	hasher.Write(bv)

	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil)[:truncate])
	return sha
}

// ------------------------------------------------------ //
