#!/bin/bash

DB_DNS="user=postgres password=postgres host=postgres-db port=5432 dbname=postgres sslmode=disable"
goose postgres "$DB_DNS" up
