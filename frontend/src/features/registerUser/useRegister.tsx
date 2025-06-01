import { registerUser } from "../../shared/api";
import { useState } from "react";
import { toast } from "react-toastify";

export function useRegister() {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");

  const handleRegister = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await registerUser(name, email, password);
      toast.success("Account created successfully!");
      setMessage("");
    } catch (err: any) {
      toast.error("Registration failed: " + err.message);
      setMessage(err.message);
    }
  };

  return {
    name,
    setName,
    email,
    setEmail,
    password,
    setPassword,
    handleRegister,
    message,
  };
}
