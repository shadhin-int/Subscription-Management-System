#!/bin/bash

echo '
package main

import "github.com/shadhin-int/Subscription-Management-System.git/migrations"

func main() {
    migrations.DoMigrate()
}
' > temp.go
go run temp.go
rm temp.go