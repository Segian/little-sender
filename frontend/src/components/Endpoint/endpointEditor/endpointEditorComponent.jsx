import './endpointEditorStyle.css';
import { useState } from "react";

import Grid from "@mui/material/Grid";
import Button from '@mui/material/Button';
import Box from '@mui/material/Box';
import FormControl from '@mui/material/FormControl';
import InputLabel from '@mui/material/InputLabel';
import Select from '@mui/material/Select';
import MenuItem from '@mui/material/MenuItem';
import TextField from '@mui/material/TextField';
import EndpointPickerComponent from '../endpointPicker/endpointPicker';

function EndpointEditorComponent() {
  const [endpoint, setEndpointSelected] = useState(null);

  const handleEndpointSelect = (selectedEndpoint) => {
    setEndpointSelected(selectedEndpoint);
    console.log("Endpoint selected in editor:", selectedEndpoint);
  };
  return (
    <Box component="form" noValidate className="endpoint-editor-component">
      <Grid container spacing={2} sx={{height: '100%'}}>
        <Grid item xs={4} sx={{ width: '61.5vh' }}>
            <EndpointPickerComponent onSelect={handleEndpointSelect} />
        </Grid>

        <Grid item xs={4} sx={{ width: '86vh'}}>
          <Box sx={{ display: 'flex', flexDirection: 'row', gap: 2, height: 40, alignItems: 'center' }}>
            <Box sx={{ flex: 1.8 }}>
              <FormControl fullWidth size="small">
                <InputLabel id="method-label">Method</InputLabel>
                <Select
                  labelId="method-label"
                  id="method"
                  label="Method"
                  value={endpoint?.method || ''}
                >
                  <MenuItem value="GET">GET</MenuItem>
                  <MenuItem value="POST">POST</MenuItem>
                  <MenuItem value="PUT">PUT</MenuItem>
                  <MenuItem value="DELETE">DELETE</MenuItem>
                </Select>
              </FormControl>
            </Box>

            <TextField
              variant="outlined"
              placeholder="barra de direcciones"
              size="small"
              value={endpoint?.url || ''}
              sx={{ flex: 7 }}
            />
          </Box>

          <Box sx={{ mt: 5 }} />
        </Grid>

        <Grid item xs={2} sx={{ width: '40vh'}}>
          <Box sx={{ display: 'flex', flexDirection: 'row', gap: 2 }}>
            <Button type="button" variant="contained" color="primary">
              Guardar
            </Button>
            <Button type="submit" variant="outlined" color="secondary">
              Ejecutar
            </Button>
          </Box>
        </Grid>
      </Grid>
    </Box>
  );
}

export default EndpointEditorComponent;