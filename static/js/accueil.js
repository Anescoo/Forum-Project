var textarea = document.getElementsByClassName("Comment-Zone") 

function HideTextArea() {
    textarea.toggle("hidden")
}

function TogglePopup(id) {
    const popup = document.getElementById(id)
    popup.toggleAttribute("hidden")
}