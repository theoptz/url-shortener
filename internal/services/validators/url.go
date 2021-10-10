package validators

import "net/url"

type URLValidator struct {
}

func (u *URLValidator) Validate(val string) bool {
	v, err := url.Parse(val)
	return err == nil && (v.Scheme == "http" || v.Scheme == "https") && v.Host != ""
}

func NewURLValidator() *URLValidator {
	return &URLValidator{}
}
