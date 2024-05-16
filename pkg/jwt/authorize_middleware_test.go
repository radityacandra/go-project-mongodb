package jwt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/radityacandra/go-project-mongodb/internal/application/user/model"
	"github.com/radityacandra/go-project-mongodb/internal/server/types"
	"github.com/radityacandra/go-project-mongodb/mocks/internal_/application/user/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInitMiddleware(t *testing.T) {
	middleware := NewJwtMiddlewre(&types.Dependency{})

	assert.NotNil(t, middleware)
	assert.NotNil(t, middleware.repository)
}

func handlerFunc(c echo.Context) error {
	type response struct {
		Message string `json:"message"`
	}
	return c.JSON(http.StatusAccepted, response{})
}

func TestAuthorize(t *testing.T) {
	t.Setenv("JWT_PRIVATE_KEY", `
-----BEGIN RSA PRIVATE KEY-----
MIIG4wIBAAKCAYEAsdUuJC7bTgGQYqBRnv+VVoue6CcNwD3Ps6hGmBt8zHwjS7Ev
t6QLJ7dGv5ZsTqoNCp0uctEzl1XOzIoBCcf4vkHzFnvKx119D+/mYAH9dIO7TwRa
JtjrPEhrhGzGoFvoZPaW78io7UPnN3K6aEJkAfp9xb7F+ZdKGMyS0WyjN6nYFy4h
2A6y6CrjR5Z6rtHFbRczRufN7g9QA0G19hbFtKaLG2Vd4w5LgZoP0pnWVeqHSbzq
Y0saj4r/rhYQ+sk/ogRp1uzTMyByUQuNl74lJD/Ip4tFlQp+ISHjvMvL006qb+ph
zwmmnZoqfwmDKUY5TTeS8/KStJRtgcY3hMZzk0O6dWioZ2JLhd7uq/Sh+oXH2/gM
MKBURohxdGKGzhrTMlF2LI+E+fi5tgLatDVAocrYRVyphCy373iSIvUdCT1c5sGY
Yu2y1hc4RIhKBd8rvZWricYVK9Mgrskhi1am6HFVdrWdl+fizueWZlKwW9GMDrE1
rFMQjM6Mok7hQ5ZBAgMBAAECggGAZwOi8vrht1JYnYlZPs23aKcAqmLVKGOmCeCU
5FAn/xx6JcLCbZLtk1gr95ffrcH8RAXBfmmJhUUDwqC+8TeR9ESn2IJleQ/C/pRt
03fTpscYnjFbN924hvc8sT2B4irbeIEP9l863BbVvd0L9pFe9XZxnTdh36+wEKWQ
9xgDhJ8yulrS6CG7qZoOhs24y8RgpU0nOw/Uzu633usVyTRv2rIPKxORS5JgUvc3
6zuKWcK+FA0HifVbsYNHkUTfGSeyaYHQSBzWKI80Z4eFFSJYHGVbjRfSBLL9o3sG
MrADzKTOzbC431ZnCKGsWIo5eUqrSDFhgWuhw+mcKtR5RqJ7ybCEQqmfky9LO1pl
Aair/fa+ZegIweZ1tF3c54wVZwOr2ol+k7kUZDk8mlbWKNmmOFuJX3QKkSUi7Ww2
nTT2iwbfO3w0G/voj3XUBjKQAzvk0i59Fxe37HNJyHQwA3mZW/ljjPbiLXr6OlA3
ErCUYaJVG7LmV0/tWccq4IpBK4ABAoHBAN0FT2TGRAIwzfnNKi3waiIZ7W9xHObA
zPQ+wTLWDznO56+gVcmW1YhsgWCbQD7g6FDg1LTkoqtu8xoX+Ad9u4bnCkZpRsUn
kxGdEnmkIb9sN//O8b2wukkhSnIkBy7cIIL1nOjWu/sX0/9Bm/Pupgktpj+7Y7c6
D7Q/eR2nAxH0iHCw6ewBaCjL23nB5aDAkLc/rp0kJSUGXgQ95gSgN8hNpb+mf2nR
73m6o6fg0g3NbUWSphe+sgMHJgOyTE9+gQKBwQDN+hnroQ4W3v5ABUFsSfyB0+m2
zFkX/OlQ+Pi/Df7eHTSWRMUb6/11rsX98B9XckJtkRL58TCjkhISHkQNx5VOVfV0
oQPX6cIsattHcy1MXvzhUq13Hz0FqNmPXBu/VeAXrlMKAIhHXujDwnSePLzyNMxr
+5B7CHpOYS3jYEGCIa1OvhDxAJlA2N2rJS/kwYxyBDYD7oU2cHiSxHdGBUwapTar
9O7XEDUOmRLEpqI4m7dwTCiHZsfUOH5wSqBnt8ECgcEAqYX4Jmh4I2IBqnc1wfSj
wzI1zoNo0ojQ6wnzL6XGGeXcCPYmtq9tau0mss2Bknj5V4eLPpAbQPihl8MIp95F
ZqxxUh5PfMrk396lHA3LIMLFlPzKqF8UEyFos3F03PlUKn5u8pw5pNba8O8Gyiui
yGluGXYGfrQW9X61zAucrDnuKNlZIn7VwGzUE2ioUwtANH2w1bCymNpwZDqB4cxt
3MHJvtza33R5hNmlyw8CrjmBdZqMsKvAZ9gkzTMFmveBAoHAC3JMonWOzOKTiodA
PX5XE/fs0wXEJbseVCuh1yw176370CX+NjEFItcVlakUdM2at3AKd+1ZYJ2rd2pL
KeDfgTzxqQuRpRwOeF1v1iama8oTj4oCrc4EnB3oCTl7KUTicS502udwq1aw5MAX
rvt9HJCmk5GlU9ECyvxHio081rh8YxXY5yu7WIk0uGAWq0W+Qk/NmKZWrNHPQYi9
gBynQAvSX/f6leUGrcr/6gHAnhi5NvpkfjHOBRMmo6LiCyGBAoHAHfHkxtOh8jwS
TXBSUrHIM2+A5ZxaD/fqFOBxRIHx3yOCiAL/mfrqxbp6F29iZCQCmD0TzoQHJnn6
iGHRqcnOyi0dwqdzeBCBacGCYlRPPimcVWnuC/7fJkWV2S0g3rhQ++NwGlMOsQXl
IKGy2XMLebnnEAG/H2ra+1v/zDW9dTROgKrQckrSbIkKQke/FnMDX4h75iOArNKw
f9ssIXGETX3SAASpToGeHh/iqcMcOahE6K7niu47lMqAbwvmNugk
-----END RSA PRIVATE KEY-----
	`)

	t.Setenv("JWT_PUBLIC_KEY", `
-----BEGIN PUBLIC KEY-----
MIIBojANBgkqhkiG9w0BAQEFAAOCAY8AMIIBigKCAYEAsdUuJC7bTgGQYqBRnv+V
Voue6CcNwD3Ps6hGmBt8zHwjS7Evt6QLJ7dGv5ZsTqoNCp0uctEzl1XOzIoBCcf4
vkHzFnvKx119D+/mYAH9dIO7TwRaJtjrPEhrhGzGoFvoZPaW78io7UPnN3K6aEJk
Afp9xb7F+ZdKGMyS0WyjN6nYFy4h2A6y6CrjR5Z6rtHFbRczRufN7g9QA0G19hbF
tKaLG2Vd4w5LgZoP0pnWVeqHSbzqY0saj4r/rhYQ+sk/ogRp1uzTMyByUQuNl74l
JD/Ip4tFlQp+ISHjvMvL006qb+phzwmmnZoqfwmDKUY5TTeS8/KStJRtgcY3hMZz
k0O6dWioZ2JLhd7uq/Sh+oXH2/gMMKBURohxdGKGzhrTMlF2LI+E+fi5tgLatDVA
ocrYRVyphCy373iSIvUdCT1c5sGYYu2y1hc4RIhKBd8rvZWricYVK9Mgrskhi1am
6HFVdrWdl+fizueWZlKwW9GMDrE1rFMQjM6Mok7hQ5ZBAgMBAAE=
-----END PUBLIC KEY-----
	`)

	type testCase struct {
		name               string
		claims             map[string]interface{}
		getUserBySessionId *model.User
		expectCode         int
		expectErrMessage   string
	}

	testCases := []testCase{
		{
			name: "should return error if no active session found",
			claims: map[string]interface{}{
				"userId":    "userId1",
				"sessionId": "sessionId1",
			},
			expectCode:       http.StatusUnauthorized,
			expectErrMessage: "no active session is found",
		},
		{
			name:             "should return invalid token if claims doesn't complete",
			expectCode:       http.StatusUnauthorized,
			expectErrMessage: "invalid token",
		},
		{
			name: "should return success",
			claims: map[string]interface{}{
				"userId":    "userId1",
				"sessionId": "sessionId1",
			},
			getUserBySessionId: &model.User{},
			expectCode:         http.StatusAccepted,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			e := echo.New()
			repository := repository.NewMockIRepository(t)
			repository.EXPECT().GetUserBySessionId(mock.Anything, "userId1", "sessionId1").
				Return(testCase.getUserBySessionId).Maybe()

			middleware := &JwtMiddleware{
				repository: repository,
			}

			e.POST("/test", handlerFunc, middleware.Authorize)

			token, _, _ := BuildToken(testCase.claims)

			req := httptest.NewRequest(http.MethodPost, "/test", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", token))
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			authorizeFunc := middleware.Authorize(handlerFunc)
			err := authorizeFunc(c)

			assert.NoError(t, err)

			assert.Equal(t, testCase.expectCode, rec.Code)

			var body map[string]interface{}
			json.Unmarshal(rec.Body.Bytes(), &body)
			assert.Equal(t, testCase.expectErrMessage, body["message"])
		})
	}
}
