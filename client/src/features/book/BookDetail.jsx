import { useLoaderData } from "react-router-dom";

function BookDetailPage() {
  const book = useLoaderData();
  const {
    title,
    author,
    lang,
    pages,
    book_cover,
    publishing_house,
    publishing_year,
    price,
    stock,
    img_path,
  } = book;

  if (book.error) {
    return <div>{book.error}</div>;
  }

  return (
    <div className="flex justify-center gap-8 mt-12">
      <img className="w-[250px] h-auto" src={img_path} alt={`photo${title}`} />
      <div>
        <h1 className="text-3xl font-semibold mb-2">{title}</h1>
        <div className="flex gap-4 text-xs mb-2">
          <p>Yazar: {author}</p>
          <p>Yayınevi: {publishing_house}</p>
        </div>
        <p className="text-2xl font-semibold">{price} TL</p>
        <div className="flex justify-center items-center gap-2 mt-4 border h-16 w-32 border-red-400">
          <img className="h-12" src={img_path} />
          <div className="font-medium text-sm">
            <p>{book_cover} Kapak</p>
            <p>{price} TL</p>
          </div>
        </div>
        <p className="mt-6 font-semibold"> Öne çıkan bilgiler</p>
        <div className="flex gap-24 mt-2">
          <div>
            <p>
              <strong> Dil: </strong> {lang}
            </p>
            <p>
              <strong> Sayfa Sayısı: </strong> {pages}
            </p>
          </div>
          <div>
            <p>
              <strong> Yayınlanma Tarihi: </strong> {publishing_year}
            </p>
            <p>
              <strong> Stok: </strong> {stock}
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}

export default BookDetailPage;
