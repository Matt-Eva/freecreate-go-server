document.addEventListener("DOMContentLoaded", () => {
  const formOne = document.getElementById("form-1");
  const formTwo = document.getElementById("form-2");
  const csrfToken = document.getElementsByName("gorilla.csrf.Token")[0].value;
  console.log(formOne);
  console.log(formTwo);
  console.log(csrfToken);
  formOne.addEventListener("submit", (e) => {
    e.preventDefault();

    fetch("/web-api/test", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "X-CSRF-Token": csrfToken,
      },
      body: JSON.stringify({
        formAction: formOne["form_action"].value,
      }),
    })
      .then(async (res) => {
        if (res.ok) {
          const data = await res.json();
          console.log(data);
        } else {
          const errMessage = await res.text();
          console.error(res.status, errMessage);
        }
      })
      .catch(async (err) => {
        console.error(err);
      });
  });
  formTwo.addEventListener("submit", (e) => {
    e.preventDefault();
    console.log("form two default prevented");
    fetch("/web-api/test", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "X-CSRF-Token": csrfToken,
      },
      body: JSON.stringify({
        formAction: formTwo["form_action"].value,
      }),
    })
      .then(async (res) => {
        console.log(res);
        const data = await res.json();
        console.log(data);
      })
      .catch(async (err) => {
        const message = await err.text();
        console.log(message);
      });
  });
});
