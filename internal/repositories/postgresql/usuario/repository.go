package usuario

import "gorm.io/gorm"

type Repository struct {
	DataBase *gorm.DB
}
