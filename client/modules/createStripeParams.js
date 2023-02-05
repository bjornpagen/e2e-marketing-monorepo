// converts a user object into an array of params to be appended to stripe checkout link
// this prefills info such as email
const createStripeParams = (user) => {
  const params = [];
  if (user.email) {
    params.push(["prefilled_email", user.email]);
  }
  return params;
}

export { createStripeParams };