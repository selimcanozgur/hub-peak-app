import { FaBookOpen } from "react-icons/fa";
import PropTypes from "prop-types";

const Logo = ({ color = "text-black", size = "text-3xl" }) => {
  return (
    <div className="flex items-center gap-4">
      <FaBookOpen className="text-red-600 text-4xl" />
      <p className={`font-poppins ${color} font-medium ${size}`}>Hub Peak</p>
    </div>
  );
};

Logo.propTypes = {
  color: PropTypes.string,
  size: PropTypes.string,
};

export default Logo;
