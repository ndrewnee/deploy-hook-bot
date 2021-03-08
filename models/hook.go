package models

import "time"

type (
	Hook struct {
		Action          string          `json:"action"`
		Actor           Actor           `json:"actor"`
		CreatedAt       time.Time       `json:"created_at"`
		ID              string          `json:"id"`
		Data            Data            `json:"data"`
		PreviousData    PreviousData    `json:"previous_data"`
		PublishedAt     time.Time       `json:"published_at"`
		Resource        string          `json:"resource"`
		Sequence        interface{}     `json:"sequence"`
		UpdatedAt       time.Time       `json:"updated_at"`
		Version         string          `json:"version"`
		WebhookMetadata WebhookMetadata `json:"webhook_metadata"`
	}

	Actor struct {
		Email string `json:"email"`
		ID    string `json:"id"`
	}

	App struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	User struct {
		Email string `json:"email"`
		ID    string `json:"id"`
	}

	Data struct {
		App         App       `json:"app"`
		CreatedAt   time.Time `json:"created_at"`
		Description string    `json:"description"`
		Status      string    `json:"status"`
		ID          string    `json:"id"`
		Slug        string    `json:"slug"`
		UpdatedAt   time.Time `json:"updated_at"`
		User        User      `json:"user"`
		Version     int       `json:"version"`
	}

	PreviousData struct {
		Status string `json:"status"`
	}

	Attempt struct {
		ID string `json:"id"`
	}

	Delivery struct {
		ID string `json:"id"`
	}

	Event struct {
		ID      string `json:"id"`
		Include string `json:"include"`
	}

	Webhook struct {
		ID string `json:"id"`
	}

	WebhookMetadata struct {
		Attempt  Attempt  `json:"attempt"`
		Delivery Delivery `json:"delivery"`
		Event    Event    `json:"event"`
		Webhook  Webhook  `json:"webhook"`
	}
)
