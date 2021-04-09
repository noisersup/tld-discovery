package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func main(){
	if len(os.Args) < 2 {
		log.Fatalf("Provide target as argument ")
	}
	tlds, err := getLines("tlds-alpha-by-domain.txt") ; 
	if err != nil {
		log.Fatal(err)
	}

	length := len(tlds)
	
	var succeedTlds []string
	for i,tld := range tlds {
		if err := ping(os.Args[1]+"."+tld,80); err == nil {
			succeedTlds = append(succeedTlds, tld)
		}
		
		fmt.Print("\033[H\033[2J")
		for _,sTld := range succeedTlds {
			fmt.Printf("%s.%s\n",os.Args[1],sTld)
		}
		fmt.Printf("%d/%d scanned TLDs\n",i,length)
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

func ping(uri string,port int) (error){
	if _,err := net.DialTimeout("tcp",uri+":"+strconv.Itoa(port),1 * time.Second) ; err != nil { return err}
	return nil
}