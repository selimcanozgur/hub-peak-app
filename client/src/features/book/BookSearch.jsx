import { useState } from "react";
import { useNavigate } from "react-router-dom";

const BookSearch = () => {
  const [title, setTitle] = useState("");
  const navigate = useNavigate();

  const handleSubmit = (e) => {
    e.preventDefault();

    if (title.length > 2) {
      navigate(`/books/title/${title}`);
    }

    setTitle("");
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        className="border border-zinc-300 px-4 py-2 focus:outline-none rounded-sm w-54"
        type="text"
        placeholder="Kitap ara..."
      />
      <button
        type="submit"
        className="ml-2 px-4 py-2 bg-blue-500 text-white rounded-sm"
      >
        Ara
      </button>
    </form>
  );
};

export default BookSearch;
