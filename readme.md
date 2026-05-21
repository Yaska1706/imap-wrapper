# IMAP Wrapper

A lightweight wrapper library built on top of `github.com/emersion/go-imap` that simplifies IMAP access for clients.

This repository is intended to help developers connect to an IMAP server, browse mailboxes, read messages, manage message state, and optionally save sent messages without exposing low-level IMAP protocol details.

## What this wrapper should do

The wrapper is designed to provide a higher-level abstraction over `go-imap` with these responsibilities:

- Manage IMAP connection/session lifecycle
- Authenticate and negotiate server capabilities
- List and manage mailboxes
- Search and fetch messages
- Read headers, body parts, and attachments
- Update message flags and move/copy messages
- Expose a simple, client-friendly public API
- Hide `go-imap` low-level constructs like `SeqSet`, `FetchItem`, and raw body section names

> Note: IMAP is not an SMTP sender. The wrapper should provide a send/save abstraction, but actual message transmission requires a separate SMTP transport or adapter.

## Intended API surface

The wrapper should expose methods such as:

- `Connect(config)`
- `Login(username, password)`
- `ListMailboxes()`
- `SelectMailbox(name)`
- `Search(criteria)`
- `FetchMessage(uid, options)`
- `FetchMessages(criteria, options)`
- `AddFlags(uid, flags)`
- `RemoveFlags(uid, flags)`
- `MoveMessage(uid, mailbox)`
- `CopyMessage(uid, mailbox)`
- `AppendMessage(message, mailbox)`
- `Logout()`

## Key concepts

### Connection Management

The wrapper should handle:

- TLS / StartTLS support
- Login and logout
- Reconnect if the session drops
- Capability detection like `IDLE`, `UIDPLUS`, and `CONDSTORE`

### Mailbox Management

The wrapper should provide mailbox-level functions such as:

- `ListMailboxes`
- `SelectMailbox`
- `CreateMailbox`
- `DeleteMailbox`
- `RenameMailbox`
- `GetMailboxStatus`

### Message Retrieval

The wrapper should expose message retrieval in a friendly form:

- `FetchMessage(uid, options)` returns a domain `Message` object
- `FetchHeaders(uid)` returns only message headers
- `FetchBody(uid, part)` returns a specific body part
- `FetchAttachments(uid)` returns parsed attachment metadata

A high-level `Message` should include:

- UID
- Subject
- From / To / Cc / Bcc
- Date
- Flags
- BodyText / BodyHTML
- Attachments

### Search and Filters

The wrapper should make searching easy by mapping client criteria to IMAP search:

- unread messages
- date ranges
- sender or subject filters
- flagged, answered, or draft state

### Message State Updates

The wrapper should support message state updates like:

- mark as read/unread
- add or remove flags
- copy or move messages
- delete and expunge

### Send / Save Sent Message

Because IMAP does not transmit outgoing email, the wrapper should support one of the following:

- a send helper that uses SMTP underneath and optionally appends the sent mail to a `Sent` mailbox
- or a simple `AppendMessage` API that stores a built message to a mailbox

## Getting started

### Clone the repository

```bash
git clone https://github.com/<owner>/imap-wrapper.git
cd imap-wrapper
```

### Install dependencies

This is a Go project. Use Go modules to install dependencies:

```bash
go mod tidy
```

### Build

```bash
go build ./...
```

### Run tests

```bash
go test ./...
```

## Usage example

A consuming client should be able to do something like:

```go
client := NewIMAPClient(config)
err := client.Connect()
if err != nil { panic(err) }

err = client.Login(username, password)
if err != nil { panic(err) }

mailboxes, err := client.ListMailboxes()
_ = mailboxes

err = client.SelectMailbox("INBOX")

uids, err := client.Search(SearchCriteria{Unseen: true})

messages, err := client.FetchMessages(uids, FetchOptions{HeadersOnly: false})
for _, msg := range messages {
    fmt.Println(msg.Subject)
}

err = client.AddFlags(uids, []string{"\Seen"})

err = client.Logout()
```

## Contributing

Contributions should keep the wrapper focused on abstraction and usability.

- Avoid exposing raw `go-imap` internals
- Keep the public API stable and domain-oriented
- Add tests for each wrapper behavior

## Project goals

The repo is meant to be a foundation for mail clients that need IMAP access without protocol plumbing. It should let developers build features like inbox browsing, message reading, search, flag management, and sent-mail storage through a clear interface.

## License

Add your project license here.

