import { useState } from "react";
import styled from "styled-components";
import { signup } from "../../services/userAPI";

const Input = styled.input`
  background-color: #efefef;
  border-radius: 2px;
  padding: 8px 10px;
  outline-color: #93bcef;
`;

const Label = styled.label`
  font-size: 14px;
  margin-bottom: 4px;
  display: block;
`;

const Signup = () => {
  const [email, setEmail] = useState();
  const [password, setPassword] = useState();

  async function handleSubmit(e) {
    e.preventDefault();
    const newUser = { email, password };
    signup(newUser);
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
        />
      </div>

      <Label htmlFor="password">Şifre</Label>
      <Input
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        type="password"
        id="password"
        name="password"
      />
      <button>Kayıt Ol</button>
    </form>
  );
};

export default Signup;
