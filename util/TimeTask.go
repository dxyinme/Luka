package util

import (
	"github.com/robfig/cron/v3"
)


// 开始一个新的TimeTask ,
// 建议数值为 */5 * * * * ? ,
// task 必须是 func()
func NewTimeTask(spec string, task func() ) error {
	c := cron.New(cron.WithSeconds())
	_, err := c.AddFunc(spec, task)
	c.Start()
	return err
}