import { useEffect, useState } from "react";
import Logo from "../components/Logo";
import Navbar from "../components/Navbar";
import BookSearch from "../features/book/BookSearch";
import { CiLogin } from "react-icons/ci";

const Header = () => {
  const [isScrolled, setIsScrolled] = useState(false);

  useEffect(() => {
    const handleScroll = () => {
      const scrollTop = window.scrollY;
      if (scrollTop > 36) {
        // 1 tab yüksekliğine yakın bir değer
        setIsScrolled(true);
      } else {
        setIsScrolled(false);
      }
    };

    window.addEventListener("scroll", handleScroll);
    return () => {
      window.removeEventListener("scroll", handleScroll);
    };
  }, []);

  return (
    <header
      className={`flex justify-around h-24 items-center sticky top-0 transition-colors z-50 shadow-md duration-200 ${
        isScrolled ? "bg-zinc-900 text-stone-200" : "bg-white"
      }`}
    >
      <Logo />
      <Navbar />

      <div className="flex gap-3 items-center">
        <BookSearch />
        <CiLogin className="text-2xl" />
      </div>
    </header>
  );
};

export default Header;
