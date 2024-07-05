/github-repos/personel/gorestclient/malformcerts (main)
$ openssl genrsa -out privkey.pem 2048

/github-repos/personel/gorestclient/malformcerts (main)
$ ls
cert.pem  key.pem  privkey.pem

/github-repos/personel/gorestclient/malformcerts (main)
$ openssl req -new -key privkey.pem -out cert.csr \
  -subj "/C=US/ST=State/L=City/O=Organization/OU=Department/CN=example.com"
req: subject name is expected to be in the format /type0=value0/type1=value1/type2=... where characters may be escaped by \. This name is not in that format: 'C:/Program Files/Git/C=US/ST=State/L=City/O=Organization/OU=Department/CN=example.com'

/github-repos/personel/gorestclient/malformcerts (main)
$ openssl req -new -key privkey.pem -out cert.csr -subj "/C=US/ST=State/L=City/O=Organization/OU=Department/CN=example.com"
req: subject name is expected to be in the format /type0=value0/type1=value1/type2=... where characters may be escaped by \. This name is not in that format: 'C:/Program Files/Git/C=US/ST=State/L=City/O=Organization/OU=Department/CN=example.com'

/github-repos/personel/gorestclient/malformcerts (main)
$ openssl req -new -key privkey.pem -out cert.csr -subj "//C=US//ST=State//L=City//O=Organization//OU=Department//CN=example.com"
req warning: Skipping unknown subject name attribute "/C"
req warning: Skipping unknown subject name attribute "/ST"
req warning: Skipping unknown subject name attribute "/L"
req warning: Skipping unknown subject name attribute "/O"
req warning: Skipping unknown subject name attribute "/OU"
req warning: Skipping unknown subject name attribute "/CN"

/github-repos/personel/gorestclient/malformcerts (main)
$ ls
cert.csr  cert.pem  key.pem  privkey.pem

/github-repos/personel/gorestclient/malformcerts (main)
$ openssl req -x509 -key privkey.pem -in cert.csr -out cert.pem -days 365
Warning: No -copy_extensions given; ignoring any extensions in the request

/github-repos/personel/gorestclient/malformcerts (main)
$ ls
cert.csr  cert.pem  malcert.pem  malkey.pem  privkey.pem
