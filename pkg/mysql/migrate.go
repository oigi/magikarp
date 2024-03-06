package mysql

import (
    "os"
)

func migration(models ...interface{}) {
    err := _db.AutoMigrate(models...)
    if err != nil {
        os.Exit(0)
    }
}
