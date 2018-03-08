# S3Bucket

## Development

Inside ```s3bucket``` directory type: 

    export GOPATH=`pwd`


and then...

    go get gopkg.in/kyokomi/emoji.v1
    go get -u github.com/aws/aws-sdk-go




... export AWS_KEY and AWS_SECRET with full access on S3

    export AWS_KEY=“YOUR_KEY_HERE"
    export AWS_SECRET="YOUR_SECRET_HERE"



... to create a bucket, run:


    go run s3bucket.go create BUCKET_NAME



... and to delete a bucket, run:


    go run s3bucket.go delete BUCKET_NAME
    




