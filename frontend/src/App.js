import { Route, Routes } from "react-router-dom";
import Login from "./screen/authentication/login";
import Register from "./screen/authentication/Register";

const App = () => {
  return (
    <div className="App">
      <Routes>
        <Route path="/login" component={<Login />} />
        <Route path="/register" component={<Register />} />
      </Routes>
    </div>
  );
};

export default App;
