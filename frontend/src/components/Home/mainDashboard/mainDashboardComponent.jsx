import { useState, useEffect } from "react";
import { getEndpoints } from "../../../services/endpointServices";
import "./mainDashboardStyle.css";

import Button from "@mui/material/Button";
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';

function MainDashboardComponent() {
  const [endpoints, setEndpoints] = useState([]);

  useEffect(() => {
    const fetchEndpoints = async () => {
      try {
        const data = await getEndpoints();
        setEndpoints(data);
      } catch (error) {
        console.error("Error fetching endpoints:", error);
      }
    };

    fetchEndpoints();
  }, []);

  return (
    <div>
      <h1 className="titulo">Lista de endpoints</h1>
      <div className="table-container">
      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 650, backgroundColor: '#213547'}} aria-label="Endpoints disponibles">
          <TableHead>
            <TableRow sx={{borderBottom: '2px solid #ffffffff'}}>
              <TableCell sx={{color: 'white'}}>Name</TableCell>
              <TableCell align="right" sx={{color: 'white'}}>URL</TableCell>
              <TableCell align="right" sx={{color: 'white'}}>Method</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {endpoints.map((endpoint) => (
              <TableRow key={endpoint.id}>
                <TableCell sx={{color: 'white', borderBottom: '2px solid #1c212dff'}}>{endpoint.name}</TableCell>
                <TableCell align="right" sx={{color: 'white', borderBottom: '2px solid #1c212dff'}}>{endpoint.url}</TableCell>
                <TableCell align="right" sx={{color: 'white', borderBottom: '2px solid #1c212dff'}}>{endpoint.method}</TableCell>
                <TableCell align="center" sx={{borderBottom: '2px solid #1c212dff'}}>
                  <Button variant="contained" color="primary" style={{marginRight: '8px'}}>Ejecutar</Button>
                  <Button variant="contained" color="warning">Eliminar</Button>
                  </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
      </div>
    </div>
  );
}

export default MainDashboardComponent;