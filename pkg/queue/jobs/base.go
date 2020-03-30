package jobs

import "encoding/json"

type Job interface {
	Execute(interface{}) error //执行任务
}