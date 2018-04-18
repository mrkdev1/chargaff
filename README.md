# chargaff *path* 

## Overview

The program **chargaff** calculates the percentages of the nucleobases (cytosine [C], guanine [G], adenine [A] or thymine [T])in a DNA sequence and generates a table of the results in markdown format.  

## `chargaff.go`

The program's source written in [Go](https://golang.org/) . 

## *sequence*`.fasta`

The input DNA sequence provided as a text file in [FASTA format](https://en.wikipedia.org/wiki/FASTA_format) . 

## Use

`chargaff some.fasta`

## `table.md` 

An output file containing the result as a table in github-flavored markdown format. 

# Examples

### >NC_001422.1 Enterobacteria phage phiX174 sensu lato, complete genome

5386bp

| A% | G% | C% | T% | A/T | G/C | G+C | A+T |
|  ---  |  ---  |  ---  |  ---  |  ---  |  ---  |  ---  |  ---  |
| 23.97 | 23.28 | 21.48 | 31.27 | 0.77 | 1.08 | 44.76 | 55.24 |


### >HG530134.1 Genlisea margaretae chloroplast, complete genome

141252bp

| A% | G% | C% | T% | A/T | G/C | G+C | A+T |
|  ---  |  ---  |  ---  |  ---  |  ---  |  ---  |  ---  |  ---  |
| 30.18 | 19.02 | 19.32 | 31.48 | 0.96 | 0.98 | 38.34 | 61.66 |


### >E. coli K-12 MG1655 U00096.2 (1 to 4639675 = 4639675 bp)

4639675bp

| A% | G% | C% | T% | A/T | G/C | G+C | A+T |
|  ---  |  ---  |  ---  |  ---  |  ---  |  ---  |  ---  |  ---  |
| 24.62 | 25.37 | 25.42 | 24.59 | 1.00 | 1.00 | 50.79 | 49.21 |

### >FR847113.2 Caenorhabditis briggsae complete genome, chrII, whole genome shotgun sequence

16060613bp

| A% | G% | C% | T% | A/T | G/C | G+C | A+T |
|  ---  |  ---  |  ---  |  ---  |  ---  |  ---  |  ---  |  ---  |
| 31.22 | 18.68 | 18.72 | 31.38 | 1.00 | 1.00 | 37.40 | 62.60 |

### >J01415.2 Homo sapiens mitochondrion, complete genome

16568bp

| A% | G% | C% | T% | A/T | G/C | G+C | A+T |
|  ---  |  ---  |  ---  |  ---  |  ---  |  ---  |  ---  |  ---  |
| 30.93 | 13.09 | 31.27 | 24.71 | 1.25 | 0.42 | 44.36 | 55.64 |

### >EF064321.1 Homo sapiens isolate 5_U6a1(Tor270) mitochondrion, complete genome

16569bp

| A% | G% | C% | T% | A/T | G/C | G+C | A+T |
|  ---  |  ---  |  ---  |  ---  |  ---  |  ---  |  ---  |  ---  |
| 30.86 | 13.16 | 31.28 | 24.70 | 1.25 | 0.42 | 44.43 | 55.57 |