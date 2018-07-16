package models

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const secret = "name"

type Jwt struct {
	Header
	Playoad
}

type Header struct {
	Typ string `json:"typ"`
	Alg string `json:"alg"`
}

type Playoad struct {
	UserId int       `json:"userid"`
	Exp    time.Time `json:"exp"`
}

// var JwtToken = Jwt{
// 	// Typ: "JWT",
// 	// Alg: "HS256",
// 	// UserId: 0,
// 	// Exp: time.Now().AddDate(0, 0, 30),
// 	Header: &Header{Typ: "JWT", Alg: "HS256"},
// 	Playoad: &Playoad{UserId: 0, Exp: time.Now().AddDate(0,0,30)},
// }

func (self *Jwt) Token() string {

	header, _ := json.Marshal(self.Header)
	playoad, _ := json.Marshal(self.Playoad)

	h := base64.URLEncoding.EncodeToString(header)
	p := base64.URLEncoding.EncodeToString(playoad)

	end := fmt.Sprintf("%s.%s%s", h, p, secret)

	result := sha256.Sum256([]byte(end))
	return fmt.Sprintf("%s.%s.%s", h, p, hex.EncodeToString(result[:]))

}

func (self *Jwt) Checktoken(s string) bool {
	r := strings.Split(s, ".")
	if len(r) == 3 {
		header := r[0]
		playoad := r[1]
		final := r[2]

		h, _ := base64.URLEncoding.DecodeString(header)
		p, _ := base64.URLEncoding.DecodeString(playoad)

		json.Unmarshal(h, &self.Header)
		json.Unmarshal(p, &self.Playoad)
		if self.Exp.After(time.Now()) {
			end := fmt.Sprintf("%s.%s%s", header, playoad, secret)

			sum := sha256.Sum256([]byte(end))
			if hex.EncodeToString(sum[:]) == final {
				return true
			} else {
				return false
			}
		} else {
			return false
		}

	} else {
		return false
	}
}
