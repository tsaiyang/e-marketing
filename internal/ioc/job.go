package ioc

import (
	"e-marketing/internal/job"

	"github.com/robfig/cron/v3"
)

func InitJobs(
	notInstalledJob *job.NotInstalledJob,
) *cron.Cron {
	builder := job.NewJobBuilder()

	expr := cron.New(cron.WithSeconds())
	if _, err := expr.AddJob("*/5 * * * * *", builder.Build(notInstalledJob)); err != nil {
		panic(err)
	}

	return expr
}
