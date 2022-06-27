import React from "react";
import { useNavigate } from "react-router-dom";
import "./NoMatchRoute.css";

function NoMatchRoute() {
  const navigate = useNavigate();
  return (
    <div className="no-match-route__container">
      <div className="no-match-route__image">
        <img
          src="images/no-match-route-illustration.jpg"
          alt="page not found"
        />
      </div>
      <div className="no-match-route__body">
        <div className="no-match-route__message">
          <p>Kemungkinan halaman telah dihapus, atau Anda salah menulis URL</p>
        </div>
        <button
          onClick={() => navigate("/")}
          className={"no-match-route__button"}
        >
          Kembali Ke halaman awal
        </button>
      </div>
    </div>
  );
}

export default NoMatchRoute;
