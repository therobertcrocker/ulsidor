package main

import (
	"encoding/json"
	"fmt"

	"go.etcd.io/bbolt"
)

type Entity interface {
	ID() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ID() string {
	return fmt.Sprintf("person_%s", p.Name)
}

type Book struct {
	Title  string
	Author string
}

func (b *Book) ID() string {
	return fmt.Sprintf("book_%s", b.Title)
}

type PersonCodex struct {
	People map[string]*Person
	Log    []string
}

type BookCodex struct {
	Books map[string]*Book
	Log   []string
}

func main() {

	db, err := bbolt.Open("test.db", 0600, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// create a person codex
	personCodex := &PersonCodex{
		People: map[string]*Person{},
		Log:    []string{},
	}

	// create a book codex
	bookCodex := &BookCodex{
		Books: map[string]*Book{},
		Log:   []string{},
	}

	// add a few people to the codex
	personCodex.People["person_1"] = &Person{
		Name: "Bob",
		Age:  42,
	}

	//log
	personCodex.Log = append(personCodex.Log, "Added person_1")

	personCodex.People["person_2"] = &Person{
		Name: "Alice",
		Age:  32,
	}

	//log
	personCodex.Log = append(personCodex.Log, "Added person_2")

	// add a few books to the codex
	bookCodex.Books["book_1"] = &Book{
		Title:  "The Hobbit",
		Author: "J.R.R. Tolkien",
	}

	//log
	bookCodex.Log = append(bookCodex.Log, "Added book_1")

	bookCodex.Books["book_2"] = &Book{
		Title:  "The Fellowship of the Ring",
		Author: "J.R.R. Tolkien",
	}

	//log
	bookCodex.Log = append(bookCodex.Log, "Added book_2")

	// save the codices
	err = db.Update(func(tx *bbolt.Tx) error {

		// create a bucket for the codices
		codices, err := tx.CreateBucketIfNotExists([]byte("codices"))
		if err != nil {
			return err
		}

		// serialize the person codex
		personCodexBytes, err := json.Marshal(personCodex)
		if err != nil {
			return err
		}

		// serialize the book codex
		bookCodexBytes, err := json.Marshal(bookCodex)
		if err != nil {
			return err
		}

		// save the codices
		err = codices.Put([]byte("person_codex"), personCodexBytes)
		if err != nil {
			return err
		}

		err = codices.Put([]byte("book_codex"), bookCodexBytes)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

}
