## PasskieCredentialEncryption

PasskieCredentialEncryption is an encryption library that will be used by Passkie password manager.

This Golang module will provide encryption functions useful for encrypting and decyrpting credentials as well as hashing functions.

### Running the program

This module is meant to be consumed by other modules as a library of functions.  The main.go function is simply a primitive test function that initializes some data, encrypts it, then decrypts it all while printing out the responses.

This can be run with 

```
go run .
```

### Testing the program

See above to test the program via main.go (not recommended).  Unit tests have also been written to do testing.

```
go test
```
