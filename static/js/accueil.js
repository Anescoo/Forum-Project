var textarea = document.getElementsByClassName("Comment-Zone") 
const btn = document.getElementById("btn")

function HideTextArea() {
    textarea.toggle("hidden")
}

function TogglePopup(id) {
    const popup = document.getElementById(id)
    popup.toggleAttribute("hidden")
    document.getElementById("popupForm").style.display="none";
    btn.style.display="flex"
}



function openForm() {
    document.getElementById("popupForm").style.display="block";
    btn.style.display="none"
}

function closeForm() {
    document.getElementById("popupForm").style.display="none";
    btn.style.display="flex"
}