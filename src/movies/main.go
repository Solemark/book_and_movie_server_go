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
	ln, e := net.Listen("tcp", ":8002")
	checkError(e)
	defer ln.Close()
	fmt.Println("Listening at localhost:8002")
	e = os.MkdirAll("data", 0700)
	checkError(e)
	for {
		conn, e := ln.Accept()
		checkError(e)

		m, e := bufio.NewReader(conn).ReadString('\n')
		checkError(e)
		fmt.Printf("Recieving: %s\nmessage: %s", conn.RemoteAddr(), m)

		movie := strings.Split(m[0:len(m)-1], ",")
		writeToMovies(movie)
		conn.Write([]byte("Movie saved to file!\n"))
		conn.Close()
	}
}

func writeToMovies(m []string) {
	movies := getMovies()
	movies = append(movies, m)

	f, e := os.Create("data/movies.csv")
	checkError(e)
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	w.WriteAll(movies)
}

func getMovies() [][]string {
	f, e := os.Open("data/movies.csv")
	if e != nil {
		log.Println("data/movies.csv does not exist")
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
