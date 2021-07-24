package password

import (
	"bytes"
	"time"

	"github.com/yongcycchen/mall-api/common/crypt"
	"github.com/yongcycchen/mall-api/common/random"
)

func GenerateSalt() string {
	ranStr := time.Now().String() + random.KrandAll(32)

	var cipherTxt bytes.Buffer
	cipherTxt.WriteString(ranStr)
	return crypt.SHA1(cipherTxt.Bytes())
}

func GeneratePassword(password string, salt string) string {
	var str = "pw:" + password + ":salt:" + salt + ":github:common"

	var cipherTxt bytes.Buffer
	cipherTxt.WriteString(str)
	return crypt.SHA1(cipherTxt.Bytes())
}

func Check(dbPassword string, dbSalt string, password string) bool {

	return dbPassword == GeneratePassword(password, dbSalt)
}
