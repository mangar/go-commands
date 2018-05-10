# S3User


Create a User at AWS


## Development


Into ```go-commands```

    export GOPATH=`pwd`

    
and then...

    
    go get gopkg.in/kyokomi/emoji.v1
    go get -u github.com/aws/aws-sdk-go


... export AWS_KEY and AWS_SECRET with full access on S3

    export AWS_KEY="YOUR_KEY_HERE"
    export AWS_SECRET="YOUR_SECRET_HERE"




## Running


__Create a Bucket__

... to create a bucket, go into ```s3user``` after the __Development__ steps and run:

    go run main.go create USERNAME BUCKET1,BUCKET2,BUCKET3



__Delete a bucket__

... and to delete a bucket, run:

    go run main.go delete USERNAME
    
    
    
    
__List of buckets__


    go run main.go list    
    
    

## Generating 

    go build main.go
    mv main s3bucket


