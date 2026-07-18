document.addEventListener("DOMContentLoaded", () => {
  console.log("running from signup page");

  const csrfToken = document.getElementsByName("gorilla.csrf.Token")[0].value;
  const requestOtpForm = document.getElementById("signup_request_otp_form");
  const requestOtpErrorBlock = document.getElementById("requestOtpErrorBlock");

  requestOtpForm.addEventListener("submit", handleRequestOtpForm);

  function handleRequestOtpForm(e) {
    e.preventDefault();

    const emailInput = document.getElementById("enter_email");
    if (!emailInput) {
      console.error("emailInputElement is null!");
      return;
    }

    const emailValue = emailInput.value;
    if (!emailValue) {
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
        const errorMessage = await res.text();
        throw new Error(`\n status: ${res.status} \n message: ${errorMessage}`);
      }

      renderOtpInputForm();
    } catch (error) {
      console.error(error);
      renderRequestOtpErrors(error.message);
    }
  }

  function renderRequestOtpErrors(errorMessage) {
    if (!errorMessage || typeof errorMessage !== "string") {
      console.error("error is not valid: error ===", errorMessage);
      return;
    }

    const errorP = document.createElement("p");
    if (!errorP) {
      console.error("could not create error paragraph!");
      return;
    }

    errorP.value = errorMessage;

    requestOtpErrorBlock.append(errorP);
  }

  function renderOtpInputForm() {
    console.log("rendering input form");
  }
});
