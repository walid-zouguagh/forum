async function validlogin(e) {
  e.preventDefault();
  var email = document.getElementById("email");
  var password = document.getElementById("password");
  var emailRegex = /^[a-zA-Z0-9]{4,30}@[a-zA-Z]{3,8}[.][a-zA-Z]{3,5}/;
  var isValidusrname = true;
  var isValidpassword = true;
  if (email.value === "") {
    document.getElementById("erroremail").innerHTML = "email is null";
    document.getElementById("erroremail").style.color = "red";
    isValidusrname = false;
  } else if (email.value !== "" && !emailRegex.test(email.value)) {
    document.getElementById("erroremail").innerHTML = "email is not valid";
    document.getElementById("erroremail").style.color = "red";
    isValidusrname = false;
  } else {
    document.getElementById("erroremail").innerHTML = "";
    isValidusrname = true;
  }

  if (password.value === "") {
    document.getElementById("errorPassword").innerHTML = "Password is null";
    document.getElementById("errorPassword").style.color = "red";
    isValidpassword = false;
  } else {
    document.getElementById("errorPassword").innerHTML = "";
    isValidpassword = true;
  }

  if (isValidpassword && isValidusrname) {
    try {
      // Send the POST request
      const response = await fetch(`/user/login`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json; charset=UTF-8",
        },
        body: JSON.stringify({ email: email.value, password: password.value }),
      });
      console.log(
        JSON.stringify({ email: email.value, password: password.value })
      );
      if (response.ok) {
        window.location.href = "/";
      } else {
        const data = await response.json();
        document.getElementById("errorPassword").innerHTML = data.error;
        document.getElementById("errorPassword").style.color = "red";
      }
    } catch (error) {
      console.error("Network error:", error);
      alert("Unable to submit the comment due to a network issue.");
    }
  }
}
