# Writing Structure

All writing is structured such that a piece of writing is split into various "chapters", which are stored in Mongo.
Writing can also have "notes" documents.
Primary / shard key for chapters and notes in Mongo is a piece of writing's uuid.

# Database Error Handling

Ok, to properly handle error codes from database queries, we'll need to first check if the error is nil. If it's not, we'll then need to convert it into a \*pgconn.PgError as mentioned here: https://github.com/go-gorm/gorm/issues/4135

From there, we can get the error code via pgErr.code. We can check this error code against the list of postgres error codes to determine what ultimately caused the problem, then handle the code effectively: https://www.postgresql.org/docs/current/errcodes-appendix.html

- Uniqueness constraint code: 23505

# Query Pattern

Currently the way we have the uuid's set up for users, creators, and pieces of writing, we're going to have to query for users, etc. and retrieve their id's instead of uuids in order to create content. This involves round trip interactions between the api and the database, but makes the creation queries simpler, as we can populate a struct and benefit from its zero value input. Otherwise the values would be set to null.

Pros: makes it harder to accidentally insert empty values into the database.

Cons: Can dramatically slow down creation queries by a factor of 3 or so.

Decision: just go for the simpler option that works for the time being. We aren't concerned about absolute maximum performance at this point, and this can be addressed / changed with relative ease in the future when maximum performance becomes more critical.
