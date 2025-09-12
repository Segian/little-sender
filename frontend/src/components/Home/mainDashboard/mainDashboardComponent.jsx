import EndpointListComponent from "../../Endpoint/endpointList/endpointListComponent";

import "./mainDashboardStyle.css";


function MainDashboardComponent() {

  return (
    <div>
      <h1 className="titulo">Lista de endpoints</h1>
      <div className="table-container">
        <EndpointListComponent/>
      </div>
    </div>
  );
}

export default MainDashboardComponent;