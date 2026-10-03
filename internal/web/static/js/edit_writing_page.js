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
  const editWritingDiv = document.getElementById("edit_writing_div");
  const submitUpdateButton = document.getElementById("submit_update_button");
  const writingTitleInput = document.getElementById("writing_title_input");
  const writingSubtitleInput = document.getElementById(
    "writing_subtitle_input",
  );
  const creatorSelect = document.getElementById("creator_select");
  const writingTypeSelect = document.getElementById("writing_type_select");
  const radioIsAdult = document.getElementById("is_adult_radio_true");
  const radioAllAges = document.getElementById("is_adult_radio_false");
  const descriptionInput = document.getElementById("writing_description_input");
  const writingUuid = document.getElementById("writing_uuid").value;
  const tagContainer = document.getElementById("tag_container");
  const addTagForm = document.getElementById("add_tag_form");
  const tagInput = document.getElementById("tag_input");
  const udpateWritingMessageBlock = document.getElementById(
    "update_writing_message_block",
  );

  // Event Listeners

  enableEditButton.addEventListener("click", enableEdit);

  cancelEditButton.addEventListener("click", cancelEdit);

  submitUpdateButton.addEventListener("click", handleUpdateWriting);

  addTagForm.addEventListener("submit", addTag);

  // Starting State

  const startingState = {
    title: writingTitleInput.value,
    subtitle: writingSubtitleInput.value,
    creator: creatorSelect.value,
    writingType: writingTypeSelect.value,
    adult: radioIsAdult.checked,
    allAges: radioAllAges.checked,
    description: descriptionInput.value,
    tags: [],
  };

  for (const child of tagContainer.children) {
    const tag = child.textContent;
    startingState.tags.push(tag);
  }

  // Tag state

  let tags = [];

  // Enable Edit

  function enableEdit() {
    const children = editWritingDiv.children;
    cancelEditButton.hidden = false;
    enableEditButton.hidden = true;
    for (const child of children) {
      child.disabled = false;
      if (child.children.length > 0) {
        console.log(child.children);
        for (const chi of child.children) {
          chi.disabled = false;
        }
      }
    }
  }

  // CancelEdit

  function cancelEdit() {
    const children = editWritingDiv.children;
    cancelEditButton.hidden = true;
    enableEditButton.hidden = false;
    for (const child of children) {
      child.disabled = true;
      if (child.children.length > 0) {
        for (const chi of child.children) {
          chi.disabled = true;
        }
      }
    }

    writingTitleInput.value = startingState.title;
    writingSubtitleInput.value = startingState.subtitle;
    creatorSelect.value = startingState.creator;
    writingTypeSelect.value = startingState.writingType;
    radioIsAdult.checked = startingState.adult;
    radioAllAges.checked = startingState.allAges;
    descriptionInput.value = startingState.description;
  }

  // Add Tag

  function addTag(e) {
    e.preventDefault();
    if (tagInput.value === "") return;
    for (const tag of tags) {
      if (tagInput.value === tag) {
        tagInput.value = "";
        return;
      }
    }
    const rawTag = tagInput.value;
    const tag = rawTag.replace(" ", "-");
    tags.push(tag);

    const tagSpan = document.createElement("span");
    tagSpan.textContent = tag;
    tagSpan.addEventListener("click", () => {
      tags = tags.filter((tag) => tag !== tagSpan.textContent);
      tagSpan.remove();
    });

    tagContainer.append(tagSpan);

    tagInput.value = "";
  }

  // UpdateWriting

  function handleUpdateWriting(e) {
    const creator = creatorSelect.value;
    const title = writingTitleInput.value;
    const subtitle = writingSubtitleInput.value;
    const writingType = writingTypeSelect.value;
    const isAdult = radioIsAdult.checked;
    const description = descriptionInput.value;
    const topicCheckboxes = document.getElementsByClassName("topic_checkbox");
    const topics = [];
    for (const checkbox of topicCheckboxes) {
      if (checkbox.checked) {
        topics.push(checkbox.value);
      }
    }

    updateWriting(
      creator,
      title,
      subtitle,
      writingType,
      isAdult,
      description,
      topics,
      tags,
    );
  }

  async function updateWriting(
    creator,
    title,
    subtitle,
    writingType,
    isAdult,
    description,
    topics,
    tags,
  ) {
    const requestBody = {
      creator,
      title,
      subtitle,
      writingType,
      isAdult,
      description,
      topics,
      tags,
    };

    const requestObject = {
      method: "PATCH",
      headers: {
        "Content-Type": "application/json",
        "X-CSRF-Token": csrfToken,
      },
      body: JSON.stringify(requestBody),
    };

    try {
      const res = await fetch(`/web-api/writing/${writingUuid}`, requestObject);
      if (!res.ok) {
        const error = await res.text();
        throw new Error(error);
      } else {
        const data = await res.json();
        console.log(data);
      }
    } catch (e) {
      console.error(e);
      udpateWritingMessageBlock.textContent = e.message;
    }
  }
});
