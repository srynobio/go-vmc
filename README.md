# go-vmc
Go library and toolkit using the Variation Modelling Collaboration (VMC) Data model for exchanging sequence variation.



# notes on how different version of the reference will create 
# different version of vmc sequence_id
// http://hgdownload.soe.ucsc.edu/goldenPath/hg19/chromosomes/
// data/chr19.fa.gz
// >chr19
VMC:GS_XF-Zesay434DzyLqFkrrbQfXfiNSWzaL


// ftp://ftp.ensembl.org/pub/grch37/current/fasta/homo_sapiens/dna/
// Homo_sapiens.GRCh37.dna.chromosome.19.fa.gz
// >19 dna:chromosome chromosome:GRCh37:19:1:59128983:1 REF
VMC:GS_ItRDD47aMoioDCNW_occY5fWKZBKlxCX


// Example use in VMC doc:
//Homo sapiens chromosome 19, GRCh38.p7 Primary Assembly
// https://www.ncbi.nlm.nih.gov/nucleotide/NC_000019.10
// >NC_000019.10 Homo sapiens chromosome 19, GRCh38.p7 Primary Assembly
VMC:GS_IIB53T8CNeJJdUqzn9V_JnRtQadwWCbl


