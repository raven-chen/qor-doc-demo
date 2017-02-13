package main

import (
	"github.com/raven-chen/qor_doc_demo/config"
	"github.com/raven-chen/qor_doc_demo/config/routes"
	_ "github.com/raven-chen/qor_doc_demo/db"
	_ "github.com/raven-chen/qor_doc_demo/db/migration"
	"github.com/theplant/appkit/log"
	"github.com/theplant/appkit/server"
)

func main() {
	logger := log.Default()

	httpLogger := logger.With("origin", "http")
	server.ListenAndServe(config.Config.HTTP, httpLogger, routes.Mux(httpLogger))
}

// func setupExchange(db *gorm.DB) {
// 	// Define Resource
// 	product := exchange.NewResource(&models.Product{}, exchange.Config{PrimaryField: "Code"})
// 	// Define columns are exportable/importable
// 	product.Meta(&exchange.Meta{Name: "Code"})
// 	product.Meta(&exchange.Meta{Name: "Name"})
// 	product.Meta(&exchange.Meta{Name: "Price"})

// 	// Define context environment
// 	context := &qor.Context{DB: db}

// 	// Import products.csv into database
// 	product.Import(csv.New("products.csv"), context)

// 	// Export products into products.csv
// 	product.Export(csv.New("products.csv"), context)
// }
