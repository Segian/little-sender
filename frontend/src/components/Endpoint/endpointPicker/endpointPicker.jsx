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

function EndpointRow({ endpoint, isSelected, onSelect }) {
  return (
    <TableRow 
        hover
        selected={isSelected}
        onClick={() => onSelect?.(endpoint)}
        sx={{
        cursor: 'pointer',
        '&.Mui-selected td': { backgroundColor: '#2b3a55' },
        '&.Mui-selected:hover td': { backgroundColor: '#2b3a55' },
        '&:hover td': { backgroundColor: '#263040' },
      }}
    >
      <TableCell sx={{color: 'white', borderBottom: '2px solid #1c212dff'}}>{endpoint.name}</TableCell>
      <TableCell align="right" sx={{color: 'white', borderBottom: '2px solid #1c212dff'}}>{endpoint.url}</TableCell>
      <TableCell align="right" sx={{color: 'white', borderBottom: '2px solid #1c212dff'}}>{endpoint.method}</TableCell>
    </TableRow>
  );
}

function EndpointPickerComponent({ onSelect }) {
  const [endpoints, setEndpoints] = useState([]);
  const [selectedId, setSelectedId] = useState(null);

  const handleSelect = (endpoint) => {
    setSelectedId(endpoint.id);
    onSelect?.(endpoint);
    console.log("Selected endpoint:", endpoint); 
  }

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
    <TableContainer component={Paper} sx={{ maxHeight: 500, backgroundColor: '#213547', boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)' }}>
      <Table sx={{ backgroundColor: '#213547'}} aria-label="Endpoints disponibles" stickyHeader>
        <TableHead
            sx={{ '& .MuiTableCell-stickyHeader': { backgroundColor: '#1c212d', borderBottom: '2px solid #ffffffff' } }}
        >
          <TableRow sx={{borderBottom: '2px solid #ffffffff'}}>
            <TableCell sx={{color: 'white'}}>Name</TableCell>
              <TableCell align="right" sx={{color: 'white'}}>URL</TableCell>
              <TableCell align="right" sx={{color: 'white'}}>Method</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {endpoints.map((endpoint) => (
              <EndpointRow 
              key={endpoint.id} 
              endpoint={endpoint} 
              isSelected={endpoint.id === selectedId}
              onSelect={handleSelect}
              />
            ))}
          </TableBody>
        </Table>
      </TableContainer>
  );
}

export default EndpointPickerComponent;