package main

import (
	"fmt"
	"strings"

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

func (gdb *GifDb) AddTags(gif *Gif, names []string) {
	tagsMap := make(map[string]bool)

	var tags []Tag
	gdb.db.Where("name IN ?", names).Find(&tags)
	for _, tag := range tags {
		tagsMap[tag.Name] = true
	}
	for _, name := range names {
		if !tagsMap[name] {
			fmt.Println("Created new tag ", name)
			gdb.db.Model(gif).Association("Tags").Append(&Tag{Name: name})
		}
	}

}

func (gdb *GifDb) GetTags(gif *Gif) []string {
	var tagObjs []Tag
	tags := []string{}
	gdb.db.Model(gif).Association("Tags").Find(&tagObjs)
	for _, tag := range tagObjs {
		tags = append(tags, tag.Name)
	}
	fmt.Println("Returning tags: ", strings.Join(tags, ","))
	return tags
}

func (gdb *GifDb) init() {
	// db
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	panicIf(err)
	db.AutoMigrate(&Gif{}, &Tag{})
	gdb.db = db
}

func (gdb *GifDb) GetOrCreate(path string) (*Gif, bool) {
	var gif Gif
	created := false
	result := gdb.db.First(&gif, "path = ?", path)
	if result.Error == gorm.ErrRecordNotFound {
		gif.Path = path
		gdb.db.Create(&gif)
		created = true
	}
	return &gif, created
}
