package mysql

import (
    "github.com/oigi/Magikarp/zmodels old/user214"
    "os"
)

func migration() {
    err := _db.AutoMigrate(
        user214.User{},
        //Todo 添加其他的
    )
    if err != nil {
        os.Exit(0)
    }
}
