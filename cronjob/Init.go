package cronjob

import "github.com/robfig/cron"

func Init() {

	c := cron.New()
	// refresh every day at 00:00
	spec := "0 0 0 * * *"
	err := c.AddFunc(spec, refresh)
	if err != nil {
		return
	}
	c.Start()
}
