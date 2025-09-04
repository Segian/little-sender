import "./headerStyles.css";
import { Link } from "react-router-dom";


function HeaderComponent() {
  return (
    <header>
        <nav>
            <div className="logo">
              <Link className="nav-logo" to="/">Little sender</Link>
              <div className="nav-links">
                <Link className="NavBton" to="/">Inicio</Link>
                <Link className="NavBton" to="/create">Crear y/o ejecutar</Link>
              </div>
            </div>
        </nav>
    </header>
  );
}

export default HeaderComponent;
