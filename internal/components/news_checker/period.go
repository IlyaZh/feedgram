package news_checker

import "time"

func (c *Component) Period() time.Duration {
	return c.config.GetValues().NewsChecker.Period
}
