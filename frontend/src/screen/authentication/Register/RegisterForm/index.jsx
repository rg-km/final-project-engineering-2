import { Spinner } from "@chakra-ui/react";
import axios from "axios";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import ENV from "../../../../.env";
import AlertModal from "../../../../components/AlertModal";
import Form from "../../../../components/data-entry/Form";
import "../../../../styles/css/main.css";

const RegisterForm = () => {
  const navigate = useNavigate();
  const [error, setError] = useState(false);
  const [loading, setLoading] = useState(false);

  const {
    register,
    control,
    formState: { errors },
    handleSubmit,
  } = useForm();

  const submit = (value) => {
    setLoading(true);
    axios({
      method: "post",
      url: `${ENV.API_URL}/api/register`,
      data: value,
      mode: "cors",
      credentials: "include",
    })
      .then((res) => {
        setError(false);
        if (res) {
          navigate("/login");
        }
      })
      .catch((err) => {
        if (err) setError(true);
      });
  };

  const navigateToLogin = () => navigate("/login");

  return (
    <div className="column-flex container">
      <h3 className="xl-4">Register</h3>
      <form>
        <Form
          forms={[
            {
              label: "Email",
              name: "email",
              placeholder: "Enter your email",
              type: "text-input",
            },
            {
              label: "Name",
              name: "nama",
              placeholder: "Enter your Full Name",
            },
            {
              label: "Password",
              name: "password",
              placeholder: "Enter your password",
              inputType: "password",
            },
            {
              label: "Education",
              name: "jenjang_pendidikan",
              placeholder: "Enter your Education Level",
            },
            {
              label: "NIK",
              name: "nik",
              placeholder: "Enter your NIK",
            },
            {
              label: "City",
              name: "tempat_lahir",
              placeholder: "Enter your City of Birth",
            },
            {
              label: "Date of Birth",
              name: "tanggal_lahir",
              placeholder: "YYYY-MM-DD",
            },
          ]}
          control={control}
          register={register}
          errors={errors}
        />
        {error && (
          <AlertModal title="Error Login" errorMsg="Email or Password Wrong" />
        )}
        <div className="column-flex container">
          <div className="row-flex spacing-text-button">
            <p className="md-4">Sudah punya akun? </p>
            <p className="md-4-semibold text" onClick={navigateToLogin}>
              Masuk Disini
            </p>
          </div>
          <button className="button" onClick={handleSubmit(submit)}>
            {loading ? <Spinner /> : "Submit"}
          </button>
        </div>
      </form>
    </div>
  );
};

export default RegisterForm;
