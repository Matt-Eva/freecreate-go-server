document.addEventListener("DOMContentLoaded", (e) => {
  console.log("profile page running");

  const csrfToken = document.getElementsByName("gorilla.csrf.Token")[0].value;
  if (!csrfToken) {
    console.error("could not retrieve csrfToken!", csrfToken);
    return;
  }

  const logoutButton = document.getElementById("logout_button");
  const logoutErrorBlock = document.getElementById("logout_error_block");

  logoutButton.addEventListener("click", logout);

  async function logout() {
    logoutErrorBlock.textContent = "";

    const requestObject = {
      method: "DELETE",
      headers: {
        "X-CSRF-Token": csrfToken,
      },
    };

    try {
      const res = await fetch("/web-api/logout", requestObject);
      if (!res.ok) {
        const err = await res.text();
        throw new Error(err);
      } else if (res.redirected) {
        window.location.href = res.url;
      }
    } catch (error) {
      console.error(error);
      const msg = error.message;
      logoutErrorBlock.textContent = msg;
    }
  }
});
