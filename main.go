package main

import "net/http"

func main() {
	http.HandleFunc("/", helloworld)
	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		panic(err)
	}
}

func helloworld(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Hello World vgreen!!!"))
}
