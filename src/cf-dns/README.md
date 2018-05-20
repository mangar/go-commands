# Cloudflare - DNS

## Development


Into ```go-commands```

    export GOPATH=`pwd`

    
and then...

    
    go get gopkg.in/kyokomi/emoji.v1



... export AWS_KEY and AWS_SECRET with full access on S3

    export CF_EMAIL="YOUR_EMAIL_HERE"
    export CF_KEY="YOUR_API_KEY_HERE"


## Running


__Creating an entry__

... to create a bucket, go into ```cf-dns``` after the __Development__ steps and run:

    go run main.go create DNS_NAME, RECORD_TYPE, VALUE
    go run main.go create www, CNAME, www-bucket-aws


## Generating 

    go build main.go
    mv main cfdns


