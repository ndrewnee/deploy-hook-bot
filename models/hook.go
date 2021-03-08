package models

import "time"

type (
	Hook struct {
		Action      string    `json:"action"`
		PublishedAt time.Time `json:"published_at"`
		Data        Data      `json:"data"`
	}

	Data struct {
		ID     string `json:"id"`
		Status string `json:"status"`
		App    App    `json:"app"`
		Slug   Slug   `json:"slug"`
	}

	App struct {
		Name string `json:"name"`
	}

	Slug struct {
		Commit string `json:"commit"`
	}
)

type (
	HookResponse struct {
		Error   string `json:"error,omitempty"`
		Message string `json:"message,omitempty"`
	}
)
