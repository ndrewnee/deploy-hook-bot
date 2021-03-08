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
	"github.com/ndrewnee/deploy-hook-bot/models"
)

func TestServer_HooksHandler(t *testing.T) {
	tgbot, err := bot.New()
	require.NoError(t, err)

	server := NewServer(tgbot)

	ts := httptest.NewServer(server.Mux())
	defer ts.Close()

	type args struct {
		body []byte
	}

	tests := []struct {
		name   string
		args   args
		status int
		want   models.HookResponse
	}{
		{
			name: "Should fail because body is invalid",
			args: args{
				body: []byte("invalid"),
			},
			status: http.StatusBadRequest,
			want: models.HookResponse{
				Error: "Unmarshal request body failed",
			},
		},
		{
			name: "Should get no content because action is not update",
			args: args{
				body: []byte("{}"),
			},
			status: http.StatusNoContent,
		},
		{
			name: "Should fail because send message to telegram failed",
			args: args{
				body: []byte(`{"action":"update","data":{"status":"[invalid"}}`),
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
			},
			status: http.StatusOK,
			want: models.HookResponse{
				Message: `[Build](https://dashboard.heroku.com/apps/lesswrong-bot/activity/builds/3e3572d2-3584-4b30-933a-8837177767ef)
App: lesswrong-bot
Commit: 53e36e65992352df09ac6610364a288239e11cb4
Status: succeeded
Published at: 2021-03-08 06:43:46 +0000 UTC`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := http.Post(ts.URL+"/hooks", "application/json", bytes.NewBuffer(tt.args.body))
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
