#!/bin/bash

echo '
package main

import "subscription_management_system/migrations"

func main() {
    migrations.LoadData()
}
' > temp.go
go run temp.go
rm temp.go