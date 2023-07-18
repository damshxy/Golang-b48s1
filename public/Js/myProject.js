function emptyFormAlert() {
  let projectName = document.getElementById("projectName").value;
  let startDate = document.getElementById("input-start-date").value;
  let endDate = document.getElementById("input-end-date").value;
  let description = document.getElementById("input-description").value;
  let image = document.getElementById("input-imgae").value;

  if (projectName = " ") {
    return alert("Please input your project name");
  } else if (startDate = " ") {
    return alert("When did you start this project ?");
  } else if (endDate) {
    return alert("When did you end this project ?");
  } else if (description = " ") {
    return alert("Please describe this project");
  } else if (image = " ") {
    return alert("Please input your image project");
  }
}

let projectData = [];

function postProject(event) {
  event.preventDefault();

  let projectName = document.getElementById("projectName").value;
  let startDate = document.getElementById("input-start-date").value;
  let endDate = document.getElementById("input-end-date").value;
  let description = document.getElementById("input-description").value;
  let image = document.getElementById("input-image").files;

  const nodeIcon = '<i class="fa-brands fa-node"></i>';
  const reactIcon = '<i class="fa-brands fa-react"></i>';
  const golangIcon = '<i class="fa-brands fa-golang"></i>';
  const jsIcon = '<i class="fa-brands fa-square-js"></i>';

  let nodeIconChecked = document.getElementById("nodeCheck").checked
    ? nodeIcon
    : "";
  let reactIconChecked = document.getElementById("reactCheck").checked
    ? reactIcon
    : "";
  let golangIconChecked = document.getElementById("goCheck").checked
    ? golangIcon
    : "";
  let jsIconChecked = document.getElementById("jsCheck").checked ? jsIcon : "";

  image = URL.createObjectURL(image[0]);
  console.log(image);

  function distanceTime() {
    let startDate = new Date(document.getElementById("input-start-date").value);
    let endDate = new Date(document.getElementById("input-end-date").value);
    let dist = new Date(endDate) - new Date(startDate);
    let days = Math.floor(dist / (1000 * 3600 * 24));
    let weeks = Math.floor(dist / (1000 * 3600 * 24 * 7));
    let months = Math.floor(dist / (1000 * 3600 * 24 * 30));
    let years = Math.floor(dist / (1000 * 3600 * 24 * 30 * 12));

    if (years == 1) {
      return `${years} year`;
    } else if (years > 0) {
      return `${years} years`;
    } else if (months == 1) {
      return `${months} month`;
    } else if (months > 0) {
      return `${months} months`;
    } else if (weeks == 1) {
      return `${weeks} week`;
    } else if (weeks > 0) {
      return `${weeks} weeks`;
    } else if (days == 1) {
      return `${days} day`;
    } else if (days > 0) {
      return `${days} days`;
    }
  }

  let getDistanceTime = distanceTime();

  let projectPreviewCard = {
    projectName,
    getDistanceTime,
    description,
    nodeIconChecked,
    reactIconChecked,
    golangIconChecked,
    jsIconChecked,
    image,
  };

  projectData.push(projectPreviewCard);
  console.log(projectData);

  pushProject();
}

function pushProject() {
  document.getElementById("card-project").innerHTML = "";

  for (let i = 0; i < projectData.length; i++) {
    document.getElementById("card-project").innerHTML += `
    <div id="card-project" class="my-project" style="padding-top: 50px;">
    <div class="project-flex">
    <div class="project-flex-box">
    <img src="${projectData[i].image}">
    <a href="detailProject.html" target="_blank">
    <h3 class="project-name">${projectData[i].projectName}</h3>
    </a>
    <h5 class="duration">Duration: ${projectData[i].getDistanceTime}</h5>
    <p class="descriotion">${projectData[i].description}</p>
    <p id="timeAgo" style="float: right; margin: 10px;"></p>
    <p class="project-icon">
            ${projectData[i].nodeIconChecked}
            ${projectData[i].reactIconChecked}
            ${projectData[i].golangIconChecked}
            ${projectData[i].jsIconChecked}
            </p>
            <div class="project-button">
            <button type="button">edit</button>
            <button type="button">delete</button>
            </div>
            </div>
            </div>
            </div>
            `;
  }
}