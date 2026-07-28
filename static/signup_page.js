document.addEventListener("DOMContentLoaded", () => {
  console.log("running from signup page");

  // ========= Declaring Global Variables =========

  const csrfToken = document.getElementsByName("gorilla.csrf.Token")[0].value;
  if (!csrfToken) {
    console.error("could not retrieve csrfToken!", csrfToken);
    return;
  }

  const requestOtpForm = document.getElementById("request_otp_form");
  const requestOtpFormFieldset = document.getElementById(
    "request_otp_form_fieldset",
  );
  const emailInput = document.getElementById("enter_email");
  const requestOtpMessageBlock = document.getElementById(
    "request_otp_error_block",
  );

  const submitOtpContainer = document.getElementById("submit_otp_container");
  const submitOtpForm = document.getElementById("submit_otp_form");
  const otpInput = document.getElementById("enter_otp");
  const submitOtpErrorBlock = document.getElementById("submit_otp_error_block");

  const resetFlowButton = document.getElementById("reset_flow_button");

  // ========= Adding Event Listeners =======

  requestOtpForm.addEventListener("submit", handleRequestOtpForm);

  submitOtpForm.addEventListener("submit", handleSubmitOtpForm);

  resetFlowButton.addEventListener("click", resetFlow);

  // ============ Request Otp Form Functionality =========

  function handleRequestOtpForm(e) {
    e.preventDefault();

    requestOtpMessageBlock.textContent = "Processing request...";

    const emailValue = emailInput.value;
    if (!emailValue) {
      console.error("emailValue is empty string!");
      return;
    }

    requestOtp(emailValue);
  }

  async function requestOtp(email) {
    if (!email) {
      renderRequestOtpErrors("Email cannot be empty.");
    }

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
        throw new Error(errorMessage);
      }

      renderOtpInputForm();
    } catch (error) {
      console.error(error);
      renderRequestOtpErrors(error.message);
    }
  }

  function renderRequestOtpErrors(errorMessage) {
    requestOtpMessageBlock.textContent = errorMessage;
  }

  function renderOtpInputForm() {
    requestOtpFormFieldset.disabled = true;
    requestOtpMessageBlock.textContent = "";
    submitOtpContainer.hidden = false;
  }

  // ========== submit otp functionality ===========

  function handleSubmitOtpForm(e) {
    e.preventDefault();

    submitOtpErrorBlock.textContent = "Processing request...";

    const otp = otpInput.value;

    const email = emailInput.value;

    submitOtp(email, otp);
  }

  async function submitOtp(email, otp) {
    if (!otp) {
      renderSubmitOtpError("One Time Password cannot be empty.");
      return;
    } else if (otp.length !== 8) {
      renderSubmitOtpError("One Time Password must be 8 characters in length.");
      return;
    }

    if (!email) {
      renderSubmitOtpError("Email cannot be empty.");
      return;
    }

    const requestBody = {
      otp: otp,
      email: email,
    };

    const requestObject = {
      method: "POST",
      headers: {
        "X-CSRF-Token": csrfToken,
        "Content-Type": "application/json",
      },
      body: JSON.stringify(requestBody),
    };

    try {
      const res = await fetch("/web-api/signup/submit-otp", requestObject);
      if (!res.ok) {
        const errorMessage = await res.text();
        throw new Error(errorMessage);
      } else if (res.redirected) {
        window.location.href = res.url;
      }
    } catch (error) {
      console.error(error);
      renderSubmitOtpErrors(error.message);
    }
  }

  function renderSubmitOtpErrors(errorMessage) {
    submitOtpErrorBlock.textContent = errorMessage;
  }

  // ========== reset page flow functionality ===========

  function resetFlow() {
    emailInput.value = "";
    requestOtpFormFieldset.disabled = false;
    submitOtpContainer.hidden = true;
    otpInput.value = "";
  }
});
