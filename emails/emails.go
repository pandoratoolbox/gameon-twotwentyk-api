package emails

import (
	"context"
	"fmt"
	"gameon-twotwentyk-api/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mrz1836/postmark"
)

const (
	EMAIL_ADDRESS_NOREPLY           = "admin@pandoratoolbox.com"
	POSTMARK_API_TOKEN              = "74d5dba3-e53e-4ef0-86e0-5249d6eb8a99"
	POSTMARK_ACCOUNT_TOKEN          = "a04c0dc0-e985-42c0-af71-f07e7653cfb1"
	TEMPLATE_ID_CREATE_USER         = 1
	TEMPLATE_ID_VERIFY_USER         = 2
	TEMPLATE_ID_RESET_PASSWORD      = 1
	TEMPLATE_ID_CLAIM_STATUS_UPDATE = 1
)

var Postmark *postmark.Client

func init() {

	// auth := &http.Client{
	// 	Transport: &postmark.AuthTransport{
	// 		Token: POSTMARK_TOKEN,
	// 	},
	// }

	Postmark = postmark.NewClient(POSTMARK_API_TOKEN, POSTMARK_ACCOUNT_TOKEN)
}

type CustomClaims struct {
	Id int64 `json:"id"`
	UserId int64 `json:"userid"`
	jwt.StandardClaims
}

type ResetPasswordClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func generateToken(verification *models.Verification) (string, error) {
	claims := CustomClaims{
		Id: *verification.Id,
		UserId: *verification.UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(), // Token expiration time
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("h1l32b"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func SendResetEmail(to string, verification *models.Verification, user *models.User) error {
	claims := ResetPasswordClaims{
		Email: to,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(), // Token expiration time
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("h1l32b"))

	if err != nil {
		return err
	}
	
	email := postmark.Email{
		From:       EMAIL_ADDRESS_NOREPLY,
		To:         to,
		Subject:    "Reset Password",
		HTMLBody:   fmt.Sprintf(`<div>Hi, %s, 

		Click Button to reset password

		<button><a href="http://%s">Confirm Email Address</a></button>
		</div>`, *user.Username, signedToken),
	}

	fmt.Println(signedToken)

	_, err = Postmark.SendEmail(context.Background(), email)

	if err != nil {
		panic(err)
	}

	return err;
}

func SendVerifyEmail(to string, verification *models.Verification, user *models.User) error {
	token, _ := generateToken(verification);
	
	email := postmark.Email{
		From:       EMAIL_ADDRESS_NOREPLY,
		To:         to,
		Subject:    "Verify Your Email",
		HTMLBody:   fmt.Sprintf(`<div>Hi, %s, 

		There’s one quick step you need to complete before creating your account. Please enter this verification code to confirm this is your email. 
		
		%s 
		
		Verification codes expire after 30 minutes. 

		<button><a href="http://%s">Confirm Email Address</a></button>
		</div>`, *user.Username, *verification.Code, token),
	}

	fmt.Println(token)

	_, err := Postmark.SendEmail(context.Background(), email)

	if err != nil {
		panic(err)
	}

	return err;
}

// func SendTemplate(from string, to string, template_id int, template_model map[string]interface{}) error {

// 	res, err := Postmark.SendTemplatedEmail(context.Background(), postmark.TemplatedEmail{
// 		From:          from,
// 		To:            to,
// 		TemplateID:    int64(template_id),
// 		TemplateModel: template_model,
// 	})

// 	return nil
// }
