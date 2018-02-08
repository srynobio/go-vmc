package main

import (
	"database/sql"
	"fmt"

	"github.com/alexflint/go-arg"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shenwei356/bio/seqio/fastx"
	"github.com/srynobio/go-vmc/vmc"
)

func main() {

	var args struct {
		Fasta    string `arg:"required,help:Reference fasta file to create VMC_Sequence_ID record."`
		DATABASE string `arg:"required,help:Database file to build or add records to."`
	}
	arg.MustParse(&args)

	// open connection to db
	db, err := sql.Open("sqlite3", args.DATABASE)
	if err != nil {
		panic("cannot open connection to sqlite3 database.")
	}
	defer db.Close()

	// create needed table in database.
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `VMC_Reference_Sequence` (`ID` INTEGER PRIMARY KEY AUTOINCREMENT, `Description_Line` TEXT NOT NULL, `VMC_Sequence_ID` TEXT NOT NULL UNIQUE)")
	if err != nil {
		panic("Could not create needed database table.")
	}

	// create insert statement
	stmt, err := db.Prepare("INSERT OR IGNORE INTO VMC_Reference_Sequence(Description_Line, VMC_Sequence_ID) values(?,?)")
	if err != nil {
		panic(err)
	}

	// Incoming fastq file.
	reader, err := fastx.NewDefaultReader(args.Fasta)
	if err != nil {
		panic(err)
	}

	for chunk := range reader.ChunkChan(5000, 5) {
		if chunk.Err != nil {
			panic(chunk.Err)
		}

		for _, record := range chunk.Data {
			digestID := vmc.Digest(record.Seq.Seq, 24)
			_, err := stmt.Exec(record.Name, digestID)
			if err != nil {
				panic(err)
			}

			fmt.Println("Added to the Database:")
			fmt.Println(string(record.Name), digestID)
		}
	}

}
