import "../../../styles/css/main.css";
import RegisterForm from "./RegisterForm";

const Register = () => {
  return (
    <div className="row-flex container bottom-spacing">
      <div className="img-left" />
      <div className="column-flex content-container spacing-form-right">
        <RegisterForm />
      </div>
    </div>
  );
};

export default Register;
