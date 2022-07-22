package nitriding

import (
	"crypto/rand"
	"errors"
	"testing"
	"time"
)

func TestSbKeyRandomness(t *testing.T) {
	k1, _ := newSbKey()
	k2, _ := newSbKey()

	// It's notoriously difficult to test if something is truly random.  Here,
	// we simply make sure that two subsequently generated key pairs are not
	// identical.  That's a low bar to pass but better than nothing.
	if k1.nonce == k2.nonce {
		t.Errorf("Nonces of two separate secretbox keys are identical.")
	}
	if k1.key == k2.key {
		t.Errorf("Keys of two separate secretbox keys are identical.")
	}
}

func TestSbKeySerialization(t *testing.T) {
	k1, _ := newSbKey()
	k2, _ := newSbKeyFromBytes(k1.Bytes())

	if k1.nonce != k2.nonce || k1.key != k2.key {
		t.Errorf("Key no longer identical after encoding and decoding.")
	}

	_, err := newSbKeyFromBytes([]byte{})
	if err == nil {
		t.Errorf("Expected an error because no bytes were provided.")
	}
}

func TestNonce(t *testing.T) {
	n1, _ := newNonce()
	n2, _ := newNonce()
	if n1 == n2 {
		t.Errorf("Two separately generated nonces are identical.")
	}
	if n1.B64() == n2.B64() {
		t.Errorf("Two separately generated Base64-encoded nonces are identical.")
	}
}

func TestErrors(t *testing.T) {
	cryptoRead = func(b []byte) (n int, err error) {
		return 0, errors.New("not enough randomness")
	}
	defer func() {
		cryptoRead = rand.Read
	}()

	if _, err := newSbKey(); err == nil {
		t.Errorf("Failed to return error")
	}
	if _, err := newNonce(); err == nil {
		t.Errorf("Failed to return error")
	}
}

func attDocNoFieldsTime() time.Time {
	t, _ := time.Parse(time.RFC3339, "2022-07-18T21:00:00Z")
	return t
}

