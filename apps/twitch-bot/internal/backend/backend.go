package backend

import (
	"context"

	"github.com/senchabot-dev/monorepo/apps/twitch-bot/internal/models"
)

type Backend interface {
	GetTwitchChannels(ctx context.Context) ([]*models.TwitchChannel, error)
	CreateTwitchChannel(ctx context.Context, channelId string, channelName string, userId *string) (bool, error)

	GetTwitchBotConfig(ctx context.Context, twitchChannelId string, configKey string) (*models.TwitchBotConfig, error)

	GetBotCommand(ctx context.Context, commandName string, twitchChannelId string) (*models.BotCommand, error)
	CreateBotCommand(ctx context.Context, commandName string, commandContent string, twitchChannelId string, createdBy string) (*string, error)
	CheckCommandExists(ctx context.Context, commandName string, twitchChannelId string) (*string, error)
	UpdateBotCommand(ctx context.Context, commandName string, commandContent string, twitchChannelId string, updatedBy string) (*string, *string, error)
	DeleteBotCommand(ctx context.Context, commandName string, twitchChannelId string) (*string, *string, error)

	CreateBotActionActivity(ctx context.Context, botPlatformType string, botActivity string, twitchChannelId string, commandAuthor string) error

	GetCommandAlias(ctx context.Context, commandAlias string, twitchChannelId string) (*string, error)
	CreateCommandAliases(ctx context.Context, commandName string, aliases []string, twitchChannelId string, createdBy string) (*string, error)
	CheckCommandAliasExist(ctx context.Context, commandAlias string, twitchChannelId string) (*string, error)
	DeleteCommandAlias(ctx context.Context, commandAlias string, twitchChannelId string) (*string, error)
}
