import ENV from "../../.env";
import axios from "axios";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import "../../styles/css/main.css";
import { useEffect } from "react";

const Header = () => {
  const hasToken = Boolean(localStorage.getItem("token"));
  const [data, setData] = useState([]);
  const navigate = useNavigate();

  const token = localStorage.getItem("token");

  useEffect(() => {
    axios({
      method: "get",
      url: `${ENV.API_URL}/api/siswa/token`,
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }).then((res) => {
      setData(res);
    });
  }, [token]);

  return (
    <div className="header-container row-flex">
      <h4 className="title md-4-semibold">Raih Beasiswa</h4>
      <div>
        {hasToken ? (
          <p>{data?.data?.siswa?.nama}</p>
        ) : (
          <button className="button" onClick={() => navigate("/login")}>
            Login
          </button>
        )}
      </div>
    </div>
  );
};

export default Header;
