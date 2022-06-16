import Login from "./screen/authentication/login";
import { BrowserRouter as Router } from "react-router-dom";

const App = () => {
  return (
    <Router>
      <div className="App">
        <Login />
      </div>
    </Router>
  );
};

export default App;
