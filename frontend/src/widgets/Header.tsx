import React from "react";
import logo from "../shared/assets/bb94c9f31e1687067757d6a34f6e56a705b7a796.png";
import phone from "../shared/assets/material-symbols_phone-in-talk-rounded.svg";
import cart from "../shared/assets/shopping_cart_24dp_000000_FILL0_wght400_GRAD0_opsz24.png";
import { Link } from "react-router-dom";
import { useState } from "react";

export const Header = () => {
  const [isAuthenticated, setIsAuthenticated] = useState(
    !!localStorage.getItem("token")
  );

  return (
    <header className="px-12">
      <div className="flex flex-row items-center p-5 justify-between">
        <img src={logo} className="w-32 h-10" />
        <nav className="flex gap-11">
          <Link to="/menu">Menu</Link>
          <Link to="/menu">Menu</Link>
          <Link to="/menu">Menu</Link>
          <Link to="/menu">Menu</Link>
        </nav>
        <div className="flex flex-row">
          <img src={phone} alt="" />
          <span>88005553535</span>
        </div>
        <Link to="/cart" className="flex flex-row">
          <img src={cart} alt="" />
          <span>Cart</span>
        </Link>
        {isAuthenticated ? (
          <Link to="/profile">
            <button className="bg-[#FFA800] px-7 py-2 rounded-3xl text-white">
              Profile
            </button>
          </Link>
        ) : (
          <Link to="/login">
            <button className="bg-[#FFA800] px-7 py-2 rounded-3xl text-white">
              Sign In
            </button>
          </Link>
        )}
      </div>
    </header>
  );
};
