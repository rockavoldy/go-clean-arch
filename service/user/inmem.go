package user

import (
	"go-clean-arch/entity"
	"strings"
	"time"
)

type inmem struct {
	m map[entity.ID]*entity.User
}

func newInmem() *inmem {
	var m = map[entity.ID]*entity.User{}
	return &inmem{
		m: m,
	}
}

// Create an user
func (r *inmem) Create(e *entity.User) (entity.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

// Get an user
func (r *inmem) Get(id entity.ID) (*entity.User, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}

	if r.m[id].DeletedAt.IsZero() == false {
		return nil, entity.ErrNotFound
	}

	return r.m[id], nil
}

// Get user lists
func (r *inmem) List() ([]*entity.User, error) {
	var ret []*entity.User
	for _, val := range r.m {
		if val.DeletedAt.IsZero() {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// Search user
func (r *inmem) Search(query string) ([]*entity.User, error) {
	var ret []*entity.User
	for _, val := range r.m {
		if strings.Contains(strings.ToLower(val.Name), strings.ToLower(query)) {
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

// Update an User
func (r *inmem) Update(e *entity.User) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}

	r.m[e.ID] = e

	return nil
}

// Delete an User
func (r *inmem) Delete(id entity.ID) error {
	_, err := r.Get(id)
	if err != nil {
		return entity.ErrNotFound
	}

	r.m[id].DeletedAt = time.Now()

	return nil
}

// Restore an Deleted User
func (r *inmem) Restore(id entity.ID) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}

	if r.m[id].DeletedAt.IsZero() == true {
		return entity.ErrCannotBeRestored
	}

	r.m[id].DeletedAt = time.Time{}

	return nil
}