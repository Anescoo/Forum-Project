// let contener = document.getElementById("pagination")

// let test = document.getElementsByTagName("form")
// let btnNbr = Math.ceil((test.length -1)/ 10)
var BtnForComments = document.querySelectorAll('.CommsBtn')

// ButtonNumber(btnNbr)
ComsBtn()

// function ButtonNumber(param) {
//     contener.innerHTML = ""
//     for (let index = 1; index <= param; index++) {
//         let page = document.createElement("button")
//         page.setAttribute("onclick", "ChangePage(" + index + ")")
//         page.id = "page" + index
//         page.innerText = "Page n° " + index
//         contener.appendChild(page)
//     }
// }

// Get the popup
var popup = document.getElementById("popupBox");

// Get the button that opens the popup
// var btn = document.getElementById("OpenPopup0");

// Get the <span> element that closes the popup
var span = document.getElementsByClassName("close")[0];

// When the user clicks on the button, open the popup
// btn.onclick = function() {
//   popup.style.display = "block";
// }


// When the user clicks on <span> (x), close the popup
span.onclick = function() {
  popup.style.display = "none";
}

// When the user clicks anywhere outside of the popup, close it
window.onclick = function(event) {
  if (event.target == popup) {
    popup.style.display = "none";
  }
}

window.addEventListener('keydown', (e) => {
  if (e.key == "Escape") {
    popup.style.display = "none";
  }
})

function ComsBtn(){
  for (let i = 0; i < BtnForComments.length; i++){
    BtnForComments[i].id = "OpenPopup" + i
    BtnForComments[i].setAttribute("onclick", "OpenPopup()")
  }
}

function OpenPopup(){
  popup.style.display = "block";
}

btn.onclick = function() {
  popup.style.display = "block";
}

// When the user clicks on <span> (x), close the popup
span.onclick = function() {
  popup.style.display = "none";
}