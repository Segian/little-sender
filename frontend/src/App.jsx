import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import mainDashboardComponent from "./components/Home/mainDashboard/mainDashboardComponent";
import HeaderComponent from "./components/Home/header/headerComponent";
import componentViewer from "./components/componentViewer";

function App() {

  return (
    <Router>
      <HeaderComponent />
      <div className="app">
        <Routes>
          <Route path="/" element={mainDashboardComponent()} />
          <Route path="/componentViewer" element={componentViewer()} />
        </Routes>
      </div>
    </Router>
  );
}


export default App;
