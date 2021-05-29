package book

import (
	"go-clean-arch/entity"
	"strings"
	"time"
)

type inmem struct {
	m map[entity.ID]*entity.Book
}

func newInmem() *inmem {
	var m = map[entity.ID]*entity.Book{}

	return &inmem{
		m: m,
	}
}

// Create a Book
func (r *inmem) Create(e *entity.Book) (entity.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

// Get a Book
func (r *inmem) Get(id entity.ID) (*entity.Book, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}

	if !r.m[id].DeletedAt.IsZero() {
		return nil, entity.ErrNotFound
	}

	return r.m[id], nil
}

// Get Book lists
func (r *inmem) List() ([]*entity.Book, error) {
	var ret []*entity.Book
	if len(r.m) == 0 {
		return nil, entity.ErrNotFound
	}

	for _, val := range r.m {
		if val.DeletedAt.IsZero() {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// Search a Book by Title
func (r *inmem) Search(query string) ([]*entity.Book, error) {
	var ret []*entity.Book

	for _, val := range r.m {
		if strings.Contains(strings.ToLower(val.Title), strings.ToLower(query)) {
			if val.DeletedAt.IsZero() {
				ret = append(ret, val)
			}
		}
	}

	if len(ret) == 0 {
		return nil, entity.ErrNotFound
	}

	return ret, nil
}

// Update a Book
func (r *inmem) Update(e *entity.Book) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}

	r.m[e.ID] = e
	return nil
}

// Delete a Book
func (r *inmem) Delete(id entity.ID) error {
	_, err := r.Get(id)
	if err != nil {
		return err
	}

	r.m[id].DeletedAt = time.Now()

	return nil
}

// Get deleted Book
func (r *inmem) GetDeletedBook(id entity.ID) (*entity.Book, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}

	if r.m[id].DeletedAt.IsZero() {
		return nil, entity.ErrCannotBeRestored
	}

	return r.m[id], nil
}

// Restore a Deleted Book
func (r *inmem) Restore(id entity.ID) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}

	if r.m[id].DeletedAt.IsZero() {
		return entity.ErrCannotBeRestored
	}

	r.m[id].DeletedAt = time.Time{}

	return nil
}
