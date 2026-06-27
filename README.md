# FreeCreate

FreeCreate is a donations-based platform for writers. Its aim is to provide a mindful, best-in-class experience for writers and readers alike, and offer an alternative to the current attention, membership, and advertisement driven content platforms.

Eventually, the hope is to have FreeCreate host all forms of content - from art, to music, to video - but for now my goal is to just make it a great platform for readers and writers.

Together, we can make the internet a beautiful place!

## Written in Go

I'm currently building FreeCreate in parallel right now, following two separate paradigms, for my own edification.

This version of FreeCreate is built with Go, following a minimalist architecture that primarily relies on barebones, native technologies, with a focus on simplicity and peformance.

The other version is being built with Ruby on Rails. You can check it out <a href="https://github.com/Matt-Eva/freecreate-rails">here</a>.

## Architecture.

Databases: FreeCreate uses Postgres for its core relational database, Valkey as an in memory cache, and Postgres again as a horizontally scalable content store.

Database Drivers: `pgx`, `valkey-go`.

Migration Manager: `dbmate`.

Routing: `chi`.

Session management: `gorilla/sessions`.

Rich text editor: `Lexical`.

Rendering: Go's `html/template` library. And vanilla CSS and JavaScript :).

Email: `resend`.

CSRF Protection: `gorilla/csrf`.

Realtime (for the future): While not yet implemented (or necessary), Kafka will be used for any realtime feature development.
