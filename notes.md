# Writing Structure

All writing is structured such that a piece of writing is split into various "chapters", which are stored in Mongo.
Writing can also have "notes" documents.
Primary / shard key for chapters and notes in Mongo is a piece of writing's uuid.

# Database Error Handling

Ok, to properly handle error codes from database queries, we'll need to first check if the error is nil. If it's not, we'll then need to convert it into a \*pgconn.PgError as mentioned here: https://github.com/go-gorm/gorm/issues/4135

From there, we can get the error code via pgErr.code. We can check this error code against the list of postgres error codes to determine what ultimately caused the problem, then handle the code effectively: https://www.postgresql.org/docs/current/errcodes-appendix.html
