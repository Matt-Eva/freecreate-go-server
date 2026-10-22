document.addEventListener("DOMContentLoaded", () => {
  console.log("running edit writing page");

  // Get CSRF Token

  const csrfToken = document.getElementsByName("gorilla.csrf.Token")[0].value;
  if (!csrfToken) {
    console.error("could not retrieve csrfToken!", csrfToken);
    return;
  }

  // Writing UUID

  const writingUuid = document.getElementById("writing_uuid").value;

  // ========== EDIT WRITING SECTION ===================================

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

  const topicsContainer = document.getElementById("topics_container");
  const tagContainer = document.getElementById("tag_container");
  const addTagForm = document.getElementById("add_tag_form");
  const tagInput = document.getElementById("tag_input");
  const udpateWritingMessageBlock = document.getElementById(
    "update_writing_message_block",
  );
  const tagSpans = document.getElementsByClassName("tag_span");

  // Event Listeners

  enableEditButton.addEventListener("click", enableEdit);

  cancelEditButton.addEventListener("click", cancelEdit);

  submitUpdateButton.addEventListener("click", handleUpdateWriting);

  addTagForm.addEventListener("submit", handleAddTag);

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
    topics: [],
  };

  function updateStartingState(data) {
    startingState.title = data.title;
    startingState.subtitle = data.subtitle;
    startingState.description = data.description;
    startingState.topics = data.topics;
    startingState.tags = data.tags;
    startingState.adult = data.isAdult;
    startingState.allAges = !data.isAdult;
    startingState.writingType = data.writingType;
    startingState.creator = data.creatorId;
  }

  for (const child of tagContainer.children) {
    const tag = child.textContent;
    startingState.tags.push(tag);
  }

  for (const child of topicsContainer.children) {
    if (child.checked) {
      startingState.topics.push(child.value);
    }
  }

  // Tag state

  let tags = [...startingState.tags];

  // Enable Edit

  function enableEdit() {
    const children = editWritingDiv.children;
    cancelEditButton.hidden = false;
    enableEditButton.hidden = true;
    for (const child of children) {
      child.disabled = false;
      if (child.children.length > 0) {
        for (const chi of child.children) {
          chi.disabled = false;
        }
      }
    }

    for (const element of tagSpans) {
      element.addEventListener("click", removeTag);
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
    tagInput.value = "";

    for (const child of topicsContainer.children) {
      const value = child.value;
      const inStartingState = startingState.topics.find(
        (topic) => topic === value,
      );
      if (inStartingState) {
        child.checked = true;
      } else {
        child.checked = false;
      }
    }

    tagContainer.innerHTML = "";
    tags = [...startingState.tags];
    for (const tag of tags) {
      addTag(tag);
    }

    for (const child of tagContainer.children) {
      child.removeEventListener("click", removeTag);
    }
  }

  // Add Tag

  function handleAddTag(e) {
    e.preventDefault();
    udpateWritingMessageBlock.textContent = "";

    if (tagInput.value === "") return;
    for (const tag of tags) {
      if (tagInput.value === tag) {
        tagInput.value = "";
        return;
      }
    }

    if (tags.length >= 20) {
      udpateWritingMessageBlock.textContent = "You already have 20 tags.";
      tagInput.value = "";
      return;
    }

    const tag = addTag(tagInput.value);

    tags.push(tag);
  }

  function addTag(tagValue) {
    const tag = tagValue.replace(" ", "-");

    const tagSpan = document.createElement("span");
    tagSpan.textContent = tag;
    tagSpan.className = "tag_span";
    tagSpan.addEventListener("click", removeTag);

    tagContainer.append(tagSpan);

    tagInput.value = "";

    return tag;
  }

  // Remove Tag

  function removeTag(e) {
    tags = tags.filter((tag) => tag !== e.target.textContent);
    e.target.remove();
  }

  // UpdateWriting

  function handleUpdateWriting(e) {
    udpateWritingMessageBlock.textContent = "";

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
    const creatorId = parseInt(creator);

    const requestBody = {
      creatorId,
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
        updateStartingState(data);
        cancelEdit();
      }
    } catch (e) {
      console.error(e);
      udpateWritingMessageBlock.textContent = e.message;
    }
  }

  // ===================== ADD NEW CHAPTER SECTION =======================

  const newChapterForm = document.getElementById("new_chapter_form");
  const newChapterTitleInput = document.getElementById(
    "new_chapter_title_input",
  );
  const newChapterNumberInput = document.getElementById(
    "new_chapter_number_input",
  );
  const newChapterMessageBlock = document.getElementById(
    "new_chapter_message_block",
  );
  const chapterContainer = document.getElementById("chapter_container");

  newChapterForm.addEventListener("submit", addNewChapter);

  async function addNewChapter(e) {
    e.preventDefault();

    newChapterMessageBlock.textContent = "";

    const title = newChapterTitleInput.value;
    const chapterNumber = parseInt(newChapterNumberInput.value);

    if (title === "") {
      newChapterMessageBlock.textContent = "Title cannot be empty.";
    } else if (chapterNumber === "") {
      newChapterMessageBlock.textContent = "Chapter number cannot be empty.";
    }

    const requestBody = {
      title,
      chapterNumber,
      writingUuid,
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
      const res = await fetch("/web-api/chapter", requestObject);
      if (!res.ok) {
        const error = await res.text();
        throw new Error(error);
      } else if (res.redirected) {
        window.location.href = res.url;
      } else {
        const data = await res.json();

        newChapterTitleInput.value = "";
        newChapterNumberInput.value = "";

        renderNewChapter(data);
      }
    } catch (error) {
      console.error(error);
      newChapterMessageBlock.textContent = error.message;
    }
  }

  function renderNewChapter(data) {
    const chapterCard = document.createElement("article");

    const title = document.createElement("p");
    title.textContent = data.title;

    const number = document.createElement("p");
    number.textContent = "Chapter Number: " + data.chapterNumber;

    const published = document.createElement("p");
    if (data.published) {
      published.textContent = "Published";
    } else {
      published.textContent = "Unpublished";
    }

    chapterCard.append(title, number, published);

    chapterContainer.prepend(chapterCard);
  }

  // ====================== EDIT CHAPTER NUMBER SECTION ==================

  const updateChapterNumberForms = document.getElementsByClassName(
    "update_chapter_number_form",
  );

  const revealUpdateChapterNumberFormButtons = document.getElementsByClassName(
    "reveal_update_chapter_number_form_button",
  );

  for (const form of updateChapterNumberForms) {
    form.addEventListener("submit", handleUpdateChapterNumber);
  }

  for (const button of revealUpdateChapterNumberFormButtons) {
    button.addEventListener("click", toggleUpdateChapterNumberForm);
  }

  function toggleUpdateChapterNumberForm(e) {
    const formId = e.target.name + "_update_chapter_number_form";
    const updateChapterNumberForm = document.getElementById(formId);
    if (updateChapterNumberForm.hidden) {
      updateChapterNumberForm.hidden = false;
      e.target.textContent = "Cancel";
    } else {
      updateChapterNumberForm.hidden = true;
      updateChapterNumberForm["new_chapter_number_input"].value = "";
      e.target.textContent = "Edit Chapter Number";
    }
  }

  async function handleUpdateChapterNumber(e) {
    e.preventDefault();

    const chapterNumber = parseInt(e.target["new_chapter_number_input"].value);
    const chapterUuid = e.target["chapter_uuid_input"].value;

    const requestBody = { chapterNumber };

    const requestObject = {
      method: "PATCH",
      headers: {
        "Content-Type": "application/json",
        "X-CSRF-Token": csrfToken,
      },
      body: JSON.stringify(requestBody),
    };

    try {
      const res = await fetch(
        `/web-api/chapter/${chapterUuid}/number`,
        requestObject,
      );
      if (!res.ok) {
        const error = await res.text();
        throw new Error(error);
      } else if (res.redirected) {
        window.location.href = res.url;
      } else {
        const data = await res.json();
        const chapterNumber = document.getElementById(
          `${chapterUuid}_chapter_number`,
        );
        chapterNumber.textContent = data.newChapterNumber;
        e.target.hidden = true;
        e.target["new_chapter_number_input"].value = "";
        const revealButton = document.getElementById(
          `${chapterUuid}_reveal_update_chapter_number_form_button`,
        );
        revealButton.textContent = "Edit Chapter Number";
      }
    } catch (e) {
      console.error(e);
    }
  }
});
