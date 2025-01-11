async function validregister(e) {
  e.preventDefault();

  var username = document.getElementById("username");
  var email = document.getElementById("email");
  var password = document.getElementById("password");
  var password_config = document.getElementById("password_config");

  var emailRegex = /^[a-zA-Z0-9]{4,30}@[a-zA-Z]{3,8}[.][a-zA-Z]{3,5}/;

  var isValidusername = true;
  var isValidemail = true;
  var isValidpassword = true;

  if (username.value === "") {
    document.getElementById("errorUsername").innerHTML = "username is null";
    document.getElementById("errorUsername").style.color = "red";
    isValidusername = false;
  } else {
    document.getElementById("errorUsername").innerHTML = "";
    isValidusername = true;
  }

  if (email.value === "") {
    document.getElementById("erroremail").innerHTML = "email is null";
    document.getElementById("erroremail").style.color = "red";
    isValidemail = false;
  } else if (email.value !== "" && !emailRegex.test(email.value)) {
    document.getElementById("erroremail").innerHTML = "email is not valid";
    document.getElementById("erroremail").style.color = "red";
    isValidemail = false;
  } else {
    document.getElementById("erroremail").innerHTML = "";
    isValidemail = true;
  }

  if (password.value === "") {
    document.getElementById("errorPassword").innerHTML =
      "password is not valid";
    document.getElementById("errorPassword").style.color = "red";
    isValidpassword = false;
  } else {
    document.getElementById("errorPassword").innerHTML = "";
  }

  if (password_config.value === "") {
    document.getElementById("errorPassword_config").innerHTML =
      "password config is not valid";
    document.getElementById("errorPassword_config").style.color = "red";
    isValidpassword = false;

    return;
  } else {
    document.getElementById("errorPassword_config").innerHTML = "";
  }

  if (password.value !== password_config.value) {

    document.getElementById("errorPassword_config").innerHTML = "password is not config";
    document.getElementById("errorPassword_config").style.color = "red";
    isValidpassword = false;
  } else {
    document.getElementById("errorPassword_config").innerHTML = "";
    document.getElementById("errorPassword").innerHTML = "";
    isValidpassword = true;
  }

  if (isValidpassword && isValidusername && isValidemail) {
    try {
      // Send the POST request
      const response = await fetch(`/user/register`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json; charset=UTF-8",
        },
        body: JSON.stringify({
          username: username.value,
          email: email.value,
          password: password.value,
          password_config: password_config.value,
        }),
      });
      
      if (response.ok) {
         window.location.href = "/";
      } else {
        const data = await response.json();
        document.getElementById("errorPassword").innerText = data.error;
        document.getElementById("errorPassword").style.color = "red";
      }

      // Handle JSON response
    } catch (error) {
      // Handle network errors
      console.error("Network error:", error);
      alert("Unable to submit the comment due to a network issue.");
    }
  }
}
