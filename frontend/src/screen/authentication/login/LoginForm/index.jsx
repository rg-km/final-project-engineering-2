import axios from "axios";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import ENV from "../../../../.env";
import Form from "../../../../components/data-entry/Form";
import AlertModal from "../../../../components/AlertModal";
import useAuth from "../../../../hooks/useAuth";
import "../../../../styles/css/main.css";
import { useState } from "react";
import { Spinner } from "@chakra-ui/react";

const LoginForm = () => {
  const navigate = useNavigate();
  const [error, setError] = useState(false);
  const [loading, setLoading] = useState(false);

  const {
    register,
    control,
    formState: { errors },
    handleSubmit,
  } = useForm({});

  const setAuthToken = useAuth((state) => state.setAuthToken);

  const submit = (value) => {
    setLoading(true);
    axios({
      method: "post",
      url: `${ENV.API_URL}/api/login`,
      data: value,
      mode: "cors",
      credentials: "include",
    })
      .then((res) => {
        setError(false);
        setAuthToken(res.data.token);
        navigate("/");
      })
      .catch((err) => {
        if (err) setError(true);
      });
    setLoading(false);
  };

  const navigateToRegister = () => navigate("/register");

  return (
    <div className="column-flex container">
      <h3 className="xl-4">Login</h3>
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
              label: "Password",
              name: "password",
              placeholder: "Enter your password",
              inputType: "password",
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
            <p className="md-4">Belum punya akun? </p>
            <p className="md-4-semibold text" onClick={navigateToRegister}>
              Daftar Disini
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

export default LoginForm;
