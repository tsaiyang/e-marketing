package dao

import (
	"e-marketing/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	InitDB()
}

func Test_GetScenarioByCode(t *testing.T) {
	testCases := []struct {
		// 测试名字
		name string
		// 预期输入
		code string
		// 预期输出
	}{
		{
			name: "获取未安装场景",
			code: "not_installed",
		},
	}

	dao := NewScenarioDAO(db)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s, err := dao.GetScenarioByCode(ctx, tc.code)
			assert.NoError(t, err)
			utils.PrintObj(s)
		})
	}
}

func Test_GetTriggerRuleByScenarioId(t *testing.T) {
	testCases := []struct {
		// 测试名字
		name string
		// 预期输入
		sid int64
	}{
		{
			name: "获取场景 id 为 1 的触发规则，正确将其 days 参数解析",
			sid:  1,
		},
	}

	dao := NewScenarioDAO(db)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rule, err := dao.GetTriggerRuleByScenarioId(ctx, tc.sid)
			assert.NoError(t, err)

			utils.PrintObj(rule)
			res := rule.Params["days"].([]any)
			days := make([]int, 0, len(res))
			for _, day := range res {
				days = append(days, int(day.(float64)))
			}

			t.Log("days", days)
		})
	}
}
