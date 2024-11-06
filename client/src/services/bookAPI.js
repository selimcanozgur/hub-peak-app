const API = "http://localhost:8080/books";

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
      const errorData = await res.json();
      throw new Error(errorData.message || "Sunucudan bir hata oluştu.");
    }

    const data = await res.json();
    return data;
  } catch (err) {
    console.error("Hata:", err.message);
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
