import { BrowserRouter, Routes, Route } from "react-router-dom";
import { LoginPage } from "../pages/login/LoginPage";
import RegistrationPage from "../pages/registration/RegistrationPage";
import React from "react";
import "./styles/reset.css";
import HomePage from "../pages/home/HomePage";
import { Header } from "../widgets/Header";

/*
Главная-рекомендации и чат 
Меню-меню сервис
Профиль,регистрация,логин,-юзер сервисе
*/

export default function AppRoutes() {
  return (
    <BrowserRouter>
      <Header />

      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/register" element={<RegistrationPage />} />

        {/* <Route path="*" element={<ErrorPage />} /> */}
      </Routes>

      {/* <Footer /> */}
    </BrowserRouter>
  );
}
