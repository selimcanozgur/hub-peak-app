import { createContext, useContext, useState } from "react";
import PropTypes from "prop-types";

const Context = createContext();

const ContextProvider = function ({ children }) {
  const [modal, setModal] = useState(false);
  const [signup, setSignup] = useState(false);
  const [login, setLogin] = useState(false);

  return (
    <Context.Provider
      value={{ modal, setModal, signup, setSignup, login, setLogin }}
    >
      {children}
    </Context.Provider>
  );
};

ContextProvider.propTypes = {
  children: PropTypes.node.isRequired, // children prop'unu tanımlayın
};

const useHubPeak = function () {
  const contextHub = useContext(Context);
  if (!contextHub) throw Error("Use hub tanımsız.");
  return contextHub;
};

export { ContextProvider, useHubPeak };
