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
}

func getClient(config *oauth2.Config, code string) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tok := getTokenFromWeb(config, code)

	return config.Client(oauth2.NoContext, tok)
}

func GetUrls() {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailSendScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	// client := getClient(config)

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	authURL = strings.ReplaceAll(authURL, "\u0026", "&")
	openbrowser(authURL)
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
		b, err := os.ReadFile("credentials.json")
		var Code string
		var client *http.Client
		// var er error
		if err != nil {
			log.Fatalf("Unable to read client secret file: %v", err)
		}

		// If modifying these scopes, delete your previously saved token.json.
		config, err := google.ConfigFromJSON(b, gmail.GmailSendScope)
		if err != nil {
			log.Fatalf("Unable to parse client secret file to config: %v", err)
		}

		Code = c.QueryParam("code")
		if Code == "" {
			resCode, err := gh.srv.GetCode()
			if err != nil {
				GetUrls()
				return c.JSON(http.StatusAccepted, SuccessResponse("berhasil", "redirect"))
			}
			Code = resCode.Code
			tok := FromDomain(resCode)
			client = config.Client(oauth2.NoContext, tok)
		} else {
			tok, err := config.Exchange(oauth2.NoContext, Code)
			if err != nil {
				GetUrls()
				return c.JSON(http.StatusAccepted, SuccessResponse("berhasil", "redirect"))
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
		// mail := "khalidrianda22@gmail.com"
		// Compose the message
		messageStr := []byte(
			"From: khalidrianda12@gmail.com\r\n" +
				"To: khalidrianda22@gmail.com\r\n" +
				"Subject: Your form apply for ranger accepted\r\n\r\n" +
				"Selamat anda diterima menjadi ranger untuk aplikasi kami\n Mohon bantuan dan kerja samanya\n Terima kasih!")

		// Place messageStr into message.Raw in base64 encoded format
		message.Raw = base64.URLEncoding.EncodeToString(messageStr)

		// Send the message
		_, err = gmailService.Users.Messages.Send("me", &message).Do()
		if err != nil {
			log.Printf("Error: %v", err)
		} else {
			return c.JSON(http.StatusAccepted, SuccessResponse("berhasil", "message sent!"))
		}
		return c.JSON(http.StatusAccepted, SuccessResponse("berhasil", "err"))
	}
}
