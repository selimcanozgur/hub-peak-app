import PropTypes from "prop-types";

const InputComp = ({ placeholder, value, onChange }) => {
  return (
    <input
      type="text"
      onChange={onChange}
      value={value}
      placeholder={placeholder}
      className="border focus:outline-none px-4 py-4 mb-4 w-96"
    />
  );
};

InputComp.propTypes = {
  placeholder: PropTypes.string,
  value: PropTypes.oneOfType([PropTypes.string, PropTypes.number]).isRequired,
  onChange: PropTypes.func.isRequired,
};

export default InputComp;
