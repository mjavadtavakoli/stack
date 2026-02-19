package main

import (
	cachemap "dependencyinversion/cache-map"
	cacheslice "dependencyinversion/cache-slice"
	"dependencyinversion/domain"
	"fmt"
)

type Store interface {
	Write(*domain.Book)
	Read(string) *domain.Book
}

func storeBook(storage Store, book *domain.Book) {
	storage.Write(book)
}

func readBook(storage Store, name string) *domain.Book {
	return storage.Read(name)
}

func main() {
	writeSlice()
	writeMap()
}

func writeSlice() {
	microservice := domain.Book{
		Name:   "microservice",
		Writer: "ali",
	}
	guidCode := domain.Book{
		Name:   "guid code",
		Writer: "asghar",
	}
	cacheSlice := cacheslice.NewCacheSlice()

	storeBook(cacheSlice, &microservice)
	storeBook(cacheSlice, &guidCode)

	expectMicroservice := readBook(cacheSlice, microservice.Name)

	fmt.Println(expectMicroservice)
}

func writeMap() {
	microservice := domain.Book{
		Name:   "microservice",
		Writer: "ali",
	}
	guidCode := domain.Book{
		Name:   "guid code",
		Writer: "asghar",
	}
	cacheSlice := cachemap.NewCacheMap()

	storeBook(cacheSlice, &microservice)
	storeBook(cacheSlice, &guidCode)

	expectMicroservice := readBook(cacheSlice, microservice.Name)

	fmt.Println(expectMicroservice)
}
