package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App        *App
	PgDB       *PgDB
	Email      *Email
	GoogleAuth *GoogleAuth
	S3         *S3
	Stripe     *Stripe
}

type App struct {
	AppUrl             string
	AppPort            string
	AppName            string
	FrontendURL        string
	ProductFrontendURL string
	TokenSecretKey     string
}

type PgDB struct {
	Host     string
	Username string
	Password string
	DbName   string
}

type Email struct {
	Name     string
	Address  string
	Password string
}

type GoogleAuth struct {
	ClientId     string
	ClientSecret string
	CallbackURL  string
}

type S3 struct {
	AwsRegion            string
	AwsAccessKeyID       string
	AwsSecretKey         string
	AwsBucketProfileName string
	AwsBucketEventName   string
}

type Stripe struct {
	PublicKey string
	SecretKey string
}

func NewConfig(path string) (*Config, error) {
	if err := godotenv.Load(path); err != nil {
		log.Println("Error loading .env file: ", err)
		return nil, err
	}
	appPort := os.Getenv("PROD_PORT")
	isProd := os.Getenv("IS_PROD")
	if isProd == "false" {
		appPort = os.Getenv("DEV_PORT")
	}
	return &Config{
		App: &App{
			AppUrl:             os.Getenv("SERVER_HOST"),
			AppPort:            appPort,
			AppName:            os.Getenv("APP_NAME"),
			FrontendURL:        os.Getenv("FRONTEND_URL"),
			ProductFrontendURL: os.Getenv("PRODUCT_FRONTEND_URL"),
			TokenSecretKey:     os.Getenv("TOKEN_SECRET_KEY"),
		},
		PgDB: &PgDB{
			Host:     os.Getenv("PG_HOST"),
			Username: os.Getenv("PG_USER"),
			Password: os.Getenv("PG_PASSWORD"),
			DbName:   os.Getenv("PG_DB"),
		},
		Email: &Email{
			Name:     os.Getenv("EMAIL_SENDER_NAME"),
			Address:  os.Getenv("EMAIL_SENDER_ADDRESS"),
			Password: os.Getenv("EMAIL_SENDER_PASSWORD"),
		},

		GoogleAuth: &GoogleAuth{
			ClientId:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			CallbackURL:  os.Getenv("GOOGLE_CALLBACK_URL"),
		},
		S3: &S3{
			AwsRegion:            os.Getenv("AWS_REGION"),
			AwsAccessKeyID:       os.Getenv("AWS_ACCESS_KEY_ID"),
			AwsSecretKey:         os.Getenv("AWS_SECRET_ACCESS_KEY"),
			AwsBucketProfileName: os.Getenv("AWS_BUCKET_PROFILE_NAME"),
			AwsBucketEventName:   os.Getenv("AWS_BUCKET_EVENT_NAME"),
		},
		Stripe: &Stripe{
			PublicKey: os.Getenv("STRIPE_PUBLIC_KEY"),
			SecretKey: os.Getenv("STRIPE_SECRET_KEY"),
		},
	}, nil
}
