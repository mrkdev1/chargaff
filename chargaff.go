package main

import (
	"os"
	"fmt"
	"strconv"	
	"strings"
	"log"
    "io/ioutil"

)

func main() {
    file := string(os.Args[1])
	
	
	// Read file to byte slice
    data, err := ioutil.ReadFile(file)
    if err != nil {
        log.Fatal(err)
    }
	
	str := string(data)

	counta := strings.Count(str, "A")
	countg := strings.Count(str, "G")
	countc := strings.Count(str, "C")
	countt := strings.Count(str, "T")
	length := len(str)

	fmt.Printf("%v bp\n", length)

	fmt.Printf("A: %.1f%%\n", 100.0*(float64(counta)/float64(length)))
	fmt.Printf("G: %.1f%%\n", 100.0*(float64(countg)/float64(length)))
	fmt.Printf("C: %.1f%%\n", 100.0*(float64(countc)/float64(length)))
	fmt.Printf("T: %.1f%%\n", 100.0*(float64(countt)/float64(length)))

	fmt.Printf("A/T: %.2f\n", (float64(counta) / float64(countt)))
	fmt.Printf("G/C: %.2f\n", (float64(countg) / float64(countc)))
	
	fmt.Printf("G+C: %.1f%%\n", 100.0*((float64(countg)+float64(countc))/float64(length)))	
	fmt.Printf("A+T: %.1f%%\n", 100.0*((float64(counta)+float64(countt))/float64(length)))

	a := 100.0*(float64(counta)/float64(length))
	g := 100.0*(float64(countg)/float64(length))
	c := 100.0*(float64(countc)/float64(length))
	t := 100.0*(float64(countt)/float64(length))
			
	fmt.Printf("Type: %T Value: %v\n", a, a)
	
	sa := strconv.FormatFloat(a, 'f', 2, 64)
	sg := strconv.FormatFloat(g, 'f', 2, 64)
	sc := strconv.FormatFloat(c, 'f', 2, 64)
	st := strconv.FormatFloat(t, 'f', 2, 64)

	sadt := strconv.FormatFloat((a/t), 'f', 2, 64)
	sgdc := strconv.FormatFloat((g/c), 'f', 2, 64)
	sgc := strconv.FormatFloat((g+c), 'f', 2, 64)
	sat := strconv.FormatFloat((a+t), 'f', 2, 64)	
	
	fmt.Printf("Type: %T Value: %v\n", sa, sa)

	head := ("| " + "A%" + " | " + "G%" +  " | " + "C%" +  " | " + "T%" +  " | " + "A/T" +  " | " + "G/C" +  " | " + "G+C" +  " | " + "A+T" +  " |")
	bar := ("| " + " --- " + " | " + " --- " +  " | " + " --- " +  " | " + " --- " +  " | " + " --- " +  " | " + " --- " +  " | " + " --- " +  " | " + " --- " +  " |")
    row := 	("| " + sa + " | " + sg +  " | " + sc +  " | " + st +  " | " + sadt +  " | " + sgdc +  " | " + sgc +  " | " + sat +  " |")
	
    err = ioutil.WriteFile("output.md", []byte(head + "\n" + bar + "\n" + row + "\n"), 0666)
    if err != nil {
        log.Fatal(err)
    }	
}
