package job

import (
	"fmt"
	"github.com/alissonphp/go-update-file/internal/application/use_cases/check_update"
	"github.com/robfig/cron/v3"
	"time"
)

func JobCheckUpdateImage(c *cron.Cron) {
	_, err := c.AddFunc("* * * * *", func() {
		err := check_update.CheckApi()
		if err != nil {
			panic(err)
		}
		fmt.Println("job executado", time.Now())
	})

	if err != nil {
		fmt.Println(err)
	}
}
