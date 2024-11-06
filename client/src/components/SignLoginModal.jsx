import { useEffect } from "react";
import { useHubPeak } from "../context/HubContext";
import { FaTimes } from "react-icons/fa";
import { FcGoogle } from "react-icons/fc";
import { AiOutlineMail } from "react-icons/ai";
import Signup from "../features/user/Signup";
import Logo from "./Logo";
import Login from "../features/user/Login";

const SignLoginModal = () => {
  const { setModal, signup, setSignup, login, setLogin } = useHubPeak();

  useEffect(() => {
    const handleKeyDown = (event) => {
      if (event.key === "Escape") {
        setModal(false);
      }
    };

    document.addEventListener("keydown", handleKeyDown);

    return () => document.removeEventListener("keydown", handleKeyDown);
  }, [setModal]);

  const handleOutsideClick = (event) => {
    if (event.target.id === "modal-overlay") {
      setModal(false);
    }
  };

  return (
    <>
      <div
        id="modal-overlay"
        className="fixed inset-0 z-50 flex items-center justify-around bg-black bg-opacity-90 backdrop-blur-sm font-roboto"
        onClick={handleOutsideClick}
      >
        <div>
          <Logo color="text-white" size="text-4xl" />
        </div>
        <div
          className="w-[600px] h-[650px] bg-white shadow-xl"
          onClick={(e) => e.stopPropagation()}
        >
          <button
            onClick={() => setModal(false)}
            className="float-right mt-4 mr-4 border shadow-xl px-2 py-2 rounded-full"
          >
            <FaTimes className="text-red-500 text-md" />
          </button>
          <div className="text-center">
            {signup ? (
              <Signup />
            ) : login ? (
              <Login />
            ) : (
              <>
                <h1 className=" text-3xl mt-32 mb-24">Bize Katıl.</h1>
                <button
                  className="mb-8 border-2 rounded-full hover:bg-slate-100 px-6 py-3"
                  onClick={() => setSignup(!signup)}
                >
                  <p className="flex items-center gap-4">
                    <AiOutlineMail className="text-xl" /> E Posta ile kaydolun
                  </p>
                </button>
                <br />
                <button className="mb-24 border-2 rounded-full hover:bg-slate-100  px-6 py-3">
                  <p className="flex items-center gap-4">
                    <FcGoogle className="text-xl" /> Google ile kaydolun
                  </p>
                </button>
                <div>
                  Zaten bir hesabınız var mı?
                  <strong
                    className="text-blue-500"
                    onClick={() => setLogin(!login)}
                  >
                    Giriş yap
                  </strong>
                </div>
              </>
            )}
          </div>
        </div>
      </div>
    </>
  );
};

export default SignLoginModal;
