# Get started

```zsh

# expected go installed on machine

git clone https://github.com/babdikaarov/go_og_service.git .

cd go_og_service

# run locally

go run main.go

# test case

curl "http://localhost:8080/og?url=asd,asdfsdf"

# returns:
[
    {
    "title":"Error",
    "description":"Try to check the URL OG is available or assign manually",
    "image":"null",
    "original_url":"asd"
    },
    {"title":"Error",
    "description":"Try to check the URL OG is available or assign manually",
    "image":"null",
    "original_url":"asdfsdf"
    }
]

curl "https://youtu.be/0RKpf3rK57I?si=pdehNIEfRnj2sB3q"

# returns

[
   {
      "title": "Hugo in 100 Seconds",
      "description": "Hugo is an extremely fast static site generator for building websites with markdown. It is written in the Go programming language and provides a large collec...",
      "image": "https://i.ytimg.com/vi/0RKpf3rK57I/maxresdefault.jpg",
      "original_url": "https://youtu.be/0RKpf3rK57I?si=pdehNIEfRnj2sB3q"
   }
]


```
