package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Body struct {
	Money      int64  `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int64  `json:"candyCount"`
}

type Response struct {
	Thanks string `json:"thanks,omitempty"`
	Change int64  `json:"change,omitempty"`
	Err    string `json:"error,omitempty"`
}

func getBody() (ret Body) {
	m := flag.Int64("m", 0, "money")
	c := flag.Int64("c", 0, "count")
	k := flag.String("k", "", "type")

	flag.Parse()

	ret = Body{
		Money:      *m,
		CandyType:  *k,
		CandyCount: *c,
	}
	return
}

func main() {
	caCertPath := "security/minica.pem"          // Корневой сертификат CA
	clientCertPath := "security/client/cert.pem" // Клиентский сертификат
	clientKeyPath := "security/client/key.pem"   // Закрытый ключ клиента

	caCert, err := os.ReadFile(caCertPath)
	if err != nil {
		log.Fatalf("Root Certificate err: %v", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	clientCert, err := tls.LoadX509KeyPair(clientCertPath, clientKeyPath)
	if err != nil {
		log.Fatalf("Client Certificate err: %v", err)
	}

	tlsConfig := &tls.Config{
		RootCAs:      caCertPool,
		Certificates: []tls.Certificate{clientCert},
		// InsecureSkipVerify: true,
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}
    params := getBody()
	request, err := json.Marshal(params)
   
    if err != nil {
        log.Fatalf("Err creating a body of request:%v", err)
    }

	resp, err := client.Post("https://candy.tld:8889/buy_candy", "application/json", bytes.NewBuffer(request))
	if err != nil {
		log.Fatalf("Err exec request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

    log.Println(string(respBody))
    var ResponseData Response
    if err = json.Unmarshal(respBody, &ResponseData); err != nil {
        log.Fatal(err)
    }

    if ResponseData.Err == "" {
        fmt.Println(ResponseData.Thanks + " Your change is " + strconv.Itoa(int(ResponseData.Change)))
    } else {
        fmt.Println(ResponseData.Err)
    }
}
