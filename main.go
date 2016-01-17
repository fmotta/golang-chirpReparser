package main

/* The source CSV file contains a distance from a given location embedded in the Comments field (a parameter with an default of 13)
	The distance (in miles) is represented as the first bit of the comment like (1mi WNW).

	So, I load the records and parse out the '1' (as shown in the above example) as a distance and use it for filtering against 'max' variable.
	These filtered records are output to stdout along with the header record

File format example:
Location,Name,Frequency,Duplex,Offset,Tone,rToneFreq,cToneFreq,DtcsCode,DtcsPolarity,Mode,TStep,Skip,Comment,URCALL,RPT1CALL,RPT2CALL
1,K6HN,440.075000,+,5.000000,Tone,114.8,88.5,023,NN,FM,5.00,,(19mi WNW) ,,,,
2,WB6ECE,441.300000,+,5.000000,Tone,131.8,88.5,023,NN,FM,5.00,,(18mi SSW) ,,,,
3,K9GVF,441.025000,+,5.000000,Tone,156.7,88.5,023,NN,FM,5.00,,(5mi SSE) ,,,,
4,W6AMT,440.125000,+,5.000000,Tone,114.8,88.5,023,NN,FM,5.00,,(5mi SW) ,,,,

*/

 import (
	"encoding/csv"
	"fmt"
	"os"
	 "strings"
	"strconv"
	"flag"
 )

var distance = flag.Int("distance", 999, "max distance away")
var csvfile = flag.String("infile", "radio-near-home.csv", "CSV file to read")
var help = flag.Bool("help", false, "help")
var commentFld = flag.Int("comment", 13, "comment field number")
var maxFreq = flag.Int("maxfreq", 440, "max frequency to extract")
var minFreq = flag.Int("minfreq", 100, "min frequency to extract")

var freqField = 2;
var max = 8;

var Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
        flag.PrintDefaults()
	fmt.Printf("Ignore the distance value for now.  I overwrite it with %d in code\n", max)
}

 func main() {


//	var commentField = *commentFld

	flag.Parse()

	if (*help == true) {
		Usage()
		os.Exit(0)
	}


	csvfile, err := os.Open(*csvfile)

	if err != nil {
		fmt.Println(err)
		return
	}


	if (*distance != 999) {  // distance flag was used - so overwrite default
		max = *distance
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	reader.FieldsPerRecord = -1 // see the Reader struct information below

	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}



	var cnt = 0;		// field counter to separate the field header
	var matchcnt = 1;	// counter to put into 'Location' field replacing the one in the input file with an enumerator
	for _, each := range rawCSVdata {
		if (cnt > 0 ) { // do not process the header
			// Parse out the distance from the Comment field (default to 13th)
			// yep this can be done easier... I am still playing around
			distfirst :=  strings.SplitAfter(each[*commentFld], ")")
			distsecond :=  strings.SplitAfter(distfirst[0], ")")
			dist3 := strings.Trim(distsecond[0], ")")
			dist4 := strings.Trim(dist3, "(")
			dist5 := strings.Trim(dist4, "(")
			dist6 := strings.SplitAfter(dist5, "mi")
			dist7 := strings.Trim(dist6[0], "mi")
			dist, err := strconv.Atoi(dist7)
			freq, err := strconv.ParseFloat(each[freqField], 64)
			if err == nil {		// seems backward
				if (dist <= max && (freq <= float64(*maxFreq) && freq >= float64(*minFreq))) {
					for i := 0; i < len(each); i++ {
						if (i == (len(each) -1)) {
							fmt.Println(each[i])
						} else {
							if (i == 0) {
								fmt.Printf("%d,", matchcnt)
							} else {
								fmt.Printf("%s,", each[i])
							}
						}
					}
					matchcnt++
				}
			}
		} else {
			// Print the header
			for i := 0; i < len(each); i++ {
				if (i == (len(each) -1)) {
					fmt.Println(each[i])
				} else {
					fmt.Printf("%s,", each[i])
				}
			}
		}
		cnt++
         }
 }

