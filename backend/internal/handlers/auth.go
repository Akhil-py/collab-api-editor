package handlers

import (
    "net/http"
    "os"
    "github.com/labstack/echo/v4"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
)

var googleOauthConfig = &oauth2.Config{
    RedirectURL:  os.Getenv("GOOGLE_CALLBACK_URL"),
    ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
    ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
    Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
    Endpoint:     google.Endpoint,
}

func GoogleLogin(c echo.Context) error {
    url := googleOauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
    return c.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleCallback(c echo.Context) error {
    // TODO: Exchange code, get user info, create JWT, etc.
    return c.String(http.StatusOK, "Logged in!")
} 