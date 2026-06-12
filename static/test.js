document.addEventListener("DOMContentLoaded", () => {
  const formOne = document.getElementById("form-1");
  const formTwo = document.getElementById("form-2");
  const csrfToken = document.getElementsByName("gorilla.csrf.Token")[0].value;
  console.log(formOne);
  console.log(formTwo);
  console.log(csrfToken);
  formOne.addEventListener("submit", (e) => {
    e.preventDefault();
    console.log("form one default prevented");
    fetch("/test", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "X-CSRF-Token": csrfToken,
      },
    })
      .then(async (res) => {
        console.log(res);
        const data = await res.json();
        console.log(data);
      })
      .catch((err) => {
        console.error(err);
      });
  });
  formTwo.addEventListener("submit", (e) => {
    e.preventDefault();
    console.log("form two default prevented");
    fetch("/test", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "X-CSRF-Token": csrfToken,
      },
    })
      .then(async (res) => {
        console.log(res);
        const data = await res.json();
        console.log(data);
      })
      .catch((err) => {
        console.error(err);
      });
  });
});
