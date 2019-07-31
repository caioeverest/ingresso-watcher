package client

import (
	"encoding/gob"
	"fmt"
	"os"
	"time"

	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	whatsapp "github.com/Rhymen/go-whatsapp"
)

type WppConnection struct {
	conn *whatsapp.Conn
}

func InitWhatsAppConnection() (*WppConnection, error) {
	wac, err := whatsapp.NewConn(10 * time.Second)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating connection: %v\n", err)
		return nil, err
	}

	err = login(wac)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error logging in: %v\n", err)
		return nil, err
	}

	return &WppConnection{wac}, err
}

func (wpp *WppConnection) Send(phoneNumber, text string) (string, error) {
	message := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: fmt.Sprintf("%s@s.whatsapp.net", phoneNumber),
		},
		Text: text,
	}
	resp, err := wpp.conn.Send(message)
	if err != nil {
		if session, err := readSession(); err == nil {
			if session, err = wpp.conn.RestoreWithSession(session); err != nil {
				return nil, fmt.Errorf("restoring failed: %v\n", err)
			}

			return wpp.conn.Send(message)
		}
		return nil, err
	}

	return resp, nil
}

func login(wac *whatsapp.Conn) error {
	session, err := readSession()
	if err == nil {
		session, err = wac.RestoreWithSession(session)
		if err != nil {
			return fmt.Errorf("restoring failed: %v\n", err)
		}
	} else {
		qr := make(chan string)
		go func() {
			terminal := qrcodeTerminal.New2(qrcodeTerminal.ConsoleColors.BrightBlack,
				qrcodeTerminal.ConsoleColors.BrightWhite,
				qrcodeTerminal.QRCodeRecoveryLevels.Low)
			terminal.Get(<-qr).Print()
		}()
		session, err = wac.Login(qr)
		if err != nil {
			return fmt.Errorf("error during login: %v\n", err)
		}
	}

	err = writeSession(session)
	if err != nil {
		return fmt.Errorf("error saving session: %v\n", err)
	}
	return nil
}

func readSession() (whatsapp.Session, error) {
	session := whatsapp.Session{}
	file, err := os.Open("./whatsappSession.gob")
	if err != nil {
		return session, err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&session)
	if err != nil {
		return session, err
	}
	return session, nil
}

func writeSession(session whatsapp.Session) error {
	file, err := os.Create("./whatsappSession.gob")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(session)
	if err != nil {
		return err
	}
	return nil
}
