# Auth

Have a user submit their email. Then, re-render the page to display the email (hidden, read only) and the OTP field. Store the OTP in valkey related to both the email and the session ID to ensure maximum security. And, of course, make sure Cookie is HTTP only, secure, and samesite strict.

Implement blind_indexing and SHA encryption for storage of user emails.

Try to ensure that user creation doesn't occur until OTP has been successfully verified upon signup.

To keep session state persistent between guest logins and logouts, save the session data from the login (if needed), then destroy the old cookie session and create a new cookie session with / without the userId.

# Writing Structure

All writing is structured such that a piece of writing is split into various "chapters", which are stored in the postgres content DB.
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

# CSRF

Look into Gorilla CSRF more to ensure protection against CSRF attacks.

# Styling

It looks like the downloading of separate CSS files without preloading approximately doubles the apparent render time, as the longest wait for the page to render is usually the roundtrip to the server.

This can be ammeliorated by using the preload attribute for css stylesheets, then dynamically changing their attributes once they're loaded to be regular stylesheets, which will allow the html to begin rendering immediately and also allow for caching of most of the css.

However, this will likely cause some layout / appearance shift of the website. So maybe embedding styles in the head is just the way to go? Either that or enable long caching in the browser.

# dbmate

to create a migration in a specific folder using dbmate, run the command `dbmate -d "./[location of my folder]" new name_of_my_migration_file`.

To run a migration for a specific database instance, run this command `dbmate -d "./db/pg_core/migrations" -s "./db/pg_core/schema.sql" --url "postgres://matte:code@localhost:5432/freecreate_go?sslmode=disable" migrate`.

To run the rollback, run `dbmate -d "./db/pg_core/migrations" -s "./db/pg_core/schema.sql" --url "postgres://matte:code@localhost:5432/freecreate_go?sslmode=disable" rollback`.

Note that for the custom schema sql argument, you must pass the full path as an argument, otherwise it will not write to the file.

To Do: Create some bash shorthand commands for this.

# creator tags / topics vs writing tags / topics

store writing tags and topics in a gin array - 23 values max - on the table itself.

because creator tags and topics will include all of the tags and topics they've ever added to their pieces of writing, we will store that information in a join table between a tag and topic in another table (for simplicity, they should just be stored in the tags table).

Because of this, we will be running a join operation on them.

However, we will store the writing types they've written across within a GIN index, since that value cannot span as many unique values.

# revealing hidden content, no JS

the following are two ways to reveal hidden content without using any JS.

One is by using the details HTML element, the other is by using css and a checkbox:

```html
<details>
  <summary>Submit otp</summary>
  <form action="/login" method="POST">
    {{ .CSRFToken }}
    <input type="hidden" name="form_action" value="submit_otp" />
    <input type="text" />
    <input type="submit" />
  </form>
</details>

<style>
  #toggle {
    display: none;
  }

  #checkbox {
    /* display: none; */
  }

  #checkbox:checked ~ #toggle {
    display: block;
  }
</style>
<label for="checkbox">show form</label>
<input type="checkbox" name="checkbox" id="checkbox" />
<p id="toggle">this is hidden content</p>
```

# Post - Redirect - Get request pattern

When rendering web pages in the traditional way without javascript rendering on the frontend,
we follow a post - redirect - get paradigm.

Basically, if any post requests are made on a specific page, upon their success they will redirect to the desired page - it may even be the same page, just rendered as a get, if that's the desired functionality.

We want to redirect rather than just handle the page directly because a page rendered by a post will cause the post request to run again if the user hard refreshes the page.

However, upon a failed post request, the post page will need to re-render the page.

Problem: For pages where we want to have multiple post requests available without massive UI penalties - aka an author or piece of content route, with multiple buttons for user interaction - managing backend rendering logic becomes an absolute nightmare. Plus, for key functionality, like creating content or making a donation, users will need to have javascript enabled anyway, in order to create content via the Lexical Editor or post a donation using the Stripe Api.

Aka, just default to using JavaScript.

We got there folks!

# Database index + view increment strategy 

Database rank index update strategy

Have several fields that keep track of pending updates:

Pending rank change - for changing the total rank. Keeps track of everything - views, flags, likes, etc

Pending view change - keeps track of just view change to calculate rel rank

Pending positive change - keeps track of all non user positive interactions and the impact they will have on the change

When an update runs, check the values of these fields. If they cause a shift in rank or rel rank that reaches a certain percentile threshold change of the record’s relative rank or absolute rank, trigger an update on the rank itself, which will cause an index update.

If the rank change is not significant enough to trigger an update for this item, it simply won’t, preventing excessive high writes to the index

Ensure that this logic runs in the database itself such that the database’s queuing mechanisms work correctly. Running this logic api side will result in multiple attempts at a parallel write, which will cause conflict.

For the view counter, set up Valkey to increment the writes, then use a lua script as a way to check and flush the writes to a postgres instance when they reach a certain threshold.
