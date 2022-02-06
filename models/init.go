package models

import (
    "github.com/glebarez/sqlite"
    "github.com/project-wind/wind-comment/config"
    "gorm.io/gorm"
)

var DB *gorm.DB

// Init initializes the database.
func Init() (err error) {
    DB, err = gorm.Open(sqlite.Open(config.Conf.Database), &gorm.Config{})
    if err != nil {
        return
    }

    err = DB.AutoMigrate(&Comment{})
    if err != nil {
        return
    }

    return
}
