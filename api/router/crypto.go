package router

import (
	"fmt"
	"crypto/sha1"
	"encoding/hex"
	"strings"
	"errors"
	"github.com/masisiliani/bitBybit/types"
	"encoding/json"
	b64 "encoding/base64"
)

func newSession(username string) string{
	session := fmt.Sprintf(`{"user": "%s"}`, username)
	sha1Hash := getHash(session)
	
	hashedSession := fmt.Sprintf("%s.%s", session, sha1Hash )

	b64Session := b64.StdEncoding.EncodeToString([]byte(hashedSession))
	return b64Session
}

func getHash(content string) string{
	h := sha1.New()
	h.Write([]byte(content))
	sha1Hash := hex.EncodeToString(h.Sum(nil))
	return sha1Hash
}

func decodeSession(session string) (types.User, error) {
	dSession, err := b64.StdEncoding.DecodeString(session)
	u := types.User{}
	if err != nil{
		return u, err
	}
	s := strings.Split(string(dSession), ".")
	if len(s) != 2{
		return u, errors.New("invalid session")
	}
	hashedSession := getHash(s[0])
	if hashedSession != s[1]{
		return u, errors.New("invalid session")
	}
	
	err = json.Unmarshal([]byte(s[0]), &u)
	return u, err
	
}