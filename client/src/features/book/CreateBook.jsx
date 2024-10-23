import { useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { addBook } from "./bookSlice";
import InputComp from "../../components/InputComp";

const BookForm = () => {
  const dispatch = useDispatch();
  const { loading, error } = useSelector((state) => state.book);

  const [title, setTitle] = useState("");
  const [author, setAuthor] = useState("");
  const [publishing_house, setPublishingHouse] = useState("");
  const [publishing_year, setPublishingYear] = useState("");
  const [price, setPrice] = useState("");
  const [pages, setPages] = useState("");
  const [summary, setSummary] = useState("");
  const [lang, setLang] = useState("");
  const [cover_type, setCoverType] = useState("");
  const [image_url, setImageUrl] = useState("");
  const [book_id, setBookId] = useState("");
  const [stock, setStock] = useState("");

  const handleSubmit = (e) => {
    e.preventDefault();

    const bookData = {
      title,
      author,
      publishing_house,
      summary,
      lang,
      cover_type,
      image_url,
      pages: parseInt(pages),
      book_id: parseInt(book_id),
      stock: parseInt(stock),
      publishing_year: parseInt(publishing_year),
      price: parseFloat(price),
    };

    dispatch(addBook(bookData)); // Dispatch the addBook thunk
  };

  return (
    <form onSubmit={handleSubmit}>
      <div className="grid grid-cols-3 justify-items-center">
        <InputComp
          placeholder="Kitap Adı"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
        />

        <InputComp
          placeholder="Yazar Adı"
          value={author}
          onChange={(e) => setAuthor(e.target.value)}
        />

        <InputComp
          placeholder="Yayınevi"
          value={publishing_house}
          onChange={(e) => setPublishingHouse(e.target.value)}
          required
        />

        <InputComp
          placeholder="Yayınlanma Tarihi"
          value={publishing_year}
          onChange={(e) => setPublishingYear(e.target.value)}
          required
        />

        <InputComp
          placeholder="Fiyatı"
          value={price}
          onChange={(e) => setPrice(e.target.value)}
        />

        <InputComp
          placeholder="Sayfa Sayısı"
          value={pages}
          onChange={(e) => setPages(e.target.value)}
        />

        <InputComp
          placeholder="Kitap Dili"
          value={lang}
          onChange={(e) => setLang(e.target.value)}
        />

        <InputComp
          placeholder="Kitap Türü"
          value={cover_type}
          onChange={(e) => setCoverType(e.target.value)}
        />

        <InputComp
          placeholder="Kitap Resmi"
          value={image_url}
          onChange={(e) => setImageUrl(e.target.value)}
        />

        <InputComp
          placeholder="Kitap ID"
          value={book_id}
          onChange={(e) => setBookId(e.target.value)}
        />

        <InputComp
          placeholder="Stok"
          value={stock}
          onChange={(e) => setStock(e.target.value)}
        />
      </div>
      <div className="flex justify-center">
        <textarea
          className="border w-[600px] h-48"
          placeholder="Kitap Özeti"
          value={summary}
          onChange={(e) => setSummary(e.target.value)}
        />
      </div>

      <button type="submit" disabled={loading}>
        {loading ? "Yükleniyor..." : "Kitap Ekle"}
      </button>

      {error && <p style={{ color: "red" }}>{error}</p>}
    </form>
  );
};

export default BookForm;
