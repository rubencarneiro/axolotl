package worker

import (
	"errors"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/nanu-c/textsecure-qml/contact"
	"github.com/nanu-c/textsecure-qml/ui"
)

func (Api *TextsecureAPI) SendContactAttachment(to, message string, file string) error {
	phone, err := contact.PhoneFromVCardFile(file)
	if err != nil {
		log.Println(err)
		return err
	}
	return Api.SendMessage(to, phone)
}
func (Api *TextsecureAPI) Test() {
	log.Printf("SendAttachmentApi")

}
func (Api *TextsecureAPI) SendAttachmentToApi(to, message string, file string) error {
	// Do not allow sending attachments larger than 100M for now
	var maxAttachmentSize int64 = 100 * 1024 * 1024
	log.Printf("SendAttachmentApi")
	fi, err := os.Stat(file)
	if err != nil {
		return err
	}
	if fi.Size() > maxAttachmentSize {
		ui.ShowError(errors.New("Attachment too large, not sending"))
		return nil
	}

	go SendMessageHelper(to, message, file)
	return nil
}