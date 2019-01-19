# GoPhish

Created by [DeadVCR](http://deadvcr.com)

A login phishing tool created in Go.

## Disclamer

Using this without consent from both ends is illegal. I am not responsible or liable for any damage or misuse caused by this program. Use for educational purposes only.

## Usage

Simply build using the Go compiler
`go build main.go`
Then execute it.

1. Install
    1. `github.com/89apt89/GoPhish`
2. Build the program
    1. `cd $GOPATH/src/github.com/89apt89/GoPhish`
    2. `go build`
3. Run
    1. `./GoPhish` (in bash) or `GoPhish.exe` (in cmd prompt)

Choose the the template you want from the list, enter the bind IP and that's it! Provided you aren't a complete brainlet and typed in simple numbers correctly, the webserver should be running.

All captured logins are stored in nice little .json files in the 'pwned' directory.

## Configuration

You can now change the default vaules. Simply open the `defaults.json` file and update accordingly.
Example:

```json
{
    "Redirect": "amazon.com",
    "BindIP": "127.0.0.1",
    "BindPort": "8000"
}
```
