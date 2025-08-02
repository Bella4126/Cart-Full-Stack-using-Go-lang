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
	"context"
	"log"
	"net"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("DSN environment variable not found")
	}

	// Force IPv4-only resolver
	dialer := &net.Dialer{
		Timeout:   5 * time.Second,
		KeepAlive: 5 * time.Second,
		Resolver: &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				return net.Dial("tcp4", address)
			},
		},
	}

	sqlDB, err := dialer.Dial("tcp", "db.qhsdkmcdsrodcbvifrzh.supabase.co:5432")
	if err != nil {
		log.Fatalf("Custom TCP dial failed: %v", err)
	}
	defer sqlDB.Close()

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("✅ Database connected successfully via IPv4")
}
