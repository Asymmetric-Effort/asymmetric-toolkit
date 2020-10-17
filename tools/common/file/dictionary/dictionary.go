package dictionary

import (
	"asymmetric-effort/asymmetric-toolkit/tools/common/errors"
	"encoding/base64"
	"time"
)

type netProtocols int
type portNumber int
type toolId int

const (
	any  netProtocols = 0
	tcp  netProtocols = 1
	udp  netProtocols = 2
	icmp netProtocols = 3
)

const (
	any toolId = 0
	//
	//Attack tools
	//
	dnsEnum         toolId = 1
	vladTheInjector toolId = 2
	//
	//Support tools
	//
	dnsProxy toolId = 65536
	vpnProxy toolId = 65537
	torProxy toolId = 65538
	i2pProxy toolId = 65539
)

type Dictionary struct {
	lastHit  int64                //unix timestamp of the last hit by any tool in the toolkit.
	rank     int                  //Aggregated score calculated periodically
	word     string               //Base64-encoded word
	created  int64                //unix timestamp when entry was created
	firstHit int64                //unix timestamp when any first hit with the entry by any tool in the toolkit.
	portsHit map[portNumber]int   //Number of hits by port number
	protoHit map[netProtocols]int //Number of hits by L3 protocol (tcp/udp/icmp)
	toolHit  map[toolId]int       //Number of hits by asymmetric-toolkit tool.
}

func (o *Dictionary) SetWord(s string) {
	errors.Assert(o.word == "", "Dictionary::SetWord() can only be used once to set the word value.")
	o.word = base64.StdEncoding.EncodeToString([]byte(s)) //ToDo: Compress/Encrypt word.
	o.created = time.Now().Unix()
	o.firstHit = 0
	o.lastHit = 0
	o.rank = 0
	o.portsHit = make(map[portNumber]int, 1)
	o.protoHit = make(map[netProtocols]int, 1)
	o.toolHit = make(map[toolId]int, 1)
}

func (o *Dictionary) GetWord() string {
	w, err := base64.StdEncoding.DecodeString(o.word)
	errors.Assert(err == nil, "Dictionary::GetWord() an error occurred getting the word content of a dictionary")
	return string(w)
}

func (o *Dictionary) AddHit(timestamp int64, port portNumber, protocol netProtocols, tool toolId) {
	if o.firstHit == 0 {
		o.firstHit = timestamp
	}
	o.lastHit = timestamp
	o.portsHit[port]++
	o.protoHit[protocol]++
	o.toolHit[tool]++
}

func (o *Dictionary) GetRank() int {
	return o.rank
}
