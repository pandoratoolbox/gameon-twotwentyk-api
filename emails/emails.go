package emails

// import (
// 	"net/http"

// 	"github.com/mattevans/postmark-go"
// )

// const (
// 	EMAIL_ADDRESS_NOREPLY   = "noreply@audienceviral.com"
// 	POSTMARK_TOKEN          = "96fb75df-a707-48c9-8b7a-77691ade2914"
// 	TEMPLATE_ID_CREATE_USER = 1
// )

// var Postmark *postmark.Client

// func Init() error {

// 	auth := &http.Client{
// 		Transport: &postmark.AuthTransport{
// 			Token: POSTMARK_TOKEN,
// 		},
// 	}

// 	Postmark = postmark.NewClient(auth)

// 	return nil
// }

// func SendTemplate(from string, to string, template_id int, template_model map[string]interface{}) error {

// 	email := &postmark.Email{
// 		From:          from,
// 		To:            to,
// 		TemplateID:    template_id,
// 		TemplateModel: template_model,
// 		TrackOpens:    true,
// 	}

// 	_, _, err := Postmark.Email.Send(email)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
