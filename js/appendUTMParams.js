function appendUTMParams(domains) {
  // Get all the links on the page
  var links = document.getElementsByTagName("a");

  // Get the UTM parameters from the URL
  var search = window.location.search;
  var params = new URLSearchParams(search);

  // Array of allowed UTM parameters
  var allowedUTMs = ["utm_source", "utm_medium", "utm_campaign", "utm_term", "utm_content"];

  // Loop through each link
  for (var i = 0; i < links.length; i++) {
    var link = links[i];
    var href = link.getAttribute("href");

    // Check if the link is in the list of domains
    for (var j = 0; j < domains.length; j++) {
      if (href.indexOf(domains[j]) !== -1) {
        // Append the UTM parameters to the link
        var utmParams = "";
        for (var [key, value] of params.entries()) {
          var newKey = allowedUTMs.includes(key) ? key : 'utm_' + key;
          if (allowedUTMs.includes(newKey)) {
            utmParams += newKey + "=" + value + "&";
          }
        }
        if (utmParams) {
          href += "?" + utmParams.slice(0, -1);
        }
        link.setAttribute("href", href);
        break;
      }
    }
  }
}

