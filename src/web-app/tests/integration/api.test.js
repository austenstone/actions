const request = require('supertest');
const app = require('../../src/index');

describe('Integration Tests', () => {
  describe('API Flow', () => {
    it('should handle complete API workflow', async () => {
      // Check health first
      await request(app)
        .get('/health')
        .expect(200);

      // Get app status
      const statusResponse = await request(app)
        .get('/api/status')
        .expect(200);

      expect(statusResponse.body.status).toBe('running');

      // Get demo information
      const demoResponse = await request(app)
        .get('/api/demo')
        .expect(200);

      expect(demoResponse.body.features).toContain('Automated testing');
    });

    it('should maintain consistent response format', async () => {
      const endpoints = ['/', '/health', '/api/status', '/api/demo'];
      
      for (const endpoint of endpoints) {
        const response = await request(app)
          .get(endpoint)
          .expect(200);

        // All endpoints should return JSON
        expect(response.headers['content-type']).toMatch(/json/);
        
        // All endpoints should have some form of timestamp or status
        expect(
          response.body.timestamp || 
          response.body.status || 
          response.body.uptime !== undefined
        ).toBeTruthy();
      }
    });
  });

  describe('Error Handling', () => {
    it('should handle invalid JSON gracefully', async () => {
      // Since we don't have POST endpoints, we'll test 404 handling
      const response = await request(app)
        .post('/invalid-endpoint')
        .send({ invalid: 'data' })
        .expect(404);

      expect(response.body).toHaveProperty('error');
    });
  });

  describe('Security Headers', () => {
    it('should include security headers', async () => {
      const response = await request(app)
        .get('/')
        .expect(200);

      // Check for helmet security headers
      expect(response.headers).toHaveProperty('x-frame-options');
      expect(response.headers).toHaveProperty('x-content-type-options');
    });
  });
});
