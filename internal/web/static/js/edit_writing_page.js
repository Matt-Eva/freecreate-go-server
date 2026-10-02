document.addEventListener("DOMContentLoaded", () => {
  console.log("running edit writing page");

  // Get CSRF Token

  const csrfToken = document.getElementsByName("gorilla.csrf.Token")[0].value;
  if (!csrfToken) {
    console.error("could not retrieve csrfToken!", csrfToken);
    return;
  }

  // DOM Elements

  const enableEditButton = document.getElementById("enable_edit_button");
  const cancelEditButton = document.getElementById("cancel_edit_button");
  const editWritingForm = document.getElementById("edit_writing_form");
  const writingTitleInput = document.getElementById("writing_title_input");
  const writingSubtitleInput = document.getElementById(
    "writing_subtitle_input",
  );
  const creatorSelect = document.getElementById("creator_select");
  const writingTypeSelect = document.getElementById("writing_type_select");

  // Event Listeners

  editWritingForm.addEventListener("submit", handleUpdateWriting);

  enableEditButton.addEventListener("click", enableEdit);

  cancelEditButton.addEventListener("click", cancelEdit);

  // Starting State

  const startingState = {
    title: writingTitleInput.value,
    subtitle: writingSubtitleInput.value,
    creator: creatorSelect.value,
    writingType: writingTypeSelect.value,
  };

  // Enable Edit

  function enableEdit() {
    const children = editWritingForm.children;
    cancelEditButton.hidden = false;
    enableEditButton.hidden = true;
    for (const child of children) {
      child.disabled = false;
    }
  }

  // CancelEdit

  function cancelEdit() {
    const children = editWritingForm.children;
    cancelEditButton.hidden = true;
    enableEditButton.hidden = false;
    for (const child of children) {
      child.disabled = true;
    }

    writingTitleInput.value = startingState.title;
    writingSubtitleInput.value = startingState.subtitle;
    creatorSelect.value = startingState.creator;
    writingTypeSelect.value = startingState.writingType;
  }

  // UpdateWriting

  function handleUpdateWriting(e) {
    e.preventDefault();
    console.log(creatorSelect.value);
    console.log(writingTitleInput.value);
    console.log(writingSubtitleInput.value);
    console.log(writingTypeSelect.value);
  }
});
