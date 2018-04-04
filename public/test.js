console.log("object")

var dragged;
// document.addEventListener("touchmove", function (e) {
//     e.preventDefault()
//     return false
// })

/* events fired on the draggable target */
document.addEventListener("drag", function (event) {
    console.log("drag")
}, false);

document.addEventListener("dragstart", function (event) {
    // store a ref. on the dragged elem
    dragged = event.target;
    console.log("drag start")
    // make it half transparent
    event.target.style.opacity = .5;
}, false);

document.addEventListener("dragend", function (event) {
    // reset the transparency
    console.log("end", event.target.innerText)
    console.log("end drop target", event.srcElement)
    event.target.style.opacity = "";
}, false);

/* events fired on the drop targets */
document.addEventListener("dragover", function (event) {
    // prevent default to allow drop
    // console.log(event.target)
    console.log("over")
    event.preventDefault();
}, false);

document.addEventListener("dragenter", function (event) {
    // highlight potential drop target when the draggable element enters it

    console.log("enter", event.srcElement.innerText)
    if (event.target.className == "dropzone") {
        event.target.style.background = "purple";
    }

}, false);

document.addEventListener("dragleave", function (event) {
    // reset background of potential drop target when the draggable element leaves it
    console.log("leave")
    if (event.target.className == "dropzone") {
        event.target.style.background = "";
    }

}, false);

document.addEventListener("drop", function (event) {
    // prevent default action (open as link for some elements)
    event.preventDefault();
    console.log("drop")
    // move dragged elem to the selected drop target
    if (event.target.className == "dropzone") {
        event.target.style.background = "";
        dragged.parentNode.removeChild(dragged);
        event.target.appendChild(dragged);
    }

}, false);