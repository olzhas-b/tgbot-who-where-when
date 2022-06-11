#!/bin/bash

DBSTRING="
user=prgsupezuxwnyl
password=37e3acc567a1544135dc73692490c626b7dfeb885d74268eba1fcde46e91e8a3
host=ec2-34-248-169-69.eu-west-1.compute.amazonaws.com
port=5432
dbname=demo702cgr2ogu
sslmode=disable"
db="postgres://postgres:postgres@localhost:5432/postgres"
goose postgres "$db" up
#database:
#  host: "ec2-34-248-169-69.eu-west-1.compute.amazonaws.com"
#  name: "demo702cgr2ogu"
#  port: "5432"
#  user: "prgsupezuxwnyl"
#  password: "37e3acc567a1544135dc73692490c626b7dfeb885d74268eba1fcde46e91e8a3!"
