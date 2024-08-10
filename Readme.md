# Encryption and Decryption Tool

This Go program allows users to encrypt and decrypt messages using three different methods: **ROT13**, **Reverse**, and **Pairs**. The user can choose to either encrypt or decrypt a message and then select one of the available methods.

## Features

- **Encrypt and Decrypt Messages**: The program can handle both encryption and decryption.
- **Multiple Encryption Methods**:
  - **ROT13**: Shifts each letter by 13 positions in the alphabet.
  - **Reverse**: Reverses the position of the letters in the alphabet.
  - **Pairs**: Swaps adjacent letters in a predefined pair set.

## How It Works

1. **User Input**: The user is prompted to choose between encryption and decryption, select an encryption method, and then input the message.
2. **Processing**: Based on the user's choices, the message is either encrypted or decrypted using the selected method.
3. **Output**: The result is displayed to the user.

## Methods Overview

### 1. `encrypt_pairs` and `decrypt_pairs`
These functions encrypt or decrypt a message by swapping letters according to a predefined mapping:
- 'a' ↔ 'b', 'c' ↔ 'd', and so on.
- Capital letters are also supported.

### 2. `encrypt_rot13` and `decrypt_rot13`
These functions implement the ROT13 cipher:
- Each letter is shifted by 13 positions in the alphabet.
- For example, 'a' becomes 'n', and 'n' becomes 'a'.

### 3. `encrypt_reverse` and `decrypt_reverse`
These functions reverse the position of each letter in the alphabet:
- 'a' becomes 'z', 'b' becomes 'y', and so on.

## Code Structure

- **Main Functions**:
  - `main`: The entry point of the program. It handles user input and calls the appropriate functions based on the user's choices.
  - `getInput`: Manages user interaction and retrieves input.

- **Encryption/Decryption Functions**:
  - `encrypt_pairs`, `decrypt_pairs`: Handle the Pairs encryption/decryption.
  - `encrypt_rot13`, `decrypt_rot13`: Handle the ROT13 encryption/decryption.
  - `encrypt_reverse`, `decrypt_reverse`: Handle the Reverse encryption/decryption.

## Example Usage

```bash
Hello User, this tool will help you encrypt or decrypt your message.
Select Operation (1/2)
 1.Encrypt
 2.Decrypt
# Choose 1 for Encrypt or 2 for Decrypt

Select method:
 1.Rot13
 2.Reverse
 3.Pairs
# Choose an encryption method

Enter your message: HelloWorld
# Enter the message to be encrypted or decrypted