package main

import (
	"fmt"

	"github.com/blevesearch/bleve"
)

type product struct {
	ID        int
	Name      string
	BrandName string
	// Brand     brand
}

// type brand struct {
// 	BrandName string
// }

func main() {
	// index := bleve.NewDocumentMapping()
	// brandMapping := bleve.NewTextFieldMapping()
	// brandMapping.Store

	// mapping := bleve.NewIndexMapping()
	// index, err := bleve.New("example.bleve", mapping)
	// if err != nil {
	// 	fmt.Println(err)
	// 	// return
	// }
	// p1 := &product{1, "eder", "teste"}
	// p2 := &product{2, "maria", "teste dois"}
	// p3 := &product{3, "josé", "teste"}
	// index.Index(strconv.Itoa(p1.ID), p1)
	// index.Index(strconv.Itoa(p2.ID), p2)
	// index.Index(strconv.Itoa(p3.ID), p3)

	// NOTE: keyword é o analyser pra marca (não quebrar o token em palavras)

	// 0,5,10,15,20,25,30,35,40,45,50,55 * * * *
	// https://godoc.org/gopkg.in/robfig/cron.v2

	index, err := bleve.Open("example.bleve")
	if err != nil {
		fmt.Println(err)
	}
	defer index.Close()

	// d, err := index.Document("2")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("=> %s\n", d.Fields[1].Value())

	// search for some text
	query := bleve.NewMatchQuery("teste")

	stylesFacet := bleve.NewFacetRequest("BrandName", 5)
	search := bleve.NewSearchRequest(query)
	search.AddFacet("Brands", stylesFacet)
	search.Fields = []string{"Name"}
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("=> %+v\n", searchResults)

	// fmt.Println(searchResults)

}
