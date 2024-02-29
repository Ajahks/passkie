## Passkie 

Passkie is an open source password manager application.  It will encrypt passwords using a master password and can store them per base url.

Passkie is broken up into multiple sub modules

### CredentialEncryption

CredentialEncryption is an encryption module that will be used by Passkie password manager.

This Golang module will provide encryption functions useful for encrypting and decyrpting credentials as well as hashing functions.

### PasswordVerification

PasswordVerification module will be used for verifying and storing the master password hash.

### Storage

Storage module and submodules will include functions that will interface with different storage options.  

Will initially support a local DB via file IO.  But will eventually support plug in play DB's of your choice

### Testing the program

Unit tests have also been written to do testing.

```
go test ./[module folder]

go test ./credentialEncryption/encryption
go test ./credentialEncryption/hash
```

You can also run all the tests in this repo via
```
go test ./...
```
