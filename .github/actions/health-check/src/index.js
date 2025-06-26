const core = require('@actions/core');
const http = require('http');
const https = require('https');
const { URL } = require('url');

/**
 * Performs a health check on the specified URL
 * @param {string} url - The URL to check
 * @param {number} timeout - Timeout in milliseconds
 * @param {number} expectedStatus - Expected HTTP status code
 * @returns {Promise<Object>} Health check results
 */
async function performHealthCheck(url, timeout, expectedStatus) {
  return new Promise((resolve, reject) => {
    const startTime = Date.now();
    const parsedUrl = new URL(url);
    const client = parsedUrl.protocol === 'https:' ? https : http;
    
    const options = {
      hostname: parsedUrl.hostname,
      port: parsedUrl.port,
      path: parsedUrl.pathname + parsedUrl.search,
      method: 'GET',
      timeout: timeout,
      headers: {
        'User-Agent': 'GitHub-Actions-Health-Check/1.0'
      }
    };
    
    const req = client.request(options, (res) => {
      const responseTime = Date.now() - startTime;
      let data = '';
      
      res.on('data', (chunk) => {
        data += chunk;
      });
      
      res.on('end', () => {
        const result = {
          url,
          statusCode: res.statusCode,
          responseTime,
          headers: res.headers,
          bodySize: data.length,
          timestamp: new Date().toISOString()
        };
        
        if (res.statusCode === expectedStatus) {
          result.status = 'healthy';
          result.message = `Health check passed (${res.statusCode})`;
        } else {
          result.status = 'unhealthy';
          result.message = `Health check failed. Expected ${expectedStatus}, got ${res.statusCode}`;
        }
        
        resolve(result);
      });
    });
    
    req.on('error', (error) => {
      const responseTime = Date.now() - startTime;
      resolve({
        url,
        status: 'unhealthy',
        message: `Health check failed: ${error.message}`,
        error: error.message,
        responseTime,
        timestamp: new Date().toISOString()
      });
    });
    
    req.on('timeout', () => {
      req.destroy();
      const responseTime = Date.now() - startTime;
      resolve({
        url,
        status: 'unhealthy',
        message: `Health check timed out after ${timeout}ms`,
        responseTime,
        timestamp: new Date().toISOString()
      });
    });
    
    req.end();
  });
}

/**
 * Main action function
 */
async function run() {
  try {
    // Get inputs
    const url = core.getInput('url', { required: true });
    const timeout = parseInt(core.getInput('timeout') || '30') * 1000; // Convert to milliseconds
    const retryCount = parseInt(core.getInput('retry-count') || '3');
    const expectedStatus = parseInt(core.getInput('expected-status') || '200');
    
    core.info(`Starting health check for ${url}`);
    core.info(`Configuration: timeout=${timeout}ms, retries=${retryCount}, expected-status=${expectedStatus}`);
    
    let lastResult = null;
    let attempt = 0;
    
    // Retry logic
    while (attempt < retryCount) {
      attempt++;
      core.info(`Health check attempt ${attempt}/${retryCount}`);
      
      lastResult = await performHealthCheck(url, timeout, expectedStatus);
      
      if (lastResult.status === 'healthy') {
        core.info(`✅ Health check passed on attempt ${attempt}`);
        break;
      } else {
        core.warning(`❌ Health check failed on attempt ${attempt}: ${lastResult.message}`);
        
        if (attempt < retryCount) {
          const delay = Math.min(1000 * Math.pow(2, attempt - 1), 10000); // Exponential backoff, max 10s
          core.info(`Waiting ${delay}ms before retry...`);
          await new Promise(resolve => setTimeout(resolve, delay));
        }
      }
    }
    
    // Set outputs
    core.setOutput('status', lastResult.status);
    core.setOutput('response-time', lastResult.responseTime.toString());
    core.setOutput('details', JSON.stringify(lastResult, null, 2));
    
    // Create summary
    const summary = core.summary
      .addHeading('🏥 Health Check Results')
      .addTable([
        [
          { data: 'URL', header: true },
          { data: 'Status', header: true },
          { data: 'Response Time', header: true },
          { data: 'Status Code', header: true }
        ],
        [
          url,
          lastResult.status === 'healthy' ? '✅ Healthy' : '❌ Unhealthy',
          `${lastResult.responseTime}ms`,
          lastResult.statusCode ? lastResult.statusCode.toString() : 'N/A'
        ]
      ]);
    
    if (lastResult.status === 'unhealthy') {
      summary.addRaw(`\n**Error Details:**\n\`\`\`\n${lastResult.message}\n\`\`\``);
    }
    
    await summary.write();
    
    // Log final result
    if (lastResult.status === 'healthy') {
      core.info(`🎉 Health check completed successfully after ${attempt} attempt(s)`);
    } else {
      core.setFailed(`Health check failed after ${attempt} attempt(s): ${lastResult.message}`);
    }
    
  } catch (error) {
    core.setFailed(`Action failed: ${error.message}`);
    core.debug(error.stack);
  }
}

// Run the action
if (require.main === module) {
  run();
}

module.exports = { run, performHealthCheck };
