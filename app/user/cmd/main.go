package cmd

import (
    "github.com/oigi/Magikarp/initialize"
    "github.com/oigi/Magikarp/pkg/loading"
)

func main() {
    loading.Loading()
    router := initialize.Routers()
    router.Run(":8080")
}
