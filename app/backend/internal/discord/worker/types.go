package worker

const (
	authURL      string = "https://discord.com/api/oauth2/authorize"
	tokenURL     string = "https://discord.com/api/v10/oauth2/token"
	userEndpoint string = "https://discord.com/api/v10/users/@me"
	baseURL      string = "https://discord.com/api/v10"
)

type Service interface {
	Me(token string) error
	GuildChannels(guildID string) error
	Guilds() error
}

type Handler struct {
	Service
}

type Model struct {
}
