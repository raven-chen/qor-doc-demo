package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/publish2"
)

type User struct {
	gorm.Model
	Name   string
	Gender string
	Active bool
}

// Create another GORM-backend model
type Product struct {
	gorm.Model
	Name        string
	Description string
	// Image       media_library.MediaLibraryStorage `sql:"size:4294967295;" media_library:"url:/backend/{{class}}/{{primary_key}}/{{column}}.{{extension}};path:./private"`
	// Image aws.S3
	publish2.Version
	publish2.Schedule
	publish2.Visible
}

type Parameter struct {
	Height int
	Weight int
}

type Order struct {
	State string
	gorm.Model
}

type ColorVariation struct {
	gorm.Model
	ProductID uint
	Product   Product
	ColorName string
}