var attDocNoFields string = `
hEShATgioFkQwKlpbW9kdWxlX2lkeCdpLTA4MDk4NDk3MTBiZjFiNjFiLWVuYzAxODIxMmY0ZWE3YmE
1Y2JmZGlnZXN0ZlNIQTM4NGl0aW1lc3RhbXAbAAABghL08UhkcGNyc7AAWDAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABWDAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACWDAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAADWDAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAEWDDYqOju6W2Bt3olFBC3qbGAeHZT8SXR28p5aFyT+4hbM14LjRcsmCGoYlFaYDzDOr
IFWDAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGWDAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHWDAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIWDAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAJWDAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAKWDAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAALWDAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAMWDAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAANWDAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAOWDAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAPWDAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABrY2VydGlmaWNhdGVZAn8wggJ7MIICAaADAgECAhAB
ghL06nulywAAAABi1b+6MAoGCCqGSM49BAMDMIGOMQswCQYDVQQGEwJVUzETMBEGA1UECAwKV2FzaGl
uZ3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEPMA0GA1UECgwGQW1hem9uMQwwCgYDVQQLDANBV1MxOTA3Bg
NVBAMMMGktMDgwOTg0OTcxMGJmMWI2MWIudXMtZWFzdC0yLmF3cy5uaXRyby1lbmNsYXZlczAeFw0yM
jA3MTgyMDE2NTVaFw0yMjA3MTgyMzE2NThaMIGTMQswCQYDVQQGEwJVUzETMBEGA1UECAwKV2FzaGlu
Z3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEPMA0GA1UECgwGQW1hem9uMQwwCgYDVQQLDANBV1MxPjA8BgN
VBAMMNWktMDgwOTg0OTcxMGJmMWI2MWItZW5jMDE4MjEyZjRlYTdiYTVjYi51cy1lYXN0LTIuYXdzMH
YwEAYHKoZIzj0CAQYFK4EEACIDYgAE2RBGub+ROH138AA4aWLkldgjY/iqtz6BAFS9pQTD+Ivys7m7F
BM+9gZ3nJFE9qJkvVbZA69hvbRtyrF2dwpuJUl2HCkWGo8g3d3vsxaSHd2vZG2xbz9l72mDbrKRNFLX
ox0wGzAMBgNVHRMBAf8EAjAAMAsGA1UdDwQEAwIGwDAKBggqhkjOPQQDAwNoADBlAjAy+ejsnvnxHII
VSSC5kG0G4J3mJYZJkvui83OvmV/wfvpNZmCMgvyFa70ZyCC2Bz4CMQDxq70/eN7VevXbP9X0w8ZD/S
VFknKAbFVxHnV1MsgYfQqO/wfOz3Trl49MaCEiyeZoY2FidW5kbGWEWQIVMIICETCCAZagAwIBAgIRA
PkxdWgbkK/hHUbMtOTn+FYwCgYIKoZIzj0EAwMwSTELMAkGA1UEBhMCVVMxDzANBgNVBAoMBkFtYXpv
bjEMMAoGA1UECwwDQVdTMRswGQYDVQQDDBJhd3Mubml0cm8tZW5jbGF2ZXMwHhcNMTkxMDI4MTMyODA
1WhcNNDkxMDI4MTQyODA1WjBJMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGQW1hem9uMQwwCgYDVQQLDA
NBV1MxGzAZBgNVBAMMEmF3cy5uaXRyby1lbmNsYXZlczB2MBAGByqGSM49AgEGBSuBBAAiA2IABPwCV
OumCMHzaHDimtqQvkY4MpJzbolL//Zy2YlES1BR5TSksfbb48C8WBoyt7F2Bw7eEtaaP+ohG2bnUs99
0d0JX28TcPQXCEPZ3BABIeTPYwEoCWZEh8l5YoQwTcU/9KNCMEAwDwYDVR0TAQH/BAUwAwEB/zAdBgN
VHQ4EFgQUkCW1DdkFR+eWw5b6cp3PmanfS5YwDgYDVR0PAQH/BAQDAgGGMAoGCCqGSM49BAMDA2kAMG
YCMQCjfy+Rocm9Xue4YnwWmNJVA44fA0P5W2OpYow9OYCVRaEevL8uO1XYru5xtMPWrfMCMQCi85sWB
bJwKKXdS6BptQFuZbT73o/gBh1qUxl/nNr12UO8Yfwr6wPLb+6NIwLz3/ZZAsMwggK/MIICRaADAgEC
AhEAuP/xq2yeQ6e2R7pyxUUjyjAKBggqhkjOPQQDAzBJMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGQW1
hem9uMQwwCgYDVQQLDANBV1MxGzAZBgNVBAMMEmF3cy5uaXRyby1lbmNsYXZlczAeFw0yMjA3MTUwNz
Q4MDhaFw0yMjA4MDQwODQ4MDhaMGQxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZBbWF6b24xDDAKBgNVB
AsMA0FXUzE2MDQGA1UEAwwtNTBhNzJiY2I1NTY0NGE4Ny51cy1lYXN0LTIuYXdzLm5pdHJvLWVuY2xh
dmVzMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEArXj+qxdWCAHUCpxLQXIYVD4lXLGEkki1o8E/cTtbMN
DGtTTP2dACg4V84LL6sPQShd620u6jN7YFQ2Pc6eiUyXoqVBC8eZYguLRWKZGlHegpzLyTFlvtVH2XI
MKNnH5o4HVMIHSMBIGA1UdEwEB/wQIMAYBAf8CAQIwHwYDVR0jBBgwFoAUkCW1DdkFR+eWw5b6cp3Pm
anfS5YwHQYDVR0OBBYEFLJ56rXN1CLr9eI43YPJgKsJs3fYMA4GA1UdDwEB/wQEAwIBhjBsBgNVHR8E
ZTBjMGGgX6BdhltodHRwOi8vYXdzLW5pdHJvLWVuY2xhdmVzLWNybC5zMy5hbWF6b25hd3MuY29tL2N
ybC9hYjQ5NjBjYy03ZDYzLTQyYmQtOWU5Zi01OTMzOGNiNjdmODQuY3JsMAoGCCqGSM49BAMDA2gAMG
UCMQDou5jD4fBOwCQh1J2spOtl0A4cGzoak+j3Zv0lf0oOZcbtYpgYbZdctVbfjcJ70EACMA4dzOJhU
racC7LvgexEA54m9OXjhWtm15pDvNbISPYL+OPCyHfKfNgT1X1FeWs5nFkDGTCCAxUwggKaoAMCAQIC
EGNeRZo3hQR7jYDiHAg5/lUwCgYIKoZIzj0EAwMwZDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBkFtYXp
vbjEMMAoGA1UECwwDQVdTMTYwNAYDVQQDDC01MGE3MmJjYjU1NjQ0YTg3LnVzLWVhc3QtMi5hd3Mubm
l0cm8tZW5jbGF2ZXMwHhcNMjIwNzE4MDYxODEwWhcNMjIwNzI0MDAxODEwWjCBiTE8MDoGA1UEAwwzZ
mUyODdmOGY1ZGI4NTEwMS56b25hbC51cy1lYXN0LTIuYXdzLm5pdHJvLWVuY2xhdmVzMQwwCgYDVQQL
DANBV1MxDzANBgNVBAoMBkFtYXpvbjELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAldBMRAwDgYDVQQHDAd
TZWF0dGxlMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAExNht09kEVN1ohbh2jPoDVRqLYeYN8x18Zf4KQ9
sZYUYjYQyL+yjAdJ57uMskkaYe1BEvy8Q4l5Y7BreP2X9cYLnZSYKZbvjVmiEvSpkTeA8dJEK++dkxB
iuknuqiN/DCo4HqMIHnMBIGA1UdEwEB/wQIMAYBAf8CAQEwHwYDVR0jBBgwFoAUsnnqtc3UIuv14jjd
g8mAqwmzd9gwHQYDVR0OBBYEFOyBxTMuxGryk4jXL7+Xf0FmUJTfMA4GA1UdDwEB/wQEAwIBhjCBgAY
DVR0fBHkwdzB1oHOgcYZvaHR0cDovL2NybC11cy1lYXN0LTItYXdzLW5pdHJvLWVuY2xhdmVzLnMzLn
VzLWVhc3QtMi5hbWF6b25hd3MuY29tL2NybC9jY2ZjNzA5OC0zZDhjLTQ0MWQtYWNmNS0yZGYwYmFhZ
DM0NTMuY3JsMAoGCCqGSM49BAMDA2kAMGYCMQD6Ym3kyGA8sovqy7DlfCErl6b30KqlRj8MRjHwTwX8
R1KqyiIEESsaEm05zHU7ChYCMQC6oTx++35mHVXp0WQYdXAMwHU/aSRDuKGt5qfM8uijIBipTVTUBAl
aPZJczlTq8TBZAoIwggJ+MIICBaADAgECAhUAvibRQZS9C/1gQYKqmFD+cxIXt1MwCgYIKoZIzj0EAw
MwgYkxPDA6BgNVBAMMM2ZlMjg3ZjhmNWRiODUxMDEuem9uYWwudXMtZWFzdC0yLmF3cy5uaXRyby1lb
mNsYXZlczEMMAoGA1UECwwDQVdTMQ8wDQYDVQQKDAZBbWF6b24xCzAJBgNVBAYTAlVTMQswCQYDVQQI
DAJXQTEQMA4GA1UEBwwHU2VhdHRsZTAeFw0yMjA3MTgxNzIyMDVaFw0yMjA3MTkxNzIyMDVaMIGOMQs
wCQYDVQQGEwJVUzETMBEGA1UECAwKV2FzaGluZ3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEPMA0GA1UECg
wGQW1hem9uMQwwCgYDVQQLDANBV1MxOTA3BgNVBAMMMGktMDgwOTg0OTcxMGJmMWI2MWIudXMtZWFzd
C0yLmF3cy5uaXRyby1lbmNsYXZlczB2MBAGByqGSM49AgEGBSuBBAAiA2IABHauNrI7BTIweN+zwPt+
cchEnzuRwHLILTAHh3OTa47tKPrx5siwKIwhkjOvzAN82o4MzgUmqtfQ0yrntfrox2be5qzKx7U26aa
tS5GJR/STHSjtoeKZn5FLMYysMJM00KMmMCQwEgYDVR0TAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBA
MCAgQwCgYIKoZIzj0EAwMDZwAwZAIwKGmJdAp6HiDJ95MEagMsG3apZ0jMf0kL6x4nlPn20HbtmsLh/
7AsXKL6t17GFW9gAjBTNZdXSfIqZ49UEI5A2GrVA5Ix7oFI5/JEGWZvCBXDrgOm45m5PzFP90fn+nwE
N4xqcHVibGljX2tlefZpdXNlcl9kYXRh9mVub25jZfZYYCj0RorPrgq04lgfQBrXW7Bxzjgy2HjgTtA
NOWwsTEhfcF/Z6xLJ2WwOAqW5L0zp45gVfDsupaPGygoEQzFHFnHvwD9aJmyMa7OlrTP7jCiCPMHH7Z
mH57o46I2wuzIN/g==
`

