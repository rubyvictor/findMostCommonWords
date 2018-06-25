# How to run this server on your local machine

Before cloning this, please take time to set up your Go workspace and the correct $GOPATH environment variable. Visit this link for a guide:

[GOLANG Docs](https://golang.org/doc/code.html#next)

Once you have set up your GOPATH, clone this repo with the command:

`git clone https://github.com/rubyvictor/findMostCommonWords.git`

Next:

`$ cd $GOPATH/src/github.com/yourUserName/findMostCommonWords`:

`$ go install`

`$ go run findMostCommon.go message.go`

## View in your browser:

Home page URL:  ```http://localhost:3000```

Form page URL:  ```http://localhost:3000/text```