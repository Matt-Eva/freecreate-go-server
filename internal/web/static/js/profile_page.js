document.addEventListener("DOMContentLoaded", (e) => {
  console.log("profile page running");

  // ============== CSRF Token ===============

  const csrfToken = document.getElementsByName("gorilla.csrf.Token")[0].value;
  if (!csrfToken) {
    console.error("could not retrieve csrfToken!", csrfToken);
    return;
  }

  // ============= DOM Elements ==============

  const logoutButton = document.getElementById("logout_button");
  const logoutErrorBlock = document.getElementById("logout_error_block");

  const newCreatorForm = document.getElementById("new_creator_form");
  const newCreatorNameInput = document.getElementById("new_creator_name_input");
  const newCreatorHandleInput = document.getElementById(
    "new_creator_handle_input",
  );
  const newCreatorMessageBlock = document.getElementById(
    "new_creator_message_block",
  );

  // ============ Event Listeners ================

  logoutButton.addEventListener("click", logout);

  newCreatorForm.addEventListener("submit", handleCreateCreator);

  // =========== Logout functionality ============

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

  // === new Creator Form functionality ===

  function handleCreateCreator(e) {
    e.preventDefault();

    newCreatorMessageBlock.textContent = "";

    const newCreatorName = newCreatorNameInput.value;
    const newCreatorHandle = newCreatorHandleInput.value;
    postCreator(newCreatorName, newCreatorHandle);
  }

  async function postCreator(name, handle) {
    if (!name || !handle) {
      renderCreateCreatorMessage("Name and handle cannot be empty.");
      return;
    }

    const requestBody = {
      name: name,
      handle: handle,
    };

    console.log(requestBody);

    const requestObject = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "X-CSRF-Token": csrfToken,
      },
      body: JSON.stringify(requestBody),
    };

    try {
      const res = await fetch("/web-api/creator", requestObject);
      if (!res.ok) {
        const err = await res.text();
        throw new Error(err);
      } else {
        const data = await res.json();
        console.log(data);
      }
    } catch (error) {
      console.error(error);
      renderCreateCreatorMessage(error.message);
    }
  }

  function renderCreateCreatorMessage(message) {
    newCreatorMessageBlock.textContent = message;
  }
});
