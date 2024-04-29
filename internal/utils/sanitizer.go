package utils

import "github.com/microcosm-cc/bluemonday"

type Sanitizer interface {
	SanitizeHTML(text string) string
}

type Component struct {
	policy *bluemonday.Policy
}

var ptr *Component

func NewSanitizer() Sanitizer {
	if ptr == nil {
		ptr = &Component{
			policy: bluemonday.StrictPolicy(),
		}
	}
	return ptr
}

func (c *Component) SanitizeHTML(text string) string {
	return c.policy.Sanitize(text)
}
