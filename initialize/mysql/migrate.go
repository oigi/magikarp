package mysql

import (
    "github.com/oigi/Magikarp/models/user"
    "os"
)

func migration() {
    err := _db.AutoMigrate(
        user.User{},
        //Todo 添加其他的
    )
    if err != nil {
        os.Exit(0)
    }
}
