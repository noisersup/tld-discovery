package main

import (
	"bufio"
	"log"
	"os"
)

func main(){
	if tlds, err := getLines("tlds-alpha-by-domain.txt") ; err != nil {
		log.Fatal(err)
	}
}

func getLines(path string) ([]string, error){
	file, err := os.Open(path)
	if err != nil {return nil, err}

	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)

	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	if err = file.Close() ; err != nil {
		return nil, err
	}
	return lines,nil
}