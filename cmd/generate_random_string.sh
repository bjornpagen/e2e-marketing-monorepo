#!/bin/sh
# Generate a cryptographically secure random string of 64 characters

# Define the character set to use
charset="ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

# Generate the random string
result=""
for i in $(seq 1 64)
do
  result="$result${charset:$(od -An -N1 -tu1 /dev/urandom | tr -d ' ') % ${#charset}:1}"
done

# Print the result
echo "$result"
