package schedule_config

import (
	"time"

	schedule "github.com/richhh7g/back-term-monitor/internal/app/scheduler"
	schedule_job "github.com/richhh7g/back-term-monitor/internal/app/scheduler/job"
	"github.com/robfig/cron/v3"
)

type Scheduler struct{}

func NewSchedulerConfig() *Scheduler {
	return &Scheduler{}
}

func (s *Scheduler) Configure() {
	schedulerManager := schedule.NewScheduleManager()

	go schedulerManager.AddSchedule(cron.Every(10*time.Second), schedule_job.NewProcessCompetitors())
}
