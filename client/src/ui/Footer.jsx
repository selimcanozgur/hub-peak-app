import Logo from "../components/Logo";

import { MdEmail } from "react-icons/md";
import { FaPhoneAlt } from "react-icons/fa";
import { FaHome } from "react-icons/fa";
import { FaGithub } from "react-icons/fa";
import { FaLinkedin } from "react-icons/fa";
import { FaInstagram } from "react-icons/fa";

import styled from "styled-components";

const P = styled.p`
  display: flex;
  align-items: center;
  gap: 10px;
`;

const Footer = () => {
  return (
    <footer className="bg-zinc-900 py-4">
      <div className="flex justify-around text-white/80 h-36 items-center">
        <div className="w-52">
          <P>
            <FaHome /> Eskişehir/Türkiye
          </P>
          <P>
            <FaPhoneAlt /> 554-655 5661
          </P>
          <P>
            <MdEmail /> selimozgur26@gmail.com
          </P>
        </div>
        <div>
          <Logo color="white" />
        </div>
        <div className="w-48">
          <P>
            <FaGithub /> GitHub
          </P>
          <P>
            <FaLinkedin /> LinkedIn
          </P>
          <P>
            <FaInstagram /> Instagram
          </P>
        </div>
      </div>
      <p className="text-center text-white/60 text-sm">
        ©2024 Selimcan Özgür. Tüm hakları saklıdır. İzinsiz kullanılamaz.
      </p>
    </footer>
  );
};

export default Footer;
