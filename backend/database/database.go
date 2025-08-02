// package database

// import (
// 	"log"
// 	"shopping-cart/models"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func Connect() {
// 	dsn := "host=db.qhsdkmcdsrodcbvifrzh.supabase.co user=postgres password=Java@1234@/#? dbname=postgres port=5432 sslmode=require TimeZone=Asia/Kolkata"

// 	var err error
// 	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Failed to connect to database:", err)
// 	}
// }

// func Migrate() {
// 	DB.AutoMigrate(&models.User{}, &models.Item{}, &models.Cart{}, &models.Order{})
// }

package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"shopping-cart/models"
)

var DB *gorm.DB
func Connect() {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("DSN environment variable not found")
	}

	// Force DNS to resolve via Go to avoid IPv6 problems
	os.Setenv("GODEBUG", "netdns=go")

	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully")
}


func Migrate() {
	DB.AutoMigrate(&models.User{}, &models.Item{}, &models.Cart{}, &models.Order{})
}
