import { NavLink } from "react-router-dom";

const Navbar = () => {
  return (
    <div>
      <ul className="flex gap-8 font-montserrat font-medium text-sm tracking-wide">
        <li>
          <NavLink to="/">Ana Sayfa</NavLink>
        </li>
        <li>
          <NavLink to="/books/all">Kitaplar</NavLink>
        </li>
        <li>
          <NavLink to="/">Hakkımızda</NavLink>
        </li>
        <li>
          <NavLink to="/">İletişim</NavLink>
        </li>
      </ul>
    </div>
  );
};

export default Navbar;
