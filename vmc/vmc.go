package vmc

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/brentp/vcfgo"
)

const Version = "v1.0.0"

type Identifier struct {
	accession string
	namespace string
}

type Interval struct {
	start uint64
	end   uint64
}
type Location struct {
	id          string
	interval    string
	sequence_id string
}
type Allele struct {
	id          string
	location_id string
	state       string
}

type VMC struct {
	Id         string
	Identifier Identifier
	Interval   Interval
	Location   Location
	Allele     Allele
}

// ------------------------- //
// VMC functions
// ------------------------- //

func VMCMarshal(v *vcfgo.Variant, namespace string) *VMC {
	vmc := &VMC{}

	vmc.LocationDigest(v, "VMC")
	vmc.AlleleDigest(v, "VMC")

	return vmc
}

// ------------------------------------------------------ //

func (v *VMC) SequenceDigest() {}

// ------------------------------------------------------ //

///////// update to use interbased coordinate system.

func (v *VMC) LocationDigest(vcf *vcfgo.Variant, namespace string) {

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
}

// ------------------------------------------------------ //

func (v *VMC) AlleleDigest(vcf *vcfgo.Variant, namespace string) {

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
}

// ------------------------------------------------------ //

func Digest(bv []byte, truncate int) string {
	hasher := sha512.New()
	hasher.Write(bv)

	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil)[:truncate])
	return sha
}

// ------------------------------------------------------ //

func LocationID(vmc *VMC) string {
	return vmc.Location.id
}

// ------------------------------------------------------ //

func AlleleID(vmc *VMC) string {
	return vmc.Allele.id
}

// ------------------------------------------------------ //
