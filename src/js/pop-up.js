

var btnPopup = document.getElementById('btnPopup');
var overlay = document.getElementById('overlay');
var btnclose = document.getElementByID('btnclose');

btnPopup.addEventListener('click',openModal);

btnclose.addEventListener('click',closePopup);

function openModal() {
    overlay.style.display='none';
}

function closePopUP(){
    overlay.style.display = 'none';
}

function test(){
    alert("Bonjour")
}