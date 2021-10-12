### Usage

This is just an example of cosign signing & verification in golang code

```
1. generate cosign key pair
$ cosign generate-key-pair

2. sign your file
$ go run sign/sign.go sample-text-file cosign.key
Using payload from: sample-text-file
...

3. verify the file with a signature
$ go run verify/verify.go sample-text-file signature cosign.pub
Verified OK
```