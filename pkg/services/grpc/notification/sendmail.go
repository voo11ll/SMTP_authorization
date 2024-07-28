package notification

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"strings"
	"time"
)

type Mail interface {
	Auth()
	Send(message Message) error
}

type SendMail struct {
	user     string
	password string
	host     string
	port     string
	auth     smtp.Auth
}

type Attachment struct {
	name        string
	fileName    string
	contentType string
	withFile    bool
}

type Message struct {
	from        string
	to          []string
	cc          []string
	bcc         []string
	body        []byte
	subject     string
	contentType string
	attachment  Attachment
}

func (mail *SendMail) Auth() {
	mail.auth = smtp.PlainAuth("", mail.user, mail.password, mail.host)
}

func (mail SendMail) Send(message Message) error {
	mail.Auth()

	buffer := bytes.NewBuffer(nil)
	boundary := "BnetBoundary"
	Header := make(map[string]string)
	Header["From"] = message.from
	Header["To"] = strings.Join(message.to, ";")
	Header["Cc"] = strings.Join(message.cc, ";")
	Header["Bcc"] = strings.Join(message.bcc, ";")
	Header["Subject"] = message.subject
	Header["Content-Type"] = "multipart/mixed;boundary=" + boundary
	Header["Mime-Version"] = "1.0"
	Header["Date"] = time.Now().String()
	mail.writeHeader(buffer, Header)

	if message.attachment.withFile {

		attachment := "\r\n--" + boundary + "\r\n"
		attachment += "Content-Transfer-Encoding:base64\r\n"
		attachment += "Content-Type:" + message.attachment.contentType + ";name=\"" + message.attachment.name + "\"\r\n"
		attachment += "Content-ID: <" + message.attachment.name + "> \r\n\r\n"
		buffer.WriteString(attachment)

		defer func() {
			if err := recover(); err != nil {
				fmt.Printf(err.(string))
			}
		}()
		mail.writeFile(buffer, message.attachment.fileName)
	}

	mimeHeaders := "\r\n--" + boundary + "\r\n"
	mimeHeaders += "Content-Type:text/html; charset=\"UTF-8\"\r\n"
	buffer.WriteString(mimeHeaders)

	buffer.Write(message.body)

	if err := smtp.SendMail(mail.host+":"+mail.port, mail.auth, message.from, message.to, buffer.Bytes()); err != nil {
		return err
	}
	return nil
}

func (mail SendMail) writeHeader(buffer *bytes.Buffer, Header map[string]string) string {
	header := ""
	for key, value := range Header {
		header += key + ":" + value + "\r\n"
	}
	header += "\r\n"
	buffer.WriteString(header)
	return header
}

// read and write the file to buffer
func (mail SendMail) writeFile(buffer *bytes.Buffer, fileName string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	payload := make([]byte, base64.StdEncoding.EncodedLen(len(file)))
	base64.StdEncoding.Encode(payload, file)
	buffer.WriteString("\r\n")
	for index, line := 0, len(payload); index < line; index++ {
		buffer.WriteByte(payload[index])
		if (index+1)%76 == 0 {
			buffer.WriteString("\r\n")
		}
	}
}
