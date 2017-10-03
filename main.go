package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/blevesearch/bleve"
)

const datastorePath = "example.bleve"

// NOTE: keyword é o analyser pra marca (não quebrar o token em palavras)
// 0,5,10,15,20,25,30,35,40,45,50,55 * * * *
// https://godoc.org/gopkg.in/robfig/cron.v2

type product struct {
	ID        int
	Name      string
	BrandName string
}

func main() {

	index, err := openIndex()
	if err != nil {
		if err == bleve.ErrorIndexPathDoesNotExist {
			in, err := createIndex()
			if err != nil {
				panic(err)
			}
			index = in
		}
	}

	query := bleve.NewMatchQuery("teste")
	search := bleve.NewSearchRequest(query)

	stylesFacet := bleve.NewFacetRequest("BrandName", 5)
	search.AddFacet("Brands", stylesFacet)
	search.Fields = []string{"Name", "BrandName"}

	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(searchResults)

}

func openIndex() (bleve.Index, error) {
	index, err := bleve.Open(datastorePath)
	if err != nil {
		return nil, err
	}
	return index, nil
}

func createIndex() (bleve.Index, error) {
	// index := bleve.NewDocumentMapping()
	// brandMapping := bleve.NewTextFieldMapping()

	mapping := bleve.NewIndexMapping()
	index, err := bleve.New(datastorePath, mapping)
	if err != nil {
		return nil, err
	}
	p1 := &product{1, "eder", "teste"}
	p2 := &product{2, "maria", "teste dois"}
	p3 := &product{3, "josé", "teste"}
	index.Index(strconv.Itoa(p1.ID), p1)
	index.Index(strconv.Itoa(p2.ID), p2)
	index.Index(strconv.Itoa(p3.ID), p3)

	return index, nil
}
