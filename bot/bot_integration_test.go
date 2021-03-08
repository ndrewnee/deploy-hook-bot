// +build integration

package bot

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ndrewnee/deploy-hook-bot/config"
)

func TestBot_GetUpdatesChan(t *testing.T) {
	type args struct {
		config config.Config
	}

	tests := []struct {
		name    string
		args    args
		want    require.ValueAssertionFunc
		wantErr require.ErrorAssertionFunc
	}{
		{
			name: "Shouldn't get webhook chan because webhook host is empty",
			args: args{
				config: config.Config{
					Webhook:     true,
					WebhookHost: "",
				},
			},
			want:    require.Nil,
			wantErr: require.Error,
		},
		{
			name: "Should get webhook chan",
			args: args{
				config: config.Config{
					Webhook:     true,
					WebhookHost: "https://deploy-hook-bot.herokuapp.com",
				},
			},
			want:    require.NotNil,
			wantErr: require.NoError,
		},
		{
			name:    "Should get polling chan",
			want:    require.NotNil,
			wantErr: require.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tgbot, err := New()
			require.NoError(t, err)

			tgbot.config = tt.args.config

			got, err := tgbot.GetUpdatesChan()
			tt.wantErr(t, err)
			tt.want(t, got)
			// To avoid error "Too Many Requests: retry after 1"
			time.Sleep(time.Second)
		})
	}
}
