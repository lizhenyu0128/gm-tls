// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"github.com/lizhenyu0128/gm-tls/src/tls"
	"github.com/lizhenyu0128/gm-tls/src/x509"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {

	root1, _ := ioutil.ReadFile("/Users/lizhenyu/go/src/gm-tls/src/sm2.oca (2).pem")
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(root1)
	c := &tls.Config{InsecureSkipVerify: true,
		GMSupport: &tls.GMSupport{},
	}
	c.CipherSuites = append(c.CipherSuites, tls.GMTLS_SM2_WITH_SM4_SM3)
	c.MinVersion = c.GMSupport.GetVersion() //only gm support VersionGMSSL
	c.RootCAs = pool
	//tr, err := tls.Dial("tcp", "sm2test.ovssl.cn:https",
	//	c)

	var netTransport = &http.Transport{
		DialTLSContext: func(ctx context.Context, network, address string) (net.Conn, error) {
			var dial, err = tls.Dial(network, address, c)
			if err != nil {
				fmt.Println(err.Error())
			}
			return dial, err
		},
		TLSHandshakeTimeout: 100 * time.Second,
	}
	var netClient = &http.Client{
		Timeout:   time.Second * 100,
		Transport: netTransport,
	}
	resp, err := netClient.Get("https://sm2test.ovssl.cn")

	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
