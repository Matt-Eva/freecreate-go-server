document.addEventListener("DOMContentLoaded", () => {
  console.log("running my creators page");
  const csrfToken = document.getElementsByName("gorilla.csrf.Token")[0].value;
  if (!csrfToken) {
    console.error("could not retrieve csrfToken!", csrfToken);
    return;
  }

  const createCreatorForm = document.getElementById("create_creator_form");
  const createCreatorNameInput = document.getElementById(
    "create_creator_name_input",
  );
  const createCreatorMessageBlock = document.getElementById(
    "create_creator_message_block",
  );

  createCreatorForm.addEventListener("submit", handleCreateCreator);

  function handleCreateCreator(e) {
    e.preventDefault();

    createCreatorMessageBlock.textContent = "";

    const newCreatorName = createCreatorNameInput.value;
    postCreator(newCreatorName);
  }

  async function postCreator(name) {
    if (!name) {
      renderCreateCreatorMessage("Name cannot be empty.");
      return;
    }

    const requestBody = {
      name: name,
    };

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
    } catch (error) {
      console.error(error);
      renderCreateCreatorMessage(error.message);
    }
  }

  function renderCreateCreatorMessage(message) {
    createCreatorMessageBlock.textContent = message;
  }
});
