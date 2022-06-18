import React from "react";
import "../../../styles/css/main.css";

const TextInput = ({
  label,
  name,
  register,
  inputType,
  placeholder,
  errorText,
}) => {
  const r = register(name);
  return (
    <div className="row-flex spacing">
      <form>
        <p className="lg-1">{label}</p>
        <input
          className="input-base-style md-4 spacing"
          name={r?.name}
          type={inputType}
          onBlur={r?.onBlur}
          onChange={r?.onChange}
          value={r?.value}
          // ref={r?.ref}
          placeholder={placeholder}
        />
        <p className="md-4 error-text">{Boolean(errorText) && errorText}</p>
      </form>
    </div>
  );
};

export default TextInput;
