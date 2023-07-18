function contactSubmit(event) {
  event.preventDefault();

  let name = document.getElementById("name").value;
  let email = document.getElementById("email").value;
  let phone = document.getElementById("phoneNumber").value;
  let subject = document.getElementById("subject").value;
  let message = document.getElementById("inputTextArea").value;

  if (name == "") {
    return alert("Nama Harus Diisi!!!");
  } else if (email == "") {
    return alert("Email Harus Diisi!!!");
  } else if (phone == "") {
    return alert("PhoneNumber Harus Diisi!!!");
  } else if (subject == "") {
    return alert("Subject Harus Diisi!!!");
  } else if (message == "") {
    return alert("Message Harus Diisi!!!");
  }

  let emailReceiver = "nczeesh@gmail.com";
  let a = document.createElement("a");
  a.href = `mailto:${emailReceiver}?subject=${subject}&body=Halo, nama saya ${name}, ${message}. Silahkan kontak saya di nomor ${phone}, terimakasih.`;
  a.click();

  let objectData = {
    name,
    email,
    phone,
    subject,
    message,
  };

  console.log(objectData);
}
