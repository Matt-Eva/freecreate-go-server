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
  const newWritingCreatorSelect = document.getElementById(
    "new_writing_creator",
  );
  const newWritingTypeSelect = document.getElementById("new_writing_type");

  // Event Listeners

  newWritingForm.addEventListener("submit", handleNewWriting);

  // Create Writing

  function handleNewWriting(e) {
    e.preventDefault();
    console.log("handling submit");

    const creator = newWritingCreatorSelect.value;

    const writingType = newWritingTypeSelect.value;

    console.log(creator);
    console.log(writingType);
  }
});
