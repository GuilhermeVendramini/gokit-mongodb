package profilesvc

import (
	"context"
	"errors"
	"sync"

	"github.com/GuilhermeVendramini/gokit-mongodb/config"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Service is a simple CRUD interface for user profiles.
type Service interface {
	PostProfile(ctx context.Context, p Profile) error
	GetProfile(ctx context.Context, id string) (Profile, error)
	PatchProfile(ctx context.Context, id string, p Profile) error
	DeleteProfile(ctx context.Context, id string) error
}

// Profile represents a single user profile.
// ID should be globally unique.
type Profile struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

type inmomService struct {
	mtx sync.RWMutex
	c   *mgo.Collection
}

func NewInmomService() Service {
	return &inmomService{
		c: config.Profiles,
	}
}

func (s *inmomService) PostProfile(ctx context.Context, p Profile) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	// TODO - verify if profile already exists. Use ErrAlreadyExists error
	err := s.c.Insert(p)
	if err != nil {
		return errors.New("internal server error" + err.Error())
	}
	return nil
}

func (s *inmomService) GetProfile(ctx context.Context, id string) (Profile, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	p := Profile{}

	err := s.c.Find(bson.M{"id": id}).One(&p)
	return p, err
}

func (s *inmomService) PatchProfile(ctx context.Context, id string, p Profile) error {
	if p.ID != "" && id != p.ID {
		return ErrInconsistentIDs
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()

	// Keep the same id
	p.ID = id

	err := s.c.Update(bson.M{"id": id}, &p)
	return err
}

func (s *inmomService) DeleteProfile(ctx context.Context, id string) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	err := s.c.Remove(bson.M{"id": id})
	return err
}
