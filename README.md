# gRPC Address Book Service

This project implements a gRPC service for managing an address book. The service allows clients to list and add persons to the address book.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Proto File](#proto-file)
- [Server Implementation](#server-implementation)
- [Client Implementation](#client-implementation)
- [License](#license)

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/grpc-addressbook.git
    cd grpc-addressbook
    ```

2. Install the necessary dependencies:
    ```sh
    go mod tidy
    ```

3. Generate the Go code from the protobuf file:
    ```sh
    protoc --go_out=. --go-grpc_out=. addressbook.proto
    ```

## Usage

1. Start the gRPC server:
    ```sh
    go run main.go
    ```

2. The server will start listening on `localhost:50051`.

## Proto File

The [addressbook.proto](http://_vscodecontentref_/0) file defines the gRPC service and messages:

```proto
syntax = "proto3";
package addressbook;

import "google/protobuf/timestamp.proto";

option go_package="go/server/addressbook";

message Person {
   string name = 1;
   int32 id = 2;  // Unique ID number for this person.
   string email = 3;

   message PhoneNumber {
     string number = 1;
     PhoneType type = 2;
   }

   repeated PhoneNumber phones = 4;

   google.protobuf.Timestamp last_updated = 5;
}

enum PhoneType {
   PHONE_TYPE_UNSPECIFIED = 0;
   PHONE_TYPE_MOBILE = 1;
   PHONE_TYPE_HOME = 2;
   PHONE_TYPE_WORK = 3;
}

message AddressBook {
   repeated Person people = 1;
}

service AddressBookService {
   rpc ListPerson(ListPersonRequest) returns (ListPersonResponse);
   rpc AddPerson(Person) returns (AddPersonResponse);
}

message ListPersonRequest {}
message ListPersonResponse {
  repeated Person Person = 1;
}

message AddPersonResponse {
  string message = 1;
}
