# awssdkforgotutor
AWS SDK for Go tutorial

https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/welcome.html

- Intall SDK into GOPATH

 $ go get -u github.com/aws/aws-sdk-go

- Determine Region for AWS services

```
$ export AWS_REGION=us-east-1
```

- file for Credentials 

```
   $ mkdir $HOME/.aws
   $ vi $HOME/.aws/credentials
```

```
[default]
aws_access_key_id = ACCESSKEYIDFORTHESERVICE
aws_secret_access_key = ACCESSKEY
```

-- Examples

   Performing Basic Amazon S3 Bucket Operations 

-  s3listitem.go

   List Bucket / List Bucket Items 

https://s3.amazonaws.com/nkvd/s3listitem.go

https://s3.amazonaws.com/nkvd/s3listitem.exe



- s3upfile.go 
   
     Upload a File to a Bucket

https://s3.amazonaws.com/nkvd/s3upfile.go

https://s3.amazonaws.com/nkvd/s3upfile.exe



- s3makebkt.go

     create a bucket

https://s3.amazonaws.com/nkvd/s3makebkt.go

https://s3.amazonaws.com/nkvd/s3makebkt.exe

- s3lsbkt.go

   list Buckets

https://s3.amazonaws.com/nkvd/s3lsbkt.go

https://s3.amazonaws.com/nkvd/s3lsbkt.exe



https://s3.amazonaws.com/nkvd/



