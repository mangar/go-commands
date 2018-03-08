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
    





__Building__

The command bellow will make the binary and install it at: ```/usr/local/bin```.
After the instalation you can test it by running: ```/usr/local/bin/s3bucket version```


	make
	make install


__Developing__

The command bellow will install the depencies needed to run the project


	make development



__Installing__


	curl -sSL https://github.com/mangar/go-commands/archive/master.zip -o go-commands-master.zip; unzip go-commands-master.zip; cd go-commands-master/src/s3bucket/src; make; make install



---




