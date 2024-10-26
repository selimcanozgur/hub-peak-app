import { useRouteError } from "react-router-dom";

const Error = () => {
  const error = useRouteError();
  console.log(error);
  return (
    <div>
      <h1>Bir şeyler ters gitti 😢</h1>
      <p>{error.data || error.message}</p>
    </div>
  );
};

export default Error;
