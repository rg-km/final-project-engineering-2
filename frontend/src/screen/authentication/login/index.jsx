import { useForm } from "react-hook-form";
import Form from "../../../components/data-entry/Form";

const Login = () => {
  const {
    register,
    control,
    formState: { errors },
    handleSubmit,
  } = useForm();

  return (
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
          type: "text-input",
        },
      ]}
      control={control}
      register={register}
    />
  );
};

export default Login;
