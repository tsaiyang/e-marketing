package dao

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error
var ctx = context.Background()

func InitDB() {
	db, err = gorm.Open(mysql.Open("root:jianyi1234@tcp(localhost:13316)/e_marketing?charset=utf8mb4&parseTime=True&loc=UTC"))
}

func init() {
	InitDB()
}

func Test_GetSenderListByPurpose(t *testing.T) {
	testCases := []struct {
		// 测试名字
		name string
		// 预期输入
		purpose string
		// 预期输出
		wantLen int
	}{
		{
			name:    "得到 purpose 为 universe 的 sender",
			purpose: "universe",
			wantLen: 1,
		},
	}

	senderDAO := NewSenderDAO(db)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			senders, er := senderDAO.GetSenderListByPurpose(ctx, tc.purpose)
			assert.NoError(t, er)

			t.Log(senders[0].Email)
			assert.Equal(t, tc.wantLen, len(senders))
		})
	}
}

func Test_GetEmailCountAndLimitTheDay(t *testing.T) {
	testCase := []struct {
		// 测试名字
		name string
		// 预期输入
		senderId int64
		// 预期输出
	}{
		{
			name:     "当天属于第一周，获得第一周的 limit 和当天的 count 数",
			senderId: 1,
		},
	}

	senderDAO := NewSenderDAO(db)

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			count, limit, err := senderDAO.GetEmailCountAndLimitTheDay(ctx, 1)
			assert.NoError(t, err)
			t.Log(count, limit)
		})
	}
}
