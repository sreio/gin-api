package models

import "fmt"

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