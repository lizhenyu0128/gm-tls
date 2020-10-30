gm-tls

## 功能

* 基于Golang（基于国密算法的TLS / SSL代码库）的GM TLS / SSL  基于了Hyperledger-TWGC/ccs-gm库进行修改，实现tls的客户端和gmssl的服务端进行通讯不报解密失败错误

## 使用
tls.Config结构体中GMSupport变量选项支持打开国密
    
## 修改内容

修改了Hyperledger-TWGC/ccs-gm库的加密预主秘钥方法，在sm2enc.go中修改了
func Encrypt(rand io.Reader, key *PublicKey, msg []byte) (cipher []byte, err error) {
	x, y, c2, c3, err := doEncrypt(rand, key, msg)
	if err != nil {
		return nil, err
	}
	enkey := EncryptClientKey{
		X:    x,
		Y:    y,
		Hash: c3,
		Data: c2,
	}
	cipher, err = asn1EncryptClientKey(enkey)
	return
}

对加密结果进行x,y,c2,c3进行asn.1编码返回
原方法是直接返回 c1||c2||c3

