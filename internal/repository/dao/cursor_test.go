package dao

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	InitDB()
}

func Test_Get(t *testing.T) {
	testCases := []struct {
		name       string
		nam        string
		wantOffset int64
	}{
		{
			name:       "拿到未安装第 0 天已发送的偏移量",
			nam:        "not_installed_0",
			wantOffset: 100,
		},
	}

	dao := NewCursorDAO(db)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			offset, err := dao.Get(ctx, tc.nam)
			assert.NoError(t, err)

			assert.Equal(t, tc.wantOffset, offset)
		})
	}
}
