var textArea = document.getElementById("modiftextarea")
var button = document.getElementById("buttonUpdate")

function AddHidden() {
    textArea.setAttribute("hidden", "")
    button.setAttribute("hidden", "")
}

function RemoveHidden() {
    textArea.removeAttribute("hidden")
    button.removeAttribute("hidden")
}

console.log("test")