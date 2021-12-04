package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Gif struct {
	gorm.Model
	Path string
	Tags []Tag `gorm:"many2many:gif_tags;"`
}

type Tag struct {
	gorm.Model
	Name string
}

type GifDb struct {
	db *gorm.DB
}

func (gdb *GifDb) init() {
	// db
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	panicIf(err)
	db.AutoMigrate(&Gif{}, &Tag{})
	gdb.db = db
}

func (gdb *GifDb) GetOrCreate(path string) *Gif {
	var gif Gif
	result := gdb.db.First(&gif, "path = ?", path)
	if result.Error == gorm.ErrRecordNotFound {
		gif.Path = path
		gdb.db.Create(&gif)
	}
	return &gif
}
