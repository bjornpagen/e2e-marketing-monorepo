const appendUTMParams = domains => {
  // Get all the links on the page
  const links = document.getElementsByTagName("a");

  // Get the UTM parameters from the URL
  const search = window.location.search;
  const params = new URLSearchParams(search);

  // Array of allowed UTM parameters
  const allowedUTMs = ["utm_source", "utm_medium", "utm_campaign", "utm_term", "utm_content"];

  // Loop through each link
  for (const link of links) {
    let href = link.getAttribute("href");

    // Check if the link is in the list of domains
    for (const domain of domains) {
      if (href.includes(domain)) {
        // Append the UTM parameters to the link
        let utmParams = "";
        for (const [key, value] of params) {
          const newKey = allowedUTMs.includes(key) ? key : `utm_${key}`;
          if (allowedUTMs.includes(newKey)) {
            utmParams += `${newKey}=${value}&`;
          }
        }
        if (utmParams) {
          href += `?${utmParams.slice(0, -1)}`;
        }
        link.setAttribute("href", href);
        break;
      }
    }
  }
};

// example usage:
//
//  window.addEventListener("DOMContentLoaded", () => {
//    appendUTMParams(["example.com", "example.org"]);
//  });

export { appendUTMParams };