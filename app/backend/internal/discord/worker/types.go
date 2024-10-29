package worker

const (
	authURL      string = "https://discord.com/api/oauth2/authorize"
	tokenURL     string = "https://discord.com/api/v10/oauth2/token"
	userEndpoint string = "https://discord.com/api/v10/users/@me"
	baseURL      string = "https://discord.com/api/v10"
)

type Service interface {
	Me(token string) (*Me, error)
	GuildChannels(guildID string) ([]Channel, error)
}

type Handler struct {
	Service
}

type Model struct {
}

type Channel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type Me struct {
	DiscordId string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Avatar string `json:"avatar"`
}

type DiscordMe struct {
	Id                    string      `json:"id"`
	Username              string      `json:"username"`
	Avatar                *string     `json:"avatar"
	Discriminator         string      `json:"discriminator"`
	PublicFlags           int         `json:"public_flags"`
	Flags                 int         `json:"flags"`
	Banner                *string     `json:"banner"
	AccentColor           *string     `json:"accent_color"
	GlobalName            string      `json:"global_name"`
	AvatarDecorationData  *string     `json:"avatar_decoration_data"
	BannerColor           *string     `json:"banner_color"
	Clan                  *string     `json:"clan"
	MFAEnabled            bool        `json:"mfa_enabled"`
	Locale                string      `json:"locale"`
	PremiumType           int         `json:"premium_type"`
	Email                 string      `json:"email"`
	Verified              bool        `json:"verified"`
}
