package discord

type StopCtx string
type WorkspaceCtx string
type DiscordReactionCtx string

const (
	AuthURL      string = "https://discord.com/api/oauth2/authorize"
	TokenURL     string = "https://discord.com/api/v10/oauth2/token"
	UserEndpoint string = "https://discord.com/api/v10/users/@me"
	BaseURL      string = "https://discord.com/api/v10"

	ChannelIdCtxKey DiscordReactionCtx = DiscordReactionCtx("ChannelIdCtxKey")
	AccessTokenCtxKey WorkspaceCtx = WorkspaceCtx("AuthorizationCtxKey")
	DiscordEventCtxKey WorkspaceCtx = WorkspaceCtx("DiscordEventCtxKey")
	WorkspaceCtxKey WorkspaceCtx = WorkspaceCtx("WorkspaceCtxKey")
	StopCtxKey = StopCtx("StopCtxKey")
)



