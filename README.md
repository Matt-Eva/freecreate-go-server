# FreeCreate

FreeCreate is a donations-based platform for writers. Its aim is to provide a mindful, best-in-class experience for writers and readers alike, and offer an alternative to the current attention, membership, and advertisement driven content platforms.

Eventually, the hope is to have FreeCreate host all forms of content - from art, to music, to video - but for now my goal is to just make it a great platform for readers and writers.

Together, we can make the internet a beautiful place!

## Code Guide

To get a working knowledge of the codebase and its various languages, tools, and design philosophies, please read the `code-guide` document within the `docs` folder.

## Written in Go

FreeCreate is writting with Go, HTML (Go Templates), JavaScript, and CSS.

## Architecture.

Databases: FreeCreate uses Postgres for its core relational database, Valkey as an in memory cache, and Postgres again as a horizontally scalable content store.

Database Drivers: `pgx`, `valkey-go`.

Migration Manager: `dbmate`.

Routing: `chi`.

Session management: `gorilla/sessions`.

CSRF Protection: `gorilla/csrf`.

Rich text editor: `Lexical`.

Rendering: Go's `html/template` library. And vanilla CSS and JavaScript :).

Email: `resend`.

Rich Text Editing: Meta's `lexical` editor.

Realtime (for the future): While not yet implemented (or necessary), Kafka will be used for any realtime feature development.
