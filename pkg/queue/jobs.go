package queue

import (
 "planet/pkg/gqueue/jobs"
)

var Jobs = map[string]jobs.Job{
"TestJob": jobs.TestJob{},
}