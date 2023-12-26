# [Tempsy Client](https://tempsy.afifurrohman.my.id)

> Tempsy Client is a web client for [Tempsy API](https://github.com/afifurrohman-id/tempsy.git)

## Usage

## Requirements

- [x] Git Bash for Windows (version >= 2.41.0)
  > only need if you're using windows OS
- [ ] Docker (version >= 24.0.x)
  > optional, only if you want to build image)
- [x] Go (version >= 1.21.x)

## Installation

- Clone this repository

```sh
git clone https://github.com/afifurrohman-id/tempsy-client.git
```

- Go to project directory

```sh
cd tempsy-client
```

- Insert Variable to `.env` file

```sh

cat <<EOENV > configs/.env

OAUTH2_CONFIG={"clientId": "EXAMPLE_CLIENT_ID","clientSecret": "EXAMPLE_SECRET","callbackUrl": "https://example.com/auth","scopes": ["https://www.googleapis.com/auth/userinfo.profile"]}
APP_ENV=testing
API_SERVER_URI=https://api.example.com
PORT=8080

EOENV

```

- Download dependencies

```sh
go mod tidy
```

## Run

- Run the app

```sh
go run main.go
```

- Build

```sh
go build -o tempsy-client main.go
```

- Build Image

```sh
docker build -f build/package/Containerfile -t tempsy .
```
