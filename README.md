# Base Module

A simple example of how to create a reusable Go module with commonly used tools.

## Installation

`go get -u github.com/Base-module/golang-base-module`

## Included Tools

### JSON Processing

- [X] Read JSON
- [X] Write JSON
- [X] Produce a JSON encoded error response
- [X] Post JSON to a remote service

### File Operations

- [X] Upload a file to a specified directory
- [X] Upload multiple files
- [X] Download a static file
- [X] Create a directory, including all parent directories, if it does not already exist

### String Utilities

- [X] Create a URL safe slug from a string

### Random Data Generation

#### Strings

- [X] `RandomString(n)` - Generate random string of length n
- [X] `RandomStringFromCharset(n, charset)` - Generate random string from custom charset
- [X] `RandomAlphaString(n)` - Generate alphabetic string (letters only)
- [X] `RandomAlphaNumericString(n)` - Generate alphanumeric string
- [X] `RandomNumericString(n)` - Generate numeric string
- [X] `RandomLowerString(n)` - Generate lowercase string
- [X] `RandomUpperString(n)` - Generate uppercase string
- [X] `RandomHexString(n)` - Generate hexadecimal string

#### Numbers

- [X] `RandomInt(min, max)` - Generate random integer in range
- [X] `RandomInt64(min, max)` - Generate random int64 in range
- [X] `RandomFloat64(min, max)` - Generate random float64 in range
- [X] `RandomBool()` - Generate random boolean

#### Identifiers

- [X] `RandomUUID()` - Generate UUID v4
- [X] `RandomUsername()` - Generate random username
- [X] `RandomPassword(length)` - Generate strong password (mixed chars)

#### Contact Info

- [X] `RandomEmail()` - Generate random email address
- [X] `RandomEmailWithDomain(domain)` - Generate email with specified domain
- [X] `RandomPhone()` - Generate random phone number (XXX-XXX-XXXX)
- [X] `RandomPhoneWithPrefix(prefix)` - Generate phone with country code

#### Network

- [X] `RandomIPv4()` - Generate random IPv4 address
- [X] `RandomIPv6()` - Generate random IPv6 address
- [X] `RandomMAC()` - Generate random MAC address
- [X] `RandomURL()` - Generate random URL

#### Date & Time

- [X] `RandomDate(start, end)` - Generate random date in range
- [X] `RandomDateString(start, end, format)` - Generate formatted random date

#### Slice Operations

- [X] `RandomSliceElement[T](slice)` - Pick random element from slice
- [X] `RandomSliceElements[T](slice, n)` - Pick n random elements (with duplicates)
- [X] `RandomUniqueSliceElements[T](slice, n)` - Pick n unique random elements
- [X] `ShuffleSlice[T](slice)` - Shuffle a slice

#### Bytes

- [X] `RandomBytes(n)` - Generate n random bytes
