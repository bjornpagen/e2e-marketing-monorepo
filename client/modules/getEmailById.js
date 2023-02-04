const ENDPOINT_DOMAIN = "my-server.com";

const getEmailById = async id => {
  const response = await fetch(`https://${ENDPOINT_DOMAIN}/idlookup`, {
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