# simplewebserver
Simple web server for serving static files locally.

## Usage

```
simplewebserver public /mnt/t/simplewebserver/https_cert /mnt/t/simplewebserver/https_key /mnt/t/torrents
```

```
rm -r /mnt/t/simplewebserver
mkdir /mnt/t/simplewebserver
openssl req -x509 -newkey rsa:4096 -noenc -keyout /mnt/t/simplewebserver/https_key -out /mnt/t/simplewebserver/https_cert
```

```
openssl pkey -in /mnt/t/simplewebserver/https_key -pubout -outform DER | sha256sum
```
