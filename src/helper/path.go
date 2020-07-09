package helper

import (
	"fmt"
)

// BasePath Baseのパスを取得
func BasePath(path string) string {
	return fmt.Sprintf("../%s", path)
}
