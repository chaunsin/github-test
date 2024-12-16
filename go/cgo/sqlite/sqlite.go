package sqlite

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	Id        int64 `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Client struct {
	db *gorm.DB
}

func New(dsn string) *Client {
	var c Client

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to SQLite:", err)
	}
	c.db = db
	fmt.Println("SQLite connection established!")

	// 2. 自动迁移（AutoMigrate）：创建或更新表结构
	if err = db.AutoMigrate(&User{}); err != nil {
		log.Fatalln("Failed to migrate database schema:", err)
	}
	fmt.Println("Database schema migrated successfully!")
	return &c
}

func (c *Client) Insert(ctx context.Context, data *User) (int64, error) {
	tx := c.db.WithContext(ctx).Create(data)
	return tx.RowsAffected, tx.Error
}

func (c *Client) Query(ctx context.Context, id int) (*User, error) {
	var u User
	if err := c.db.WithContext(ctx).Where("id = ?", id).First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}
