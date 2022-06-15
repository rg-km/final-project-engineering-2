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
    <div className="row-flex">
      <form>
        <p className="lg-1">{label}</p>
        <input
          className="input-base-style md-4"
          name={r?.name}
          type={inputType}
          onBlur={r?.onBlur}
          onChange={r?.onChange}
          value={r?.value}
          ref={r?.ref}
          placeholder={placeholder}
        />
      </form>
    </div>
  );
};

export default TextInput;
