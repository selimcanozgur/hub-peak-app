import { Outlet } from "react-router-dom";
import DashboardNavbar from "../components/DashboardNavbar";

const DashBoardLayout = () => {
  return (
    <div>
      <DashboardNavbar />
      <main>
        <Outlet />
      </main>
    </div>
  );
};

export default DashBoardLayout;
