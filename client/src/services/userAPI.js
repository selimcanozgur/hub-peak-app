const API = "http://localhost:8080/signup";

export async function signup(newUser) {
  try {
    const res = await fetch(API, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(newUser),
    });

    if (!res.ok) {
      throw new Error(`Error: ${res.status} ${res.statusText}`);
    }

    const data = await res.json();
    return data;
  } catch (err) {
    console.error("Signup failed:", err);
    throw err;
  }
}
