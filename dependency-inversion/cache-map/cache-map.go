package cachemap

import "dependencyinversion/domain"

type CacheMap struct {
	data map[string]*domain.Book
}

func NewCacheMap() *CacheMap {
	return &CacheMap{
		data: make(map[string]*domain.Book),
	}
}

func (c *CacheMap) Write(book *domain.Book) {
	c.data[book.Name] = book
}

func (c *CacheMap) Read(name string) *domain.Book {
	return c.data[name]
}
