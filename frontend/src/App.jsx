import { useEffect, useState } from "react";
import ShowDetails from "./components/showDetails/showDetails";

function App() {
  const [endpoints, setEndpoints] = useState([]);
  const [selectedEndpoint, setSelectedEndpoint] = useState(null);

  useEffect(() => {
    fetch("http://localhost:8080/endpoint")
      .then(res => {
        console.log(res);
        return res.json();
      })
      .then(data => setEndpoints(data))
      .catch(err => console.error(err));
  }, []);

  return (
    <div>
      <h1>Little Sender UI</h1>
      <h2>Endpoints</h2>
      <ul>
        {endpoints.map((ep, i) => (
          <li key={i}>
            <strong>{ep.name}</strong> → {ep.url}
            <button onClick={() => setSelectedEndpoint(ep.id)}>Show Details</button>
          </li>
        ))}
      </ul>
      {selectedEndpoint && <ShowDetails id={selectedEndpoint} />}
    </div>
  );
}

export default App;
