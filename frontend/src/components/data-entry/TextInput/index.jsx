import React from "react";
import "../../../styles/core/styles.css";
import "./styles.css";

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
        <input
          className="input-base-style"
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
