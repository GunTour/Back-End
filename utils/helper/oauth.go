package helper

import (
	admin "GunTour/features/admin/domain"
	booking "GunTour/features/booking/domain"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/gmail/v1"
)

func AuthConfig() *oauth2.Config {
	config := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/gmail.send", "https://www.googleapis.com/auth/calendar.events", "https://www.googleapis.com/auth/calendar"},
		Endpoint:     google.Endpoint,
		RedirectURL:  os.Getenv("REDIRECT_GMAIL"),
	}
	return config
}

func AuthConfigCalendar() *oauth2.Config {
	config := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/gmail.send", "https://www.googleapis.com/auth/calendar.events", "https://www.googleapis.com/auth/calendar"},
		Endpoint:     google.Endpoint,
		RedirectURL:  os.Getenv("REDIRECT_CALENDAR"),
	}
	return config
}

// FUNC TO SEND EMAIL
func SendMail(resCode admin.Code, msg admin.PesanCore) error {
	var messageStr []byte
	config := AuthConfig()

	tok := &oauth2.Token{AccessToken: resCode.AccessToken, TokenType: resCode.TokenType, RefreshToken: resCode.RefreshToken, Expiry: resCode.Expiry}
	client := config.Client(oauth2.NoContext, tok)

	gmailService, err := gmail.New(client)
	if err != nil {
		return errors.New("cannot create client")
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
		return errors.New("cannot send mail")
	}

	return nil
}

// FUNC TO ADD EVENT ON GOOGLE CALENDAR
func EventCalendar(resCode booking.Code, book booking.Core) error {
	config := AuthConfigCalendar()
	tok := &oauth2.Token{AccessToken: resCode.AccessToken, TokenType: resCode.TokenType, RefreshToken: resCode.RefreshToken, Expiry: resCode.Expiry}
	client := config.Client(oauth2.NoContext, tok)

	calendarService, err := calendar.New(client)
	if err != nil {
		return errors.New("cannot create event")
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
	event, err = calendarService.Events.Insert(calendarID, event).Do()
	if err != nil {
		return errors.New("cannot insert to calendar")
	}
	return nil
}

// FUNC TO GET GMAIL AUTH
func GetUrls() string {
	config := AuthConfig()
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	authURL = strings.ReplaceAll(authURL, "\u0026", "&")

	return authURL
}

// FUNC TO GET CALENDAR AUTH
func GetUrlsCal() string {
	config := AuthConfigCalendar()
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	authURL = strings.ReplaceAll(authURL, "\u0026", "&")

	return authURL
}
