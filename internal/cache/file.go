package cache

import (
	"fmt"
	"os"
)

func init() {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		fmt.Println("Failed to get user cache directory: ", err)
		return
	}

	fmt.Println("User cache directory: ", cacheDir)
}
