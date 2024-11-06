import { useState } from "react";
import { Input } from "../../components/Input";
import { useDispatch, useSelector } from "react-redux";
import Loader from "../../components/Loader";
import { useHubPeak } from "../../context/HubContext";
import { userLoading } from "../user/userSlice";

const Login = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const { loading } = useSelector((state) => state.user);
  const { setModal } = useHubPeak();

  const dispatch = useDispatch();

  async function handleSubmit(e) {
    e.preventDefault();
    const loginUser = { email, password };

    try {
      dispatch(userLoading());
      const res = await fetch("http://localhost:8080/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(loginUser),
        credentials: "include",
      });

      if (!res.ok) {
        console.log("sunucu hatası");
      } else {
        setModal(false);
      }
    } catch (err) {
      console.log(err);
    }
  }

  return (
    <form onSubmit={handleSubmit}>
      <h1 className="text-3xl mt-32 mb-20">Giriş Yap</h1>
      <Input
        className="mb-6 mx-auto"
        value={email}
        type="email"
        onChange={(e) => setEmail(e.target.value)}
        placeholder="E-Posta"
        required
      />
      <Input
        className="mb-6 mx-auto"
        value={password}
        type="password"
        onChange={(e) => setPassword(e.target.value)}
        placeholder="Şifre"
        required
      />

      <button
        className="bg-blue-500 px-4 py-2 rounded-sm text-white font-medium"
        type="submit"
      >
        {loading ? (
          <p className="flex items-center gap-2">
            Giriş yapılıyor <Loader />
          </p>
        ) : (
          <p>Giriş</p>
        )}
      </button>
    </form>
  );
};

export default Login;
