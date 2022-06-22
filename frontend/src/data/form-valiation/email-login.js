import * as Yup from "yup";

const emailLoginValidation = Yup.object().shape({
  email: Yup.string().required("Email Required").email("Invalid Email"),
  password: Yup.string().required("Password Required"),
});

export default emailLoginValidation;
