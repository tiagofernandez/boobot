package main

import (
	"html/template"
	"sort"
	"strings"

	"github.com/kataras/iris"
	"github.com/kataras/iris/config"
)

// register your auth via configs, providers with non-empty values will be registered to goth automatically by Iris
var oauth = config.OAuth{
	Path: "/auth", //defaults to /auth

	GithubKey:    "YOUR_GITHUB_KEY",
	GithubSecret: "YOUR_GITHUB_SECRET",
	GithubName:   "github", // defaults to github

	FacebookKey:    "YOUR_FACEBOOK_KEY",
	FacebookSecret: "YOUR_FACEBOOK_KEY",
	FacebookName:   "facebook", // defaults to facebook
}

func main() {
	// set the configs
	iris.Config.OAuth = oauth

	m := make(map[string]string)
	m[oauth.GithubName] = "Github"
	m[oauth.FacebookName] = "Facebook"

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	providerIndex := &ProviderIndex{Providers: keys, ProvidersMap: m}

	// if user succeed to logged in
	// client comes here from: localhost:3000/auth/lowercase_provider_name/callback
	iris.OnUserOAuth(func(ctx *iris.Context) {
		// ctx.OAuthUser() returns the authenticated goth.User
		// if user couldn't validate then server sends StatusUnauthorized, which you can handle by: iris.OnError(iris.StatusUnauthorized, func(ctx *iris.Context){})
		user := ctx.OAuthUser()

		// you can get the url by the predefined-named-route 'oauth'
		println("came from " + iris.URL("oauth", strings.ToLower(user.Provider)))

		t, _ := template.New("foo").Parse(userTemplate)
		ctx.ExecuteTemplate(t, user)
	})

	iris.Get("/", func(ctx *iris.Context) {
		t, _ := template.New("foo").Parse(indexTemplate)
		ctx.ExecuteTemplate(t, providerIndex)
	})

	iris.Listen(":3000")
}

// ProviderIndex ...
type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

var indexTemplate = `{{range $key,$value:=.Providers}}
    <p><a href="` + oauth.Path + `/{{$value}}">Log in with {{index $.ProvidersMap $value}}</a></p>
{{end}}`

var userTemplate = `
<p>Name: {{.Name}}</p>
<p>Email: {{.Email}}</p>
<p>NickName: {{.NickName}}</p>
<p>Location: {{.Location}}</p>
<p>AvatarURL: {{.AvatarURL}} <img src="{{.AvatarURL}}"></p>
<p>Description: {{.Description}}</p>
<p>UserID: {{.UserID}}</p>
<p>AccessToken: {{.AccessToken}}</p>
<p>ExpiresAt: {{.ExpiresAt}}</p>
<p>RefreshToken: {{.RefreshToken}}</p>
`
