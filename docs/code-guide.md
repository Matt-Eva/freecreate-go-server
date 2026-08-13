# Code Guide

This document serves as a guide for the various tools, technologies, and programming conventions employed by this codebase.

Please read to familiarize yourself with the standard patterns and practices, and to gain a basic understanding of the various languages used in this codebase.

# Table of Contents

- [Code Philosophy](#code-philosophy)
- [Programming Languages - Overview](#programming-languages-and-more)
- [Databases - Overview](#databases)
- [Requisit Tooling](#requisite-tooling)
- [Codebase Structure](#codebase-structure)
- [Go](#go)
- [Error Handling in Go](#error-handling-in-go)
- [JavaScript](#javascript)
- [Error Handling in JavaScript](#error-handling-in-javascript)
- [Postgres](#postgres)
- [SQL](#sql)
- [Dbmate](#dbmate)
- [Valkey](#valkey)

# Code Philosophy

First and foremost, code with calm.

Don't rush. Don't panic. Be curious. Be patient. Take breaks. Persist. Think big. Be free. Be diligent. Be methodical.

In other words, code slow, code thoughtfully, and code happy :).

In addition to that, FreeCreate's codebase and architectural choices are guided by four foundational goals:

1. Maintainability
2. Exentsibility
3. Reliability
4. Performance

In accordance with these goals, there are also a few fundamental coding practices. They are:

- Simplicity
- Clarity
- Specificity
- Locality of Behavior
- Abstraction
- Separation of Concerns

## Maintainability and Extensibility

Maintainability - how easy it is to work within an existing codebase, implement changes and updates, understand past code, and manage dependencies.

Extensibility - how easy it is to add new features, expand into new use cases, and ultimately evolve the project / codebase as time progresses.

These two principles are key to developing long-lasting projects that can and will grow, in multiple capacities.

## Reliability

Your code doesn't break easily. And if it does, it's quick and easy to find out why. And easy to fix!

## Performance

This one is pretty self-explanatory. But basically, you want to aim for optimal performance balanced against these other principles. Nobody likes a slow app or a slow website.

However, If you find yourself compromising your other principles by adding needless complexity for the sake of squeezing out some last little drop of "performance", you've strayed from the path.

There may be edge cases where this is truly worthwhile, but err on the side of caution and temperance.

## Coding Practices

You may feel that there is some tension between some of coding practices listed above.

For example, Simplicity, Clarity, Specificty, and Locality of Behavior may initially seem at odds with Abstraction and Separation of Concerns.

Indeed, we could even split these foundational philosophies into two groups

Preferred:

- Simplicity
- Clarity
- Specificity
- Locality of Behavior

Essential:

- Abstraction
- Separation of Concerns

What does that mean? Did that clear up the picture at all? How can some things be preferred and others be essential?

Basically, when writing software, we should <em>favor</em> writing things in a simple, clear, specific manner with minimal abstraction or opaque code-cleverness, keeping all relevant behavior local to its actual purpose - it's better to have code that works together to accomplish the same thing be in the same place. Separating it out into unecessary packages, hemming and hawing over function length or DRYness, or pursuing some sort of "clever" (aka often difficult) solution are all red herring pitfalls that waste time and can make code harder to work with, write, and debug.

HOWEVER!!!

Abstraction and Separation of Concerns are two ESSENTIAL components of writing good, maintainble, extensible, reliable, performant software.

Heck, every programming language itself is an Abstraction - or did you want to build your website with machine code? (Honestly, if you do that, more power to you, follow your heart!)

If we forgo intelligently applying abstraction and separation of concerns in our codebase, it can become endlessly verbose, frustratingly repetitive, brittle, and ultimately very hard to maintain, update, or build on top of. Separating out core, repetitive functionality makes your code MORE robust and EASIER to work with. Breaking long, complicated functions into more tolerable bite-sized pieces can make them easier to comprehend and manage. Isolating and decoupling independent pieces of your code base from each other allows developers to work better in parallel and keeps your code easier to update and expand.

So what are we to do?

Well, as you write code, prefer keeping things simple, clear, specific, and local. Whenever you're diving into some new feature, start by writing your code this way. Yes, it's helpful to take a minute and assess your problem and see if an abstraction would be helpful or if certain things should be de-coupled from the get go, but don't seek out abstraction. If an abstraction is useful or necessary, it will reveal itself to you as you write your code.

This means that we should embrace refactoring as part of the process of writing code. Not that there is some need to refactor an entire codebase on a routine basis (in fact, I personally would like to avoid doing that whenever possible), but rather than we should refactor our code as we're building out something new.

It's very similar to writing - there may be some novelists, essayists, journalists, or critics who can just execute an absolutely perfect first draft. But the reality is that the vast majority of writers will churn out a semi-coherent first draft, then spend a much longer time combing through and editing and restructuring.

Take this approach with writing code. If you're doing something new, just start making something! Then, as you gain a better understanding of what you're making, or start running into issues created by how you've been writing your code, refactor it into a cleaner, better structure that employs appropriate abstraction and separation of concerns.

Don't be a stubborn purist. Recognize, and adhere to, the value of good code philosophy, but don't choose some "camp" out of ideological zeal - stay open, stay humble, and embrace improvement and enlightenment.

## In Summary

The choices and concepts you will encounter throughout the rest of this guide are all representations of these core philosophies and practices. This is still a work in progress, so if you see room for improvement, your suggestions are most welcome!

[Top](#table-of-contents)

# Programming Languages (and More)

FreeCreate uses two primary programming lanaguages, a markup language, a style sheet language, and a database query language.

## Programming Languages - Go and JavaScript

### Go

Go is the primary language used to built FreeCreate. The web server is written in Go and rendered using Go's native `html/template` package. You can learn more about Go, and how it's used in FreeCreate, in the Go section.

### JavaScript

FreeCreate is also built use JavaScript. While we do use a bundler - Vite - for bundling necessary packages and serving them to our users (as of now, the only package we're bundling and serving is the open source rich text editor, `Lexical`, created by Meta), the rest of our JavaScript is written as standard JavaScript. No React, no TypeScript, just basic JavaScript. More on that later.

## Markup Language - HTML

HTML is the standard markup language for structuring websites. FreeCreate's website is written using plain HTML, rendered via Go's `html/template` package. You will see some Go code interpolated in these HTML files, to handle conditional rendering, rendering of arrays of data, and other programmatic functions.

## Styling Language - CSS

FreeCreate uses plain CSS for its styling. (You may be noticing a theme here.) If we need something fancier / more advanced / "better" at some point, we'll use that. But for now, plain CSS is more than enough.

## Database Query Language - SQL

FreeCreate's main database, horizontally scalable content store, and horizontally scalable donations store are all handled by Postgres instances. All queries are written in `SQL`, the standard query language for relational databases.

While we also do use Valkey as an in-memory cache database, communication with valkey is handled by the Go driver for valkey, `valkey-go`, maintained by Valkey itself. This code is essentially Go code, so there is not a need to learn a specific query language, although understanding the fundamentals of how queries are executed and constructed will still be necessary.

[Top](#table-of-contents)

# Databases

As mentioned, FreeCreate uses Postgres and Valkey as its databases - Postgres for persistent data, and Valkey to handle session caching, query caching, and any other form of ephemeral data.

## Database Drivers

- Postgres: `pgx`
- Valkey: `valkey-go`

## Migration Manager

- FreeCreate does not use an ORM, but it does use an amazingly simple (and amazingly effective), migration manager - `dbmate`. Thank you, `dbmate`, for making an excellent migration manager that is free of ORM overhead and just works!

## Postgres Structure

While many platforms reach for NoSQL databases in order to solve the problem of horizontal scalability, FreeCreate takes the simple approach of leveraging multiple Postgres instances to handle data volume. The structure of FreeCreate's databases is as follows:

- Core DB:
  - Where most of FreeCreate's data lives, particularly data that will many relationships with records in other tables.
  - Primarily, these will be users, creators, their "writing", their writings' "chapters", join tables connecting users to writing in various ways, and any metadata / descriptive data to assist with querying.
  - The goal is to keep the size of individual records low, to allow for scale without the need to shard the database due to data volume.
  - The goal is to keep this database unsharded as long as possible.

- Content DB(s):
  - While writing, chapters of writing, and associated metadata will all be stored in the Core DB, the actual content of writing and chapters will be stored in a separate Postgres database.
  - Content will be linked to a chapter, writing, creator, and user by their various UUIDs, indexed appropriately.
  - The assumption is that there will need to be a need for multiple content db's to account for the volume of written content (this is a far-off future but a possible one)
  - For that reason, each chapter with which content is associated will be given a "shard key" - essentially the name of the database in which its content will be stored.
  - This pattern replicates the sharding pattern found in true sharded databases, but without the management overhead.
  - Connection to multiple content DBs will simply be handled by multiple pgx driver instances.

- Donations DB(s):
  - The Donations DB follows the same pattern as the Content DB - instead of storing donations in the core database, they are stored in separate postgres instances.
  - The rationale for doing this for donations is that they could potentially grow to infinite number, or at least very high volume, and are not needed as part of relational querying. A donation will still be allocated to a piece of content's "rank", but it most likely won't be needed for specific relational queries. And, if it is needed, using the embedded uuid of its related pieces of data from the core db should still be relatively efficient.

## Valkey - Cache

Caching - for session data as well as query caching - will be managed by Valkey, which is an open source redis spinoff. You can read more about it below in the dedicated section on Valkey.

## Realtime (for the future!)

FreeCreate does not currently have any need for realtime features, but in the event that it does in the future - and it just might ;) - the plan is to use Kafka alongside either Timescale DB or either Cassandra or Scylla DB.

Timescale would be convenient, given that it is just Postgres with an extension, and can be managed / queried using traditional postgres techniques.

Cassandra or Scylla will likely be a little more robust, as they are designed to operate as a sharded, scalable cluster, 

## Search Engine - Open Search?

FreeCreate's search functionality does not currently employ a full-scale search engine, but when it does it plans to use Open Search.

[Top](#table-of-contents)

# Requisite Tooling

[Top](#table-of-contents)

# Postgres

## Interacting with your local postgres instance
To access your postgres installation locally, you can run the following command: `psql postgres`.

This will give you access to your entire local postgres instance.

To see all available databases for your local instance, enter the command `\l`.

To access a specific database, enter the command `psql [your_database_name]`.

Once in your specific database instance, enter the command `\dt` to view all tables.

To view information about a specific table within your databae, run the command `\d [your_table_name]`.

## Creating Requisite Databases

To create the necessary databases you'll need for development, first enter your local postgres instance: `psql postgres`.

Then, we'll create the requisite databases using the `CREATE DATABASE` command:

- `CREATE DATABASE freecreate;`
- `CREATE DATABASE freecreate_writing_content_one;`
- `CREATE DATABASE freecreate_donations_one;`

The reason we're adding "one" to the end of our writing content and donations databases is to accomodate for the possibility of multiple content and donations databases.

Once you've created your databases, make sure to add the following environment variables to your `.env` file (make sure this file is also in your .gitignore):

- 

# Dbmate

to create a migration in a specific folder using dbmate, run the command `dbmate -d "./[location of my folder]" new [name_of_my_migration_file]`. This will create a new, timestamped migration in the folder of your choosing.

Example:

`dbmate -d "./internal/db/pg_core/migrations" new create_users`.

Commands for migrating up and rolling back migrations for specific databases are located in the `cmd` folder.

Please make sure to enable the permissions of these scripts to ensure you're able to run them.

Example command:

`./internal/cmd/migrate_pg_core.sh`