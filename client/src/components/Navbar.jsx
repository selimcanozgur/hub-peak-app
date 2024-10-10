import { NavLink } from "react-router-dom";

const Navbar = () => {
  return (
    <div>
      <ul className="flex gap-8 font-montserrat font-medium text-sm tracking-wide">
        <li>
          <NavLink to="/">Home</NavLink>
        </li>
        <li>
          <NavLink to="/">Books</NavLink>
        </li>
        <li>
          <NavLink to="/">About</NavLink>
        </li>
        <li>
          <NavLink to="/">Context</NavLink>
        </li>
      </ul>
    </div>
  );
};

export default Navbar;
