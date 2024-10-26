import { useLoaderData, useNavigate } from "react-router-dom";
import BookItem from "./BookItem";
import { useEffect, useState } from "react";

const Book = () => {
  const data = useLoaderData();
  const [filter, setFilter] = useState("all");
  const navigate = useNavigate();

  useEffect(() => {
    navigate(`/books/${filter}`);
  }, [filter, navigate]);

  return (
    <div>
      <div className="mt-20">
        <select
          onChange={(e) => setFilter(e.target.value)}
          className="absolute right-36 top-32 focus:outline-none bg-blue-100/80 shadow-lg px-4 py-2"
        >
          <option value="all"> Tüm Kitaplar</option>
          <option value="asc">En Düşük Fiyat</option>
          <option value="desc">En Yüksek Fiyat</option>
        </select>
        <ul className="flex flex-wrap justify-center">
          {data.map((book) => (
            <BookItem key={book.id} book={book} />
          ))}
        </ul>
      </div>
    </div>
  );
};

export default Book;
