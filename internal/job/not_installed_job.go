package job

import (
	"context"
	"e-marketing/internal/model"
	"e-marketing/internal/service"
	"e-marketing/pkg/logger"
	"fmt"
	"time"

	"go.uber.org/zap"
)

// todo 测试改一下，后面改回来
const sendLimitOneTime = 2

type NotInstalledJob struct {
	shopService      service.ShopService
	emailService     service.EmailService
	emailSendService service.EmailSendService
	timeout          time.Duration
}

var _ Job = &NotInstalledJob{}

func NewUninstalledJob(
	shopService service.ShopService,
	emailService service.EmailService,
	emailSendService service.EmailSendService,
) *NotInstalledJob {
	return &NotInstalledJob{
		shopService:      shopService,
		emailService:     emailService,
		emailSendService: emailSendService,
		timeout:          10 * time.Minute,
	}
}

func (job *NotInstalledJob) Name() string {
	return "uninstalled"
}

func (job *NotInstalledJob) Run() error {
	ctx, cancel := context.WithTimeout(context.Background(), job.timeout)
	defer cancel()

	// 获取 sender 列表
	senders, err := job.emailService.GetSenderListByPurpose(ctx, model.SenderPurposeUniverse)
	if err != nil {
		return err
	}

	scenario, err := job.emailService.GetScenarioByCode(ctx, model.ScenarioCodeNotInstalled)
	if err != nil {
		return err
	}

	triggerRule, err := job.emailService.GetTriggerRuleByScenarioId(ctx, scenario.Id)
	if err != nil {
		return err
	}

	res := triggerRule.Params["days"].([]any)
	days := make([]int, 0, len(res))
	for _, day := range res {
		days = append(days, int(day.(float64)))
	}

	sended := make([]bool, len(senders))
	for _, day := range days {
		for i, sender := range senders {
			if sended[i] {
				continue
			}

			count, limit, err := job.emailService.GetEmailCountAndLimitTheDay(ctx, sender.Id)
			if err != nil {
				continue
			}

			// 这个邮箱当天发送的邮件数超过限制数，该邮箱就不提供发送服务
			if count >= limit {
				continue
			}

			offset, err := job.emailService.GetNotInstalledOffset(ctx, fmt.Sprintf("not_installed_%d", day))
			if err != nil {
				return err
			}

			recipients, err := job.emailService.GetRecipientList(ctx, int(offset), sendLimitOneTime)
			if err != nil || len(recipients) == 0 {
				continue
			}

			for _, recipient := range recipients {
				// todo 发送邮件
				err := job.emailSendService.Send(ctx, sender, recipient, "this is a test email", "1111", "")
				if err != nil {
					logger.Logger.Error("email sent failed",
						zap.String("from", sender.Email),
						zap.String("to", recipient.Email),
					)
				}
			}

			// 发送成功后，更新该邮箱今日已发送数，以及收件人列表的偏移量
			_ = job.emailService.IncrCountStatTheDay(ctx, sender.Id, len(recipients))
			_ = job.emailService.IncrNotInstalledOffset(ctx,
				fmt.Sprintf("not_installed_%d", day), len(recipients))

			sended[i] = true
		}
	}

	return nil
}
