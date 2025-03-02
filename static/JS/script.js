dragElement(document.getElementById("window"));

function dragElement(elmnt) {
  var pos1 = 0, pos2 = 0, pos3 = 0, pos4 = 0;
  if (document.getElementById(elmnt.id + "header")) {
    // if present, the header is where you move the DIV from:
    document.getElementById(elmnt.id + "header").onmousedown = dragMouseDown;
  } else {
    // otherwise, move the DIV from anywhere inside the DIV:
    elmnt.onmousedown = dragMouseDown;
  }

  function dragMouseDown(e) {
    e = e || window.event;
    e.preventDefault();
    // get the mouse cursor position at startup:
    pos3 = e.clientX;
    pos4 = e.clientY;
    document.onmouseup = closeDragElement;
    // call a function whenever the cursor moves:
    document.onmousemove = elementDrag;
  }

  function elementDrag(e) {
    e = e || window.event;
    e.preventDefault();
    // calculate the new cursor position:
    pos1 = pos3 - e.clientX;
    pos2 = pos4 - e.clientY;
    pos3 = e.clientX;
    pos4 = e.clientY;
    // set the element's new position:
    elmnt.style.top = (elmnt.offsetTop - pos2) + "px";
    elmnt.style.left = (elmnt.offsetLeft - pos1) + "px";
  }

  function closeDragElement() {
    // stop moving when mouse button is released:
    document.onmouseup = null;
    document.onmousemove = null;
  }
}

window.addEventListener("load", () => {
  clock();
  function clock() {
    const today = new Date();


    const hours = today.getHours();
    const minutes = today.getMinutes();

    const hour = hours < 10 ? "0" + hours : hours;
    const minute = minutes < 10 ? "0" + minutes : minutes;
   
    const hourTime = hour > 12 ? hour - 12 : hour;

    const ampm = hour < 12 ? "AM" : "PM";
   
    const time = hourTime + ":" + minute + "  " + ampm;

    const dateTime = time;

    document.getElementById("date-time").innerHTML = dateTime;
    setTimeout(clock, 1000);
  }
});

function addResultElement(i_name, i_filename, i_hash, passed){
  
  results_box = document.getElementById('results-box');

  row = document.createElement("div");
  row.classList.add("row");

  modname = document.createElement("p");
  modname.classList.add("modname");
  modname.innerHTML = i_name;
  row.appendChild(modname);

  filename = document.createElement("p");
  filename.classList.add("modname");
  filename.innerHTML = i_filename;
  row.appendChild(filename);

  hash = document.createElement("p");
  hash.classList.add("hash");
  hash.innerHTML = i_hash;
  row.appendChild(hash);

  icon = document.createElement("img");
  icon.classList.add("icon")
  if (passed) {
    icon.src = "/IMG/check.png";
  } else {
    icon.src = "/IMG/no.png";
  }
  row.appendChild(icon) 

  results_box.appendChild(row);
  
}


const input = document.getElementById('fileinput');


const upload = (file) => {
  const formData  = new FormData();
  
  formData.append("file", file);

  fetch(window.location + 'upload', { 
    method: 'POST',
    body: formData,
  }).then(
    response => {
      response.json().then(
        data => {
          for (i in data) {
            item = data[i]
            addResultElement(item["Name"],item["Filename"],item["Hash"],item["Found"]);
          }
        } 
      )
    }
  ).catch(
    error => console.log(error)
  );
};

// Event handler executed when a file is selected
const onSelectFile = () => upload(input.files[0]);

// Add a listener on your input
// It will be triggered when a file will be selected
input.addEventListener('change', onSelectFile, false);
