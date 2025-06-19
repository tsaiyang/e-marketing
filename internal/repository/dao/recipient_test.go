package dao

import (
	"e-marketing/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetRecipientList(t *testing.T) {
	testCases := []struct {
		name   string
		offset int
		limit  int
	}{
		{
			name:   "获取收件人列表",
			offset: 0,
			limit:  50,
		},
	}

	dao := NewRecipientDAO(db)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			recipients, err := dao.GetRecipientList(ctx, int(tc.offset), tc.limit)
			assert.NoError(t, err)

			for _, r := range recipients {
				utils.PrintObj(r)
			}
		})
	}
}
