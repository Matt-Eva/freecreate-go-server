import checkTagLength from "./lib/check_tag_length.js";

document.addEventListener("DOMContentLoaded", () => {
  console.log("running search page");

  // Get csrf token

  const csrfToken = document.getElementsByName("gorilla.csrf.Token")[0].value;
  if (!csrfToken) {
    console.error("could not retrieve csrfToken!", csrfToken);
    return;
  }

  // declare dom elements variables

  const searchTypeSelect = document.getElementById("search_type_select");
  const writingTypeSelect = document.getElementById("writing_type_select");
  const datePostedSelect = document.getElementById("date_posted_select");
  const topicCheckboxes = document.getElementsByClassName("topic_checkbox");
  const allTopicsContainer = document.getElementById("all_topics_container");
  const fictionTopicsContainer = document.getElementById(
    "fiction_topics_container",
  );
  const essayAndBlogTopicsContainer = document.getElementById(
    "essay_and_blog_topics_container",
  );
  const memoirAndAutobiograhyTopicsContainer = document.getElementById(
    "memoir_and_autobiography_topics_container",
  );
  const poetryTopicsContainer = document.getElementById(
    "poetry_topics_container",
  );
  const addTagForm = document.getElementById("add_tag_form");
  const addTagInput = document.getElementById("add_tag_input");
  const addTagMessageBlock = document.getElementById("add_tag_message_block");
  const addedTagBlock = document.getElementById("added_tag_block");

  const searchSubmitForm = document.getElementById("search_submit_form");

  // declare global search state

  let tags = [];

  for (const child of addedTagBlock.children) {
    tags.push(child.textContent);
    child.addEventListener("click", removeTag);
  }

  let writingType = writingTypeSelect.value;

  // event listeners

  writingTypeSelect.addEventListener("change", handleWritingTypeChange);

  addTagForm.addEventListener("submit", addTag);

  searchSubmitForm.addEventListener("submit", search);

  // populate global search state

  function populateSearchState() {}

  populateSearchState();

  // handle writing type change

  function handleWritingTypeChange(e) {
    allTopicsContainer.hidden = true;
    fictionTopicsContainer.hidden = true;
    essayAndBlogTopicsContainer.hidden = true;
    memoirAndAutobiograhyTopicsContainer.hidden = true;
    poetryTopicsContainer.hidden = true;

    if (e.target.value === "All") {
      allTopicsContainer.hidden = false;
    } else if (
      e.target.value === "Fiction" ||
      e.target.value === "Short Stories" ||
      e.target.value === "Novelettes" ||
      e.target.value === "Novellas" ||
      e.target.value === "Novels"
    ) {
      fictionTopicsContainer.hidden = false;
      if (
        writingType === "Fiction" ||
        writingType === "Short Stories" ||
        writingType === "Novelettes" ||
        writingType === "Novellas" ||
        writingType === "Novels"
      ) {
        writingType = e.target.value;
        return;
      }
    } else if (e.target.value === "Essays and Blogs") {
      essayAndBlogTopicsContainer.hidden = false;
    } else if (e.target.value === "Memoir and Autobiography") {
      memoirAndAutobiograhyTopicsContainer.hidden = false;
    } else if (e.target.value === "Poetry") {
      poetryTopicsContainer.hidden = false;
    }
    writingType = e.target.value;
    for (const checkbox of topicCheckboxes) {
      checkbox.checked = false;
    }
  }

  // handle add tag

  function addTag(e) {
    e.preventDefault();

    let tag = addTagInput.value;

    addTagMessageBlock.textContent = "";

    if (checkTagLength(tag)) {
      addTagMessageBlock.textContent =
        "Tag length cannot be longer than 30 charaters.";
      return;
    }

    tag = tag.replaceAll(" ", "-");

    addTagInput.value = "";

    if (tags.length >= 5) {
      return;
    }

    if (tags.find((t) => t === tag)) {
      return;
    }

    tags.push(tag);

    const span = document.createElement("span");
    span.className = "tag_span";
    span.textContent = tag;

    span.addEventListener("click", removeTag);

    addedTagBlock.append(span);
  }

  // handle remove tag

  function removeTag(e) {
    tags = tags.filter((tag) => tag !== e.target.textContent);
    e.target.remove();
  }

  // handle search

  function search(e) {
    e.preventDefault();

    const searchTypeString = "searchType=" + searchTypeSelect.value;

    const writingTypeString = "writingType=" + writingTypeSelect.value;

    const datePostedString = "datePosted=" + datePostedSelect.value;

    const topics = [];

    for (let i = 0; i < topicCheckboxes.length; i++) {
      const checkbox = topicCheckboxes[i];
      if (checkbox.checked === true && topics.length < 3) {
        topics.push(checkbox.value);
      }
    }

    let topicString = "";

    topics.forEach((topic, index) => {
      topicString += "topic=" + topic;
      if (index < topics.length - 1) {
        topicString += "&";
      }
    });

    let tagString = "";

    tags.forEach((tag, index) => {
      tagString = tagString + "tag=" + tag;

      if (index < tags.length - 1) {
        tagString += "&";
      }
    });

    let searchString =
      "/search?" +
      searchTypeString +
      "&" +
      writingTypeString +
      "&" +
      datePostedString;

    if (topicString.length > 0) {
      searchString += "&" + topicString;
    }

    if (tagString.length > 0) {
      searchString += "&" + tagString;
    }

    window.location.href = searchString;
  }
});
