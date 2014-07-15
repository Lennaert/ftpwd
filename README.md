# Ftpwd
Ftpwd is a small Go tool that can recover your FTP password from your saved ftp
connections. I had saved some FTP favourites in my FTP client, but I forgot the password.

Since FTP is a very unsecure protocol, I just could have sniffed the traffic
with Wireshark or something. However I also like to learn Golang so I decided
to create a very little Go app that acts like an FTP server.

![Alt text](/screenshots/01-original.png?raw=true)
![Alt text](/screenshots/02-replacehost.png?raw=true)
![Alt text](/screenshots/03-result.png?raw=true)


## Usage:
	Run ftpwd. It listens on port 2121
	In your FTP client, change the "server/host" to 127.0.0.1 and the "port" to 2121.
	