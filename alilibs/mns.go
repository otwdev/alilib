package alilibs

import (
	"github.com/otwdev/galaxylib"
	"github.com/weikaishio/ali_mns"
)

type Mns struct {
	name   string
	client ali_mns.MNSClient
}

func NewMns(name string) *Mns {
	m := &Mns{}
	m.name = name
	url := galaxylib.GalaxyCfgFile.MustValue("alimns", "url")
	id := galaxylib.GalaxyCfgFile.MustValue("alimns", "id")
	secret := galaxylib.GalaxyCfgFile.MustValue("alimns", "secret")

	m.client = ali_mns.NewAliMNSClient(url, id, secret)

	return m
}

func (m *Mns) Send(body string) (ret string, err error) {

	msg := ali_mns.MessageSendRequest{
		MessageBody:  []byte(body),
		DelaySeconds: 0,
		Priority:     8}

	queue := ali_mns.NewMNSQueue(m.name, m.client)
	res, err := queue.SendMessage(msg)
	ret = res.MessageId
	return
}
