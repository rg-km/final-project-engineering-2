import {
  Alert,
  AlertDescription,
  AlertIcon,
  AlertTitle,
} from "@chakra-ui/react";

const AlertModal = ({ title, errorMsg }) => (
  <Alert status="error">
    <AlertIcon />
    <AlertTitle>{title}</AlertTitle>
    <AlertDescription>{errorMsg}</AlertDescription>
  </Alert>
);

export default AlertModal;
