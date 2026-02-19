package cacheslice

import "dependencyinversion/domain"

type CacheSlice struct {
	data []*domain.Book
}

func NewCacheSlice() *CacheSlice {
	return &CacheSlice{
		data: make([]*domain.Book, 0),
	}
}

func (c *CacheSlice) Write(book *domain.Book) {
	c.data = append(c.data, book)
}

func (c *CacheSlice) Read(name string) *domain.Book {
	for _, book := range c.data {
		if book.Name == name {
			return book
		}
	}
	return nil
}
