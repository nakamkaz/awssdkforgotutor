# awssdkforgotutor
AWS SDK for Go tutorial

https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/welcome.html

- Intall SDK into GOPATH

 $ go get -u github.com/aws/aws-sdk-go

- Determine Region for AWS services

  $ export AWS_REGION=us-east-1

- file for Credentials 
   $ mkdir $HOME/.aws
   $ vi $HOME/.aws/credential
  
```
[default]
aws_access_key_id = ACCESSKEYIDFORTHESERVICE
aws_secret_access_key = ACCESSKEY
```

!Example 

!! Performing Basic Amazon S3 Bucket Operations 

-  lists3.go

    List Bucket / List Bucket Items 
  
  - s3upfile.go 
   
     Upload a File to a Bucket
     
     
     
