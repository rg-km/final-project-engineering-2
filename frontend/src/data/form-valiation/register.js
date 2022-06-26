import * as Yup from "yup";

const registerValidation = Yup.object().shape({
  email: Yup.string().required("Email Required").email("Invalid Email"),
  nama: Yup.string().required("Name Required"),
  password: Yup.string().required("Password Required"),
  jenjang_pendidikan: Yup.string().required("Education Required"),
  nik: Yup.string().required("NIK Required"),
  tempat_lahir: Yup.string().required("City of Birth required"),
  tanggal_lahir: Yup.string().required("Date of Birth required"),
  kota_domisili: Yup.string().required("City required"),
});

export default registerValidation;
