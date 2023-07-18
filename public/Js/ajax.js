const promise = new Promise((resolve, reject) => {
  const xhr = new XMLHttpRequest();

  xhr.open("GET", "https://api.npoint.io/dd3924f19142075d57dd", true);
  xhr.onload = () => {
    if (xhr.status === 200) {
      resolve(JSON.parse(xhr.response));
    } else {
      reject("Error loading data.");
    }
  };
  xhr.onerror = () => {
    reject("Network error.");
  };
  xhr.send();
});

async function getAllTestimonials() {
  const response = await promise;

  let testimonials = "";
  response.forEach((item) => {
    testimonials += `
        <div class="testimonial">
         <img src="${item.image}" alt="" class="profile-testimonial" />
          <p class="quote">
             ${item.quote}
          </p>
          <p class="author">- ${item.author}</p>
          <p class="author"> ${item.rating}  <i class="fa-solid fa-star"></i></p>
        </div>
        `;
  });
  document.getElementById("testimonials").innerHTML = testimonials;
}
getAllTestimonials()

async function getFilteredTestimonials(rating) {
  const response = await promise;

  const testimonialsFiltered = response.filter((item) => {
    return item.rating === rating;
  });

  let testimonialsHTML = "";

  if (testimonialsFiltered.length === 0) {
    testimonialsHTML = "<h1>Data Not Found</h1>";
  } else {
    testimonialsFiltered.forEach((item) => {
      testimonialsHTML += `
            <div class="testimonial">
         <img src="${item.image}" alt="" class="profile-testimonial" />
          <p class="quote">
             ${item.quote}
          </p>
          <p class="author">- ${item.author}</p>
          <p class="author"> ${item.rating}  <i class="fa-solid fa-star"></i></p>
        </div>
            `;
    });
  }
  document.getElementById("testimonials").innerHTML = testimonialsHTML;
}