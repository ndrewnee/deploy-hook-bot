package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ndrewnee/deploy-hook-bot/bot"
	"github.com/ndrewnee/deploy-hook-bot/config"
	"github.com/ndrewnee/deploy-hook-bot/models"
)

func TestServer_HooksHandler(t *testing.T) {
	config := config.Parse()

	tgbot, err := bot.New()
	require.NoError(t, err)

	server := NewServer(config, tgbot)

	ts := httptest.NewServer(server.Mux())
	defer ts.Close()

	type args struct {
		authToken string
		body      []byte
	}

	tests := []struct {
		name   string
		args   args
		status int
		want   models.HookResponse
	}{
		{
			name: "Should fail because auth token is not valid",
			args: args{
				body:      []byte("{}"),
				authToken: "invalid",
			},
			status: http.StatusForbidden,
			want: models.HookResponse{
				Error: "Authorization is invalid",
			},
		},
		{
			name: "Should fail because body is invalid",
			args: args{
				body:      []byte("invalid"),
				authToken: config.AuthToken,
			},
			status: http.StatusBadRequest,
			want: models.HookResponse{
				Error: "Unmarshal request body failed",
			},
		},
		{
			name: "Should get no content because action is not update",
			args: args{
				body:      []byte("{}"),
				authToken: config.AuthToken,
			},
			status: http.StatusNoContent,
		},
		{
			name: "Should fail because send message to telegram failed",
			args: args{
				body:      []byte(`{"action":"update","data":{"status":"[invalid"}}`),
				authToken: config.AuthToken,
			},
			status: http.StatusInternalServerError,
			want: models.HookResponse{
				Error: "Send message to telegram failed",
			},
		},
		{
			name: "Should sent message to telegram",
			args: args{
				body: func() []byte {
					file, err := ioutil.ReadFile("testdata/hook_request.json")
					require.NoError(t, err)

					return file
				}(),
				authToken: config.AuthToken,
			},
			status: http.StatusOK,
			want: models.HookResponse{
				Message: `🛠 [Build](https://dashboard.heroku.com/apps/lesswrong-bot/activity/builds/3e3572d2-3584-4b30-933a-8837177767ef)

*App*: lesswrong-bot

*Commit*: 53e36e65992352df09ac6610364a288239e11cb4

*Status*: succeeded

*Published*: 2021-03-08 11:43:46`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, ts.URL+"/hooks", bytes.NewBuffer(tt.args.body))
			require.NoError(t, err)

			req.Header.Set("Content-Type", "application/json")
			if tt.args.authToken != "" {
				req.Header.Set("Authorization", "Bearer "+tt.args.authToken)
			}

			response, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			require.Equal(t, tt.status, response.StatusCode)

			defer response.Body.Close()

			if response.StatusCode != http.StatusNoContent {
				var got models.HookResponse
				err = json.NewDecoder(response.Body).Decode(&got)
				require.NoError(t, err)

				require.Equal(t, tt.want, got)
			}
		})
	}
}
