function toggleNavbar() {
  var x = document.querySelector('header nav ul');
  if (x.style.display === "block") {
    x.style.display = "none";
  } else {
    x.style.display = "block";
  }
}

// Sticky Navbar Functionality
window.onscroll = function() {makeSticky()};
var navbar = document.querySelector('header nav');
var sticky = navbar.offsetTop;
function makeSticky() {
  if (window.pageYOffset >= sticky) {
    navbar.classList.add("sticky");
  } else {
    navbar.classList.remove("sticky");
  }
}
