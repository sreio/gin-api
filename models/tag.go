package models

import (
	"fmt"
	_ "time"

	_ "github.com/jinzhu/gorm"
)

type Tag struct{
	Model
	
	Name string `gorm:"column:name" json:"name"`
    CreatedBy string `gorm:"column:created_by" json:"created_by"`
    ModifiedBy string `gorm:"column:modified_by" json:"modified_by"`
    State int `gorm:"column:state" json:"state"`
}

func GetTags(PageNum int, PageSize int, maps interface {}) (tags []*Tag) {
	db.Where(maps).Offset(PageNum).Limit(PageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface {}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool {
    var tag Tag
    db.Select("id").Where("name = ?", name).First(&tag)
	fmt.Println(tag)
    return tag.ID > 0
}

func AddTag(name string, state int, createdBy string) error{
    tag := Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}
	if err := db.Model(&Tag{}).Create(&tag).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}

func EditTag(id int, data interface{}) bool{
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	return true
}

func ExistTagByID(id int) bool {
	var tag Tag
    db.Select("id").Where("id = ?", id).First(&tag)
	fmt.Println(tag)
    return tag.ID > 0
}

// func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
// 	scope.SetColumn("CreatedOn", time.Now().Unix())
// 	return nil
// }

// func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
// 	scope.SetColumn("ModifiedOn", time.Now().Unix())
// 	return nil
// }

func CleanAllTag() bool {
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{})
	return true
}