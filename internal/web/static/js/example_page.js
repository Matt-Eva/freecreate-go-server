// While we should be including defer within all of our script tags in our header,
// we will also wrap all of our functionality with a "DOMContentLoaded" event listener, to help
// ensure we don't accidentally create global variables that may pollute across files

document.addEventListener("DOMContentLoaded", () => {
  console.log("hello from example js!");

  // Get CSRF Token from dom
  const csrfToken = document.getElementsByName("gorilla.csrf.Token")[0].value;
  if (!csrfToken) {
    console.error("could not retrieve csrfToken!", csrfToken);
    return;
  }

  //   Declare global variables

  const examplePostForm = document.getElementById("example_post_form");
  const examplePostFormInput = document.getElementById(
    "example_post_form_input",
  );
  const examplePostFormMessageBlock = document.getElementById(
    "example_post_form_message_block",
  );
  const fetchedDataContainer = document.getElementById(
    "fetched_data_container",
  );

  //   Declare Event Listeners

  examplePostForm.addEventListener("submit", handleExamplePost);

  // Declare Function Section - Handle Example Post Form

  function handleExamplePost(e) {
    e.preventDefault();

    const postInput = examplePostFormInput.value;

    examplePostFormMessageBlock.textContent = "";

    post(postInput);
  }

  async function post(postInput) {
    if (!postInput) {
      const msg = "post input cannot be empty.";
      console.error(msg, postInput);
      renderPostFormMessage;
      return;
    }

    const requestBody = {
      postInput: postInput,
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
      const res = await fetch("/web-api/example", requestObject);

      if (!res.ok) {
        const error = await res.text();
        throw new Error(error);
      } else if (res.redirected) {
        window.location.href = res.url;
      } else {
        const data = await res.json();
        renderFetchData(data);
      }
    } catch (error) {
      console.error(error);
      renderPostFormMessage(error.message);
    }
  }

  function renderPostFormMessage(msg) {
    examplePostFormMessageBlock = msg;
  }

  function renderFetchData(data) {}
});
