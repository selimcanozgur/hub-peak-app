import Header from "./Header";
import { Outlet } from "react-router-dom";
import SignLoginModal from "../components/SignLoginModal";
import { useHubPeak } from "../context/HubContext";
import Footer from "./Footer";

const AppLayout = () => {
  const { modal } = useHubPeak();
  return (
    <div>
      <Header />
      {modal && <SignLoginModal />}
      <main>
        <Outlet />
      </main>
      <Footer />
    </div>
  );
};

export default AppLayout;
