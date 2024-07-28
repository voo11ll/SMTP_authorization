package notification

import (
	"auth/auth_back/pkg/globalvars"
	notificationRepository "auth/auth_back/pkg/repositories/notification"
	"bytes"
	context "context"
	"fmt"
	"text/template"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type GrpcServer struct {
	NotifyRepo *notificationRepository.NotificationRepository
	UnimplementedNotificationServiceServer
}

func (s *GrpcServer) SendMailConfirmLink(ctx context.Context, in *SendMailConfirmLinkRequest) (*SendMailConfirmLinkResponse, error) {

	var mail Mail
	var body bytes.Buffer

	linkId, err := uuid.Parse(in.GetLinkId())
	if err != nil {
		return &SendMailConfirmLinkResponse{
			StatusENUM: globalvars.ErrMailNotSend.Enum,
			Message:    err.Error(),
		}, nil
	}

	linkData := s.NotifyRepo.FindLinkById(linkId)
	if linkData == nil {
		return &SendMailConfirmLinkResponse{
			StatusENUM: globalvars.ErrMailNotSend.Enum,
			Message:    "link not found",
		}, nil
	}

	url := viper.GetString("api.url") + "user/email-confirm?id=" + linkData.ID.String() + "&key=" + linkData.HashKey

	t, _ := template.ParseFiles("./static/confirmMail.html")

	t.Execute(&body, struct {
		FirstName string
		Url       string
	}{
		FirstName: in.GetFirstName(),
		Url:       url,
	})

	mail = &SendMail{user: viper.GetString("mail.auth.user"), password: viper.GetString("mail.auth.password"), host: viper.GetString("mail.auth.host"), port: viper.GetString("mail.auth.port")}

	message := Message{from: viper.GetString("mail.auth.user"),
		to:          []string{in.Email},
		cc:          []string{},
		bcc:         []string{},
		subject:     "Bnet24: Подтверждение e-mil адреса",
		body:        body.Bytes(),
		contentType: "text/html;charset=utf-8",
	}

	err = mail.Send(message)

	if err != nil {
		return &SendMailConfirmLinkResponse{
			StatusENUM: globalvars.ErrMailNotSend.Enum,
			Message:    err.Error(),
		}, nil
	} else {
		return &SendMailConfirmLinkResponse{
			StatusENUM: "OK",
			Message:    "Email with confirmation link send--- " + url,
		}, nil
	}
}

func (s *GrpcServer) MailConfirmation(ctx context.Context, in *MailConfirmationRequest) (*MailConfirmationResponse, error) {
	linkId, err := uuid.Parse(in.GetLinkId())
	if err != nil {
		return &MailConfirmationResponse{
			StatusENUM: "ERR",
			Message:    "Link parse error",
		}, err
	}

	linkData := s.NotifyRepo.FindLinkById(linkId)
	if linkData == nil {
		return &MailConfirmationResponse{
			StatusENUM: "ERR",
			Message:    "Can't find Link in repo",
		}, fmt.Errorf("can't find Link in repo")
	}

	if in.GetHashKey() != linkData.HashKey {
		return &MailConfirmationResponse{
			StatusENUM: "ERR",
			Message:    "Can't find Link in repo",
		}, fmt.Errorf("can't find Link in repo")
	} else {
		_, err := s.NotifyRepo.UpdateLink(&notificationRepository.UpdateMailConfrimationLink{
			SendTry:   linkData.SendTry + 1,
			Confirmed: true,
		}, linkData.ID)
		if err != nil {
			return &MailConfirmationResponse{
				StatusENUM: "ERR",
				Message:    "Can't update Link in repo",
			}, fmt.Errorf("can't update Link in repo")
		}

	}

	return &MailConfirmationResponse{
		StatusENUM: "OK",
		Message:    "OK",
	}, nil
}
