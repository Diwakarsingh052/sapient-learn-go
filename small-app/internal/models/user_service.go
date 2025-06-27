package models

import (
	"database/sql"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

type CacheStore map[string]User

type Conn struct {
	db    *sql.DB
	cache CacheStore
	mu    *sync.RWMutex
}

func NewConn() Conn {
	return Conn{
		cache: make(CacheStore, 100),
		mu:    new(sync.RWMutex),
	}
}

func (c *Conn) CreateUser(n NewUser) (User, error) {

	passHash, err := bcrypt.GenerateFromPassword([]byte(n.Password), bcrypt.DefaultCost)
	if err != nil {

		return User{}, err
	}
	us := User{
		Id:           uuid.NewString(),
		Email:        n.Email,
		Name:         n.Name,
		Age:          n.Age,
		PasswordHash: string(passHash),
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	// put the new user in the map
	c.cache[n.Email] = us
	return us, nil
}

func (c *Conn) FetchUsers() map[string]User {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.cache
}
