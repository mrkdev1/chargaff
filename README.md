# chargaff *filepath* 

## Overview

The program **chargaff** calculates the percentages of the nucleobases (cytosine [C], guanine [G], adenine [A] and thymine [T]) in a DNA sequence. A FASTA input file is required. The results are displayed in the console and written into an output file named `table.md`.  

## `chargaff.go`

The program's source is written in the [Go](https://golang.org/) language. After installing Go and organizing a workspace environment as described in ["How to Write Go Code"](https://golang.org/doc/code.html) , you can create a new package directory named `chargaff` and clone the repository `https://github.com/mrkdev1/chargaff.git` into this directory. If you keep your source code in a GitHub repository, and you have a [GitHub](https://github.com/) account at github.com/*user*, and are using the default the GOPATH environmental variable, the location of the package on a Windows system would be:

`C:\Go\src\github.com\`*user*`\chargaff`

Then to build `chargaff.exe` move to the package directory and type:

`go build`    

## *sequence*`.fasta`

The DNA sequence input must be in [FASTA format](https://en.wikipedia.org/wiki/FASTA_format). FASTA format begins with a single-line description followed by lines of ascii characters representing the sequence of nucleobases .) The description line is distinguished from the sequence data by a greater-than (">") symbol at the beginning. All lines of text should be shorter than 80 characters in length. Blank lines are not allowed in the middle of FASTA input. Place this input file in the same directory as `chargaff.exe`. If you clone the github repository, you will also get a sample FASTA input file (5.4 KB), `phi.fasta`, representing the complete DNA genome of [Bacteriophage phiX174](http://pdb101.rcsb.org/motm/2) . 

## Command Syntax

The following command computes the proportions for the DNA sequence provided as a FASTA file located at the file path: *filepath*. If your input file is in the package directory then from that directory you can simply type the following to input, for example, the `phi.fasta` sequence.   

`chargaff phi.fasta`

## `table.md` 

On success, **chargaff** will display the result in the console and create a text file named `table.md` inside the same directory as `chargaff.exe`. Results get written into `table.md` as a table in github-flavored markdown format. For example, for the Bacteriophage phiX174 sequence `table.md` would contain markup text similar to the following. 

```
### >NC_001422.1 Enterobacteria phage phiX174 sensu lato, complete genome

5386bp

| A% | G% | C% | T% | A/T | G/C | G+C | A+T |
|  ---  |  ---  |  ---  |  ---  |  ---  |  ---  |  ---  |  ---  |
| 23.97 | 23.28 | 21.48 | 31.27 | 0.77 | 1.08 | 44.76 | 55.24 |
```


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