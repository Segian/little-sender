import './endpointEditorStyle.css';
import Grid from "@mui/material/Grid";
import Paper from '@mui/material/Paper';

function EndpointEditorComponent() {
  return (
    <div className="endpoint-editor-component">
      <Grid container spacing={2} sx={{ flex: 1 }}>
        <Grid item size={4} sx={{ height: '100%' }}>
          <Paper sx={{ height: '100%' }}>
            aqui va la lista de endpoints
          </Paper>
        </Grid>
        <Grid item xs={1} sx={{ height: '100%' }}>
          <Paper>
            aqui va el selector de tipo de endpoint
          </Paper>
        </Grid>
        
      </Grid>
    </div>
  );
}

export default EndpointEditorComponent;