package news_checker

func (c *Component) Finish() {
	c.metrics.SyncDone()
	close(c.out)
}
