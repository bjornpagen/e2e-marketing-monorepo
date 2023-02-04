// endpoint is the domain name of the endpoint, no https:// or trailing slash
const getEmailById = async (endpoint, id) => {
  const response = await fetch(`https://${endpoint}/lookup`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ id: id }),
  });

  if (!response.ok) {
    return null;
  }

  const data = await response.json();
  return data.email;
}

export { getEmailById };