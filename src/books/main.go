package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	ln, e := net.Listen("tcp", ":8001")
	checkError(e)
	defer ln.Close()
	fmt.Println("Listening at localhost:8001")
	e = os.MkdirAll("data", 0700)
	checkError(e)
	for {
		conn, e := ln.Accept()
		checkError(e)

		m, e := bufio.NewReader(conn).ReadString('\n')
		checkError(e)
		fmt.Printf("Recieving: %s\nmessage: %s", conn.RemoteAddr(), m)

		book := strings.Split(m[0:len(m)-1], ",")
		writeToBooks(book)
		conn.Write([]byte("Book saved to file!\n"))
		conn.Close()
	}
}

func writeToBooks(b []string) {
	books := getBooks()
	books = append(books, b)

	f, e := os.Create("data/books.csv")
	checkError(e)
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	w.WriteAll(books)
}

func getBooks() [][]string {
	f, e := os.Open("data/books.csv")
	if e != nil {
		log.Println("data/books.csv does not exist")
		return [][]string{}
	}
	r := csv.NewReader(f)
	res, e := r.ReadAll()
	checkError(e)
	return res
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
