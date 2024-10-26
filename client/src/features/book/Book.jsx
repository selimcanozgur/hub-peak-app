import { useLoaderData } from "react-router-dom";
import BookItem from "./BookItem";

const Book = () => {
  const data = useLoaderData();

  return (
    <div className="mt-12">
      <ul className="flex flex-wrap justify-center">
        {data.map((book) => (
          <BookItem key={book.id} book={book} />
        ))}
      </ul>
    </div>
  );
};

export default Book;
