# pritunl-link
Working progress installation guide

## Installing
1. Install buildessentials 
2. Download Go
3. Extract go to /usr/local
4. Add go bin to $PATH
5. Do go mod innit
6. Install Dependencies

```
go get github.com/Azure/azure-sdk-for-go/services/compute/mgmt/
go get github.com/Azure/go-autorest/autorest
go get github.com/Azure/go-autorest/autorest/azure/auth
go get github.com/aws/aws-sdk-go/aws
go get github.com/aws/aws-sdk-go/aws/ec2metadata
go get github.com/aws/aws-sdk-go/aws/session
go get github.com/aws/aws-sdk-go/service/ec2
go get github.com/dropbox/godropbox/container/set
go get github.com/dropbox/godropbox/errors
go get github.com/oracle/oci-go-sdk/common
go get github.com/oracle/oci-go-sdk/core
go get github.com/sirupsen/logrus
go get golang.org/x/net/context
go get golang.org/x/oauth2/google
go get google.golang.org/api/compute/v1
```
