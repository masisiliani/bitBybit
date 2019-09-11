package router

import (
	"fmt"
	"crypto/sha1"
	"encoding/hex"
	"strings"
	"errors"
	"github.com/masisiliani/bitBybit/types"
	"encoding/json"

)

func newCookie(username string) string{
	cookie := fmt.Sprintf(`{"user": %s}`, username)
	sha1Hash := getHash(cookie)
	
	hashedCookie := fmt.Sprintf("%s.%s", cookie, sha1Hash )
	return hashedCookie
}

func getHash(content string) string{
	h := sha1.New()
	h.Write([]byte(content))
	sha1Hash := hex.EncodeToString(h.Sum(nil))
	return sha1Hash
}

func decodeCookie(cookie string) (types.User, error) {
	u := types.User{}
	s := strings.Split(cookie, ".")
	hashedCookie := getHash(s[0])
	if hashedCookie != s[1]{
		return u, errors.New("invalid cookie")
	}
	
	err := json.Unmarshal([]byte(s[0]), &u)
	return u, err
	
}