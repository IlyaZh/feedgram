package news_checker

func (c *Component) Finish() {
	close(c.out)
}
