// frontend/src/config/api.js
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8000';

export const API_CONFIG = {
  baseURL: API_BASE_URL,
  endpoints: {
    health: '/health',
    users: '/api/v1/users',
    auth: '/api/v1/auth',
    announcements: '/api/v1/announcements',
    publicDocs: '/api/v1/public-docs',
    jobPositions: '/api/v1/job-positions',
  }
};

export default API_CONFIG;