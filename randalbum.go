package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"math/rand"
	"os"
)

// Album holds album data
type Album struct {
	Title  string
	Artist string
}

var albums []Album

func getAlbum() Album {
	n := rand.Intn(len(albums))
	return albums[n]
}

func init() {
	file, err := os.Open("./assets/docs/albums.csv")
	if err != nil {
		log.Println("randalbum: init(): Could not open csv file.")
	}
	reader := csv.NewReader(bufio.NewReader(file))
	for {
		line, errR := reader.Read()
		if errR == io.EOF {
			break
		} else if errR != nil {
			log.Println("randalbum: init(): Not nil error on csv file read")
		}
		albums = append(albums, Album{
			Title:  line[0],
			Artist: line[1],
		})
	}
}
