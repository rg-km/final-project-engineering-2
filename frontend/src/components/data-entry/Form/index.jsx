import TextInput from "../TextInput";

const getComponent = (type) => {
  switch (type) {
    case "text-input":
      return TextInput;
    default:
      return TextInput;
  }
};

const Form = ({ forms = [], register, control, errors }) => {
  const Forms = forms.map((form) => {
    const Comp = getComponent(form.type);
    return (
      <Comp
        key={form.name}
        label={form.label}
        register={register}
        control={control}
        errors={errors}
        inputType={form.inputType}
        {...form}
      />
    );
  });
  return Forms;
};

export default Form;
