import React from "react";
import axios from "axios";
import ENV from "../../../.env";
import ScholarshipsItem from "../ScholarshipItem/index";
import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import "./ScholarshipList.css";

function ScholarshipsList() {
  const [data, setData] = useState([]);
  const navigate = useNavigate();

  const token = localStorage.getItem("token");

  useEffect(() => {
    axios({
      method: "get",
      url: `${ENV.API_URL}/api/beasiswa/all`,
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }).then((res) => {
      setData(res.data);
    });
  }, [token]);

  return (
    <div className="scholarships-list">
      {data?.beasiswa?.map(
        ({
          id,
          nama: name,
          jenjang_pendidikan: faculty,
          jenis_beasiswa: level,
          lama_program: completionTime,
          url_gambar: imageUrl,
        }) => (
          <ScholarshipsItem
            onClickScholarship={() => navigate(`/scholarship/${id}`)}
            key={id}
            id={id}
            name={name}
            faculty={faculty}
            level={level}
            completionTime={completionTime}
            imageUrl={imageUrl}
          />
        )
      )}
    </div>
  );
}

export default ScholarshipsList;
