# GoPhish

Created by DeadVCR (http://deadvcr.com)

**Description**
A login phishing tool created in Go.

# Disclamer
Using this without consent from both ends is illegal. I am not responsible or liable for any damage or misuse caused by this program. Use for educational purposes only.

# Usage
Simply build using the Go compiler
`go build main.go`
Then execute it.
```
./main
or
main.exe
```

Choose the the template you want from the list, enter the bind IP and that's it! Provided you aren't a complete brainlet and typed in simple numbers correctly, the webserver should be running.

All captured logins are stored in nice little .json files in the 'pwned' directory.

Demo video: https://youtu.be/9UoGnnm_aiI

# Configuration
You can now change the default vaules. Simply open the `defaults.json` file and update accordingly.
Example:
```
{
    "Redirect": "amazon.com",
    "BindIP": "127.0.0.1:8000"
}
```
