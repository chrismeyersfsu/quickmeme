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
	Name string `gorm:"unique"`
}

type GifDb struct {
	db *gorm.DB
}

func (gdb *GifDb) SetTags(gif *Gif, names []string) {
	tagsMap := make(map[string]Tag)
	notYetCreatedTagNames := []string{}

	var tags []Tag
	gdb.db.Where("name IN ?", names).Find(&tags)
	for _, tag := range tags {
		tagsMap[tag.Name] = tag
	}
	// Find tags that don't exist
	for _, name := range names {
		if _, ok := tagsMap[name]; !ok {
			notYetCreatedTagNames = append(notYetCreatedTagNames, name)
		}
	}
	// Create them
	fmt.Println(gif.ID, " Created new tag(s): [", strings.Join(notYetCreatedTagNames, ","), "]")
	createdTags := gdb.createTags(notYetCreatedTagNames)
	// Update map with newly created tags
	for _, tag := range createdTags {
		tagsMap[tag.Name] = tag
	}
	tags2append := []Tag{}
	for _, tag := range tagsMap {
		tags2append = append(tags2append, tag)
	}
	gdb.db.Model(gif).Association("Tags").Clear()
	gdb.db.Model(gif).Association("Tags").Append(&tags2append)
}

func (gdb *GifDb) GetTags(gif *Gif) []string {
	var tagObjs []Tag
	tags := []string{}
	gdb.db.Model(gif).Association("Tags").Find(&tagObjs)
	for _, tag := range tagObjs {
		tags = append(tags, tag.Name)
	}
	//fmt.Println(gif.ID, " Returning tags: ", strings.Join(tags, ","))
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

func (gdb *GifDb) createTags(names []string) []Tag {
	tags := []Tag{}

	for _, name := range names {
		tag := Tag{Name: name}
		gdb.db.Create(&tag)
		tags = append(tags, tag)
	}
	return tags
}
