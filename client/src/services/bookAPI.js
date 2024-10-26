const API = "http://localhost:8080/books";
// http://localhost:8080/books?title=nut

export async function getAllBook({ params }) {
  try {
    const { all, title } = params;
    const queryParams = new URLSearchParams();

    if (all) {
      queryParams.append("filter", all);
    }

    if (title) {
      queryParams.append("title", title);
    }

    const res = await fetch(`${API}?${queryParams.toString()}`);

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

export async function getOneBook({ params }) {
  try {
    const { id } = params;
    const res = await fetch(`${API}/${id}`);
    const data = await res.json();
    return data;
  } catch (err) {
    console.log(err);
  }
}
