package multiprovider

import (
	"os"
	"strconv"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

type Provider interface {
}

func CreateMultiProvider(providers ...goth.Provider) {
	key := os.Getenv("AUTH_KEY")
	maxAge64, err := strconv.ParseInt(os.Getenv("AUTH_MAX_AGE"), 10, 64)
	if err != nil {
		maxAge64 = int64(86400 * 30)
	}
	isProd, err := strconv.ParseBool(os.Getenv("AUTH_IS_PROD"))
	if err != nil {
		isProd = false
	}

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(int(maxAge64))
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		providers...,
	)
}
