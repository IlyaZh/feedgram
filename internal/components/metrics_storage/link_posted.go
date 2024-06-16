package metrics_storage

func (c *Component) LinkPosted(amount uint) {
	c.metrics.postedLinksCount.Add(float64(amount))
}

func (c *Component) LinkPostError(amount uint) {
	c.metrics.linkPostErrors.Add(float64(amount))
}
