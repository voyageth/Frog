sudo yum -y install go
export GOPATH=/home/ec2-user/go
go get -t github.com/revel/revel
go get -t github.com/revel/cmd/revel
go get -t github.com/go-gorp/gorp
go get -t github.com/go-sql-driver/mysql
go get -t golang.org/x/crypto/bcrypt
