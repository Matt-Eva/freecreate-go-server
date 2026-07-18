document.addEventListener("DOMContentLoaded", () => {
  console.log("running from signup page");

  const csrfToken = document.getElementsByName("gorilla.csrf.Token")[0].value;
  const requestOtpForm = document.getElementById("signup_request_otp_form");

  requestOtpForm.addEventListener("submit", handleRequestOtpForm);

  function handleRequestOtpForm(e) {
    e.preventDefault();

    const emailInput = document.getElementById("enter_email");
    if (emailInput === null) {
      console.error("emailInputElement is null!");
      return;
    }

    const emailValue = emailInput.value;
    if (emailValue === "") {
      console.error("emailValue is empty string!");
      return;
    }

    requestOtp(emailValue);
  }

  async function requestOtp(email) {
    const requestBody = {
      email: email,
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
      const res = await fetch("/web-api/signup/request-otp", requestObject);

      if (!res.ok) {
        const error = await res.text();
        console.error(error);
        return;
      }

      const data = await res.json();
      console.log(data);
    } catch (error) {
      console.error(error);
    }
  }

  function renderOtpInputForm() {}
});
