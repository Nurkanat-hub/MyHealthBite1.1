import React, { useState } from "react";
import { loginUser } from "../../shared/api";
import { ToastContainer, toast } from "react-toastify";

export const LoginPage = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleLogin = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const res = await loginUser(email, password);
      toast.success("Login successful!");
    } catch (err) {
      if (err instanceof Error) {
        toast.error("Login failed: " + err.message);
      } else {
        toast.error("Login failed: Unknown error");
      }
    }

    return (
      <form onSubmit={handleLogin}>
        <input
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <input
          placeholder="Password"
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <button type="submit">Login</button>
        <ToastContainer />
      </form>
    );
  };
};

export default LoginPage;
