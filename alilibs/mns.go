package alilibs

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/denverdino/aliyungo/mns"

	"github.com/otwdev/galaxylib"
)

type Mns struct {
	name   string
	client *mns.Client
	queue  *mns.Queue
}

func NewMns(section string) *Mns {
	m := &Mns{}

	sct, _ := galaxylib.GalaxyCfgFile.GetSection(section)

	m.name = sct["name"]
	url := sct["url"]
	id := sct["id"]
	secret := sct["secret"]

	m.client = mns.NewClient(id, secret, url)
	m.queue = &mns.Queue{
		Client:    m.client,
		QueueName: m.name,
		Base64:    false,
	}

	return m
}

func (m *Mns) Send(body interface{}) (ret string, err error) {

	buf, _ := json.Marshal(body)
	//body = string(buf)

	//ali_mns.NewMNSQueue(m.name, m.client)

	sendBody := mns.Message{
		MessageBody: string(buf),
	}

	xmlBuf, err := xml.Marshal(sendBody)
	if err != nil {
		return "", err
	}

	msg, err := m.queue.Send(mns.GetCurrentUnixMicro(), xmlBuf) //queue.SendMessage(msg)
	ret = msg.MessageId
	return ret, err
}

func (m *Mns) Receiver(msgChan chan mns.MsgReceive, errChan chan error) {

	m.queue.Receive(msgChan, errChan)
}

func (m *Mns) Delete(msgID string) {
	errChan := make(chan error)
	//fmt.Println(msgID)
	go func() {
		select {
		case err := <-errChan:
			{
				if err != nil {
					fmt.Printf("Delte error %s\n", err)
					return
				}
				fmt.Println("deleted....")
			}
		}
	}()
	m.queue.Delete(msgID, errChan)
}
