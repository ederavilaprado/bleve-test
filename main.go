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
	ID   int
	Name string `json:"name"`
	Type string
	Tags []string
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

	query := bleve.NewMatchQuery("Flats")
	query.SetField("name")

	search := bleve.NewSearchRequest(query)
	search.Fields = []string{"Name", "Type", "Tags"}

	// Facet
	search.AddFacet("Types", bleve.NewFacetRequest("Type", 5))

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
	indexMapping := bleve.NewIndexMapping()

	productMapping := bleve.NewDocumentMapping()

	nameFieldMapping := bleve.NewTextFieldMapping()
	nameFieldMapping.Analyzer = "en"
	productMapping.AddFieldMappingsAt("name", nameFieldMapping)

	// Add product mapping to indexMaping
	indexMapping.AddDocumentMapping("product", productMapping)

	index, err := bleve.New(datastorePath, indexMapping)
	if err != nil {
		return nil, err
	}

	p1 := &product{1, "180 flat", "fotolivro", []string{"best", "new"}}
	p2 := &product{
		ID:   2,
		Name: "Prime",
		Type: "fotolivro",
	}
	p3 := &product{3, "Revista", "fotolivreto", []string{"new"}}

	index.Index(strconv.Itoa(p1.ID), p1)
	index.Index(strconv.Itoa(p2.ID), p2)
	index.Index(strconv.Itoa(p3.ID), p3)

	return index, nil
}
