document.addEventListener("DOMContentLoaded", () => {
  console.log("running from signup page");

  // ========= Declaring Global Variables =========

  const csrfToken = document.getElementsByName("gorilla.csrf.Token")[0].value;

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
  const submitOtpErrorBlock = document.getElementById("submitOtpErrorBlock");

  const resetFlowButton = document.getElementById("reset_flow_button");

  // ========= Adding Event Listeners =======

  requestOtpForm.addEventListener("submit", handleRequestOtpForm);

  submitOtpForm.addEventListener("submit", handleSubmitOtpForm);

  resetFlowButton.addEventListener("click", resetFlow);

  // ============ Request Otp Form Functionality =========

  function handleRequestOtpForm(e) {
    e.preventDefault();
    if (!requestOtpMessageBlock) {
      console.error(
        "request otp error block is not valid HTML!",
        requestOtpMessageBlock,
      );
      return;
    }
    requestOtpMessageBlock.textContent = "Processing request...";

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
        throw new Error(errorMessage);
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

    if (!requestOtpMessageBlock) {
      console.error(
        "requestOtpMessageBlock is not valid HTML element!",
        requestOtpMessageBlock,
      );
      return;
    }
    requestOtpMessageBlock.textContent = errorMessage;
  }

  function renderOtpInputForm() {
    if (!requestOtpFormFieldset) {
      console.error(
        "requestOtpFormFieldset is not valid html!",
        requestOtpFormFieldset,
      );
      return;
    }
    requestOtpFormFieldset.disabled = true;

    if (!requestOtpMessageBlock) {
      console.error(
        "requestOptMessage block is not valid html",
        requestOtpMessageBlock,
      );
      return;
    }
    requestOtpMessageBlock.textContent = "";

    if (!submitOtpContainer) {
      console.error(
        "submit otp container is not valid html",
        submitOtpContainer,
      );
      return;
    }
    submitOtpContainer.hidden = false;
  }

  // ========== submit otp functionality ===========

  function handleSubmitOtpForm(e) {
    e.preventDefault();

    if (!otpInput) {
      console.error("otp input is not valid html", otpInput);
      return;
    }
    const otp = otpInput.value;

    if (!emailInput) {
      console.error("email input is not valid html!", emailInput);
      return;
    }
    const email = emailInput.value;

    submitOtp(email, otp);
  }

  async function submitOtp(email, otp) {
    const requestBody = {
      otp: otp,
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
      const res = await fetch("/signup/submit-otp", requestObject);
      if (!res.ok) {
        const errorMessage = await res.text();
        throw new Error(errorMessage);
      }
    } catch (error) {
      console.error(error);
      renderSubmitOtpErrors(error.message);
    }
  }

  function renderSubmitOtpErrors(errorMessage) {
    if (!errorMessage || typeof errorMessage !== "string") {
      console.error("error is not valid: error ===", errorMessage);
      return;
    }

    if (!submitOtpErrorBlock) {
      console.error(
        "submitOtpErrorBlock is not valid HTML element!",
        submitOtpErrorBlock,
      );
      return;
    }
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
