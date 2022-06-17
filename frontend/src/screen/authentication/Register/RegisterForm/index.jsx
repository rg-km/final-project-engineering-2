import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import Form from "../../../../components/data-entry/Form";
import "../../../../styles/css/main.css";

const RegisterForm = () => {
  const navigate = useNavigate();

  const {
    register,
    control,
    formState: { errors },
    handleSubmit,
  } = useForm();

  const submit = async (value) => {};

  const navigateToLogin = () => navigate("/login");

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
              label: "Full name",
              name: "fullName",
              placeholder: "Enter your Full Name",
            },
            {
              label: "Password",
              name: "password",
              placeholder: "Enter your password",
              inputType: "password",
            },
            {
              label: "Confirm Password",
              name: "confirmpassword",
              placeholder: "Enter your password",
              inputType: "password",
            },
          ]}
          control={control}
          register={register}
          errors={errors}
        />
        <div className="column-flex container">
          <div className="row-flex spacing-text-button">
            <p className="md-4">Sudah punya akun? </p>
            <p className="md-4-semibold text" onClick={navigateToLogin}>
              Masuk Disini
            </p>
          </div>
          <button className="button" onClick={handleSubmit(submit)}>
            Submit
          </button>
        </div>
      </form>
    </div>
  );
};

export default RegisterForm;
