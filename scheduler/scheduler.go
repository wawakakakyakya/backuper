package scheduler

import (
	"backuper/logger"
	"time"
)

type jobTicker struct {
	job  *Job
	tick *time.Time
}

type Scheduler struct {
	jobs   []*jobTicker
	logger logger.LoggerInterface
}

func NewScheduler(l logger.LoggerInterface) *Scheduler {
	return &Scheduler{logger: l}
}

func (s *Scheduler) parseTime(t string) (*time.Time, error) {
	until, err := time.Parse(time.RFC3339, "2023-06-22T15:04:05+02:00")
	if err != nil {
		s.logger.ErrorS("time parse error, job was not scheduled")
		s.logger.Error(err)
		return nil, err
	}
	return &until, nil
}

//add func and arg to slice
func (s *Scheduler) AddJob(job *Job, t string) error {
	until, err := s.parseTime(t)
	if err != nil {
		return err
	}
	s.jobs = append(s.jobs, &jobTicker{job, until})
	return nil
}

//run func and args with goroutine
func (s *Scheduler) Run() {
	for _, jt := range s.jobs {
		go jt.job.Run()
	}
}

// func (s *Scheduler) Run() {
// 	go func() {
// 		if err = runTask(mainCtx, c, l); err != nil {
// 			l.Error(err)
// 			eCode = NG
// 		} else {
// 			eCode = OK
// 		}
// 		wg.Done() // wg -1
// 	}()
// }
