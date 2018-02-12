## go-vmc
Go library and toolkit using the Variation Modelling Collaboration (VMC) Data model for exchanging sequence variation.

### This file will be updated on release.

* This codebase is still in development and will/can change based on changes to the [VMC Spec](https://docs.google.com/document/d/12E8WbQlvfZWk5NrxwLytmympPby6vsv60RxCeD5wc1E/edit)


### gofastq example

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

#### Current sqlite3 schema for VMC_Reference_Sequence

```
`VMC_Reference_Sequence` 
	`ID` INTEGER PRIMARY KEY AUTOINCREMENT, 
	`Chromosome` TEXT NOT NULL UNIQUE, 
	`Description_Line` TEXT NOT NULL UNIQUE, 
	`VMC_Sequence_ID` TEXT NOT NULL
);
```


#### Processing test runtimes.

* Test using GRCh37 primary assembly

```
$> gofastq-vmc-osx -database vmc.sequence.db -fasta data/Homo_sapiens.GRCh37.dna.primary_assembly.fa.gz
real	13m16.808s

$> sqlite3 vmc.sequence.db

sqlite> select * from VMC_Reference_Sequence;
1|1 dna:chromosome chromosome:GRCh37:1:1:249250621:1 REF|S_KjnFVz-FE7M0W6yoaUDgYxLPc1jyWU
2|10 dna:chromosome chromosome:GRCh37:10:1:135534747:1 REF|-BOZ8Esn8J88qDwNiSEwUr5425UXdiGX
3|11 dna:chromosome chromosome:GRCh37:11:1:135006516:1 REF|XXi2_O1ly-CCOi3HP5TypAw7LtC6niFG
4|12 dna:chromosome chromosome:GRCh37:12:1:133851895:1 REF|105bBysLoDFQHhajooTAUyUkNiZ8LJEH
5|13 dna:chromosome chromosome:GRCh37:13:1:115169878:1 REF|Ewb9qlgTqN6e_XQiRVYpoUfZJHXeiUfH
6|14 dna:chromosome chromosome:GRCh37:14:1:107349540:1 REF|5Ji6FGEKfejK1U6BMScqrdKJK8GqmIGf
7|15 dna:chromosome chromosome:GRCh37:15:1:102531392:1 REF|zIMZb3Ft7RdWa5XYq0PxIlezLY2ccCgt
8|16 dna:chromosome chromosome:GRCh37:16:1:90354753:1 REF|W6wLoIFOn4G7cjopxPxYNk2lcEqhLQFb
9|17 dna:chromosome chromosome:GRCh37:17:1:81195210:1 REF|AjWXsI7AkTK35XW9pgd3UbjpC3MAevlz
10|18 dna:chromosome chromosome:GRCh37:18:1:78077248:1 REF|BTj4BDaaHYoPhD3oY2GdwC_l0uqZ92UD
11|19 dna:chromosome chromosome:GRCh37:19:1:59128983:1 REF|ItRDD47aMoioDCNW_occY5fWKZBKlxCX
12|2 dna:chromosome chromosome:GRCh37:2:1:243199373:1 REF|9KdcA9ZpY1Cpvxvg8bMSLYDUpsX6GDLO
13|20 dna:chromosome chromosome:GRCh37:20:1:63025520:1 REF|iy_UbUrvECxFRX5LPTH_KPojdlT7BKsf
14|21 dna:chromosome chromosome:GRCh37:21:1:48129895:1 REF|LpTaNW-hwuY_yARP0rtarCnpCQLkgVCg
15|22 dna:chromosome chromosome:GRCh37:22:1:51304566:1 REF|XOgHwwR3Upfp5sZYk6ZKzvV25a4RBVu8
16|3 dna:chromosome chromosome:GRCh37:3:1:198022430:1 REF|4NK8W1wiBx1e2iknrBRL7xJgTSPnZBPs
17|4 dna:chromosome chromosome:GRCh37:4:1:191154276:1 REF|iy7Zfceb5_VGtTQzJ-v5JpPbpeifHD_V
18|5 dna:chromosome chromosome:GRCh37:5:1:180915260:1 REF|vbjOdMfHJvTjK_nqvFvpaSKhZillW0SX
19|6 dna:chromosome chromosome:GRCh37:6:1:171115067:1 REF|KqaUhJMW3CDjhoVtBetdEKT1n6hM-7Ek
20|7 dna:chromosome chromosome:GRCh37:7:1:159138663:1 REF|IW78mgV5Cqf6M24hy52hPjyyo5tCCd86
21|8 dna:chromosome chromosome:GRCh37:8:1:146364022:1 REF|tTm7wmhz0G4lpt8wPspcNkAD_qiminj6
22|9 dna:chromosome chromosome:GRCh37:9:1:141213431:1 REF|HBckYGQ4wYG9APHLpjoQ9UUe9v7NxExt
23|MT dna:chromosome chromosome:GRCh37:MT:1:16569:1 REF|k3grVkjY-hoWcCUojHw6VU6GE3MZ8Sct
24|X dna:chromosome chromosome:GRCh37:X:1:155270560:1 REF|v7noePfnNpK8ghYXEqZ9NukMXW7YeNsm
25|Y dna:chromosome chromosome:GRCh37:Y:2649521:59034049:1 REF|fbS5kAwZUB5-1xVpa7xZ4s_lyDpLPVUo

```






