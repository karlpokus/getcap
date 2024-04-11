# getcap
P0rtable TPM2 getcap binary

Since https://github.com/tpm2-software/tpm2-tools comes in a 1000 pieces, is dynamically linked and can't just be shipped to a remote host.

# usage

````sh
# build the thing (assuming compatible arch and os)
$ go build
# ship it
$ scp getcap user@host:/tmp
# run it
$ ssh user@host "sudo /tmp/getcap"
````
