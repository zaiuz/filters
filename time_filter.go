package filters

import "time"
import z "github.com/zaiuz/zaiuz"

const TimeFilterPrefix = "z.time"

var zeroTime = time.Time{}

func TimeFilter() z.Filter {
	return func(action z.Action) z.Action {
		return func(c *z.Context) z.Result {
			startTime := time.Now()
			c.Set(TimeFilterPrefix+".start", startTime)
			result := action(c)

			finishTime := time.Now()
			duration := finishTime.Sub(startTime)

			c.Set(TimeFilterPrefix+".finish", finishTime)
			c.Set(TimeFilterPrefix+".duration", duration)
			return result
		}
	}
}

func GetStartTime(c *z.Context) time.Time {
	time_, ok := c.GetOk(TimeFilterPrefix+".start")
	if !ok {
		return zeroTime
	}

	time, ok := time_.(time.Time)
	if !ok {
		return zeroTime
	}

	return time
}

func GetFinishTime(c *z.Context) time.Time {
	time_, ok := c.GetOk(TimeFilterPrefix+".finish")
	if !ok {
		return zeroTime
	}

	time, ok := time_.(time.Time)
	if !ok {
		return zeroTime
	}

	return time
}

func GetDuration(c *z.Context) time.Duration {
	d_, ok := c.GetOk(TimeFilterPrefix+".duration")
	if !ok {
		return 0
	}

	d, ok := d_.(time.Duration)
	if !ok {
		return 0
	}

	return d
}
