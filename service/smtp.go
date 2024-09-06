package service

import (
	"crypto/tls"
	"net"
	"net/smtp"
	"strconv"
)

type SMTPClient interface {
	SendMail(from string, to []string, msg []byte) error
}

type smtpConnection struct {
	Host     string
	Port     int
	Username string
	Password string
	UseTLS   bool
}

func NewSMTPClient(host string, port int, username, password string, useTLS bool) SMTPClient {
	return &smtpConnection{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		UseTLS:   useTLS,
	}
}

func (ms *smtpConnection) SendMail(from string, to []string, msg []byte) error {
	addr := net.JoinHostPort(ms.Host, strconv.Itoa(ms.Port))

	// Set up authentication information.
	auth := smtp.PlainAuth("", ms.Username, ms.Password, ms.Host)

	// Connect to the server.
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Create a new client.
	client, err := smtp.NewClient(conn, ms.Host)
	if err != nil {
		return err
	}
	defer client.Quit()

	// Enable STARTTLS if specified.
	if ms.UseTLS {
		tlsConfig := &tls.Config{ServerName: ms.Host}
		if err = client.StartTLS(tlsConfig); err != nil {
			return err
		}
	}

	// Authenticate.
	if err = client.Auth(auth); err != nil {
		return err
	}

	// Set the sender and recipient.
	if err = client.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = client.Rcpt(addr); err != nil {
			return err
		}
	}

	// Send the email body.
	wc, err := client.Data()
	if err != nil {
		return err
	}
	if _, err = wc.Write(msg); err != nil {
		return err
	}
	if err = wc.Close(); err != nil {
		return err
	}

	return nil
}
