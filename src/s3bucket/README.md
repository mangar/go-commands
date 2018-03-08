# S3Bucket

## Development


Inside the base dir, same level as ```src``` directory, type:

    go get gopkg.in/kyokomi/emoji.v1
    go get -u github.com/aws/aws-sdk-go


... export AWS_KEY and AWS_SECRET with full access on S3

    export AWS_KEY=“YOUR_KEY_HERE"
    export AWS_SECRET="YOUR_SECRET_HERE"




## Running


__Create a Bucket__

... to create a bucket, run:


    go run s3bucket.go create BUCKET_NAME



__Delete a bucket__

... and to delete a bucket, run:


    go run s3bucket.go delete BUCKET_NAME
    


