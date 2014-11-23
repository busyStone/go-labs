# !/bin/bash

# use go get github.com/yvasiyarov/swagger
# then in $GOPATH/bin you have swagger

baseDir=/github.com/busyStone/go-labs/revel-swagger

# The package that implements the API controllers, relative to $GOPATH/src
apiPackage=$baseDir/app/controllers

# you need change 
# $baseDir/swagger-ui/index.html
# line 28 url:""
# to 
# url:http://basePath/docs
# Web service base path
basePath=127.0.0.1:9000

mainApiFile=$baseDir/app/init.go
output=$GOPATH/src/$baseDir/modules/swagger/app/controllers

$GOPATH/bin/swagger -apiPackage=$apiPackage -basePath=$basePath -mainApiFile=$mainApiFile -output=$output

echo 'change docs.go content'
sed 's/package main/package controllers/' $output/docs.go > $output/docstemp.go
rm $output/docs.go
mv $output/docstemp.go $output/docs.go
