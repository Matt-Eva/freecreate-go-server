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
  const addedTagBlock = document.getElementById("added_tag_block");

  const searchSubmitForm = document.getElementById("search_submit_form");

  // declare global search state

  let tags = [];

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
    span.textContent = tag;
    span.style.margin = "5px";
    span.style.height = "20px";
    span.style.display = "flex";
    span.style.padding = "0px 3px";
    span.style.textAlign = "center";
    span.style.border = "solid 1px";
    span.style.borderRadius = "2px";

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

    let topicString = "";
    let topicCounter = 1;

    for (const checkbox of topicCheckboxes) {
      if (checkbox.checked && topicCounter < 3) {
        topicString = topicString + "topic=" + checkbox.value + "&";
        topicCounter += 1;
      }
    }

    let tagString = "";

    tags.forEach((tag, index) => {
      tagString = tagString + "tag=" + tag + "&";
      // if (index < tags.length - 1) {
      // } else {
      //   tagString = tagString + tag;
      // }
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
