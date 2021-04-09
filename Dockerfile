FROM golang
WORKDIR /src/
COPY . /src/.
COPY go.mod /src/go.mod


CMD ["go", "run", "/src/main.go"]