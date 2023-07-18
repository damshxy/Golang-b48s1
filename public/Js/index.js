let responsiveOpen = false;

function responsiveNavbar() {
  let responsiveNav = document.getElementById("btn-itemMenu");

  if (!responsiveOpen) {
    console.log(responsiveOpen);
    responsiveNav.style.display = "flex";
    responsiveOpen = true;
  } else {
    console.log(responsiveOpen);
    responsiveNav.style.display = "none";
    responsiveOpen = false;
  }
}
