import axios from 'axios';

export const getEndpoints = async () => {
  try {
    const response = await axios.get('/api/endpoint');
    return response.data;
  } catch (error) {
    console.error('Error fetching endpoints:', error);
    throw error;
  }
};

export const createEndpoint = async (endpointData) => {
  try {
    const response = await axios.post('/api/endpoint', endpointData);
    return response.data;
  } catch (error) {
    console.error('Error creating endpoint:', error);
    throw error;
  }
};

export const updateEndpoint = async (endpointId, endpointData) => {
  try {
    const response = await axios.put(`/api/endpoint/${endpointId}`, endpointData);
    return response.data;
  } catch (error) {
    console.error('Error updating endpoint:', error);
    throw error;
  }
};

export const deleteEndpoint = async (endpointId) => {
  try {
    const response = await axios.delete(`/endpoint/${endpointId}`);
    return response.data;
  } catch (error) {
    console.error('Error deleting endpoint:', error);
    throw error;
  }
};
