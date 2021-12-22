package main

import (
	"fmt"
	"net/smtp"

	"github.com/joho/godotenv"
)

var log_email string
var passwd_email string
var auth smtp.Auth

func loadCredentials(cred string) string {
	var env_file map[string]string
	godotenv.Load(".env")
	env_file, err := godotenv.Read()

	if err != nil {
		fmt.Println("Erro loading .env")
	}

	return env_file[cred]
}

func sendMail(name string, adress string) {
	to := []string{adress}
	msg := []byte("To:" + adress + "\r\n" +
		"Subject: cheiro de prego pqp!\r\n" +
		"\r\n" +
		`Bolsonaro-chan eu te amo! 
		P-por que vc está t-tirando minhas roupas Bolsonaro-chan..
		AWNNNNNNNNNMMNN
		CONTINUA B-BOLSONARO-CHAN
		OWNNNNN NAO PARA VAI
		AAAAAAAENNNnnnn...
		e-está doendo um pouco ainda Bolsonaro chan...
		v-você pode passar uma pomada pra mim oppa?` + "\r\n")

	err := smtp.SendMail("smtp.gmail.com:587", auth, log_email, to, msg)

	if err != nil {
		fmt.Println(err)
	}
}

func sendToSingleMail(emails map[string]string) {
	for name, adress := range emails {
		sendMail(name, adress)
	}
}

func main() {
	log_email = loadCredentials("email")
	passwd_email = loadCredentials("passwd")
	auth = smtp.PlainAuth("", log_email, passwd_email, "smtp.gmail.com")

	emails := map[string]string{
		"Luca":    "smooqyoficial@outlook.com",
		"Sora":    "s0ra@tuta.io",
		"Raphael": "z1burg3@gmail.com",
	}

	sendToSingleMail(emails)
}
