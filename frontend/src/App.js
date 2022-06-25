import { BrowserRouter } from "react-router-dom";
import { Route, Routes } from "react-router-dom";
import Header from "./components/Header";
import Login from "./screen/authentication/login";
import Register from "./screen/authentication/Register";
import Footer from "./components/Footer/index";

const App = () => {
  return (
    <BrowserRouter>
      <div className="App">
        <Header />
        <Routes>
          <Route path="/" element={<h1>test</h1>} />
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
        </Routes>
        <Footer />
      </div>
    </BrowserRouter>
  );
};

export default App;
