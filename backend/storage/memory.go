package storage

import (
	"sparepart-api/models"
	"sync"
)

var (
	Users = map[string]string{
		"admin": "admin123",
	}

	Spareparts = []models.Sparepart{}
	LastID     = 0
	Mutex      = &sync.Mutex{}
)
