import { BrowserRouter } from "react-router-dom";
import { Route, Routes } from "react-router-dom";
import Header from "./components/Header";
import Login from "./screen/authentication/login";
import Register from "./screen/authentication/Register";
import Footer from "./components/Footer/index";
import ScholarshipsList from "./components/card/ScholarshipList";
import Jumbotron from "./components/Jumbotron/Jumbotron";
<<<<<<< HEAD
import NoMatchRoute from "./components/card/NoMatchRoute";
=======
import Description from "./screen/Description";
>>>>>>> 1f036914647e4dcbee24293b2d2e092bc4f57c11

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
<<<<<<< HEAD
          <Route path="*" element={<NoMatchRoute />} />
=======
          <Route path="/scholarship/:id" element={<Description />} />
>>>>>>> 1f036914647e4dcbee24293b2d2e092bc4f57c11
        </Routes>
        <Footer />
      </div>
    </BrowserRouter>
  );
};

export default App;
