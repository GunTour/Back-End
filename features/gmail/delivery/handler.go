package delivery

import (
	"GunTour/features/gmail/domain"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
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
	e.GET("/calendar", handler.GoCalendar())
	e.POST("/calendar/send", handler.GoCalendar())
}

func GetUrls() string {
	config := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/gmail.send", "https://www.googleapis.com/auth/calendar.events", "https://www.googleapis.com/auth/calendar"},
		Endpoint:     google.Endpoint,
		RedirectURL:  os.Getenv("REDIRECT_GMAIL"),
	}

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	authURL = strings.ReplaceAll(authURL, "\u0026", "&")
	// log.Print(authURL)
	return authURL
}

func GetUrlsCal() string {
	config := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/gmail.send", "https://www.googleapis.com/auth/calendar.events", "https://www.googleapis.com/auth/calendar"},
		Endpoint:     google.Endpoint,
		RedirectURL:  os.Getenv("REDIRECT_CALENDAR"),
	}

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	authURL = strings.ReplaceAll(authURL, "\u0026", "&")
	// log.Print(authURL)
	return authURL
}

func (gh *gmailHandler) GoSend() echo.HandlerFunc {
	return func(c echo.Context) error {
		var Code string
		var client *http.Client
		var messageStr []byte

		config := &oauth2.Config{
			ClientID:     os.Getenv("CLIENT_ID"),
			ClientSecret: os.Getenv("CLIENT_SECRET"),
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/gmail.send", "https://www.googleapis.com/auth/calendar.events", "https://www.googleapis.com/auth/calendar"},
			Endpoint:     google.Endpoint,
			RedirectURL:  os.Getenv("REDIRECT_GMAIL"),
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
			toks := config.TokenSource(oauth2.NoContext, tok)
			s, err := toks.Token()
			client = config.Client(oauth2.NoContext, s)
			res := ToDomain(s, Code)
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
					"Assalamu'alaikum Wr. Wb.\n\nMohon maaf, kami tidak dapat menerima permintaan anda.\nAnda tidak memenuhi syarat yang dibutuhkan.\nAnda dapat mengajukan diri kembali dilain waktu.\n\nTerima kasih!")
		} else {
			messageStr = []byte(
				"From: Guntour@gmail.com\r\n" +
					fmt.Sprintf("To: %v\r\n", msg.Email) +
					"Subject: Your form apply for ranger accepted\r\n\r\n" +
					"Assalamu'alaikum Wr. Wb.\n\nSelamat anda diterima menjadi ranger untuk aplikasi kami.\nMohon bantuan dan kerja samanya.\n\nTerima kasih!")
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

func (gh *gmailHandler) GoCalendar() echo.HandlerFunc {
	return func(c echo.Context) error {
		var Code string
		var client *http.Client

		config := &oauth2.Config{
			ClientID:     os.Getenv("CLIENT_ID"),
			ClientSecret: os.Getenv("CLIENT_SECRET"),
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/gmail.send", "https://www.googleapis.com/auth/calendar.events", "https://www.googleapis.com/auth/calendar"},
			Endpoint:     google.Endpoint,
			RedirectURL:  os.Getenv("REDIRECT_CALENDAR"),
		}

		Code = c.QueryParam("code")
		book := gh.srv.GetPesanCal()
		if Code == "" {
			resCode, err := gh.srv.GetCode()
			if err != nil {
				authURL := GetUrlsCal()
				log.Print(authURL)
				return c.JSON(http.StatusCreated, SuccessResponseBooking("success make booking", ToResponseGagal(book, authURL, "book")))
			}
			Code = resCode.Code
			tok := FromDomain(resCode)
			client = config.Client(oauth2.NoContext, tok)
		} else {
			tok, err := config.Exchange(oauth2.NoContext, Code)
			if err != nil {
				authURL := GetUrlsCal()
				// c.Redirect(http.Redirect(w, r, ))
				return c.JSON(http.StatusCreated, SuccessResponseBooking("success make booking", ToResponseGagal(book, authURL, "book")))
			}
			toks := config.TokenSource(oauth2.NoContext, tok)
			s, err := toks.Token()
			client = config.Client(oauth2.NoContext, s)
			res := ToDomain(s, Code)
			gh.srv.UpdateCode(res)
		}

		calendarService, err := calendar.New(client)
		if err != nil {
			GetUrlsCal()
			return c.JSON(http.StatusCreated, SuccessResponse("berhasil", "redirect"))
		}

		event := &calendar.Event{
			Summary:     "Your Climbing Day",
			Location:    "Taman Nasional Gunung Gede",
			Description: "Prepare for your greatest adventure.",
			Start: &calendar.EventDateTime{
				DateTime: fmt.Sprintf("%vT07:20:50.52Z", book.DateStart.Format("2006-01-02")),
				TimeZone: "Asia/Jakarta",
			},
			End: &calendar.EventDateTime{
				DateTime: fmt.Sprintf("%vT07:20:50.52Z", book.DateEnd.Format("2006-01-02")),
				TimeZone: "Asia/Jakarta",
			},
			Attendees: []*calendar.EventAttendee{
				{Email: fmt.Sprintf("%v", book.Email)},
			},
		}

		calendarID := "primary"
		event, err = calendarService.Events.Insert(calendarID, event).Do()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("unable to create an event."))
		}
		return c.JSON(http.StatusCreated, SuccessResponseBooking("success make booking", ToResponse(book, "book")))

	}
}
