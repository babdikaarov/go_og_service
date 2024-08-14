# Get started

TODO:

-  doc pending to write

## deprecated!! Get started

works on commit - ec36c64

```zsh
# Ensure Go is installed on your machine

git clone https://github.com/babdikaarov/go_og_service.git .

cd go_og_service

# Run the service locally
go run main.go

# Test case
curl "http://localhost:8080/og?url=asd,asdfsdf"

# Expected response:
[
    {
        "title": "Error",
        "description": "Try to check the URL OG is available or assign manually",
        "image": "null",
        "original_url": "asd"
    },
    {
        "title": "Error",
        "description": "Try to check the URL OG is available or assign manually",
        "image": "null",
        "original_url": "asdfsdf"
    }
]

curl "https://youtu.be/0RKpf3rK57I?si=pdehNIEfRnj2sB3q"

# Expected response:
[
    {
        "title": "Hugo in 100 Seconds",
        "description": "Hugo is an extremely fast static site generator for building websites with markdown. It is written in the Go programming language and provides a large collection of features.",
        "image": "https://i.ytimg.com/vi/0RKpf3rK57I/maxresdefault.jpg",
        "original_url": "https://youtu.be/0RKpf3rK57I?si=pdehNIEfRnj2sB3q"
    }
]

```

## demo api

```go
url := "https://go-og-service.onrender.com"
getPath := "/og"
params := "url" // Multiple values should be comma-separated
example := `https://go-og-service.onrender.com/og?url=https://youtu.be/0RKpf3rK57I?si=pdehNIEfRnj2sB3q,http://localhost:8080/og?url=asd,asdfsdf`
```
