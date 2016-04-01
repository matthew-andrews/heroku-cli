package main

import (
	"time"

	"github.com/franela/goreq"
)

// App is a Heroku app
type App struct {
	Name string `json:"name"`
}

func apps() ([]App, error) {
	var apps []App
	err := cacheFetch("apps.json", &apps, 20*time.Minute, func(o *interface{}) error {
		apps := make([]App, 0, 100)
		err := apiPartialRequests(apiToken(), "/apps", func(res *goreq.Response) error {
			var next []App
			if err := res.Body.FromJsonTo(&next); err != nil {
				return err
			}
			apps = append(apps, next...)
			return nil
		})
		*o = apps
		return err
	})
	return apps, err
}
