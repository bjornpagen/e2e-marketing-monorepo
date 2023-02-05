const appendParams = (paramsArr, domains) => {
  // Get all links on the page
  const links = document.getElementsByTagName("a");

  // Loop through each link
  for (const link of links) {
    let href = link.href;

    // Check if the link is in the list of domains
    for (const domain of domains) {
      if (href.includes(domain)) {
        let params = new URLSearchParams(paramsArr);
        // Add in the existing params on the link
        const existingParams = new URLSearchParams(new URL(href).search) ?? [];
        for (const [key, value] of existingParams) {
          // Check if the param is already in the UTM object, if it is, override it with new value
          if (params.has(key)) {
            continue;
          } else {
            params.append(key, value);
          }
        }

        // Append the new parameters to the link
        let paramStr = "";
        for (const [key, value] of params) {
          paramStr += `${key}=${value}&`;
        }
        // if href already had parameters, cut them all out
        if (href.includes("?")) {
          href = href.slice(0, href.indexOf("?"));
        }
        if (paramStr) {
          href += `?${paramStr.slice(0, -1)}`;
        }
        link.setAttribute("href", href);
        break;
      }
    }
  }
};

export { appendUTMParams };