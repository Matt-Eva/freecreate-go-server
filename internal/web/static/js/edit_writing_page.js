document.addEventListener("DOMContentLoaded", () => {
  console.log("running edit writing page");
  // Get CSRF Token
  const csrfToken = document.getElementsByName("gorilla.csrf.Token")[0].value;
  if (!csrfToken) {
    console.error("could not retrieve csrfToken!", csrfToken);
    return;
  }
});
