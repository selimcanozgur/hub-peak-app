import Header from "./Header";
import { Outlet } from "react-router-dom";
import SignLoginModal from "../components/SignLoginModal";
import { useHubPeak } from "../context/HubContext";

const AppLayout = () => {
  const { modal } = useHubPeak();
  return (
    <div>
      <Header />
      {modal && <SignLoginModal />}
      <main>
        <Outlet />
      </main>
    </div>
  );
};

export default AppLayout;
