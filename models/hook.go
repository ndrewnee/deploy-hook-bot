package models

import "time"

type (
	HookAPIBuild struct {
		ID              string          `json:"id"`
		Action          string          `json:"action"`
		Version         string          `json:"version"`
		Resource        string          `json:"resource"`
		Sequence        interface{}     `json:"sequence"`
		CreatedAt       time.Time       `json:"created_at"`
		UpdatedAt       time.Time       `json:"updated_at"`
		PublishedAt     time.Time       `json:"published_at"`
		Actor           Actor           `json:"actor"`
		Data            Data            `json:"data"`
		PreviousData    PreviousData    `json:"previous_data"`
		WebhookMetadata WebhookMetadata `json:"webhook_metadata"`
	}

	App struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	Slug struct {
		ID                string `json:"id"`
		Commit            string `json:"commit"`
		CommitDescription string `json:"commit_description"`
	}

	User struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}

	Release struct {
		ID      string `json:"id"`
		Version int    `json:"version"`
	}

	Metadata struct{}

	Buildpacks struct {
		URL string `json:"url"`
	}

	SourceBlob struct {
		URL      string `json:"url"`
		Version  string `json:"version"`
		Checksum string `json:"checksum"`
	}

	Data struct {
		ID              string       `json:"id"`
		App             App          `json:"app"`
		Slug            Slug         `json:"slug"`
		User            User         `json:"user"`
		Stack           string       `json:"stack"`
		Status          string       `json:"status"`
		Release         Release      `json:"release"`
		Metadata        Metadata     `json:"metadata"`
		Buildpacks      []Buildpacks `json:"buildpacks"`
		CreatedAt       time.Time    `json:"created_at"`
		UpdatedAt       time.Time    `json:"updated_at"`
		SourceBlob      SourceBlob   `json:"source_blob"`
		OutputStreamURL string       `json:"output_stream_url"`
	}

	Actor struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}

	PreviousData struct{}

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
