// scans params for utm_content (which is the user id)
// then calls f to get the user info
const getStripeUser = async (params, f) => {
  const utmParams = params.filter((param) => param[0] === "utm_content");
  if (utmParams.length === 0) {
    return null;
  }
  const user = await f(utmParams[0][1]);
  if (!user) {
    return null;
  }
  return user;
}

export { getStripeUser };