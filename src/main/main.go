package main

import (
	"fmt"
	"net/http"
)

func main() {
	//	var dst []byte
	//	base64.StdEncoding.Decode(dst, []byte("jKmSAvMOge8sHR2mDUeAylUanU5GpBtukcRIy2ruucdjJFl+mqM7VOiHISWgB5Li+/5owccaZ4Owob0iSh9EHw=="))
	/*
		by1, _ := base64.StdEncoding.DecodeString("jKmSAvMOge8sHR2mDUeAylUanU5GpBtukcRIy2ruucdjJFl+mqM7VOiHISWgB5Li+/5owccaZ4Owob0iSh9EHw==")
		by, _ := crypt_b.AesDecrypt(by1, []byte("CWELWRREW4567i1o"))
		fmt.Println(string(by))

		aby1, _ := crypt_b.AesEncrypt([]byte(`{"token":"8dd2afbbb2e0c3b614858ec72d93b51b"}`), []byte("CWELWRREW4567i1o"))

		astr1 := base64.StdEncoding.EncodeToString(aby1)
		fmt.Println(astr1) */
	http.HandleFunc("/hello", SayHello)
	http.ListenAndServe("0.0.0.0:8001", nil)
}
func SayHello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("访问")
	//w.Write([]byte("hello"))
	w.Header().Set("Location", "http://127.0.0.1:8001/hello")
}
