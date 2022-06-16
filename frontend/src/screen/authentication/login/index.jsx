import "../../../styles/css/main.css";
import LoginForm from "./LoginForm";

const Login = () => {
  return (
    <div className="row-flex container">
      <div className="img-left" />
      <div className="column-flex content-container spacing-form-right">
        <LoginForm />
      </div>
    </div>
  );
};

export default Login;
