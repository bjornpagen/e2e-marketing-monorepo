package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

// a simple csv to json converter
// input csv has header row:
//   ID, First Name, Last Name, Company Email
// output json is a map of ID to objects containng the other fields:
//   {
//     "1": {
//       "first": "John",
//       "last": "Doe",
//       "email": "john.doe@example.com"
//     }
//   }

func main() {
	// usage: csv2json <input.csv> <output.json>
	if len(os.Args) != 3 {
		fmt.Println("usage: csv2json <input.csv> <output.json>")
		os.Exit(1)
	}
	in := os.Args[1]
	out := os.Args[2]

	// open input file
	f, err := os.Open(in)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// read input file
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// create output map
	m := make(map[string]map[string]string)

	// loop over records
	for i, record := range records {
		// skip header row
		if i == 0 {
			continue
		}

		// create object
		o := make(map[string]string)
		o["first"] = record[1]
		o["last"] = record[2]
		o["email"] = record[3]

		// add to map
		m[record[0]] = o
	}

	// open output file
	f, err = os.Create(out)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// write output file
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	f.Write(b)
}
