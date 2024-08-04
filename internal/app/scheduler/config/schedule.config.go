package schedule_config

import (
	schedule "github.com/richhh7g/back-term-monitor/internal/app/scheduler"
)

type Scheduler struct{}

func NewSchedulerConfig() *Scheduler {
	return &Scheduler{}
}

func (s *Scheduler) Configure() {
	schedule.NewScheduleManager()
}
