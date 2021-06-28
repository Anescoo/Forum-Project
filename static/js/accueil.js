let contener = document.getElementById("pagination")

let test = document.getElementsByTagName("form")
let btnNbr = Math.ceil((test.length -1)/ 10)
console.log(btnNbr)
ButtonNumber(btnNbr)

function ButtonNumber(param) {
    contener.innerHTML = ""
    for (let index = 1; index <= param; index++) {
        let page = document.createElement("button")
        page.setAttribute("onclick", "ChangePage(" + index + ")")
        page.id = "page" + index
        page.innerText = "Page n° " + index
        contener.appendChild(page)
    }
}

function ChangePage(){
    
}

function openForm() {
    document.getElementById("Posts").style.display="block";
}

function closeForm() {
    document.getElementById("Posts").style.display="none";
}
