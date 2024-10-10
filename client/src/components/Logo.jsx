import { FaBookOpen } from "react-icons/fa";

const Logo = () => {
  return (
    <div className="flex items-center gap-4">
      <FaBookOpen className="text-red-600 text-4xl" />
      <p className="font-poppins uppercase text-3xl ">Hub Peak</p>
    </div>
  );
};

export default Logo;
