# Code Guide

This document serves as a guide for the various tools, technologies, and programming conventions employed by this codebase.

Please read to familiarize yourself with the standard patterns and practices, and to gain a basic understanding of the various languages used in this codebase.

# Table of Contents

- [Code Philosophy](#code-philosophy)
- [Programming Languages - Overview](#programming-languages-and-more)
- [Databases - Overview](#databases)
- [Requisite Tooling](#requisite-tooling)
- [Codebase Structure](#codebase-structure)
- [Team Roles](#team-roles)
- [Go](#go)
- [Error Handling in Our Go Code](#error-handling-in-our-go-code)
- [JavaScript](#javascript)
- [FreeCreate JavaScript Conventions](#freecreate-javascript-conventions)
- [Error Handling in JavaScript](#error-handling-in-javascript)
- [Postgres](#postgres)
- [SQL](#sql)
- [Dbmate](#dbmate)
- [Valkey](#valkey)
- [Queries](#Queries)

# Code Philosophy

First and foremost, code with calm.

Don't rush. Don't panic. Be curious. Be patient. Take breaks. Persist. Think big. Be free. Be diligent. Be methodical.

In other words, code slow, code thoughtfully, and code happy :).

In addition to that, FreeCreate's codebase and architectural choices are guided by four foundational goals:

1. Maintainability
2. Extensibility
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

It also performs consistently - certain functions are well defined and do not deviate, regardless of when or where they are called.

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

If we forgo intelligently applying abstraction and separation of concerns in our codebase, it can become endlessly verbose, frustratingly repetitive, brittle, and ultimately very hard to maintain, update, or build on top of. 

Separating out core, repetitive functionality makes your code MORE robust and EASIER to work with. 

Breaking long, complicated functions into more tolerable bite-sized pieces can make them easier to comprehend and manage. 

Isolating and decoupling independent pieces of your code base from each other allows developers to work better in parallel and keeps your code easier to update and expand. It also makes it much more reliable and predictable.

So what are we to do?

Well, as you write code, prefer keeping things simple, clear, specific, and local. Whenever you're diving into some new feature, start by writing your code this way. Yes, it's helpful to take a minute and assess your problem and see if an abstraction would be helpful or if certain things should be de-coupled from the get-go, but don't seek out abstraction. If an abstraction is useful or necessary, it will reveal itself to you as you write your code.

This means that we should embrace refactoring as part of the process of writing code. Not that there is some need to refactor an entire codebase on a routine basis (in fact, I personally would like to avoid doing that whenever possible - Matt), but rather that we should refactor our code as we're building out something new.

It's very similar to writing - there may be some novelists, essayists, journalists, or critics who can just execute an absolutely perfect first draft. But the reality is that the vast majority of writers will churn out a semi-coherent first draft, then spend a much longer time combing through and editing and restructuring.

Take this approach with writing code. If you're doing something new, just start making something! Then, as you gain a better understanding of what you're making, or start running into issues created by how you've been writing your code, refactor it into a cleaner, better structure that employs appropriate abstraction and separation of concerns.

Don't be a stubborn purist. Recognize, and adhere to, the value of good code philosophy, but don't choose some "camp" out of ideological zeal - stay open, stay humble, and embrace improvement and enlightenment.

That being said, there are DEFINITELY some patterns and conventions to the code base that you should follow, for your own sake and the sake of all others. BUT, you can always write your code first, get it working, then make sure it follows conventions (and is still working). Just don't forget to update it to follow the conventions. (PLEASE.)

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

Here are the tools you will need installed to work on the project (working on getting this Dockerized):

## Core Technologies:
- Go version 1.27.0
  - [Installing Go](https://go.dev/doc/install)
  - [Managing multiple Go versions](https://go.dev/doc/manage-install)
- Postgresql version 18.6
- Valkey server 8.1.3
- dbmate version 2.33.0
  - [Installing dbmate](https://github.com/amacneil/dbmate)
- React Native version(?)
- Electron version(?)

## Tooling:
- Go - VS Code extension
- Draw.io - VS Code extension
- Go Template Support - VS Code extension
  - author: jinliming2
- Prettier - VS Code extension
- Prettier Plugin Go Template - Prettier Plugin
  - (Reference)[https://github.com/NiklasPor/prettier-plugin-go-template]

## Env file
- Please refer to the sample.env file to determine the environment variables you'll need for the project.
- Please BE CERTAIN not to commit your own .env file to Git history. If you do, please tell Matt as soon as you realize you have. (Don't worry, it'll be ok!)


[Top](#table-of-contents)

# Codebase Structure

## Entry Points

There are three main entry points to the codebase that the API is designed to handle - web, desktop, and mobile.

Currently, web is being handled with traditional SSR and javascript, while desktop is being handled by Electron and mobile is being handled by React Native.

Each of these entrypoints will have separate API endpoints. This is to simplify the workings of different teams across different platforms, as well as keep the code simpler and easier to extend, tweak, and build on.

For example, each different platform will likely have a different auth strategy - giving each platform its own sub-api, router, and set of handlers will make it easier to incorporate and manage those different auth strategies.

As mentioned, this also makes it much easier for different teams to work in parallel. If a feature needs to be changed or created for react-native, the team handling the react-native frontend and API endpoints and handlers can easily do so without accidentally compromising the code of the electron team or the web team.

## Query Handlers and Shared Functionality

However, some code can and should be shared across all teams. 

This functionality is separated and shared to ensure that there is a Single Way of Doing Something for certain types of actions - this is the "business logic" of the application.

For example, there should be one and only one function for generating OTPs. Similarly, logic for sending email should be extrapolated out into its own shared package.

This not only ensures consistent user experience, but enforces the intended design patterns of the codebase and makes the codebase more resilient to errors.

This may slow down developers somewhat, but we actually WANT that in these instances. For cases, we want to be thoughtful, deliberate, and careful when making any changes, to ensure we stay bug-free and deliver a uniform user experience.

### Query Handlers

This is particularly true for handling database logic - creation queries, search queries, update queries, etc.

When it comes to interacting with the database(s) and cache(s), there should be ONE way to interact with them.

Moreover, sometimes a single "query" - like a search query - may actually involve a whole number of queries, or could potentially run a host of various queries.

To ensure that these occur in the correct order and follow the correct patterns every time, it's better to extrapolate this logic out into consolidated, shared functions, rather than trying to replicate them across every platform. We don't want "web search" to somehow differ from "desktop search" in its functionality - search should be search.

Moreover, the schema is pretty complex. Having an established query interface that can be more easily dropped into and integrated with api endpoints makes it easier for the developers of various platforms to create and integrate features. Instead of everybody having to be masters of the schema and query patterns in addition to being masters of their own platform, they can focus more on delivering features and developing competence in their domain.

This also allows a person - or team - to focus solely on become schema and query experts. The schema is a bit complicated - and will only become moreso with time - although much of that complexity is just due to information / table volume and variety rather than innate complexity. 

This team can then also spend more time thinking about and working on ways to improve data storage and optimize queries, whether considering scale, performance, or feature buildout.

# Team Roles

Given this pattern, this would be the ideal team structure:

Graphic Design / UI / UX Design:
- This team can focus entirely on look and layout. 
- They may dabble in some CSS or at least become experts in determining color codes to ensure universal experience cross platform

Web Team (Fullstack, Frontend Emphasis):
- Handles web API endpoints and Auth (Go), as well as frontend (Go templates + JavaScript).
- Implements the CSS rules that will create the look and layout implemented by the design team.

Mobile Team (React-Native) (Fullstack, Frontend Emphasis):
- Handles Mobile API endpoints and Auth, as well as frontend (React-Native or other frontend mobile language / framework - Swift (IOS), Kotlin (Android), Flutter (cross-platform), etc.).
- Also implements whatever visual instructions - CSS analog - needed to create the look and layout created by the design team.

Desktop Team (Electrion) (Fullstack, Frontend Emphasis):
- Handles Desktop API endpoints and Auth, as well as desktop frontend (Electron, Tauri, etc.)
- Implements styling instructions - CSS analog - needed to create the look and layout created by the design team.

Backend Team - System Architecture, Database Administration, Query Handling, and Business Logic
- This team is reponsible for designing and implementing the database structure, overall backend application structure - caching, storage, horizontal scaling, realtime (eventually), etc. - and query handling for all of these various operations.
- They will also handle other cross-team backend functionality, like email sending, otp generation, and the like.
- If other teams want a feature implemented, or have questions about a feature, they should reach out to this team. They should NOT jump into this portion of the codebase and start making changes.
- This team will need to have a rock solid understanding of how the application is designed to function, and should have a certain level of awareness of the other team's needs, so that they can help serve as a go-between between the different platform teams
- Note from Matt: 
  - I'm probably going to be the one handling this role for the time being, as I currently have the most in-depth knowledge of planned features and the overall schema.
  - There's definitely opportunity to hop in and contribute here, but the learning curve will likely be a little bit steeper and the implementation will need to be a little bit stricter (aka I will probably want to have more oversight of any feature being implemented).
  - I also can tend to get a little carried away with schema design (in a GOOD way, I promise), and it can get a little bit nuts back there. If you want to dive into the muck with me, you're welcome to! But be warned...


[Top](#table-of-contents)

# Go 

## Why Go?

Why are we using Go as our primary backend language? There are several reasons:

- Go is statically typed
- Go is compiled
- Go has built in concurrency and parallelism that is easy to implement (and often implemented out of the box in core package functionality)
- Go has a built in formatter that dictates formatting rules: `go fmt ./...`.
- Go is simple
  - The founding philosophy of Go is that there is "one way to write things"
  - While this is no longer true, Go still has relatively little syntactic sugar, and is easy to read and logic through.
  - It's very explicit and clear and keeps things predictable and easy to logic through.
- Go uses errors as values
  - Many developers hate this about Go. You will see a lot of people complaing about writing `if err != nil` over and over and over again.
  - I, personally, love errors as values (Matt). 
  - It makes error handling a first order of business when writing code, rather than a secondary order of business.
  - For more on this, please see [error handling in our Go code](#error-handling-in-our-go-code).

[Top](#table-of-contents)

# Error Handling in Our Go Code

As mentioned in the above Go section, Go treats errors as values, and has the native `error` type.

This is great, but there is one downside and one consideration:

1. The Downside: 
  - Go's errors don't automatically do stack tracing when they occur. You can print them, and get the error and corresponding message, but only doing that doesn't really give you a clue as to where exactly the error occurred in your codebase. This is really frustrating, and I don't like it (Matt).
  - However, we've fortunately been able to build our OWN stacktrace logger. Yay!
  - This is in the `/lib/logger` folder and is the `logger` pacakge.
  - To log an error, just pass a native go error type to `logger.Log()`, and you'll get your handy-dandy error with the corresponding stack trace. All is well :).

2. The Consideration:
  - Typically, Go errors propogate up through the function chain - a function generates an error and passes it to its parent function, which passes it to its parent function, and so on and so forth until we get to our API handler and our user endpoint.
  - This is also great - it's the whole errors as values philosophy - but we also don't want to be sending native error messages back to the frontend, where they may be communicated to our users. Those errors might have sensitive data! Nooooo!
  - Instead, we've created our own `api_error` error type - you can find it in the `/lib/api_error` folder.
  - `api_error`s have three fields: 
    - `Code` - the http.status code that should be sent back to the frontend based on the nature of the error that occurred - `http.StatusInternalServerError`, `http.StatusNotFound`, etc.
    - `Message` - this will be the user facing message that we DO want to display to the user.
    - `Error` - this will be the error that triggered this `api_error` in the first place. Currently, there isn't a plan to do anything with it, but it seems prudent to make it part of the `api_error` in case it becomes handy for debugging at a future point.

## How We Handle Errors

Given these two considerations, the way we handle errors is as follows:

1. Functions that we ourselves are writing do not return a native Go `error` type.
    - Instead, they return a pointer to an `api_error.Error`.
      - Why a pointer and not the struct itself? So that we can follow Go's built in `if err != nil {}` pattern! This is idiomatic to Go, and we want to preserve that convention.
2. When an error is triggered by something that goes wrong, we log that native Go `error`:
    - `logger.Log(err)`
3. Then, we construct our `api_error`:
    - ```Go
      logger.Log(err)

      apiErr := &api_error.Error{
        Code: http.StatusInternalServerError,
        Message: api_error.InternalServerErrorMessage,
        Error: err,
      }

      return apiErr
    
4. This error propogates up the function chain until it finally reaches the api endpoint. There, we send the message back, as well as the Code:
    - ```Go
      apiErr := package.MyFunction(parameters)
      if apiErr != nil{
        http.Error(w, apiErr.Message, apiErr.Code)
        return
      }

And voila! We've logged our error where it originally occurred, which is good for us developers, and we've also communicated the correct status code and message to our frontend users.

One thing you may have noticed is that the `apiErr` created in the example uses a pre-established message defined in the `api_error` package.

Currently, there is only one such pre-established message, which is reserved for the 500, "internal server error" messages.

These errors occur when something has gone wrong on the server itself (or with the database, etc.), and as such, we want to have a single, consistent message to display to users when these types of errors occur.

There may be more to come, but in other instances you'll write your own error messages:

```Go
if err != nil {
  logger.Log(err)
  
  msg :="We could not find a user with that email address."

  apiErr := &api_error.Error{
    Code: http.StatusNotFound,
    Message: msg,
    Error: err,
  }

  return apiErr
}

```

[Top](#table-of-contents)

# JavaScript

# FreeCreate JavaScript Conventions

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

# Dbmate

to create a migration in a specific folder using dbmate, run the command `dbmate -d "./[location of my folder]" new [name_of_my_migration_file]`. This will create a new, timestamped migration in the folder of your choosing.

Example:

`dbmate -d "./internal/db/pg_core/migrations" new create_users`.

Commands for migrating up and rolling back migrations for specific databases are located in the `cmd` folder.

Please make sure to enable the permissions of these scripts to ensure you're able to run them.

Change permissions:

- `chmod +x ./internal/cmd/_migrate_pg_core.sh`

Example command:

- `./internal/cmd/migrate_pg_core.sh`

- `./internal/cmd/rollback_pg_core.sh`

# Queries

The following is a comprehensive list of queries made by the platform, as well as strategies for handling them.

## Writing

Writing will need to be queried in a variety of ways.

- Individual pieces of writing will need to be loaded, along with its creators and chapters, for both viewing and editing
- For writing that is loaded for editing, it will need to load both published and unpublished chatpers.