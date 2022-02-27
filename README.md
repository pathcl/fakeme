# fakeme

fakeme is a simple Go program that makes HTTP request constantly in order to generate random HTTP/DNS traffic noise.

> fakeme is a project inspired by [Whisperer](https://github.com/daniarlert/Whisperer).

## Example

```bash
fakeme -v -d 3s
```

## Install

### Go

```bash
go get github.com/pathcl/fakeme
```

### Cloning the repository

```bash
# First, clone the repository
git clone https://github.com/pathcl/fakeme

# Then navigate into the fakeme directory
cd fakeme

# Run
go run main.go
```

## Docker

### Pulling Image

To use fakeme as a Docker container you can pull the image with the following command:

```bash
docker image pull pathcl/fakeme
```

> Note that the image ```pathcl/fakeme``` uses the urls file from this repository. So it is not a valid option if you want to customize the URLs that fakeme is going to visit.

### Building Image

```bash
# First, clone the repository
git clone https://github.com/pathcl/fakeme

# Then navigate into the fakeme directory
cd fakeme

# Modify the urls.txt file if you want
vim urls.txt

# Build the Docker Image from the Dockerfile inside the repository
docker image build -t fakeme .

# Run
docker container run fakeme
```

## Options

fakeme can accept a number of command line arguments:

```text
$ fakeme --help
fakeme makes HTTP request constantly in order to generate random HTTP/DNS traffic noise.

Usage:
  fakeme [flags]

Flags:
  -a, --agent string       user agent (default "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:67.0) Gecko/20100101 Firefox/67.0")
      --debug              prints error messages
  -d, --delay duration     delay between requests (default 1s)
  -g, --goroutines int     number of goroutines (default 1)
  -h, --help               help for fakeme
  -p, --proxy string       proxy URL
  -r, --random             random delay between requests
  -t, --timeout duration   max time to wait for a response before canceling the request (default 3s)
      --urls string        simple .txt file with URL's to visit (default "./urls.txt")
  -v, --verbose            enables verbose mode
```

## URLs file

This file is from which fakeme will extract the different URLs that will be visiting.

> You can see an example of how this file should be [here](https://github.com/pathcl/fakeme/blob/master/urls.txt).

## Help is always welcome!

If you know about anything else I can improve or add please, don't hesitate to let me know!
