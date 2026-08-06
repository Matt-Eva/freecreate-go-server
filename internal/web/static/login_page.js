document.addEventListener("DOMContentLoaded", () => {
  console.log("running login page");

  const csrfToken = document.getElementsByName("gorilla.csrf.Token")[0].value;
  if (!csrfToken) {
    console.error("could not retrieve csrfToken!", csrfToken);
    return;
  }

  const requestOtpForm = document.getElementById("request_otp_form");
  const requestOtpFieldset = document.getElementById("request_otp_fieldset");
  const emailInput = document.getElementById("email_input");
  const requestOtpErrBlock = document.getElementById("request_otp_err_block");

  const submitOtpContainer = document.getElementById("submit_otp_container");
  const submitOtpForm = document.getElementById("submit_otp_form");
  const otpInput = document.getElementById("otp_input");
  const submitOtpErrBlock = document.getElementById("submit_otp_err_block");
  const resetButton = document.getElementById("reset_button");

  requestOtpForm.addEventListener("submit", handleRequestOtpForm);

  submitOtpForm.addEventListener("submit", handleSubmitOtpForm);

  resetButton.addEventListener("click", resetFlow);

  function handleRequestOtpForm(e) {
    e.preventDefault();

    requestOtpErrBlock.textContent = "Processing request...";

    const email = emailInput.value;

    requestOtp(email);
  }

  async function requestOtp(email) {
    if (!email) {
      renderRequestOtpError("Email cannot be empty.");
      return;
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
      const res = await fetch("/web-api/login/request-otp", requestObject);

      if (!res.ok) {
        const err = await res.text();
        throw new Error(err);
      }

      renderSubmitOtpForm();
    } catch (error) {
      console.error(error);
      renderRequestOtpError(error.message);
    }
  }

  function renderRequestOtpError(error) {
    requestOtpErrBlock.textContent = error;
  }

  function renderSubmitOtpForm() {
    requestOtpFieldset.disabled = true;
    requestOtpErrBlock.textContent = "";
    submitOtpContainer.hidden = false;
  }

  function handleSubmitOtpForm(e) {
    e.preventDefault();

    submitOtpErrBlock.textContent = "Processing request...";

    const email = emailInput.value;

    const otp = otpInput.value;

    submitOtp(otp, email);
  }

  async function submitOtp(otp, email) {
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
        "Content-Type": "application/json",
        "X-CSRF-Token": csrfToken,
      },
      body: JSON.stringify(requestBody),
    };

    try {
      const res = await fetch("/web-api/login/submit-otp", requestObject);
      if (!res.ok) {
        const error = await res.text();
        throw new Error(error);
      } else if (res.redirected) {
        window.location.href = res.url;
      }
    } catch (error) {
      console.error(error);
      renderSubmitOtpError(error.message);
    }
  }

  function renderSubmitOtpError(error) {
    submitOtpErrBlock.text = error;
  }

  function resetFlow() {
    submitOtpContainer.hidden = true;
    otpInput.value = "";
    submitOtpErrBlock.textContent = "";

    requestOtpFieldset.disabled = false;
    emailInput.value = "";
    requestOtpErrBlock.textContent = "";
  }
});
