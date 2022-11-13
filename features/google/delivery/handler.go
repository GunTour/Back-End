package delivery

import (
	"GunTour/features/google/domain"
	"GunTour/utils/helper"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/gmail/v1"
)

type googleHandler struct {
	srv domain.Services
}

func New(e *echo.Echo, srv domain.Services) {
	handler := googleHandler{srv: srv}
	e.GET("/gmail", handler.GoSend())
	e.PUT("/gmail/send", handler.GoSend())
	e.GET("/calendar", handler.GoCalendar())
	e.POST("/calendar/send", handler.GoCalendar())
}

// HANDLER TO TAKE CALLBACK FROM GOOGLE OAUTH2
func (gh *googleHandler) GoSend() echo.HandlerFunc {
	return func(c echo.Context) error {
		var Code string
		var client *http.Client
		var messageStr []byte
		config := helper.AuthConfig()

		Code = c.QueryParam("code")
		msg, ranger := gh.srv.GetPesan()
		tok, err := config.Exchange(oauth2.NoContext, Code)
		if err != nil {
			authURL := helper.GetUrls()
			return c.JSON(http.StatusAccepted, SuccessResponseRanger("success update status ranger", ToResponseGagal(ranger, authURL, "ranger")))
		}
		toks := config.TokenSource(oauth2.NoContext, tok)
		s, _ := toks.Token()
		client = config.Client(oauth2.NoContext, s)
		res := ToDomain(s, Code)
		gh.srv.InsertCode(res)

		gmailService, err := gmail.New(client)
		if err != nil {
			authURL := helper.GetUrls()
			return c.JSON(http.StatusAccepted, SuccessResponseRanger("cannot get client", ToResponseGagal(ranger, authURL, "ranger")))
		}
		var message gmail.Message

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

		message.Raw = base64.URLEncoding.EncodeToString(messageStr)

		_, err = gmailService.Users.Messages.Send("me", &message).Do()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("unable to send an email."))
		} else {
			return c.JSON(http.StatusAccepted, SuccessResponseRanger("success update status ranger", ToResponse(ranger, "ranger")))
		}
	}
}

// HANDLER TO TAKE CALLBACK FROM GOOGLE OAUTH2
func (gh *googleHandler) GoCalendar() echo.HandlerFunc {
	return func(c echo.Context) error {
		var Code string
		var client *http.Client
		config := helper.AuthConfigCalendar()

		Code = c.QueryParam("code")
		book := gh.srv.GetPesanCal()

		tok, err := config.Exchange(oauth2.NoContext, Code)
		if err != nil {
			authURL := helper.GetUrlsCal()
			return c.JSON(http.StatusCreated, SuccessResponseBooking("success make booking", ToResponseGagal(book, authURL, "book")))
		}
		toks := config.TokenSource(oauth2.NoContext, tok)
		s, _ := toks.Token()
		client = config.Client(oauth2.NoContext, s)
		res := ToDomain(s, Code)
		gh.srv.InsertCode(res)

		calendarService, err := calendar.New(client)
		if err != nil {
			authURL := helper.GetUrlsCal()
			return c.JSON(http.StatusAccepted, SuccessResponseRanger("cannot get client", ToResponseGagal(book, authURL, "ranger")))
		}

		event := &calendar.Event{
			Summary:     "Your Climbing Day",
			Location:    "Taman Nasional Gunung Gede",
			Description: "Prepare for your greatest adventure.",
			Start: &calendar.EventDateTime{
				DateTime: fmt.Sprintf("%vT00:20:50.52Z", book.DateStart.Format("2006-01-02")),
				TimeZone: "Asia/Jakarta",
			},
			End: &calendar.EventDateTime{
				DateTime: fmt.Sprintf("%vT01:20:50.52Z", book.DateEnd.Format("2006-01-02")),
				TimeZone: "Asia/Jakarta",
			},
			Attendees: []*calendar.EventAttendee{
				{Email: fmt.Sprintf("%v", book.Email)},
			},
		}

		calendarID := "primary"
		_, err = calendarService.Events.Insert(calendarID, event).Do()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("unable to create an event."))
		}
		return c.JSON(http.StatusCreated, SuccessResponseBooking("success make booking", ToResponse(book, "book")))

	}
}
