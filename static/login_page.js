document.addEventListener("DOMContentLoaded", () => {
  console.log("running login page");

  const csrfToken = document.getElementsByName("gorilla.csrf.Token")[0].value;
  if (!csrfToken) {
    console.error("could not retrieve csrfToken!", csrfToken);
    return;
  }

  const requestOtpForm = document.getElementById("request_otp_form");
  const requestOtpErrBlock = document.getElementById("request_otp_err_block");

  const submitOtpForm = document.getElementById("submit_otp_form");
  const submitOtpErrBlock = document.getElementById("submit_otp_err_block");

  requestOtpForm.addEventListener("submit", handleRequestOtpForm);

  submitOtpForm.addEventListener("submit", handleSubmitOtpForm);

  function handleRequestOtpForm(e) {
    e.preventDefault();
  }

  function requestOtp() {}

  function handleSubmitOtpForm(e) {
    e.preventDefault();
  }

  function submitOtp() {}
});
