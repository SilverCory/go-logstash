# go-logstash
**Ez quick way of combining multiple log files into one, from multiple remote instances.**

## CAUTION
Be careful, there is an unresolved issue where if you use a folder as a file it may overwrite the file.

## Usage
Run the program with the authkey flag set to a very secure password.

The following command will demonstrate usage.
```bash
curl -X POST \
  http://localhost:9090/[folder1]/[folder2]/[logfile] \
  -H 'auth: [authkey]' \
  -d '[logentry]'
 ```
 output file will be `$(pwd)/[folder1]/[folder2]/[logfile]`

If you can't figure this out well, try harder.

Responses are very much dependent on http status codes, 20x is good, anything else is bad. If it 500's an error message will be supplied.
We're not checking or setting content-type's, you can supply raw text to the endpoint and it will append \n and write.

## Building and installation.
You need go dep to install this.
- `git clone https://github.com/SilverCory/go-logstash.git`
- `cd go-logstash`
- `export GOPATH="$(pwd)"`
- `cd src/github.com/SilverCory/go-logstash/`
- `dep ensure` This may take a while.
- `go build main.go`
- `cp main ~/logstash`
- `cd ~`
- `./logstash -authkey "kekeman1234dfjgisiofdjgiosdfjgiojsdfjudghsdrhgkujsgsjfdfasdjl"`

## Enjoy.
