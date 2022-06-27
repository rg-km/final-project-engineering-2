import { BrowserRouter } from "react-router-dom";
import { Route, Routes } from "react-router-dom";
import Header from "./components/Header";
import Login from "./screen/authentication/login";
import Register from "./screen/authentication/Register";
import Footer from "./components/Footer/index";
import ScholarshipsList from "./components/card/ScholarshipList";
import Jumbotron from "./components/Jumbotron/Jumbotron";
import NoMatchRoute from "./components/card/NoMatchRoute";

const App = () => {
  return (
    <BrowserRouter>
      <div className="App">
        <Header />
        <Routes>
          <Route
            path="/"
            element={
              <>
                <Jumbotron />
                <ScholarshipsList />
              </>
            }
          />
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route path="*" element={<NoMatchRoute />} />
        </Routes>
        <Footer />
      </div>
    </BrowserRouter>
  );
};

export default App;
