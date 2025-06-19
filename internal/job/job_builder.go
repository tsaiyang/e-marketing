package job

import (
	"e-marketing/pkg/logger"
	"time"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type JobBuilder struct {
}

func NewJobBuilder() *JobBuilder {
	return &JobBuilder{}
}

func (builder *JobBuilder) Build(job Job) cron.Job {
	return cronJobAdapterFunc(func() {
		logger.Logger.Debug("定时任务开始运行",
			zap.String("name", job.Name()),
			zap.String("time", time.Now().UTC().Format(time.DateTime)),
		)

		if err := job.Run(); err != nil {
			logger.Logger.Error("任务执行失败",
				zap.Error(err),
				zap.String("name", job.Name()),
			)
		}
	})
}

type cronJobAdapterFunc func()

func (c cronJobAdapterFunc) Run() {
	c()
}
