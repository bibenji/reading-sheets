package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"

	"../gotenv"
)

const facebookKey = ""
const facebookSecret = ""
const sessionSecret = ""

func init() {
	env, err := gotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(env)

	key := sessionSecret // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30 // 30 days
	isProd := false      // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store
}

// newProviderIndex return a ProviderIndex
func newProviderIndex() *providerIndex {
	goth.UseProviders(
		facebook.New(os.Getenv("FACEBOOK_KEY"), os.Getenv("FACEBOOK_SECRET"), "http://localhost:3000/auth/facebook/callback"),
		// facebook.New(facebookKey, facebookSecret, "http://localhost:8080/auth/facebook/callback"),
	)

	m := make(map[string]string)
	m["facebook"] = "Facebook"

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	providerIndex := &providerIndex{Providers: keys, ProvidersMap: m}

	return providerIndex
}

// providerIndex providerIndex struct
type providerIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}
