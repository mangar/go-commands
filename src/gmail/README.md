# Gmail

## Development


Into ```go-commands```

    export GOPATH=`pwd`

    
and then...

    
    go get gopkg.in/kyokomi/emoji.v1



... export AWS_KEY and AWS_SECRET with full access on S3

    export GMAIL_EMAIL="YOUR_EMAIL_HERE"
    export GMAIL_PWD="YOUR_PASSWORD_HERE"


## Running


__Sending an email__

... to create a bucket, go into ```gmail``` after the __Development__ steps and run:

    go run main.go send TO_EMAIL, SUBJECT, CONTENT


## Generating 

    go build main.go
    mv main gmail


