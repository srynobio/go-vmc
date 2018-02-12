package vmc

import (
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"path"

	"github.com/brentp/vcfgo"
	_ "github.com/mattn/go-sqlite3"
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

func VMCMarshal(v *vcfgo.Variant, fastaFile, database, namespace string) *VMC {
	vmc := &VMC{}

	// search for and set seq id first.
	vmcID, err := vmc.getSeqID(v, fastaFile, database)
	if err != nil {
		fmt.Println(err)
	}
	vmc.Location.sequence_id = vmcID

	vmc.LocationDigest(v, "VMC")
	vmc.AlleleDigest(v, "VMC")

	return vmc
}

// ------------------------------------------------------ //

func (v *VMC) getSeqID(vcf *vcfgo.Variant, fastaFile, database string) (string, error) {

	// get basename of fastqFile for search
	fastaBase := path.Base(fastaFile)
	dataPathBase := path.Base(database)

	// open connection to db
	db, err := sql.Open("sqlite3", dataPathBase)
	if err != nil {
		panic("cannot open connection to sqlite3 database.")
	}
	defer db.Close()

	/// current location where need to handle nul returns from db.

	// fetch from db
	rows, err := db.Query("SELECT VMC_Sequence_ID FROM VMC_Reference_Sequence WHERE File_Name = (?) AND Chromosome = (?)", fastaBase, vcf.Chromosome)
	if err != nil {
		return "", errors.New("shit when down.")
		//		log.Panicf("Chromosome %s not found in fasta file %s\n", vcf.Chromosome, fastaBase)
	}
	defer rows.Close()

	var VMC_Sequence_ID string
	for rows.Next() {
		err := rows.Scan(&VMC_Sequence_ID)
		if err != nil {
			panic("Fatal error retrieving data from database.")
		}
	}
	return VMC_Sequence_ID, nil
}

// ------------------------------------------------------ //

///////// update to use interbased coordinate system.

func (v *VMC) LocationDigest(vcf *vcfgo.Variant, namespace string) {

	seqID := v.Location.sequence_id

	intervalString := fmt.Sprint(uint64(vcf.Start())) + ":" + fmt.Sprint(uint64(vcf.End()))
	location := "<Location:<Identifier:" + seqID + ">:<Interval:" + intervalString + ">>"
	DigestLocation := Digest([]byte(location), 24)

	v.Location.interval = intervalString
	v.Location.sequence_id = seqID
	v.Location.id = DigestLocation

	switch namespace {
	case "VMC":
		identifier := namespace + ":GL_" + DigestLocation
		v.Location.id = identifier
	default:
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

func SequenceID(vmc *VMC) string {
	return vmc.Location.sequence_id
}

// ------------------------------------------------------ //
