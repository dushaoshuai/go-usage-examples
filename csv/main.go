package main

import (
	"encoding/csv"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

var (
	in = `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
# lines beginning with a # character are ignored
"Robert","Griesemer","gri"
`
	records = [][]string{
		{"first name", "last name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}
)

func main() {
	r := csv.NewReader(strings.NewReader(in))
	r.Comment = '#'

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			logrus.Fatal(err)
		}
		fmt.Println(record)
	}

	w := csv.NewWriter(os.Stdout)
	err := w.WriteAll(records)
	if err != nil {
		logrus.Fatal("error writing csv:", err)
	}
}
