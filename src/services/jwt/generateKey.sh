#!/bin/bash

# access
openssl ecparam -outform PEM -genkey -name prime256v1 -noout -out access.private.pem
openssl ec -inform PEM -outform PEM -in access.private.pem -pubout -out access.public.pem

# refresh
openssl ecparam -outform PEM -genkey -name secp384r1 -noout -out refresh.private.pem
openssl ec -inform PEM -outform PEM -in refresh.private.pem -pubout -out refresh.public.pem
