import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

const BookSearch = () => {
  const [title, setTitle] = useState("");
  const navigate = useNavigate();

  useEffect(() => {
    navigate(`/books/title/${title}`);
  }, [title, navigate]);
  return (
    <input
      onChange={(e) => setTitle(e.target.value)}
      className="border border-zinc-300 px-4 py-2 focus:outline-none rounded-sm w-54"
      type="text"
      placeholder="Kitap ara..."
    />
  );
};

export default BookSearch;
