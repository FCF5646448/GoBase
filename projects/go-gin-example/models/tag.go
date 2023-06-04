package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// Tag结构体，用于给Gorm使用，并给与了附属属性json，这样子在c.JSON的时候就会自动转换格式
type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// 这里的return 后面没有跟着变量。这是因为函数的声明中
// 已经明确显示了返回值是tags。
// db从哪里来？因为在同一个models包下，因此DB可以直接使用
func GetTags(pageNum int, pageSize int, maps interface{}) ([]Tag, error) {
	fmt.Printf("xxxxxxxx fcf gettags —— page: %d, size: %d", pageNum, pageSize)
	var (
		tags []Tag
		err  error
	)
	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Find(&tags).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&tags).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tags, nil
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

// 判断Tag是否存在
func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ? AND deleted_on = ?", name, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	return tag.ID > 0
}

func AddTag(name string, state int, createdBy string) error {
	tag := Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}
	db.AutoMigrate(&Tag{})
	if err := db.Create(&tag).Error; err != nil {
		fmt.Println("xxxxxxfcf: addtag 失败", err)
		return err
	}
	return nil
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	return true
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}
