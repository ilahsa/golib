//
//常用的方法函数
package common

import (
	"crypto/md5"
	crand "crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

//GetMd5String is
//生成32位md5字串
func GetMd5String(b []byte) string {
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

//GetCurrentDir is
//获取当前目录
func GetCurrentDir() string {
	file, _ := exec.LookPath(os.Args[0])
	str := filepath.Dir(file)
	if !FileExist(str) {
		os.Mkdir(str, 0777)
	}
	return str
}

//FileExist is
//判断文件是否存在
func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

//GetLocalFileMd5String is
//计算文件的32位md5字串
func GetLocalFileMd5String(localFile string) string {
	file, inerr := os.Open(localFile)
	defer file.Close()

	if inerr == nil {
		md5h := md5.New()
		io.Copy(md5h, file)
		strMD5 := hex.EncodeToString(md5h.Sum(nil))
		return strMD5
	}
	return ""
}

//GetRandom is
//生成随机数
func GetRandom() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(250000)

}

//MakeGuid is
//生成guid 字符串
func MakeGuid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String([]byte(base64.URLEncoding.EncodeToString(b)))
}
