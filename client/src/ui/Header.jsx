import { useEffect, useState } from "react";
import Logo from "../components/Logo";
import Navbar from "../components/Navbar";
import BookSearch from "../features/book/BookSearch";
import { CiLogin } from "react-icons/ci";
import { useHubPeak } from "../context/HubContext";
import { userLoaded, userError } from "../features/user/userSlice";
import { useDispatch, useSelector } from "react-redux";
import { CiUser } from "react-icons/ci";
import { Link } from "react-router-dom";

const Header = () => {
  const { status } = useSelector((state) => state.user);
  const dispatch = useDispatch();

  useEffect(() => {
    async function getUser() {
      try {
        const res = await fetch("http://localhost:8080/profile", {
          credentials: "include",
        });
        const data = await res.json();
        if (res.status === 401) {
          dispatch(userError("Kullanıcı giriş yapmadı"));
          return;
        }
        dispatch(userLoaded(data));
      } catch (err) {
        console.log(err);
      }
    }
    getUser();
  }, [dispatch]);

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
        {status ? (
          <Link to="/dashboard/profile">
            <CiUser />
          </Link>
        ) : (
          <CiLogin onClick={handleClick} className="text-2xl cursor-pointer" />
        )}
      </div>
    </header>
  );
};

export default Header;
