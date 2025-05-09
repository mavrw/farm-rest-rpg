package cookieutil

import (
	"time"

	"github.com/gin-gonic/gin"
)

const RefreshCookieName = "refresh_token"

func SetRefreshCookie(c *gin.Context, token string, maxAge time.Duration) {
	c.SetCookie(
		RefreshCookieName,     // name
		token,                 // value
		int(maxAge.Seconds()), // maxAge
		"/",                   // path
		"",                    // domain
		true,                  // secure - set to false within dev as needed, maybe move out to var
		true,                  // httpOnly
	)
}

func ClearRefreshTokenCookie(c *gin.Context) {
	c.SetCookie(
		RefreshCookieName, // name
		"",                // value
		-1,                // maxAge
		"/",               // path
		"",                // domain
		true,              // secure - set to false within dev as needed, maybe move out to var
		true,              // httpOnly
	)
}
