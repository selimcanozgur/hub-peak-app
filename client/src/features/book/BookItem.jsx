import PropTypes from "prop-types";

const BookItem = ({ book }) => {
  const { title, author, publishing_house, price, img_path } = book;
  return (
    <li className="w-72 h-96 justify-center shadow-lg bg-zinc-200/45 hover:bg-zinc-200/75 rounded-md mx-4 my-4 flex flex-col items-center cursor-pointer">
      <img className="w-20 mb-6 hover:w-24 duration-300" src={img_path} />
      <h1 className="text-2xl font-semibold mb-2">{title}</h1>
      <h2 className="text-lg font-medium text-gray-700 mb-2">{author}</h2>
      <p className="font-bold text-xl text-green-600 mb-4">{price} TL</p>
      <p className="text-sm text-gray-600 italic">{publishing_house}</p>
    </li>
  );
};

BookItem.propTypes = {
  book: PropTypes.object.isRequired,
};

export default BookItem;
