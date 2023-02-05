const getUTMParams = () => {
  // Get the UTM parameters from the URL
  const search = window.location.search;
  const params = new URLSearchParams(search);

  // Array of allowed UTM parameters
  const allowedUTMs = ["utm_source", "utm_medium", "utm_campaign", "utm_term", "utm_content"];

  // Append the UTM parameters to the link
  let utmParams = [];
  for (const [key, value] of params) {
    const newKey = allowedUTMs.includes(key) ? key : `utm_${key}`;
    if (allowedUTMs.includes(newKey)) {
      utmParams.push([newKey, value]);
    }
  }

  return utmParams;
}

export { getUTMParams };