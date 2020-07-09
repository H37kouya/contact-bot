package helper_test

import (
	"contact-bot/pkg/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasePath(t *testing.T) {
	t.Run("BasePathを取得できる", func(t *testing.T) {
		assert.Exactly(t, helper.BasePath(".env"), "../.env")
	})
}
