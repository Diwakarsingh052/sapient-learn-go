package graph

import (
	"gql/author/models"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {

	// Mutex to protect concurrent access to our data maps
	mutex sync.RWMutex
	
	// Map to store our data (in-memory database)
	authors     map[string]*models.Author // nil
	authorsList []*models.Author          // nil
}

func NewResolver() *Resolver {
	return &Resolver{
		authors:     make(map[string]*models.Author),
		authorsList: make([]*models.Author, 0, 100),
	}
}
