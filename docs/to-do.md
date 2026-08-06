# To-Do

- Setup Login Route
- Reformat Auth Flow [x]
  - We can have two types of sessions
    - One is the guest session, which is only used by the login and signup pages for storing the otp and creating a session uuid for doing so
    - the other is the auth session, which is created once the user acutally logs in
    - Instead of having "check login" or "get session" functionality, we just have a single get user function.
      - Yes, we will run it for every page for rendering purposes, resulting in greater queries to the valkey cache, but it will simplify everything
      - If an error occurs during the get user function, we want to invalidate the whole session.
- Set Up Email sender [x]
- Create OTP Generator for Auth [x]
- Store UserID in Valkey using Session ID [x]
- Set Up Create User Route [x]
