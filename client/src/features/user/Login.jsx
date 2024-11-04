import { useState } from "react";
import { Input, Label } from "../../components/Input";
import { useHubPeak } from "../../context/HubContext";

const Login = () => {
  const { userData } = useHubPeak();
  console.log(userData);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  async function handleSubmit(e) {
    e.preventDefault();
    const newUser = { email, password };
    try {
      const res = await fetch("http://localhost:8080/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(newUser),
        credentials: "include",
      });

      if (!res.ok) {
        throw new Error(`Error: ${res.status} ${res.statusText}`);
      }

      const data = await res.json();
      return data;
    } catch (err) {
      console.error("Signup failed:", err);
      throw err;
    }
  }

  return (
    <form onSubmit={handleSubmit}>
      <h1>Giriş</h1>
      <div>
        <Label htmlFor="email">E-Posta</Label>
        <Input
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          type="email"
          id="email"
          name="email"
          required
        />
      </div>

      <Label htmlFor="password">Şifre</Label>
      <Input
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        type="password"
        id="password"
        name="password"
        required
      />
      <button>Giriş yap</button>
    </form>
  );
};

export default Login;
