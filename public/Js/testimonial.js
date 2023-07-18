// class Testimonial {
//   #quote = "";
//   #image = "";

//   constructor(quote, image) {
//     this.#quote = quote;
//     this.#image = image;
//   }

//   get quote() {
//     return this.#quote;
//   }

//   get image() {
//     return this.#image;
//   }

//   // This is an abstract method that subclasses will implement
//   get author() {
//     throw new Error("getAuthor() method must be implemented");
//   }

//   // This is a polymorphic method that can take any subclasses of Testimonial
//   get testimonialHTML() {
//     return `<div class="testimonial">
//                     <img src="${this.image}" alt="" class="profile-testimonial" />
//                     <p class="quote">
//                         ${this.quote}
//                     </p>
//                     <p class="author">- ${this.author}</p>
//                 </div>
//           `;
//   }
// }

// class AuthorTestimonials extends Testimonial {
//   #author = "";

//   constructor(author, quote, image) {
//     super(quote, image);
//     this.#author = author;
//   }

//   get author() {
//     return this.#author;
//   }
// }

// const testimonial1 = new AuthorTestimonials(
//   "Batuhan",
//   "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam egetlectus eu libero venenatis gravida eget vel odio.",
//   "https://images.unsplash.com/photo-1686041673559-a763bd15d2f9?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=687&q=80"
// );
// const testimonial2 = new AuthorTestimonials(
//   "Lucas",
//   "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam egetlectus eu libero venenatis gravida eget vel odio.",
//   "https://images.unsplash.com/photo-1685627298101-01b11fe1b17c?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=687&q=80"
// );

// const testimonial3 = new AuthorTestimonials (
//     "Chuko",
//     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam egetlectus eu libero venenatis gravida eget vel odio.",
//     "https://images.unsplash.com/photo-1686781483909-a5fb625ca043?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=765&q=80"
// )

// const testimonial4 = new AuthorTestimonials(
//   "Jayson",
//   "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam egetlectus eu libero venenatis gravida eget vel odio.",
//   "https://images.unsplash.com/photo-1676178294406-5b907fb04b11?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=686&q=80"
// );

// const testimonial5 = new AuthorTestimonials(
//   "Sergey",
//   "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam egetlectus eu libero venenatis gravida eget vel odio.",
//   "https://images.unsplash.com/photo-1633190323610-49a306581216?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=687&q=80"
// );

// let testimonialData = [testimonial1, testimonial2, testimonial3, testimonial4, testimonial5];
// let testimonialHTML = "";

// for (let i = 0; i < testimonialData.length; i++) {
//   testimonialHTML += testimonialData[i].testimonialHTML;
// }

// document.getElementById("testimonials").innerHTML = testimonialHTML;

// const testimonialData = [
//   {
//     author: "Batuhan",
//     quote:
//       "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam egetlectus eu libero venenatis gravida eget vel odio.",
//     rating: 5,
//     image:
//       "https://images.unsplash.com/photo-1686041673559-a763bd15d2f9?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=687&q=80",
//   },
//   {
//     author: "Lucas",
//     quote:
//       "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam egetlectus eu libero venenatis gravida eget vel odio.",
//     rating: 4,
//     image:
//       "https://images.unsplash.com/photo-1685627298101-01b11fe1b17c?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=687&q=80",
//   },
//   {
//     author: "Chuko",
//     quote:
//       "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam egetlectus eu libero venenatis gravida eget vel odio.",
//     rating: 3,
//     image:
//       "https://images.unsplash.com/photo-1686781483909-a5fb625ca043?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=765&q=80",
//   },
//   {
//     author: "Sergey",
//     quote:
//       "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam egetlectus eu libero venenatis gravida eget vel odio.",
//     rating: 2,
//     image:
//       "https://images.unsplash.com/photo-1633190323610-49a306581216?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=687&q=80",
//   },
// ];

// function getAllTestimonials() {
//   let testimonialHTML = ''

//   testimonialData.forEach((card) => {
//     testimonialHTML += `
//         <div class="testimonial">
//          <img src="${card.image}" alt="" class="profile-testimonial" />
//           <p class="quote">
//              ${card.quote}
//           </p>
//           <p class="author">- ${card.author}</p>
//           <p class="author">- ${card.rating}  <i class="fa-solid fa-star"></i></p>
//         </div>
//     `;
//   })

//   document.getElementById('testimonials').innerHTML = testimonialHTML
// }

// getAllTestimonials()

// function getFilteredTestimonials(rating) {
//   let filteredTestimonialHTML = ''

//   const filteredData = testimonialData.filter((card) => {
//     return card.rating === rating;
//   })

//   filteredData.forEach((card) => {
//     filteredTestimonialHTML += `
//     <div class="testimonial">
//          <img src="${card.image}" alt="" class="profile-testimonial" />
//           <p class="quote">
//              ${card.quote}
//           </p>
//           <p class="author">- ${card.author}</p>
//           <p class="author"> ${card.rating}  <i class="fa-solid fa-star"></i></p>
//         </div>
//     `;
//   })
//   document.getElementById('testimonials').innerHTML = filteredTestimonialHTML
// }