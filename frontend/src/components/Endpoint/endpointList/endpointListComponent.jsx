import { useState, useEffect } from "react";
import { getEndpoints } from "../../../services/endpointServices";

import Button from "@mui/material/Button";
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';

function EndpointRow({ endpoint }) {
  return (
    <TableRow key={endpoint.id}>
      <TableCell sx={{color: 'white', borderBottom: '2px solid #1c212dff'}}>{endpoint.name}</TableCell>
      <TableCell align="right" sx={{color: 'white', borderBottom: '2px solid #1c212dff'}}>{endpoint.url}</TableCell>
      <TableCell align="right" sx={{color: 'white', borderBottom: '2px solid #1c212dff'}}>{endpoint.method}</TableCell>
      <TableCell align="center" sx={{borderBottom: '2px solid #1c212dff'}}>
        <Button variant="contained" color="primary" style={{marginRight: '8px'}}>Ejecutar</Button>
        <Button variant="contained" color="warning">Eliminar</Button>
      </TableCell>
    </TableRow>
  );
}

function EndpointListComponent() {
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
              <EndpointRow key={endpoint.id} endpoint={endpoint} />
            ))}
          </TableBody>
        </Table>
      </TableContainer>
  );
}

export default EndpointListComponent;