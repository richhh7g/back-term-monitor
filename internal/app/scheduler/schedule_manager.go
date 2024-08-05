package schedule

import (
	"log"

	"github.com/robfig/cron/v3"
)

type scheduleManagerImpl struct {
	driver *cron.Cron
}

func NewScheduleManager() *scheduleManagerImpl {
	cronLog := cron.VerbosePrintfLogger(log.Default())

	cronOptions := []cron.Option{
		cron.WithSeconds(),
		cron.WithChain(cron.SkipIfStillRunning(cronLog)),
		cron.WithLogger(cronLog),
	}
	cron := cron.New(cronOptions...)

	cron.Start()

	return &scheduleManagerImpl{
		driver: cron,
	}
}

func (s *scheduleManagerImpl) AddSchedule(schedule cron.Schedule, job cron.Job) cron.EntryID {
	return s.driver.Schedule(schedule, job)
}

func (s *scheduleManagerImpl) Stop() {
	s.driver.Stop()
}
