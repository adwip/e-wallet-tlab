package drivers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase(url string) (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open(url), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
