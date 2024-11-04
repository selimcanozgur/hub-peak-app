import { useState } from "react";
import { useHubPeak } from "../../context/HubContext";
import { signupService } from "../../services/userAPI";
import { Input, Label } from "../../components/Input";

const Signup = () => {
  const { setSignup, setLogin } = useHubPeak();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [first_name, setFirstName] = useState("");
  const [last_name, setLastName] = useState("");
  const [birth_date, setBirthDate] = useState("");

  async function handleSubmit(e) {
    e.preventDefault();
    const newUser = { email, password, first_name, last_name, birth_date };
    signupService(newUser);
    setSignup(false);
    setLogin(true);
  }

  return (
    <form onSubmit={handleSubmit}>
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

      <Label htmlFor="password">Adınız</Label>
      <Input
        value={first_name}
        onChange={(e) => setFirstName(e.target.value)}
        type="text"
        id="first_name"
        name="first_name"
        required
      />

      <Label htmlFor="password">Soyadınız</Label>
      <Input
        value={last_name}
        onChange={(e) => setLastName(e.target.value)}
        type="text"
        id="last_name"
        name="last_name"
        required
      />

      <Label htmlFor="password">Doğum Tarihiniz</Label>
      <Input
        value={birth_date}
        onChange={(e) => setBirthDate(e.target.value)}
        type="text"
        id="birth_date"
        name="birth_date"
        required
      />
      <button>Kayıt Ol</button>
      <br />
    </form>
  );
};

export default Signup;
