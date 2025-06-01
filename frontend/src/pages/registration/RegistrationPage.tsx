import React, { useState } from "react";
import { registerUser } from "../../shared/api";
import { ToastContainer, toast } from "react-toastify";
import { useRegister } from "../../features/registerUser/useRegister";

export default function RegistrationPage() {
  const {
    handleRegister,
    name,
    setName,
    email,
    setEmail,
    password,
    setPassword,
    message,
  } = useRegister();
  return (
    <form onSubmit={handleRegister}>
      <input
        placeholder="Name"
        value={name}
        onChange={(e) => setName(e.target.value)}
        required
      />
      <input
        placeholder="Email"
        type="email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        required
      />
      <input
        placeholder="Password"
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        required
      />
      <button type="submit">Register</button>
      {message && <p>{message}</p>}
      <ToastContainer />
    </form>
  );
}
