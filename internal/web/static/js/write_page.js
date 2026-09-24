document.addEventListener("DOMContentLoaded", () => {
  console.log("running write page");
  // Get CSRF Token
  const csrfToken = document.getElementsByName("gorilla.csrf.Token")[0].value;
  if (!csrfToken) {
    console.error("could not retrieve csrfToken!", csrfToken);
    return;
  }

  // Get DOM Elements
  const newWritingForm = document.getElementById("new_writing_form");
  const newWritingTitleInput = document.getElementById("new_writing_title");
  const newWritingCreatorSelect = document.getElementById(
    "new_writing_creator",
  );
  const newWritingTypeSelect = document.getElementById("new_writing_type");

  const newWritingMessageBlock = document.getElementById(
    "new_writing_message_block",
  );

  // Event Listeners

  newWritingForm.addEventListener("submit", handleNewWriting);

  // Create Writing

  function handleNewWriting(e) {
    e.preventDefault();
    console.log("handling submit");

    newWritingMessageBlock.textContent = "";

    const title = newWritingTitleInput.value;

    const creator = parseInt(newWritingCreatorSelect.value);

    const writingType = newWritingTypeSelect.value;

    createWriting(title, creator, writingType);
  }

  async function createWriting(title, creator, writingType) {
    console.log(title);
    console.log(creator);
    console.log(writingType);

    const requestBody = {
      title,
      creator,
      writingType,
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
      const res = await fetch("/web-api/writing", requestObject);
      if (!res.ok) {
        const error = await res.text();
        throw new Error(error);
      } else {
        const data = await res.json();
        console.log(data);
        window.location.href = "";
      }
    } catch (error) {
      console.error(error);
      newWritingMessageBlock.textContent = error.message;
    }
  }
});
