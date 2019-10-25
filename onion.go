package main

import (
	"encoding/csv"
	"os"
)

//Onions holds a list for addresses
type Onions struct {
	list []string
}

//NewOnions reads and parse the CSV file
func NewOnions(csvFile string) (*Onions, error) {
	fd, err := os.Open(csvFile)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(fd)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	onions := new(Onions)
	for _, record := range records {
		if len(record[0]) > 0 {
			onions.list = append(onions.list, record[0])
		}
	}
	return onions, nil
}

//Links returns array of onion URL's
func (o Onions) Links() []string {
	return o.list
}
