package delivery

import (
	"GunTour/features/gmail/domain"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

type gmailHandler struct {
	srv domain.Services
}

func New(e *echo.Echo, srv domain.Services) {
	handler := gmailHandler{srv: srv}
	// e.GET("/gmail/url", GetUrl()) // GET LIST PENDAKI
	e.GET("/gmail", handler.GoSend())
	e.PUT("/gmail/send", handler.GoSend())
}

func getClient(config *oauth2.Config, code string) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tok := getTokenFromWeb(config, code)

	return config.Client(oauth2.NoContext, tok)
}

func GetUrls() string {
	config := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/calendar.events", "https://www.googleapis.com/auth/calendar", "https://www.googleapis.com/auth/gmail.send"},
		Endpoint:    google.Endpoint,
		RedirectURL: os.Getenv("REDIRECT_GMAIL"),
	}

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	authURL = strings.ReplaceAll(authURL, "\u0026", "&")
	// log.Print(authURL)
	return authURL
}

func openbrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config, code string) *oauth2.Token {

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func (gh *gmailHandler) GoSend() echo.HandlerFunc {
	return func(c echo.Context) error {
		var Code string
		var client *http.Client
		var messageStr []byte

		config := &oauth2.Config{
			ClientID:     os.Getenv("CLIENT_ID"),
			ClientSecret: os.Getenv("CLIENT_SECRET"),
			Scopes: []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile",
				"https://www.googleapis.com/auth/calendar.events", "https://www.googleapis.com/auth/calendar", "https://www.googleapis.com/auth/gmail.send"},
			Endpoint:    google.Endpoint,
			RedirectURL: os.Getenv("REDIRECT_GMAIL"),
		}

		Code = c.QueryParam("code")
		msg, ranger := gh.srv.GetPesan()
		if Code == "" {
			resCode, err := gh.srv.GetCode()
			if err != nil {
				authURL := GetUrls()
				return c.JSON(http.StatusAccepted, SuccessResponseRanger("success update status ranger", ToResponseGagal(ranger, authURL, "ranger")))
			}
			Code = resCode.Code
			tok := FromDomain(resCode)
			client = config.Client(oauth2.NoContext, tok)
		} else {
			tok, err := config.Exchange(oauth2.NoContext, Code)
			if err != nil {
				authURL := GetUrls()
				// c.Redirect(http.Redirect(w, r, ))
				return c.JSON(http.StatusAccepted, SuccessResponseRanger("success update status ranger", ToResponseGagal(ranger, authURL, "ranger")))
			}
			client = config.Client(oauth2.NoContext, tok)
			res := ToDomain(tok, Code)
			gh.srv.UpdateCode(res)
		}

		gmailService, err := gmail.New(client)
		if err != nil {
			GetUrls()
			return c.JSON(http.StatusAccepted, SuccessResponse("berhasil", "redirect"))
		}
		var message gmail.Message

		// Compose the message
		if msg.Status == "rejected" {
			messageStr = []byte(
				"From: Guntour@gmail.com\r\n" +
					fmt.Sprintf("To: %v\r\n", msg.Email) +
					"Subject: Your form apply for ranger rejected\r\n\r\n" +
					"Mohon maaf, kami tidak dapat menerima permintaan anda\n Anda tidak memenuhi syarat yang dibutuhkan\n Terima kasih!")
		} else {
			messageStr = []byte(
				"From: Guntour@gmail.com\r\n" +
					fmt.Sprintf("To: %v\r\n", msg.Email) +
					"Subject: Your form apply for ranger accepted\r\n\r\n" +
					"Selamat anda diterima menjadi ranger untuk aplikasi kami\n Mohon bantuan dan kerja samanya\n Terima kasih!")
		}

		// Place messageStr into message.Raw in base64 encoded format
		message.Raw = base64.URLEncoding.EncodeToString(messageStr)

		// Send the message
		_, err = gmailService.Users.Messages.Send("me", &message).Do()
		if err != nil {
			return c.JSON(http.StatusAccepted, SuccessResponse("gagal", "permintaan sedang diproses"))
		} else {
			return c.JSON(http.StatusAccepted, SuccessResponseRanger("success update status ranger", ToResponse(ranger, "ranger")))
		}
	}
}
