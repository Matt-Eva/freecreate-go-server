document.addEventListener("DOMContentLoaded", () => {
  console.log("edit profile page running");

  // ============== CSRF Token ===============

  const csrfToken = document.getElementsByName("gorilla.csrf.Token")[0].value;
  if (!csrfToken) {
    console.error("could not retrieve csrfToken!", csrfToken);
    return;
  }

  // ============== DOM Elements ============
  const updateProfileForm = document.getElementById("update_profile_form");

  const updateProfileMessageBlock = document.getElementById(
    "update_profile_message_block",
  );
  const cancelUpdateBtn = document.getElementById("cancel_update_btn");

  // ============== Event Listeners =========

  updateProfileForm.addEventListener("submit", handleUpdateProfile);

  cancelUpdateBtn.addEventListener("click", cancelUpdate);

  // ============== Update Profile ===========

  function handleUpdateProfile(e) {
    e.preventDefault();

    updateProfileMessageBlock.textContent = "";

    const username = document.getElementById("update_username_input").value;

    const userHandle = document.getElementById("update_userhandle_input").value;

    let isAdult = false;
    const isAdultInput = document.querySelector(
      "input[name='is_adult']:checked",
    );
    if (isAdultInput.value === "yes") {
      isAdult = true;
    }

    let readingHistory = false;
    const readingHistoryInput = document.querySelector(
      "input[name='reading_history']:checked",
    );
    if (readingHistoryInput.value === "on") {
      readingHistory = true;
    }

    updateProfile(username, userHandle, isAdult, readingHistory);
  }

  async function updateProfile(username, userHandle, isAdult, readingHistory) {
    const requestBody = {
      username,
      userHandle,
      isAdult,
      readingHistory,
    };

    const requestObject = {
      method: "PATCH",
      headers: {
        "Conte-Type": "application/json",
        "X-CSRF-Token": csrfToken,
      },
      body: JSON.stringify(requestBody),
    };

    try {
      const res = await fetch("/web-api/user/profile", requestObject);
      if (!res.ok) {
        const error = await res.text();
        throw new Error(error);
      } else if (res.redirected) {
        window.location.href = res.url;
      } else {
        window.location.href = "/profile";
      }
    } catch (error) {
      console.error(error);
      updateProfileMessageBlock.textContent = error.message;
    }
  }

  // ========= Cancel Update ==============

  function cancelUpdate() {
    window.location.href = "/profile";
  }
});
