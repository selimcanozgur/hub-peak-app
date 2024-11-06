import { NavLink } from "react-router-dom";
import styled from "styled-components";

const Li = styled.li`
  background-color: #7d7d7d;
  border-radius: 30px;
  padding: 8px 36px;
  color: white;
  font-weight: 500;
  margin-bottom: 8px;
`;

const DashboardNavbar = () => {
  return (
    <ul className="h-screen border-r bg-slate-50 w-80 px-4 py-8">
      <NavLink to="profile">
        <Li> Profil </Li>
      </NavLink>
      <NavLink to="users">
        <Li> Kullanıcılar </Li>
      </NavLink>
      <NavLink to="users">
        <Li> Kitaplar </Li>
      </NavLink>
    </ul>
  );
};

export default DashboardNavbar;
