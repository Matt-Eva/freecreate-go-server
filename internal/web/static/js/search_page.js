document.addEventListener("DOMContentLoaded", () => {
  console.log("running search page");

  // Get csrf token

  const csrfToken = document.getElementsByName("gorilla.csrf.Token")[0].value;
  if (!csrfToken) {
    console.error("could not retrieve csrfToken!", csrfToken);
    return;
  }

  // declare global search state

  const tags = [];

  // declare dom elements variables

  const searchTypeSelect = document.getElementById("search_type_select");
  const writingTypeSelect = document.getElementById("writing_type_select");
  const datePostedSelect = document.getElementById("date_posted_select");
  const topicCheckboxes = document.getElementsByClassName("topic_checkbox");
  const addTagForm = document.getElementById("add_tag_form");
  const addTagInput = document.getElementById("add_tag_input");
  const addedTagBlock = document.getElementById("added_tag_block");

  const searchSubmitForm = document.getElementById("search_submit_form");

  // event listeners

  addTagForm.addEventListener("submit", addTag);
  searchSubmitForm.addEventListener("submit", search);

  // populate global search state

  function populateSearchState() {}

  populateSearchState();

  // handle add tag

  function addTag(e) {
    e.preventDefault();

    const tag = addTagInput.value;

    addTagInput.value = "";

    if (tags.find((t) => t === tag)) {
      return;
    }

    tags.push(tag);

    const span = document.createElement("span");
    span.textContent = tag;

    addedTagBlock.append(span);
  }

  // handle search

  function search(e) {
    e.preventDefault();

    const searchTypeString = "searchType=" + searchTypeSelect.value;

    const writingTypeString = "writingType=" + writingTypeSelect.value;

    const datePostedString = "datePosted=" + datePostedSelect.value;

    let topicString = "";
    let topicCounter = 1;

    for (const checkbox of topicCheckboxes) {
      if (checkbox.checked) {
        if (topicCounter === 1) {
          topicString = topicString + "topicOne=" + checkbox.value + "&";
        } else if (topicCounter === 2) {
          topicString = topicString + "topicTwo=" + checkbox.value + "&";
        } else if (topicCounter === 3) {
          topicString = topicString + "topicThree=" + checkbox.value + "&";
          break;
        }
        topicCounter += 1;
      }
    }

    let tagString = "tags=";

    tags.forEach((tag, index) => {
      if (index < tags.length - 1) {
        tagString = tagString + tag + "&";
      } else {
        tagString = tagString + tag;
      }
    });

    console.log(tagString);

    window.location.href =
      "/search?" +
      searchTypeString +
      "&" +
      writingTypeString +
      "&" +
      datePostedString +
      "&" +
      topicString +
      "&" +
      tagString;
  }
});
