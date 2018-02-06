## go-vmc
Go library and toolkit using the Variation Modelling Collaboration (VMC) Data model for exchanging sequence variation.

### This file will be updated on release.

* This codebase is still in development and will/can change based on changes to the [VMC Spec](https://docs.google.com/document/d/12E8WbQlvfZWk5NrxwLytmympPby6vsv60RxCeD5wc1E/edit)


#### gofastq example

```
$> gofastq-vmc-osx -database vmc.sequence.db -fasta data/chr19.fa.gz
$> gofastq-vmc-osx -database vmc.sequence.db -fasta data/NC_000019.10.fasta

```

```
$> sqlite3 vmc.sequence.db

sqlite> select * from VMC_Reference_Sequence;
1|chr19|XF-Zesay434DzyLqFkrrbQfXfiNSWzaL
3|1 dna:chromosome chromosome:GRCh37:1:1:249250621:1 REF|hDCv74x2LDrTU2wveDvXnqyBIuWrfDm8

```

#### Current sqlite3 schema for VMC_Reference_Sequence_

```
`VMC_Reference_Sequence` (
    `ID` INTEGER PRIMARY KEY AUTOINCREMENT,
    `Description_Line` TEXT NOT NULL,
    `VMC_Sequence_ID` TEXT NOT NULL UNIQUE
);
```

