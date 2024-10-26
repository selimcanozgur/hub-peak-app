const API = "http://localhost:8080/books";

export async function getAllBook() {
  try {
    const res = await fetch(API);
    if (!res.ok) {
      throw new Error(`HTTP error! Status: ${res.status}`);
    }
    const data = await res.json();
    return data;
  } catch (err) {
    console.log("API Hatası:", err.message);
    return { error: err.message };
  }
}
