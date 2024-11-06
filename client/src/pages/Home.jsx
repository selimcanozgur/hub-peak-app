import Slick from "../components/Slick";

const Home = () => {
  return (
    <div>
      <div className="w-full h-[605px] homeBanner">
        <img src="/img/book.png" />
      </div>
      <div className="mt-12">
        <h1 className="text-4xl">Favori Kitaplar</h1>
        <Slick />
      </div>
    </div>
  );
};

export default Home;
