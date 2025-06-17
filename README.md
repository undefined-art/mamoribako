# Mamoribako

## Overview

A secure command-line password manager written in Golang. Passwords are encrypted with AES-256 using a key derived from a master password via PBKDF2.

## Features

- Add, get, list, and delete passwords
- Master password protected with key derivation (PBKDF2)
- Secure input and memory protection using memguard
- Encrypted password storage in JSON file

## Usage

Initialize the password manager (generate salt):

    mamoribako init

Add a password:

    mamoribako add <key> <password>

Retrieve a password:

    mamoribako get <key>

List all keys:

    mamoribako list

Delete a password:

    mamoribako delete <key>

## Security Notes

- Use a strong master password.
- The master password is never stored; only a derived key is used.
- Passwords are encrypted with AES-GCM.
- Sensitive data is protected in memory using memguard.
- Salt is stored locally to derive the key consistently.

## Building

    go build -o mamoribako

## License

MIT License
