import "../../styles/css/main.css";
import { useParams } from "react-router-dom";
import axios from "axios";
import { useEffect } from "react";
import ENV from "../../.env";
import { useState } from "react";
import { useToast } from "@chakra-ui/react";

const Description = () => {
  const { id } = useParams();
  const [data, setData] = useState([]);
  const [message, setMessage] = useState("");

  const token = localStorage.getItem("token");
  const idSiswa = localStorage.getItem("id_siswa");

  const toast = useToast();

  useEffect(() => {
    axios({
      method: "get",
      url: `${ENV.API_URL}/api/beasiswa?id=${id}`,
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }).then((res) => {
      setData(res.data);
    });
  }, [id, token]);

  const applyBeasiswa = () => {
    axios({
      method: "post",
      url: `${ENV.API_URL}/api/pendaftaran/create`,
      headers: {
        Authorization: `Bearer ${token}`,
      },
      data: {
        id_beasiswa: parseInt(id),
        id_siswa: parseInt(idSiswa),
      },
      mode: "cors",
      credentials: "include",
    })
      .then((res) => {
        if (res)
          toast({
            title: "Daftar Success.",
            description: "Menunggu Hasil.",
            status: "success",
            duration: 9000,
            isClosable: true,
          });
      })
      .catch((err) => {
        if (err)
          toast({
            title: "Error Daftar.",
            description: "Mohon coba beberapa saat lagi.",
            status: "error",
            duration: 9000,
            isClosable: true,
          });
        console.log(err);
      });
  };

  const isApply = () => {
    axios({
      method: "post",
      url: `${ENV.API_URL}/api/pendaftaran/update`,
      headers: {
        Authorization: `Bearer ${token}`,
      },
      data: {
        id_beasiswa: parseInt(id),
        Status: "Menungggu",
      },
      mode: "cors",
      credentials: "include",
    }).then((res) => setMessage(res.data.status));
  };

  return (
    <>
      <div className="row-flex spacing-header">
        {data?.beasiswa?.map((item) => (
          <>
            <img src={item.url_gambar} alt={item.nama} width={200} />
            <div className="column-flex spacing-content justify-center">
              <div>
                <h2 className="xl-4">{item.nama}</h2>
                <p className="lg-1">{item.program_pendidikan}</p>
              </div>
              <button className="button" onClick={applyBeasiswa && isApply}>
                {Boolean(message) ? message : "Daftar"}
              </button>
            </div>
          </>
        ))}
      </div>
      <div className="line" />
      <div className="column-flex spacing-header">
        <h4 className="lg-1">Description:</h4>
        {data?.beasiswa?.map((item) => (
          <p className="md-4 spacing-description">
            {item.deskripsi}, untuk menempuh jenjang pendidikan{" "}
            {item.jenjang_pendidikan} {item.jenis_beasiswa} program ini akan
            dimulai pada tanggal {item.tanggal_mulai} dan akan berakhir pada
            tanggal {item.tanggal_akhir}
          </p>
        ))}
      </div>
    </>
  );
};

export default Description;
