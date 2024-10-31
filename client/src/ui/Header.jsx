import { useEffect, useState } from "react";
import Logo from "../components/Logo";
import Navbar from "../components/Navbar";
import BookSearch from "../features/book/BookSearch";
import { CiLogin } from "react-icons/ci";
import { useHubPeak } from "../context/HubContext";

const Header = () => {
  const [isScrolled, setIsScrolled] = useState(false);
  const { modal, setModal } = useHubPeak();

  function handleClick() {
    setModal(!modal);
  }

  useEffect(() => {
    function handleScroll() {
      const scrollTop = window.scrollY;
      if (scrollTop > 36) {
        setIsScrolled(true);
      } else {
        setIsScrolled(false);
      }
    }

    window.addEventListener("scroll", handleScroll);
    return () => {
      window.removeEventListener("scroll", handleScroll);
    };
  }, []);

  return (
    <header
      className={`flex justify-around h-24 items-center sticky top-0 transition-colors z-10 shadow-md duration-200 ${
        isScrolled ? "bg-zinc-900 text-stone-200" : "bg-white"
      }`}
    >
      <Logo color={`${isScrolled && "text-stone-200"}`} />
      <Navbar />

      <div className="flex gap-3 items-center">
        <BookSearch />
        <CiLogin onClick={handleClick} className="text-2xl cursor-pointer" />
      </div>
    </header>
  );
};

export default Header;
