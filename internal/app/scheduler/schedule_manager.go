package schedule

import (
	"log"

	"github.com/robfig/cron/v3"
)

type scheduleManagerImpl struct {
	driver *cron.Cron
}

func NewScheduleManager() *scheduleManagerImpl {
	cronOptions := []cron.Option{
		cron.WithSeconds(),
		cron.WithLogger(cron.VerbosePrintfLogger(log.Default())),
	}
	cron := cron.New(cronOptions...)

	return &scheduleManagerImpl{
		driver: cron,
	}
}

func (s *scheduleManagerImpl) AddSchedule(schedule cron.Schedule, job cron.Job) cron.EntryID {
	return s.driver.Schedule(schedule, job)
}

func (s *scheduleManagerImpl) Start() {
	s.driver.Start()
}
