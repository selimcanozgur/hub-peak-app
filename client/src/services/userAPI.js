const API = "http://localhost:8080/users/admin";

export async function getAllUser() {
  try {
    const res = await fetch(API, {
      credentials: "include",
    });

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