func attDocNoSbTime() time.Time {
	return time.Unix(1658177474, 0)
}

var attDocNoSb string = `
hEShATgioFkQ1alpbW9kdWxlX2lkeCdpLTA4MDk4NDk3MTBiZjFiNjFiLWVuYzAxODIxMzE0NGEwMjI
1YzlmZGlnZXN0ZlNIQTM4NGl0aW1lc3RhbXAbAAABghMUUNJkcGNyc7AAWDAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABWDAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACWDAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAADWDAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAEWDDYqOju6W2Bt3olFBC3qbGAeHZT8SXR28p5aFyT+4hbM14LjRcsmCGoYlFaYDzDOr
IFWDAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGWDAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHWDAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIWDAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAJWDAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAKWDAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAALWDAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAMWDAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAANWDAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAOWDAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAPWDAAAAAAAAAAAAAAAAAAAAAAAAAAA
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABrY2VydGlmaWNhdGVZAoAwggJ8MIICAaADAgECAhAB
ghMUSgIlyQAAAABi1cfCMAoGCCqGSM49BAMDMIGOMQswCQYDVQQGEwJVUzETMBEGA1UECAwKV2FzaGl
uZ3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEPMA0GA1UECgwGQW1hem9uMQwwCgYDVQQLDANBV1MxOTA3Bg
NVBAMMMGktMDgwOTg0OTcxMGJmMWI2MWIudXMtZWFzdC0yLmF3cy5uaXRyby1lbmNsYXZlczAeFw0yM
jA3MTgyMDUxMTFaFw0yMjA3MTgyMzUxMTRaMIGTMQswCQYDVQQGEwJVUzETMBEGA1UECAwKV2FzaGlu
Z3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEPMA0GA1UECgwGQW1hem9uMQwwCgYDVQQLDANBV1MxPjA8BgN
VBAMMNWktMDgwOTg0OTcxMGJmMWI2MWItZW5jMDE4MjEzMTQ0YTAyMjVjOS51cy1lYXN0LTIuYXdzMH
YwEAYHKoZIzj0CAQYFK4EEACIDYgAEuv+y9TIepZu/s8XUSYXHSCj+5x3riR9MvinDCz0kUOAGvqNEN
JYUCvhi1Rw9AcAH49AJEpE4Ehg32WPafKvmc6WFGG4+9Q57GtqkCTQ7m5vqFTB4nOPrHhEj9XezTcuD
ox0wGzAMBgNVHRMBAf8EAjAAMAsGA1UdDwQEAwIGwDAKBggqhkjOPQQDAwNpADBmAjEAmrHJFdWFL8G
0A4JtlUElG9qUSJQBE9g6woNh/2Jh6QUEVdIkoqsReVbJa4PP4rT5AjEA9vWFVWSgPgQPpf2EVzDuO4
8Cyl05mBwy/hQM4+j4z2HD6Z/GOLFLMqEvlyL1hJ8laGNhYnVuZGxlhFkCFTCCAhEwggGWoAMCAQICE
QD5MXVoG5Cv4R1GzLTk5/hWMAoGCCqGSM49BAMDMEkxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZBbWF6
b24xDDAKBgNVBAsMA0FXUzEbMBkGA1UEAwwSYXdzLm5pdHJvLWVuY2xhdmVzMB4XDTE5MTAyODEzMjg
wNVoXDTQ5MTAyODE0MjgwNVowSTELMAkGA1UEBhMCVVMxDzANBgNVBAoMBkFtYXpvbjEMMAoGA1UECw
wDQVdTMRswGQYDVQQDDBJhd3Mubml0cm8tZW5jbGF2ZXMwdjAQBgcqhkjOPQIBBgUrgQQAIgNiAAT8A
lTrpgjB82hw4prakL5GODKSc26JS//2ctmJREtQUeU0pLH22+PAvFgaMrexdgcO3hLWmj/qIRtm51LP
fdHdCV9vE3D0FwhD2dwQASHkz2MBKAlmRIfJeWKEME3FP/SjQjBAMA8GA1UdEwEB/wQFMAMBAf8wHQY
DVR0OBBYEFJAltQ3ZBUfnlsOW+nKdz5mp30uWMA4GA1UdDwEB/wQEAwIBhjAKBggqhkjOPQQDAwNpAD
BmAjEAo38vkaHJvV7nuGJ8FpjSVQOOHwND+VtjqWKMPTmAlUWhHry/LjtV2K7ucbTD1q3zAjEAovObF
gWycCil3UugabUBbmW0+96P4AYdalMZf5za9dlDvGH8K+sDy2/ujSMC89/2WQLDMIICvzCCAkWgAwIB
AgIRALj/8atsnkOntke6csVFI8owCgYIKoZIzj0EAwMwSTELMAkGA1UEBhMCVVMxDzANBgNVBAoMBkF
tYXpvbjEMMAoGA1UECwwDQVdTMRswGQYDVQQDDBJhd3Mubml0cm8tZW5jbGF2ZXMwHhcNMjIwNzE1MD
c0ODA4WhcNMjIwODA0MDg0ODA4WjBkMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGQW1hem9uMQwwCgYDV
QQLDANBV1MxNjA0BgNVBAMMLTUwYTcyYmNiNTU2NDRhODcudXMtZWFzdC0yLmF3cy5uaXRyby1lbmNs
YXZlczB2MBAGByqGSM49AgEGBSuBBAAiA2IABAK14/qsXVggB1AqcS0FyGFQ+JVyxhJJItaPBP3E7Wz
DQxrU0z9nQAoOFfOCy+rD0EoXettLuoze2BUNj3OnolMl6KlQQvHmWILi0VimRpR3oKcy8kxZb7VR9l
yDCjZx+aOB1TCB0jASBgNVHRMBAf8ECDAGAQH/AgECMB8GA1UdIwQYMBaAFJAltQ3ZBUfnlsOW+nKdz
5mp30uWMB0GA1UdDgQWBBSyeeq1zdQi6/XiON2DyYCrCbN32DAOBgNVHQ8BAf8EBAMCAYYwbAYDVR0f
BGUwYzBhoF+gXYZbaHR0cDovL2F3cy1uaXRyby1lbmNsYXZlcy1jcmwuczMuYW1hem9uYXdzLmNvbS9
jcmwvYWI0OTYwY2MtN2Q2My00MmJkLTllOWYtNTkzMzhjYjY3Zjg0LmNybDAKBggqhkjOPQQDAwNoAD
BlAjEA6LuYw+HwTsAkIdSdrKTrZdAOHBs6GpPo92b9JX9KDmXG7WKYGG2XXLVW343Ce9BAAjAOHcziY
VK2nAuy74HsRAOeJvTl44VrZteaQ7zWyEj2C/jjwsh3ynzYE9V9RXlrOZxZAxkwggMVMIICmqADAgEC
AhBjXkWaN4UEe42A4hwIOf5VMAoGCCqGSM49BAMDMGQxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZBbWF
6b24xDDAKBgNVBAsMA0FXUzE2MDQGA1UEAwwtNTBhNzJiY2I1NTY0NGE4Ny51cy1lYXN0LTIuYXdzLm
5pdHJvLWVuY2xhdmVzMB4XDTIyMDcxODA2MTgxMFoXDTIyMDcyNDAwMTgxMFowgYkxPDA6BgNVBAMMM
2ZlMjg3ZjhmNWRiODUxMDEuem9uYWwudXMtZWFzdC0yLmF3cy5uaXRyby1lbmNsYXZlczEMMAoGA1UE
CwwDQVdTMQ8wDQYDVQQKDAZBbWF6b24xCzAJBgNVBAYTAlVTMQswCQYDVQQIDAJXQTEQMA4GA1UEBww
HU2VhdHRsZTB2MBAGByqGSM49AgEGBSuBBAAiA2IABMTYbdPZBFTdaIW4doz6A1Uai2HmDfMdfGX+Ck
PbGWFGI2EMi/sowHSee7jLJJGmHtQRL8vEOJeWOwa3j9l/XGC52UmCmW741ZohL0qZE3gPHSRCvvnZM
QYrpJ7qojfwwqOB6jCB5zASBgNVHRMBAf8ECDAGAQH/AgEBMB8GA1UdIwQYMBaAFLJ56rXN1CLr9eI4
3YPJgKsJs3fYMB0GA1UdDgQWBBTsgcUzLsRq8pOI1y+/l39BZlCU3zAOBgNVHQ8BAf8EBAMCAYYwgYA
GA1UdHwR5MHcwdaBzoHGGb2h0dHA6Ly9jcmwtdXMtZWFzdC0yLWF3cy1uaXRyby1lbmNsYXZlcy5zMy
51cy1lYXN0LTIuYW1hem9uYXdzLmNvbS9jcmwvY2NmYzcwOTgtM2Q4Yy00NDFkLWFjZjUtMmRmMGJhY
WQzNDUzLmNybDAKBggqhkjOPQQDAwNpADBmAjEA+mJt5MhgPLKL6suw5XwhK5em99CqpUY/DEYx8E8F
/EdSqsoiBBErGhJtOcx1OwoWAjEAuqE8fvt+Zh1V6dFkGHVwDMB1P2kkQ7ihreanzPLooyAYqU1U1AQ
JWj2SXM5U6vEwWQKCMIICfjCCAgWgAwIBAgIVAL4m0UGUvQv9YEGCqphQ/nMSF7dTMAoGCCqGSM49BA
MDMIGJMTwwOgYDVQQDDDNmZTI4N2Y4ZjVkYjg1MTAxLnpvbmFsLnVzLWVhc3QtMi5hd3Mubml0cm8tZ
W5jbGF2ZXMxDDAKBgNVBAsMA0FXUzEPMA0GA1UECgwGQW1hem9uMQswCQYDVQQGEwJVUzELMAkGA1UE
CAwCV0ExEDAOBgNVBAcMB1NlYXR0bGUwHhcNMjIwNzE4MTcyMjA1WhcNMjIwNzE5MTcyMjA1WjCBjjE
LMAkGA1UEBhMCVVMxEzARBgNVBAgMCldhc2hpbmd0b24xEDAOBgNVBAcMB1NlYXR0bGUxDzANBgNVBA
oMBkFtYXpvbjEMMAoGA1UECwwDQVdTMTkwNwYDVQQDDDBpLTA4MDk4NDk3MTBiZjFiNjFiLnVzLWVhc
3QtMi5hd3Mubml0cm8tZW5jbGF2ZXMwdjAQBgcqhkjOPQIBBgUrgQQAIgNiAAR2rjayOwUyMHjfs8D7
fnHIRJ87kcByyC0wB4dzk2uO7Sj68ebIsCiMIZIzr8wDfNqODM4FJqrX0NMq57X66Mdm3uasyse1Num
mrUuRiUf0kx0o7aHimZ+RSzGMrDCTNNCjJjAkMBIGA1UdEwEB/wQIMAYBAf8CAQAwDgYDVR0PAQH/BA
QDAgIEMAoGCCqGSM49BAMDA2cAMGQCMChpiXQKeh4gyfeTBGoDLBt2qWdIzH9JC+seJ5T59tB27ZrC4
f+wLFyi+rdexhVvYAIwUzWXV0nyKmePVBCOQNhq1QOSMe6BSOfyRBlmbwgVw64DpuOZuT8xT/dH5/p8
BDeManB1YmxpY19rZXn2aXVzZXJfZGF0YfZlbm9uY2VU77Ofqd0vUmm6t89uu4vRtxpHXmZYYNAVRyS
iloDHhy5KwRUC5Z/lxLndg8AP91CDx6QUXbPdUS1mvig5z8fzPnrJJ70dIA3cDWbPEjfJwso3nOCUMx
WAjquh0RAQJ4rf89PhMTmfi6COnO6+J1FDkNak7ZFG1A==
`