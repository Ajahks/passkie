## Passkie 

Passkie is an open source password manager application.  It will encrypt passwords using a master password and can store them per base url.

## CLI Application

### Building

To build the Passkie CLI application from the source code run:

```
go build -o passkie ./app/cli
```

### Usage

This will create an executable named 'passkie'.  To run the application via the exectuable run.  This will provide usage details for the application.

```
./passkie
```

For first time users, a user will need to be initialized.  To initialize a user as 'default' omit the user parameter
```
./passkie init --user <username (default)>
```

Once the user has been initialized, credentials can now be stored per site
```
./passkie store --user <username (default)> --site <base url (required)>
```

To retrieve the credentials
```
./passkie retrieve --user <username (default)> --site <base url (required)>
```

## Passkie Framework

The passkie framework is broken up into multiple sub modules

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
