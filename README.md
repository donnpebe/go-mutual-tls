# go-mutual-tls

This is an example of doing mutual tls authentication in go.

To setup certs for the example, install [certstrap](https://github.com/square/certstrap) via `go get -u https://github.com/square/certstrap` and run `./generate-certs.sh`

## valid server + client
This uses server and client keypairs signed by the valid CA, YoloSwagginsCA.

* server

  ```
  ➜  go-mutual-tls go run server/main.go --ca-crt=valid/YoloSwagginsCA.crt --crt=valid/server.crt --key=valid/server.key
  2016/08/15 23:07:09 https server crt: valid/server.crt
  2016/08/15 23:07:09 https server key: valid/server.key
  2016/08/15 23:07:09 ca crt: valid/YoloSwagginsCA.crt
  2016/08/15 23:07:09 starting server at 127.0.0.1:8443
  2016/08/15 23:07:10 client "client" ["127.0.0.1"]: signed by YoloSwagginsCA
  ```

* client

  ```
  ➜  go-mutual-tls go run client/main.go --ca-crt=valid/YoloSwagginsCA.crt --crt=valid/client.crt --key=valid/client.key
  2016/08/15 23:07:10 https client crt: valid/client.crt
  2016/08/15 23:07:10 https client key: valid/client.key
  2016/08/15 23:07:10 ca crt: valid/YoloSwagginsCA.crt
  2016/08/15 23:07:10 response: "hello world"
  ```

## valid server + invalid client
This uses the valid server keypair but client keypair signed by the invalid CA, failing to authenticate.

* server

  ```
  ➜  go-mutual-tls go run server/main.go --ca-crt=valid/YoloSwagginsCA.crt --crt=valid/server.crt --key=valid/server.key      
  2016/08/15 23:10:10 https server crt: valid/server.crt
  2016/08/15 23:10:10 https server key: valid/server.key
  2016/08/15 23:10:10 ca crt: valid/YoloSwagginsCA.crt
  2016/08/15 23:10:10 starting server at 127.0.0.1:8443
  2016/08/15 23:10:13 http: TLS handshake error from 127.0.0.1:51794: tls: client didn't provide a certificate
  ```

* client
  ```
  ➜  go-mutual-tls go run client/main.go --ca-crt=valid/YoloSwagginsCA.crt --crt=invalid/client-bad.crt --key=invalid/client-bad.key
  2016/08/15 23:10:12 https client crt: invalid/client-bad.crt
  2016/08/15 23:10:12 https client key: invalid/client-bad.key
  2016/08/15 23:10:12 ca crt: valid/YoloSwagginsCA.crt
  2016/08/15 23:10:13 failed to get http response: Get https://127.0.0.1:8443: remote error: tls: bad certificate
  ```
