CURLFS is a lightweight, cross-platform file server written in Go. It allows users to upload and download files via HTTP or HTTPS using simple curl commands. The server automatically serves over HTTPS if TLS certificates are present, falling back to HTTP otherwise.

🚀 Features
	•	File Serving: Serve files from the current directory.
	•	File Uploading: Upload files via POST requests.
	•	HTTPS Support: Automatically uses HTTPS if cert.pem and key.pem are available.
	•	HTTP Fallback: Falls back to HTTP if TLS certificates are missing.
	•	Cross-Platform: Build for multiple operating systems using the provided Makefile.

🛠️ Installation

### Prerequisites
	•	Go installed on your system.

### Build from Source
```sh
git clone https://github.com/theenawman/curlfs.git
cd curlfs
go build -o curlfs main.go
```

### Cross-Compile for Multiple Platforms

Use the provided Makefile to build for various platforms:
```sh
make build
```
Binaries will be placed in the build/ directory.

🔒 HTTPS Configuration

To enable HTTPS:
	1. ```openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes```
 	2. Place your cert.pem and key.pem files in the same directory as the executable.
	3. Start the server ```./curlfs```

The server will listen on port 8443 for HTTPS. If the certificates are not found, it will default to HTTP on port 8080.

📂 Usage

Start the Server

./curlfs

### Upload a File

Over HTTPS
```
curl -k -F "file=@yourfile.txt" https://localhost:8443/
```
Over HTTP
```
curl -F "file=@yourfile.txt" http://localhost:8080/
```
### Download a File

Over HTTPS
```
curl -k https://localhost:8443/yourfile.txt -O
```
Over HTTP
```
curl http://localhost:8080/yourfile.txt -O
```
Note: The -k flag is used to allow connections to SSL sites without certs signed by a trusted CA.

📌 Notes
	•	All files are served and saved in the same directory the server runs from.
	•	Maximum upload size is 10MB (modifiable in the code).
	•	Ensure port 8443 or 8080 is open and not blocked by firewalls.

🛠️ Customization

You can change:
	•	Default ports (httpsPort, httpPort)
	•	Upload size limits
	•	Directory path (directoryPath)
