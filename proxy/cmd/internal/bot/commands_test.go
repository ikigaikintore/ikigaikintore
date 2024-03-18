package bot

import (
	"github.com/ikigaikintore/ikigaikintore/proxy/cmd/internal/config"
	"gopkg.in/telebot.v3"
	"testing"
)

func Test_botServer_secure(t *testing.T) {
	type args struct {
		envCfg config.Envs
		msg    *telebot.Message
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "sender is registered user",
			args: args{
				envCfg: config.Envs{
					Telegram: config.Telegram{WebhookUserID: 1},
				},
				msg: &telebot.Message{
					ID:     123,
					Sender: &telebot.User{ID: 1},
				},
			},
			wantErr: false,
		},
		{
			name: "sender is bot",
			args: args{
				envCfg: config.Envs{
					Telegram: config.Telegram{WebhookUserID: 1},
				},
				msg: &telebot.Message{
					ID:     123,
					Sender: &telebot.User{ID: 12, IsBot: true},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &botServer{}
			if err := b.secure(tt.args.envCfg, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("secure() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
